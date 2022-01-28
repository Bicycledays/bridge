package service

type SystemType int

const (
	SystemWindows SystemType = 1
	SystemLinux   SystemType = 2
	SystemMac     SystemType = 3
)

func DetectSystem() SystemType {
	return SystemWindows
}
