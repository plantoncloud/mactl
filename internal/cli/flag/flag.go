package flag

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Flag string
type EnvVar string

const (
	Debug     Flag = "debug"
	Host      Flag = "host"
	Workspace Flag = "workspace"
	Username  Flag = "username"
	Email     Flag = "email"
	SshKey    Flag = "ssh-key"
	Copy      Flag = "copy"
	Show      Flag = "show"
	OutFile   Flag = "out-file"
	Name      Flag = "name"
)

func HandleFlagErrAndVal(err error, flag Flag, flagVal string) {
	if err != nil {
		log.Fatalf("error parsing %s flag. err %v", flag, err)
	}
	if flagVal == "" {
		log.Fatalf("empty val not allowed for %s flag", flag)
	}
}

func HandleFlagErr(err error, flag Flag) {
	if err != nil {
		log.Fatalf("error parsing %s flag. err %v", flag, err)
	}
}

func MarkAllFlagsRequired(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := cmd.MarkPersistentFlagRequired(f.Name); err != nil {
			log.Fatalf("failed to mark %s flag as required", f.Name)
		}
	})
}

func MarkFlagsRequired(cmd *cobra.Command, flags ...Flag) {
	for _, f := range flags {
		if err := cmd.MarkPersistentFlagRequired(string(f)); err != nil {
			log.Fatalf("failed to mark %s flag as required", f)
		}
	}
}
