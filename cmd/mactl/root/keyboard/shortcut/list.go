package shortcut

import (
	"github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "list shortcuts",
	Run:   listHandler,
}

func listHandler(cmd *cobra.Command, args []string) {
	list, err := openapp.List()
	if err != nil {
		log.Fatalf("%v", err)
	}
	openapp.PrintList(list)
}
