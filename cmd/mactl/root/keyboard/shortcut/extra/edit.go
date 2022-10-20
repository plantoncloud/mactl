package extra

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp/extra"
)

var Edit = &cobra.Command{
	Use:   "edit",
	Short: "edit extra shortcuts",
	Run:   editHandler,
}

func editHandler(cmd *cobra.Command, args []string) {
	if err := extra.Edit(); err != nil {
		log.Fatalf("%v", err)
	}
}
