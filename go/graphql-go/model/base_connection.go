package model

type BaseConnection struct {
	Edges    []*BaseEdge
	PageInfo *PageInfo
}

type BaseEdge struct {
	Cursor string
	Node   any
}

func NewBaseConnection(edges []*BaseEdge, pageInfo *PageInfo) *BaseConnection {
	return &BaseConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}
}
