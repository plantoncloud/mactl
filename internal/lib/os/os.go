package os

import "runtime"

type Arch string

const (
	AMD64   Arch = "amd64"
	ARM64   Arch = "arm64"
	UNKNOWN Arch = "unknown"
)

func IsArmArch() bool {
	return runtime.GOARCH == "arm64"
}

func IsAmdArch() bool {
	return runtime.GOARCH == "amd64"
}

func GetArch() Arch {
	if runtime.GOARCH == "amd64" {
		return AMD64
	}
	if runtime.GOARCH == "arm64" {
		return ARM64
	}
	return UNKNOWN
}

func IsUnknown() bool {
	return GetArch() == UNKNOWN
}
