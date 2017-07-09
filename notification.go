package notificationForMac

import (
	"os/exec"
	"regexp"
)

func Enable() error {
	flag, err := IsEnable()
	if err != nil {
		return err
	}
	if !flag {
		Toggle()
	}
	return nil
}
func Disable() error {
	flag, err := IsEnable()
	if err != nil {
		return err
	}
	if flag {
		Toggle()
	}
	return nil
}
func Toggle() error {
	cmd := `
osascript <<EOD
	tell application "System Events" to tell process "SystemUIServer"
		try
			key down option
			click menu bar item 1 of menu bar 1
		on error error_message
			key up option
			display dialog error_message
			set theResult to result
			if not (button returned of theResult = "") then
				tell application "System Preferences"
    				--get a reference to the Security & Privacy preferences pane
    				set securityPane to pane id "com.script.preference.security"
  					--tell that pane to navigate to its "Accessibility" section under its Privacy tab
   					--(the anchor name is arbitrary and does not imply a meaningful hierarchy.)
    				tell securityPane to reveal anchor "Privacy_Accessibility"
    				--open the preferences window and make it frontmost
    				activate
				end tell
			end if
		end try
		key up option
	end tell
EOD`
	err := exec.Command("bash", "-c", cmd).Run()
	return err
}
func GetNotificationCenterUI() ([]byte, error) {
	script := "plutil -convert xml1 -o - ~/Library/Preferences/ByHost/com.apple.notificationcenterui.*.plist"
	cmd := exec.Command("sh", "-c", script)
	xml, err := cmd.Output()
	return xml, err
}
func IsEnable() (bool, error) {
	xml, err := GetNotificationCenterUI()
	if err != nil {
		return false, err
	}
	pattern := "<key>doNotDisturb</key>\\n\\s*<true/>"
	flag, err := regexp.Match(pattern, xml)
	return flag, err
}
