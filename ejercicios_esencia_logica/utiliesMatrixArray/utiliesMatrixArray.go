package utiliesMatrixArray

import ( 
	// "fmt"
	"math"
	// "sort"
	// "strconv"
	//"strings"
	"../arrayMetodos"
)

func unionNumber(before int,after int)int{
	return before + after;
}
func SumArray(number[]int)int{
	return arrayMetodos.ReduceInt(number,unionNumber);
}

type PowerMatrixMapArray struct{
	mapNumber func(int)int
}
func(powerMatrixMapArray*PowerMatrixMapArray)MapArray(numbers []int)[]int{
	var arrayConvertido []int;
	for number := 0;number < len(numbers);number++ {
		arrayConvertido = append(arrayConvertido,powerMatrixMapArray.mapNumber(numbers[number]));
	}
	return arrayConvertido;
}
func(powerMatrixMapArray*PowerMatrixMapArray)MapMatrix(numbers [][]int)[][]int{
	var colectionNumber  [][]int;
	for _,element := range numbers{		        
		colectionNumber = append(colectionNumber,powerMatrixMapArray.MapArray(element));
	}
	return colectionNumber;
}
func NewPowerMatrixMapArray(mapNumber func(int)int)PowerMatrixMapArray{
	 return PowerMatrixMapArray{mapNumber:mapNumber};
}

type PowerMatrixArray struct{
	shearNumber func(int)bool
}
func(powerMatrixArray*PowerMatrixArray)Some(numbers []int)bool{
 	return	arrayMetodos.Some(numbers,powerMatrixArray.shearNumber);
}
func(powerMatrixArray*PowerMatrixArray) Every(numbers []int) bool{
	return arrayMetodos.Every(numbers,powerMatrixArray.shearNumber);
}
func (powerMatrixArray*PowerMatrixArray) Filter(numbers []int) []int{
	return arrayMetodos.Filter(numbers,powerMatrixArray.shearNumber);	
}
func(powerMatrixArray*PowerMatrixArray) FilterMatrix(numbers [][]int)[][]int{
	return arrayMetodos.FilterMatrix(numbers,powerMatrixArray.shearNumber);	
}
func(powerMatrixArray*PowerMatrixArray) FilterRowsMatrix(numbers [][]int)[][]int{	
	return arrayMetodos.FilterRowsMatrix(numbers,powerMatrixArray.shearNumber);		
}
func(powerMatrixArray*PowerMatrixArray) FindIndex(numbers []int)int{
	return arrayMetodos.FindIndex(numbers,powerMatrixArray.shearNumber);		
}
func(powerMatrixArray*PowerMatrixArray)FindArrayIndex(numbers []int )[]int{
	return arrayMetodos.FindArrayIndex(numbers,powerMatrixArray.shearNumber);			
}
func(powerMatrixArray*PowerMatrixArray) SumNumbersCondition(numbers []int)int{
	numberfilter := arrayMetodos.Filter(numbers,powerMatrixArray.shearNumber);
	return SumArray(numberfilter);
}
func NewPowerMatrixArray(shearNumber func(int)bool)PowerMatrixArray{
	return PowerMatrixArray{shearNumber:shearNumber};
}

func  ifPrimeFloat(numero float64 )bool{
	var respuesta float64;
	var esPrimo bool;
	if numero != 2{
		respuesta = math.Mod( math.Pow(2,numero-1) , numero);
	}else {
		respuesta = 1;
	}
	if respuesta == 1 {
		esPrimo =  true;
	}
	return esPrimo;
}
func ifPrimeInt(numero int )bool{
	return ifPrimeFloat(float64(numero));
}
type NewPrime struct{
	powerMatrixArray PowerMatrixArray
}
func(NewPrime*NewPrime)Prime()PowerMatrixArray{
	return PowerMatrixArray{shearNumber:ifPrimeInt};
}

func evenNumber(number int)bool{
	return number % 2 == 0;
}
func NewEvenNumber()PowerMatrixArray{
	return PowerMatrixArray{shearNumber:evenNumber};	
}