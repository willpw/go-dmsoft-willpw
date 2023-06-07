// 后台

package dmsoft

func (com *Dmsoft) BindWindow(hwnd int, display string, mouse string, keypad string, mode int) int {
	ret, _ := com.dm.CallMethod("BindWindow", hwnd, display, mouse, keypad, mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) BindWindowEx(hwnd int, display string, mouse string, keypad string, public string, mode int) int {
	ret, _ := com.dm.CallMethod("BindWindowEx", hwnd, display, mouse, keypad, public, mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) DownCPU(rate int) int {
	ret, _ := com.dm.CallMethod("DownCpu", rate)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableBind(enable int) int {
	ret, _ := com.dm.CallMethod("EnableBind", enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableFakeActive(enable int) int {
	ret, _ := com.dm.CallMethod("EnableFakeActive", enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableIme(enable int) int {
	ret, _ := com.dm.CallMethod("EnableIme", enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableKeypadMsg(enable int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadMsg", enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableKeypadPatch(enable int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadPatch", enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableKeypadSync(enable, timeOut int) int {
	ret, _ := com.dm.CallMethod("EnableKeypadSync", enable, timeOut)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableMouseMsg(enable int) int {
	ret, _ := com.dm.CallMethod("EnableMouseMsg", enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableMouseSync(enable, timeOut int) int {
	ret, _ := com.dm.CallMethod("EnableMouseSync", enable, timeOut)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableRealKeypad(enable int) int {
	ret, _ := com.dm.CallMethod("EnableRealKeypad", enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableRealMouse(enable, mousedelay, mousestep int) int {
	ret, _ := com.dm.CallMethod("EnableRealMouse", enable, mousedelay, mousestep)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) EnableSpeedDx(enable int) int {
	ret, _ := com.dm.CallMethod("EnableSpeedDx", enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ForceUnBindWindow() int {
	ret, _ := com.dm.CallMethod("ForceUnBindWindow")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetBindWindow() int {
	ret, _ := com.dm.CallMethod("GetBindWindow")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetFps() int {
	ret, _ := com.dm.CallMethod("GetFps")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) HackSpeed(rate int) int {
	ret, _ := com.dm.CallMethod("HackSpeed", rate)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) IsBind(hwnd int) int {
	ret, _ := com.dm.CallMethod("IsBind", hwnd)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) LockDisplay(lock int) int {
	ret, _ := com.dm.CallMethod("LockDisplay", lock)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) LockInput(lock int) int {
	ret, _ := com.dm.CallMethod("LockInput", lock)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) LockMouseRect(x1, y1, x2, y2 int) int {
	ret, _ := com.dm.CallMethod("LockMouseRect", x1, y1, x2, y2)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetAero(enable int) int {
	ret, _ := com.dm.CallMethod("SetAero", enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayDelay(time int) int {
	ret, _ := com.dm.CallMethod("SetDisplayDelay", time)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayRefreshDelay(time int) int {
	ret, _ := com.dm.CallMethod("SetDisplayRefreshDelay", time)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SwitchBindWindow(hwnd int) int {
	ret, _ := com.dm.CallMethod("SwitchBindWindow", hwnd)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) UnBindWindow() int {
	ret, _ := com.dm.CallMethod("UnBindWindow")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetInputDm(dm_id, rx, ry int) int {
	ret, _ := com.dm.CallMethod("SetInputDm", dm_id, rx, ry)
	defer ret.Clear()
	return int(ret.Val)
}
