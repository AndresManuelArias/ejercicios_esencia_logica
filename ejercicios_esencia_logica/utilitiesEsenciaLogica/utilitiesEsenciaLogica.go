package utilitiesEsenciaLogica

import ( 
	"fmt"
	"math"
	"sort"
	// "strconv"
	//"strings"
	"../arrayMetodos"
)


func ContarDigitosLog(numeroEntero int) int{
	digitos := int(math.Log10(float64(numeroEntero)))+1;
	if digitos <  0{
		digitos = 1;
	}
	return digitos;
}
func  ArrayDigitosNumero(valorEntero int) []int{	
	var particionado = valorEntero;

	var colecionDigitos = make([]int,ContarDigitosLog(valorEntero));			
	for  position := len(colecionDigitos)-1; position >= 0 ;position--{
		sobrante := particionado%10;
		colecionDigitos[position] = sobrante;	
		particionado = particionado / 10 ;	
	}
	return colecionDigitos;
}

func ShearIndexNumbers( numbers []int, shear int )[]int{
	shearNumber := func (number int) bool{
		return shear == number
	}
	return arrayMetodos.FindArrayIndex(numbers,shearNumber);
}
func ColectionNumber(number int)func(int)[]int{
	var colection []int;
	return func (number int)[]int{
		colection = append(colection,number);
		return colection;
	}
}
func saveSearchValue(numbers []int,function func([]int)int)func(int)bool{
	numberSearch := function(numbers);
	return func(number int)bool{
		return number == numberSearch;
	}		
}
func NumberSame(numberOne int,numberTwo int ) bool {
	return numberOne ==	numberTwo;
}
func UniqueNumbers(numbers []int) []int {
	colectionNumber := ColectionNumber(numbers[0]);
	var colection []int;
	comprobateNumber := func(number int){
		if	!arrayMetodos.Some( colection,func(number2 int)bool{return number2 == number }){
			colection = colectionNumber(number)
		}
	}
	for _,number := range numbers{
		comprobateNumber(number);
   	}
	return 	colection
}
func SameNumbers(numbers []int) []int {
	return UniqueNumbers(arrayMetodos.CompareAllTheArray(numbers,NumberSame));
}

func NumberMaxMatrix(numeros [][]int) int {
	numbersMax := arrayMetodos.ApplyRowCol(numeros,NumberMaxArray)
	return NumberMaxArray(numbersMax);
}
func NumberMinMatrix(numeros [][]int) int {
	numbersMax := arrayMetodos.ApplyRowCol(numeros,NumberMinArray)
	return NumberMinArray(numbersMax);
}
func NumberMaxArray(numeros []int) int {
	shearNumberMax := func(numberMax int, numberMin  int)int{
		if numberMax < numberMin {
			numberMax =  numberMin;
			fmt.Println("numeroMayor1",numberMax)			
		}
		return numberMax
	}
	return arrayMetodos.ReduceIntT(numeros,shearNumberMax);
}
func NewNumberMax(numbers []int)ShowInNumbers{
	numberMax := saveSearchValue(numbers,NumberMaxArray)
	return ShowInNumbers{shearNumber:numberMax};
}
func NumberMinArray(numeros []int) int {
	shearNumberMin := func(numberMax int, numberMin  int)int{		
		if numberMax < numberMin {
			numberMin = numberMax;
			// fmt.Println("numeroMayor1",numeroMayor)			
		}
		return numberMin
	}
	return arrayMetodos.ReduceIntT(numeros,shearNumberMin);
}
func NewNumberMin(numbers []int)ShowInNumbers{
	numberMin := saveSearchValue(numbers,NumberMinArray)
	return ShowInNumbers{shearNumber:numberMin};
}

func numberMore(number int,numberThis int)bool{
	return number > numberThis;
}
func numberMin(number int,numberThis int)bool{
	return number < numberThis;
}
func countsDigit(number int)int{
	return len(ArrayDigitosNumero(number))
}

func countDigit(digit int,option func(int,int)bool)func( int)bool{
	return func(number int)bool{
		return option( len(ArrayDigitosNumero(number)),digit);
	}
}
func NewCountMoreDigit(shearchDigit int)ShowInNumbers{
	shearchCountDigit := countDigit(shearchDigit,numberMore);
	return newShowInNumbers(shearchCountDigit);
}
func NewCountMinDigit(shearchDigit int)ShowInNumbers{
	shearchCountDigit := countDigit(shearchDigit,numberMin);
	return newShowInNumbers(shearchCountDigit);
}
func multipleNumber(numberOne int, numberTwo int)int{
	if 	numberOne == 0 {
		numberOne = 1;
	}
	return numberOne * numberTwo;
}
func MultipleNumberArray(numbers []int)int{	
	return arrayMetodos.ReduceInt(numbers,multipleNumber);
}
func unionNumber(before int,after int)int{
	return before + after;
}
func SumArray(number[]int)int{
	return arrayMetodos.ReduceInt(number,unionNumber);
}
func SumDigitArray(numbers[]int)[]int{
	matrixDigitos := arrayMetodos.MapTransformDataConvertToMatrix(numbers,ArrayDigitosNumero);
	return arrayMetodos.MapTransformMatrixDataInIntArray(matrixDigitos,SumArray);	
}
func NewSumDigitArray()NumberOperate{
	return NumberOperate{function:SumDigitArray}
}
func CounDigitArray(numbers[]int)[]int{
	return arrayMetodos.MapTransformData(numbers,countsDigit)
}
func NewCounDigitArray()NumberOperate{
	return NumberOperate{function:CounDigitArray}
}
func SumDigitEven(number int)int{
	arrayDigit := ArrayDigitosNumero(number);
	arrayDigitEven :=arrayMetodos.Filter(arrayDigit,EvenNumber);
	return SumArray(arrayDigitEven);
}



type NumberOperate struct {
	function func([]int)[]int
}
func(numberOperate*NumberOperate)IndexMaxArray(numbers []int)[]int{
	transforNumber := numberOperate.function(numbers);
	maxNumber := NumberMaxArray(transforNumber); 
	return arrayMetodos.FindArrayIndex(transforNumber,func(number int )bool{return maxNumber == number})
}
func(numberOperate*NumberOperate)MaxArray(numbers []int)[]int{
	transforNumber := numberOperate.function(numbers);	
	max := NumberMaxArray( transforNumber);
	return arrayMetodos.Filter(transforNumber,func(number int)bool{return number == max;})
}

func newSumArray(numbers[]int)ShowInNumbers{
	numbersSum := saveSearchValue(numbers,SumArray);
	return ShowInNumbers{shearNumber:numbersSum};
}

func AverageArray(number[]int)int{
	return SumArray(number)/len(number);
}

func NewAverageArray(numbers[]int)ShowInNumbers{
	numbersAverage := saveSearchValue(numbers,AverageArray);
	return ShowInNumbers{shearNumber:numbersAverage};
}

func shearchMultiple(number int,multiple int)bool{
	return 	number%multiple == 0;
}

func NewMultiples(numberMultiple int)ShowInNumbers{
	multiples := func( numbersArray int)bool{
		return shearchMultiple(numbersArray , numberMultiple )		
	}
	return ShowInNumbers{shearNumber:multiples};
}
func numberNegative(number int )bool{
	return number < 0;
}
func NewNumberNegative()ShowInNumbers{
	return ShowInNumbers{shearNumber:numberNegative};
}
func numberPositive(number int )bool{
	return number > 0;
}
func NewNumberPositive()ShowInNumbers{
	return ShowInNumbers{shearNumber:numberPositive};
}
func NewNumberPrime()ShowInNumbers{
	return ShowInNumbers{shearNumber:IfPrimeInt};
}
type ShowInNumbers struct{
	shearNumber func(int)bool
	shearNumberInt func([]int)int
}
func(showInNumbers*ShowInNumbers)ShearchNumbersInArray(numbers []int)[]int{
	return arrayMetodos.Filter(numbers,showInNumbers.shearNumber);	
}
func(showInNumbers*ShowInNumbers)IndexNumberInArray(numbers []int) []int{
	return arrayMetodos.FindArrayIndex(numbers,showInNumbers.shearNumber);	
}

func(showInNumbers*ShowInNumbers)IndexNumbersConditionArray(numbers []int) []int{
	conditionNumbers := showInNumbers.ShearchNumbersInArray(numbers);
	conditionNumbersMax := showInNumbers.shearNumberInt(conditionNumbers);
	return ShearIndexNumbers(numbers,conditionNumbersMax);
}
func newShowInNumbers(shearNumber func(int)bool)ShowInNumbers{
	return ShowInNumbers{shearNumber:shearNumber};
}

func NewShearNumber(shearNumber func(int)bool)ShowInNumbers{
	return ShowInNumbers{shearNumber:shearNumber};
}

func EvenNumber(number int)bool{
	return number % 2 == 0;
}
func NewEvenNumberMax()ShowInNumbers{
	return ShowInNumbers{shearNumber:EvenNumber,shearNumberInt:NumberMaxArray};
}

func IfPrime(number int)bool{
	prime := true;
	var n int;
	var count int;
	for  n = 1; n <= number;n++{
		if number % n ==  0 {
			count++
		}
		if count > 2 {
			prime = false;
			break;
		}
	}
	return prime;
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
func IfPrimeInt(numero int )bool{
	return ifPrimeFloat(float64(numero));
}	
func NewPrimeNumberMax()ShowInNumbers{
	return ShowInNumbers{shearNumber:IfPrimeInt,shearNumberInt:NumberMaxArray};
}
func NewPrimeNumberMin()ShowInNumbers{
	return ShowInNumbers{shearNumber:IfPrimeInt,shearNumberInt:NumberMinArray};
}

func GenerateFibonacci() func()int {
	var after int;
	before := 1	 
	var sequence int;
	return func ()int{
		sequence = after + before;
		after = before;
		before = sequence;
		return sequence;
	}
}

func GeneratePrimes() func()int {
	var prime int;
	return func ()int{
		prime++;
		for ;IfPrimeInt(prime) == false;{
			prime++;
		}		
		return prime;
	}
}
func GenerateNumber(count int) func()int {
	var number int;
	return func ()int{
		number = number+ count
		return number;
	}
}

func generateSeveralNumbersWithInformation(until int, condition func()int)[]int{
	var colection []int;
	var creation int;
	for; creation < until; {
		creation = condition();
		colection = append(colection,creation);		
	}
	return colection;
}
func GenerateNumbersbetweenNumbers(numbers []int,condition func()func()int)[][]int{
	sort.Ints(numbers);	
	var colectionNumber  [][]int;
	generator := condition();	
	for _,element := range numbers{
		colectionNumber = append(colectionNumber,generateSeveralNumbersWithInformation(element,generator));
	}
	return colectionNumber;
}
func GenerateNumberUntil(until int, condition func()func()int)[]int{
	var colection  = make([]int, until);
	generator := condition();
	for start :=1; start <= until; start++{
		colection[start-1] = generator();
	}
	return colection;
}
func GenerateNumbersUntilNumber(until int, condition func()func()int)[]int{
	var colection []int;
	var creation int;
	generator := condition();
	for; creation < until; {
		creation = generator() 
		colection = append(colection,creation);		
	}
	return colection;
}

func LastDigitNumber(number int)int{
	return number %10 ;
}
func LastDigitNumberCountDigit(number int,count int)int{
	return number % int( math.Pow(10,float64( count)) );
}
func StartDigitNumberCountDigit(number int,count int)int{
	cantidadDigitos  := ContarDigitosLog(number)-count;
	return number / int(math.Pow(10,float64(cantidadDigitos)));	
}
type ShowDigit struct{
	ultimeDigit bool
}
func (showDigit*ShowDigit)showCousinOrLastDigit(number int)int{
	digits := ArrayDigitosNumero(number);
	var positionDigit int;
	if showDigit.ultimeDigit{
		positionDigit = digits[len(digits)-1];
	}else{
		positionDigit = digits[0];
	}
	return positionDigit;
}
func (showDigit*ShowDigit)showCousinOrLastDigitCount(number int,count int)int{
	digits := ArrayDigitosNumero(number);
	var positionDigit []int;
	if count > len(digits){
		count = len(digits);
	}	
	if showDigit.ultimeDigit{
		positionDigit = digits[(len(digits)- count):(len(digits)-1)];
	}else{
		positionDigit = digits[0:count];
	}
	reduceNumber := arrayMetodos.ReduceInt( positionDigit,func(number1 int,number2 int)int{
		number := number1 * 10;
		return number + number2; 
	});
	return reduceNumber;				
}
func (showDigit*ShowDigit)transforNumbersCousinOrLastCount(numbers []int,count int )[]int{
	tranformNumbers := func(number int)int{
		return	 showDigit.showCousinOrLastDigitCount(number,count);
	}
	return  arrayMetodos.MapTransformData(numbers,tranformNumbers);		
}
func (showDigit*ShowDigit)transforNumbersCousinOrLast(numbers []int )[]int{
	tranformNumbers := func(number int)int{
		return	 showDigit.showCousinOrLastDigit(number);
	}
	return  arrayMetodos.MapTransformData(numbers,tranformNumbers);		
}

func (showDigit*ShowDigit)NumbersDigitCousinOrLast(numbers []int,shearchNumber int)[]int{
	digits := showDigit.transforNumbersCousinOrLast(numbers);
	shearDigit := func(number int)bool{			
		return number== shearchNumber;
	}
	return arrayMetodos.Filter(digits,shearDigit);	
}
func (showDigit*ShowDigit)NumbersDigitCousinOrLastCount(numbers []int,shearchNumber int,count int)[]int{
	digits := showDigit.transforNumbersCousinOrLastCount(numbers,count);			
	shearDigit := func(number int)bool{			
		return number == shearchNumber;
	}
	return arrayMetodos.Filter(digits,shearDigit);	
}
func (showDigit*ShowDigit)IndexNumbersDigitCousinOrLast(numbers []int,shearchNumber int) []int{
	digits := showDigit.transforNumbersCousinOrLast(numbers);	
	return ShearIndexNumbers(digits,shearchNumber);
}

func NewShowDigitShowDigit(ultimeDigit bool)ShowDigit{
	return ShowDigit{ultimeDigit:ultimeDigit}
}

