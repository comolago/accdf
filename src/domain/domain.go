package domain

type Domain interface {
	GetId() string
	SetId(i string)
	GetName() string
	SetName(n string)
}
