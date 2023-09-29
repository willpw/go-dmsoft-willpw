package dmsoft

import "github.com/go-ole/go-ole"

/*
string AiYoloDetectObjects(x1, y1, x2, y2,prob,iou)
long AiYoloDetectObjectsToDataBmp(x1, y1, x2, y2,prob,iou,data,size,mode)
long AiYoloDetectObjectsToFile(x1, y1, x2, y2,prob,iou,file,mode)
long AiYoloFreeModel(index)
string AiYoloObjectsToString(objects)
long AiYoloSetModel(index,file,pwd)
long AiYoloSetModelMemory(index,data,size,pwd)
long AiYoloSetVersion(ver)
string AiYoloSortsObjects(objects,height)
long AiYoloUseModel(index)
long LoadAi(file)
long LoadAiMemory(data,size)

*/

func (com *Dmsoft) AiYoloDetectObjects(x1, y1, x2, y2 int, prob, iou float32) string {
	ret, _ := com.dm.CallMethod(DllM["AiYoloDetectObjects"], x1, y1, x2, y2, prob, iou)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) AiYoloDetectObjectsToDataBmp(x1, y1, x2, y2 int, prob, iou float32, data, size *int, mode int) int {
	d := ole.NewVariant(ole.VT_I4, int64(*data))
	s := ole.NewVariant(ole.VT_I4, int64(*size))
	ret, _ := com.dm.CallMethod(DllM["AiYoloDetectObjectsToDataBmp"], x1, y1, x2, y2, prob, iou, &d, &s, mode)
	*data = int(d.Val)
	*size = int(s.Val)
	d.Clear()
	s.Clear()
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AiYoloDetectObjectsToFile(x1, y1, x2, y2 int, prob, iou float32, file string, mode int) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloDetectObjectsToFile"], x1, y1, x2, y2, prob, iou, file, mode)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AiYoloFreeModel(index int) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloFreeModel"], index)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AiYoloObjectsToString(objects string) string {
	ret, _ := com.dm.CallMethod(DllM["AiYoloObjectsToString"], objects)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) AiYoloSetModel(index int, file, pwd string) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloSetModel"], index, file, pwd)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AiYoloSetModelMemory(index, data, size int, pwd string) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloSetModelMemory"], index, data, size, pwd)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AiYoloSetVersion(ver string) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloSetVersion"], ver)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) AiYoloSortsObjects(objects string, height int) string {
	ret, _ := com.dm.CallMethod(DllM["AiYoloSortsObjects"], objects, height)
	defer ret.Clear()
	return ret.ToString()
}

func (com *Dmsoft) AiYoloUseModel(index int) int {
	ret, _ := com.dm.CallMethod(DllM["AiYoloUseModel"], index)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) LoadAi(file string) int {
	ret, _ := com.dm.CallMethod(DllM["LoadAi"], file)
	defer ret.Clear()
	return int(ret.Val)
}

func (com *Dmsoft) LoadAiMemory(data, size int) int {
	ret, _ := com.dm.CallMethod(DllM["LoadAiMemory"], data, size)
	defer ret.Clear()
	return int(ret.Val)
}
