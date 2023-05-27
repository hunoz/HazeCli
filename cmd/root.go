package cmd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/fatih/color"
	"github.com/hunoz/haze/cmd/update"
	"github.com/hunoz/haze/request"
	"github.com/hunoz/spark/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getTokenFromAllOptions() string {
	sparkConfig, configErr := config.GetCognitoConfig()
	tokenArg := viper.GetString(FlagKey.Token)
	if (configErr != nil || sparkConfig.IdToken == "") && tokenArg == "" {
		color.Red("Token could not be found. Please pass in the token via CLI or Spark config")
		os.Exit(1)
	} else if tokenArg != "" {
		return tokenArg
	}

	return sparkConfig.IdToken
}

var RootCmd = &cobra.Command{
	Use:   "haze <url>",
	Short: "Haze can call Cognito-backed endpoints using JWT",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return errors.New("No URL specified")
		}

		if _, err := url.ParseRequestURI(args[0]); err != nil {
			return errors.Wrap(err, "Invalid URL")
		}

		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag(FlagKey.Method, cmd.Flags().Lookup(FlagKey.Method))
		viper.BindPFlag(FlagKey.Data, cmd.Flags().Lookup(FlagKey.Data))
		viper.BindPFlag(FlagKey.Token, cmd.Flags().Lookup(FlagKey.Token))
	},
	Run: func(cmd *cobra.Command, args []string) {
		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			return
		}

		if version {
			fmt.Println(update.CmdVersion)
			os.Exit(0)
		}

		var method string
		if viper.GetString(FlagKey.Method) == "" {
			method = http.MethodGet
		} else {
			method = viper.GetString(FlagKey.Method)
		}

		token := getTokenFromAllOptions()

		response := request.MakeRequest(&request.Request{
			Method: method,
			Url:    args[0],
			Token:  token,
			Data:   viper.GetString(FlagKey.Data),
		})

		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		fmt.Printf("%+v\n", string(body))
	},
}

func init() {
	RootCmd.Flags().BoolP("version", "v", false, "Current version of Haze")
	RootCmd.Flags().StringP(FlagKey.Method, string(FlagKey.Method[0]), "", "HTTP Method (ex. GET, POST, PATCH, DELETE, etc.)")
	RootCmd.Flags().StringP(FlagKey.Data, string(FlagKey.Data[0]), "", "HTTP POST data")
	RootCmd.Flags().StringP(FlagKey.Token, string(FlagKey.Token[0]), "", "JWT token")
	RootCmd.AddCommand(update.UpdateCmd)
}
