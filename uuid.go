package gochan

var gochanUUID int

func defualtUUID() int {
	gochanUUID++
	return gochanUUID
}
