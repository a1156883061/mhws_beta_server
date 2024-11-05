package backend

type Hunter struct {
	HunterSaveList []HunterSave `msgpack:"HunterSaveList"`
	UsingSaveSlot  int          `msgpack:"UsingSaveSlot"`
}

type HunterSave struct {
	HunterId   string `msgpack:"HunterId"`
	HunterName string `msgpack:"HunterName"`
	OtomoName  string `msgpack:"OtomoName"`
	SaveSlot   int    `msgpack:"SaveSlot"`
}
