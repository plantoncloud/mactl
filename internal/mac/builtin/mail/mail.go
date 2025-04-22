package mail

const (
	AppName        = "mail"
	MacAppFilePath = "/System/Applications/Mail.app"
)

func GetPath() string {
	return MacAppFilePath
}
