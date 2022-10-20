package host

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kevinburke/ssh_config"
	lbnutil "github.com/leftbin/go-util/pkg/table"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

const ScmHostComment = "#mactl-scm-host"

func DefaultHosts() []string {
	return []string{"github.com", "gitlab.com", "bitbucket.org"}
}

func List() (hosts []string, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get home dir")
	}
	sshConfigFile := filepath.Join(homeDir, ".ssh", "config")
	f, err := os.Open(sshConfigFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open %s file", sshConfigFile)
	}
	cfg, err := ssh_config.Decode(f)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decode ssh config file %s", sshConfigFile)
	}
	for _, host := range cfg.Hosts {
		if !isScmHost(host) {
			continue
		}
		for _, n := range host.Nodes {
			if strings.Contains(n.String(), "HostName") {
				hn := strings.TrimSpace(n.String())
				if strings.Split(hn, " ")[1] != "" {
					hosts = append(hosts, strings.Split(hn, " ")[1])
				}
			}
		}
	}
	return
}

func isScmHost(h *ssh_config.Host) bool {
	if len(h.Patterns) == 0 {
		return false
	}
	for _, n := range h.Nodes {
		log.Debugf("evaluating node %s", strings.TrimSpace(n.String()))
		if strings.TrimSpace(n.String()) != ScmHostComment {
			continue
		}
		return true
	}
	return false
}

func IsExists(hostname string) (bool, error) {
	hosts, err := List()
	if err != nil {
		return false, errors.Wrap(err, "failed to get existing list of hosts")
	}
	for _, h := range hosts {
		if hostname == h {
			return true, nil
		}
	}
	return false, nil
}

func Match(hostname string, h *ssh_config.Host) bool {
	if !isScmHost(h) {
		return false
	}
	for _, n := range h.Nodes {
		if !strings.Contains(n.String(), "HostName") {
			continue
		}
		hn := strings.TrimSpace(n.String())
		if strings.Split(hn, " ")[1] != hostname {
			continue
		}
		return true
	}
	return false
}

func PrintList(hosts []string) {
	header := table.Row{"#", "name"}
	rows := make([]table.Row, 0)
	for index, h := range hosts {
		rows = append(rows, table.Row{index + 1, h})
	}
	lbnutil.PrintTable(header, rows)
}
