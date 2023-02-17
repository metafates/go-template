package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/metafates/go-template/app"
	"github.com/metafates/go-template/style"
	"golang.org/x/exp/slices"

	levenshtein "github.com/ka-weihe/fast-levenshtein"
	"github.com/metafates/go-template/color"
	"github.com/metafates/go-template/config"
	"github.com/metafates/go-template/filesystem"
	"github.com/metafates/go-template/icon"
	"github.com/metafates/go-template/where"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// errUnknownKey will generate error for key that was not found and will provide a hint
func errUnknownKey(key string) error {
	closest := lo.MinBy(lo.Keys(config.Default), func(a string, b string) bool {
		return levenshtein.Distance(key, a) < levenshtein.Distance(key, b)
	})
	msg := fmt.Sprintf(
		"unknown key %s, did you mean %s?",
		style.Fg(color.Red)(key),
		style.Fg(color.Yellow)(closest),
	)

	return errors.New(msg)
}

func completionConfigKeys(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
	return lo.Keys(config.Default), cobra.ShellCompDirectiveNoFileComp
}

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Various config commands",
}

func init() {
	configCmd.AddCommand(configInfoCmd)
	configInfoCmd.Flags().StringP("key", "k", "", "The key to get the value for")
	_ = configInfoCmd.RegisterFlagCompletionFunc("key", completionConfigKeys)
}

var configInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show the info for each config field with description",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			key    = lo.Must(cmd.Flags().GetString("key"))
			fields = lo.Values(config.Default)
		)

		if key != "" {
			if field, ok := config.Default[key]; ok {
				fields = []*config.Field{field}
			} else {
				handleErr(errUnknownKey(key))
			}
		}

		sort.Slice(fields, func(i, j int) bool {
			return fields[i].Key < fields[j].Key
		})

		for i, field := range fields {
			fmt.Println(field.Pretty())

			if i < len(fields)-1 {
				fmt.Println()
			}
		}
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
	configSetCmd.Flags().StringP("key", "k", "", "The key to set the value for")
	lo.Must0(configSetCmd.MarkFlagRequired("key"))
	_ = configSetCmd.RegisterFlagCompletionFunc("key", completionConfigKeys)

	configSetCmd.Flags().StringP("value", "v", "", "The value to set")
	lo.Must0(configSetCmd.MarkFlagRequired("value"))
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a config value",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			key   = lo.Must(cmd.Flags().GetString("key"))
			value = lo.Must(cmd.Flags().GetString("value"))
		)

		if _, ok := config.Default[key]; !ok {
			handleErr(errUnknownKey(key))
		}

		var v any
		switch config.Default[key].DefaultValue.(type) {
		case string:
			v = value
		case int:
			parsedInt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				handleErr(fmt.Errorf("invalid integer value: %s", value))
			}

			v = int(parsedInt)
		case bool:
			parsedBool, err := strconv.ParseBool(value)
			if err != nil {
				handleErr(fmt.Errorf("invalid boolean value: %s", value))
			}

			v = parsedBool
		}

		viper.Set(key, v)
		switch err := viper.WriteConfig(); err.(type) {
		case viper.ConfigFileNotFoundError:
			handleErr(viper.SafeWriteConfig())
		default:
			handleErr(err)
		}

		fmt.Printf(
			"%s set %s to %s\n",
			style.Fg(color.Green)(icon.Check),
			style.Fg(color.Purple)(key),
			style.Fg(color.Yellow)(fmt.Sprintf("%v", v)),
		)
	},
}

func init() {
	configCmd.AddCommand(configGetCmd)
	configGetCmd.Flags().StringP("key", "k", "", "The key to get the value for")
	lo.Must0(configGetCmd.MarkFlagRequired("key"))
	_ = configGetCmd.RegisterFlagCompletionFunc("key", completionConfigKeys)
}

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a config value",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			key = lo.Must(cmd.Flags().GetString("key"))
		)

		if _, ok := config.Default[key]; !ok {
			handleErr(errUnknownKey(key))
		}

		fmt.Println(viper.Get(key))
	},
}

func init() {
	configCmd.AddCommand(configEnvCmd)
}

var configEnvCmd = &cobra.Command{
	Use:   "env",
	Short: "Show the env for each config field",
	Run: func(cmd *cobra.Command, args []string) {
		fields := lo.Values(config.Default)
		fields = append(fields, &config.Field{Key: where.EnvConfigPath})

		slices.SortFunc(fields, func(a, b *config.Field) bool {
			return a.Key < b.Key
		})

		for _, field := range fields {
			envValue, isSet := os.LookupEnv(field.Env())

			var value string

			if isSet {
				value = envValue
			} else {
				value = style.Faint("unset")
			}

			_, err := fmt.Fprintf(cmd.OutOrStdout(), "%s=%s\n", field.Env(), value)
			handleErr(err)
		}
	},
}

func init() {
	configCmd.AddCommand(configWriteCmd)
	configWriteCmd.Flags().BoolP("force", "f", false, "Force overwrite of existing config file")
}

var configWriteCmd = &cobra.Command{
	Use:   "write",
	Short: "Write current config to the file",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			force          = lo.Must(cmd.Flags().GetBool("force"))
			configFilePath = filepath.Join(
				where.Config(),
				fmt.Sprintf("%s.%s", app.Name, config.ConfigFormat),
			)
		)

		if force {
			exists, err := filesystem.Api().Exists(configFilePath)
			handleErr(err)

			if exists {
				err := filesystem.Api().Remove(configFilePath)
				handleErr(err)
			}
		}

		handleErr(viper.SafeWriteConfig())
		fmt.Printf(
			"%s wrote config to %s\n",
			style.Fg(color.Green)(icon.Check),
			configFilePath,
		)
	},
}

func init() {
	configCmd.AddCommand(configDeleteCmd)
}

var configDeleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete the config file",
	Aliases: []string{"remove"},
	Run: func(cmd *cobra.Command, args []string) {
		configFilePath := filepath.Join(where.Config(), fmt.Sprintf("%s.%s", app.Name, config.ConfigFormat))

		exists, err := filesystem.Api().Exists(configFilePath)
		handleErr(err)

		if !exists {
			fmt.Printf(
				"%s nothing to delete\n",
				style.Fg(color.Green)(icon.Check),
			)
			return
		}

		err = filesystem.Api().Remove(configFilePath)

		handleErr(err)
		fmt.Printf(
			"%s deleted config\n",
			style.Fg(color.Green)(icon.Check),
		)
	},
}
