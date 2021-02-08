package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	magicbell "github.com/tizz98/magicbell-go"
	"github.com/tizz98/magicbell-go/version"
)

type rootOptions struct {
	verbose        bool
	jsonLogging    bool
	configLocation string
}

var (
	rootCmd = &cobra.Command{
		Use:     "mbctl",
		Short:   "Control MagicBell from the command line!",
		Version: fmt.Sprintf("%s (%s)", version.BuildVersion, version.BuildCommit),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Read viper config
			if rootOpts.configLocation != "" {
				f, err := os.Open(rootOpts.configLocation)
				if err != nil {
					return fmt.Errorf("unable to open config file %s: %w", rootOpts.configLocation, err)
				}
				defer f.Close()

				if err := viper.ReadConfig(f); err != nil {
					return fmt.Errorf("unable to read config file %s: %w", rootOpts.configLocation, err)
				}
			} else {
				if err := viper.ReadInConfig(); err != nil {
					if _, ok := err.(viper.ConfigFileNotFoundError); ok {
						logrus.Debugf("config file not found, using default values")
					} else {
						return fmt.Errorf("error reading config: %w", err)
					}
				}
			}

			logrus.SetFormatter(&logrus.TextFormatter{
				FullTimestamp:   true,
				TimestampFormat: time.RFC822,
			})
			if rootOpts.verbose {
				logrus.SetLevel(logrus.DebugLevel)
			}
			if rootOpts.jsonLogging {
				logrus.SetFormatter(&logrus.JSONFormatter{
					TimestampFormat: time.RFC3339Nano,
				})
			}

			timeout := viper.GetDuration("Timeout")
			api = magicbell.New(magicbell.Config{
				APIKey:    viper.GetString("APIKey"),
				APISecret: viper.GetString("APISecret"),
				BaseURL:   viper.GetString("BaseURL"),
				Timeout:   &timeout,
			})

			return nil
		},
	}

	rootOpts = &rootOptions{}
	// api should be used by all sub-commands
	api magicbell.IAPI
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&rootOpts.verbose, "verbose", "v", false, "Enable verbose/debug logging")
	rootCmd.PersistentFlags().BoolVar(&rootOpts.jsonLogging, "json", false, "Print logs in JSON")
	rootCmd.PersistentFlags().StringVarP(&rootOpts.configLocation, "config", "c", "", "Specify an explicit config location. Defaults to $CWD/config.yaml or ~/.config/magicbell/config.yaml")

	// magicbell.Config flags
	rootCmd.PersistentFlags().String("api-key", "", "The MagicBell API key to use in requests")
	rootCmd.PersistentFlags().String("api-secret", "", "The MagicBell API secret to use in requests")
	rootCmd.PersistentFlags().String("base-url", "https://api.magicbell.io", "The MagicBell API URL to use instead of the default")
	rootCmd.PersistentFlags().Duration("timeout", 5*time.Second, "How long to wait for HTTP requests to timeout")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/magicbell")
	viper.AddConfigPath(".")

	_ = viper.BindPFlag("APIKey", rootCmd.Flag("api-key"))
	_ = viper.BindPFlag("APISecret", rootCmd.Flag("api-secret"))
	_ = viper.BindPFlag("BaseURL", rootCmd.Flag("base-url"))
	_ = viper.BindPFlag("Timeout", rootCmd.Flag("timeout"))
}

// Execute runs the mbctl root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
