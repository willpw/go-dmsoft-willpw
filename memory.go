// 内存操作暂时不需要

package dmsoft

// string DoubleToData(value)
// string FindData(hwnd, addr_range, data)
// string FindDataEx(hwnd, addr_range, data,step,multi_thread,mode)
// string FindDouble(hwnd, addr_range, double_value_min, double_value_max)
// string FindDoubleEx(hwnd, addr_range, double_value_min, double_value_max,step,multi_thread,mode)
// string FindFloat(hwnd, addr_range, float_value_min, float_value_max)
// string FindFloatEx(hwnd, addr_range, float_value_min, float_value_max,step,multi_thread,mode)
// string FindInt(hwnd, addr_range, int_value_min, int_value_max,type)
// string FindIntEx(hwnd, addr_range, int_value_min, int_value_max,type,step,multi_thread,mode)
// string FindString(hwnd, addr_range, string_value,type)
// string FindStringEx(hwnd, addr_range, string_value,type,step,multi_thread,mode)
// string FloatToData(value)
// long FreeProcessMemory(hwnd)

func (com *Dmsoft) GetCommandLine(hwnd int) string {
	ret, _ := com.dm.CallMethod("GetCommandLine", hwnd)
	defer ret.Clear()
	return ret.ToString()
}

// LONGLONG GetModuleBaseAddr(hwnd,module)
// long GetModuleSize(hwnd,module)
// LONGLONG GetRemoteApiAddress(hwnd,base_addr,fun_name)
// long Int64ToInt32(value)
// string IntToData(value,type)

func (com *Dmsoft) OpenProcess(pid int) int {
	ret, _ := com.dm.CallMethod("OpenProcess", pid)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ReadData(hwnd int, addr string, le int) string {
	ret, _ := com.dm.CallMethod("ReadData", hwnd, addr, le)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) ReadDataAddr(hwnd int, addr int64, le int) string {
	ret, _ := com.dm.CallMethod("ReadDataAddr", hwnd, addr, le)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) ReadDataAddrToBin(hwnd int, addr int64, le int) int {
	ret, _ := com.dm.CallMethod("ReadDataAddrToBin", hwnd, addr, le)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ReadDataToBin(hwnd int, addr string, le int) int {
	ret, _ := com.dm.CallMethod("ReadDataToBin", hwnd, addr, le)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) ReadDouble(hwnd int, addr string) float64 {
	ret, _ := com.dm.CallMethod("ReadDouble", hwnd, addr)
	defer ret.Clear()
	return ret.Value().(float64)
}

func (com *Dmsoft) ReadDoubleAddr(hwnd int, addr int64) float64 {
	ret, _ := com.dm.CallMethod("ReadDoubleAddr", hwnd, addr)
	defer ret.Clear()
	return ret.Value().(float64)
}

func (com *Dmsoft) ReadFloat(hwnd int, addr string) float32 {
	ret, _ := com.dm.CallMethod("ReadFloat", hwnd, addr)
	defer ret.Clear()
	return ret.Value().(float32)
}

func (com *Dmsoft) ReadFloatAddr(hwnd int, addr int64) float32 {
	ret, _ := com.dm.CallMethod("ReadFloatAddr", hwnd, addr)
	defer ret.Clear()
	return ret.Value().(float32)
}

func (com *Dmsoft) ReadInt(hwnd int, addr string, typ int) int64 {
	ret, _ := com.dm.CallMethod("ReadInt", hwnd, addr, typ)
	defer ret.Clear()
	return ret.Val

}

func (com *Dmsoft) ReadIntAddr(hwnd int, addr int64, typ int) int64 {
	ret, _ := com.dm.CallMethod("ReadIntAddr", hwnd, addr, typ)
	defer ret.Clear()
	return ret.Val

}

func (com *Dmsoft) ReadString(hwnd int, addr string, typ, le int) string {
	ret, _ := com.dm.CallMethod("ReadString", hwnd, addr, typ, le)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) ReadStringAddr(hwnd int, addr int64, typ, le int) string {
	ret, _ := com.dm.CallMethod("ReadStringAddr", hwnd, addr, typ, le)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) SetMemoryFindResultToFile(file string) int {
	ret, _ := com.dm.CallMethod("SetMemoryFindResultToFile", file)
	defer ret.Clear()
	return int(ret.Val)

}

func (com *Dmsoft) SetMemoryHwndAsProcessId(en int) int {
	ret, _ := com.dm.CallMethod("SetMemoryHwndAsProcessId", en)
	defer ret.Clear()
	return int(ret.Val)
}

// string StringToData(value,type)
// long TerminateProcess(pid)

func (com *Dmsoft) VirtualAllocEx(hwnd int, addr int64, size, typ int) int64 {
	ret, _ := com.dm.CallMethod("VirtualAllocEx", hwnd, addr, size, typ)
	defer ret.Clear()
	return ret.Val
}

func (com *Dmsoft) VirtualFreeEx(hwnd int, addr int64) int {
	ret, _ := com.dm.CallMethod("VirtualFreeEx", hwnd, addr)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) VirtualProtectEx(hwnd int, addr int64, size, typ, oldProtect int) int {
	ret, _ := com.dm.CallMethod("VirtualProtectEx", hwnd, addr, size, typ, oldProtect)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) VirtualQueryEx(hwnd int, addr int64, pmbi int) string {
	ret, _ := com.dm.CallMethod("VirtualQueryEx", hwnd, addr, pmbi)
	defer ret.Clear()
	return ret.ToString()

}

func (com *Dmsoft) WriteData(hwnd int, addr, data string) int {
	ret, _ := com.dm.CallMethod("WriteData", hwnd, addr, data)
	defer ret.Clear()
	return int(ret.Val)

}

func (com *Dmsoft) WriteDataAddr(hwnd int, addr, data string) int {
	ret, _ := com.dm.CallMethod("WriteDataAddr", hwnd, addr, data)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteDataAddrFromBin(hwnd int, addr int64, data, le int) int {
	ret, _ := com.dm.CallMethod("WriteDataAddrFromBin", hwnd, addr, data, le)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteDataFromBin(hwnd int, addr string, data, le int) int {
	ret, _ := com.dm.CallMethod("WriteDataFromBin", hwnd, addr, data, le)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteDouble(hwnd int, addr string, v float64) int {
	ret, _ := com.dm.CallMethod("WriteDouble", hwnd, addr, v)
	defer ret.Clear()
	return int(ret.Val)

}

func (com *Dmsoft) WriteDoubleAddr(hwnd int, addr int64, v float64) int {
	ret, _ := com.dm.CallMethod("WriteDoubleAddr", hwnd, addr, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteFloat(hwnd int, addr string, v float32) int {
	ret, _ := com.dm.CallMethod("WriteFloat", hwnd, addr, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteFloatAddr(hwnd int, addr int64, v float32) int {
	ret, _ := com.dm.CallMethod("WriteFloatAddr", hwnd, addr, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteInt(hwnd int, addr string, typ int, v int64) int {
	ret, _ := com.dm.CallMethod("WriteInt", hwnd, addr, typ, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteIntAddr(hwnd int, addr int64, typ int, v int64) int {
	ret, _ := com.dm.CallMethod("WriteIntAddr", hwnd, addr, typ, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteString(hwnd int, addr string, typ int, v string) int {
	ret, _ := com.dm.CallMethod("WriteString", hwnd, addr, typ, v)
	defer ret.Clear()
	return int(ret.Val)

}
func (com *Dmsoft) WriteStringAddr(hwnd int, addr int64, typ int, v string) int {
	ret, _ := com.dm.CallMethod("WriteStringAddr", hwnd, addr, typ, v)
	defer ret.Clear()
	return int(ret.Val)

}
