package openapp

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	lbnutil "github.com/leftbin/go-util/pkg/table"
	"github.com/pkg/errors"
	karabinercfg "github.com/plantoncloud/mactl/internal/app/keyboard/karabiner/config"
	"github.com/plantoncloud/mactl/internal/keyboard/keys"
	_default "github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp/default"
	"github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp/extra"
	"sort"
)

func Setup() error {
	c, err := karabinercfg.GetDefault()
	if err != nil {
		return errors.Wrapf(err, "failed to get karabiner config")
	}
	extraShortcuts, err := extra.List()
	if err != nil {
		return errors.Wrapf(err, "failed to get extra shortcuts")
	}
	merged := mergeShortcuts(_default.DefaultShortcuts, extraShortcuts)
	if err := addShortcuts(c, merged); err != nil {
		return errors.Wrapf(err, "failed to add shortcuts to open apps")
	}
	if err := c.Save(); err != nil {
		return errors.Wrapf(err, "failed to save config with shortcuts")
	}
	return nil
}

func List() (map[keys.KeyBoardKey]*keys.AppShortcut, error) {
	extraShortcuts, err := extra.List()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get extraShortcuts shortcuts")
	}
	return mergeShortcuts(_default.DefaultShortcuts, extraShortcuts), nil
}

func PrintList(shortcuts map[keys.KeyBoardKey]*keys.AppShortcut) {
	header := table.Row{"key", "app name", "app path"}
	rows := make([]table.Row, 0)
	hotkeys := make([]string, 0)
	for k, _ := range shortcuts {
		hotkeys = append(hotkeys, string(k))
	}
	sort.Strings(hotkeys)
	for _, r := range hotkeys {
		rows = append(rows, table.Row{shortcuts[keys.KeyBoardKey(r)].Key, shortcuts[keys.KeyBoardKey(r)].AppName, shortcuts[keys.KeyBoardKey(r)].AppFilePath})
	}
	lbnutil.PrintTable(header, rows)
}

func addShortcuts(c *karabinercfg.Config, shortcuts map[keys.KeyBoardKey]*keys.AppShortcut) error {
	defaultProfile, err := c.GetDefaultProfile()
	if err != nil {
		return errors.Wrapf(err, "failed to get default profile")
	}
	rules := make([]*karabinercfg.ComplexModRule, 0)
	for _, s := range shortcuts {
		rules = append(rules, &karabinercfg.ComplexModRule{
			Description: fmt.Sprintf("open %s app", s.AppName),
			Manipulators: []*karabinercfg.ComplexModRuleManipulator{
				{
					From: &karabinercfg.ComplexModRuleManipulatorFrom{
						KeyCode: string(s.Key),
						Modifiers: &karabinercfg.ComplexModRuleManipulatorFromModifiers{
							Mandatory: karabinercfg.CapsLockMofifierKeys,
						},
					},
					To: []*karabinercfg.ComplexModRuleManipulatorTo{
						{
							ShellCommand: fmt.Sprintf("open -a '%s'", s.AppFilePath),
						},
					},
					Type: karabinercfg.ComplexModRuleManipulatorTypeBasic,
				},
			},
		})
	}
	defaultProfile.ComplexMod.Rules = append(defaultProfile.ComplexMod.Rules, rules...)
	return nil
}

func mergeShortcuts(defaultShortcuts, extraShortcuts []*keys.AppShortcut) map[keys.KeyBoardKey]*keys.AppShortcut {
	merged := make(map[keys.KeyBoardKey]*keys.AppShortcut, 0)
	for _, s := range defaultShortcuts {
		merged[s.Key] = s
	}
	for _, s := range extraShortcuts {
		merged[s.Key] = s
	}
	return merged
}
