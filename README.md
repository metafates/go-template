# Go CLI Project Template ☄️

Powerful template for Go CLI applications with advanced config management

## Features

- Advanced config management with [viper](https://github.com/spf13/viper) and
  useful config commands such as setting config values directly from CLI (like this `config set -k logs.write -v true`),
  reading env variables and file-based configuration (either TOML or YAML). Also, configuration is self-documented, type `config info` to show every config field available with description for each.

- Polished CLI experience with [cobra](https://github.com/spf13/cobra) + [coloredcobra](https://github.com/ivanpirog/coloredcobra) to make things look pretty

- [Afero](https://github.com/spf13/afero) filesystem for various fs utils, abstractions and in-memory fs for testing.
  For example, instead of `os.Remove("file")` use `filesystem.Api().Remove("file")`

- Easy to use path management with `where` package

- Logging to file

- Icons!

- Predefined lipgloss colors
