package main
//go:generate go run gen_ast.go

func main() {
	genFunc("Hello, World!")
}

func genFunc(t interface{}) {
	println(
