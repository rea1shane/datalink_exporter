package api

type Overviews struct {
	AaData          []Overview `json:"aaData"`
	Draw            int        `json:"draw"`
	Length          int        `json:"length"`
	PageNum         int        `json:"pageNum"`
	PageSize        int        `json:"pageSize"`
	Pages           int        `json:"pages"`
	RecordsFiltered int        `json:"recordsFiltered"`
	RecordsTotal    int        `json:"recordsTotal"`
	Size            int        `json:"size"`
	Start           int        `json:"start"`
}

type Overview struct {
	GroupCount   int `json:"groupCount"`
	MappingCount int `json:"mappingCount"`
	MsCount      int `json:"msCount"`
	TaskCount    int `json:"taskCount"`
	WorkerCount  int `json:"workerCount"`
}

type Statis struct {
	MsCount                int       `json:"msCount"`
	TaskNameList           []string  `json:"taskNameList"`
	WorkerNetTrafficList   []float64 `json:"workerNetTrafficList"`
	MappingCount           int       `json:"mappingCount"`
	WorkerJvmUsedList      []float64 `json:"workerJvmUsedList"`
	WorkerYoungGCCountList []float64 `json:"workerYoungGCCountList"`
	WorkerNameListNet      []string  `json:"workerNameListNet"`
	WorkerCount            int       `json:"workerCount"`
	TaskCount              int       `json:"taskCount"`
	TaskRecordsList        []int     `json:"taskRecordsList"`
	GroupCount             int       `json:"groupCount"`
	TaskSizeList           []int     `json:"taskSizeList"`
	TaskNameListDelay      []string  `json:"taskNameListDelay"`
	TaskDelayTimeList      []int     `json:"taskDelayTimeList"`
	WorkerNameListGC       []string  `json:"workerNameListGC"`
	WorkerNameList         []string  `json:"workerNameList"`
}

type Groups struct {
	AaData          []Group `json:"aaData"`
	Draw            int     `json:"draw"`
	Length          int     `json:"length"`
	PageNum         int     `json:"pageNum"`
	PageSize        int     `json:"pageSize"`
	Pages           int     `json:"pages"`
	RecordsFiltered int     `json:"recordsFiltered"`
	RecordsTotal    int     `json:"recordsTotal"`
	Size            int     `json:"size"`
	Start           int     `json:"start"`
}

type Group struct {
	CreateTime         int64  `json:"createTime"`
	GenerationID       int    `json:"generationId"`
	GroupDesc          string `json:"groupDesc"`
	GroupName          string `json:"groupName"`
	GroupState         string `json:"groupState"`
	ID                 int    `json:"id"`
	LastRebalancedTime string `json:"lastRebalancedTime"`
	ModifyTime         int64  `json:"modifyTime"`
}

type Workers struct {
	AaData          []Worker `json:"aaData"`
	Draw            int      `json:"draw"`
	Length          int      `json:"length"`
	PageNum         int      `json:"pageNum"`
	PageSize        int      `json:"pageSize"`
	Pages           int      `json:"pages"`
	RecordsFiltered int      `json:"recordsFiltered"`
	RecordsTotal    int      `json:"recordsTotal"`
	Size            int      `json:"size"`
	Start           int      `json:"start"`
}

type Worker struct {
	CreateTime    int64  `json:"createTime"`
	GroupName     string `json:"groupName"`
	ID            int    `json:"id"`
	ModifyTime    int64  `json:"modifyTime"`
	RestPort      int    `json:"restPort"`
	StartTime     string `json:"startTime"`
	WorkerAddress string `json:"workerAddress"`
	WorkerName    string `json:"workerName"`
	WorkerState   string `json:"workerState"`
}

type MysqlTasks struct {
	AaData          []MysqlTask `json:"aaData"`
	Draw            int         `json:"draw"`
	Length          int         `json:"length"`
	PageNum         int         `json:"pageNum"`
	PageSize        int         `json:"pageSize"`
	Pages           int         `json:"pages"`
	RecordsFiltered int         `json:"recordsFiltered"`
	RecordsTotal    int         `json:"recordsTotal"`
	Size            int         `json:"size"`
	Start           int         `json:"start"`
}

type MysqlTask struct {
	CurrentLogFile                      string `json:"currentLogFile"`
	CurrentLogPosition                  int    `json:"currentLogPosition"`
	CurrentTimeStamp                    string `json:"currentTimeStamp"`
	Detail                              string `json:"detail"`
	GroupID                             int    `json:"groupId"`
	ID                                  int    `json:"id"`
	LatestEffectSyncLogFileName         string `json:"latestEffectSyncLogFileName"`
	LatestEffectSyncLogFileOffset       string `json:"latestEffectSyncLogFileOffset"`
	ListenedState                       string `json:"listenedState"`
	ReaderIP                            string `json:"readerIp"`
	ShadowCurrentTimeStamp              string `json:"shadowCurrentTimeStamp"`
	ShadowLatestEffectSyncLogFileName   string `json:"shadowLatestEffectSyncLogFileName"`
	ShadowLatestEffectSyncLogFileOffset string `json:"shadowLatestEffectSyncLogFileOffset"`
	StartTime                           string `json:"startTime"`
	TargetState                         string `json:"targetState"`
	TaskDesc                            string `json:"taskDesc"`
	TaskName                            string `json:"taskName"`
	TaskSyncStatus                      string `json:"taskSyncStatus"`
	WorkerID                            int    `json:"workerId"`
}
