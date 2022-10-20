package generate

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	gcloud "github.com/plantoncloud/mactl/internal/app/build/cloud/gcloud/zshrc"
	golang "github.com/plantoncloud/mactl/internal/app/build/code/lang/golang/zshrc"
	java "github.com/plantoncloud/mactl/internal/app/build/code/lang/java/zshrc"
	nodejs "github.com/plantoncloud/mactl/internal/app/build/code/lang/javascript/nodejs/zshrc"
	python "github.com/plantoncloud/mactl/internal/app/build/code/lang/python/zshrc"
	terraform "github.com/plantoncloud/mactl/internal/app/build/iac/terraform/zshrc"
	kubectl "github.com/plantoncloud/mactl/internal/app/build/kubernetes/kubectl/zshrc"
	gitr "github.com/plantoncloud/mactl/internal/app/build/scm/gitr/zshrc"
	ssh "github.com/plantoncloud/mactl/internal/git/ssh/zshrc"
	git "github.com/plantoncloud/mactl/internal/git/zshrc"
	brew "github.com/plantoncloud/mactl/internal/installer/brew/zshrc"
	"github.com/plantoncloud/mactl/internal/lib/shell"
	"github.com/plantoncloud/mactl/internal/zshrc/override"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	GeneratedZshrcFileName = ".zshrc.mactl.generated"
)

//Cre creates .zshrc.mactl.generated file
func Cre() error {
	generatedContent, err := getGenerated()
	if err != nil {
		return errors.Wrapf(err, "failed to get generated content")
	}
	log.Debug("ensuring generated zshrc file")
	if err := writeGeneratedFile(generatedContent); err != nil {
		return errors.Wrap(err, "failed to ensure zshrc file")
	}
	log.Debug("ensured generated zshrc file")
	return nil
}

func Show() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrapf(err, "failed to get home dir")
	}
	f := filepath.Join(homeDir, GeneratedZshrcFileName)
	if err := shell.RunCmd(exec.Command("code", f)); err != nil {
		return errors.Wrapf(err, "failed to run command to open %s in vs code. is vs code installed?", f)
	}
	return nil
}

func getGenerated() ([]byte, error) {
	var zb strings.Builder
	zb.WriteString("## brew env vars begin\n")
	zb.WriteString(brew.Get())
	zb.WriteString("## brew env vars begin\n")
	zb.WriteString("## common env vars begin\n")
	zb.WriteString(GenericEnvVars)
	zb.WriteString("## common env vars end\n")
	zb.WriteString("## un-aliases begin\n")
	zb.WriteString(UnAliases)
	zb.WriteString("## un-aliases end\n")
	zb.WriteString("## common env aliases begin\n")
	zb.WriteString(GenericAliases)
	zb.WriteString("## common env aliases end\n")
	zb.WriteString("## cd aliases begin\n")
	zb.WriteString(CdAliases)
	zb.WriteString("## cd aliases end\n")
	zb.WriteString("## common functions begin\n")
	appendCommonFunctions(&zb)
	zb.WriteString("## common functions end\n")
	zb.WriteString("## gcloud env vars begin\n")
	zb.WriteString(gcloud.Get())
	zb.WriteString("## gcloud env vars end\n")
	zb.WriteString("## ssh begin\n")
	zb.WriteString(ssh.Get())
	zb.WriteString("## ssh end\n")
	zb.WriteString("## git begin\n")
	zb.WriteString(git.Get())
	zb.WriteString("## git end\n")
	zb.WriteString("## gitr begin\n")
	zb.WriteString(gitr.Get())
	zb.WriteString("## gitr end\n")
	zb.WriteString("## kubectl begin\n")
	zb.WriteString(kubectl.Get())
	zb.WriteString("## kubectl end\n")
	zb.WriteString("## terraform begin\n")
	zb.WriteString(terraform.Get())
	zb.WriteString("## terraform end\n")
	zb.WriteString("## golang begin\n")
	zb.WriteString(golang.Get())
	zb.WriteString("## golang end\n")
	zb.WriteString("## python begin\n")
	zb.WriteString(python.Get())
	zb.WriteString("## python end\n")
	zb.WriteString("## java begin\n")
	javaEnvVars, err := java.Get()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get java env vars")
	}
	zb.WriteString(javaEnvVars)
	zb.WriteString("## java end\n")
	zb.WriteString("## nodejs begin\n")
	zb.WriteString(nodejs.GetEnvVars())
	zb.WriteString("## nodejs end\n")
	if err := appendOverride(&zb); err != nil {
		return nil, errors.Wrap(err, "failed to append extra zshrc")
	}
	return []byte(zb.String()), nil
}

func appendCommonFunctions(zb *strings.Builder) {
	zb.WriteString(UtilFunctions)
	zb.WriteString("\n")
	zb.WriteString(ChromeFunctions)
	zb.WriteString("\n")
	zb.WriteString(GodaddyFunctions)
	zb.WriteString("\n")
	zb.WriteString(BlockChainFunctions)
	zb.WriteString("\n")
	zb.WriteString(LoadTestFunctions)
	zb.WriteString("\n")
	zb.WriteString(LocalHostFunction)
	zb.WriteString("\n")
}

func appendOverride(zb *strings.Builder) error {
	overrideContentBytes, err := override.Get()
	if err != nil {
		return errors.Wrapf(err, "failed to get override content")
	}
	zb.WriteString(string(overrideContentBytes))
	zb.WriteString("\n")
	return nil
}

func writeGeneratedFile(generatedContent []byte) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrapf(err, "failed to get user home dir")
	}
	f := filepath.Join(homeDir, GeneratedZshrcFileName)
	if err := os.WriteFile(f, generatedContent, os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write %s file", f)
	}
	return nil
}
