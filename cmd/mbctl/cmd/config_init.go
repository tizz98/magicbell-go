package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/manifoldco/promptui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"

	magicbell "github.com/tizz98/magicbell-go"
)

func notEmptyValidator(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return fmt.Errorf("cannot be empty")
	}
	return nil
}

var (
	configInitCmd = &cobra.Command{
		Use:   "init",
		Short: "Interactively initialize a config.yaml file for use with this cli",
		RunE: func(cmd *cobra.Command, args []string) error {
			var possibleConfigPaths []string

			if homePath, err := homedir.Expand("~/.config/magicbell/config.yaml"); err == nil {
				possibleConfigPaths = append(possibleConfigPaths, homePath)
			}

			if cwd, err := os.Getwd(); err == nil {
				possibleConfigPaths = append(possibleConfigPaths, path.Join(cwd, "config.yaml"))
			}

			selector := promptui.Select{
				Label: "Config output location",
				Items: possibleConfigPaths,
			}
			_, configPath, err := selector.Run()
			if err != nil {
				return err
			}

			// ensure parent directory exists
			if err := os.MkdirAll(path.Dir(configPath), 0755); err != nil { // #nosec G301
				return fmt.Errorf("unable to create config parent directory: %w", err)
			}

			apiKeyPrompt := promptui.Prompt{
				Label:     "API Key",
				AllowEdit: true,
				Validate:  notEmptyValidator,
			}
			apiSecretPrompt := promptui.Prompt{
				Label:     "API Secret",
				AllowEdit: true,
				Validate:  notEmptyValidator,
			}

			apiKey, err := apiKeyPrompt.Run()
			if err != nil {
				return err
			}
			apiSecret, err := apiSecretPrompt.Run()
			if err != nil {
				return err
			}

			config := magicbell.Config{
				APIKey:    apiKey,
				APISecret: apiSecret,
			}

			encoded, err := yaml.Marshal(config)
			if err != nil {
				return fmt.Errorf("unable to marshal config to yaml: %w", err)
			}

			if err := ioutil.WriteFile(configPath, encoded, 0644); err != nil { // #nosec G306
				return fmt.Errorf("unable to write config file: %w", err)
			}

			logrus.Infof("Wrote config to %s !", configPath)
			return nil
		},
	}
)

func init() {
	configCmd.AddCommand(configInitCmd)
}
