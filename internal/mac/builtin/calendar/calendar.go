package calendar

const (
	AppName        = "calendar"
	MacAppFilePath = "/System/Applications/Calendar.app"
)

func GetPath() string {
	return MacAppFilePath
}
