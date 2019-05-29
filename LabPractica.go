/*	ALUMNO: Juan Carlos Paye Rivera
        
	CURSO: Paralelismo y Sistemas Distribuidos
	NOTA: Al ingresar '0' (cero) se termina la entrada de enteros
	*/
package main

import (
	"fmt"
	"sync"
)
var cant sync.WaitGroup
func Dividir(enteros []int) {
	tamp := len(enteros)/4
	
	go Ordenar(enteros[:tamp])
	go Ordenar(enteros[tamp:tamp*2])
	go Ordenar(enteros[tamp*2:tamp*3])
	go Ordenar(enteros[tamp*3:])
	Ordenar(enteros)

}
func Ordenar(enteros []int) {
	tamp := len(enteros)
	for i := 0; i < tamp; i++ {
		for j := tamp - 1; j >= i+1; j-- {
			if enteros[j] < enteros[j-1] {
				enteros[j], enteros[j-1] = enteros[j-1], enteros[j]
			}
		}
	}
	fmt.Println(enteros)
	cant.Done()
}
func main() {
	var numero int
	i := 0
	

	enteros:= [20] int{}	
	// --------------- INSERTAMOS NUMEROS ---------------
	fmt.Println("ingresar enteros: ")
		for{
		fmt.Scanf("%d\n",&numero)
		if(numero == 0){
			break
		}else{

			enteros[i] = numero
			i++
		}
	}
	
	cant.Add(5)
	fmt.Println("Matriz: ",enteros, len(enteros))
	temp:=enteros[:i]
	Dividir(temp)

	cant.Wait()
}
