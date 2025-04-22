package config

const ComplexModRuleManipulatorTypeBasic = "basic"

type Config struct {
	Global   *Global    `json:"global"`
	Profiles []*Profile `json:"profiles"`
}

type Global struct {
	CheckForUpdatesOnStartup bool `json:"check_for_updates_on_startup"`
	ShowInMenuBar            bool `json:"show_in_menu_bar"`
	ShowProfileNameInMenuBar bool `json:"show_profile_name_in_menu_bar"`
}

type ProfileParameters struct {
	DelayMillisecondsBeforeOpenDevice int `json:"delay_milliseconds_before_open_device"`
}

type Profile struct {
	Name               string              `json:"name,omitempty"`
	Selected           bool                `json:"selected,omitempty"`
	Parameters         *ProfileParameters  `json:"parameters,omitempty"`
	ComplexMod         *ComplexMod         `json:"complex_modifications,omitempty"`
	SimpleMod          []interface{}       `json:"simple_modifications"`
	FnFunctionKeys     []*FnFunctionKeys   `json:"fn_function_keys,omitempty"`
	VirtualHidKeyboard *VirtualHidKeyboard `json:"virtual_hid_keyboard,omitempty"`
}

type ComplexModParameters struct {
	BasicSimultaneousThresholdMilliseconds int `json:"basic.simultaneous_threshold_milliseconds,omitempty"`
	BasicToDelayedActionDelayMilliseconds  int `json:"basic.to_delayed_action_delay_milliseconds,omitempty"`
	BasicToIfAloneTimeoutMilliseconds      int `json:"basic.to_if_alone_timeout_milliseconds,omitempty"`
	BasicToIfHeldDownThresholdMilliseconds int `json:"basic.to_if_held_down_threshold_milliseconds,omitempty"`
	MouseMotionToScrollSpeed               int `json:"mouse_motion_to_scroll.speed,omitempty"`
}

type VirtualHidKeyboard struct {
	CountryCode                     int  `json:"country_code,omitempty"`
	IndicateStickyModifierKeysState bool `json:"indicate_sticky_modifier_keys_state,omitempty"`
	MouseKeyXyScale                 int  `json:"mouse_key_xy_scale,omitempty"`
}

type ComplexModRuleManipulatorFromModifiers struct {
	Mandatory []string `json:"mandatory,omitempty"`
	Optional  []string `json:"optional,omitempty"`
}

type ComplexModRuleManipulatorFrom struct {
	KeyCode   string                                  `json:"key_code,omitempty"`
	Modifiers *ComplexModRuleManipulatorFromModifiers `json:"modifiers,omitempty"`
}

type ComplexModRuleManipulatorTo struct {
	ShellCommand string   `json:"shell_command,omitempty"`
	KeyCode      string   `json:"key_code,omitempty"`
	Modifiers    []string `json:"modifiers,omitempty"`
}

type ComplexModRuleManipulator struct {
	Description string                         `json:"description,omitempty"`
	From        *ComplexModRuleManipulatorFrom `json:"from,omitempty"`
	To          []*ComplexModRuleManipulatorTo `json:"to,omitempty"`
	Type        string                         `json:"type,omitempty"`
}

type ComplexModRule struct {
	Description  string                       `json:"description,omitempty"`
	Manipulators []*ComplexModRuleManipulator `json:"manipulators,omitempty"`
}

type ComplexMod struct {
	Parameters *ComplexModParameters `json:"parameters,omitempty"`
	Rules      []*ComplexModRule     `json:"rules,omitempty"`
}

type FnFunctionKeyFrom struct {
	KeyCode string `json:"key_code,omitempty"`
}

type FnFunctionKeyTo struct {
	KeyCode                    string `json:"key_code,omitempty"`
	ConsumerKeyCode            string `json:"consumer_key_code,omitempty"`
	AppleVendorKeyboardKeyCode string `json:"apple_vendor_keyboard_key_code,omitempty"`
}

type FnFunctionKeys struct {
	From *FnFunctionKeyFrom `json:"from,omitempty"`
	To   []*FnFunctionKeyTo `json:"to,omitempty"`
}
