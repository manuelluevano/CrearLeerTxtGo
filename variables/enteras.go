package variables

import (
	"fmt"
	"strconv"
)

func MuestroEnteros() {
	var valor1 int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun = ", valor1)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)

}

func ConvertoTexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero)
	return true, texto
}
