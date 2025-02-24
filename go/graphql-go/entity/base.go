package entity

type BaseNode struct {
	id string
}

func (b *BaseNode) ID() string {
	return b.id
}
