// agregar y mostrar datos
package agregarMostrarValores

import ( "fmt"
	"math"
)

func AgregarValoresVariablesABCD()(float64,float64,float64,float64){
	var a ,b ,c,d float64;
	a ,b ,c =  AgregarValoresVariablesABC()
	fmt.Println("escriba el valor de la letra d")
	fmt.Scanf("%f\n", &d)  
	return a,b,c,d
}
func AgregarValoresVariablesABC()(float64,float64,float64){
	var a ,b ,c float64;
	a,b =  AgregarValoresVariablesAB()
	fmt.Println("escriba el valor de la letra c")
	fmt.Scanf("%f\n", &c) 
	return a,b,c
}
func AgregarValoresVariablesAB()(float64,float64){
	var a ,b  float64;
	fmt.Println("escriba el valor de la letra a")
	fmt.Scanf("%f\n", &a) 
	fmt.Println("escriba el valor de la letra b")
	fmt.Scanf("%f\n", &b) 
	return a,b
}
func AgregarValoreVariableEntera()int{
	var a   int;
	fmt.Println("escriba el valor del numero entero")
	fmt.Scanf("%d\n", &a) 
	return a
}
func DeterminarCuantosDatosIngresar(numeroDatosIngresar int,funcionIngresarDatos func()(  int )) []int {
	var numeroGuardado []int;
	var numeroGuardado1 = make([]int,  numeroDatosIngresar);
	
	for ingreso := 1;ingreso <= numeroDatosIngresar; ingreso++ {	
		numeroGuardado1[ingreso-1] =	funcionIngresarDatos();
		fmt.Println("Numeros ingresados ",numeroGuardado1)
	}
	numeroGuardado = numeroGuardado1[0:numeroDatosIngresar];	
	return numeroGuardado;
}
func AgregarValoreVariableEnteraMatrix(filas int,columnas int)[][]int{
	var colectionNumber = make([][]int,filas);
	for fila := 0; fila < filas; fila++ {
		for columna := 0; columna < columnas; columna++ {
			fmt.Println("--fila:",fila,"Columna:",columna,"--")	
			colectionNumber[fila] =	append(colectionNumber[fila],AgregarValoreVariableEntera());
		}				
	}
	return colectionNumber;
}
func DatosIngresar(numeroDatosIngresar func([]int) bool,funcionIngresarDatos func()(  int )) []int {
	var numeroGuardado []int;
	var terminarEjercicios bool;	
	for ingreso := 0;terminarEjercicios == false;ingreso++   {	
		numeroGuardado =	append(numeroGuardado,funcionIngresarDatos());
		fmt.Println("Numeros ingresados ",numeroGuardado);
		terminarEjercicios = numeroDatosIngresar(numeroGuardado);
	}	
	return numeroGuardado;
}

func AgregarNumerosHastaIngresarCero( numeroIngresado []int )bool{
	var terminarEjercicio bool;		
	if numeroIngresado[len(numeroIngresado)- 1 ] == 0 {
		terminarEjercicio = true;
	}
	return terminarEjercicio;
}
func AgregarHastaTantosNumeros( numeroIngresado []int, cantidad int )bool{
	var terminarEjercicio bool;		
	if len(numeroIngresado) == cantidad {
		terminarEjercicio = true;
	}
	return terminarEjercicio;
}
func AgregarHastaUnoNumeros( numeroIngresado []int )bool{
	return AgregarHastaTantosNumeros( numeroIngresado, 1  );
}
func AgregarHastaDosNumeros( numeroIngresado []int )bool{
	return AgregarHastaTantosNumeros( numeroIngresado, 2  );
}
func AgregarHastaTresNumeros( numeroIngresado []int )bool{
	return AgregarHastaTantosNumeros( numeroIngresado, 3  );
}
func AgregarHastaDiezNumeros( numeroIngresado []int )bool{
	return AgregarHastaTantosNumeros( numeroIngresado, 10  );
}
func SinDigitos()int {
	return AgregarValoreVariableEnteraDefinirCantidadDigito(0);
}
func IngresarDosDigitos()int {
	return AgregarValoreVariableEnteraDefinirCantidadDigito(2);
}
func IngresarTresDigitos()int {
	return AgregarValoreVariableEnteraDefinirCantidadDigito(3);
}
func IngresarCuatroDigitos()int {
	return AgregarValoreVariableEnteraDefinirCantidadDigito(4);
}
func IngresarCincoDigitos()int {
	return AgregarValoreVariableEnteraDefinirCantidadDigito(5);
}
func AgregarValoreVariableEnteraDefinirCantidadDigito(cantidadDigitos int)int{
	 cantidadDigitosApropiada := true;
	var numero int;
	for cantidadDigitosApropiada {
		numero = AgregarValoreVariableEntera();
		if !determinarCantidadDigitos(numero,cantidadDigitos) {
			fmt.Println("El numero no posee la cantidad de digitos exigida que es", cantidadDigitos ," y este posee ",contarDigitos(numero)," digitos");
		}else{
			cantidadDigitosApropiada = false;
			break;
		}
	}	
	return numero;
}

func NumeroPositivoMenorDe100()int{
	return	AgregarValoreVariableEnteraDefinirCondicion(condicionEntradaNumeroPositivoMenor100);
}
func condicionEntradaNumeroPositivoMenor100(numero int) bool {
	return numero >0 && numero <=100;
}
func NumeroPositivoMenorDe50()int{
	return	AgregarValoreVariableEnteraDefinirCondicion(condicionEntradaNumeroPositivoMenor50);
}
func condicionEntradaNumeroPositivoMenor50(numero int) bool {
	return numero >0 && numero <=50;
}
func NumeroPositivoMenorDeMil()int{
	return	AgregarValoreVariableEnteraDefinirCondicion(condicionEntradaNumeroPositivoMenor1000);
}
func condicionEntradaNumeroPositivoMenor1000(numero int) bool {
	return numero >0 && numero <=1000;
}
func AgregarValoreVariableEnteraDefinirCondicion(  determinarCondicion func(int) bool  )int{
	condicionNoCumplida := true;
	var numero int;
	for condicionNoCumplida {
	   numero = AgregarValoreVariableEntera();
	   if !determinarCondicion(numero) {
		   fmt.Println("No cumple con la condicion");
	   }else{
		   condicionNoCumplida = false;
		   break;
		}
	}	
	return numero;
}
	
func determinarCantidadDigitos(numero int,cantidadDigitos int) bool{
	return contarDigitos(numero) == cantidadDigitos;
}
func RespuestaEjercicio(valor float64) {
	fmt.Println("El resultado es ",valor)		
}
func RespuestaEjercicioString( respuesta string) {
	fmt.Println(respuesta)
}

type Ejercicio struct{
	titulo string
	valorEntero int
	respuestaEntero int
}
func ClaseEjercicio(titulo string)Ejercicio{
	return Ejercicio{titulo:titulo}
}
func (ejercicio*Ejercicio) RespuestaEjercicioEntero() {
	fmt.Println("El resultado es ",ejercicio.respuestaEntero);		
}
func (ejercicio Ejercicio) MostrarTituloEjercicio() {
	fmt.Println(ejercicio.titulo);
}
func (ejercicio*Ejercicio) GuardarValorRespuestaEntero(numero int) {
	ejercicio.respuestaEntero = numero;
}

func contarDigitos(numeroEntero int) int{
	digitos := int(math.Log10(float64(numeroEntero)))+1;
	// fmt.Println("digitos log ",math.Log10(float64(numeroEntero)));
	return digitos;
}
