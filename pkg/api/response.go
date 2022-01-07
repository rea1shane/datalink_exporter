package api

type Overview struct {
	AaData          []AaData `json:"aaData"`
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

type AaData struct {
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
