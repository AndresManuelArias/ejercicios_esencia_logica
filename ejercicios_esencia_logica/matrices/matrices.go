package matrices
// ejemplo :https://docslide.net/documents/ejercicios-resueltos-de-algoritmos-55b9fbab8d6b9.html
import ( "fmt"
// "math"
// "sort"
"strconv"
// "strings"
"../controlarFuncionamientoPrograma"
// "../agregarMostrarValores"
"../arrayMetodos"
"../ejecutarEjercicio"
"../utilities"
"../utilitiesEsenciaLogica"
"../utiliesMatrixArray"

)
func ejercicios(){
	var numeroPregunta string
	posicion := [...] interface{} {	
		ejecutarEjercicio1,
		ejecutarEjercicio2,
		ejecutarEjercicio3,
		ejecutarEjercicio4,
		ejecutarEjercicio5,
		ejecutarEjercicio6,
		ejecutarEjercicio7,
		ejecutarEjercicio8,
		ejecutarEjercicio9,
		ejecutarEjercicio10,
		ejecutarEjercicio11,
		ejecutarEjercicio12,
		ejecutarEjercicio13,
		ejecutarEjercicio14,
		ejecutarEjercicio15,
		ejecutarEjercicio16,
		ejecutarEjercicio17,
		ejecutarEjercicio18,
		ejecutarEjercicio19,
		ejecutarEjercicio20,
		ejecutarEjercicio21,
		ejecutarEjercicio22,
		ejecutarEjercicio23,
		ejecutarEjercicio24,
	}
	fmt.Println("Escoja el número de la pregunta; que son el 1 al ",len(posicion)) ;
	/* 
		hacer una funcion 
		que ejecute la funcion indifinidamente y esta se detenga cuando oprima una tecla especial
	*/
	fmt.Scanf("%s\n", &numeroPregunta);
	numeroPreguntaEntero,_ := strconv.Atoi(numeroPregunta)
	if numeroPreguntaEntero <= len(posicion) && numeroPreguntaEntero > 0 {
		posicion[numeroPreguntaEntero-1].(func())();
	}else {
		fmt.Println("el  valor ",numeroPregunta,"  que digito, no se encuentra como ejercicio");		
	}
}
/*
Notas Aclaratorias:
a.
b.
En los siguientes enunciados cuando se diga Leer una matriz mxn entera significa leer mxn
datos enteros y almacenarlos en m filas y n columnas para cualquier valor positivo de m y de n.
Cuando el enunciado diga Posición Exacta se refiere a a
l fila y a la columna del dato
especificado.
1. Leer una matriz 4x4 entera y determinar en qué fila y en qué columna se encuentra el número
mayor.*/
type Ejercicio1 struct{}
func (ejercicio1*Ejercicio1)rowColNumberMax(matrix [][]int)[][]int{
	numbersMax := arrayMetodos.ApplyRowCol(matrix,utilitiesEsenciaLogica.NumberMaxArray);		
	numberMax := utilitiesEsenciaLogica.NumberMaxArray(numbersMax);		
	filtro := func(num int)bool{
		return num  == numberMax;
	}
	return arrayMetodos.ApplyFindIndexRowColInIndex(matrix,filtro);
}
func(ejercicio1*Ejercicio1)answer(numbers [][][]int)string{	
	return "La fila y la columna del numero mayor es "+ utilities.MatrixToString(ejercicio1.rowColNumberMax(numbers[0][:][:]),",");
}
func ejecutarEjercicio1(){
	ejercicio1 := Ejercicio1{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"1. Leer una matriz 4x4 entera y determinar en qué fila y en qué columna se encuentra el número	mayor.",
	[2]int{4,4},
	ejercicio1.answer,
	1);
}
/*
2. Leer una matriz 4x4 entera y determinar cuántas veces se repita en ella el número mayor.*/
type Ejercicio2 struct{}
func (ejercicio2*Ejercicio2)repeatNumberMax(matrix [][]int)int{
	numberMax := utilitiesEsenciaLogica.NumberMaxMatrix(matrix);			
	filtro := func(num int)bool{
		return num  == numberMax;
	}
	return len(arrayMetodos.ApplyFindIndexRowColInIndex(matrix,filtro));
}

func(ejercicio2*Ejercicio2)answer(numbers [][][]int)string{	
	return "El numero mayor se repite :"+ strconv.Itoa(ejercicio2.repeatNumberMax(numbers[0][:][:]));
}
func ejecutarEjercicio2(){
	ejercicio2 := Ejercicio2{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"2. Leer una matriz 4x4 entera y determinar cuántas veces se repita en ella el número mayor.",
	[2]int{4,4},
	ejercicio2.answer,
	1);
}
/*
3. Leer una matriz 3x4 entera y determinar en qué posiciones exactas se encuentran los números
pares.*/
type Ejercicio3 struct{}
func (ejercicio3*Ejercicio3)positionNumberEven(matrix [][]int)[][]int{
	return arrayMetodos.ApplyFindIndexRowColInIndex(matrix,utilitiesEsenciaLogica.EvenNumber);
}
func(ejercicio3*Ejercicio3)answer(numbers [][][]int)string{	
	return "Las posisicone donde se encuentran los numeros pares son: "+ utilities.MatrixToString(ejercicio3.positionNumberEven(numbers[0][:][:]),",");
}
func ejecutarEjercicio3(){
	ejercicio3 := Ejercicio3{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"3. Leer una matriz 3x4 entera y determinar en qué posiciones exactas se encuentran los números pares.",
	[2]int{3,4},
	ejercicio3.answer,
	1);
}
/*
4. Leer una matriz 4x3 entera y determinar en qué posiciones exactas se encuentran los números
primos.*/
type Ejercicio4 struct{}
func (ejercicio4*Ejercicio4)positionNumberPrime(matrix [][]int)[][]int{		
	return arrayMetodos.ApplyFindIndexRowColInIndex(matrix,utilitiesEsenciaLogica.IfPrime);
}
func(ejercicio4*Ejercicio4)answer(numbers [][][]int)string{	
	return "Las posisicone donde se encuentran los numeros primos son: "+ utilities.MatrixToString(ejercicio4.positionNumberPrime(numbers[0][:][:]),",");
}
func ejecutarEjercicio4(){
	ejercicio4 := Ejercicio4{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"4. Leer una matriz 4x3 entera y determinar en qué posiciones exactas se encuentran los números	primos.",
	[2]int{4,3},
	ejercicio4.answer,
	1);
}
/*
5. Leer una matriz 4x3 entera, calcular la suma de los elementos de cada fila y determinar cuál es
la fila que tiene la mayor suma.*/
type Ejercicio5 struct{}
func (ejercicio5*Ejercicio5)rowColNumberMax(matrix [][]int)[]int{
	sumRow := arrayMetodos.ApplyRowCol(matrix,utilitiesEsenciaLogica.SumArray);		
	numberMax := utilitiesEsenciaLogica.NumberMaxArray(sumRow);		
	filtro := func(num int)bool{
		return num  == numberMax;
	}
	return arrayMetodos.FindArrayIndex(sumRow,filtro);
}
func(ejercicio5*Ejercicio5)answer(numbers [][][]int)string{	
	return "La  columna la mayor suma "+ utilities.ArrayToString(ejercicio5.rowColNumberMax(numbers[0][:][:]),",");
}
func ejecutarEjercicio5(){
	ejercicio5 := Ejercicio5{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"5. Leer una matriz 4x3 entera, calcular la suma de los elementos de cada fila y determinar cuál es	la fila que tiene la mayor suma.",
	[2]int{4,3},
	ejercicio5.answer,
	1);
}
/*
6. Leer una matriz 4x4 entera y calcular el promedio de los números mayores de cada fila.*/
type Ejercicio6 struct{}
func (ejercicio6*Ejercicio6)averageNumberMaxRow(matrix [][]int)int{
	numbersMax := arrayMetodos.ApplyRowCol(matrix,utilitiesEsenciaLogica.NumberMaxArray);	
	return utilitiesEsenciaLogica.AverageArray(numbersMax);		
}

func(ejercicio6*Ejercicio6)answer(numbers [][][]int)string{	
	return "El promedio enterio de los numeros mayores es :"+ strconv.Itoa(ejercicio6.averageNumberMaxRow(numbers[0][:][:]));
}
func ejecutarEjercicio6(){
	ejercicio6 := Ejercicio6{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"6. Leer una matriz 4x4 entera y calcular el promedio de los números mayores de cada fila.",
	[2]int{4,4},
	ejercicio6.answer,
	1);
}
/*
7. Leer una matriz 4x4 entera y determinar en qué posiciones están los enteros terminados en 0.*/
type Ejercicio7 struct{}
func (ejercicio7*Ejercicio7)positionNumberTerminateZero(matrix [][]int)[][]int{
	termianteInzero := func(num int)bool{
		return utilitiesEsenciaLogica.LastDigitNumberCountDigit(num,1) == 0;
	}
	return arrayMetodos.ApplyFindRowColInIndexArgument(matrix,termianteInzero);
}
func(ejercicio7*Ejercicio7)answer(numbers [][][]int)string{	
	return "Las posisicones donde se encuentran los numeros primos son: "+ utilities.MatrixToString(ejercicio7.positionNumberTerminateZero(numbers[0][:][:]),",");
}
func ejecutarEjercicio7(){
	ejercicio7 := Ejercicio7{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"7. Leer una matriz 4x4 entera y determinar en qué posiciones están los enteros terminados en 0.",
	[2]int{4,4},
	ejercicio7.answer,
	1);
}
/*
8. Leer una matriz 4x4 entera y determinar cuántos enteros terminados en 0 hay almacenados en
ella.*/
type Ejercicio8 struct{
	ejercicio7 Ejercicio7
}
func (ejercicio8*Ejercicio8)countNumberTerminateInZero(matrix [][]int)int{
	return len( ejercicio8.ejercicio7.positionNumberTerminateZero(matrix));		
}

func(ejercicio8*Ejercicio8)answer(numbers [][][]int)string{	
	return "La cantidad de numeros terminados en cero es  :"+ strconv.Itoa(ejercicio8.countNumberTerminateInZero(numbers[0][:][:]));
}
func ejecutarEjercicio8(){
	ejercicio8 := Ejercicio8{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"8. Leer una matriz 4x4 entera y determinar cuántos enteros terminados en 0 hay almacenados en ella.",
	[2]int{4,4},
	ejercicio8.answer,
	1);
}
/*
9. Leer una matriz 3x4 entera y determinar cuántos de los números almacenados son primos y
terminan en 3.*/
type Ejercicio9 struct{}
func (ejercicio9*Ejercicio9)positionNumberPrimeTerminateThree(matrix [][]int)int{
	termianteThree := func(num int)bool{
		return utilitiesEsenciaLogica.LastDigitNumberCountDigit(num,1) == 3;
	}
	return len(arrayMetodos.ApplyFindRowColInIndexArgument(matrix,utilitiesEsenciaLogica.IfPrime,termianteThree));
}
func(ejercicio9*Ejercicio9)answer(numbers [][][]int)string{	
	return "La cantidad que son numeros primos y terminan en 3 son: "+ strconv.Itoa(ejercicio9.positionNumberPrimeTerminateThree(numbers[0][:][:]));
}
func ejecutarEjercicio9(){
	ejercicio9 := Ejercicio9{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"9. Leer una matriz 3x4 entera y determinar cuántos de los números almacenados son primos y	terminan en 3.",
	[2]int{3,4},
	ejercicio9.answer,
	1);
}
/*
10. Leer una matriz 5x3 entera y determinar en qué fila está el mayor número primo.*/
type Ejercicio10 struct{}
func (ejercicio10*Ejercicio10)rowPrimeNumberMax(matrix [][]int)[]int{
	numbersPrime := arrayMetodos.ApplyFindRowColInArgument(matrix,utilitiesEsenciaLogica.IfPrime);
	primeMax := utilitiesEsenciaLogica.NumberMaxMatrix(numbersPrime); 
	return arrayMetodos.ApplyFindRowIndexArgument(matrix,func(num int)bool{
		return num == primeMax;
	});
}
//[11 1 2 3] [4 5 6 7] [8 9 90 45]
func(ejercicio10*Ejercicio10)answer(numbers [][][]int)string{	
	return "Las filas donde se encuentran el mayor numero primo son: "+ utilities.ArrayToString(ejercicio10.rowPrimeNumberMax(numbers[0][:][:]),",");
}
func ejecutarEjercicio10(){
	ejercicio10 := Ejercicio10{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"10. Leer una matriz 5x3 entera y determinar en qué fila está el mayor número primo.",
	[2]int{3,4},
	ejercicio10.answer,
	1);
}
/*
11. Leer una matriz 5x3 entera y determinar en qué columna está el menor número par.*/
type Ejercicio11 struct{}
func (ejercicio11*Ejercicio11)colNumberEvenMin(matrix [][]int)[]int{
	numbersEven := arrayMetodos.ApplyFindRowColInArgument(matrix,utilitiesEsenciaLogica.EvenNumber);
	numbersMinEven := arrayMetodos.ApplyRowCol(numbersEven,utilitiesEsenciaLogica.NumberMinArray)
	MinEven := utilitiesEsenciaLogica.NumberMinArray(numbersMinEven);
	fmt.Println("MinEven",MinEven);
	return utilitiesEsenciaLogica.UniqueNumbers(arrayMetodos.ApplyFindColInIndexArgument(matrix,func(num int)bool{
		return num == MinEven;
	}));
}
// [[[3 5 7] [2 4 6] [1 2 3] [2 3 4] [3 4 5]]]
func(ejercicio11*Ejercicio11)answer(numbers [][][]int)string{	
	return "Las columnas donde se encuentran el menor numero par son: "+ utilities.ArrayToString(ejercicio11.colNumberEvenMin(numbers[0][:][:]),",");
}
func ejecutarEjercicio11(){
	ejercicio11 := Ejercicio11{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"11. Leer una matriz 5x3 entera y determinar en qué columna está el menor número par.",
	[2]int{5,3},
	ejercicio11.answer,
	1);
}
/*
12. Leer una matriz 5x5 entera y determinar en qué fila está el mayor número terminado en 6.*/
type Ejercicio12 struct{}
func (ejercicio12*Ejercicio12)rowNumberMaxTerminateInSix(matrix [][]int)[]int{
	termianteSix := func(num int)bool{
		return utilitiesEsenciaLogica.LastDigitNumberCountDigit(num,1) == 6;
	}
	numbersTerminateSix := arrayMetodos.ApplyFindRowColInArgument(matrix,termianteSix);
	terminateSixMax :=utilitiesEsenciaLogica.NumberMaxMatrix(numbersTerminateSix);
	return arrayMetodos.ApplyFindRowIndexArgument(matrix,func(num int)bool{
		return num == terminateSixMax;
	});
}
func(ejercicio12*Ejercicio12)answer(numbers [][][]int)string{	
	return "Las filas donde se encuentran el mayor numero terminado en 6: "+ utilities.ArrayToString(ejercicio12.rowNumberMaxTerminateInSix(numbers[0][:][:]),",");
}
func ejecutarEjercicio12(){
	ejercicio12 := Ejercicio12{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"12. Leer una matriz 5x5 entera y determinar en qué fila está el mayor número terminado en 6.",
	[2]int{5,5},
	ejercicio12.answer,
	1);
}
/*
13. Leer una matriz 5x3 entera y determinar en qué columna está el mayor número que comienza
con el dígito 4.*/
type Ejercicio13 struct{}
func (ejercicio13*Ejercicio13)colNumberStarInFour(matrix [][]int)[]int{
	startThree := func(num int)bool{
		return utilitiesEsenciaLogica.StartDigitNumberCountDigit(num,1) == 4;
	}
	numbersStartFour := arrayMetodos.ApplyFindRowColInArgument(matrix,startThree);
	startFourMax := utilitiesEsenciaLogica.NumberMaxMatrix(numbersStartFour);
	return utilitiesEsenciaLogica.UniqueNumbers(arrayMetodos.ApplyFindColInIndexArgument(matrix,func(num int)bool{
		return num == startFourMax;
	}));
}

func(ejercicio13*Ejercicio13)answer(numbers [][][]int)string{	
	return "Las filas donde se encuentran el mayor numero que comienza por 4: "+ utilities.ArrayToString(ejercicio13.colNumberStarInFour(numbers[0][:][:]),",");
}
func ejecutarEjercicio13(){
	ejercicio13 := Ejercicio13{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"13. Leer una matriz 5x3 entera y determinar en qué columna está el mayor número que comienza	con el dígito 4.",
	[2]int{5,3},
	ejercicio13.answer,
	1);
}
/*
14. Leer una matriz 5x5 entera y determinar cuántos números almacenados en ella tienen mas de
3 dígitos.*/
type Ejercicio14 struct{}
func (ejercicio14*Ejercicio14)numberMaxthreeDigit(matrix [][]int)int{
	startThree := func(num int)bool{
		return utilitiesEsenciaLogica.ContarDigitosLog(num) >= 3;
	}
	return len(arrayMetodos.ApplyFindRowColInIndexArgument(matrix,startThree));
}

func(ejercicio14*Ejercicio14)answer(numbers [][][]int)string{	
	return "Los numeros que tienen mas de tres digitos son: "+ strconv.Itoa(ejercicio14.numberMaxthreeDigit(numbers[0][:][:]));
}
func ejecutarEjercicio14(){
	ejercicio14 := Ejercicio14{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"14. Leer una matriz 5x5 entera y determinar cuántos números almacenados en ella tienen mas de	3 dígitos.",
	[2]int{5,5},
	ejercicio14.answer,
	1);
}
/*
15. Leer una matriz 5x4 entera y determinar cuántos números almacenados en ella terminan en 34.*/
type Ejercicio15 struct{}
func (ejercicio15*Ejercicio15)numberTerminateInThirtyFour(matrix [][]int)int{
	startThree := func(num int)bool{
		return utilitiesEsenciaLogica.LastDigitNumberCountDigit(num,2) == 34;
	}
	fmt.Println(arrayMetodos.ApplyFindRowColInIndexArgument(matrix,startThree))
	return len(arrayMetodos.ApplyFindRowColInIndexArgument(matrix,startThree));
}

func(ejercicio15*Ejercicio15)answer(numbers [][][]int)string{	
	return "Los numeros que terminan en 34 son: "+ strconv.Itoa(ejercicio15.numberTerminateInThirtyFour(numbers[0][:][:]));
}
func ejecutarEjercicio15(){
	ejercicio15 := Ejercicio15{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"15. Leer una matriz 5x4 entera y determinar cuántos números almacenados en ella terminan en 34.",
	[2]int{5,5},
	ejercicio15.answer,
	1);
}
/*
16. Leer una matriz 5x4 entera y determinar cuántos números almacenados en ella tienen un solo
dígito.*/
type Ejercicio16 struct{}
func (ejercicio16*Ejercicio16)numberOneDigit(matrix [][]int)int{
	oneDigit := func(num int)bool{
		return utilitiesEsenciaLogica.ContarDigitosLog(num) == 1;
	}
	fmt.Println(arrayMetodos.ApplyFindRowColInIndexArgument(matrix,oneDigit))
	return len(arrayMetodos.ApplyFindRowColInIndexArgument(matrix,oneDigit));
}

func(ejercicio16*Ejercicio16)answer(numbers [][][]int)string{	
	return "Los numeros que tienen un solo digito: "+ strconv.Itoa(ejercicio16.numberOneDigit(numbers[0][:][:]));
}
func ejecutarEjercicio16(){
	ejercicio16 := Ejercicio16{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"16. Leer una matriz 5x4 entera y determinar cuántos números almacenados en ella tienen un solo	dígito.",
	[2]int{5,4},
	ejercicio16.answer,
	1);
}
/*
17. Leer una matriz 5x4 entera y determinar cuántos múltiplos de 5 hay almacenados en ella.*/
type Ejercicio17 struct{}
func (ejercicio17*Ejercicio17)countMultipleFive(matrix [][]int)int{
	multiplefive := func(num int)bool{
		return num%5 == 0;
	}
	fmt.Println(arrayMetodos.ApplyFindRowColInIndexArgument(matrix,multiplefive))
	return len(arrayMetodos.ApplyFindRowColInIndexArgument(matrix,multiplefive));
}

func(ejercicio17*Ejercicio17)answer(numbers [][][]int)string{	
	return "Los numeros multiplos de 5 son "+ strconv.Itoa(ejercicio17.countMultipleFive(numbers[0][:][:]));
}
func ejecutarEjercicio17(){
	ejercicio17 := Ejercicio17{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"17. Leer una matriz 5x4 entera y determinar cuántos múltiplos de 5 hay almacenados en ella.",
	[2]int{5,4},
	ejercicio17.answer,
	1);
}
/*
18. Leer una matriz 5x5 entera y determinar en qué posición exacta se encuentra el mayor múltiplo
de 8.*/
type Ejercicio18 struct{}
func (ejercicio18*Ejercicio18)numberMaxMultipleOfEight(matrix [][]int)[][]int{
	multipleOfEight := func(num int)bool{
		return num%8 == 0;
	}
	numbersMultipleEight := arrayMetodos.ApplyFindRowColInArgument(matrix,multipleOfEight);
	numberMaxMultipleEigth := utilitiesEsenciaLogica.NumberMaxMatrix(numbersMultipleEight);
	return arrayMetodos.ApplyFindRowColInIndexArgument(matrix,func(num int)bool{
		return num == numberMaxMultipleEigth;
	});
}

func(ejercicio18*Ejercicio18)answer(numbers [][][]int)string{	
	return "las posisicones donde se encuentra el mayor multiplo de 8 son : "+ utilities.MatrixToString(ejercicio18.numberMaxMultipleOfEight(numbers[0][:][:]),",");
}
// [[[8 80 5 80 7] [24 67 87 42 40] [72 80 12 26 16] [3 4 5 2 1] [7 6 5 4 3]]]
func ejecutarEjercicio18(){
	ejercicio18 := Ejercicio18{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"18. Leer una matriz 5x5 entera y determinar en qué posición exacta se encuentra el mayor múltiplo	de 8.",
	[2]int{5,5},
	ejercicio18.answer,
	1);
}
/*
19. Leer dos matrices 4x5 entera y determinar si sus contenidos son exactamente iguales.*/
type Ejercicio19 struct{}
func (ejercicio19*Ejercicio19)theTwoMatricesAreEqual(matrix3D [][][]int)bool{
	equal := true;
	for row, elementRow := range  matrix3D[0] {	
		for col, elementCol := range  elementRow { 				
			fmt.Println(matrix3D[1][row][col],elementCol);		
			if elementCol != matrix3D[1][row][col] {
				equal = false;
				break
			}
		}	
	}		
	return equal;
}

func(ejercicio19*Ejercicio19)answer(numbers [][][]int)string{	
	answer :="Las dos matrices no son iguales";
	if ejercicio19.theTwoMatricesAreEqual(numbers) {
		answer = "Las dos matrices si son iguales"
	}
	return "El resultado es que: "+answer;
}

func ejecutarEjercicio19(){
	ejercicio19 := Ejercicio19{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"19. Leer dos matrices 4x5 entera y determinar si sus contenidos son exactamente iguales.",
	[2]int{4,5},
	ejercicio19.answer,
	2);
}

/*
20. Leer dos matrices 4x5 entera, luego leer un entero y determinar si cada uno de los elementos
de una de las matrices es igual a cada uno de los elementos de la otra matriz multiplicado por
el entero leído.*/
type Ejercicio20 struct{
	ejercicio19 Ejercicio19
}
func (ejercicio20*Ejercicio20)theTwoMatricesAreEqualMultipleNumber(matrix3D [][][]int,numberMultiple int)bool{
	conver :=  utiliesMatrixArray.NewPowerMatrixMapArray(func(n int)int{return n*numberMultiple });
	matrixConver := conver.MapMatrix(matrix3D[0]);
	var newMatrix [][][]int;
	newMatrix = append(newMatrix,matrixConver,matrix3D[1]);
	return ejercicio20.ejercicio19.theTwoMatricesAreEqual(newMatrix);
}

func(ejercicio20*Ejercicio20)answer(numbers [][][]int)string{	
	var multipleNumber int
	fmt.Println("Escriba un numero ")
	fmt.Scanf("%d\n", &multipleNumber) 	
	answer :="Las dos matrices no son iguales";
	if ejercicio20.theTwoMatricesAreEqualMultipleNumber(numbers,multipleNumber) {
		answer = "Las dos matrices si son iguales"
	}
	return "El resultado es que: "+answer;
}

func ejecutarEjercicio20(){
	ejercicio20 := Ejercicio20{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"20. Leer dos matrices 4x5 entera, luego leer un entero y determinar si cada uno de los elementos de una de las matrices es igual a cada uno de los elementos de la otra matriz multiplicado por el entero leído.",
	[2]int{4,5},
	ejercicio20.answer,
	2);
}
/*
21. Leer dos matrices 4x5 enteras y determinar cuántos datos tienen en común.*/
type Ejercicio21 struct{}
func (ejercicio21*Ejercicio21)whatNumberHasTheSame(matrix3D [][][]int)[]int{
	var colection []int;
	arrayMetodos.ForMatrix(matrix3D[0] ,func(n0 int,row int,_ int){
		if arrayMetodos.Some(matrix3D[1][row] ,func(n1 int)bool{return n0 == n1}){
			colection = append(colection,n0);			
		}	
	})			
	return colection;
}
func(ejercicio21*Ejercicio21)answer(numbers [][][]int)string{	
	return "Los elementos parecidos son  "+utilities.ArrayToString(ejercicio21.whatNumberHasTheSame(numbers),",");
}

func ejecutarEjercicio21(){
	ejercicio21 := Ejercicio21{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"21. Leer dos matrices 4x5 enteras y determinar cuántos datos tienen en común.",
	[2]int{4,5},
	ejercicio21.answer,
	2);
}
/*
22. Leer dos matrices 4x5 enteras y determinar si el número mayor almacenado en la primera está
en la segunda.*/
type Ejercicio22 struct{}
func (ejercicio22*Ejercicio22)higherNumberInTheOtherMatrix(matrix3D [][][]int)bool{
	var numberEqual bool
	numberMax := utilitiesEsenciaLogica.NumberMaxMatrix(matrix3D[0])
	arrayMetodos.ForMatrix(matrix3D[1] ,func(n0 int,row int,_ int){
		if n0 == numberMax {
			numberEqual = true;
		}
	})
	return numberEqual
}

func(ejercicio22*Ejercicio22)answer(numbers [][][]int)string{
	answer := "no esta en la otra matrix"	
	if ejercicio22.higherNumberInTheOtherMatrix(numbers) {
		answer = "si se encuentra en la otra matrix"
	}
	return "El numero mayor de la primera matrix   "+answer;
}

func ejecutarEjercicio22(){
	ejercicio22 := Ejercicio22{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"22. Leer dos matrices 4x5 enteras y determinar si el número mayor almacenado en la primera está en la segunda.",
	[2]int{4,5},
	ejercicio22.answer,
	2);
}
/*
23. Leer dos matrices 4x5 enteras y determinar si el número mayor de una de las matrices es igual
al número mayor de la otra matriz.*/
type Ejercicio23 struct{}
func (ejercicio23*Ejercicio23)higherNumberEqualToHigherNumberinOtherMatrix(matrix3D [][][]int)bool{
	numberMax0 := utilitiesEsenciaLogica.NumberMaxMatrix(matrix3D[0])
	numberMax1 := utilitiesEsenciaLogica.NumberMaxMatrix(matrix3D[1])
	return numberMax0 == numberMax1;
}

func(ejercicio23*Ejercicio23)answer(numbers [][][]int)string{
	answer := "no es igual al numero mayor de la segunda matrix"	
	if ejercicio23.higherNumberEqualToHigherNumberinOtherMatrix(numbers) {
		answer = "es igual al numero mayor de la segunda matrix"
	}
	return "El numero mayor de la primera matrix " +answer;
}

func ejecutarEjercicio23(){
	ejercicio23 := Ejercicio23{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"23. Leer dos matrices 4x5 enteras y determinar si el número mayor de una de las matrices es igual al número mayor de la otra matriz.",
	[2]int{4,5},
	ejercicio23.answer,
	2);
}
/*
24. Leer dos matrices 4x5 enteras y determinar si el mayor número primo de una de las matrices
también se encuentra en la otra matriz.*/
type Ejercicio24 struct{}
func (ejercicio24*Ejercicio24)majorCousinNumberTheSameIsInsideTheOtherMatrix(matrix3D [][][]int)bool{
	primeNumber := utiliesMatrixArray.NewPrime{};
	numberEqual := false;
	primeNumberFilter := primeNumber.Prime();
	numberMaxPrime0 := utilitiesEsenciaLogica.NumberMaxMatrix(primeNumberFilter.FilterMatrix(matrix3D[0]))
	arrayMetodos.ForMatrix(matrix3D[1] ,func(n0 int,row int,_ int){
		if n0 == numberMaxPrime0 {
			numberEqual = true;
		}
	})
	return numberEqual ;
}

func(ejercicio24*Ejercicio24)answer(numbers [][][]int)string{
	answer := "El numero mayor primo de la primera matrix no esta en la segunda matrix"	
	if ejercicio24.majorCousinNumberTheSameIsInsideTheOtherMatrix(numbers) {
		answer = "El numero mayor primo de la primera matrix si esta en la segunda matrix"
	}
	return answer;
}
func ejecutarEjercicio24(){
	ejercicio24 := Ejercicio24{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"24. Leer dos matrices 4x5 enteras y determinar si el mayor número primo de una de las matrices	también se encuentra en la otra matriz.",
	[2]int{4,5},
	ejercicio24.answer,
	2);
}
/*
25. Leer dos matrices 4x5 enteras y determinar si el mayor número primo de una de las matrices
es también el mayor número primo de la otra matriz.*/
type Ejercicio25 struct{}
func (ejercicio25*Ejercicio25)higherCousinNumberEqualToLargerCousinNumberInTheOtherMatrix(matrix3D [][][]int)bool{
	primeNumber := utiliesMatrixArray.NewPrime{};
	primeNumberFilter := primeNumber.Prime();
	numberMaxPrime0 := utilitiesEsenciaLogica.NumberMaxMatrix(primeNumberFilter.FilterMatrix(matrix3D[0]))
	numberMaxPrime1 := utilitiesEsenciaLogica.NumberMaxMatrix(primeNumberFilter.FilterMatrix(matrix3D[1]))
	return numberMaxPrime0 == numberMaxPrime1;
}

func(ejercicio25*Ejercicio25)answer(numbers [][][]int)string{
	answer := "no es igual al numero mayor primo de la segunda matrix"	
	if ejercicio25.higherCousinNumberEqualToLargerCousinNumberInTheOtherMatrix(numbers) {
		answer = "es igual al numero mayor primo de la segunda matrix"
	}
	return answer;
}
func ejecutarEjercicio25(){
	ejercicio25 := Ejercicio25{};
	ejecutarEjercicio.NewEjecucionEjercicioMatrices(
	"25. Leer dos matrices 4x5 enteras y determinar si el mayor número primo de una de las matrices	es también el mayor número primo de la otra matriz.",
	[2]int{4,5},
	ejercicio25.answer,
	2);
}
/*
26. Leer dos matrices 4x5 enteras y determinar si la cantidad de números pares almacenados en
una matriz es igual a la cantidad de números pares almacenados en la otra matriz.
27. Leer dos matrices 4x5 enteras y determinar si la cantidad de números primos almacenados en
una matriz es igual a la cantidad de números primos almacenados en la otra matriz.
28. Leer una matriz 4x6 entera y determinar en qué posiciones se encuentran los números cuyo
penúltimo dígito sea el 5.
29. Leer una matriz 4x6 entera y determinar si alguno de sus números está repetido al menos 3
veces.
30. Leer una matriz 4x6 entera y determinar cuántas veces está en ella el número menor.
31. Leer una matriz 4x6 entera y determinar en qué posiciones están los menores por fila.
32. Leer una matriz 4x6 entera y determinar en qué posiciones están los menores primos por fila.
33. Leer una matriz 4x6 entera y determinar en qué posiciones están los menores pares por fila.
34. Leer una matriz 4x6 entera y determinar cuántos de los números almacenados en ella
pertenecen a los 100 primeros elementos de la serie de Fibonacci.
35. Leer dos matrices 4x6 enteras y determinar cuál es el mayor dato almacenado en ella que
pertenezca a la Serie de Fibonacci.
36. Leer dos matrices 4x6 enteras y determinar si el mayor número almacenado en una de ellas
que pertenezca a la Serie de Fibonacci es igual al mayor número almacenado en la otra matriz
que pertenezca a la Serie de Fibonacci.
37. Leer dos matrices 4x6 enteras y determinar si el número mayor de una matriz se encuentra en
la misma posición exacta en la otra matriz.
38. Leer dos matrices 4x6 enteras y determinar si el mayor número primo de una matriz está
repetido en la otra matriz.
39. Leer dos matrices 4x6 enteras y determinar si el promedio de las “esquinas” de una matriz es
igual al promedio de las “esquinas” de la otra matriz.
40. Leer dos matrices 5x5 enteras y determinar si el promedio entero de los elementos de la
diagonal de una matriz es igual al promedio de los elementos de la diagonal de la otra matriz.306
Capítulo 10 - Matrices
41. Leer dos matrices 5x5 enteras y determinar si el promedio entero de todos los elementos que
no están en la diagonal de una matriz es igual al promedio entero de todos los elementos que
no están en la diagonal de la otra matriz.
42. Leer dos matrices 5x5 enteras y determinar si el promedio entero de los números primos de
una matriz se encuentra almacenado en la otra matriz.
43. Leer dos matrices 5x5 enteras y determinar si el promedio entero de los números pares de una
matriz es igual al promedio de los números pares de la otra matriz.
44. Leer dos matrices 5x5 enteras y determinar si el promedio entero de los números terminados
en 4 de una matriz se encuentra al menos 3 veces en la otra matriz.
45. Leer dos matrices 5x5 enteras y determinar si el promedio entero de los números mayores de
cada fila de una matriz es igual al promedio de los números mayores de cada fila de la otra
matriz.
46. Leer dos matrices 5x5 enteras y determinar si el promedio entero de los números menores
cada fila de una matriz corresponde a alguno de los datos almacenados en las “esquinas” de la
otra matriz.
47. Leer dos matrices 5x5 enteras y determinar si el promedio de los mayores números primos por
cada fila de una matriz es igual al promedio de los mayores números primos por cada columna
de la otra matriz.
48. Leer dos matrices 5x5 entera y determinar si el promedio de los mayores elementos que
pertenecen a la serie de Fibonacci de cada fila de una matriz es igual al promedio de los
mayores elementos que pertenecen a la serie de Fibonacci de cada fila de la otra matriz.
49. Leer una matriz 3x3 entera y determinar si el promedio de todos los datos almacenados en ella
se encuentra también almacenado.
50. Leer una matriz 5x5 y determinar si el promedio de los elementos que se encuentran en su
diagonal está almacenado en ella. Mostrar en pantalla en qué posiciones exactas se encuentra
dicho dato.
*/

func EjecutarPrograma(){
	controlarFuncionamiento := controlarFuncionamientoPrograma.ConstruirTitulo("condicionales");	
	fmt.Println("Hola, bienvenido al programa de condicionales");
	controlarFuncionamiento.Administrar(ejercicios);
}