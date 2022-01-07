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

type WorkerInitWorkerParam struct {
	ServerUrl string
	Start     int
	Length    int
}

type MysqlTaskMysqlTaskDatasParam struct {
	ServerUrl string
	Start     int
	Length    int
}

type HbaseTaskInitHbaseTaskListParam struct {
	ServerUrl string
	Start     int
	Length    int
}

