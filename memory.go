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

// string ReadDataAddr(hwnd,addr,len)
// long ReadDataAddrToBin(hwnd,addr,len)
// long ReadDataToBin(hwnd,addr,len)
// double ReadDouble(hwnd,addr)
// double ReadDoubleAddr(hwnd,addr)
// float ReadFloat(hwnd,addr)
// float ReadFloatAddr(hwnd,addr)
// LONGLONG ReadInt(hwnd,addr,type)
// LONGLONG ReadIntAddr(hwnd,addr,type)
// string ReadString(hwnd,addr,type,len)
// string ReadStringAddr(hwnd,addr,type,len)
// long SetMemoryFindResultToFile(file)
func (com *Dmsoft) SetMemoryHwndAsProcessId(en int) int {
	ret, _ := com.dm.CallMethod("SetMemoryHwndAsProcessId", en)
	defer ret.Clear()
	return int(ret.Val)
}

// string StringToData(value,type)
// long TerminateProcess(pid)
// LONGLONG VirtualAllocEx(hwnd,addr,size,type)
// long VirtualFreeEx(hwnd,addr)
// long VirtualProtectEx(hwnd,addr,size,type,old_protect)
// string VirtualQueryEx(hwnd,addr,pmbi)
func (com *Dmsoft) WriteData(hwnd int, addr, data string) int {
	ret, _ := com.dm.CallMethod("WriteData", hwnd, addr, data)
	defer ret.Clear()
	return int(ret.Val)
}

// long WriteDataAddr(hwnd,addr,data)
// long WriteDataAddrFromBin(hwnd,addr,data,len)
// long WriteDataFromBin(hwnd,addr,data,len)
// long WriteDouble(hwnd,addr,v)
// long WriteDoubleAddr(hwnd,addr,v)
// long WriteFloat(hwnd,addr,v)
// long WriteFloatAddr(hwnd,addr,v)
// long WriteInt(hwnd,addr,type,v)
// long WriteIntAddr(hwnd,addr,type,v)
// long WriteString(hwnd,addr,type,v)
// long WriteStringAddr(hwnd,addr,type,v)
