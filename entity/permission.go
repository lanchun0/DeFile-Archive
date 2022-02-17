package entity

type Permission uint8
type PermissionLevel uint8

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
