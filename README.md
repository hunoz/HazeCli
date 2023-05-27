# Haze
## Description
Haze CLI will allow for easily calling Cognito backed endpoints such as API Gateway endpoints using a Cognito authorizer. This CLI operates using the [Spark CLI](https://github.com/hunoz/spark) information. Spark is NOT required, as this tool offers the passing of the ID token via the `--token` parameter, otherwise it will read it from the Spark config file.

## Installation
1. Navigate to the [releases page](https://github.com/hunoz/haze/releases) and download the binary for your operating system. If you do not see your operating system, please submit an issue with your OS and ARCH so that it can be added.
2. Place the binary in a location in your PATH (e.g. /usr/local/bin/haze)
3. Run `haze help` to see the list of options

## Usage
### Request Without POST Data
```
haze <url>
```

### Request With POST Data
```
haze -d '{"key": "value"}' -m POST <url>
```

### Request With Token From CLI
```
haze --token <token> <url>
```

## Roadmap
* Add cookie flag to allow passing in cookies
* Add output flag to allow passing the response body to a file