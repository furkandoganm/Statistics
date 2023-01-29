package model

//ConnectionString

type Connection struct {
	ConnectionString []ConnectionString `json:"connectionString"`
}

type ConnectionString struct {
	Name    string `json:"name"`
	CString string `json:"cString"`
}

//RegionCode

type RegionCodeList struct {
	List []RegionCode `json:"list"`
}

type RegionCode struct {
	Country string `json:"country"`
	Code    string `json:"code"`
}

// channels and video Ids

type Ids struct {
	Ids []Id `json:"ids"`
}

type Id struct {
	Id string `json:"id"`
}
