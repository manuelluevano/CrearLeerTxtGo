package ejercicios

import (
	"strconv"
)

func ConvNumerico(valor string) (int, string) {

	//CONVERTIR EL STRING A INT

	num, err := strconv.Atoi(valor)

	//MANEJOR DE ERROR
	if err != nil {
		return 0, "Hubo un error: " + err.Error()
	}
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}
