package extra

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp"
)

var Reload = &cobra.Command{
	Use:   "reload",
	Short: "reload extra shortcuts",
	Run:   reloadHandler,
}

func reloadHandler(cmd *cobra.Command, args []string) {
	if err := openapp.Setup(); err != nil {
		log.Fatalf("%v", err)
	}
	log.Info("success!")
}
