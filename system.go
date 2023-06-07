// 系统

package dmsoft

func (com *Dmsoft) Beep(f, duration int) int {
	ret, _ := com.dm.CallMethod("Beep", f, duration)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) CheckFontSmooth() int {
	ret, _ := com.dm.CallMethod("CheckFontSmooth")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) CheckUAC() int {
	ret, _ := com.dm.CallMethod("CheckUAC")
	defer ret.Clear()
	return int(ret.Val)
}

// Delay 函数不推荐使用
// 建议time.Sleep()
func (com *Dmsoft) Delay(mis int) int {
	ret, _ := com.dm.CallMethod("Delay")
	defer ret.Clear()
	return int(ret.Val)
}
func (com *Dmsoft) Delays(misMin, misMax int) int {
	ret, _ := com.dm.CallMethod("Delays", misMin, misMax)
	defer ret.Clear()
	return int(ret.Val)
}

// DisableCloseDisplayAndSleep 不支持xp系统
func (com *Dmsoft) DisableCloseDisplayAndSleep() int {
	ret, _ := com.dm.CallMethod("DisableCloseDisplayAndSleep")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) DisableFontSmooth() int {
	ret, _ := com.dm.CallMethod("DisableFontSmooth")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) DisablePowerSave() int {
	ret, _ := com.dm.CallMethod("DisablePowerSave")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) DisableScreenSave() int {
	ret, _ := com.dm.CallMethod("DisableScreenSave")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableFontSmooth() int {
	ret, _ := com.dm.CallMethod("EnableFontSmooth")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ExitOs(types int) int {
	ret, _ := com.dm.CallMethod("ExitOs", types)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetClipboard() string {
	ret, _ := com.dm.CallMethod("GetClipboard")
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetCPUType() int {
	ret, _ := com.dm.CallMethod("GetCpuType")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetCPUUsage() int {
	ret, _ := com.dm.CallMethod("GetCpuUsage")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetDir(types int) string {
	ret, _ := com.dm.CallMethod("GetDir", types)
	defer ret.Clear()
	return ret.ToString()
}

// GetDiskModel 需要管理员权限
func (com *Dmsoft) GetDiskModel(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskModel", index)
	defer ret.Clear()
	return ret.ToString()
}

// GetDiskReversion 需要管理员权限
func (com *Dmsoft) GetDiskReversion(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskReversion", index)
	defer ret.Clear()
	return ret.ToString()
}

// GetDiskSerial 需要管理员权限
func (com *Dmsoft) GetDiskSerial(index int) string {
	ret, _ := com.dm.CallMethod("GetDiskSerial", index)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetDisplayInfo() string {
	ret, _ := com.dm.CallMethod("GetDisplayInfo")
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetDPI() int {
	ret, _ := com.dm.CallMethod("GetDPI")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetLocale() int {
	ret, _ := com.dm.CallMethod("GetLocale")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetMachineCode() string {
	ret, _ := com.dm.CallMethod("GetMachineCode")
	defer ret.Clear()
	return ret.ToString()
}

// GetMachineCodeNoMac 需要管理员权限
func (com *Dmsoft) GetMachineCodeNoMac() string {
	ret, _ := com.dm.CallMethod("GetMachineCodeNoMac")
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetMemoryUsage() int {
	ret, _ := com.dm.CallMethod("GetMemoryUsage")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetNetTime() string {
	ret, _ := com.dm.CallMethod("GetNetTime")
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetNetTimeByIP(ip string) string {
	ret, _ := com.dm.CallMethod("GetNetTimeByIp", ip)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetOsBuildNumber() int {
	ret, _ := com.dm.CallMethod("GetOsBuildNumber")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetOsType() int {
	ret, _ := com.dm.CallMethod("GetOsType")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetScreenDepth() int {
	ret, _ := com.dm.CallMethod("GetScreenDepth")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetScreenHeight() int {
	ret, _ := com.dm.CallMethod("GetScreenHeight")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetScreenWidth() int {
	ret, _ := com.dm.CallMethod("GetScreenWidth")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetTime() int {
	ret, _ := com.dm.CallMethod("GetTime")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) Is64Bit() int {
	ret, _ := com.dm.CallMethod("Is64Bit")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) IsSurrpotVt() int {
	ret, _ := com.dm.CallMethod("IsSurrpotVt")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) Play(mediaFile string) int {
	ret, _ := com.dm.CallMethod("Play", mediaFile)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) RunApp(appPath string, mode int) int {
	ret, _ := com.dm.CallMethod("RunApp", appPath, mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetClipboard(value string) int {
	ret, _ := com.dm.CallMethod("SetClipboard", value)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayAcceler(level int) int {
	ret, _ := com.dm.CallMethod("SetDisplayAcceler", level)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetLocale() int {
	ret, _ := com.dm.CallMethod("SetLocale")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetScreen(width, height, depth int) int {
	ret, _ := com.dm.CallMethod("SetScreen", width, height, depth)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetUAC(enable int) int {
	ret, _ := com.dm.CallMethod("SetUAC", enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ShowTaskBarIcon(hwnd, isShow int) int {
	ret, _ := com.dm.CallMethod("ShowTaskBarIcon", hwnd, isShow)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) Stop(id int) int {
	ret, _ := com.dm.CallMethod("Stop", id)
	defer ret.Clear()
	return int(ret.Val)
}
