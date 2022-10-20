package extra

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp/extra"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "list shortcuts",
	Run:   listHandler,
}

func listHandler(cmd *cobra.Command, args []string) {
	list, err := extra.List()
	if err != nil {
		log.Fatalf("%v", err)
	}
	if len(list) == 0 {
		log.Infof("no extra shortcuts")
		return
	}
	extra.PrintList(list)
}
