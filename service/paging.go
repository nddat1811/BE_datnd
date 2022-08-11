package service


func getPageSize(pageSize uint) int {
	if pageSize < 1 || pageSize > 20 {
		return 20
	}
	return int(pageSize)
}

func CalcPagination(page, pageSize uint) (limit, offset int) {
	limit = getPageSize(pageSize)
	offset = int(pageSize * (page - 1))

	return limit, offset
}
