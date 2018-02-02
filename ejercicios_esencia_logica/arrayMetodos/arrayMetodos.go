// agregar y mostrar datos
package arrayMetodos


func  MapTransformDataString(numeros []int, function func(int)string) []string {
	var arrayConvertido []string;
	for numero := 0;numero < len(numeros);numero++ {
		arrayConvertido = append(arrayConvertido,string(function(numeros[numero])));
	}
	return arrayConvertido;
}
func  MapTransformDataFloat32(numeros []int, function func(int)float32) []float32 {
	var arrayConvertido []float32;
	for numero := 0;numero < len(numeros);numero++ {
		arrayConvertido = append(arrayConvertido,function(numeros[numero]));
	}
	return arrayConvertido;
}
func MapTransformDataConvertToMatrix(numbers []int,funcion func(int)[]int)[][]int{
	var colectionNumber = make([][]int,len(numbers));
	for numero := 0;numero < len(numbers);numero++ {		
		colectionNumber[numero] = funcion(numbers[numero]);
	}	
	return colectionNumber;
}
func  MapTransformData(numeros []int, function func(int)int) []int {
	var arrayConvertido []int;
	for numero := 0;numero < len(numeros);numero++ {
		arrayConvertido = append(arrayConvertido,function(numeros[numero]));
	}
	return arrayConvertido;
}
func  MapTransformMatrixDataInIntArray(numbers [][]int,funcion func([]int)int)[]int{
	var colectionNumber = make([]int,len(numbers));
	for numero := 0;numero < len(numbers);numero++ {		
		colectionNumber[numero] = funcion(numbers[numero]);
	}	
	return colectionNumber;
}
func  MapTransformMatrixDataInMatrix(numbers [][]int,funcion func([]int)[]int)[][]int{
	var colectionNumber = make([][]int,len(numbers));
	for numero := 0;numero < len(numbers);numero++ {		
		colectionNumber[numero] = funcion(numbers[numero]);
	}	
	return colectionNumber;
}
func Some(numeros []int, function func(int)bool) bool {
	var resultado bool;
	for _,element := range numeros{		        
		resultado = function(element);
		if resultado {
			break;
		}
	}
	return resultado;
}
func SomeFloat32(numeros [] float32, function func( float32)bool) bool {
	var resultado bool;
	for _,element := range numeros{		        
		resultado = function(element);
		if resultado {
			break;
		}
	}
	return resultado;
}
func Every(numeros []int, function func(int)bool) bool{
	var resultado bool;
	for _,element := range numeros{		        
		resultado = function(element);
		if !resultado  {
			break;
		}
	}
	return resultado;
}
func TraverseArrayWorkWithTwoData(numeros []int, accion func(int,int) int) int{
	var resultado  int;
	for _,element := range numeros{
		resultado = accion(resultado,element);// toca arreglar
	}
	return resultado;
}
func ReduceInt(numeros []int, accion func(int,int) int) int{
	var resultado  int;
	for _,element := range numeros{
		resultado = accion(resultado,element);
	}
	return resultado;
}
func ReduceIntT(numeros []int, accion func(int,int) int) int{
	var resultado int;
	if len(numeros)>0 {
		resultado  = numeros[0];
		for _,element := range numeros{
			resultado = accion(resultado,element);
		}
	}
	return resultado;
}
func ReduceRightInt(numeros []int, accion func(int,int) int )int {
	var resultado  int;
	for numero := len(numeros)-1; numero >= 0; numero-- {
		resultado = accion(resultado,numeros[numero]);
	}
	return resultado;
}
func ReduceIntS(numeros []int, accion func(int) int) int{
	resultado  := numeros[0];
   for _,number := range numeros{
		resultado = accion(number);
   }
   return resultado;
}
func Filter(numeros []int, accion func(int) bool) []int{
	var resultadoArray []int;	
	for numero := 0; numero < len(numeros); numero++ {
		if accion(numeros[numero]){
			resultadoArray = append(resultadoArray,numeros[numero]);
		}
	}
	return resultadoArray;
}
func FilterMatrix(numbers [][]int,funcion func(int)bool)[][]int{
	var colectionNumber  [][]int;
	for numero := 0; numero < len(numbers); numero++ {
		if Some(numbers[numero],funcion){
			colectionNumber = append(colectionNumber,numbers[numero]);
		}
	}
	return colectionNumber;
}
func FilterRowsMatrix(numbers [][]int,funcion func(int)bool)[][]int{	
	var colectionNumber  [][]int;
	for numero := 0; numero < len(numbers); numero++ {
		colectionNumber = append(colectionNumber,Filter(numbers[numero],funcion));
	}
	return colectionNumber;
}
func CompareAllTheArray( numeros []int, function func(int,int)bool)[]int{
	var resultadoArray 	[]int;
	for posicion,element := range numeros{
		for posisionComparado,elementComparado := range numeros{	
			if posicion != posisionComparado {
				if function( element,elementComparado ){
					resultadoArray = append(resultadoArray,element);
				}
			}
	   }
	} 
	return resultadoArray;
}
func FindIndex(array []int, function func(int)bool )int{
	var position int;
	for numero := 0; numero < len(array); numero++ {
		if function(array[numero]) {
			position = numero;
			break
		}
	}
	return position+1;
}

func FindArrayIndex(array []int, function func(int)bool )[]int{
	var positions []int;
	for index := 0; index < len(array); index++ {
		if function(array[index]) {
			positions = append(positions,index);					
		}
	}
	return positions;
}

func ApplyFindIndexRowColInIndex(matrix [][]int,function func(int)bool)[][]int{
	var colectionIndex [][]int;
	for row, element := range matrix {
		indexCol := FindArrayIndex(element,function);		
		for _, elementCol := range indexCol { 
			colectionIndex = append(colectionIndex,[]int{row,elementCol});				
		}	
	}
	return colectionIndex
}

func ApplyRowCol(matrix [][]int,function func([]int)int)[]int{
	var colectionIndex []int;
	for _, element := range matrix {
		if len(element) > 0 {	
			colectionIndex = append(colectionIndex,function(element ));
		}		
	}
	return colectionIndex;
}

func FindArrayIndexArgument(array []int, functions []func(int)bool )[]int{
	var positions []int;
	for index, element := range array {
		var noSelection bool;
		for _, function := range functions {
			if !function(element) {
				noSelection = true;
				break;
			}
		}
		if !noSelection {
			positions = append(positions,index);					
		}
	}
	return positions;
}

func ApplyFindRowColInIndexArgument(matrix [][]int,function ...func(int)bool)[][]int{
	var colectionIndex [][]int;
	for row, element := range matrix {
		indexCol := FindArrayIndexArgument(element,function);		
		for _, elementCol := range indexCol { 
			colectionIndex = append(colectionIndex,[]int{row,elementCol});				
		}	
	}
	return colectionIndex
}
func ApplyFindRowIndexArgument(matrix [][]int,function ...func(int)bool)[]int{
	var colectionIndex []int;
	for row, element := range matrix {
		if len(FindArrayIndexArgument(element,function)) > 0{
			colectionIndex = append(colectionIndex,row);				
		}	
	}
	return colectionIndex
}
func ApplyFindColInIndexArgument(matrix [][]int,function ...func(int)bool)[]int{
	var colectionIndex []int;
	for _, element := range matrix {
		indexCol := FindArrayIndexArgument(element,function);		
		for _, elementCol := range indexCol { 
			colectionIndex = append(colectionIndex,elementCol);				
		}	
	}
	return colectionIndex
}

func FindArrayArgument(array []int, functions []func(int)bool )[]int{
	var colection []int;
	for _, element := range array {
		var noSelection bool;
		for _, function := range functions {
			if !function(element) {
				noSelection = true;
				break;
			}
		}
		if !noSelection {
			colection = append(colection,element);					
		}
	}
	return colection;
}

func ApplyFindRowColInArgument(matrix [][]int,function ...func(int)bool)[][]int{
	var colection [][]int;
	for _, element := range matrix {
		colection = append(colection,FindArrayArgument(element,function));				
	}
	return colection
}