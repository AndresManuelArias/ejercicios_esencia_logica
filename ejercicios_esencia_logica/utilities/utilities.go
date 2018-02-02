package utilities

import ( "fmt"
	//"math"
	// "sort"
	// "strconv"
	"strings"
	"../arrayMetodos"
	"../condicionales"
	)

func ArrayToInt(numbers []int) int {
	var unionNumber int;
	for number := 0 ;number < len(numbers);number++{
		unionNumber = unionNumber *10 + numbers[number];
	}	
	return unionNumber;
}

func  ContarDigitos(numbers int) int{
	condition := condicionales.ExportarObjetosCondicionales();
	return condition.Ejercicio2.ContarDigitos(numbers);
}
func ArrayDigitosNumero(numbers int) []int{
	condition := condicionales.ExportarObjetosCondicionales();
	return condition.Ejercicio4.ArrayDigitosNumero(numbers)
}
func TransformNumber(numbers []int, function func( []int)[][]int )[]int{
	arrayNumber := function(numbers);
	return arrayMetodos.MapTransformMatrixDataInIntArray(arrayNumber,ArrayToInt);
}	
func ArrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
func ArrayToStringFloat32(a []float32, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
func MatrixToString( matrix [][]int,delimColumn string )string{
	var colectionNumber [] string;	
	for _,element := range matrix{		
		colectionNumber = append(colectionNumber,ArrayToString(element,delimColumn));
	}
	return strings.Join(colectionNumber, " | ");
}
func NumerIfPrime(number int)bool{
	condition := condicionales.ExportarObjetosCondicionales();
	return condition.Ejercicio6.DeterminarSiPrimoInt(number);
}