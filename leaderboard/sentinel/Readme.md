# Sentinel

Sentinel is an application for Openhack - DevOps. 
This app do healthcheck for the all service enpoints.

# Usage 

## Configuration 

Create a CosmosDB account. 
Then Rename `config.json.example` to `config.json`. Then edit the `config.json` to fit your CosmosDB account.

`config.json`

```
{
  "url": "https://YOUR_COSMOSDB_ACCOUNT.documents.azure.com:443",
  "masterKey": "YOUR_PRIMARY_KEY"
}

```

## Setup Initial Data  (Production)

This command will remove/recreate the db and insert Team data. However, challenges and Services is null. This is for production setting.

```
sentinel init
```

## Setup Test Data (Test)

This subcommand will create some sample data for testing. 
This subcommand also remove/recreate database.Then install sample data. 

```
sentinel init -t
```

## Setup Test Endpoint

For testing purpose, We provide an endpoint for testing. We develop it as Azure Functions. Please refer the `infrastructure/test/Readme.md`.


## Start Sentinel

This subcommand start sentinel to keep health checking.

```
sentinel start
```

If you want to stop the sentinel, just `ctl + c` or `kill` it.

## Help

```
sentinel --help
NAME:
   Sentinel - Status checking tool for OpenHack - DevOps

USAGE:
   sentinel.exe [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     init, i   Initialize Sentinel. Insert initial Data for an Open Hack
     start, s  Start Sentinel App and monitor the endpoint
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

# Development 

If you want to contribute this tool, you need these tools. 

* [golang](https://golang.org/)
* [golang / dep](https://github.com/golang/dep)

## Setup

Clone this repo then

```
dep ensure
```
This will install whole dependencies. 

## Build

```
go build
```

## Test

```
go test
```

## NOTE

This application uses [a8m/documentdb-go](https://github.com/a8m/documentdb-got). However, it doesn't support `upsert`. I contribute to send [pull request](https://github.com/a8m/documentdb-go/pull/7) to them. However, it is not merged yet. Until then, I publish the folk project. I currently use this. However, once the pull request is merged, I'd like to change to the original one.

