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
