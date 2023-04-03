# mage
MageGames Case Study
## Clone and Dicrectory

Clone the project

```
  git clone https://github.com/yclyldrm/mage.git
```

Go to the project directory

```bash
  cd mage
```

## Run Locally

Change/Add configuration file

Need to change/add .env file into /pkg/configs/envs/.env

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run main.go
```

## Makefile

To generate protobuf files run the following command:

```bash
  make proto
```

This will generate .go files for all .proto files in the project.