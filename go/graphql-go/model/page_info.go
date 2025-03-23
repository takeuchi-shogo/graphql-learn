package model

type PageInfo struct {
	HasNextPage     bool
	HasPreviousPage bool
	StartCursor     string
	EndCursor       string
}

func NewPageInfo(hasNextPage, hasPreviousPage bool, startCursor, endCursor string) *PageInfo {
	return &PageInfo{
		HasNextPage:     hasNextPage,
		HasPreviousPage: hasPreviousPage,
	}
}
