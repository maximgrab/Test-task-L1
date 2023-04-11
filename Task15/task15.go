/*
var justString string

	func someFunc(){
		v:=createHugeString(1<<10)
		justString=v[:100]
	}

	func main(){
		someFunc()
	}

	1. Глобальные переменные - bad practice
	2. Лишние аллокации и копирования
	3. Взятие среза корректно только при размере символа стринга в один байт, например при использовании кириллицы будем иметь размер символа в 2 байта
	и в justString скопируются только первые 50 символов, вперемешку в одном стринге разные символы могут иметь разный размер
*/
package main

var justString []rune

func someFunc() {
	justString = []rune(createHugeString(1 << 10))
	justString = justString[:100]
}
func main() {
	someFunc()
}
