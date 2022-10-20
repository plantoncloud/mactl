package preview

const (
	AppName        = "preview"
	MacAppFilePath = "/System/Applications/Preview.app"
)

func GetPath() string {
	return MacAppFilePath
}
