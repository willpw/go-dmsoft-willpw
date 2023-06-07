// 汇编

package dmsoft

func (com *Dmsoft) AsmAdd(asmIns string) int {
	ret, _ := com.dm.CallMethod("AsmAdd", asmIns)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AsmCall(hwnd, mode int) int {
	ret, _ := com.dm.CallMethod("AsmCall", hwnd, mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AsmCallEx(hwnd, mode int, baseAddr string) int64 {
	ret, _ := com.dm.CallMethod("AsmCallEx", hwnd, mode, baseAddr)
	defer ret.Clear()
	return ret.Val
}

func (com *Dmsoft) AsmClear() int {
	ret, _ := com.dm.CallMethod("AsmClear")
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AsmSetTimeout(timeOut, param int) int {
	ret, _ := com.dm.CallMethod("AsmSetTimeout", timeOut, param)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) Assemble(timeOut int64, is64bit int) string {
	ret, _ := com.dm.CallMethod("Assemble", timeOut, is64bit)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) DisAssemble(asmCode string, baseAddr int64, is64bit int) string {
	ret, _ := com.dm.CallMethod("DisAssemble", asmCode, baseAddr, is64bit)
	defer ret.Clear()
	return ret.ToString()
}
