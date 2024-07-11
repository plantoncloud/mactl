package _default

import (
	"github.com/plantoncloud/mactl/internal/app/browser/chrome"
	"github.com/plantoncloud/mactl/internal/app/build/api/grpc/client/wombat"
	"github.com/plantoncloud/mactl/internal/app/build/api/rest/client/postman"
	"github.com/plantoncloud/mactl/internal/app/build/code/lang/golang/goland"
	"github.com/plantoncloud/mactl/internal/app/build/code/lang/java/intellij"
	"github.com/plantoncloud/mactl/internal/app/build/code/lang/javascript/webstorm"
	"github.com/plantoncloud/mactl/internal/app/build/code/lang/python/pycharm"
	"github.com/plantoncloud/mactl/internal/app/build/code/lang/sql/datagrip"
	"github.com/plantoncloud/mactl/internal/app/build/code/vscode"
	"github.com/plantoncloud/mactl/internal/app/build/scm/github"
	"github.com/plantoncloud/mactl/internal/app/build/terminal/iterm"
	"github.com/plantoncloud/mactl/internal/app/build/ux/figma"
	"github.com/plantoncloud/mactl/internal/app/comm/slack"
	"github.com/plantoncloud/mactl/internal/app/comm/telegram"
	"github.com/plantoncloud/mactl/internal/app/comm/whatsapp"
	"github.com/plantoncloud/mactl/internal/app/mac/builtin/activitymonitor"
	"github.com/plantoncloud/mactl/internal/app/mac/builtin/calendar"
	"github.com/plantoncloud/mactl/internal/app/mac/builtin/finder"
	"github.com/plantoncloud/mactl/internal/app/mac/builtin/mail"
	"github.com/plantoncloud/mactl/internal/app/mac/builtin/preview"
	"github.com/plantoncloud/mactl/internal/app/mac/builtin/syspref"
	"github.com/plantoncloud/mactl/internal/app/tool/snagit"
	"github.com/plantoncloud/mactl/internal/keyboard/keys"
)

var DefaultShortcuts = []*keys.AppShortcut{
	{
		AppName:     slack.AppName,
		Key:         keys.A,
		AppFilePath: slack.GetPath(),
	}, {
		AppName:     vscode.AppName,
		Key:         keys.C,
		AppFilePath: vscode.GetPath(),
	}, {
		AppName:     chrome.AppName,
		Key:         keys.D,
		AppFilePath: chrome.GetPath(),
	}, {
		AppName:     goland.AppName,
		Key:         keys.G,
		AppFilePath: goland.GetPath(),
	}, {
		AppName:     github.AppName,
		Key:         keys.H,
		AppFilePath: github.GetPath(),
	}, {
		AppName:     snagit.AppName,
		Key:         keys.I,
		AppFilePath: snagit.GetPath(),
	}, {
		AppName:     postman.AppName,
		Key:         keys.K,
		AppFilePath: postman.GetPath(),
	}, {
		AppName:     preview.AppName,
		Key:         keys.P,
		AppFilePath: preview.GetPath(),
	}, {
		AppName:     intellij.AppName,
		Key:         keys.R,
		AppFilePath: intellij.GetPath(),
	}, {
		AppName:     iterm.AppName,
		Key:         keys.S,
		AppFilePath: iterm.GetPath(),
	}, {
		AppName:     finder.AppName,
		Key:         keys.T,
		AppFilePath: finder.GetPath(),
	}, {
		AppName:     pycharm.AppName,
		Key:         keys.V,
		AppFilePath: pycharm.GetPath(),
	}, {
		AppName:     webstorm.AppName,
		Key:         keys.W,
		AppFilePath: webstorm.GetPath(),
	}, {
		AppName:     datagrip.AppName,
		Key:         keys.B,
		AppFilePath: datagrip.GetPath(),
	}, {
		AppName:     mail.AppName,
		Key:         keys.Q,
		AppFilePath: figma.GetPath(),
	}, {
		AppName:     whatsapp.AppName,
		Key:         keys.Z,
		AppFilePath: whatsapp.GetPath(),
	}, {
		AppName:     telegram.AppName,
		Key:         keys.X,
		AppFilePath: telegram.GetPath(),
	}, {
		AppName:     syspref.AppName,
		Key:         keys.NumberOne,
		AppFilePath: syspref.GetPath(),
	}, {
		AppName:     calendar.AppName,
		Key:         keys.NumberTwo,
		AppFilePath: calendar.GetPath(),
	}, {
		AppName:     wombat.AppName,
		Key:         keys.NumberThree,
		AppFilePath: wombat.GetPath(),
	}, {
		AppName:     activitymonitor.AppName,
		Key:         keys.NumberNine,
		AppFilePath: activitymonitor.GetPath(),
	},
}
