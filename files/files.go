package files

import (
	"bufio"
	"example/hello_word/ejercicios"
	"fmt"
	"os"
	// "bufio"
	// "ioutil"
)

var filename string = "./files/txt/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.TablaMultiplicar()

	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Printf("Error creating archive")
		return
	}

	fmt.Fprintln(archivo, texto)

	archivo.Close()

}

func SumaTabla() {
	var texto string = ejercicios.TablaMultiplicar()

	if Append(filename, texto) == false {
		fmt.Println("Error al concatenar contenido")
	}
	// if !Append(filename, texto) {
	// 	fmt.Println("Error al concatenar contenido")
	// }
}

func Append(file string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto + "\n")

	//AGREGAR SALTO DE LINEA

	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Hubo un error leyendo el archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

	archivo.Close()
}
