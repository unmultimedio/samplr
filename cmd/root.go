package cmd

import (
	"os"

	"github.com/op/go-logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfigFile = ".samplr" // Without extension
)

var (
	logger *logging.Logger
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "",
		Short: "A sample files generator for the current directory",
		Long: `Samplr is a CLI that generates sample copies for any file
that wants to easily offuscate sensitive information to avoid check it
in repositories.`,
		Run: func(cmd *cobra.Command, args []string) {
			Sample()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./"+defaultConfigFile+".yml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose mode (default false)")
	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.SetDefault("verbose", false)
}

func initConfig() {
	logger = logging.MustGetLogger("samplr")
	loggerBackend := logging.NewLogBackend(os.Stderr, "", 0)
	leveledLogger := logging.AddModuleLevel(loggerBackend)
	leveledLogger.SetLevel(logging.ERROR, "")
	if viper.GetBool("verbose") {
		leveledLogger.SetLevel(logging.INFO, "")
	}
	logger.SetBackend(leveledLogger)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory with name (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName(defaultConfigFile)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		logger.Info("Using config file:", viper.ConfigFileUsed())
	}
}
