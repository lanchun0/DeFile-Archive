package entity

type Permission byte
type PermissionLevel byte

const (
	Administrator Permission = 0b00000001
	Writer        Permission = 0b00000010
	Reader        Permission = 0b00000100
)
const (
	L_0 PermissionLevel = iota
	L_1
	L_2
	L_3
	L_4
)
