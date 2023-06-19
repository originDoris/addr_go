package main

import "getAddr/generate/autoCode"

//go:generate go run main.go
func main() {
	autoCode.AutoAreaMap()
}
