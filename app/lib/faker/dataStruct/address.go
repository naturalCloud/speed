package dataStruct

//省份数据
type Province struct {
	Pk uint8 //省级码
	Pv string
}

//市级数据
type City struct {
	Pk uint8
	Ck uint16 //市级码
	Cv string
}

//县级数据
type Country struct {
	Ck  uint16
	Cyk uint16 // 县级码
	Cyv string
}

//镇级别数据
type Town struct {
	Cyk uint16
	Tk  uint16 //镇级别码
	Tv  string
}
