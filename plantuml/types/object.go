package types

type objectType string

const (
	Object = "object"
	Map    = "map"
)

// примеры: https://www.plantuml.com/plantuml/umla/fPB1IWCn48RFpLCCFIvUr5lHqfg3Y7eguZcsQHlOJKfsUYWYM2zU18--HWKBgrPzXTatSf8MMKhrO2aDamdpVn_3jd9K7iTPAbBTHq07tWo623tZTHUDiq9aQWHIZGogL1gSmBMeNPp2OHFEdEqvoxTEErmlLJuqTe2XTQnzfkmaj8-qJuqLDsAPL3_d8qcai-ZZUoKwIUgWSf3o_uo8uAC1p3auiQW7skz01psFDMFyEDS-3uJTCDfEJLUBsbwCSLUYjX-bXeeJTKYq3lLgAQw_6XRqH2yqBwVb0w31L3V-WWgQ0HMqA1-1ldZgbkQquFy7_pxf5SefpRWHopikd-cTYlAUtgXeB852snu4aNKTAJUfn1N3Bpk6P2iXHbPirdGgd8qxghwpHR65HPNquzFYYh_bRm00

type ObjectModel struct {
	name   string
	title  string
	fields []string
	Type   string
}

type MapModel struct {
	name   string
	title  string
	fields map[string]int
	Type   string
}

type JsonModel struct {
	name  string
	title string
	json  string //json type
}
