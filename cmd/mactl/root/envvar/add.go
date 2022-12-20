package envvar

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/cli/clierr"
	"github.com/plantoncloud/mactl/internal/envvar"
	"github.com/spf13/cobra"
)

var Add = &cobra.Command{
	Use:   "add",
	Short: "add environment variable to environment. pass variable name and value as two command arguments to add.",
	Run:   addHandler,
}

func addHandler(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		clierr.HandleDefault(errors.New("two arguments variable name and value are required."))
	}
	clierr.HandleDefault(envvar.Add(args[0], args[1]))
}
