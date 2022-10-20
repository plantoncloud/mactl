package extra

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "init extra shortcuts",
	Run:   initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	if err := openapp.Setup(); err != nil {
		log.Fatalf("%v", err)
	}
	log.Info("success!")
}
