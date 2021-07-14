package helpers

import (
	"github.com/jorgec815/models"
)


func CalcularDefinitiva(id int, nota1 float64, nota2 float64, nota3 float64) (def float64){

	def = ((nota1*0.35)+(nota2*0.35)+(nota3*0.3))

	return def
}