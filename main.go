package main

import (
	// 	"example/hello_word/variables"

	"example/hello_word/files"
)

func main() {
	// estado, texto := variables.ConvertoTexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)

	//FUNCTION EJERCICIOS
	// numero, texto := ejercicios.ConvNumerico("100")
	// fmt.Println(numero)
	// fmt.Println(texto)

	//OBTENER EL SISTEMA OPERATIVO
	// if os := runtime.GOOS; os == "Linux." || os == "OS X." || os == "darwin" {
	// 	fmt.Println("Es un Mac o Linux")
	// } else {
	// 	fmt.Println("Esto es un windows")
	// }

	// fmt.Println(ejercicios.TablaMultiplicar())

	// files.GrabarTabla()
	files.LeoArchivo()
}
