# GO CLI Project Template 👾

Template for Go CLI applications with advanced config management

## Features

- Advanced config management with viper and
  useful config commands such as setting config values directly from CLI (like this `config set -k logs.write -v true`),
  reading env variables and file-based configuration (either TOML or YAML). Also, configuration is self-documented, type `config info` to show every config field available with description for each.

- Polished CLI experience with cobra library + coloredcobra to make things look pretty

- Afero filesystem for various fs utils, abstractions and in-memory fs for testing

- Easy to use path management with `where` module

- Logging to file

- Icons!

- Predefined lipgloss colors

