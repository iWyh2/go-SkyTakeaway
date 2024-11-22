package pageResult

// Page 分页查询数据装载模型
type Page[T any] struct {
	Total   int64          `json:"total"`
	Records PageRecords[T] `json:"records"`
}

// PageRecords 分页查询数据记录模型
type PageRecords[T any] []T
