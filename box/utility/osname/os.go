package osname

import "runtime"

const (
	Other int8 = -1 + iota
	Macos
	Linux
	Window
)

func Get() (osName int8) {
	switch runtime.GOOS {
	case "darwin":
		osName = Macos
	case "linux":
		osName = Linux
	case "windows":
		osName = Window
	default:
		osName = Other
	}
	return
}
