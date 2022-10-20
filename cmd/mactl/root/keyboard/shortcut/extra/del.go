package extra

import (
	"github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp/extra"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Del = &cobra.Command{
	Use:   "del",
	Short: "del extra shortcuts",
	Run:   delHandler,
}

func delHandler(cmd *cobra.Command, args []string) {
	if err := extra.Del(); err != nil {
		log.Fatalf("%v", err)
	}
}
