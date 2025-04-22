package app

import (
	"github.com/jedib0t/go-pretty/v6/table"
	lbnutil "github.com/leftbin/go-util/pkg/table"
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/app/bartender"
	"github.com/plantoncloud/mactl/internal/app/battenberg"
	"github.com/plantoncloud/mactl/internal/app/chrome"
	"github.com/plantoncloud/mactl/internal/app/dnsmasq"
	"github.com/plantoncloud/mactl/internal/app/flycut"
	"github.com/plantoncloud/mactl/internal/app/gcloud"
	"github.com/plantoncloud/mactl/internal/app/github"
	"github.com/plantoncloud/mactl/internal/app/gitr"
	"github.com/plantoncloud/mactl/internal/app/iterm"
	"github.com/plantoncloud/mactl/internal/app/karabiner"
	"github.com/plantoncloud/mactl/internal/app/mactl"
	"github.com/plantoncloud/mactl/internal/app/mas"
	"github.com/plantoncloud/mactl/internal/app/rectangle"
	"github.com/plantoncloud/mactl/internal/app/sleepcontrolcenter"
	"github.com/plantoncloud/mactl/internal/app/snagit"
	"github.com/plantoncloud/mactl/internal/app/telnet"
	"github.com/plantoncloud/mactl/internal/app/toothfairy"
	"github.com/plantoncloud/mactl/internal/build/code/lang/java"
)

type App string

const (
	bartenderApp          App = "bartender"
	battenbergApp         App = "battenberg"
	chromeApp             App = "chrome"
	dnsmasqApp            App = "dnsmasq"
	flycutApp             App = "flycut"
	gcloudApp             App = "gcloud"
	githubApp             App = "github"
	gitrApp               App = "gitr"
	itermApp              App = "iterm"
	javaApp               App = "java"
	karabinerApp          App = "karabiner"
	krispApp              App = "krisp"
	mactlApp              App = "mactl"
	masApp                App = "mas"
	rectangleApp          App = "rectangle"
	shushApp              App = "shush"
	sleepControlCenterApp App = "sleep-control-center"
	snagitApp             App = "snagit"
	telnetApp             App = "telnet"
	toothFairyApp         App = "tooth-fairy"
	INVALID               App = "invalid"
)

func Get(app string) App {
	switch app {
	case
		"bartender",
		"battenberg",
		"chrome",
		"dnsmasq",
		"flycut",
		"gcloud",
		"github",
		"gitr",
		"iterm",
		"java",
		"karabiner",
		"keyboard-maestro",
		"krisp",
		"mactl",
		"mas",
		"rectangle",
		"shush",
		"sleep-control-center",
		"snagit",
		"telnet",
		"tooth-fairy":
		return App(app)
	}
	return INVALID
}

func (e App) IsInvalid() bool {
	return e == INVALID
}

func PrintList() {
	apps := []App{
		bartenderApp,
		battenbergApp,
		chromeApp,
		dnsmasqApp,
		flycutApp,
		gcloudApp,
		githubApp,
		gitrApp,
		itermApp,
		javaApp,
		karabinerApp,
		krispApp,
		mactlApp,
		masApp,
		rectangleApp,
		shushApp,
		sleepControlCenterApp,
		snagitApp,
		telnetApp,
		toothFairyApp,
	}
	header := table.Row{"#", "name"}
	rows := make([]table.Row, 0)
	for index, app := range apps {
		rows = append(rows, table.Row{index + 1, app})
	}
	lbnutil.PrintTable(header, rows)
}

func Add(app App) error {
	switch app {
	case bartenderApp:
		return bartender.Setup()
	case chromeApp:
		return chrome.Setup()
	case flycutApp:
		return flycut.Setup()
	case gcloudApp:
		return gcloud.Setup()
	case githubApp:
		return github.Setup()
	case gitrApp:
		return gitr.Setup()
	case itermApp:
		return iterm.Setup()
	case karabinerApp:
		if err := karabiner.Install(); err != nil {
			return errors.Wrapf(err, "failed to install %s", karabinerApp)
		}
		if err := karabiner.Configure(); err != nil {
			return errors.Wrapf(err, "failed to configure karabiner")
		}
		return nil
	case mactlApp:
		return mactl.Setup()
	case masApp:
		return mas.Setup()
	case telnetApp:
		return telnet.Setup()
	case rectangleApp:
		return rectangle.Setup()
	case sleepControlCenterApp:
		return sleepcontrolcenter.Setup()
	case snagitApp:
		return snagit.Setup()
	case toothFairyApp:
		return toothfairy.Setup()
	case battenbergApp:
		return battenberg.Setup()
	case dnsmasqApp:
		return dnsmasq.Setup()
	case javaApp:
		return java.Setup()
	}
	return nil
}
