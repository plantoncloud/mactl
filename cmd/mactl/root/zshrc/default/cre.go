package _default

import (
	_default "github.com/plantoncloud/mactl/internal/zshrc/default"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cre = &cobra.Command{
	Use:   "cre",
	Short: "cre default zshrc",
	Run:   creHandler,
}

func creHandler(cmd *cobra.Command, args []string) {
	if err := _default.Cre(); err != nil {
		log.Fatalf("%v", err)
	}
}
