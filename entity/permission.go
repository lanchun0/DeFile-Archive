package entity

type Permission byte
type PermissionLevel byte

const (
	Owner Permission = 0b10000000 >> iota
	Administrator
	Writer
	Reader
)
const (
	L_0 PermissionLevel = iota
	L_1
	L_2
	L_3
	L_4
)
