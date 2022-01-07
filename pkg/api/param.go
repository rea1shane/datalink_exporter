package api

type DoLoginParam struct {
	ServerUrl  string
	LoginEmail string
	Password   string
}

type HomeCountParam struct {
	ServerUrl string
}

type HomeStatisParam struct {
	ServerUrl string
}

type GroupInitGroupParam struct {
	ServerUrl string
	Start     int
	Length    int
}
