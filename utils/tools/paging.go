package tools

import "RudderMaster/settings"

// GetPageSize 获取页面大小
func GetPageSize(size int) int {
	defaultSize := settings.Config.Application.PageSize
	maxSize := settings.Config.Application.PageMaxSize
	if size <= 0 {
		return defaultSize
	} else if size > maxSize {
		return maxSize
	}
	return size
}

// GetPageNum 获取页码
func GetPageNum(current, size int) int {
	offset := 0
	if current > 0 {
		offset = (current - 1) * size
	}
	return offset
}
