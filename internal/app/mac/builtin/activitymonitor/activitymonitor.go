package activitymonitor

const (
	AppName        = "activity-monitor"
	MacAppFilePath = "/System/Applications/Utilities/Activity Monitor.app"
)

func GetPath() string {
	return MacAppFilePath
}
