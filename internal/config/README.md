# CONFIG

Config is the main config of the application. It using [viper][1] package to load the configuration file. Config consist of few objects including `database`, `server`, `account`, `logging`, and `auth`

### File structure
```bash
|-- account/
|-- |-- account.go
|-- |-- config.go
|-- |-- config_test.go
|-- |-- database.go
|-- |-- database_test.go
|-- |-- logger.go
|-- |-- README.md
|-- |-- server.go
|-- |-- server_test.go
```

[1]:https://github.com/spf13/viper
