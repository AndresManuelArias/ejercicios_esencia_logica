package ciclos
// ejemplo :https://docslide.net/documents/ejercicios-resueltos-de-algoritmos-55b9fbab8d6b9.html
import ( "fmt"
//"math"
"sort"
"strconv"
"strings"
"../controlarFuncionamientoPrograma"
"../agregarMostrarValores"
"../arrayMetodos"
"../ejecutarEjercicio"
"../condicionales"
"../utiliesMatrixArray"
"../utilitiesEsenciaLogica"
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

func arrayToInt(numbers []int) int {
	var unionNumber int;
	for number := 0 ;number < len(numbers);number++{
		unionNumber = unionNumber *10 + numbers[number];
	}	
	return unionNumber;
}

func  contarDigitos(numbers int) int{
	condition := condicionales.ExportarObjetosCondicionales();
	return condition.Ejercicio2.ContarDigitos(numbers);
}
func arrayDigitosNumero(numbers int) []int{
	condition := condicionales.ExportarObjetosCondicionales();
	return condition.Ejercicio4.ArrayDigitosNumero(numbers)
}
func transformNumber(numbers []int, function func( []int)[][]int )[]int{
	arrayNumber := function(numbers);
	return arrayMetodos.MapTransformMatrixDataInIntArray(arrayNumber,arrayToInt);
}	
func arrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
func arrayToStringFloat32(a []float32, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
func matrixToString( matrix [][]int,delimColumn string )string{
	var colectionNumber [] string;	
	for _,element := range matrix{		
		colectionNumber = append(colectionNumber,arrayToString(element,delimColumn));
	}
	return strings.Join(colectionNumber, " | ");
}
func numerIfPrime(number int)bool{
	condition := condicionales.ExportarObjetosCondicionales();
	return condition.Ejercicio6.DeterminarSiPrimoInt(number);
}

/*
1. Leer un número entero y mostrar todos los enteros comprendidos entre 1 y el número leído.
*/
type Ejercicio1 struct{

}
func(ejercicio1*Ejercicio1) generateGrowthOfNumbers(quantity int, differenceBetweenNumbers int)[]int{
	numberGenerate := func()func()int{
		generate := utilitiesEsenciaLogica.GenerateNumber(differenceBetweenNumbers)
		return func()int{
			return generate();
		}
	 }
	return utilitiesEsenciaLogica.GenerateNumberUntil(quantity,numberGenerate)
}
func(ejercicio1*Ejercicio1)generateNumbersUp(number int)[]int{
	return ejercicio1.generateGrowthOfNumbers(number,1)
}
func(ejercicio1*Ejercicio1) generateNnumbersAccordingToCondition(numbers []int,functionFilterNumbers func ([]int) []int)[][]int{
	colectionNumber := arrayMetodos.MapTransformDataConvertToMatrix(numbers,ejercicio1.generateNumbersUp);		
	filterColectionNumber :=  arrayMetodos.MapTransformMatrixDataInMatrix(colectionNumber,functionFilterNumbers )
	return filterColectionNumber;		
}
func(ejercicio1*Ejercicio1)theNumberIncludedUpToThis(number int)string{
	return arrayToString(ejercicio1.generateNumbersUp(number),",");
}
func(ejercicio1*Ejercicio1)generateTheNumberIncludedUpToThis(numbers []int)string{
	coleccionArray := arrayMetodos.MapTransformDataString(numbers,ejercicio1.theNumberIncludedUpToThis);
	return strings.Join(coleccionArray, "| ");
}
	
func(ejercicio1*Ejercicio1)answer(numbers []int)string{
	return ejercicio1.generateTheNumberIncludedUpToThis(numbers);
}	
func ejecutarEjercicio1(){
	ejercicio1 := Ejercicio1{};
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"1. Leer un número entero y mostrar todos los enteros comprendidos entre 1 y el número leído.",
		2,
	agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicio(ejercicio1.answer);
}
/*
2. Leer un número entero y mostrar todos los pares comprendidos entre 1 y el número leído.*/
type Ejercicio2 struct{
	ejercicio1 Ejercicio1
}
func(ejercicio2*Ejercicio2) filterEvenNumbers(numbers []int) []int{
	evenNumber := utiliesMatrixArray.NewEvenNumber();
	return evenNumber.Filter(numbers);
}
func(ejercicio2*Ejercicio2) generateEvenNumbers(numbers []int)[][]int{
	return ejercicio2.ejercicio1.generateNnumbersAccordingToCondition(numbers,ejercicio2.filterEvenNumbers);
}
func(ejercicio2*Ejercicio2)answer(numbers []int)string{
	return matrixToString(ejercicio2.generateEvenNumbers(numbers),",")	
}
func ejecutarEjercicio2(){
	ejercicio2 := Ejercicio2{};
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
	"2. Leer un número entero y mostrar todos los pares comprendidos entre 1 y el número leído.",
	2,
	agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicio(ejercicio2.answer);
}
/*
3. Leer un número entero y mostrar todos los divisores exactos del número comprendidos entre 1
y el número leído.
*/
type Ejercicio3 struct{
	ejercicio1 Ejercicio1
	numberDivider int
}
func(ejercicio3*Ejercicio3)numbersDivisibleByThis(number int)bool{
	return  ejercicio3.numberDivider %number == 0;
}
func(ejercicio3*Ejercicio3)filterNnumbersDivisibleByThis(numbers []int )[]int{
	ejercicio3.numberDivider = numbers[len(numbers)-1];
	return arrayMetodos.Filter(numbers,ejercicio3.numbersDivisibleByThis);	
}	
func(ejercicio3*Ejercicio3) generateDivisorNumbers(numbers []int)[][]int{	
	return ejercicio3.ejercicio1.generateNnumbersAccordingToCondition(numbers,ejercicio3.filterNnumbersDivisibleByThis);	
}
func(ejercicio3*Ejercicio3)answer(numbers []int)string{
	return matrixToString(ejercicio3.generateDivisorNumbers(numbers),",")
}
func ejecutarEjercicio3(){
	ejercicio3 := new(Ejercicio3);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
	"3. Leer un número entero y mostrar todos los divisores exactos del número comprendidos entre 1	y el número leído.",
	2,
	agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicio(ejercicio3.answer);
}
/*
4. Leer dos números y mostrar todos los enteros comprendidos entre ellos.*/
type Ejercicio4 struct{
	ejercicio1 Ejercicio1
	firstNumber int
}
func(ejercicio4*Ejercicio4)filterNumbers(number int)bool{
	return number >ejercicio4.firstNumber;
}	
func(ejercicio4*Ejercicio4)betweenNumbers(numbers []int)[]int{
	result := arrayMetodos.Filter(numbers,ejercicio4.filterNumbers);
	var firstNumber int;
	if len(numbers) != 0{
		firstNumber = numbers[len(numbers)-1] ;
	}
	ejercicio4.firstNumber = firstNumber;
	return 		result;
}
func(ejercicio4*Ejercicio4) generateBetweenNumbers(numbers []int)[][]int{
	numberGenerate := func()func()int{
		generate := utilitiesEsenciaLogica.GenerateNumber(1)
		return func()int{
			return generate()
		}
	 }
	return utilitiesEsenciaLogica.GenerateNumbersbetweenNumbers(numbers,numberGenerate)
}
func(ejercicio4*Ejercicio4)answer(numbers []int)string{
	return matrixToString(ejercicio4.generateBetweenNumbers(numbers),",")
}
func ejecutarEjercicio4(){
	ejercicio4 := new(Ejercicio4);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
	"4. Leer dos números y mostrar todos los enteros comprendidos entre ellos.",
	3,
	agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicio(ejercicio4.answer);
}
/*
5. Leer dos números y mostrar todos los números terminados en 4 comprendidos entre ellos.*/
type Ejercicio5 struct{
	ejercicio4 Ejercicio4
	terminate int
}

func(ejercicio5*Ejercicio5)numbersTerminate(number int)bool{
	return  number % 10 == ejercicio5.terminate ;
}
	
func(ejercicio5*Ejercicio5)betweenNumbersTerminate(numbers []int)[]int{
	utilies := utiliesMatrixArray.NewPowerMatrixArray(ejercicio5.numbersTerminate)	
	return utilies.Filter(numbers)
}
func(ejercicio5*Ejercicio5) generateBetweenNumbersAndTerminate(numbers []int)[][]int{
	utilies := utiliesMatrixArray.NewPowerMatrixArray(ejercicio5.numbersTerminate);	
	generateNUmber := ejercicio5.ejercicio4.generateBetweenNumbers(numbers);
	return utilies.FilterRowsMatrix(generateNUmber);
}
func(ejercicio5*Ejercicio5)answer(numbers []int)string{
	return matrixToString(ejercicio5.generateBetweenNumbersAndTerminate(numbers),",")
}
func ejecutarEjercicio5(){
	ejercicio5 := new(Ejercicio5);
	ejercicio5.terminate = 4;
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
	"5. Leer dos números y mostrar todos los números terminados en 4 comprendidos entre ellos.",
	3,
	agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicio(ejercicio5.answer);
}
/*
6. Leer un número entero de tres dígitos y mostrar todos los enteros comprendidos entre 1 y cada
uno de los dígitos.
*/
type Ejercicio6 struct{
	ejercicio1 Ejercicio1
}
func(ejercicio6*Ejercicio6) arrayGeneratNumbersUpToArray(numbers []int)[]int{
	matrixNumbers := arrayMetodos.MapTransformDataConvertToMatrix(numbers, ejercicio6.ejercicio1.generateNumbersUp);
	return arrayMetodos.MapTransformMatrixDataInIntArray(matrixNumbers,arrayToInt);
}	
func(ejercicio6*Ejercicio6) matrixGeneratenumbersUpInArrayToMatrix(numbers [][]int)[][]int{	
	return arrayMetodos.MapTransformMatrixDataInMatrix(numbers,ejercicio6.arrayGeneratNumbersUpToArray);
}
func(ejercicio6*Ejercicio6) matrixDigitColection(numbers []int)[][]int{
	return arrayMetodos.MapTransformDataConvertToMatrix(numbers,arrayDigitosNumero);	
}
func(ejercicio6*Ejercicio6) generateNumberUpToTheDigitsOfTheNumber(numbers []int)[][]int{		
	arrayDigitColection := arrayMetodos.MapTransformDataConvertToMatrix(numbers,arrayDigitosNumero);	
	return ejercicio6.matrixGeneratenumbersUpInArrayToMatrix(arrayDigitColection)
}
func(ejercicio6*Ejercicio6)answer(numbers []int)string{
	return matrixToString(ejercicio6.generateNumberUpToTheDigitsOfTheNumber(numbers),",")
}
func ejecutarEjercicio6(){
	ejercicio6 := new(Ejercicio6);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
	"6. Leer un número entero de tres dígitos y mostrar todos los enteros comprendidos entre 1 y cada 	uno de los dígitos.",
	3,
	agregarMostrarValores.IngresarTresDigitos);
	ejercicicep.EjecutarEjercicio(ejercicio6.answer);
}
/*
7. Mostrar en pantalla todos los enteros comprendidos entre 1 y 100.*/
type Ejercicio7 struct{
	ejercicio1 Ejercicio1
}
func(ejercicio7*Ejercicio7) generateNumbersUp100()[]int{
	return 	ejercicio7.ejercicio1.generateNumbersUp(100);
}
func(ejercicio7*Ejercicio7)answer()string{
	return arrayToString(ejercicio7.generateNumbersUp100(),",");	
}
func ejecutarEjercicio7(){
	ejercicio7 := new(Ejercicio7);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("7. Mostrar en pantalla todos los enteros comprendidos entre 1 y 100.");
	ejercicicep.EjecutarEjercicioSinData(ejercicio7.answer);
}
/*
8. Mostrar en pantalla todos los pares comprendidos entre 20 y 200.*/
type Ejercicio8 struct{
	ejercicio4 Ejercicio4
}
func(ejercicio8*Ejercicio8) generateNumbersOf20to200()[]int{
	numbers :=[]int{20,200};
	numberBetwen := ejercicio8.ejercicio4.generateBetweenNumbers(numbers)
	return 	numberBetwen[1];
}
func(ejercicio8*Ejercicio8)answer()string{
	return arrayToString(ejercicio8.generateNumbersOf20to200(),",");	
}
func ejecutarEjercicio8(){
	ejercicio8 := new(Ejercicio8);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("8. Mostrar en pantalla todos los pares comprendidos entre 20 y 200.");
	ejercicicep.EjecutarEjercicioSinData(ejercicio8.answer);
}
/*
9. Mostrar en pantalla todos los números terminados en 6 comprendidos entre 25 y 205.*/
type Ejercicio9 struct{
	ejercicio5 Ejercicio5
}
func(ejercicio9*Ejercicio9) generateNumberTerminateSix()[]int{
	numbers :=[]int{25,205};
	ejercicio9.ejercicio5.terminate = 6;	
	numberTerminateInSix := ejercicio9.ejercicio5.generateBetweenNumbersAndTerminate(numbers)
	return 	numberTerminateInSix[1];
}
func(ejercicio9*Ejercicio9)answer()string{
	return arrayToString(ejercicio9.generateNumberTerminateSix(),",");	
}
func ejecutarEjercicio9(){
	ejercicio9 := new(Ejercicio9);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("9. Mostrar en pantalla todos los números terminados en 6 comprendidos entre 25 y 205.");
	ejercicicep.EjecutarEjercicioSinData(ejercicio9.answer);
}
/*
10. Leer un número entero y determinar a cuánto es igual la suma de todos los enteros
comprendidos entre 1 y el número leído.*/

type Ejercicio10 struct{
	ejercicio1 Ejercicio1
}
func(ejercicio10*Ejercicio10) unionNumber(before int,after int)int{
	return before + after;
}
func (ejercicio10*Ejercicio10)sumArray(number[]int)int{
	return arrayMetodos.ReduceInt(number,ejercicio10.unionNumber);
}
func(ejercicio10*Ejercicio10) sumArrayNumberUp(numbers []int)[]int{
	matrixNumber :=  arrayMetodos.MapTransformDataConvertToMatrix(numbers, ejercicio10.ejercicio1.generateNumbersUp);
	return arrayMetodos.MapTransformMatrixDataInIntArray(matrixNumber,ejercicio10.sumArray);
}	
func(ejercicio10*Ejercicio10) sumMatrixNumberUP(numbers [][]int)[][]int{		
	return arrayMetodos.MapTransformMatrixDataInMatrix(numbers,ejercicio10.sumArrayNumberUp);
}
func(ejercicio10*Ejercicio10) fibonanciOfArrayShowInMatrix(numbers []int)[][]int{		
	matrixDigitColection := arrayMetodos.MapTransformDataConvertToMatrix(numbers,arrayDigitosNumero);	
	return ejercicio10.sumMatrixNumberUP(matrixDigitColection);
}
func(ejercicio10*Ejercicio10)answer(numbers []int)string{
	return matrixToString(ejercicio10.fibonanciOfArrayShowInMatrix(numbers),",");
}
func ejecutarEjercicio10(){
	ejercicio10 := new(Ejercicio10);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
	"10. Leer un número entero y determinar a cuánto es igual la suma de todos los enteros 	comprendidos entre 1 y el número leído.",
	3,
	agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicio(ejercicio10.answer);
}
/*
11. Leer un número entero de dos dígitos y mostrar en pantalla todos los enteros comprendidos
entre un dígito y otro.
*/
type Ejercicio11 struct{
	ejercicio1 Ejercicio1
	ejercicio4 Ejercicio4
}
func(ejercicio11*Ejercicio11) transformNumber(numbers []int)[]int{
	matrixNumber := ejercicio11.ejercicio4.generateBetweenNumbers(numbers);
	return arrayMetodos.MapTransformMatrixDataInIntArray(matrixNumber,arrayToInt);
}	
func(ejercicio11*Ejercicio11) generateBetweenNumbersInMatrix(numbers [][]int)[][]int{		
	return arrayMetodos.MapTransformMatrixDataInMatrix(numbers,ejercicio11.transformNumber);
}
func(ejercicio11*Ejercicio11) generateNumberUpToTheDigitsOfTheNumber(numbers []int)[][]int{		
	arrayDigitColection := arrayMetodos.MapTransformDataConvertToMatrix(numbers,arrayDigitosNumero);	
	return ejercicio11.generateBetweenNumbersInMatrix(arrayDigitColection)
}
func(ejercicio11*Ejercicio11)answer(numbers []int)string{
	return matrixToString(ejercicio11.generateNumberUpToTheDigitsOfTheNumber(numbers),",")
}
func ejecutarEjercicio11(){
	ejercicio11 := new(Ejercicio11);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
	"11. Leer un número entero de dos dígitos y mostrar en pantalla todos los enteros comprendidos	entre un dígito y otro.",
	2,
	agregarMostrarValores.IngresarDosDigitos);
	ejercicicep.EjecutarEjercicio(ejercicio11.answer);
}
/*
12. Leer un número entero de 3 dígitos y determinar si tiene el dígito 1.
*/
type Ejercicio12 struct{
}
func(ejercicio12*Ejercicio12) searchNumberOne(number int)bool{
	return 	number == 1;
}	
func(ejercicio12*Ejercicio12) shearchNumberOneInArrayShowInMatrix(numbers []int)[][]int{		
	matrixDigitColection := arrayMetodos.MapTransformDataConvertToMatrix(numbers,arrayDigitosNumero);
	utilies := utiliesMatrixArray.NewPowerMatrixArray(ejercicio12.searchNumberOne)	
	return utilies.FilterMatrix(matrixDigitColection);
}
func(ejercicio12*Ejercicio12)answer(numbers []int)string{
	 result := "no hay numeros con digitos iguales a uno";
	numbersEgualOne := ejercicio12.shearchNumberOneInArrayShowInMatrix(numbers)	
	if len(numbersEgualOne) > 0 {
		result = "tiene numero con  digito iguales a uno " + matrixToString(numbersEgualOne,"")
	}
	return result
}
func ejecutarEjercicio12(){
	ejercicio12 := new(Ejercicio12);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
	"12. Leer un número entero de 3 dígitos y determinar si tiene el dígito 1.",
	3,
	agregarMostrarValores.IngresarTresDigitos);
	ejercicicep.EjecutarEjercicio(ejercicio12.answer);
}
/*
13. Leer un entero y mostrar todos los múltiplos de 5 comprendidos entre 1 y el número leído.*/
type Ejercicio13 struct{
}
func(ejercicio13*Ejercicio13)itIsMultipleOfFive(number int)bool{
	return 	number%5 == 0;
}	
func(ejercicio13*Ejercicio13) shearchNumberOneInArrayShowInMatrix(numbers []int)[][]int{		
	matrixDigitColection := arrayMetodos.MapTransformDataConvertToMatrix(numbers,arrayDigitosNumero);
	utilies := utiliesMatrixArray.NewPowerMatrixArray(ejercicio13.itIsMultipleOfFive)	
	return utilies.FilterMatrix(matrixDigitColection);
}
func(ejercicio13*Ejercicio13)answer(numbers []int)string{
	 result := "no hay digitos con multiplos de cinco";
	numbersEgualOne := ejercicio13.shearchNumberOneInArrayShowInMatrix(numbers)	
	if len(numbersEgualOne) > 0 {
		result = "tiene numeros con  digito multiplos de cinco " + matrixToString(numbersEgualOne,"")
	}
	return result
}
func ejecutarEjercicio13(){
	ejercicio13 := new(Ejercicio13);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
	"13. Leer un entero y mostrar todos los múltiplos de 5 comprendidos entre 1 y el número leído.",
	3,
	agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicio(ejercicio13.answer);
}
/*
14. Mostrar en pantalla los primeros 20 múltiplos de 3.
*/
type Ejercicio14 struct{
	ejercicio1 Ejercicio1
}

func(ejercicio14*Ejercicio14) generateTwentyFirstNumbersMultipleThree()[]int{
	return ejercicio14.ejercicio1.generateGrowthOfNumbers(20,3)
}
func(ejercicio14*Ejercicio14)answer()string{
	return arrayToString(ejercicio14.generateTwentyFirstNumbersMultipleThree(),",");	
}
func ejecutarEjercicio14(){
	ejercicio14 := new(Ejercicio14);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("14. Mostrar en pantalla los primeros 20 múltiplos de 3.");
	ejercicicep.EjecutarEjercicioSinData(ejercicio14.answer);
}
/*
15. Escribir en pantalla el resultado de sumar los primeros 20 múltiplos de 3.*/
type Ejercicio15 struct{
	ejercicio1 Ejercicio1
	ejercicio10 Ejercicio10
}
func(ejercicio15*Ejercicio15) generateNumbersAndSum(quantity int, differenceBetweenNumbers int)int{
	return ejercicio15.ejercicio10.sumArray(ejercicio15.ejercicio1.generateGrowthOfNumbers(quantity , differenceBetweenNumbers ));
}
func(ejercicio15*Ejercicio15) generateTwentyFirstSumNumbersMultipleThree()int{
	return ejercicio15.generateNumbersAndSum(20,3);
}	
func(ejercicio15*Ejercicio15)answer()string{
	return "la sumatoria de los multiplos de 3 es " + strconv.Itoa(ejercicio15.generateTwentyFirstSumNumbersMultipleThree());	
}
func ejecutarEjercicio15(){
	ejercicio15 := new(Ejercicio15);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("15. Escribir en pantalla el resultado de sumar los primeros 20 múltiplos de 3.");
	ejercicicep.EjecutarEjercicioSinData(ejercicio15.answer);
}
/*
16. Mostrar en pantalla el promedio entero de los n primeros múltiplos de 3 para un número n
leído.*/
type Ejercicio16 struct{
	ejercicio1 Ejercicio1
	ejercicio10 Ejercicio10
}
func(ejercicio16*Ejercicio16) averageInt(numbers []int)int{
	return  ejercicio16.ejercicio10.sumArray(numbers)/len(numbers); 
}
func(ejercicio16*Ejercicio16) average(numbers []int)float32{
	return float32( ejercicio16.ejercicio10.sumArray(numbers))/float32(len(numbers)); 
}
	
func(ejercicio16*Ejercicio16) generateNumbersAndAverage(quantity int, differenceBetweenNumbers int)float32{
	return ejercicio16.average(ejercicio16.ejercicio1.generateGrowthOfNumbers(quantity , differenceBetweenNumbers ))
}
func(ejercicio16*Ejercicio16) generateAverageMultipleOfThree(quantity int)float32{
	return ejercicio16.generateNumbersAndAverage(quantity , 3 );
}
func(ejercicio16*Ejercicio16) transfortArrayAverageMultipleOfThree(numbers []int)[]float32{
	return arrayMetodos.MapTransformDataFloat32(numbers, ejercicio16.generateAverageMultipleOfThree);
}
func(ejercicio16*Ejercicio16)answer(numbers []int)string{
	return "El promedio de los multiplos de 3 es :" + arrayToStringFloat32(ejercicio16.transfortArrayAverageMultipleOfThree(numbers)," ");	
}
func ejecutarEjercicio16(){
	ejercicio16 := new(Ejercicio16);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"16. Mostrar en pantalla el promedio entero de los n primeros múltiplos de 3 para un número n		leído",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicio(ejercicio16.answer);
}
/*
17. Promediar los x primeros múltiplos de 2 y determinar si ese promedio es mayor que los y
primeros múltiplos de 5 para valores de x y y leídos.
*/
type Ejercicio17 struct{
	ejercicio16 Ejercicio16
	trackerPair int
}


func(ejercicio17*Ejercicio17) generateAverageMultiple(quantity int)float32{
	multiple := 5;
	if ejercicio17.trackerPair%2 == 0 {
		multiple = 2;
	}
	ejercicio17.trackerPair = ejercicio17.trackerPair +1;
	return ejercicio17.ejercicio16.generateNumbersAndAverage(quantity , multiple );
}
func(ejercicio17*Ejercicio17) transfortArrayAverageMultipleOfTwoAndFive(numbers []int)[]float32{
	return arrayMetodos.MapTransformDataFloat32(numbers, ejercicio17.generateAverageMultiple);
}
func(ejercicio17*Ejercicio17) firstNumberGreaterThanSecond(numbers []float32) bool{
	var fristNumber float32;
	compareFristSecond := func(number float32) bool {
		result := number > fristNumber;
		fristNumber = number;
		return result;
	}
	return arrayMetodos.SomeFloat32(numbers,compareFristSecond);
}
func(ejercicio17*Ejercicio17)answer(numbers []int)string{
	averageFiveAndtwo := ejercicio17.transfortArrayAverageMultipleOfTwoAndFive(numbers);
	firstNumberGraterSecond := "Ningun promedio de los multiplos de 2 son mayores que los multiplos de 5";
	if ejercicio17.firstNumberGreaterThanSecond(averageFiveAndtwo) {
		firstNumberGraterSecond = " Un promedio de los multiplos de 2 es mayor que los multiplos de 5";
	}
	return "El promedio de los multiplos de 5 y 2 son :" + arrayToStringFloat32(averageFiveAndtwo,"  ")+firstNumberGraterSecond;	
}
func ejecutarEjercicio17(){
	ejercicio17 := new(Ejercicio17);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"17. Promediar los x primeros múltiplos de 2 y determinar si ese promedio es mayor que los y  primeros múltiplos de 5 para valores de x y y leídos.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicio(ejercicio17.answer);
}

/*
18. Leer dos números entero y mostrar todos los múltiplos de 5 comprendidos entre el menor y el
mayor.*/
type Ejercicio18 struct{
	ejercicio16 Ejercicio16
	ejercicio4 Ejercicio4
	ejercicio13 Ejercicio13
}

func(ejercicio18*Ejercicio18) betweenMultipleFive(numbers []int)[][]int{
	betweenNumbers := ejercicio18.ejercicio4.generateBetweenNumbers(numbers);
	return arrayMetodos.FilterRowsMatrix(betweenNumbers,ejercicio18.ejercicio13.itIsMultipleOfFive);
}
func(ejercicio18*Ejercicio18)answer(numbers []int)string{
	return " los multiplos de 5 :" + matrixToString(ejercicio18.betweenMultipleFive(numbers)," - ");	
}
func ejecutarEjercicio18(){
	ejercicio18 := new(Ejercicio18);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"18. Leer dos números entero y mostrar todos los múltiplos de 5 comprendidos entre el menor y el mayor",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);
	ejercicicep.EjecutarEjercicio(ejercicio18.answer);
}
/*
19. Leer un número entero y determinar si es primo.

*/

type Ejercicio19 struct{
	prime utiliesMatrixArray.NewPrime
}

func(ejercicio19*Ejercicio19) numbersPrime(numbers []int)[]int{
	prime :=  ejercicio19.prime.Prime();
	return prime.Filter(numbers);
}
func(ejercicio19*Ejercicio19)answer(numbers []int)string{
	return " Estos son los numeros primos :"+ arrayToString(ejercicio19.numbersPrime(numbers)," ");	
}
func ejecutarEjercicio19(){
	ejercicio19 := new(Ejercicio19);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"19. Leer un número entero y determinar si es primo.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio19.answer);
}
/*
20. Leer un número entero y determinar cuántos dígitos tiene.*/
type Ejercicio20 struct{
}

func(ejercicio20*Ejercicio20) countDigitOfNumber(numbers []int)[]int{	
	return arrayMetodos.MapTransformData(numbers ,contarDigitos);
}
func(ejercicio20*Ejercicio20)answer(numbers []int)string{
	return " Esta es la cantidad de digitos que tienen los numeros ingresados:"+ arrayToString(ejercicio20.countDigitOfNumber(numbers)," ");	
}
func ejecutarEjercicio20(){
	ejercicio20 := new(Ejercicio20);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"20. Leer un número entero y determinar cuántos dígitos tiene.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio20.answer);
}

/*
21. Leer un número entero y determinar a cuánto es igual al suma de sus dígitos.
*/
type Ejercicio21 struct{
	ejercicio10 Ejercicio10
}
func(ejercicio21*Ejercicio21) sumDigit(number int)int{	
	return ejercicio21.ejercicio10.sumArray(arrayDigitosNumero(number));
}
func(ejercicio21*Ejercicio21) sumArrayDigits(numbers []int)[]int{	
	return arrayMetodos.MapTransformData(numbers ,ejercicio21.sumDigit);
}
func(ejercicio21*Ejercicio21)answer(numbers []int)string{
	return " Esta es la suma de los digitos de los numeros ingresados :"+ arrayToString(ejercicio21.sumArrayDigits(numbers)," ");	
}
func ejecutarEjercicio21(){
	ejercicio21 := new(Ejercicio21);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"21. Leer un número entero y determinar a cuánto es igual al suma de sus dígitos.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio21.answer);
}
/*
22. Leer un número entero y determinar cuántas veces tiene el dígito 1.
*/
type Ejercicio22 struct{
	numberSearch int
}
func(ejercicio22*Ejercicio22) checkIfItTheNumber(number int)bool{	
	return number == ejercicio22.numberSearch;
}
func(ejercicio22*Ejercicio22) searchNumbersInArray(numbers []int)[]int{
	return arrayMetodos.Filter( numbers,ejercicio22.checkIfItTheNumber);
}
func(ejercicio22*Ejercicio22) countDigitInNumbers(number int)int{
	return len( ejercicio22.searchNumbersInArray( arrayDigitosNumero(number)));
}
func(ejercicio22*Ejercicio22) countDigitInNumbersInArray(numbers []int)[]int{	
	return arrayMetodos.MapTransformData(numbers ,ejercicio22.countDigitInNumbers);
}
func(ejercicio22*Ejercicio22)answer(numbers []int)string{
	return " Los numeros ingresados tienen el digito uno estas veces :"+ arrayToString(ejercicio22.countDigitInNumbersInArray(numbers)," ");	
}
func ejecutarEjercicio22(){
	ejercicio22 := Ejercicio22{numberSearch:1};
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"22. Leer un número entero y determinar cuántas veces tiene el dígito 1.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio22.answer);
}
/*
23. Leer un número entero y determinar si la suma de sus dígitos es también un número primo.
*/
type Ejercicio23 struct{
	ejercicio21 Ejercicio21
	ejercicio19 Ejercicio19
}

func(ejercicio23*Ejercicio23) addDigitsIfPrimeNumbersArray(numbers []int)[]int{
	return ejercicio23.ejercicio19.numbersPrime(ejercicio23.ejercicio21.sumArrayDigits(numbers));
}
func(ejercicio23*Ejercicio23)answer(numbers []int)string{
	return "La suma de los numberos que primos son  :"+ arrayToString(ejercicio23.addDigitsIfPrimeNumbersArray(numbers)," ");	
}
func ejecutarEjercicio23(){
	ejercicio23 := new(Ejercicio23);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"23. Leer un número entero y determinar si la suma de sus dígitos es también un número primo.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio23.answer);
}
/*
24. Leer un número entero y determinar a cuánto es igual al suma de sus dígitos pares.
*/
type Ejercicio24 struct{
	ejercicio2 Ejercicio2
	ejercicio10 Ejercicio10
}

func(ejercicio24*Ejercicio24) addDigitsIfEvenNumber(number int)int{
	arrayDigits := arrayDigitosNumero(number);
	evenNumber := utiliesMatrixArray.NewEvenNumber();
	return evenNumber.SumNumbersCondition(arrayDigits);
}
func(ejercicio24*Ejercicio24) addDigitsIfEvenNumbersArray(numbers []int)[]int{
	return arrayMetodos.MapTransformData(numbers ,ejercicio24.addDigitsIfEvenNumber);
}	
func(ejercicio24*Ejercicio24)answer(numbers []int)string{
	return "La suma de los digitos pares son  :"+ arrayToString(ejercicio24.addDigitsIfEvenNumbersArray(numbers)," ");	
}
func ejecutarEjercicio24(){
	ejercicio24 := new(Ejercicio24);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"24. Leer un número entero y determinar a cuánto es igual al suma de sus dígitos pares.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio24.answer);
}
/*
25. Leer un número entero y determinar a cuánto es igual el promedio entero de sus dígitos.
*/
type Ejercicio25 struct{
	ejercicio16 Ejercicio16
}
func(ejercicio25*Ejercicio25) averageDigitNumber(number int) int{
	return ejercicio25.ejercicio16.averageInt( arrayDigitosNumero(number));
}	
func(ejercicio25*Ejercicio25) averageDigitNumbersArray(numbers []int) []int{	
	return arrayMetodos.MapTransformData(numbers ,ejercicio25.averageDigitNumber);
}
func(ejercicio25*Ejercicio25)answer(numbers []int)string{
	return "El promedio entero de los digitos pares son  :"+ arrayToString(ejercicio25.averageDigitNumbersArray(numbers)," ");	
}
func ejecutarEjercicio25(){
	ejercicio25 := new(Ejercicio25);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"25. Leer un número entero y determinar a cuánto es igual el promedio entero de sus dígitos.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio25.answer);
}
/*
26. Leer un número entero y determinar cuál es el mayor de sus dígitos.*/
type Ejercicio26 struct{
}
func(ejercicio26*Ejercicio26) NumberMaxArray(numbers []int) int{
	sort.Ints(numbers);
	return 	numbers[len(numbers)-1];
}
func(ejercicio26*Ejercicio26)digitNumberMax(number int) int{
	return ejercicio26.NumberMaxArray( arrayDigitosNumero(number));
}	
func(ejercicio26*Ejercicio26) digitNumberMaxArray(numbers []int) []int{	
	return arrayMetodos.MapTransformData(numbers ,ejercicio26.digitNumberMax);
}
func(ejercicio26*Ejercicio26)answer(numbers []int)string{
return "Sus digitos mayores son :"+ arrayToString(ejercicio26.digitNumberMaxArray(numbers)," ");	
}
func ejecutarEjercicio26(){
	ejercicio26 := new(Ejercicio26);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"26. Leer un número entero y determinar cuál es el mayor de sus dígitos.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio26.answer);
}
/*
27. Leer 2 números enteros y determinar cuál de los dos tiene mayor cantidad de dígitos.
*/
type Ejercicio27 struct{
	ejercicio6 Ejercicio6
	ejercicio26 Ejercicio26
}
func(ejercicio27*Ejercicio27)countNumbers(numbers []int) int{
	return len(numbers);
}	
func(ejercicio27*Ejercicio27) countDigitNumbers(numbers []int) []int{	
	digitsMatrix := ejercicio27.ejercicio6.matrixDigitColection(numbers);
	return arrayMetodos.MapTransformMatrixDataInIntArray(digitsMatrix,ejercicio27.countNumbers);
}
func(ejercicio27*Ejercicio27) countDigitNumbersMax(numbers []int) ([]int, int){
	countDigitNumber := ejercicio27.countDigitNumbers(numbers);
	return countDigitNumber,ejercicio27.ejercicio26.NumberMaxArray(countDigitNumber);
}
func(ejercicio27*Ejercicio27)answer(numbers []int)string{
	countDigitNumber,maxDititNumber :=  ejercicio27.countDigitNumbersMax(numbers);
	return "Su cantidad de  digitos son :"+ arrayToString(countDigitNumber," ")+" y el digito mayor es "+strconv.Itoa(maxDititNumber) ;	
}
func ejecutarEjercicio27(){
	ejercicio27 := new(Ejercicio27);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"27. Leer 2 números enteros y determinar cuál de los dos tiene mayor cantidad de dígitos.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio27.answer);
}
/*
28. Leer 2 números enteros y determinar cual de los dos tiene mayor cantidad de dígitos primos.
*/
type Ejercicio28 struct{
	ejercicio27 Ejercicio27	
}
	
func(ejercicio28*Ejercicio28) digitPrimeNumbers(numbers []int) [][]int{	
	digitsMatrix := ejercicio28.ejercicio27.ejercicio6.matrixDigitColection(numbers);
	return arrayMetodos.FilterRowsMatrix(digitsMatrix,numerIfPrime);
}
func(ejercicio28*Ejercicio28) countDigitsPrimeNumbers(numbers []int)([][]int, []int){
	digitsPrimeMatrix := ejercicio28.digitPrimeNumbers(numbers);
	return digitsPrimeMatrix,arrayMetodos.MapTransformMatrixDataInIntArray(digitsPrimeMatrix,ejercicio28.ejercicio27.countNumbers);
}
func(ejercicio28*Ejercicio28) showCountAndMaxDigitsPrime(numbers []int)([][]int, []int, int){
	digitsPrimeMatrix,countdigitPrime := ejercicio28.countDigitsPrimeNumbers(numbers);
	return digitsPrimeMatrix,countdigitPrime, ejercicio28.ejercicio27.ejercicio26.NumberMaxArray(countdigitPrime)
}
func(ejercicio28*Ejercicio28)answer(numbers []int)string{
	digitsPrimeMatrix,countdigitPrime,numberMax := ejercicio28.showCountAndMaxDigitsPrime(numbers);	
	return "Los digitos primos son "+matrixToString(digitsPrimeMatrix,",")+", la cantidad de digitos primos que tiene cada numero es "+ arrayToString(countdigitPrime," ")+",la mayor cantidad de primos es "+ strconv.Itoa(numberMax)  ;	
}

func ejecutarEjercicio28(){
	ejercicio28 := new(Ejercicio28);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"28. Leer 2 números enteros y determinar cual de los dos tiene mayor cantidad de dígitos primos.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio28.answer);
}
/*
29. Leer un número entero y determinar a cuánto es igual el primero de sus dígitos.
*/
type Ejercicio29 struct{
	ejercicio6 Ejercicio6	
}
func(ejercicio29*Ejercicio29)fristNumber(numbers []int)int{
	return numbers[0];
}	
func(ejercicio29*Ejercicio29)fristNumberMatrix(numbers [][]int)[]int{
	return arrayMetodos.MapTransformMatrixDataInIntArray(numbers,ejercicio29.fristNumber);
}
func(ejercicio29*Ejercicio29)theFristdigitOfNumber(numbers []int)[]int{
	return ejercicio29.fristNumberMatrix(ejercicio29.ejercicio6.matrixDigitColection(numbers))
}	
func(ejercicio29*Ejercicio29)answer(numbers []int)string{	
	return "Sus primeros numeros son "+ arrayToString(ejercicio29.theFristdigitOfNumber(numbers)," ") ;	
}

func ejecutarEjercicio29(){
	ejercicio29 := new(Ejercicio29);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"29. Leer un número entero y determinar a cuánto es igual el primero de sus dígitos.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio29.answer);
}
/*
30. Leer un número entero y mostrar todos sus componentes numéricos o sea aquellos para
quienes el sea un múltiplo.*/
type Ejercicio30 struct{
	ejercicio3 Ejercicio3	
}
func(ejercicio30*Ejercicio30)numberMultiple(numbers []int)[][]int{	
	return ejercicio30.ejercicio3.generateDivisorNumbers(numbers);
}
func(ejercicio30*Ejercicio30)answer(numbers []int)string{	
	return "Sus primeros numeros son "+ matrixToString(ejercicio30.numberMultiple(numbers)," ") ;	
}

func ejecutarEjercicio30(){
	ejercicio30 := new(Ejercicio30);
	ejercicicep := ejecutarEjercicio.HeredarObjeto(
		"30. Leer un número entero y mostrar todos sus componentes numéricos o sea aquellos para quienes el sea un múltiplo.",
		3,
		agregarMostrarValores.AgregarValoreVariableEntera);	
	ejercicicep.EjecutarEjercicio(ejercicio30.answer);
}
/*
31. Leer números hasta que digiten 0 y determinar a cuánto es igual el promedio de los números
terminados en 5.
*/
type Ejercicio31 struct{
	ejercicio5 Ejercicio5	
	ejercicio16 Ejercicio16
}


func(ejercicio31*Ejercicio31)averageTerminateInFive(numbers []int)int{	
	ejercicio31.ejercicio5.terminate = 5;
	numberTerminateInfive := ejercicio31.ejercicio5.betweenNumbersTerminate(numbers);
	return ejercicio31.ejercicio16.averageInt(numberTerminateInfive);
}

func(ejercicio31*Ejercicio31)answer(numbers []int)string{	
	return "El promedio de los numeros terminados en  5 son  "+ strconv.Itoa( ejercicio31.averageTerminateInFive(numbers)) ;	
}

func ejecutarEjercicio31(){
	ejercicio31 := new(Ejercicio31);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"31. Leer números hasta que digiten 0 y determinar a cuánto es igual el promedio de los números		terminados en 5.",	
		agregarMostrarValores.AgregarNumerosHastaIngresarCero ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio31.answer);
}
/*
32. Leer números hasta que digiten 0 y determinar a cuanto es igual el promedio entero de los
número primos leídos.*/
type Ejercicio32 struct{
	ejercicio16 Ejercicio16
}

func(ejercicio32*Ejercicio32)averageTerminateInFive(numbers []int)int{	
	numbersPrime :=  arrayMetodos.Filter(numbers,numerIfPrime);	
	return ejercicio32.ejercicio16.averageInt(numbersPrime);
}

func(ejercicio32*Ejercicio32)answer(numbers []int)string{	
	return "El promedio de los numeros primos son  "+ strconv.Itoa( ejercicio32.averageTerminateInFive(numbers)) ;	
}

func ejecutarEjercicio32(){
	ejercicio32 := new(Ejercicio32);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"32. Leer números hasta que digiten 0 y determinar a cuanto es igual el promedio entero de los número primos leídos.",	
		agregarMostrarValores.AgregarNumerosHastaIngresarCero ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio32.answer);
}
/*
33. Si 32768 es el tope superior para los números entero cortos, determinar cuál es el número
primo mas cercano por debajo de él.*/
type Ejercicio33 struct{
	ejercicio1 Ejercicio1
}
func(ejercicio33*Ejercicio33) generateNumbersUp10()[]int{
	numbers := ejercicio33.ejercicio1.generateNumbersUp(32768);
	numbersPrime :=  arrayMetodos.Filter(numbers,numerIfPrime);	
	return numbersPrime;		
}
func(ejercicio33*Ejercicio33)answer()string{
	return "Los numeros primos por debajo de 32768 son "+arrayToString(ejercicio33.generateNumbersUp10(),",");	
}
func ejecutarEjercicio33(){
	ejercicio33 := new(Ejercicio33);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("33. Si 32768 es el tope superior para los números entero cortos, determinar cuál es el número		primo mas cercano por debajo de él..");
	ejercicicep.EjecutarEjercicioSinData(ejercicio33.answer);
}
/*
34. Generar los números del 1 al 10 utilizando un ciclo que vaya de 10 a 1.
*/
type Ejercicio34 struct{
	ejercicio1 Ejercicio1
}
func(ejercicio34*Ejercicio34) generateNumbersUp10()[]int{
	return 	ejercicio34.ejercicio1.generateNumbersUp(10);
}
func(ejercicio34*Ejercicio34)answer()string{
	return "numeros del 1 al 10 "+arrayToString(ejercicio34.generateNumbersUp10(),",");	
}
func ejecutarEjercicio34(){
	ejercicio34 := new(Ejercicio34);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("34. Generar los números del 1 al 10 utilizando un ciclo que vaya de 10 a 1.");
	ejercicicep.EjecutarEjercicioSinData(ejercicio34.answer);
}
/*
35. Leer dos números enteros y determinar a cuánto es igual el producto mutuo del primer dígito
de cada uno.*/
type Ejercicio35 struct{
	ejercicio6 Ejercicio6
}
func(ejercicio35*Ejercicio35)fristNumber(numberOne []int)int{	
	return numberOne[0];
}
func(ejercicio35*Ejercicio35)multipleNumber(numberOne int, numberTwo int)int{
	if 	numberOne == 0 {
		numberOne = 1;
	}
	return numberOne * numberTwo;
}
func(ejercicio35*Ejercicio35)fristDigitNumbersArray(numbers  []int)[]int{
	digitsMatrixColecion := ejercicio35.ejercicio6.matrixDigitColection(numbers);
	return arrayMetodos.MapTransformMatrixDataInIntArray(digitsMatrixColecion,ejercicio35.fristNumber)
}
func(ejercicio35*Ejercicio35)multipleNumberArray(numbers []int)int{	
	return arrayMetodos.ReduceInt(numbers,ejercicio35.multipleNumber);
}
func(ejercicio35*Ejercicio35)multipleFristNumber(numbers []int)int{	
	return  ejercicio35.multipleNumberArray(ejercicio35.fristDigitNumbersArray(numbers));
}
func(ejercicio35*Ejercicio35)answer(numbers []int)string{	
	return "La multiplicacion de los primeros digitos son  "+ strconv.Itoa(ejercicio35.multipleFristNumber(numbers));
}

func ejecutarEjercicio35(){
	ejercicio35 := new(Ejercicio35);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"35. Leer dos números enteros y determinar a cuánto es igual el producto mutuo del primer dígito de cada uno",	
		agregarMostrarValores.AgregarHastaDosNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio35.answer);
}
/*
36. Mostrar en pantalla la tabla de multiplicar del número 5.
*/
type Ejercicio36 struct{
}
func(ejercicio36*Ejercicio36)tableMultiplique(number int)string{	
	var tableMultiple string;
	for start :=  1; start <= 10 ; start++{
		multiple := start*number;
		multiplestring := strconv.Itoa( start) +" x "+ strconv.Itoa( number)+" = "+ strconv.Itoa(multiple )+"\n ";
		tableMultiple = tableMultiple + multiplestring;
	}
	return tableMultiple;
}
func(ejercicio36*Ejercicio36)tablesMultipliques(numbers []int)string{
	var resultado  string;
	for _,element := range numbers{
		fmt.Println(element);

		resultado = resultado + ejercicio36.tableMultiplique(element)+"\n";
	}
	return resultado;
}


func(ejercicio36*Ejercicio36)answer()string{	
	 numbers :=[]int {5} ;
	return "La multiplicacion es   : \n"+ ejercicio36.tablesMultipliques(numbers);
}

func ejecutarEjercicio36(){
	ejercicio36 := new(Ejercicio36);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("36. Mostrar en pantalla la tabla de multiplicar del número 5.");
	ejercicicep.EjecutarEjercicioSinData(ejercicio36.answer);
}
/*
37. Generar todas las tablas de multiplicar del 1 al 10.
*/
type Ejercicio37 struct{
	ejercicio36 Ejercicio36
}

func(ejercicio37*Ejercicio37)answer()string{	
	numbers := []int{1,2,3,4,5,6,7,8,9,10};
	return "La multiplicaciones  del 1 al 10 : \n"+ ejercicio37.ejercicio36.tablesMultipliques(numbers);
}

func ejecutarEjercicio37(){
	ejercicio37 := new(Ejercicio37);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("37. Generar todas las tablas de multiplicar del 1 al 10.");
	ejercicicep.EjecutarEjercicioSinData(ejercicio37.answer);
}
/*
38. Leer un número entero y mostrar en pantalla su tabla de multiplicar.
*/
type Ejercicio38 struct{
	ejercicio36 Ejercicio36
}

func(ejercicio38*Ejercicio38)answer(numbers []int) string{	
	return "La multiplicaciones  del 1 al 10 : \n"+ ejercicio38.ejercicio36.tablesMultipliques(numbers);
}

func ejecutarEjercicio38(){
	ejercicio38 := new(Ejercicio38);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"38. Leer un número entero y mostrar en pantalla su tabla de multiplicar.",	
		agregarMostrarValores.AgregarHastaUnoNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio38.answer);
}
/*
39. Se define la serie de Fibonacci como la serie que comienza con los dígitos 1 y 0 y va sumando
progresivamente los dos últimos elementos de la serie, así:
0 1 1 2 3 5 8 13 21 34.......
Utilizando el concepto de ciclo generar la serie de Fibonacci hasta llegar o sobrepasas el número
10000.
*/
type Ejercicio39 struct{
}
func(ejercicio39*Ejercicio39) fibonacciUntil( conditionUntil func(int,int)bool) []int {
	var colectionFibonanci []int;
	var after int;
	before := 1	 
	var sequence int;
	var stop bool;
	for  index := 0 ;stop == false; index++ {	
		sequence = after + before;
		after = before;
		before = sequence;
		colectionFibonanci = append(colectionFibonanci,sequence);
		stop = conditionUntil(index,sequence);
	}	
	return colectionFibonanci;
}

func(ejercicio39*Ejercicio39)answer()string{
	condition :=  func(index int, until int)bool{
		return until >= 10000 || index == 10000;
	}	
	return "La secuencia fibonanci que existe hasta el  10000 es  : \n"+arrayToString(ejercicio39.fibonacciUntil( condition)," ");
}

func ejecutarEjercicio39(){
	ejercicio39 := new(Ejercicio39);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("39. Se define la serie de Fibonacci como la serie que comienza con los dígitos 1 y 0 y va sumando 		progresivamente los dos últimos elementos de la serie, así:		0 1 1 2 3 5 8 13 21 34.......		Utilizando el concepto de ciclo generar la serie de Fibonacci hasta llegar o sobrepasas el número	10000.");
	ejercicicep.EjecutarEjercicioSinData(ejercicio39.answer);
}
/*
40. Leer un número de dos dígitos y determinar si pertenece a la serie de Fibonacci.*/
type Ejercicio40 struct{
	ejercicio39 Ejercicio39
	ejercicio22 Ejercicio22
	colectionFibonanci []int
}


func(ejercicio40*Ejercicio40)ifFibonancci(number int)bool{
	ejercicio40.ejercicio22.numberSearch = number;
	return len( ejercicio40.ejercicio22.searchNumbersInArray(ejercicio40.colectionFibonanci))> 0;
}
func(ejercicio40*Ejercicio40)createFibonancci(numbers []int)[]int{
	sort.Ints(numbers);	
	numberMax := numbers[len(numbers)-1];
	condition :=  func(index int, until int)bool{
		return until >= numberMax || index == numberMax;
	}
	return ejercicio40.ejercicio39.fibonacciUntil(condition);
}
func(ejercicio40*Ejercicio40)ifFibonancciColectionNumber(numbers []int)[]int{
	ejercicio40.colectionFibonanci = ejercicio40.createFibonancci(numbers);
	return arrayMetodos.Filter(numbers,ejercicio40.ifFibonancci);
}
func(ejercicio40*Ejercicio40)answer(numbers []int)string{
	return "los numeros que pertenecen a la serie de fibonanci : \n"+arrayToString(ejercicio40.ifFibonancciColectionNumber( numbers)," ");
}

func ejecutarEjercicio40(){
	ejercicio40 := new(Ejercicio40);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"40. Leer un número de dos dígitos y determinar si pertenece a la serie de Fibonacci.",	
		agregarMostrarValores.AgregarHastaUnoNumeros ,	
		agregarMostrarValores.IngresarDosDigitos ); 
	ejercicicep.EjecutarEjercicioF(ejercicio40.answer);
}
/*
41. Determinar a cuánto es igual la suma de los elementos de la serie de Fibonacci entre 0 y 100.
*/
type Ejercicio41 struct{
	ejercicio10 Ejercicio10
	ejercicio39	Ejercicio39
	ejercicio1 Ejercicio1
} 
func (ejercicio41*Ejercicio41) generatFibonancci(number int)[]int{
	condition :=  func(index int, until int)bool{
		return until > number || index == number;
	}
	filter := func(numberGet int)bool{
		return numberGet <= number
	}
	return arrayMetodos.Filter( ejercicio41.ejercicio39.fibonacciUntil(condition),filter)
}
func(ejercicio41*Ejercicio41)operationFibonancci(number int, operation func([]int)int)int{
	fibonanci := ejercicio41.generatFibonancci(number);
	fmt.Println(fibonanci);
	return operation(fibonanci);
}
func(ejercicio41*Ejercicio41)sumFibonancci(number int)int{
	return ejercicio41.operationFibonancci(number, ejercicio41.ejercicio10.sumArray);
}
func(ejercicio41*Ejercicio41)sumArrayFibonancci(numbers []int)[]int{
	return arrayMetodos.MapTransformData(numbers,ejercicio41.sumFibonancci);
}
func(ejercicio41*Ejercicio41)sumArrayFibonancciOneHundred()[]int{
	number := []int {100} ;
	return arrayMetodos.MapTransformData(number,ejercicio41.sumFibonancci);
}
func(ejercicio41*Ejercicio41)answer()string{
	return "La suma de sus numeros  de fibonanci es : \n"+arrayToString(ejercicio41.sumArrayFibonancciOneHundred( )," ");
}

func ejecutarEjercicio41(){
	ejercicio41 := new(Ejercicio41);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("41. Determinar a cuánto es igual la suma de los elementos de la serie de Fibonacci entre 0 y 100.",	
		);
	ejercicicep.EjecutarEjercicioSinData(ejercicio41.answer);
}
/*
42. Determinar a cuánto es igual el promedio entero de los elementos de la serie de Fibonacci
entre 0 y 1000.*/

type Ejercicio42 struct{
	ejercicio41	Ejercicio41
	ejercicio16 Ejercicio16
}
func(ejercicio42*Ejercicio42)averageFibonancci(number int)int{
	return ejercicio42.ejercicio41.operationFibonancci(number, ejercicio42.ejercicio16.averageInt);
}
func(ejercicio42*Ejercicio42)averageArrayFibonancciOneThousand()[]int{
	number := []int {1000} ;
	return arrayMetodos.MapTransformData(number,ejercicio42.averageFibonancci);
}
func(ejercicio42*Ejercicio42)answer()string{
	return "El promedio de sus numeros  de fibonanci es : \n"+arrayToString(ejercicio42.averageArrayFibonancciOneThousand()," ");
}

func ejecutarEjercicio42(){
	ejercicio42 := new(Ejercicio42);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("42. Determinar a cuánto es igual el promedio entero de los elementos de la serie de Fibonacci entre 0 y 1000.",	
		);
	ejercicicep.EjecutarEjercicioSinData(ejercicio42.answer);
}
/*
43. Determinar cuántos elementos de la serie de Fibonacci se encuentran entre 1000 y 2000.*/
type Ejercicio43 struct{
	ejercicio4 Ejercicio4	
	ejercicio40 Ejercicio40
}
func(ejercicio43*Ejercicio43)ifFibonancciColectionNumber(numbers []int)[]int{
	return arrayMetodos.Filter(numbers,ejercicio43.ejercicio40.ifFibonancci);
}

func(ejercicio43*Ejercicio43)fibonanciBetween(numbers []int)[][]int{
	ejercicio43.ejercicio40.colectionFibonanci = ejercicio43.ejercicio40.createFibonancci(numbers);
 	return	arrayMetodos.MapTransformMatrixDataInMatrix( ejercicio43.ejercicio4.generateBetweenNumbers(numbers ), ejercicio43.ifFibonancciColectionNumber);
}

func(ejercicio43*Ejercicio43)answer()string{
	numbers := []int{1000, 2000};
	return "Los numeros fibonanci entre "+arrayToString(numbers,",")+": \n"+matrixToString(ejercicio43.fibonanciBetween(numbers )," ");
}

func ejecutarEjercicio43(){
	ejercicio43 := new(Ejercicio43);
	ejercicicep := ejecutarEjercicio.HeredarObjetoNoIngresarDatos("43. Determinar cuántos elementos de la serie de Fibonacci se encuentran entre 1000 y 2000");
	ejercicicep.EjecutarEjercicioSinData(ejercicio43.answer);
}
/*
44. Leer un número y calcularle su factorial.
*/
type Ejercicio44 struct{
	ejercicio1 Ejercicio1
	ejercicio35 Ejercicio35
}


func(ejercicio44*Ejercicio44)factorialNumber(number int) int{
	return ejercicio44.ejercicio35.multipleNumberArray(ejercicio44.ejercicio1.generateGrowthOfNumbers(number , 1 ));
}
func(ejercicio44*Ejercicio44)factorialsNumbers(numbers []int) []int{
	return arrayMetodos.MapTransformData(numbers,ejercicio44.factorialNumber);
}
func(ejercicio44*Ejercicio44)answer(numbers []int)string{
	return "los factoriales de los numeros son : \n"+arrayToString(ejercicio44.factorialsNumbers( numbers)," ");
}

func ejecutarEjercicio44(){
	ejercicio44 := new(Ejercicio44);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"44. Leer un número y calcularle su factorial.",	
		agregarMostrarValores.AgregarHastaTresNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio44.answer);
}
/*
45. Leer un número y calcularle el factorial a todos los enteros comprendidos entre 1 y el número
leído.*/
type Ejercicio45 struct{
	ejercicio44 Ejercicio44
}
func(ejercicio45*Ejercicio45)generateNumberOfoneTo(number int) []int{
	return ejercicio45.ejercicio44.ejercicio1.generateGrowthOfNumbers(number , 1 );
}
func(ejercicio45*Ejercicio45)factorialNumberOfoneTo(numbers []int) [][]int{
	betweenNumbers := arrayMetodos.MapTransformDataConvertToMatrix(numbers,ejercicio45.generateNumberOfoneTo);
	return arrayMetodos.MapTransformMatrixDataInMatrix(betweenNumbers,ejercicio45.ejercicio44.factorialsNumbers);
}
func(ejercicio45*Ejercicio45)answer(numbers []int)string{
	return "los factoriales de los numeros son : \n"+matrixToString(ejercicio45.factorialNumberOfoneTo( numbers)," ");
}
func ejecutarEjercicio45(){
	ejercicio45 := new(Ejercicio45);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"45. Leer un número y calcularle el factorial a todos los enteros comprendidos entre 1 y el número	leído.",	
		agregarMostrarValores.AgregarHastaTresNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio45.answer);
}
/*
46. Leer un número entero y calcular el promedio entero de los factoriales de los enteros
comprendidos entre 1 y el número leído.*/
type Ejercicio46 struct{
	ejercicio45 Ejercicio45
	ejercicio16 Ejercicio16
}

func(ejercicio46*Ejercicio46)averagefactorialNumberOfoneTo(numbers []int)[]int{
	return arrayMetodos.MapTransformMatrixDataInIntArray(ejercicio46.ejercicio45.factorialNumberOfoneTo(numbers),ejercicio46.ejercicio16.averageInt)
}
func(ejercicio46*Ejercicio46)answer(numbers []int)string{
	return "los el promedio de los factoriales de los numeros son : \n"+arrayToString(ejercicio46.averagefactorialNumberOfoneTo( numbers)," ");
}
func ejecutarEjercicio46(){
	ejercicio46 := new(Ejercicio46);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"46. Leer un número entero y calcular el promedio entero de los factoriales de los enteros comprendidos entre 1 y el número leído.",	
		agregarMostrarValores.AgregarHastaTresNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio46.answer);
}
/*
47. Leer un número entero y calcular a cuánto es igual la sumatoria de todos los factoriales de los
números comprendidos entre 1 y el número leído.
*/
type Ejercicio47 struct{
	ejercicio10 Ejercicio10
	ejercicio45 Ejercicio45
}

func(ejercicio47*Ejercicio47)sumfactorialNumberOfoneTo(numbers []int)[]int{
	return arrayMetodos.MapTransformMatrixDataInIntArray(ejercicio47.ejercicio45.factorialNumberOfoneTo(numbers),ejercicio47.ejercicio10.sumArray);
}
func(ejercicio47*Ejercicio47)answer(numbers []int)string{
	return "los el promedio de los factoriales de los numeros son : \n"+arrayToString(ejercicio47.sumfactorialNumberOfoneTo( numbers)," ");
}
func ejecutarEjercicio47(){
	ejercicio47 := new(Ejercicio47);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"47. Leer un número entero y calcular a cuánto es igual la sumatoria de todos los factoriales de los números comprendidos entre 1 y el número leído.",	
		agregarMostrarValores.AgregarHastaTresNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio47.answer);
}
/*
48. Utilizando ciclos anidados generar las siguientes parejas de enteros
0 1
1 1
2 2
3 2
4 3
5 3
6 4
7 4
8 5
9 5
*/
type Ejercicio48 struct{
}
func(ejercicio48*Ejercicio48)generateIntegerUntil(number int,since int,until int)[]int{
	var colection  = make([]int, number);
	sequence := since;
	for start :=0; start < number; start++{
		if start%until == 0 {
			sequence = since;
		}
		colection[start] =  sequence;
		fmt.Println("sequence ",sequence);
		sequence++;
	}
	return colection;
}
func(ejercicio48*Ejercicio48)generateIntegerSequences(number int,secuence int)[]int{
	var colection  = make([]int, number);
	var colectionCount int;
	for start :=0; start < number; start++{
		if start%secuence == 0 {
			colectionCount++
		}
		colection[start] =  colectionCount;
	}
	return colection;
}
func(ejercicio48*Ejercicio48)generateIntegerPair(number int)[]int{
	return ejercicio48.generateIntegerSequences(number+1,2)
}
func(ejercicio48*Ejercicio48)generateUntil(numbers []int)[][]int{
	until := func(number int)[]int{
		return ejercicio48.generateIntegerUntil(number+1,0, number+1);
	}
	return arrayMetodos.MapTransformDataConvertToMatrix(numbers,until);
}
func(ejercicio48*Ejercicio48)generateNumbersIntegerPair(numbers []int)[][]int{
	return arrayMetodos.MapTransformDataConvertToMatrix(numbers,ejercicio48.generateIntegerPair);
}
func(ejercicio48*Ejercicio48)answer(numbers []int)string{
	return "La pareja de enteros de los numeros son : \n"+matrixToString( ejercicio48.generateUntil(numbers)," ")+"\n"+matrixToString(ejercicio48.generateNumbersIntegerPair( numbers)," ");
}
func ejecutarEjercicio48(){
	ejercicio48 := new(Ejercicio48);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"48. Utilizando ciclos anidados generar las siguientes parejas de enteros",	
		agregarMostrarValores.AgregarHastaTresNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio48.answer);
}
/*
49. Utilizando ciclos anidados generar las siguientes ternas de números
1 1 1
2 1 2
3 1 3
4 2 1
5 2 2
6 2 3
7 3 1
8 3 2
9 3 3*/
type Ejercicio49 struct{
	ejercicio48 Ejercicio48
}
func(ejercicio49*Ejercicio49)generateUntil(numbers []int)[][]int{
	until := func(number int)[]int{
		return ejercicio49.ejercicio48.generateIntegerUntil(number,1, number);
	}
	return arrayMetodos.MapTransformDataConvertToMatrix(numbers,until);
}
func(ejercicio49*Ejercicio49)generateUntilThree(numbers []int)[][]int{
	until := func(number int)[]int{
		return ejercicio49.ejercicio48.generateIntegerUntil(number,1, 3);
	}
	return arrayMetodos.MapTransformDataConvertToMatrix(numbers,until);
}
func(ejercicio49*Ejercicio49)generateIntegerOdd(number int)[]int{
	return ejercicio49.ejercicio48.generateIntegerSequences(number,3);
}
func(ejercicio49*Ejercicio49)generateNumbersIntegerOdd(numbers []int)[][]int{
	return arrayMetodos.MapTransformDataConvertToMatrix(numbers,ejercicio49.generateIntegerOdd);
}
func(ejercicio49*Ejercicio49)answer(numbers []int)string{
	untilNumber := matrixToString(ejercicio49.generateUntil(numbers)," ");
	numberOdd := matrixToString(ejercicio49.generateNumbersIntegerOdd(numbers)," ");
	untilThree := matrixToString(ejercicio49.generateUntilThree(numbers)," ");
	return "Generando la siguiente terna de numeros : \n"+untilNumber+"\n"+numberOdd+"\n"+untilThree+"\n";
}
func ejecutarEjercicio49(){
	ejercicio49 := new(Ejercicio49);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"49. Utilizando ciclos anidados generar las siguientes ternas de números",	
		agregarMostrarValores.AgregarHastaTresNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio49.answer);
}
/*
50. Utilizando ciclos anidados generar las siguientes parejas de números
0 1
1 1
2 1
3 1
4 2
5 2
6 2
7 2
*/

type Ejercicio50 struct{
	ejercicio48 Ejercicio48
}
func(ejercicio50*Ejercicio50)generateUntil(numbers []int)[][]int{
	until := func(number int)[]int{
		return ejercicio50.ejercicio48.generateIntegerUntil(number+1,0, number+1);
	}
	return arrayMetodos.MapTransformDataConvertToMatrix(numbers,until);
}
func(ejercicio50*Ejercicio50)generateIntegerFour(number int)[]int{
	return ejercicio50.ejercicio48.generateIntegerSequences(number+1,4);
}
func(ejercicio50*Ejercicio50)generateNumbersIntegerFour(numbers []int)[][]int{
	return arrayMetodos.MapTransformDataConvertToMatrix(numbers,ejercicio50.generateIntegerFour);
}
func(ejercicio50*Ejercicio50)answer(numbers []int)string{
	untilNumber := matrixToString(ejercicio50.generateUntil(numbers)," ");
	numberFour := matrixToString(ejercicio50.generateNumbersIntegerFour(numbers)," ");
	return "Generando la siguiente duo de numeros : \n"+untilNumber+"\n"+numberFour;
}
func ejecutarEjercicio50(){
	ejercicio50 := new(Ejercicio50);
	ejercicicep :=  ejecutarEjercicio.HeredarObjetoF(
		"50. Utilizando ciclos anidados generar las siguientes parejas de números",	
		agregarMostrarValores.AgregarHastaTresNumeros ,	
		agregarMostrarValores.AgregarValoreVariableEntera ); 
	ejercicicep.EjecutarEjercicioF(ejercicio50.answer);
}


func EjecutarPrograma(){
	controlarFuncionamiento := controlarFuncionamientoPrograma.ConstruirTitulo("ciclos");	
	fmt.Println("Hola, bienvenido al programa de ciclos");
	controlarFuncionamiento.Administrar(ejercicios);
}

type  ObjetosCiclos  struct {		
	Ejercicio26 Ejercicio26				
}

func ExportarObjetosCiclos() ObjetosCiclos{
	return ObjetosCiclos{}
}

