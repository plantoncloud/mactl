package flag

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Flag string
type EnvVar string

const (
	Copy                 Flag = "copy"
	Debug                Flag = "debug"
	Email                Flag = "email"
	ExcludeFileExtension Flag = "exclude-file-extension"
	Host                 Flag = "host"
	Name                 Flag = "name"
	OutFile              Flag = "out-file"
	Show                 Flag = "show"
	SourceDir            Flag = "source-dir"
	SshKey               Flag = "ssh-key"
	Username             Flag = "username"
	Val                  Flag = "val"
	Workspace            Flag = "workspace"
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
