package openapp

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	lbnutil "github.com/leftbin/go-util/pkg/table"
	"github.com/pkg/errors"
	karabinercfg "github.com/plantoncloud/mactl/internal/app/keyboard/karabiner/config"
	"github.com/plantoncloud/mactl/internal/keyboard/keys"
	keyboardshortcutopenappconfig "github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp/config"
	_default "github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp/default"
	"sort"
)

// Initialize reads the default karabiner config and adds the default list of keyboard shortcuts to the karabiner config.
func Initialize() error {
	if err := keyboardshortcutopenappconfig.Write(_default.DefaultShortcuts); err != nil {
		return errors.Wrapf(err, "failed to write default shortcuts to config file")
	}
	defaultShortcutMap := convertToMap(_default.DefaultShortcuts)
	c, err := karabinercfg.GetDefault()
	if err != nil {
		return errors.Wrapf(err, "failed to get karabiner config")
	}
	if err := addShortcuts(c, defaultShortcutMap); err != nil {
		return errors.Wrapf(err, "failed to add shortcuts to open apps")
	}
	if err := c.Save(); err != nil {
		return errors.Wrapf(err, "failed to save config with shortcuts")
	}
	return nil
}

// Edit opens config file in vs-code and waits until the file is closed and then reloads the karabiner config.
func Edit() error {
	if err := keyboardshortcutopenappconfig.Edit(); err != nil {
		return errors.Wrapf(err, "failed to edit the config file")
	}
	shortcuts, err := keyboardshortcutopenappconfig.List()
	if err != nil {
		return errors.Wrapf(err, "failed to get list of shortcuts from config file")
	}
	shortcutMap := convertToMap(shortcuts)
	c, err := karabinercfg.GetDefault()
	if err != nil {
		return errors.Wrapf(err, "failed to get karabiner config")
	}
	if err := addShortcuts(c, shortcutMap); err != nil {
		return errors.Wrapf(err, "failed to add shortcuts to open apps")
	}
	if err := c.Save(); err != nil {
		return errors.Wrapf(err, "failed to save config with shortcuts")
	}
	return nil
}

// List returns a map of keyboard shortcuts
func List() (map[keys.KeyBoardKey]*keys.AppShortcut, error) {
	shortcuts, err := keyboardshortcutopenappconfig.List()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get keyboard shortcuts")
	}
	return convertToMap(shortcuts), nil
}

// PrintList prints the provided shortcuts map to the console
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

// addShortcuts adds the provided map of shortcuts to karabiner config
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
							Mandatory: karabinercfg.CapsLockModifierKeys,
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

// convertToMap converts the provided shortcuts into a map
func convertToMap(shortcuts []*keys.AppShortcut) map[keys.KeyBoardKey]*keys.AppShortcut {
	shortcutsMap := make(map[keys.KeyBoardKey]*keys.AppShortcut, 0)
	for _, s := range shortcuts {
		shortcutsMap[s.Key] = s
	}
	return shortcutsMap
}
