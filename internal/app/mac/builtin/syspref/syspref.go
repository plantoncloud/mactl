package syspref

const (
	AppName        = "system-preferences"
	MacAppFilePath = "/System/Applications/System Preferences.app"
)

func GetPath() string {
	return MacAppFilePath
}
