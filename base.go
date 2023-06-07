package dmsoft

func (com *Dmsoft) EnablePicCache(enable int) int {
	ret, _ := com.dm.CallMethod("EnablePicCache", enable)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetBasePath() string {
	ret, _ := com.dm.CallMethod("GetBasePath")
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) GetDmCount() int {
	ret, _ := com.dm.CallMethod("GetDmCount")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetID() int {
	ret, _ := com.dm.CallMethod("GetID")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetLastError() int {
	ret, _ := com.dm.CallMethod("GetLastError")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) GetPath() string {
	ret, _ := com.dm.CallMethod("GetPath")
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) Reg(regCode string, verInfo string) int {
	ret, _ := com.dm.CallMethod("Reg", regCode, verInfo)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) RegEx(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegEx", regCode, verInfo, ip)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) RegExNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegExNoMac", regCode, verInfo, ip)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) RegNoMac(regCode, verInfo, ip string) int {
	ret, _ := com.dm.CallMethod("RegNoMac", regCode, verInfo, ip)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetDisplayInput(mode string) int {
	ret, _ := com.dm.CallMethod("SetDisplayInput", mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetEnumWindowDelay(delay int) int {
	ret, _ := com.dm.CallMethod("SetEnumWindowDelay", delay)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetPath(path string) int {
	ret, _ := com.dm.CallMethod("SetPath", path)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SetShowErrorMsg(show int) int {
	ret, _ := com.dm.CallMethod("SetShowErrorMsg", show)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) SpeedNormalGraphic(enable int) int {
	ret, _ := com.dm.CallMethod("SpeedNormalGraphic", enable)
	defer ret.Clear()
	return int(ret.Val)
}

// Ver get version
func (com *Dmsoft) Ver() string {
	ver, _ := com.dm.CallMethod("Ver")
	defer ver.Clear()
	return ver.ToString()
}
