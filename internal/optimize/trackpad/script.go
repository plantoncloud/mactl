package trackpad

const (
	AppleScriptToIncTrackingSpeed = `
set trackingValue to 10

--Open and activate System Preferences
tell application "System Preferences" to activate
tell application "System Preferences"
    reveal pane "com.apple.preference.trackpad"
end tell
--Attempt to change settings using System Events
tell application "System Events"
    tell process "System Preferences"
        try
            delay 1
            tell tab group 1 of window "Trackpad"
                delay 1
                set value of slider "Tracking Speed" to trackingValue
            end tell
        on error theError
            --An error occured
            display dialog ("sorry, an error occured while changing trackpad speed:" & return & theError) buttons "ok" default button "ok"
        end try
    end tell
end tell
tell application "System Preferences" to quit
`
)
