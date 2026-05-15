package vydon_cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	accounts_cmd "github.com/vydon-io/vydon/cli/internal/cmds/vydon/accounts"
	connections_cmd "github.com/vydon-io/vydon/cli/internal/cmds/vydon/connections"
	jobs_cmd "github.com/vydon-io/vydon/cli/internal/cmds/vydon/jobs"
	login_cmd "github.com/vydon-io/vydon/cli/internal/cmds/vydon/login"
	sync_cmd "github.com/vydon-io/vydon/cli/internal/cmds/vydon/sync"
	version_cmd "github.com/vydon-io/vydon/cli/internal/cmds/vydon/version"
	whoami_cmd "github.com/vydon-io/vydon/cli/internal/cmds/vydon/whoami"
	"github.com/vydon-io/vydon/cli/internal/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
)

const (
	vydonDirName           = ".vydon"
	cliSettingsFileNameNoExt = "config"
	cliSettingsFileExt       = "yaml"

	apiKeyEnvVarName = "VYDON_API_KEY" //nolint:gosec
	apiKeyFlag       = "api-key"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "vydon",
		Short: "Terminal UI that interfaces with the Vydon system.",
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			cmd.SilenceErrors = true
			cmd.SetContext(metadata.NewOutgoingContext(cmd.Context(), version.Get().GrpcMetadata()))
		},
	}

	var cfgFilePath string
	cobra.OnInitialize(
		func() { migrateOldConfig(cfgFilePath) },
		func() { initConfig(cfgFilePath) },
		func() {
			apiKey, err := rootCmd.Flags().GetString(apiKeyFlag)
			if err != nil {
				panic(err)
			}
			envApiKey := viper.GetString(apiKeyEnvVarName)
			if apiKey == "" && envApiKey != "" {
				err = rootCmd.Flags().Set(apiKeyFlag, envApiKey)
				if err != nil {
					panic(err)
				}
			}
		},
	)

	rootCmd.Version = version.Get().GitVersion
	rootCmd.SetVersionTemplate(`{{printf "%s\n" .Version}}`)

	rootCmd.PersistentFlags().StringVar(
		&cfgFilePath, "config", "", fmt.Sprintf("config file (default is $HOME/%s/%s.%s)", vydonDirName, cliSettingsFileNameNoExt, cliSettingsFileExt),
	)
	rootCmd.PersistentFlags().
		String(apiKeyFlag, "", fmt.Sprintf("Vydon API Key. Takes precedence over $%s", apiKeyEnvVarName))

	rootCmd.PersistentFlags().Bool("debug", false, "Run in debug mode")

	rootCmd.AddCommand(jobs_cmd.NewCmd())
	rootCmd.AddCommand(version_cmd.NewCmd())
	rootCmd.AddCommand(whoami_cmd.NewCmd())
	rootCmd.AddCommand(login_cmd.NewCmd())
	rootCmd.AddCommand(sync_cmd.NewCmd())
	rootCmd.AddCommand(accounts_cmd.NewCmd())
	rootCmd.AddCommand(connections_cmd.NewCmd())

	cobra.CheckErr(rootCmd.Execute())
}

// Hack: This method attempts to migrate the old vydon-cli file to the new default location
func migrateOldConfig(cfgFilePath string) {
	if cfgFilePath != "" {
		return
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	oldPath := filepath.Join(home, ".vydon-cli.yaml")
	_, err = os.Stat(oldPath)
	if errors.Is(err, fs.ErrNotExist) {
		return
	}
	err = os.Mkdir(filepath.Join(home, vydonDirName), 0755)
	if err != nil {
		return
	}
	err = os.Rename(
		oldPath,
		filepath.Join(
			home,
			vydonDirName,
			fmt.Sprintf("%s.%s", cliSettingsFileNameNoExt, cliSettingsFileExt),
		),
	)
	if err != nil {
		return
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig(cfgFilePath string) {
	if cfgFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFilePath)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		fullVydonSettingsDir := filepath.Join(home, vydonDirName)
		vydonConfigDir := os.Getenv("VYDON_CONFIG_DIR") // helpful for tools such as direnv and people who want it somewhere interesting
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")       // linux users expect this to be respected

		viper.AddConfigPath(".")
		viper.AddConfigPath(fullVydonSettingsDir)
		viper.AddConfigPath(home)
		if vydonConfigDir != "" {
			viper.AddConfigPath(vydonConfigDir)
		}
		if xdgConfigHome != "" {
			viper.AddConfigPath(xdgConfigHome)
		}

		viper.SetConfigType(cliSettingsFileExt)
		viper.SetConfigName(cliSettingsFileNameNoExt)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
			return
		}
	}
}
