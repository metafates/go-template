# Go CLI Project Template ☄️

Powerful template for Go CLI applications with advanced config management

<img width="912" alt="Screenshot 2022-10-05 at 11 14 32" src="https://user-images.githubusercontent.com/62389790/194013247-897697ee-4b32-4b5d-9667-462fcc45e161.png">


## Features

- Advanced config management with [viper](https://github.com/spf13/viper) and
  useful config commands such as setting config values directly from CLI (like this `config set -k logs.write -v true`),
  reading env variables and file-based configuration (either TOML or YAML). Also, configuration is self-documented, type `config info` to show every config field available with description for each.

- Cache & Temp files management with `clear` command

- Polished CLI experience with [cobra](https://github.com/spf13/cobra) + [coloredcobra](https://github.com/ivanpirog/coloredcobra) to make things look pretty

- [Afero](https://github.com/spf13/afero) filesystem for various fs utils, abstractions and in-memory fs for testing.
  For example, instead of `os.Remove("file")` use `filesystem.Api().Remove("file")`

- Easy to use path management with `where` package

- Logging to file

- Icons!

- Predefined lipgloss colors

## How to use

Press this shiny green button on top

<img width="203" alt="Screenshot 2022-09-30 at 13 37 30" src="https://user-images.githubusercontent.com/62389790/193252456-42b966a7-2679-4868-bf25-d862524733ee.png">

Then you would probably want to rename go mod name from `github.com/metafates/go-template` to something else.
To do this you could use your IDE refactor features or run [make](https://www.gnu.org/software/make/) target.

```shell
make rename
```

This will prompt you to type a new name and will replace every occurence of the old go mod name with the new one.

## Further usage

### Changing name of the app

Change the value of the constant `Name` at [app/meta.go](https://github.com/metafates/go-template/blob/main/app/meta.go)

### Changing config file format from TOML from YAML

Change the value of the constant `ConfigFormat` at [config/init.go](https://github.com/metafates/go-template/blob/main/config/init.go)

### Declaring new config fields

Firstly, declare a field key name as a constant inside [key/keys.go](https://github.com/metafates/go-template/blob/main/key/keys.go)

Then put them inside [config/default.go](https://github.com/metafates/go-template/blob/0a71f1da1c51415469067edbfbe4cbb90e06ef13/config/default.go#L8:L23) (take a predefined fields for logging as a reference)

For example

```go
// key/keys.go

const IconType = "icon.type"
```

```go
// config/default.go

{
  constant.IconType, // config field key
  "emoji", // default value
  "What type of icons to use", // description
}
```

### Accessing config fields

For the example above it would be `viper.GetString(key.EmojiType)`. See [viper](https://github.com/spf13/viper) for more information


## Something is not clear?

Please, [open an issue](https://github.com/metafates/go-template/issues/new) so I could document it
