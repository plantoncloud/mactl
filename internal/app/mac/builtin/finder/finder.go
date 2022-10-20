package finder

const (
	AppName        = "finder"
	MacAppFilePath = "/System/Library/CoreServices/Finder.app"
)

func GetPath() string {
	return MacAppFilePath
}
