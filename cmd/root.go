package mactl

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/plantoncloud/mactl/cmd/mactl/root"
	"github.com/plantoncloud/mactl/internal/cli/flag"
	intos "github.com/plantoncloud/mactl/internal/lib/os"
	"os"
)

var debug bool

var rootCmd = &cobra.Command{
	Use:   "mactl",
	Short: "tool to manage macbook",
	Long:  `mactl simplifies macbook management`,
}

const HomebrewAppleSiliconBinPath = "/opt/homebrew/bin"

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, string(flag.Debug), false, "set log level to debug")
	rootCmd.AddCommand(root.Version, root.Bootstrap, root.Git, root.App, root.Bundle, root.Optimize, root.Zshrc, root.Keyboard)
	rootCmd.DisableSuggestions = true
	cobra.OnInitialize(func() {
		if debug {
			log.SetLevel(log.DebugLevel)
			log.Debug("running in debug mode")
		}
		if intos.IsArmArch() {
			pathEnvVal := os.Getenv("PATH")
			if err := os.Setenv("PATH", fmt.Sprintf("%s:%s", pathEnvVal, HomebrewAppleSiliconBinPath)); err != nil {
				log.Fatalf("failed to set PATH env. err: %v", err)
			}
		}
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
