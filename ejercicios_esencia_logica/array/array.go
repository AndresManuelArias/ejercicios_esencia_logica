package array
// ejemplo :https://docslide.net/documents/ejercicios-resueltos-de-algoritmos-55b9fbab8d6b9.html
import ( "fmt"
// "math"
"sort"
"strconv"
"strings"
"../controlarFuncionamientoPrograma"
"../agregarMostrarValores"
"../arrayMetodos"
"../ejecutarEjercicio"
"../utilities"
"../utilitiesEsenciaLogica"
"../utiliesMatrixArray"

)
func ejercicios(){
	var numeroPregunta string
	posicion := [...] interface{} {	ejecutarEjercicio1,
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
		ejecutarEjercicio25,
		ejecutarEjercicio26,
		ejecutarEjercicio27,
		ejecutarEjercicio28,
		ejecutarEjercicio29,
		ejecutarEjercicio30,
		ejecutarEjercicio31,
		ejecutarEjercicio32,
		ejecutarEjercicio33,
		ejecutarEjercicio34,
		ejecutarEjercicio35,
		ejecutarEjercicio36,
		ejecutarEjercicio37,
		ejecutarEjercicio38,
		ejecutarEjercicio39,
		ejecutarEjercicio40,
		ejecutarEjercicio41,
		ejecutarEjercicio42,
		ejecutarEjercicio43,
		ejecutarEjercicio44,
		ejecutarEjercicio45,
		ejecutarEjercicio46,
		ejecutarEjercicio47,
		ejecutarEjercicio48,
		ejecutarEjercicio49,
		ejecutarEjercicio50,
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
1. Leer 10 enteros, almacenarlos en un vector y determinar en qué posición del vector está el
mayor número leído.*/
type Ejercicio1 struct{}
func(ejercicio1*Ejercicio1)indexNumberMaxArray(numbers []int)int{
	numbersMaxArray := utilitiesEsenciaLogica.NewNumberMax(numbers);
	return numbersMaxArray.IndexNumberInArray(numbers)[0];
}
func(ejercicio1*Ejercicio1)answer(numbers []int)string{
	positionNumberMax :=  strconv.Itoa(ejercicio1.indexNumberMaxArray(numbers));
	numberMax := strconv.Itoa(utilitiesEsenciaLogica.NumberMaxArray(numbers));
	return "El numero mayor es: "+numberMax+"\n Su posicion es: "+positionNumberMax;
}
func ejecutarEjercicio1(){
	ejercicio1 := new(Ejercicio1);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"1. Leer 10 enteros, almacenarlos en un vector y determinar en qué posición del vector está el	mayor número leído.",	
		agregarMostrarValores.AgregarHastaDiezNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio1.answer);
}
/*
2. Leer 10 enteros, almacenarlos en un vector y determinar en qué posición del vector está el
mayor número par leído.
*/
type Ejercicio2 struct{}
func(ejercicio2*Ejercicio2)numbersAndindexOfEvenNumber(numbers []int)([]int,int){
	evenNumber := utilitiesEsenciaLogica.NewEvenNumberMax();
	evenNumbers := evenNumber.ShearchNumbersInArray(numbers);
	positionNumberMax :=  evenNumber.IndexNumbersConditionArray(numbers);
	return evenNumbers,positionNumberMax[0];
}
func(ejercicio2*Ejercicio2)answer(numbers []int)string{
	evenNumbers,positionNumberMax := ejercicio2.numbersAndindexOfEvenNumber(numbers);
	evenNumberString := utilities.ArrayToString(evenNumbers,",");
	positionNumberMaxString :=  strconv.Itoa(positionNumberMax);
	return "Los numeros pares son : "+evenNumberString+"\n La posicion del mayor par es: "+positionNumberMaxString;
}
func ejecutarEjercicio2(){
	ejercicio2 := new(Ejercicio2);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"2. Leer 10 enteros, almacenarlos en un vector y determinar en qué posición del vector está el mayor número par leído.",
		agregarMostrarValores.AgregarHastaDiezNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio2.answer);
}
/*
3. Leer 10 enteros, almacenarlos en un vector y determinar en qué posición del vector está el
mayor número primo leído.
*/
type Ejercicio3 struct{}
func(ejercicio3*Ejercicio3)numbersAndindexOfPrimeNumber(numbers []int)([]int,int){
	objectNumberPrime := utilitiesEsenciaLogica.NewPrimeNumberMax();
	evenNumbers := objectNumberPrime.ShearchNumbersInArray(numbers);
	positionNumberMax :=  objectNumberPrime.IndexNumbersConditionArray(numbers);
	return evenNumbers,positionNumberMax[0];
}
func(ejercicio3*Ejercicio3)answer(numbers []int)string{
	primeNumbers,positionPrimeNumberMax := ejercicio3.numbersAndindexOfPrimeNumber(numbers);
	primeNumberString := utilities.ArrayToString(primeNumbers,",");
	positionPrimeNumberMaxString :=  strconv.Itoa(positionPrimeNumberMax);
	return "Los numeros primo son : "+primeNumberString+"\n La posicion del mayor par es: "+positionPrimeNumberMaxString;
}
func ejecutarEjercicio3(){
	ejercicio3 := new(Ejercicio3);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"3. Leer 10 enteros, almacenarlos en un vector y determinar en qué posición del vector está el mayor número primo leído.",
		agregarMostrarValores.AgregarHastaDiezNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio3.answer);
}
/*
4. Cargar un vector de 10 posiciones con los 10 primeros elementos de la serie de Fibonacci y
mostrarlo en pantalla.*/
type Ejercicio4 struct{
}
func(ejercicio4*Ejercicio4)answer()string{
	fibonanci := utilitiesEsenciaLogica.GenerateNumberUntil(10,utilitiesEsenciaLogica.GenerateFibonacci);
	return "Los numeros fibonancis son : "+utilities.ArrayToString(fibonanci, ",");
}
func ejecutarEjercicio4(){
	ejercicio4 := new(Ejercicio4);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("4. Cargar un vector de 10 posiciones con los 10 primeros elementos de la serie de Fibonacci y mostrarlo en pantalla.");
	ejercicicep.EjecutarEjercicioSinData(ejercicio4.answer);
}
/*
5. Almacenar en un vector de 10 posiciones los 10 números primos comprendidos entre 100 y
300. Luego mostrarlos en pantalla.*/
type Ejercicio5 struct{}
func(ejercicio5*Ejercicio5)tenNumberPrimeSinceHundredUntilThreeHundred()[]int{
	numbers := []int{100,300};
	between := utilitiesEsenciaLogica.GenerateNumbersbetweenNumbers(numbers,utilitiesEsenciaLogica.GeneratePrimes);
	hundreUntilThreeHundred := between[1:len(numbers)];
	return hundreUntilThreeHundred[0][:10];
}
func(ejercicio5*Ejercicio5)answer()string{
	primes := utilities.ArrayToString(ejercicio5.tenNumberPrimeSinceHundredUntilThreeHundred(),",");
	return "Los numeros primos son : "+primes;
}
func ejecutarEjercicio5(){
	ejercicio5 := new(Ejercicio5);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("5. Almacenar en un vector de 10 posiciones los 10 números primos comprendidos entre 100 y	300. Luego mostrarlos en pantalla.");
	ejercicicep.EjecutarEjercicioSinData(ejercicio5.answer);
}
/*
6. Leer dos números enteros y almacenar en un vector los 10 primeros números primos
comprendidos entre el menor y el mayor. Luego mostrarlos en pantalla.
*/
type Ejercicio6 struct{}
func(ejercicio6*Ejercicio6)GenerateNumbersbetweenNumbers(numbers []int)[][]int{
	between := utilitiesEsenciaLogica.GenerateNumbersbetweenNumbers(numbers,utilitiesEsenciaLogica.GeneratePrimes);
	return between[1:len(numbers)];
}
func(ejercicio6*Ejercicio6)answer(numbers []int)string{
	primes := utilities.MatrixToString(ejercicio6.GenerateNumbersbetweenNumbers(numbers),",");
	return "Los numeros primos son : "+primes;
}
func ejecutarEjercicio6(){
	ejercicio6 := new(Ejercicio6);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"6. Leer dos números enteros y almacenar en un vector los 10 primeros números primos comprendidos entre el menor y el mayor. Luego mostrarlos en pantalla.",
		agregarMostrarValores.AgregarHastaDosNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio6.answer);
}
/*
7. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones se
encuentra el número mayor.*/
type Ejercicio7 struct{}
func(ejercicio7*Ejercicio7)indexNumbersMax(numbers []int)[]int{
	numbersMaxArray := utilitiesEsenciaLogica.NewNumberMax(numbers);
	return numbersMaxArray.IndexNumberInArray(numbers);
}
func(ejercicio7*Ejercicio7)answer(numbers []int)string{
	numberMax := utilities.ArrayToString(ejercicio7.indexNumbersMax(numbers),",");
	return "Las posiciones de los numeros mayores son : "+numberMax;
}
func ejecutarEjercicio7(){
	ejercicio7 := new(Ejercicio7);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"7. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones se encuentra el número mayor.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio7.answer);
}
/*
8. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones se
encuentran los números terminados en 4.*/
type Ejercicio8 struct{} 
func(ejercicio8*Ejercicio8)numbersFinishedInFour(numbers []int)[]int{
	ShowDigit := utilitiesEsenciaLogica.NewShowDigitShowDigit(true);
	return ShowDigit.IndexNumbersDigitCousinOrLast(numbers,4);
}
func(ejercicio8*Ejercicio8)answer(numbers []int)string{
	primes := utilities.ArrayToString(ejercicio8.numbersFinishedInFour(numbers),",");
	return "Las posisiones de los numeros terminados en 4 son : "+primes;
}
func ejecutarEjercicio8(){
	ejercicio8 := new(Ejercicio8);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"8. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones se encuentran los números terminados en 4.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio8.answer);
}
/*
9. Leer 10 números enteros, almacenarlos en un vector y determinar cuántas veces está repetido
el mayor.*/
type Ejercicio9 struct{}
func(ejercicio9*Ejercicio9)countNumbersMax(numbers []int)[]int{
	numbersMaxArray := utilitiesEsenciaLogica.NewNumberMax(numbers);
	return numbersMaxArray.ShearchNumbersInArray(numbers);
}
func(ejercicio9*Ejercicio9)answer(numbers []int)string{
	numberMax := ejercicio9.countNumbersMax(numbers);
	countNumberMax := strconv.Itoa( len(numberMax));
	return "El numero mayor es "+ strconv.Itoa(numberMax[0])+" y esta repetido: "+countNumberMax;
}
func ejecutarEjercicio9(){
	ejercicio9 := new(Ejercicio9);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"9. Leer 10 números enteros, almacenarlos en un vector y determinar cuántas veces está repetido	el mayor.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio9.answer);
}
/*
10. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones se
encuentran los números con mas de 3 dígitos.*/
type Ejercicio10 struct{}
func(ejercicio10*Ejercicio10)numberMoreThreeDigit(numbers []int)[]int{
	shearchDigit :=  utilitiesEsenciaLogica.NewCountMoreDigit(3);
	return shearchDigit.ShearchNumbersInArray(numbers);
}
func(ejercicio10*Ejercicio10)answer(numbers []int)string{
	numberMaxhreeDigit := ejercicio10.numberMoreThreeDigit(numbers);
	countNumberMax := strconv.Itoa( len(numberMaxhreeDigit));
	return "El numero que tiene mas de tres digito es "+ utilities.ArrayToString(numberMaxhreeDigit,",")+" y esta repetido: "+countNumberMax;
}
func ejecutarEjercicio10(){
	ejercicio10 := new(Ejercicio10);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"10. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones se encuentran los números con mas de 3 dígitos..",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio10.answer);
}
/*
11. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números tienen, de
los almacenados allí, tienen menos de 3 dígitos.*/
type Ejercicio11 struct{}
func(ejercicio11*Ejercicio11)numberMoreThreeDigit(numbers []int)[]int{
	shearchDigit :=  utilitiesEsenciaLogica.NewCountMinDigit(3);
	return shearchDigit.ShearchNumbersInArray(numbers);
}
func(ejercicio11*Ejercicio11)answer(numbers []int)string{
	numberMaxhreeDigit := ejercicio11.numberMoreThreeDigit(numbers);
	countNumberMax := strconv.Itoa( len(numberMaxhreeDigit));
	return "El numero que tiene menos de tres digito es "+ utilities.ArrayToString(numberMaxhreeDigit,",")+" y esta repetido: "+countNumberMax;
}
func ejecutarEjercicio11(){
	ejercicio11 := new(Ejercicio11);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"11. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números tienen, de los almacenados allí, tienen menos de 3 dígitos.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio11.answer);
}
/*
12. Leer 10 números enteros, almacenarlos en un vector y determinar a cuánto es igual el
promedio entero de los datos del vector.*/
type Ejercicio12 struct{}
func(ejercicio12*Ejercicio12)averageArray(numbers []int)int{
	return utilitiesEsenciaLogica.AverageArray(numbers);
}
func(ejercicio12*Ejercicio12)answer(numbers []int)string{
	averageNumbers := strconv.Itoa(ejercicio12.averageArray(numbers));
	return "El promedio entero es :"+averageNumbers;
}
func ejecutarEjercicio12(){
	ejercicio12 := new(Ejercicio12);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"12. Leer 10 números enteros, almacenarlos en un vector y determinar a cuánto es igual el promedio entero de los datos del vector.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio12.answer);
}
/*
13. Leer 10 números enteros, almacenarlos en un vector y determinar si el promedio entero de
estos datos está almacenado en el vector.*/
type Ejercicio13 struct{}
func(ejercicio13*Ejercicio13)ifAverageInsideVector(numbers []int)[]int{
	average :=  utilitiesEsenciaLogica.NewAverageArray(numbers);
	return average.ShearchNumbersInArray(numbers);
}
func(ejercicio13*Ejercicio13)answer(numbers []int)string{
	insideVectorAverage := ejercicio13.ifAverageInsideVector(numbers);
	countNumberAverage := strconv.Itoa( len(insideVectorAverage));
	return "En el array promedio dentro del vector es:"+ utilities.ArrayToString(insideVectorAverage,",")+" y se encuentra : "+countNumberAverage+" veces";
}
func ejecutarEjercicio13(){
	ejercicio13 := new(Ejercicio13);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"13. Leer 10 números enteros, almacenarlos en un vector y determinar si el promedio entero de estos datos está almacenado en el vector.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio13.answer);
}
/*
14. Leer 10 números enteros, almacenarlos en un vector y determinar cuántas veces se repite el
promedio entero de los datos dentro del vector.*/
type Ejercicio14 struct{
	ejercicio13 Ejercicio13
}

func(ejercicio14*Ejercicio14)answer(numbers []int)string{
	insideVectorAverage := ejercicio14.ejercicio13.ifAverageInsideVector(numbers);
	countNumberAverage := strconv.Itoa( len(insideVectorAverage));
	return "En el array promedio dentro del vector es:"+ utilities.ArrayToString(insideVectorAverage,",")+" y se encuentra : "+countNumberAverage+" veces";
}
func ejecutarEjercicio14(){
	ejercicio14 := new(Ejercicio14);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"14. Leer 10 números enteros, almacenarlos en un vector y determinar cuántas veces se repite el promedio entero de los datos dentro del vector.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio14.answer);
}
/*
15. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos datos almacenados
son múltiplos de 3.*/
type Ejercicio15 struct{}
func(ejercicio15*Ejercicio15)multipleOfThree(numbers []int)[]int{
	average :=  utilitiesEsenciaLogica.NewMultiples(3);
	return average.ShearchNumbersInArray(numbers);
}
func(ejercicio15*Ejercicio15)answer(numbers []int)string{
	insideVectorMultipleThree := ejercicio15.multipleOfThree(numbers);
	countNumberMultipleThree := strconv.Itoa( len(insideVectorMultipleThree));
	return "En el array multiplo de 4 dentro del vector es:"+ utilities.ArrayToString(insideVectorMultipleThree,",")+" y se encuentra : "+countNumberMultipleThree+" veces";
}
func ejecutarEjercicio15(){
	ejercicio15 := new(Ejercicio15);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"15. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos datos almacenados son múltiplos de 3.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio15.answer);
}
/*
16. Leer 10 números enteros, almacenarlos en un vector y determinar cuáles son los datos
almacenados múltiplos de 3.*/
type Ejercicio16 struct{
	ejercicio15*Ejercicio15
}
func(ejercicio16*Ejercicio16)answer(numbers []int)string{
	insideVectorMultipleThree := ejercicio16.ejercicio15.multipleOfThree(numbers);
	countNumberMultipleThree := strconv.Itoa( len(insideVectorMultipleThree));
	return "En el array multiplo de 4 dentro del vector es:"+ utilities.ArrayToString(insideVectorMultipleThree,",")+" y se encuentra : "+countNumberMultipleThree+" veces";
}
func ejecutarEjercicio16(){
	ejercicio16 := new(Ejercicio16);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"16. Leer 10 números enteros, almacenarlos en un vector y determinar cuáles son los datos almacenados múltiplos de 3.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio16.answer);
}
/*
17. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números negativos
hay.*/
type Ejercicio17 struct{}
func(ejercicio17*Ejercicio17)negativeNumber(numbers []int)[]int{
	negativeNumber :=  utilitiesEsenciaLogica.NewNumberNegative();
	return negativeNumber.ShearchNumbersInArray(numbers);
}
func(ejercicio17*Ejercicio17)answer(numbers []int)string{
	insideVectorNegative := ejercicio17.negativeNumber(numbers);
	countNumberNegative := strconv.Itoa( len(insideVectorNegative));
	return "En el array se encuentran numeros negativos dentro del vector es:"+ utilities.ArrayToString(insideVectorNegative,",")+" y se encuentra : "+countNumberNegative+" veces";
}
func ejecutarEjercicio17(){
	ejercicio17 := new(Ejercicio17);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"17. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números negativoshay",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio17.answer);
}
/*
18. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones están los
números positivos.*/
type Ejercicio18 struct{}
func(ejercicio18*Ejercicio18)indexpositiveNumber(numbers []int)[]int{
	negativeNumber :=  utilitiesEsenciaLogica.NewNumberPositive();
	return negativeNumber.IndexNumberInArray(numbers);
}
func(ejercicio18*Ejercicio18)answer(numbers []int)string{
	insideVectorPositive := ejercicio18.indexpositiveNumber(numbers);
	countNumberPositive := strconv.Itoa( len(insideVectorPositive));
	return "En el array se encuentran en los indices numeros positivos dentro del vector es:"+ utilities.ArrayToString(insideVectorPositive,",")+" y se encuentra : "+countNumberPositive+" veces";
}
func ejecutarEjercicio18(){
	ejercicio18 := new(Ejercicio18);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"18. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones están los números positivos.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio18.answer);
}
/*
19. Leer 10 números enteros, almacenarlos en un vector y determinar cuál es el número menor.*/
type Ejercicio19 struct{}
func(ejercicio19*Ejercicio19)numberMinInArray(numbers []int)[]int{
	minNumber :=  utilitiesEsenciaLogica.NewNumberMin(numbers);
	return minNumber.ShearchNumbersInArray(numbers);
}
func(ejercicio19*Ejercicio19)answer(numbers []int)string{
	insideVectorMin := ejercicio19.numberMinInArray(numbers);
	countNumberMin := strconv.Itoa( len(insideVectorMin));
	return "El menor numero es  vector es:"+ utilities.ArrayToString(insideVectorMin,",")+" y se encuentra : "+countNumberMin+" veces";
}
func ejecutarEjercicio19(){
	ejercicio19 := new(Ejercicio19);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"19. Leer 10 números enteros, almacenarlos en un vector y determinar cuál es el número menor..",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio19.answer);
}
/*
20. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posición está el
menor número primo.*/
type Ejercicio20 struct{}
func(ejercicio20*Ejercicio20)numberIndexMinPrimeInArray(numbers []int)[]int{
	primeNumber :=  utilitiesEsenciaLogica.NewPrimeNumberMin();
	return primeNumber.IndexNumbersConditionArray(numbers);
}
func(ejercicio20*Ejercicio20)answer(numbers []int)string{
	insideMinVectorPrime := ejercicio20.numberIndexMinPrimeInArray(numbers);
	countNumberPrime := strconv.Itoa( len(insideMinVectorPrime));
	return "El menor numero  primo dentro del vector  es:"+ utilities.ArrayToString(insideMinVectorPrime,",")+" y se encuentra : "+countNumberPrime+" veces";
}
func ejecutarEjercicio20(){
	ejercicio20 := new(Ejercicio20);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"20. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posición está el menor número primo.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio20.answer);
}
/*
21. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posición está el
número cuya suma de dígitos sea la mayor.*/
type Ejercicio21 struct{}

func(ejercicio21*Ejercicio21)sumDigitMaxArrayIndex(numbers []int)[]int{
	sumdigitArray  := utilitiesEsenciaLogica.NewSumDigitArray();
	return sumdigitArray.IndexMaxArray(numbers);
}
func(ejercicio21*Ejercicio21)sumDigitArray(numbers []int)[]int{
	sumdigitArray  := utilitiesEsenciaLogica.NewSumDigitArray();
	return sumdigitArray.MaxArray(numbers);
}
func(ejercicio21*Ejercicio21)answer(numbers []int)string{
	sumDigitMaxArrayIndex := utilities.ArrayToString( ejercicio21.sumDigitMaxArrayIndex(numbers),",");
	sumDigitArray := utilities.ArrayToString( ejercicio21.sumDigitArray(numbers),",");	
	return "El indice con la mayor suma de digitos es  "+sumDigitMaxArrayIndex+"\n la suma de los digitos da los siguiente: "+sumDigitArray;
}
func ejecutarEjercicio21(){
	ejercicio21 := new(Ejercicio21);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"21. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posición está el número cuya suma de dígitos sea la mayor.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio21.answer);
}
/*
22. Leer 10 números enteros, almacenarlos en un vector y determinar cuáles son los números
múltiplos de 5 y en qué posiciones están.*/
type Ejercicio22 struct{}
func(ejercicio22*Ejercicio22)multipleOfFive(numbers []int)[]int{
	average :=  utilitiesEsenciaLogica.NewMultiples(5);
	return average.ShearchNumbersInArray(numbers);
}
func(ejercicio22*Ejercicio22)answer(numbers []int)string{
	insideVectorMultipleFive := ejercicio22.multipleOfFive(numbers);
	countNumberMultipleFive := strconv.Itoa( len(insideVectorMultipleFive));
	return "En el array los multiplo de 5 son:"+ utilities.ArrayToString(insideVectorMultipleFive,",")+" y se encuentra : "+countNumberMultipleFive+" veces";
}
func ejecutarEjercicio22(){
	ejercicio22 := new(Ejercicio22);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"22. Leer 10 números enteros, almacenarlos en un vector y determinar cuáles son los números	múltiplos de 5 y en qué posiciones están.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio22.answer);
}
/*
23. Leer 10 números enteros, almacenarlos en un vector y determinar si existe al menos un
número repetido.*/

type Ejercicio23 struct{}
func(ejercicio23*Ejercicio23)NumbersSameArray(numbers []int)[]int{
	return  utilitiesEsenciaLogica.SameNumbers(numbers);
}
func(ejercicio23*Ejercicio23)answer(numbers []int)string{
	sameNumbers := ejercicio23.NumbersSameArray(numbers);
	countSameNumber := strconv.Itoa( len(sameNumbers));
	return "En el array sus numeros repetidos son:"+ utilities.ArrayToString(sameNumbers,",")+" y se encuentra : "+countSameNumber+" veces";
}
func ejecutarEjercicio23(){
	ejercicio23 := new(Ejercicio23);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"23. Leer 10 números enteros, almacenarlos en un vector y determinar si existe al menos un número repetido.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio23.answer);
}
/*
24. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posición está el
número con mas dígitos.*/
type Ejercicio24 struct{}


func(ejercicio24*Ejercicio24)NumberMaxDigitIndexArray(numbers []int)[]int{
	countDigit := utilitiesEsenciaLogica.NewCounDigitArray();
	return countDigit.IndexMaxArray(numbers);
}
func(ejercicio24*Ejercicio24)counDigitArray(numbers []int)[]int{
	countDigit := utilitiesEsenciaLogica.NewCounDigitArray();
	return countDigit.MaxArray(numbers)
}
func(ejercicio24*Ejercicio24)answer(numbers []int)string{
	sumDigitMaxArrayIndex := utilities.ArrayToString( ejercicio24.NumberMaxDigitIndexArray(numbers),",");
	sumDigitArray := utilities.ArrayToString( ejercicio24.counDigitArray(numbers),",");	
	return "El indice con la mayor cantidad de digitos es  "+sumDigitMaxArrayIndex+"\n la cantidad de digitos que tiene es la siguiente: "+sumDigitArray;
}
func ejecutarEjercicio24(){
	ejercicio24 := new(Ejercicio24);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"24. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posición está el	número con mas dígitos.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicioF(ejercicio24.answer);
}
/*
25. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos de los números
leídos son números primos terminados en 3.*/

type Ejercicio25 struct{} 
func(ejercicio25*Ejercicio25)numbersPrimeFinishedInThree(numbers []int)[]int{
	ShowDigit := utilitiesEsenciaLogica.NewShowDigitShowDigit(true);
	primeNumber := utilitiesEsenciaLogica.NewNumberPrime();
	return ShowDigit.NumbersDigitCousinOrLast(primeNumber.ShearchNumbersInArray(numbers),3);
}
func(ejercicio25*Ejercicio25)answer(numbers []int)string{
	primes := utilities.ArrayToString(ejercicio25.numbersPrimeFinishedInThree(numbers),",");
	return "Los  numeros primos terminados en 3 son : "+primes;
}
func ejecutarEjercicio25(){
	ejercicio25 := new(Ejercicio25);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"25. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos de los números leídos son números primos terminados en 3.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio25.answer);
}
/*
26. Leer 10 números enteros, almacenarlos en un vector y calcularle el factorial a cada uno de los
números leídos almacenándolos en otro vector.*/
type Ejercicio26 struct{} 
func(ejercicio26*Ejercicio26)factorial(numbers []int)int{
	return utilitiesEsenciaLogica.MultipleNumberArray(numbers);
}
func(ejercicio26*Ejercicio26)answer(numbers []int)string{
	factorialNumber := strconv.Itoa(ejercicio26.factorial(numbers))
	return "El factorial es : "+factorialNumber;
}
func ejecutarEjercicio26(){
	ejercicio26 := new(Ejercicio26);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"26. Leer 10 números enteros, almacenarlos en un vector y calcularle el factorial a cada uno de los	números leídos almacenándolos en otro vector.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio26.answer);
}
/*
27. Leer 10 números enteros, almacenarlos en un vector y determinar a cuánto es igual el
promedio entero de los factoriales de cada uno de los números leídos.
*/
type Ejercicio27 struct{} 
func(ejercicio27*Ejercicio27)factorial(number int)int{
	numberGenerate := func()func()int{
		generate := utilitiesEsenciaLogica.GenerateNumber(1)
		return func()int{
			return generate();
		}
	}
	array := utilitiesEsenciaLogica.GenerateNumberUntil(number,numberGenerate)
	return utilitiesEsenciaLogica.MultipleNumberArray(array);
}
func(ejercicio27*Ejercicio27)factorialAverage(numbers []int)int{
	factorial := utiliesMatrixArray.NewPowerMatrixMapArray(ejercicio27.factorial);
	return utilitiesEsenciaLogica.AverageArray(factorial.MapArray(numbers));
}
func(ejercicio27*Ejercicio27)answer(numbers []int)string{
	factorialAverageNumber := strconv.Itoa(ejercicio27.factorialAverage(numbers))
	return "El factorial es : "+factorialAverageNumber;
}
func ejecutarEjercicio27(){
	ejercicio27 := new(Ejercicio27);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"27. Leer 10 números enteros, almacenarlos en un vector y determinar a cuánto es igual el promedio entero de los factoriales de cada uno de los números leídos.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio27.answer);
}
/*
28. Leer 10 números enteros, almacenarlos en un vector y mostrar en pantalla todos los enteros
comprendidos entre 1 y cada uno de los números almacenados en el vector.*/
type Ejercicio28 struct{} 
func(ejercicio28*Ejercicio28)GenerateNumberOneToOneUntil(number int)[]int{
	numberGenerate := func()func()int{
		generate := utilitiesEsenciaLogica.GenerateNumber(1)
		return func()int{
			return generate();
		}
	}
	return utilitiesEsenciaLogica.GenerateNumberUntil(number,numberGenerate)
}
func(ejercicio28*Ejercicio28)generateNumber(numbers []int)[][]int{
	return arrayMetodos.MapTransformDataConvertToMatrix(numbers,ejercicio28.GenerateNumberOneToOneUntil)
}
func(ejercicio28*Ejercicio28)answer(numbers []int)string{
	generateNumberOfOneToOne := utilities.MatrixToString(ejercicio28.generateNumber(numbers),",");	
	return "El factorial es : "+ generateNumberOfOneToOne;
}
func ejecutarEjercicio28(){
	ejercicio28 := new(Ejercicio28);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"28. Leer 10 números enteros, almacenarlos en un vector y mostrar en pantalla todos los enteros comprendidos entre 1 y cada uno de los números almacenados en el vector.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio28.answer);
}
/*
29. Leer 10 números enteros, almacenarlos en un vector y mostrar en pantalla todos los enteros
comprendidos entre 1 y cada uno de los dígitos de cada uno de los números almacenados en
el vector.*/
type Ejercicio29 struct{
	ejercicio28 Ejercicio28
} 

func(ejercicio29*Ejercicio29)generateNumberDigitUntil(number int)string{
	arrayDigit := utilitiesEsenciaLogica.ArrayDigitosNumero(number);
	return utilities.MatrixToString(ejercicio29.ejercicio28.generateNumber(arrayDigit),",");
}
func(ejercicio29*Ejercicio29)generateNumbersDigitUntil(numbers []int)[]string{
	return arrayMetodos.MapTransformDataString(numbers,ejercicio29.generateNumberDigitUntil);
}
func(ejercicio29*Ejercicio29)answer(numbers []int)string{	
	return "El factorial es : "+ strings.Join(ejercicio29.generateNumbersDigitUntil(numbers), " | ") ;
}
func ejecutarEjercicio29(){
	ejercicio29 := new(Ejercicio29);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"29. Leer 10 números enteros, almacenarlos en un vector y mostrar en pantalla todos los enteros	comprendidos entre 1 y cada uno de los dígitos de cada uno de los números almacenados en el vector.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio29.answer);
}
/*
30. Leer 10 números enteros, almacenarlos en un vector. Luego leer un entero y determinar si este
último entero se encuentra entre los 10 valores almacenados en el vector.*/
type Ejercicio30 struct{} 

func(ejercicio30*Ejercicio30)theLastNumberIsFound(numbers []int)[]int{
	lastNumber := utiliesMatrixArray.NewPowerMatrixArray(func(n int)bool{return n == numbers[len(numbers)-1]})
	return lastNumber.Filter(numbers)
}
func(ejercicio30*Ejercicio30)answer(numbers []int)string{
	foundNumber := utilities.ArrayToString(ejercicio30.theLastNumberIsFound(numbers),",");	
	return "El numero se encuentra  "+ foundNumber;
}
func ejecutarEjercicio30(){
	ejercicio30 := new(Ejercicio30);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"30. Leer 10 números enteros, almacenarlos en un vector. Luego leer un entero y determinar si este último entero se encuentra entre los 10 valores almacenados en el vector.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio30.answer);
}
// http://www.goconejemplos.com/recursion
/*
31. Leer 10 números enteros, almacenarlos en un vector. Luego leer un entero y determinar
cuantos divisores exactos tiene este último número entre los valores almacenados en el vector.*/
type Ejercicio31 struct{} 

func(ejercicio31*Ejercicio31)theLastNumberIsDivisible(numbers []int)[]int{
	lastNumber := utiliesMatrixArray.NewPowerMatrixArray(func(n int)bool{return 0 == numbers[len(numbers)-1]%n})
	return lastNumber.Filter(numbers)
}
func(ejercicio31*Ejercicio31)answer(numbers []int)string{
	foundNumber := utilities.ArrayToString(ejercicio31.theLastNumberIsDivisible(numbers),",");	
	return "Los numeros que dividen al ultimo numero son   "+ foundNumber;
}
func ejecutarEjercicio31(){
	ejercicio31 := new(Ejercicio31);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"31. Leer 10 números enteros, almacenarlos en un vector. Luego leer un entero y determinar cuantos divisores exactos tiene este último número entre los valores almacenados en el vector.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio31.answer);
}
/*
32. Leer 10 números enteros, almacenarlos en un vector. Luego leer un entero y determinar
cuántos números de los almacenados en el vector terminan en el mismo dígito que el último
valor leído.*/
type Ejercicio32 struct{} 

func(ejercicio32*Ejercicio32)theLastNumberIsDivisible(numbers []int)[]int{
	lastNumberDigit := utiliesMatrixArray.NewPowerMatrixArray(func(n int)bool{
		return utilitiesEsenciaLogica.LastDigitNumber(n) == utilitiesEsenciaLogica.LastDigitNumber(numbers[len(numbers)-1])});
	return lastNumberDigit.Filter(numbers);
}
func(ejercicio32*Ejercicio32)answer(numbers []int)string{
	foundNumber := utilities.ArrayToString(ejercicio32.theLastNumberIsDivisible(numbers),",");	
	return "Los numeros que dividen al ultimo numero son   "+ foundNumber;
}
func ejecutarEjercicio32(){
	ejercicio32 := new(Ejercicio32);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"32. Leer 10 números enteros, almacenarlos en un vector. Luego leer un entero y determinar	cuántos números de los almacenados en el vector terminan en el mismo dígito que el último	valor leído.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio32.answer);
}
/*
33. Leer 10 números enteros, almacenarlos en un vector y determinar a cuánto es igual la suma de
los dígitos pares de cada uno de los números leídos.*/
type Ejercicio33 struct{} 

func(ejercicio33*Ejercicio33)sumDigitEven(numbers []int)[]int{
	sumDigitEven  :=utiliesMatrixArray.NewPowerMatrixMapArray( utilitiesEsenciaLogica.SumDigitEven);
	return sumDigitEven.MapArray(numbers);
}

func(ejercicio33*Ejercicio33)answer(numbers []int)string{
	foundNumber := utilities.ArrayToString(ejercicio33.sumDigitEven(numbers),",");	
	return "Los numeros que dividen al ultimo numero son   "+ foundNumber;
}
func ejecutarEjercicio33(){
	ejercicio33 := new(Ejercicio33);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"33. Leer 10 números enteros, almacenarlos en un vector y determinar a cuánto es igual la suma de los dígitos pares de cada uno de los números leídos.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio33.answer);
}
/*
34. Leer 10 números enteros, almacenarlos en un vector y determinar cuántas veces en el vector
se encuentra el dígito 2. No se olvide que el dígito 2 puede estar varias veces en un mismo
número.*/
type Ejercicio34 struct{} 

func(ejercicio34*Ejercicio34)countDigitTwo(numbers []int)int{
	twoNumber := utiliesMatrixArray.NewPowerMatrixMapArray(func(nu int)int{
		arrayDigit := utilitiesEsenciaLogica.ArrayDigitosNumero(nu);
		numbersTwo := arrayMetodos.Filter(arrayDigit,func(n int )bool{return n%10 == 2});  
		return len(numbersTwo);
	});
	return utilitiesEsenciaLogica.SumArray(twoNumber.MapArray(numbers));
}

func(ejercicio34*Ejercicio34)answer(numbers []int)string{
	countNumberEven := strconv.Itoa(ejercicio34.countDigitTwo(numbers));	
	return "La cuenta de los digitos 2 es    "+ countNumberEven;
}
func ejecutarEjercicio34(){
	ejercicio34 := new(Ejercicio34);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"34. Leer 10 números enteros, almacenarlos en un vector y determinar cuántas veces en el vector	se encuentra el dígito 2. No se olvide que el dígito 2 puede estar varias veces en un mismo	número.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio34.answer);
}

/*
35. Leer 10 números enteros, almacenarlos en un vector y determinar si el promedio entero de
dichos números es un número primo.*/
type Ejercicio35 struct{} 

func(ejercicio35*Ejercicio35)averageIntIsPrime(numbers []int)(bool,int){
	average := utilitiesEsenciaLogica.AverageArray(numbers)
	return utilitiesEsenciaLogica.IfPrimeInt(average),average;
}

func(ejercicio35*Ejercicio35)answer(numbers []int)string{
	answer := "El Promedio entero no es primo";
	ifPrime,average := ejercicio35.averageIntIsPrime(numbers)
	if ifPrime{
		answer = "El Promedio entero si es primo"
	}	
	return "El resultado es "+strconv.Itoa(average)+". "+answer;
}
func ejecutarEjercicio35(){
	ejercicio35 := new(Ejercicio35);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"35. Leer 10 números enteros, almacenarlos en un vector y determinar si el promedio entero de dichos números es un número primo.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio35.answer);
}
/*
36. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos dígitos primos hay
en los números leídos.*/
type Ejercicio36 struct{} 

func(ejercicio36*Ejercicio36)numbersIsPrime(numbers []int)[]int{
	prime := utiliesMatrixArray.NewPowerMatrixArray(utilitiesEsenciaLogica.IfPrimeInt);
	return prime.Filter(numbers);
}

func(ejercicio36*Ejercicio36)answer(numbers []int)string{
	return "Los numeros primos son "+utilities.ArrayToString(ejercicio36.numbersIsPrime(numbers),",");
}
func ejecutarEjercicio36(){
	ejercicio36 := new(Ejercicio36);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"36. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos dígitos primos hay	en los números leídos.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio36.answer);
}
/*
37. Leer 10 números enteros, almacenarlos en un vector y determinar a cuántos es igual el
cuadrado de cada uno de los números leídos.*/
type Ejercicio37 struct{} 
func(ejercicio37*Ejercicio37)numbersIsPrime(numbers []int)[]int{
	pow := utiliesMatrixArray.NewPowerMatrixMapArray(func(number int)int{return number * number});
	return pow.MapArray(numbers);
}
func(ejercicio37*Ejercicio37)answer(numbers []int)string{
	return "El cuadrado de los numeros es "+utilities.ArrayToString(ejercicio37.numbersIsPrime(numbers),",");
}
func ejecutarEjercicio37(){
	ejercicio37 := new(Ejercicio37);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"37. Leer 10 números enteros, almacenarlos en un vector y determinar a cuántos es igual el cuadrado de cada uno de los números leídos.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio37.answer);
}
/*
38. Leer 10 números enteros, almacenarlos en un vector y determinar si la semisuma entre el valor
mayor y el valor menor es un número primo.*/
type Ejercicio38 struct{} 

func(ejercicio38*Ejercicio38)averageIntIsPrime(numbers []int)(bool,int){
	sort.Ints(numbers);
	sum := numbers[0]+numbers[len(numbers)-1];
	return utilitiesEsenciaLogica.IfPrimeInt(sum),sum;
}
func(ejercicio38*Ejercicio38)answer(numbers []int)string{
	answer := "La suma del mayor y meno no es primo";
	ifPrime,average := ejercicio38.averageIntIsPrime(numbers)
	if ifPrime{
		answer = "La suma del mayor y meno si es primo"
	}	
	return "El resultado es "+strconv.Itoa(average)+". "+answer;
}
func ejecutarEjercicio38(){
	ejercicio38 := new(Ejercicio38);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"38. Leer 10 números enteros, almacenarlos en un vector y determinar si la semisuma entre el valor mayor y el valor menor es un número primo.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio38.answer);
}
/*
39. Leer 10 números enteros, almacenarlos en un vector y determinar si la semisuma entre el valor
mayor y el valor menor es un número par.*/
type Ejercicio39 struct{} 
func(ejercicio39*Ejercicio39)eventLastAndFrist(numbers []int)(bool,int){
	sort.Ints(numbers);
	sum := numbers[0]+numbers[len(numbers)-1];
	return utilitiesEsenciaLogica.EvenNumber(sum),sum;
}
func(ejercicio39*Ejercicio39)answer(numbers []int)string{
	answer := "La suma del mayor y meno no es par";
	ifPrime,average := ejercicio39.eventLastAndFrist(numbers)
	if ifPrime{
		answer = "La suma del mayor y meno si es par"
	}	
	return "El resultado es "+strconv.Itoa(average)+". "+answer;
}
func ejecutarEjercicio39(){
	ejercicio39 := new(Ejercicio39);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"39. Leer 10 números enteros, almacenarlos en un vector y determinar si la semisuma entre el valor mayor y el valor menor es un número par.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio39.answer);
}
/*
40. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números de los
almacenados en dicho vector terminan en 15.*/
type Ejercicio40 struct{} 
func(ejercicio40*Ejercicio40)lastNumberInFifteen(numbers []int)[]int{
	lastNumber := utiliesMatrixArray.NewPowerMatrixArray(func(n int)bool{return 15 == utilitiesEsenciaLogica.LastDigitNumberCountDigit(n,2)});
	return lastNumber.Filter(numbers);
}
func(ejercicio40*Ejercicio40)answer(numbers []int)string{	
	return "los numeros terminados en 15 son "+utilities.ArrayToString(ejercicio40.lastNumberInFifteen(numbers),",");
}
func ejecutarEjercicio40(){
	ejercicio40 := new(Ejercicio40);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"40. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números de los	almacenados en dicho vector terminan en 15.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio40.answer);
}
/*
41. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números de los
almacenados en dicho vector comienzan con 3.*/
type Ejercicio41 struct{} 
func(ejercicio41*Ejercicio41)countStartWithThree(numbers []int)int{
	showDigit := utilitiesEsenciaLogica.NewShowDigitShowDigit(false);
	return len(showDigit.NumbersDigitCousinOrLast(numbers,3));
}
func(ejercicio41*Ejercicio41)answer(numbers []int)string{	
	return "los numeros que comienzan con 3 son "+strconv.Itoa(ejercicio41.countStartWithThree(numbers));
}
func ejecutarEjercicio41(){
	ejercicio41 := new(Ejercicio41);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"41. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números de los	almacenados en dicho vector comienzan con 3.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio41.answer);
}
/*
42. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números con
cantidad par de dígitos pares hay almacenados en dicho vector.*/
type Ejercicio42 struct{} 
func(ejercicio42*Ejercicio42)countEvenDigit(numbers []int)int{
	lastNumber := utiliesMatrixArray.NewPowerMatrixArray(func(n int)bool{
		return 0 == len(utilitiesEsenciaLogica.ArrayDigitosNumero(n))%2});
	return len(lastNumber.Filter(numbers));
}
func(ejercicio42*Ejercicio42)answer(numbers []int)string{	
	return "los digitos con numeros pares son "+strconv.Itoa(ejercicio42.countEvenDigit(numbers));
}
func ejecutarEjercicio42(){
	ejercicio42 := new(Ejercicio42);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"42. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números con cantidad par de dígitos pares hay almacenados en dicho vector.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio42.answer);
}
/*
43. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones se
encuentra el número con mayor cantidad de dígitos primos.*/
type Ejercicio43 struct{} 
func(ejercicio43*Ejercicio43)countPrimeDigit(numbers []int)[]int{
	mapNumberLenDigit := utiliesMatrixArray.NewPowerMatrixMapArray(func(n int)int{
		return len(utilitiesEsenciaLogica.ArrayDigitosNumero(n));
	});
	numberslen := mapNumberLenDigit.MapArray(numbers);
	maxNumber := utilitiesEsenciaLogica.NewPrimeNumberMax();
	return maxNumber.IndexNumbersConditionArray( numberslen);
}
func(ejercicio43*Ejercicio43)answer(numbers []int)string{	
	return "las posiciones en las que se encuentra los numeros con mayor cantidad de digitos son "+ utilities.ArrayToString(ejercicio43.countPrimeDigit(numbers),",");
}
func ejecutarEjercicio43(){
	ejercicio43 := new(Ejercicio43);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"43. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones se encuentra el número con mayor cantidad de dígitos primos.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio43.answer);
}
/*
44. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos de los números
almacenados en dicho vector pertenecen a los 100 primeros elementos de la serie de
Fibonacci.*/
type Ejercicio44 struct{} 
func(ejercicio44*Ejercicio44)filterNumberFibonancci(numbers []int)[]int{
	fibonanci := utilitiesEsenciaLogica.GenerateNumberUntil(89,utilitiesEsenciaLogica.GenerateFibonacci);	
	fiterFibonanci := utiliesMatrixArray.NewPowerMatrixArray(func(nu int)bool{
		return arrayMetodos.Some(fibonanci,func(n int)bool{return  n == nu});
	});
	return fiterFibonanci.Filter(numbers);
}
func(ejercicio44*Ejercicio44)answer(numbers []int)string{	
	return "Estos numeros pertenecen a la serie fibonanci"+ utilities.ArrayToString(ejercicio44.filterNumberFibonancci(numbers),",");
}
func ejecutarEjercicio44(){
	ejercicio44 := new(Ejercicio44);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"44. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos de los números	almacenados en dicho vector pertenecen a los 100 primeros elementos de la serie de	Fibonacci.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio44.answer);
}
/*
45. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números de los
almacenados en dicho vector comienzan por 34.*/
type Ejercicio45 struct{} 
func(ejercicio45*Ejercicio45)countStartWithThirtyFour(numbers []int)int{
	showDigit := utilitiesEsenciaLogica.NewShowDigitShowDigit(false);
	return len(showDigit.NumbersDigitCousinOrLastCount(numbers,34,2));
}
func(ejercicio45*Ejercicio45)answer(numbers []int)string{	
	return "los numeros que comienzan con 34 son "+strconv.Itoa(ejercicio45.countStartWithThirtyFour(numbers));
}
func ejecutarEjercicio45(){
	ejercicio45 := new(Ejercicio45);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"45. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números de los	almacenados en dicho vector comienzan por 34.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio45.answer);
}
/*
46. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números de los
almacenados en dicho vector son primos y comienzan por 5.*/
type Ejercicio46 struct{} 
func(ejercicio46*Ejercicio46)countStartWithFiveAndEvent(numbers []int)int{
	showDigit := utilitiesEsenciaLogica.NewShowDigitShowDigit(false);
	primeNumber := utilitiesEsenciaLogica.NewNumberPrime();
	return len(showDigit.NumbersDigitCousinOrLastCount(primeNumber.ShearchNumbersInArray(numbers),5,1));
}
func(ejercicio46*Ejercicio46)answer(numbers []int)string{	
	return "los numeros que comienzan con 5 y son primos"+strconv.Itoa(ejercicio46.countStartWithFiveAndEvent(numbers));
}
func ejecutarEjercicio46(){
	ejercicio46 := new(Ejercicio46);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"46. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números de los	almacenados en dicho vector son primos y comienzan por 5..",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio46.answer);
}
/*
47. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones se
encuentran los números múltiplos de 10. No utilizar el número 10 en ninguna operación.*/
type Ejercicio47 struct{} 
func(ejercicio47*Ejercicio47)multipleOfTeen(numbers []int)[]int{
	multipleOFTeen := utiliesMatrixArray.NewPowerMatrixArray(func(number int)bool{
		return number % (5*2) == 0;
	});
	return multipleOFTeen.FindArrayIndex(numbers);
}
func(ejercicio47*Ejercicio47)answer(numbers []int)string{	
	return "los primos de 10 son "+ utilities.ArrayToString(ejercicio47.multipleOfTeen(numbers),",");
}
func ejecutarEjercicio47(){
	ejercicio47 := new(Ejercicio47);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"47. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posiciones se encuentran los números múltiplos de 10. No utilizar el número 10 en ninguna operación.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio47.answer);
}
/*
48. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posición se
encuentra el número primo con mayor cantidad de dígitos pares.*/
type Ejercicio48 struct{} 
func(ejercicio48*Ejercicio48)countEvenDigit(numbers []int)[]int{
	numberslen := arrayMetodos.MapTransformData(numbers,func(n int)int{
		var transfor int;
		if len(utilitiesEsenciaLogica.ArrayDigitosNumero(n)) % 2 == 0 && utilitiesEsenciaLogica.IfPrime(n) {
			transfor = len(utilitiesEsenciaLogica.ArrayDigitosNumero(n));
		}
		return transfor;
	});
	lenMax := utilitiesEsenciaLogica.NumberMaxArray(numberslen);
	primeNumberDigitEvenMax := utiliesMatrixArray.NewPowerMatrixArray(func(n int)bool{
		return  utilitiesEsenciaLogica.IfPrime(n) && len(utilitiesEsenciaLogica.ArrayDigitosNumero(n)) == lenMax;
	})
	return primeNumberDigitEvenMax.FindArrayIndex(numbers);
}
func(ejercicio48*Ejercicio48)answer(numbers []int)string{	
	return "El numero primo con mayor numero de digitos par son "+ utilities.ArrayToString(ejercicio48.countEvenDigit(numbers),",");
}
func ejecutarEjercicio48(){
	ejercicio48 := new(Ejercicio48);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"48. Leer 10 números enteros, almacenarlos en un vector y determinar en qué posición se	encuentra el número primo con mayor cantidad de dígitos pares.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio48.answer);
}
/*
49. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números terminar
en dígito primo.*/
type Ejercicio49 struct{} 
func(ejercicio49*Ejercicio49)terminateDigitPrime(numbers []int)[]int{
	return arrayMetodos.FindArrayIndex(numbers,func(n int)bool{
		digitsPrime := utilitiesEsenciaLogica.ArrayDigitosNumero(n);
		return  utilitiesEsenciaLogica.IfPrimeInt(digitsPrime[len(digitsPrime)-1]);
	});
}
func(ejercicio49*Ejercicio49)answer(numbers []int)string{	
	return "Los numeros que terminan con digitos primos son "+ utilities.ArrayToString(ejercicio49.terminateDigitPrime(numbers),",");
}
func ejecutarEjercicio49(){
	ejercicio49 := new(Ejercicio49);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"49. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números terminar en dígito primo.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio49.answer);
}
/*
50. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números de los
almacenados en dicho vector comienzan en dígito primo.
*/
type Ejercicio50 struct{} 
func(ejercicio50*Ejercicio50)terminateDigitPrime(numbers []int)[]int{
	return arrayMetodos.FindArrayIndex(numbers,func(n int)bool{
		digitsPrime := utilitiesEsenciaLogica.ArrayDigitosNumero(n);
		return  utilitiesEsenciaLogica.IfPrimeInt(digitsPrime[0]);
	});
}
func(ejercicio50*Ejercicio50)answer(numbers []int)string{	
	return "Los numeros que comienzan con digitos primos son "+ utilities.ArrayToString(ejercicio50.terminateDigitPrime(numbers),",");
}
func ejecutarEjercicio50(){
	ejercicio50 := new(Ejercicio50);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"50. Leer 10 números enteros, almacenarlos en un vector y determinar cuántos números de los	almacenados en dicho vector comienzan en dígito primo.",
		agregarMostrarValores.AgregarHastaDiezNumeros,	
		agregarMostrarValores.AgregarValoreVariableEntera); 
	ejercicicep.EjecutarEjercicioF(ejercicio50.answer);
}

func EjecutarPrograma(){
	controlarFuncionamiento := controlarFuncionamientoPrograma.ConstruirTitulo("array");	
	fmt.Println("Hola, bienvenido al programa de array");
	controlarFuncionamiento.Administrar(ejercicios);
}