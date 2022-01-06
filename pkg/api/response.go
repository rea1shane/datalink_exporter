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
