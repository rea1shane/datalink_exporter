package api

type Overview struct {
	AaData          []InnerOverview `json:"aaData"`
	Draw            int             `json:"draw"`
	Length          int             `json:"length"`
	PageNum         int             `json:"pageNum"`
	PageSize        int             `json:"pageSize"`
	Pages           int             `json:"pages"`
	RecordsFiltered int             `json:"recordsFiltered"`
	RecordsTotal    int             `json:"recordsTotal"`
	Size            int             `json:"size"`
	Start           int             `json:"start"`
}

type InnerOverview struct {
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

type Group struct {
	AaData          []InnerGroup `json:"aaData"`
	Draw            int          `json:"draw"`
	Length          int          `json:"length"`
	PageNum         int          `json:"pageNum"`
	PageSize        int          `json:"pageSize"`
	Pages           int          `json:"pages"`
	RecordsFiltered int          `json:"recordsFiltered"`
	RecordsTotal    int          `json:"recordsTotal"`
	Size            int          `json:"size"`
	Start           int          `json:"start"`
}

type InnerGroup struct {
	CreateTime         int64  `json:"createTime"`
	GenerationID       int    `json:"generationId"`
	GroupDesc          string `json:"groupDesc"`
	GroupName          string `json:"groupName"`
	GroupState         string `json:"groupState"`
	ID                 int    `json:"id"`
	LastRebalancedTime string `json:"lastRebalancedTime"`
	ModifyTime         int64  `json:"modifyTime"`
}

type Worker struct {
	AaData          []InnerWorker `json:"aaData"`
	Draw            int           `json:"draw"`
	Length          int           `json:"length"`
	PageNum         int           `json:"pageNum"`
	PageSize        int           `json:"pageSize"`
	Pages           int           `json:"pages"`
	RecordsFiltered int           `json:"recordsFiltered"`
	RecordsTotal    int           `json:"recordsTotal"`
	Size            int           `json:"size"`
	Start           int           `json:"start"`
}

type InnerWorker struct {
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
