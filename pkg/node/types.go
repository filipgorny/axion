package node

type DataType int

const (
	String DataType = iota
	Integer
	Boolean
	Float
	DateTime
	Object
)
