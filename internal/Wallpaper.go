package internal

import (
	"syscall"
	"unsafe"
)

const (
	// UI action to set desktop wallpaper
	spiSetDeskWallpaper = 0x0014

	uiParam = 0x0000

	// Writes new system-wide parameter setting to the user profile
	spifUpdateINIFile = 0x01
	// Broadcast WM_SETTINGCHANGE message after updating user profile
	spifSendChange = 0x02
)

var (
	user32                = syscall.NewLazyDLL("user32.dll")
	systemParametersInfoW = user32.NewProc("SystemParametersInfoW")
)

// SetWallpaper sets desktop wallpaper.
// The parameter `filePath` must be absolute path to the file.
func SetWallpaper(filePath string) error {
	filePtr, err := syscall.UTF16PtrFromString(filePath)
	if err != nil {
		return err
	}
	systemParametersInfoW.Call(
		uintptr(spiSetDeskWallpaper),
		uintptr(uiParam),
		uintptr(unsafe.Pointer(filePtr)),
		uintptr(spifUpdateINIFile|spifSendChange),
	)
	return nil
}
