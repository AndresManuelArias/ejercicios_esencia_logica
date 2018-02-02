// los ejercicios de condicionales comienzan en la pagina 140, capitulo 8
/* si quiero exportar algunas funciones en ves de agregarle la primera inicial mayuscula
 lo mejor es crear una nueva funcion que contenga esa inicial mayuscula */
package condicionales

import ( "fmt"
"math"
"strconv"
"strings"
"../controlarFuncionamientoPrograma"
"../agregarMostrarValores"
"../arrayMetodos"
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
		ejecutarEjercicio50	}
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



func ejecutarEjercicio(titulo string, funcion func()(string,int)){
	imprimirResultado := agregarMostrarValores.ClaseEjercicio(titulo);
	imprimirResultado.MostrarTituloEjercicio();
	respuestaTexto, respuestaNumero := funcion();
	fmt.Println(respuestaTexto) ;
	imprimirResultado.GuardarValorRespuestaEntero(respuestaNumero);
	imprimirResultado.RespuestaEjercicioEntero();
}
type GuardarResultadosMostrar struct  {
	cantidadDatosIngresar int
	valorIngresar int
	valoresIngresar []int
	funcionRecibirArrayInt func([]int) (bool, int)
	funcionMostrarRespuestaUsuario func(int) string
	cantidadDigitosPermitidos func() int
	numeroresultado int
	respuestaPositiva string
	respuestaNegativa string
}
func  (guardarResultadosMostrar*GuardarResultadosMostrar) ejecutarEjercicioGuardarResultado() bool{	
	var respuesta bool;
	respuesta,guardarResultadosMostrar.numeroresultado = guardarResultadosMostrar.funcionRecibirArrayInt(guardarResultadosMostrar.valoresIngresar);
	if guardarResultadosMostrar.respuestaPositiva == "" {
		guardarResultadosMostrar.respuestaPositiva = guardarResultadosMostrar.funcionMostrarRespuestaUsuario(guardarResultadosMostrar.numeroresultado);
	}
   	return respuesta;
}
func (guardarResultadosMostrar*GuardarResultadosMostrar) mostrarRespuestaUsuario() string{
   var resultado = guardarResultadosMostrar.respuestaNegativa;	
   if guardarResultadosMostrar.ejecutarEjercicioGuardarResultado() {
	   resultado = guardarResultadosMostrar.respuestaPositiva;
   }
   return resultado;
}
func  (guardarResultadosMostrar*GuardarResultadosMostrar) ingresarValoresMostrarResultados()(string, int)  {
   guardarResultadosMostrar.valoresIngresar = agregarMostrarValores.DeterminarCuantosDatosIngresar(guardarResultadosMostrar.cantidadDatosIngresar,guardarResultadosMostrar.cantidadDigitosPermitidos);
   return guardarResultadosMostrar.mostrarRespuestaUsuario(),guardarResultadosMostrar.numeroresultado;
}
/*
1. Leer un número entero y determinar si es un número terminado en 4.
*/

type Ejercicio1 struct{
	valorEntero int
	numeroresultado int
}
func (ejercicio1*Ejercicio1) getAgregarValor( numero int){
	ejercicio1.valorEntero = numero;
}
func (ejercicio1*Ejercicio1) sacarUltimoNumero()int{
	 return ejercicio1.valorEntero %  10; 
}
func (ejercicio1*Ejercicio1) sacarUltimoNumeroGuardarResultado()int{
	ejercicio1.numeroresultado = ejercicio1.sacarUltimoNumero(); 
	return ejercicio1.numeroresultado; 
}
func (ejercicio1*Ejercicio1) determinarSiTerminaEn(numero int)bool{
	return ejercicio1.sacarUltimoNumeroGuardarResultado() ==  numero; 
}
func (ejercicio1*Ejercicio1) determinarSiTerminaEn4()bool{
	return ejercicio1.determinarSiTerminaEn(4); 
}
func (ejercicio*Ejercicio1)  mostrarRespuestaUsuario()string{
	var respuesta string;
	ejercicio.numeroresultado = ejercicio.sacarUltimoNumero();
	if ejercicio.determinarSiTerminaEn4() {
		respuesta = "El numero si termina en 4 ";	
	}else{
		respuesta = "El numero no termina en 4 ";	
	}
	return respuesta;
}
func (ejercicio1*Ejercicio1) ingresarValoresMostrarResultados()(string, int)  {
	ejercicio1.valorEntero = agregarMostrarValores.AgregarValoreVariableEntera();
	return ejercicio1.mostrarRespuestaUsuario(),ejercicio1.numeroresultado;
}
func  ejecutarEjercicio1(){
	ejercicio1 := Ejercicio1{};
	ejecutarEjercicio("1. Leer un número entero y determinar si es un número terminado en 4.",ejercicio1.ingresarValoresMostrarResultados)
}

/*
2. Leer un número entero y determinar si tiene 3 dígitos.
*/
type Ejercicio2 struct{
	valorEntero int
	numeroresultado int
}

func (ejercicio2*Ejercicio2) ContarDigitos(numeroEntero int) int{
	var particionado = numeroEntero;
	contador := 1;
	for particionado > 9 {
		particionado = particionado / 10 ;
		// fmt.Println("particionado ",particionado)
		contador++;
	}
	return contador;
}
func (ejercicio2*Ejercicio2)contarDigitosLog(numeroEntero int) int{
	digitos := int(math.Log10(float64(numeroEntero)))+1;
	// fmt.Println("digitos log ",math.Log10(float64(numeroEntero)));
	return digitos;
}
func (ejercicio2*Ejercicio2) contarDigitosGuardarResultado() int {
	ejercicio2.numeroresultado = ejercicio2.ContarDigitos(ejercicio2.valorEntero);
	return ejercicio2.numeroresultado;
}
func (ejercicio2*Ejercicio2) saberSiTieneEstaCantidadDigitos(digitos int) bool{	
	var coincidenciaDigitos bool;
	if digitos ==	ejercicio2.contarDigitosGuardarResultado() {
		coincidenciaDigitos = true;
	}
	return coincidenciaDigitos;
}
func (ejercicio2*Ejercicio2) saberSiTiene3Digitos()bool {
	return ejercicio2.saberSiTieneEstaCantidadDigitos(3);
}

func (ejercicio2*Ejercicio2) mostrarRespuestaUsuario() string{
	var respuesta string;	
	if ejercicio2.saberSiTiene3Digitos() {
		respuesta = "El numero si tiene los 3 digitos exactos";
	}else {		
		respuesta = "El numero no tiene los 3 digitos";
	}
	return respuesta;	
}
func (ejercicio2*Ejercicio2) ingresarValoresMostrarResultados()(string, int)  {
	ejercicio2.valorEntero = agregarMostrarValores.AgregarValoreVariableEntera();
	return ejercicio2.mostrarRespuestaUsuario(),ejercicio2.numeroresultado;
}
func  ejecutarEjercicio2(){
	ejercicio2 := Ejercicio2{}
	ejecutarEjercicio("2. Leer un número entero y determinar si tiene 3 dígitos.",ejercicio2.ingresarValoresMostrarResultados);
	
}

/*
3. Leer un número entero y determinar si es negativo.
*/
type Ejercicio3 struct{
	valorEntero int
	numeroresultado int
}
func (ejercicio3*Ejercicio3) determinarSiNegativo(numero int ) bool{
	return numero < 0;
}
func (ejercicio3*Ejercicio3)guardarResultadoDeterminarSiNegativo() bool {
	ejercicio3.numeroresultado = ejercicio3.valorEntero;
	return ejercicio3.determinarSiNegativo(ejercicio3.valorEntero);
} 
func (ejercicio3*Ejercicio3) mostrarRespuestaUsuario() string{
	var respuesta string;	
	if ejercicio3.guardarResultadoDeterminarSiNegativo() {
		respuesta = "El numero si es negativo";
	}else {		
		respuesta = "El numero no no es negativo";
	}
	return respuesta;	
}
func (ejercicio3*Ejercicio3) ingresarValoresMostrarResultados()(string, int)  {
	ejercicio3.valorEntero = agregarMostrarValores.AgregarValoreVariableEntera();
	return ejercicio3.mostrarRespuestaUsuario(),ejercicio3.numeroresultado;
}
func  ejecutarEjercicio3(){
	ejercicio3 := Ejercicio3{}
	ejecutarEjercicio("3. Leer un número entero y determinar si es negativo.",ejercicio3.ingresarValoresMostrarResultados);	
}
/*
4. Leer un número entero de dos dígitos y determinar a cuánto es igual la suma de sus dígitos.
*/
type Ejercicio4 struct{
	valorEntero int
	numeroresultado int
	ejercicio2 Ejercicio2
}
func (ejercicio4*Ejercicio4) sum(numberOne int, numberTwo int)int {
	return numberOne + numberTwo;
}
func (ejercicio4*Ejercicio4) mapTransformData(numeros []int, function func(int)int) []int {
	var arrayConvertido []int;
	for numero := 0;numero < len(numeros);numero++ {
		arrayConvertido = append(arrayConvertido,function(numeros[numero]));		 	
	}
	return arrayConvertido;
}
func (ejercicio4*Ejercicio4) determinarMinimoDigitos(cantidadDigitos int) bool {
	return ejercicio4.ejercicio2.ContarDigitos(ejercicio4.valorEntero) == cantidadDigitos;
}
func (ejercicio4*Ejercicio4) ArrayDigitosNumero(valorEntero int) []int{
	var particionado,cantidadDigitos int =  valorEntero, ejercicio4.ejercicio2.ContarDigitos(valorEntero);	 	
	// fmt.Println("arrayDigitosNumero entrada ",valorEntero);		
	var colecionDigitos = make([]int, cantidadDigitos);
	
	for contador := cantidadDigitos -1 ; contador >=0 ;  contador--{
		sobrante := particionado%10;
		colecionDigitos[contador] = sobrante;		
		particionado = particionado / 10 ;
	}
	// fmt.Println("arraydigitosNumero ",colecionDigitos);			
	return colecionDigitos;
}
func (ejercicio4*Ejercicio4)reduceInt(numeros []int, accion func(int,int) int) int{
	var resultado  int;
	for numero := 0; numero < len(numeros); numero++ {
		resultado = accion(resultado,numeros[numero]);
	}
	return resultado;
}
func (ejercicio4*Ejercicio4)reduce(numeros []int, accion func(int) bool) []int{
	var resultadoArray []int;	
	for numero := 0; numero < len(numeros); numero++ {
		if accion(numeros[numero]){
			resultadoArray = append(resultadoArray,numeros[numero]);
		}
	}
	return resultadoArray;
}
func (ejercicio4*Ejercicio4) sumarDigitosArray(numero int) int {
	coleccionDigitos := ejercicio4.ArrayDigitosNumero(numero);
	var sumaDigitos int;
	for cuenta := 0; cuenta < len(coleccionDigitos); cuenta++ {
		sumaDigitos += coleccionDigitos[cuenta] ;
	}
	return sumaDigitos;
}
func (ejercicio4*Ejercicio4) sumarDigitos(numero int) int {
	var sumaDigitos ,particionado,cantidadDigitos  = 0, numero, ejercicio4.ejercicio2.ContarDigitos(numero);	 	
	for contador := 0 ; contador < cantidadDigitos;  contador++{
		sobrante := particionado % 10
		sumaDigitos	+= sobrante;
		particionado = particionado / 10 ;
		// fmt.Println("particionado ",particionado,"sobrante",sobrante);
	}
	return sumaDigitos;
}
func (ejercicio4*Ejercicio4)guardarResultadoSumarDigitos() int {
	ejercicio4.numeroresultado = ejercicio4.sumarDigitosArray(ejercicio4.valorEntero);
	return ejercicio4.numeroresultado;
} 
func (ejercicio4*Ejercicio4) mostrarRespuestaUsuario() string{
	var respuesta string;	
	ejercicio4.guardarResultadoSumarDigitos();
	respuesta  = "la suma del los numeros es";
	return respuesta;	
}
func (ejercicio4*Ejercicio4) ingresarValoresMostrarResultados()(string, int)  {
	ejercicio4.valorEntero = agregarMostrarValores.AgregarValoreVariableEnteraDefinirCantidadDigito(2);
	return ejercicio4.mostrarRespuestaUsuario(),ejercicio4.numeroresultado;
}
func  ejecutarEjercicio4(){
	ejercicio4 := Ejercicio4{}
	ejecutarEjercicio("4. Leer un número entero de dos dígitos y determinar a cuánto es igual la suma de sus dígitos.",ejercicio4.ingresarValoresMostrarResultados);	
}
/*
5. Leer un número entero de dos dígitos y determinar si ambos dígitos son pares.
*/
type Ejercicio5 struct{
	valorEntero int
	numeroresultado int
	ejercicio2 Ejercicio2
	ejercicio4 Ejercicio4
}
func (ejercicio5*Ejercicio5) determinarSiPar(numero int)bool{
	var numeroPar bool;
	if numero % 2 == 0 {
		numeroPar = true;
	}
	return numeroPar;
}
func (ejercicio5*Ejercicio5) determinarSiDigitosPares()bool{
	digitosPar := true; 
	coleccionDigitos := ejercicio5.ejercicio4.ArrayDigitosNumero(ejercicio5.valorEntero);
	for digito := 0; digito < ejercicio5.ejercicio2.ContarDigitos(ejercicio5.valorEntero); digito++ {
		fmt.Println("digito saber si par ",coleccionDigitos[digito]);				
		if !ejercicio5.determinarSiPar( coleccionDigitos[digito]){
			digitosPar = false;
			break;
		}
	}
	return digitosPar;
}
func (ejercicio5*Ejercicio5) determinarSiDigitosParesGardarResultado()bool{
	 digitosPares := ejercicio5.determinarSiDigitosPares();
	if digitosPares {
		ejercicio5.numeroresultado = ejercicio5.valorEntero;
	}
	return digitosPares;
}	
func (ejercicio5*Ejercicio5) mostrarRespuestaUsuario() string{
	var respuesta string;	
	if ejercicio5.determinarSiDigitosParesGardarResultado(){
		respuesta  = "El numero tiene todos sus digitos pares";
	}else{
		respuesta  = "El numero no tiene todos sus digitos pares";
	}
	return respuesta;	
}
func (ejercicio5*Ejercicio5) ingresarValoresMostrarResultados()(string, int)  {
	ejercicio5.valorEntero = agregarMostrarValores.AgregarValoreVariableEnteraDefinirCantidadDigito(2);
	return ejercicio5.mostrarRespuestaUsuario(),ejercicio5.numeroresultado;
}
func  ejecutarEjercicio5(){
	ejercicio5 := Ejercicio5{};
	ejecutarEjercicio("5. Leer un número entero de dos dígitos y determinar si ambos dígitos son pares.",ejercicio5.ingresarValoresMostrarResultados);	
}

/*
6. Leer un número entero de dos dígitos menor que 20 y determinar si es primo.
*/
type Ejercicio6 struct{
	valorEntero int
	numeroresultado int
	ejercicio2 Ejercicio2
	ejercicio4 Ejercicio4
}
func (ejercicio6*Ejercicio6) DeterminarSiPrimoInt(numero int )bool{
	return ejercicio6.determinarSiPrimo(float64(numero));
}	
func (ejercicio6*Ejercicio6) determinarSiPrimo(numero float64 )bool{
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
func (ejercicio6*Ejercicio6) determinarSiPrimoGuardarRespuesta()bool{
	numeroFloat := float64(ejercicio6.valorEntero);
	ejercicio6.numeroresultado = ejercicio6.valorEntero;
	 return ejercicio6.determinarSiPrimo(numeroFloat);
}
func (ejercicio6*Ejercicio6) mostrarRespuestaUsuario() string {
	var respuesta string;
	if ejercicio6.determinarSiPrimoGuardarRespuesta() {
		respuesta = "El numero es primo";
	}else{
		respuesta = "El numero no es primo";		
	}
	return respuesta;
}
func (ejercicio6*Ejercicio6) ingresarValoresMostrarResultados()(string, int)  {
	ejercicio6.valorEntero = agregarMostrarValores.AgregarValoreVariableEnteraDefinirCantidadDigito(2);
	return ejercicio6.mostrarRespuestaUsuario(),ejercicio6.numeroresultado;
}
func  ejecutarEjercicio6(){
	ejercicio6 := Ejercicio6{};
	ejecutarEjercicio("6. Leer un número entero de dos dígitos menor que 20 y determinar si es primo.",ejercicio6.ingresarValoresMostrarResultados);	
}

/*
7. Leer un número entero de dos dígitos y determinar si es primo y además si es negativo.
*/
type Ejercicio7 struct{
	valorEntero int
	numeroresultado int
	ejercicio6 Ejercicio6
	ejercicio3 Ejercicio3
}

func (ejercicio7*Ejercicio7)determinarSiPrimoAdemasNegativoGuardarRespuesta()string{
	ejercicio7.numeroresultado = ejercicio7.valorEntero;
	ejercicio7.ejercicio6.valorEntero = ejercicio7.valorEntero;
	ejercicio7.ejercicio3.valorEntero = ejercicio7.valorEntero;	
	return ejercicio7.ejercicio6.mostrarRespuestaUsuario()+" y "+ ejercicio7.ejercicio3.mostrarRespuestaUsuario();
	
}
func (ejercicio7*Ejercicio7) mostrarRespuestaUsuario()string{
	return	ejercicio7.determinarSiPrimoAdemasNegativoGuardarRespuesta();
}
func (ejercicio7*Ejercicio7) ingresarValoresMostrarResultados()(string, int)  {
	ejercicio7.valorEntero = agregarMostrarValores.AgregarValoreVariableEnteraDefinirCantidadDigito(2);
	return ejercicio7.mostrarRespuestaUsuario(),ejercicio7.numeroresultado;
}
func  ejecutarEjercicio7(){
	ejercicio7 := Ejercicio7{};
	ejecutarEjercicio("7. Leer un número entero de dos dígitos y determinar si es primo y además si es negativo.",ejercicio7.ingresarValoresMostrarResultados);	
}

/*
8. Leer un número entero de dos dígitos y determinar si sus dos dígitos son primos.
*/
type Ejercicio8 struct{
	valorEntero int
	numeroresultado int
	ejercicio6 Ejercicio6
	ejercicio2 Ejercicio2
	ejercicio4 Ejercicio4

}
func(ejercicio8*Ejercicio8)siDigitosSonPrimos(numero int) bool{
	digitosPrimos := true;
	coleccionDigitos := ejercicio8.ejercicio4.ArrayDigitosNumero(ejercicio8.valorEntero);
	longitud := ejercicio8.ejercicio2.ContarDigitos(numero);
	for digito :=0; digito < longitud; digito++{
		fmt.Println("digito",coleccionDigitos[digito]);		
		if !ejercicio8.ejercicio6.determinarSiPrimo( float64( coleccionDigitos[digito]))   {
			digitosPrimos = false;
			break;
		}
	}
	return digitosPrimos;
}
func(ejercicio8*Ejercicio8)siDigitosSonPrimosGuardarResultado()bool{
	ejercicio8.numeroresultado = ejercicio8.valorEntero;
	return ejercicio8.siDigitosSonPrimos(ejercicio8.valorEntero);
}
func (ejercicio8*Ejercicio8) mostrarRespuestaUsuario()string{
	var resultado string;
	if	ejercicio8.siDigitosSonPrimosGuardarResultado() {
		resultado = "Todos sus digitos son primos";
	}else{
		resultado = "No todos sus digitos son primos";
	}
	return resultado;
}
func (ejercicio8*Ejercicio8) ingresarValoresMostrarResultados()(string, int)  {
	ejercicio8.valorEntero = agregarMostrarValores.AgregarValoreVariableEnteraDefinirCantidadDigito(2);
	return ejercicio8.mostrarRespuestaUsuario(),ejercicio8.numeroresultado;
}
func  ejecutarEjercicio8(){
	ejercicio8 := Ejercicio8{};
	ejecutarEjercicio("8. Leer un número entero de dos dígitos y determinar si sus dos dígitos son primos.",ejercicio8.ingresarValoresMostrarResultados);	
}
/*
9. Leer un número entero de dos dígitos y determinar si un dígito es múltiplo del otro.
*/
type Ejercicio9 struct{
	valorEntero int
	numeroresultado int
	ejercicio4 Ejercicio4
}
func (ejercicio9*Ejercicio9) saberSiEsteNumeroEsMultiploDeEste(dividendo , divisor int)bool{
	return dividendo % divisor == 0 ;
}
func(ejercicio9*Ejercicio9) determinarSiSusDosDigitosSonMultiplos(numero int) (int, bool){
	coleccionDigitos := ejercicio9.ejercicio4.ArrayDigitosNumero(ejercicio9.valorEntero);
	var numerorespuesta int;
	respuesta1 := ejercicio9.saberSiEsteNumeroEsMultiploDeEste(coleccionDigitos[0],coleccionDigitos[1]);
	respuesta2 := ejercicio9.saberSiEsteNumeroEsMultiploDeEste(coleccionDigitos[1],coleccionDigitos[0]);
	if respuesta1 && respuesta2 {
		numerorespuesta = ejercicio9.valorEntero;
	}else if respuesta1 {
		numerorespuesta = coleccionDigitos[0];
	} else if respuesta2 {
		numerorespuesta = coleccionDigitos[1];
	}
	return numerorespuesta , respuesta1 || respuesta2;
}
func(ejercicio9*Ejercicio9) ejecutarEjercicioGuardarResultado() bool{
	numero, esMultiplo := ejercicio9.determinarSiSusDosDigitosSonMultiplos(ejercicio9.valorEntero)
	if esMultiplo {
		ejercicio9.numeroresultado = numero;		
	}else {
		ejercicio9.numeroresultado = numero;
	}
	return esMultiplo;
}
func (ejercicio9*Ejercicio9) mostrarRespuestaUsuario()string{
	var resultado string;
	if	ejercicio9.ejecutarEjercicioGuardarResultado() {
		resultado = "uno de sus digitos son multiplos";
	}else{
		resultado = "ninguno de sus digitos son multiplos";
	}
	return resultado;
}
func (ejercicio9*Ejercicio9) ingresarValoresMostrarResultados()(string, int)  {
	ejercicio9.valorEntero = agregarMostrarValores.AgregarValoreVariableEnteraDefinirCantidadDigito(2);
	return ejercicio9.mostrarRespuestaUsuario(),ejercicio9.numeroresultado;
}
func  ejecutarEjercicio9(){
	ejercicio9 := Ejercicio9{};
	ejecutarEjercicio("9. Leer un número entero de dos dígitos y determinar si un dígito es múltiplo del otro.",ejercicio9.ingresarValoresMostrarResultados);	
}
/*
10. Leer un número entero de dos dígitos y determinar si los dos dígitos son iguales.
*/
type Ejercicio10 struct{
	valorEntero int
	numeroresultado int	
	ejercicio2 Ejercicio2
	ejercicio4 Ejercicio4
}
func (ejercicio10*Ejercicio10)determinarSiNumerosIguales(numero1 int , numero2 int)bool{
	// fmt.Println("compararion  ",numero1 , numero2,numero1 == numero2) ;		
	return numero1 == numero2;
}
func (ejercicio10*Ejercicio10)repetidoNumeroDentroDeEste(numeroParaBuscar int, buscarEnEste int) bool{
	coleccionDigitos := ejercicio10.ejercicio4.ArrayDigitosNumero(buscarEnEste);
	cantidadDigitos := ejercicio10.ejercicio2.ContarDigitos(buscarEnEste);
	var resultado bool;
	for digito := 0; digito < cantidadDigitos; digito++ {
		if ejercicio10.determinarSiNumerosIguales( numeroParaBuscar, coleccionDigitos[digito]){
			resultado = true;
			break;
		}
	}
	return resultado;	
}
func (ejercicio10*Ejercicio10)buscarColeccionarNumerosRepetidos(numeroParaBuscar int, buscarEnEste int) []int{
	var coleccionRepetidos []int;
	arrayNumero := ejercicio10.ejercicio4.ArrayDigitosNumero(buscarEnEste);
	cantidadDigitos := len(arrayNumero);
	for digito := 0; digito < cantidadDigitos; digito++ {
		// fmt.Println("buscarColeccionarNumerosRepetidos ",arrayNumero,digito,cantidadDigitos) ;	
		if ejercicio10.repetidoNumeroDentroDeEste( arrayNumero[digito],buscarEnEste) {
			coleccionRepetidos = append(coleccionRepetidos,arrayNumero[digito]);			
		}
	}
	return coleccionRepetidos;
}
func (ejercicio10*Ejercicio10)convertirArrayintToInt(arrayNumeros []int ) int{
	var coleccionNumeros []string;	
	for numero := 0; numero < len(arrayNumeros); numero++ {
		numeroConvertido := strconv.Itoa(arrayNumeros[numero]);
		// fmt.Println("dato array", numeroConvertido);
		coleccionNumeros = append(coleccionNumeros,numeroConvertido);
	}	
	respuesta,_ := strconv.Atoi(strings.Join(coleccionNumeros, ""));
	// fmt.Println("resultado conversion ",respuesta,coleccionNumeros) ;		
	return respuesta;
}
func (ejercicio10*Ejercicio10)quitarNumerosRepetidos( numero int ) []int{
	cantidadDigitos := ejercicio10.ejercicio2.ContarDigitos(numero);
	var coleccionNumeros []int;	
	arrayDigitos :=ejercicio10.ejercicio4.ArrayDigitosNumero(numero);
	// fmt.Println("colecvr cerso ",ejercicio10.convertirArrayintToInt(arrayDigitos)) ;	
	for digito := 0; digito < cantidadDigitos; digito++ { 	
		// fmt.Println("coleccion1 ",arrayDigitos[digito],numero) ;
		if ejercicio10.repetidoNumeroDentroDeEste( arrayDigitos[digito] , numero ){
			// fmt.Println("coleccion2 ",arrayDigitos[digito],ejercicio10.convertirArrayintToInt(coleccionNumeros)) ;	
			if !ejercicio10.repetidoNumeroDentroDeEste( arrayDigitos[digito] , ejercicio10.convertirArrayintToInt(coleccionNumeros) ) {
				coleccionNumeros = append(coleccionNumeros,arrayDigitos[digito]);
				// fmt.Println("coleccion3 ",ejercicio10.convertirArrayintToInt(coleccionNumeros)) ;		
			}
		}
	}
	return coleccionNumeros;
}
	
func (ejercicio10*Ejercicio10)determinarSiDigitosIguales(numero int) bool {
	var resultado bool;
	cantidadDigitos := ejercicio10.ejercicio2.ContarDigitos(numero);	
	// fmt.Println("cantidad digitos ",ejercicio10.quitarNumerosRepetidos(numero),cantidadDigitos) ;	
	if len(ejercicio10.quitarNumerosRepetidos(numero)) < cantidadDigitos {
		resultado = true;
	}
	return resultado;
}
func(ejercicio10*Ejercicio10) ejecutarEjercicioGuardarResultado() bool{
	ejercicio10.numeroresultado = ejercicio10.valorEntero;
	return  ejercicio10.determinarSiDigitosIguales(ejercicio10.valorEntero);
}
func(ejercicio10*Ejercicio10) mostrarRespuestaUsuario() string{
	var resultado string;
	if ejercicio10.ejecutarEjercicioGuardarResultado(){
		resultado = "Sus digitos son iguales";
	}else {
		resultado = "Sus digitos no son iguales";
	}
	return resultado ;
}
func (ejercicio10*Ejercicio10) ingresarValoresMostrarResultados()(string, int)  {
	ejercicio10.valorEntero = agregarMostrarValores.AgregarValoreVariableEnteraDefinirCantidadDigito(2);
	return ejercicio10.mostrarRespuestaUsuario(),ejercicio10.numeroresultado;
}
func  ejecutarEjercicio10(){
	ejercicio10 := Ejercicio10{};
	ejecutarEjercicio("10. Leer un número entero de dos dígitos y determinar si los dos dígitos son iguales.",ejercicio10.ingresarValoresMostrarResultados);	
}
/*
11. Leer dos números enteros y determinar cuál es el mayor.
*/
type Ejercicio11 struct{
	valoresEnteros []int
	numeroresultado int	
	numberMin int
}
func(ejercicio11*Ejercicio11)primerNumeroEsMayorQueSegundoNumero(numero1 int, numero2 int)bool{
	return numero1 > numero2;
}
func(ejercicio11*Ejercicio11)leerNumerosEnterosDeterminarMayor(numeros []int) int {
	numeroMayor := numeros[0];
	for numero := 0; numero <	len( numeros);numero++ {
		if(ejercicio11.primerNumeroEsMayorQueSegundoNumero(numeros[numero],numeroMayor)){
			numeroMayor =  numeros[numero];
			// fmt.Println("numeroMayor1",numeroMayor)			
		}
	}
	return numeroMayor;
}
func(ejercicio11*Ejercicio11)primerNumeroEsMenorQueSegundoNumero(numero1 int, numero2 int)bool{
	return numero1 < numero2;
}
func(ejercicio11*Ejercicio11)leerNumerosEnterosDeterminarMenor(numeros []int) int {
	numeroMenor := numeros[0];
	for numero := 0; numero <	len( numeros);numero++ {
		if(ejercicio11.primerNumeroEsMenorQueSegundoNumero(numeros[numero],numeroMenor)){
			numeroMenor =  numeros[numero];
			// fmt.Println("numeroMayor1",numeroMayor)			
		}
	}
	return numeroMenor;
}
func(ejercicio11*Ejercicio11) ejecutarEjercicioGuardarResultado() bool{
	ejercicio11.numeroresultado = ejercicio11.leerNumerosEnterosDeterminarMayor(ejercicio11.valoresEnteros);
	return true;	
}
func(ejercicio11*Ejercicio11) mostrarRespuestaUsuario() string{
	ejercicio11.ejecutarEjercicioGuardarResultado();
	return "El numero mayor es " ;
}
func (ejercicio11*Ejercicio11) ingresarValoresMostrarResultados()(string, int)  {
	ejercicio11.valoresEnteros = agregarMostrarValores.DeterminarCuantosDatosIngresar(2, agregarMostrarValores.IngresarDosDigitos);
	return ejercicio11.mostrarRespuestaUsuario(),ejercicio11.numeroresultado;
}
func  ejecutarEjercicio11(){
	ejercicio11 := Ejercicio11{};
	ejecutarEjercicio("11. Leer dos números enteros y determinar cuál es el mayor.",ejercicio11.ingresarValoresMostrarResultados);	
}
/*
12. Leer dos números enteros de dos dígitos y determinar si tienen dígitos comunes.*/
type Ejercicio12 struct {
	ejercicio11 Ejercicio11 
	ejercicio4 Ejercicio4
	ejercicio2 Ejercicio2
}
func(ejercicio12*Ejercicio12) seEncuentraNumeroDentroArray( numero int, array []int)bool{
	var seEncuentraDentro bool
	for numeros :=0; numeros < len(array);numeros++ {
	 	if	array[numeros] == numero {
			seEncuentraDentro = true;
		 }
	}
	return seEncuentraDentro;
}
func(ejercicio12*Ejercicio12) retornarEntreDosNumerosElMayoryElMenor( numero1, numero2 int) (int,int){
	var numeroMayor, numeroMenor int;
	if ejercicio12.ejercicio11.primerNumeroEsMayorQueSegundoNumero(numero1, numero2){
		numeroMayor = numero1;
		numeroMenor = numero2;
	}else {
		numeroMayor = numero2;
		numeroMenor = numero1;	
	}
	return numeroMayor, numeroMenor;
}
func(ejercicio12*Ejercicio12) digitosComunesEntreEstosDosNumeros( numero1, numero2 int) []int{
	var cuentaColeccion int;
	mayor, menor := ejercicio12.retornarEntreDosNumerosElMayoryElMenor(numero1, numero2);
	digitosMayor,digitosMenor := ejercicio12.ejercicio4.ArrayDigitosNumero(mayor),ejercicio12.ejercicio4.ArrayDigitosNumero(menor);
	var coleccionDigitosIguales = make([]int, len(digitosMayor));
	for digitoMayor := 0; digitoMayor < len(digitosMayor); digitoMayor++ {
		if	ejercicio12.seEncuentraNumeroDentroArray( digitosMayor[digitoMayor],digitosMenor) {
			if !ejercicio12.seEncuentraNumeroDentroArray(digitosMayor[digitoMayor],coleccionDigitosIguales){
				coleccionDigitosIguales[cuentaColeccion] = digitosMayor[digitoMayor];
				cuentaColeccion++
			}
		}
	}
	return coleccionDigitosIguales[0:cuentaColeccion];
}
func (ejercicio12*Ejercicio12) leerNumerosEnterosDeterminarDigitosIguales( valoresEnteros []int) []int {
	var coleccionDigitosComunes []int;
	for numero :=0; numero < len(valoresEnteros)-1; numero++{
		coleccionDigitosComunes = ejercicio12.digitosComunesEntreEstosDosNumeros(	valoresEnteros[numero],	valoresEnteros[numero+1]);
		fmt.Println(coleccionDigitosComunes)	;
	}
	return coleccionDigitosComunes;
}
func (ejercicio12*Ejercicio12) ejecutarEjercicioDeterminarDigitosIguales(valoresEnteros []int)( bool,int){
	resultado := ejercicio12.leerNumerosEnterosDeterminarDigitosIguales(valoresEnteros);
	var numeroResultado int;
	for numero := 0 ; numero < len( resultado); numero++ {
		numeroResultado = numeroResultado*10 +  resultado[numero];
	}	 
	return len(resultado) > 0, numeroResultado;
}
func  ejecutarEjercicio12(){
	ejercicio12 := Ejercicio12{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarDosDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio12.ejecutarEjercicioDeterminarDigitosIguales;
	guardarResultadosMostrar.respuestaNegativa = "no tiene digitos iguales";
	guardarResultadosMostrar.respuestaPositiva = "si tiene digitos iguales";
	guardarResultadosMostrar.cantidadDatosIngresar = 2;
	ejecutarEjercicio("12. Leer dos números enteros de dos dígitos y determinar si tienen dígitos comunes.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
13. Leer dos números enteros de dos dígitos y determinar si la suma de los dos números origina
un número par.*/
type Ejercicio13 struct {
	ejercicio4 Ejercicio4
	ejercicio5 Ejercicio5
}
func(ejercicio13*Ejercicio13)sumaDosNumeros(numero1 , numero2 int )int{
	return numero1 + numero2;
}
func(ejercicio13*Ejercicio13)sumarArrayDeNumeros(numeros []int)int{
	sumarNumeros := arrayMetodos.ReduceInt(numeros,ejercicio13.sumaDosNumeros);
	return sumarNumeros;
}
func(ejercicio13*Ejercicio13)numerosSumaGeneraPar(numeros []int)(bool,int){
	numerosSumados :=  ejercicio13.sumarArrayDeNumeros(numeros);
	return ejercicio13.ejercicio5.determinarSiPar(numerosSumados),numerosSumados;
}

func  ejecutarEjercicio13(){
	ejercicio13 := Ejercicio13{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar = 2;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarDosDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio13.numerosSumaGeneraPar;
	guardarResultadosMostrar.respuestaPositiva = "La suma de sus numeros si da un numero par";
	guardarResultadosMostrar.respuestaNegativa = "La suma de sus numeros no da un numero par";
	ejecutarEjercicio("13. Leer dos números enteros de dos dígitos y determinar si la suma de los dos números origina un número par.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*14. Leer dos números enteros de dos dígitos y determinar a cuánto es igual la suma de todos los
dígitos.*/
type Ejercicio14 struct {
	ejercicio4 Ejercicio4
	ejercicio13 Ejercicio13
}
func (ejercicio14*Ejercicio14)coleccionSumasVariosDigitos(numeros []int)[]int{
	var coleccionSumaDigitos = make([]int,  len(numeros))
	for numero :=0; numero < len(numeros);numero++ {
		coleccionSumaDigitos[numero] = ejercicio14.ejercicio4.sumarDigitosArray(numeros[numero]);
	}
	return coleccionSumaDigitos;
}
func (ejercicio14*Ejercicio14)sumarDigitosVariosNumeros(numeros []int)int{
	 coleccionSumaDigitos := ejercicio14.coleccionSumasVariosDigitos(numeros);
	 return ejercicio14.ejercicio13.sumarArrayDeNumeros(coleccionSumaDigitos);
}
func (ejercicio14*Ejercicio14) ejecutarEjercicio(numeros []int) (bool, int){
	return true,  ejercicio14.sumarDigitosVariosNumeros(numeros);
}

func  ejecutarEjercicio14(){
	ejercicio14 := Ejercicio14{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar = 2;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarDosDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio14.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "La sumas de sus digitos es";
	guardarResultadosMostrar.respuestaNegativa = "La sumas de sus digitos es";
	ejecutarEjercicio("14. Leer dos números enteros de dos dígitos y determinar a cuánto es igual la suma de todos los 	dígitos.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
15. Leer un número entero de tres dígitos y determinar a cuánto es igual la suma de sus dígitos.
*/
type Ejercicio15 struct {
	ejercicio4 Ejercicio4
}
func (ejercicio15*Ejercicio15) leerNumeroEnteroCuantoSumaDigitos(numeros []int)(bool, int){
	numero := numeros[0];
	return true, ejercicio15.ejercicio4.sumarDigitos(numero);
}
func  ejecutarEjercicio15(){
	ejercicio15 := Ejercicio15{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarTresDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio15.leerNumeroEnteroCuantoSumaDigitos;
	guardarResultadosMostrar.respuestaPositiva = "La suma de sus digitos que tiene es";
	guardarResultadosMostrar.respuestaNegativa = "La suma de sus digitos que tiene es";
	ejecutarEjercicio("15. Leer un número entero de tres dígitos y determinar a cuánto es igual la suma de sus dígitos.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}

/*
16. Leer un número entero de tres dígitos y determinar si al menos dos de sus tres dígitos son
iguales.
*/

type Ejercicio16 struct {
	ejercicio10 Ejercicio10
}

func (ejercicio16*Ejercicio16) ejecutarEjercicio(numeros []int)(bool, int){
	numerosSinRepetidos := ejercicio16.ejercicio10.convertirArrayintToInt(numeros);
	return  ejercicio16.ejercicio10.determinarSiDigitosIguales(numeros[0]),numerosSinRepetidos;
}
func  ejecutarEjercicio16(){
	ejercicio16 := Ejercicio16{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarTresDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio16.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El numero si posee numeros repetidos";
	guardarResultadosMostrar.respuestaNegativa = "El numero no posee numeros repetidos";
	ejecutarEjercicio("16. Leer un número entero de tres dígitos y determinar si al menos dos de sus tres dígitos son iguales.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
17. Leer un número entero de tres dígitos y determinar en qué posición está el mayor dígito.
*/
type Ejercicio17 struct {
	ejercicio4 Ejercicio4
	ejercicio11 Ejercicio11
	numeroComparar int
}

func (ejercicio17*Ejercicio17) sonDosNumerosParecidos(numeroEntrante int) bool{
	return numeroEntrante  == ejercicio17.numeroComparar;
}
func (ejercicio17*Ejercicio17) SearchForPositionNumberGreater(numeros []int)int{
	ejercicio17.numeroComparar = ejercicio17.ejercicio11.leerNumerosEnterosDeterminarMayor(numeros);			
	return arrayMetodos.FindIndex(numeros,ejercicio17.sonDosNumerosParecidos);
}
func (ejercicio17*Ejercicio17) encontrarPositionNumeroMayor(numeros int)int{
	arrayNumeros :=  ejercicio17.ejercicio4.ArrayDigitosNumero(numeros);
	ejercicio17.numeroComparar = ejercicio17.ejercicio11.leerNumerosEnterosDeterminarMayor(arrayNumeros);
	return arrayMetodos.FindIndex(arrayNumeros,ejercicio17.sonDosNumerosParecidos);
}

func (ejercicio17*Ejercicio17) ejecutarEjercicio(numeros []int)(bool, int){
	return  true,ejercicio17.encontrarPositionNumeroMayor(numeros[0]);
}

func  ejecutarEjercicio17(){
	ejercicio17 := Ejercicio17{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarTresDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio17.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "La posision del numero mayor es";
	guardarResultadosMostrar.respuestaNegativa = "";
	ejecutarEjercicio("17. Leer un número entero de tres dígitos y determinar en qué posición está el mayor dígito.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
18. Leer un número entero de tres dígitos y determinar si algún dígito es múltiplo de los otros.
*/
type Ejercicio18 struct {
	ejercicio4 Ejercicio4
	ejercicio9 Ejercicio9
}

func (ejercicio18*Ejercicio18)buscarPorCriterio(numeros []int, function func(int,int)bool)[]int{
	//coleccionRepetidos = append(coleccionRepetidos,arrayNumero[digito]);			
	var resultadoArray 	[]int;
	for posicion :=0; posicion <  len(numeros); posicion++ {
		for posisionComparado :=0; posisionComparado  <  len(numeros); posisionComparado ++ {
			if posicion != posisionComparado {
				if function( numeros[posicion],numeros[posisionComparado] ){
					resultadoArray = append(resultadoArray,numeros[posicion]);
				}
			}
	   }		 
	}
	return resultadoArray;	
}

func (ejercicio18*Ejercicio18)determinarNumerosMultiploOtros(numeros []int)[]int{
	coleccionNumeros :=  arrayMetodos.CompareAllTheArray(numeros,ejercicio18.ejercicio9.saberSiEsteNumeroEsMultiploDeEste);
	return coleccionNumeros;
}
func (ejercicio18*Ejercicio18)determinarSiNumeroMultiploOtros(numero int)bool{
	arrayNumeros :=  ejercicio18.ejercicio4.ArrayDigitosNumero(numero);	
	coleccionNumeros := ejercicio18.determinarNumerosMultiploOtros(arrayNumeros);
	var respuesta bool;
	if len(coleccionNumeros) >0 {
		respuesta = true;
	}
	fmt.Println("Numeros multiplo de otros ", coleccionNumeros);
	return respuesta;
}
func (ejercicio18*Ejercicio18) ejecutarEjercicio(numeros []int)(bool, int){	
	return  ejercicio18.determinarSiNumeroMultiploOtros(numeros[0]),0;
}

func  ejecutarEjercicio18(){
	ejercicio18 := Ejercicio18{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarTresDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio18.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Posee numeros multiplos";
	guardarResultadosMostrar.respuestaNegativa = "No posee numeros multiplos";
	ejecutarEjercicio("18. Leer un número entero de tres dígitos y determinar si algún dígito es múltiplo de los otros.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
19. Leer tres números enteros y determinar cuál es el mayor. Usar solamente dos variables.
*/
type Ejercicio19 struct {
	ejercicio11 Ejercicio11
}
func (ejercicio19*Ejercicio19) ejecutarEjercicio(numeros []int)(bool, int){	
	return  true,ejercicio19.ejercicio11.leerNumerosEnterosDeterminarMayor(numeros);
}
func  ejecutarEjercicio19(){
	ejercicio19 := Ejercicio19{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=3;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio19.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Numero mayor es";
	guardarResultadosMostrar.respuestaNegativa = "";
	ejecutarEjercicio("19. Leer tres números enteros y determinar cuál es el mayor. Usar solamente dos variables.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
20. Leer tres números enteros y mostrarlos ascendentemente.
*/
type Ejercicio20 struct {
	ejercicio11 Ejercicio11
	ejercicio4 Ejercicio4
	numeroMayor int
}
func(ejercicio20*Ejercicio20) quitarEsteNumeroMayor(numero int)bool{
	return	ejercicio20.numeroMayor != numero;
}
func(ejercicio20*Ejercicio20) quitarMayorNumeroArray(arrayNumeros []int) []int {
	return arrayMetodos.Filter(arrayNumeros,ejercicio20.quitarEsteNumeroMayor);
}
func(ejercicio20*Ejercicio20) organizarNumeros(arrayNumeros []int) []int {
	var numeroOrganizados []int;
	coleccionNumeros := arrayNumeros;
	for numero := 0; numero < len(arrayNumeros);numero++ {
		ejercicio20.numeroMayor = ejercicio20.ejercicio11.leerNumerosEnterosDeterminarMayor(coleccionNumeros);
		coleccionNumeros = ejercicio20.quitarMayorNumeroArray(coleccionNumeros);
		numeroOrganizados = append(numeroOrganizados,ejercicio20.numeroMayor);		
	}
	return numeroOrganizados; 
}

func (ejercicio20*Ejercicio20) ejecutarEjercicio(numeros []int)(bool, int){
	numeroOrganizado := ejercicio20.organizarNumeros(numeros)
	fmt.Println("Numeros organizados ",numeroOrganizado)	
	return  true,numeroOrganizado[0];
}
func  ejecutarEjercicio20(){
	ejercicio20 := Ejercicio20{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=2;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio20.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Numero mayor es";
	guardarResultadosMostrar.respuestaNegativa = "";
	ejecutarEjercicio("20. Leer tres números enteros y mostrarlos ascendentemente.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
21. Leer tres números enteros de dos dígitos cada uno y determinar en cuál de ellos se encuentra el mayor dígito.
*/ 
type Ejercicio21 struct{
	ejercicio11 Ejercicio11	
	ejercicio4 Ejercicio4
	ejercicio17 Ejercicio17
}
func(ejercicio21*Ejercicio21)sacarMayorDigito(numero int) int{
	return ejercicio21.ejercicio11.leerNumerosEnterosDeterminarMayor(ejercicio21.ejercicio4.ArrayDigitosNumero(numero));
}
func(ejercicio21*Ejercicio21)DeterminarPosisionMayorDigito(numeros []int)int{
	fmt.Println("no convertidos",numeros);	
	numerosMostradoDigitoMayor := arrayMetodos.MapTransformData(numeros,ejercicio21.sacarMayorDigito);
	fmt.Println("numeros convertidos",numerosMostradoDigitoMayor);
	positionNumberMax := ejercicio21.ejercicio17.SearchForPositionNumberGreater(numerosMostradoDigitoMayor);
	return 	positionNumberMax;			
}
func (ejercicio21*Ejercicio21) ejecutarEjercicio(numeros []int)(bool, int){
	posisionNumeroMayor := ejercicio21.DeterminarPosisionMayorDigito(numeros)
	fmt.Println("Numeros organizados ",posisionNumeroMayor)	
	return  true,posisionNumeroMayor;
}
func  ejecutarEjercicio21(){
	ejercicio21 := Ejercicio21{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=3;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarDosDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio21.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "posision numero mayor";
	guardarResultadosMostrar.respuestaNegativa = "";
	ejecutarEjercicio("21. Leer tres números enteros de dos dígitos cada uno y determinar en cuál de ellos se encuentra	el mayor dígito.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
22. Leer un número entero de tres dígitos y determinar si el primer dígito es igual al último.
*/
type Ejercicio22 struct{	
	ejercicio4 Ejercicio4
}
func (ejercicio22*Ejercicio22) numberFirstEqualToLast(numero int)bool{
	numeros := ejercicio22.ejercicio4.ArrayDigitosNumero(numero)
	return numeros[0] == numeros[len(numeros)-1];
}
func (ejercicio22*Ejercicio22) ejecutarEjercicio(numeros []int)(bool, int){
	fmt.Println("Numeros",numeros);
	return  ejercicio22.numberFirstEqualToLast(numeros[0]),numeros[0];
}
func  ejecutarEjercicio22(){
	ejercicio22 := Ejercicio22{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarTresDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio22.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El primer digito es igual al ultimo";
	guardarResultadosMostrar.respuestaNegativa = "El primer digito es diferetne al ultimo";
	ejecutarEjercicio("22. Leer un número entero de tres dígitos y determinar si el primer dígito es igual al último.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
23. Leer un número entero de tres dígitos y determinar cuántos dígitos primos tiene.
*/
type Ejercicio23 struct {
	ejercicio4 Ejercicio4	
	ejercicio6 Ejercicio6
}
func (ejercicio23*Ejercicio23) thisNumberIntIsPrime(number int) bool{
	return 	ejercicio23.ejercicio6.determinarSiPrimo(float64(number));	
}
func (ejercicio23*Ejercicio23) howManyPrimeCousinsDoItHave(numero int ) int{
	arrayDigitos := ejercicio23.ejercicio4.ArrayDigitosNumero(numero);
	numberPrimes := arrayMetodos.Filter(arrayDigitos,ejercicio23.thisNumberIntIsPrime);
	fmt.Println("numero primos encontrados ",numberPrimes) ;		
	return len(numberPrimes);
}
func (ejercicio23*Ejercicio23) ejecutarEjercicio(numeros []int)(bool, int){
	fmt.Println("Numeros",numeros);
	numberOfPrimeNumbers := ejercicio23.howManyPrimeCousinsDoItHave(numeros[0])
	return   numberOfPrimeNumbers > 0,numberOfPrimeNumbers;
}
func  ejecutarEjercicio23(){
	ejercicio23 := Ejercicio23{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarTresDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio23.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Si tiene digitos primos";
	guardarResultadosMostrar.respuestaNegativa = "No tiene digitos primos";
	ejecutarEjercicio("23. Leer un número entero de tres dígitos y determinar cuántos dígitos primos tiene.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
24. Leer un número entero de tres dígitos y determinar cuántos dígitos pares tiene.
*/
type Ejercicio24 struct {
	ejercicio4 Ejercicio4
	ejercicio5 Ejercicio5
}
func (ejercicio24*Ejercicio24) countDigitPairs(number int) int{
	arrayDigit := ejercicio24.ejercicio4.ArrayDigitosNumero(number);
	pairNumbers := arrayMetodos.Filter(arrayDigit,ejercicio24.ejercicio5.determinarSiPar);
	return len(pairNumbers);
}
func (ejercicio24*Ejercicio24) ejecutarEjercicio(numeros []int)(bool, int){
	fmt.Println("Numeros",numeros);
	numberOfPairsNumbers := ejercicio24.countDigitPairs(numeros[0]);
	return   numberOfPairsNumbers > 0,numberOfPairsNumbers;
}
func  ejecutarEjercicio24(){
	ejercicio24 := Ejercicio24{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarTresDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio24.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Si tiene digitos pares";
	guardarResultadosMostrar.respuestaNegativa = "No tiene digitos pares";
	ejecutarEjercicio("24. Leer un número entero de tres dígitos y determinar cuántos dígitos pares tiene.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
25. Leer un número entero de tres dígitos y determinar si alguno de sus dígitos es igual a la suma
de los otros dos.
*/
type Ejercicio25 struct {
	ejercicio4 Ejercicio4
	ejercicio13 Ejercicio13
	removeNumber int 
}
func (ejercicio25*Ejercicio25) removeThisNumber(number int) bool {
	return ejercicio25.removeNumber != number;
}
func (ejercicio25*Ejercicio25) sumNumberOtherNumber(number int) int{
	var numbersum, answerCorect int;
	arrayDigit := ejercicio25.ejercicio4.ArrayDigitosNumero(number);
	for removeNumber :=0;removeNumber < len(arrayDigit);removeNumber++{
		ejercicio25.removeNumber = arrayDigit[removeNumber];
		numbersRemove := arrayMetodos.Filter(arrayDigit,ejercicio25.removeThisNumber);
		// fmt.Println("con el numero removido",numbersRemove,ejercicio25.removeNumber);			
		numbersum = ejercicio25.ejercicio13.sumarArrayDeNumeros(numbersRemove);
		// fmt.Println("suma numeros",numbersum);	
		if numbersum == arrayDigit[removeNumber] {
			answerCorect = numbersum;
			break;
		}
	}
	return answerCorect;
}
func (ejercicio25*Ejercicio25) oneOfItsDigitsEqualToTheSumOfTheOtherTwo (number int)(bool,int){
	return ejercicio25.sumNumberOtherNumber(number) > 1,ejercicio25.sumNumberOtherNumber(number);
}
func (ejercicio25*Ejercicio25) ejecutarEjercicio(numeros []int)(bool, int){
	return   ejercicio25.oneOfItsDigitsEqualToTheSumOfTheOtherTwo(numeros[0]);
}
func  ejecutarEjercicio25(){
	ejercicio25 := Ejercicio25{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarTresDigitos
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio25.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Si tiene un numero que la suma de los otros es igual  a el";
	guardarResultadosMostrar.respuestaNegativa = "No tiene un numero que la suma de los otros es igual  a el";
	ejecutarEjercicio("25. Leer un número entero de tres dígitos y determinar si alguno de sus dígitos es igual a la suma de los otros dos.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
26. Leer un número entero de cuatro dígitos y determinar a cuanto es igual la suma de sus dígitos.
*/
type Ejercicio26 struct {
	ejercicio4 Ejercicio4
}

func (ejercicio26*Ejercicio26) ejecutarEjercicio(numeros []int)(bool, int){
	numeroSumados := ejercicio26.ejercicio4.sumarDigitosArray(numeros[0]);
	return  true ,numeroSumados;
}
func  ejecutarEjercicio26(){
	ejercicio26 := Ejercicio26{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarCuatroDigitos;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio26.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "La suma de sus numeros es ";
	guardarResultadosMostrar.respuestaNegativa = "";
	ejecutarEjercicio("26. Leer un número entero de cuatro dígitos y determinar a cuanto es igual la suma de sus dígitos.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
27. Leer un número entero de cuatro dígitos y determinar cuántos dígitos pares tiene.
*/

type Ejercicio27 struct {
	ejercicio4 Ejercicio4
	ejercicio5 Ejercicio5
}
func (ejercicio27*Ejercicio27) howManyEvenNumbersHave(number int) int {
	arrayDigitos := ejercicio27.ejercicio4.ArrayDigitosNumero(number);
	pairNUmber := arrayMetodos.Filter(arrayDigitos, ejercicio27.ejercicio5.determinarSiPar);
	fmt.Println(pairNUmber);	
	return len(pairNUmber);
}
func (ejercicio27*Ejercicio27) ejecutarEjercicio(numeros []int)(bool, int){
	countPairNumber := ejercicio27.howManyEvenNumbersHave(numeros[0]);
	return  countPairNumber > 0 ,countPairNumber;
}
func  ejecutarEjercicio27(){
	ejercicio27 := Ejercicio27{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarCuatroDigitos;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio27.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El numero tiene digitos pares ";
	guardarResultadosMostrar.respuestaNegativa = "El numero no tiene digitos pares";
	ejecutarEjercicio("27. Leer un número entero de cuatro dígitos y determinar cuántos dígitos pares tiene.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
28. Leer un número entero menor que 50 y positivo y determinar si es un número primo.
*/

type Ejercicio28 struct {
	ejercicio6 Ejercicio6
}


func (ejercicio28*Ejercicio28) ejecutarEjercicio(numeros []int)(bool, int){
	return  ejercicio28.ejercicio6.determinarSiPrimo(float64(numeros[0])) ,numeros[0];
}
func  ejecutarEjercicio28(){
	ejercicio28 := Ejercicio28{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.NumeroPositivoMenorDe50;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio28.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El numero es primo ";
	guardarResultadosMostrar.respuestaNegativa = "El numero no primo";
	ejecutarEjercicio("28. Leer un número entero menor que 50 y positivo y determinar si es un número primo.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
29. Leer un número entero de cinco dígitos y determinar si es un número capicúo. Ej. 15651,
59895.
*/
type Ejercicio29 struct {
	ejercicio4 Ejercicio4 
	multiTenInTen int
	dividiTenInTen int
}
func (ejercicio29*Ejercicio29)multiplyTenInTen(number int) int{
	multipliForTen := number * ejercicio29.multiTenInTen
	ejercicio29.multiTenInTen *= 10;
	return multipliForTen;
}	
func (ejercicio29*Ejercicio29)divisionTenInTen(number int) int{
	ejercicio29.dividiTenInTen /=  10;
	return number * ejercicio29.dividiTenInTen;
}
func (ejercicio29*Ejercicio29)knowIfNumberIsCapicúa(number int) bool{
	arrayNumber := ejercicio29.ejercicio4.ArrayDigitosNumero(number);
	ejercicio29.multiTenInTen = 1;
	converTenInTenOne := arrayMetodos.MapTransformData(arrayNumber,ejercicio29.multiplyTenInTen);
	sumOne := arrayMetodos.ReduceRightInt(converTenInTenOne,ejercicio29.ejercicio4.sum);
	ejercicio29.dividiTenInTen = ejercicio29.multiTenInTen;
	converTenInTenTwo := arrayMetodos.MapTransformData(arrayNumber,ejercicio29.divisionTenInTen);
	sumTwo := arrayMetodos.ReduceInt(converTenInTenTwo,ejercicio29.ejercicio4.sum);
	fmt.Println("Capicua", sumOne == sumTwo ,sumOne , sumTwo)
	
	return sumOne == sumTwo;
}
func (ejercicio29*Ejercicio29) ejecutarEjercicio(numeros []int)(bool, int){
	return  ejercicio29.knowIfNumberIsCapicúa(numeros[0]) ,numeros[0];
}
func  ejecutarEjercicio29(){
	ejercicio29 := Ejercicio29{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarCincoDigitos;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio29.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El numero es capicua ";
	guardarResultadosMostrar.respuestaNegativa = "El numero no capicua";
	ejecutarEjercicio("29. Leer un número entero de cinco dígitos y determinar si es un número capicúo. Ej. 15651,		59895.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
30. Leer un número entero de cuatro dígitos y determinar si el segundo dígito es igual al penúltimo.
*/
type Ejercicio30 struct{
	ejercicio4 Ejercicio4 	
}
func (ejercicio30*Ejercicio30)secondDigitEqualToTheLastLastLast(number int) bool {
	arrayDigitos := ejercicio30.ejercicio4.ArrayDigitosNumero(number);
	numberLastLastLast := arrayDigitos[len(arrayDigitos)-2];
	return arrayDigitos[1] == numberLastLastLast;
}
func (ejercicio30*Ejercicio30) ejecutarEjercicio(numeros []int)(bool, int){
	return  ejercicio30.secondDigitEqualToTheLastLastLast(numeros[0]) ,numeros[0];
}
func  ejecutarEjercicio30(){
	ejercicio30 := Ejercicio30{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarCuatroDigitos;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio30.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El segundo digito es igual al penultimo numero ";
	guardarResultadosMostrar.respuestaNegativa = "El segundo digito no es igual al penultimo numero ";
	ejecutarEjercicio("30. Leer un número entero de cuatro dígitos y determinar si el segundo dígito es igual al penúltimo.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
31. Leer un número entero y determina si es igual a 10.
*/
type Ejercicio31 struct{}

func (ejercicio31*Ejercicio31)numberEqualToTen(number int) bool {
	return number == 10;
}
func (ejercicio31*Ejercicio31) ejecutarEjercicio(numeros []int)(bool, int){
	return  ejercicio31.numberEqualToTen(numeros[0]) ,numeros[0];
}
func  ejecutarEjercicio31(){
	ejercicio31 := Ejercicio31{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio31.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El numero es igual a diez ";
	guardarResultadosMostrar.respuestaNegativa = "El numero no es igual a diez ";
	ejecutarEjercicio("31. Leer un número entero y determina si es igual a 10.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
32. Leer un número entero y determinar si es múltiplo de 7.
*/
type Ejercicio32 struct{}

func (ejercicio32*Ejercicio32)multipleNumberOfSeven(number int) bool {
	return number % 7 == 0;
}
func (ejercicio32*Ejercicio32) ejecutarEjercicio(numeros []int)(bool, int){
	return  ejercicio32.multipleNumberOfSeven(numeros[0]) ,numeros[0];
}
func  ejecutarEjercicio32(){
	ejercicio32 := Ejercicio32{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio32.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El numero es multiplo de 7 ";
	guardarResultadosMostrar.respuestaNegativa = "El numero no es multiplo de 7 ";
	ejecutarEjercicio("32. Leer un número entero y determinar si es múltiplo de 7.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
33. Leer un número entero y determinar si termina en 7.
*/
type Ejercicio33 struct{
	ejercicio4 Ejercicio4
}

func (ejercicio33*Ejercicio33)numberEndsInSeven(number int) bool {
	arrayDigitos := ejercicio33.ejercicio4.ArrayDigitosNumero(number);
	return arrayDigitos[len(arrayDigitos)-1] == 7;
}
func (ejercicio33*Ejercicio33) ejecutarEjercicio(numeros []int)(bool, int){
	return  ejercicio33.numberEndsInSeven(numeros[0]) ,numeros[0];
}
func  ejecutarEjercicio33(){
	ejercicio33 := Ejercicio33{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarCuatroDigitos;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio33.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El numero termina en 7 ";
	guardarResultadosMostrar.respuestaNegativa = "El numero no termina en 7 ";
	ejecutarEjercicio("33. Leer un número entero y determinar si termina en 7..",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
34. Leer un número entero menor que mil y determinar cuántos dígitos tiene.
*/
type Ejercicio34 struct{
	ejercicio2 Ejercicio2
}
func (ejercicio34*Ejercicio34) ejecutarEjercicio(numeros []int)(bool, int){
	return  true ,ejercicio34.ejercicio2.contarDigitosLog(numeros[0]);
}
func  ejecutarEjercicio34(){
	ejercicio34 := Ejercicio34{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio34.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El numero tiene esta cantidad de digitos";
	guardarResultadosMostrar.respuestaNegativa = "";
	ejecutarEjercicio("34. Leer un número entero menor que mil y determinar cuántos dígitos tiene.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
35. Leer un número entero de dos dígitos, guardar cada dígito en una variable diferente y luego
mostrarlas en pantalla.
*/
type Ejercicio35 struct{
	ejercicio4 Ejercicio4
	respuesta string
}

func (ejercicio35*Ejercicio35)showDigitsVariableDifferent(number int) string {
	arrayDigitos := ejercicio35.ejercicio4.ArrayDigitosNumero(number);
	numeroDigitoUnoEnTexto := strconv.Itoa(arrayDigitos[0]);
	numeroDigitoDosEnTexto := strconv.Itoa(arrayDigitos[1]);
	return  "el primer digito es "+numeroDigitoUnoEnTexto+" el segundo digito es "+numeroDigitoDosEnTexto;
}
func (ejercicio35*Ejercicio35) ejecutarEjercicio(numeros []int)(bool, int){
	return  true ,numeros[0];
}
func  ejecutarEjercicio35(){
	ejercicio35 := Ejercicio35{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio35.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "";
	guardarResultadosMostrar.respuestaNegativa = "";
	guardarResultadosMostrar.funcionMostrarRespuestaUsuario = ejercicio35.showDigitsVariableDifferent;
	ejecutarEjercicio("Leer un número entero de dos dígitos, guardar cada dígito en una variable diferente y luego mostrarlas en pantalla.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
36. Leer un número entero de 4 dígitos y determinar si tiene mas dígitos pares o impares.
*/
type Ejercicio36 struct{
	ejercicio24 Ejercicio24
	ejercicio4 Ejercicio4
}
func (ejercicio36*Ejercicio36) determinarSiImpar(numero int)bool{
	var numeroImpar bool;
	if numero % 2 != 0 {
		numeroImpar = true;
	}
	return numeroImpar;
}
func (ejercicio36*Ejercicio36) countDigitImPairs(number int) int{
	arrayDigit := ejercicio36.ejercicio4.ArrayDigitosNumero(number);
	impairNumbers := arrayMetodos.Filter(arrayDigit,ejercicio36.determinarSiImpar);
	return len(impairNumbers);
}
func (ejercicio36*Ejercicio36)hasMoreEvenDigitsThanOddDigits(number int) bool {
	numberPairs := ejercicio36.ejercicio24.countDigitPairs(number);
	numberImpairs := ejercicio36.countDigitImPairs(number);
	return numberPairs > numberImpairs;
}
func (ejercicio36*Ejercicio36) ejecutarEjercicio(numeros []int)(bool, int){
	return  ejercicio36.hasMoreEvenDigitsThanOddDigits(numeros[0]),numeros[0];
}
func  ejecutarEjercicio36(){
	ejercicio36 := Ejercicio36{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarCuatroDigitos;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio36.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Tiene mas digitos pares que impares";
	guardarResultadosMostrar.respuestaNegativa = "No tiene mas digitos pares que impares";
	ejecutarEjercicio("36. Leer un número entero de 4 dígitos y determinar si tiene mas dígitos pares o impares.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
37. Leer dos números enteros y determinar cuál es múltiplo de cuál.
*/
type Ejercicio37 struct{
	ejercicio9 Ejercicio9
	ejercicio25 Ejercicio25
}

func (ejercicio37*Ejercicio37) whichIsMultiploOfWhich(numbers []int)[]int{
	return 	arrayMetodos.CompareAllTheArray(numbers,ejercicio37.ejercicio9.saberSiEsteNumeroEsMultiploDeEste);
}
func (ejercicio37*Ejercicio37) theMultiplesNumbersOfOthersAre(numbers []int)(bool,int){
	multiploNumbers :=  ejercicio37.whichIsMultiploOfWhich(numbers);
	var colection string;
	var numbersRemove = numbers;
	for _,element := range multiploNumbers{
		colection += " "+	strconv.Itoa(element);
		ejercicio37.ejercicio25.removeNumber = element;
		numbersRemove = arrayMetodos.Filter(numbersRemove,ejercicio37.ejercicio25.removeThisNumber);
	}			
	fmt.Println( "Los numeros multiplos de otros son "+ colection);
	return len(multiploNumbers) >0,numbersRemove[0];
}
func (ejercicio37*Ejercicio37) ejecutarEjercicio(numeros []int)(bool, int){
	return  ejercicio37.theMultiplesNumbersOfOthersAre(numeros);
}
func  ejecutarEjercicio37(){
	ejercicio37 := Ejercicio37{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=2;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio37.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = " ";
	guardarResultadosMostrar.respuestaNegativa = "Ningun no numero es multiplo del otro";	
	ejecutarEjercicio("37. Leer dos números enteros y determinar cuál es múltiplo de cuál.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
38. Leer tres números enteros y determinar si el último dígito de los tres números es igual.*/
type Ejercicio38 struct{
	ejercicio4 Ejercicio4
	ejercicio25 Ejercicio25
}

func (ejercicio38*Ejercicio38) numberLastDigit(number int) int{
	firstNumberDigit := ejercicio38.ejercicio4.ArrayDigitosNumero(number);
	return firstNumberDigit[len(firstNumberDigit)-1];
}
func (ejercicio38*Ejercicio38) lastDigitEqual(numeros []int)bool {
	numberLastDigito := arrayMetodos.MapTransformData(numeros, ejercicio38.numberLastDigit);
	ejercicio38.ejercicio25.removeNumber = numberLastDigito[0];
	numbersRemove := arrayMetodos.Filter(numberLastDigito,ejercicio38.ejercicio25.removeThisNumber);			
	return len(numbersRemove) == 0;
}
func (ejercicio38*Ejercicio38) ejecutarEjercicio(numeros []int)(bool, int){
	return  ejercicio38.lastDigitEqual(numeros),ejercicio38.ejercicio25.removeNumber;
}
func  ejecutarEjercicio38(){
	ejercicio38 := Ejercicio38{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=3;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio38.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Sus ultimos digitos son iguales";
	guardarResultadosMostrar.respuestaNegativa = "Sus ultimos digitos no son iguales";	
	ejecutarEjercicio("38. Leer tres números enteros y determinar si el último dígito de los tres números es igual.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
39. Leer tres números enteros y determina si el penúltimo dígito de los tres números es igual.
*/
type Ejercicio39 struct{
	ejercicio4 Ejercicio4
	ejercicio25 Ejercicio25
}

func (ejercicio39*Ejercicio39) numberLastDigit(number int) int{
	firstNumberDigit := ejercicio39.ejercicio4.ArrayDigitosNumero(number);
	var numberLastLast int;
	if len(firstNumberDigit) > 1 {
		numberLastLast = firstNumberDigit[len(firstNumberDigit)-2]
	}else{
		numberLastLast = firstNumberDigit[len(firstNumberDigit)-1]
	}
	return numberLastLast;
}
func (ejercicio39*Ejercicio39) lastLastDigitEqual(numeros []int)bool {
	numberLastDigito := arrayMetodos.MapTransformData(numeros, ejercicio39.numberLastDigit);
	ejercicio39.ejercicio25.removeNumber = numberLastDigito[0];
	numbersRemove := arrayMetodos.Filter(numberLastDigito,ejercicio39.ejercicio25.removeThisNumber);			
	fmt.Println("El resultado es ",numbersRemove);						
	return len(numbersRemove) == 0;
}
func (ejercicio39*Ejercicio39) ejecutarEjercicio(numeros []int)(bool, int){
	return  ejercicio39.lastLastDigitEqual(numeros),ejercicio39.ejercicio25.removeNumber;
}
func  ejecutarEjercicio39(){
	ejercicio39 := Ejercicio39{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=3;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio39.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Sus antepenultimos digitos son iguales";
	guardarResultadosMostrar.respuestaNegativa = "Sus antepenultimos digitos no son iguales";	
	ejecutarEjercicio("39. Leer tres números enteros y determina si el penúltimo dígito de los tres números es igual.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
40. Leer dos números enteros y si la diferencia entre los dos es menor o igual a 10 entonces
mostrar en pantalla todos los enteros comprendidos entre el menor y el mayor de los números
leídos.
*/
type Ejercicio40 struct{
	ejercicio11 Ejercicio11 
}
func (ejercicio40*Ejercicio40) differenceNumbersLessThanOrEqualToTen(number []int) bool{
	diferent := number[0] - number[1];
	if diferent < 0 {
		diferent = diferent *-1;
	}
	return diferent <= 10 ;
}
func (ejercicio40*Ejercicio40) printNumberMinUntilMax(number []int){
	numberMax := ejercicio40.ejercicio11.leerNumerosEnterosDeterminarMayor(number);
	numberMin := ejercicio40.ejercicio11.leerNumerosEnterosDeterminarMenor(number);
	for starNumber :=  numberMin; starNumber <=  numberMax;starNumber++{
		fmt.Println("El conteo de numero es ",starNumber);					
	}
}
func (ejercicio40*Ejercicio40) intergerBetweenNumberMaxAndMin(number []int)bool{						
	var answer bool;
	if ejercicio40.differenceNumbersLessThanOrEqualToTen(number) {
		ejercicio40.printNumberMinUntilMax(number);
		answer = true;
	};
	return answer;
}

func (ejercicio40*Ejercicio40) ejecutarEjercicio(numeros []int)(bool, int){
	return ejercicio40.intergerBetweenNumberMaxAndMin(numeros),numeros[0];
}
func  ejecutarEjercicio40(){
	ejercicio40 := Ejercicio40{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=2;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio40.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "La diferencia es menor igual a 10";
	guardarResultadosMostrar.respuestaNegativa = "La diferencia es mayor a 10";	
	ejecutarEjercicio("40. Leer dos números enteros y si la diferencia entre los dos es menor o igual a 10 entonces 		mostrar en pantalla todos los enteros comprendidos entre el menor y el mayor de los números 		leídos.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
41. Leer dos números enteros y determinar si la diferencia entre los dos es un número primo.
*/
type Ejercicio41 struct{
	ejercicio23 Ejercicio23 
	colecctionNumber []int
}
func (ejercicio41*Ejercicio41)subtraction(number1 int,number2 int) int { 
	return number1 - number2;
}
func (ejercicio41*Ejercicio41)subtractionNumbers(numbers []int) int { 
	return arrayMetodos.ReduceInt(numbers,ejercicio41.subtraction);
}	
func (ejercicio41*Ejercicio41)collectSubtraction(numbers1 int,numbers2 int) int {
	subtractResult := numbers1 - numbers2;
	subtractResultabsolute := int(math.Abs( float64(subtractResult)));
	ejercicio41.colecctionNumber = append(ejercicio41.colecctionNumber,subtractResultabsolute);
	return numbers2;
}
func (ejercicio41*Ejercicio41) takeTheDifferenceBetweenNumbers(numeros []int)[]int {
	arrayMetodos.ReduceInt(numeros,ejercicio41.collectSubtraction);
	// fmt.Println("La diferencia es",ejercicio41.colecctionNumber);				
	return ejercicio41.colecctionNumber[1:];
}
func (ejercicio41*Ejercicio41) takeTheDifferenceBetweenNumbersIsPrime(numeros []int)[]int {
	colecctionNumber:= ejercicio41.takeTheDifferenceBetweenNumbers(numeros);
	primeColection := arrayMetodos.Filter(colecctionNumber,ejercicio41.ejercicio23.thisNumberIntIsPrime);
	return primeColection;	
}
	
func (ejercicio41*Ejercicio41) ejecutarEjercicio(numeros []int)(bool, int){
	resultPrime := ejercicio41.takeTheDifferenceBetweenNumbersIsPrime(numeros);
	fmt.Println("Los primos son ",resultPrime);			
	return len(resultPrime)>0,resultPrime[0];
}
func  ejecutarEjercicio41(){
	ejercicio41 := Ejercicio41{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=2;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio41.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "La diferencia de los numeros es primo";
	guardarResultadosMostrar.respuestaNegativa = "La diferencia de los numeros no es primo";	
	ejecutarEjercicio("41. Leer dos números enteros y determinar si la diferencia entre los dos es un número primo.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
42. Leer dos números enteros y determinar si la diferencia entre los dos es un número par.
*/
type Ejercicio42 struct{
	ejercicio5 Ejercicio5 	
	ejercicio41 Ejercicio41 
}
func (ejercicio42*Ejercicio42) takeTheDifferenceBetweenNumbersIsPar(numeros []int)[]int {
	diferentNumber := ejercicio42.ejercicio41.takeTheDifferenceBetweenNumbers(numeros);
	fmt.Println("diferencia ",diferentNumber);
	numberPar := arrayMetodos.Filter(diferentNumber,ejercicio42.ejercicio5.determinarSiPar);
	fmt.Println("numeros pares ",numberPar);
	return numberPar;
}	
func (ejercicio42*Ejercicio42) ejecutarEjercicio(numeros []int)(bool, int){
	resultPar := ejercicio42.takeTheDifferenceBetweenNumbersIsPar(numeros);			
	return len(resultPar)>0,resultPar[0];
}
func  ejecutarEjercicio42(){
	ejercicio42 := Ejercicio42{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=5;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio42.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "La diferencia de los numeros es pares es ";
	guardarResultadosMostrar.respuestaNegativa = "La diferencia de los numeros no es pares es ";	
	ejecutarEjercicio("42. Leer dos números enteros y determinar si la diferencia entre los dos es un número par.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
43. Leer dos números enteros y determinar si la diferencia entre los dos es un número divisor
exacto de alguno de los dos números.
*/
type Ejercicio43 struct{	
	ejercicio41 Ejercicio41 
	ejercicio9 Ejercicio9
	numbercolection []int
}
func (ejercicio43*Ejercicio43) ifDifferenceIsdiv(numberDiferent int)bool {
	var resultado bool;
	for _,element := range ejercicio43.numbercolection{		        
		resultado = ejercicio43.ejercicio9.saberSiEsteNumeroEsMultiploDeEste(element,numberDiferent);
		if resultado {
			break;
		}
	}
	return resultado;
}
	
func (ejercicio43*Ejercicio43) takeTheDifferenceBetweenNumbersIsdiv(numeros []int)[]int {
	ejercicio43.numbercolection =  numeros;
	diferentNumber := ejercicio43.ejercicio41.takeTheDifferenceBetweenNumbers(numeros);
	fmt.Println("diferencia ",diferentNumber);
	divisores := arrayMetodos.Filter(diferentNumber,ejercicio43.ifDifferenceIsdiv);	
	fmt.Println("divisores",divisores);
	return divisores;
}	
func (ejercicio43*Ejercicio43) ejecutarEjercicio(numeros []int)(bool, int){
	resultDiv := ejercicio43.takeTheDifferenceBetweenNumbersIsdiv(numeros);			
	return len(resultDiv)>0,resultDiv[0];
}
func  ejecutarEjercicio43(){
	ejercicio43 := Ejercicio43{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=5;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio43.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "La diferencia si es un divisor";
	guardarResultadosMostrar.respuestaNegativa = "La diferencia no es un divisor";	
	ejecutarEjercicio("43. Leer dos números enteros y determinar si la diferencia entre los dos es un número divisor	exacto de alguno de los dos números.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
44. Leer un número entero de 4 dígitos y determinar si el primer dígito es múltiplo de alguno de los
otros dígitos.
*/
type Ejercicio44 struct{	
	ejercicio9 Ejercicio9
	frisNumber int
	ejercicio4 Ejercicio4
}
func (ejercicio44*Ejercicio44) mumberMultipleThatNumber(number int)bool {
	return ejercicio44.ejercicio9.saberSiEsteNumeroEsMultiploDeEste(number,ejercicio44.frisNumber);
}
func (ejercicio44*Ejercicio44) theFrisNUmberMultipleOtherNumbers(number  int)[]int {
	digigitNumber := ejercicio44.ejercicio4.ArrayDigitosNumero(number);
	otherDigits := digigitNumber[1:];
	ejercicio44.frisNumber = digigitNumber[0];
	return arrayMetodos.Filter(otherDigits,ejercicio44.mumberMultipleThatNumber);
}		
func (ejercicio44*Ejercicio44) ejecutarEjercicio(numeros []int)(bool, int){
	resultDiv := ejercicio44.theFrisNUmberMultipleOtherNumbers(numeros[0]);
	fmt.Println("digitos divididos por el primer digito",resultDiv);		
	return len(resultDiv)>0,resultDiv[0];
}
func  ejecutarEjercicio44(){
	ejercicio44 := Ejercicio44{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarCuatroDigitos;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio44.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El primer digito divisor de alguno de los otros digitos";
	guardarResultadosMostrar.respuestaNegativa = "El primer digito no es divisor de alguno de los otros digitos";	
	ejecutarEjercicio("44. Leer un número entero de 4 dígitos y determinar si el primer dígito es múltiplo de alguno de los	otros dígitos..",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
45. Leer un número entero de 2 dígitos y si es par mostrar en pantalla la suma de sus dígitos, si es
primo y menor que 10 mostrar en pantalla su último dígito y si es múltiplo de 5 y menor que 30
mostrar en pantalla el primer dígito.*/
type Ejercicio45 struct{	
	ejercicio5 Ejercicio5
	ejercicio4 Ejercicio4
	ejercicio9 Ejercicio9
	ejercicio23 Ejercicio23
}
func (ejercicio45*Ejercicio45) ifItIsCoupleNumberShowSumDigits(numero int)(bool,int){
	var digitSum int;
	var respuesta bool;
	if ejercicio45.ejercicio5.determinarSiPar(numero) {
		digitSum = ejercicio45.ejercicio4.sumarDigitos(numero)
		respuesta = true;
	}
	return respuesta, digitSum;
}
func (ejercicio45*Ejercicio45) ifCousinAndLessThan10ShowOnTheScreenItsLastDigit(numero int)(bool,int){
	var lastDigit int;
	var respuesta bool;
	if ejercicio45.ejercicio23.thisNumberIntIsPrime(numero) &&  numero < 10 {
		arrayDigit := ejercicio45.ejercicio4.ArrayDigitosNumero(numero);
		lastDigit = arrayDigit[len(arrayDigit)];
		respuesta = true;
	}
	return respuesta, lastDigit;
}		


func (ejercicio45*Ejercicio45) ifItIsMultipleOf5AndLessThan30ShowFirstDigitOnScreen(numero int)(bool,int){	
	var fristDigit int;
	var respuesta bool;
	if ejercicio45.ejercicio9.saberSiEsteNumeroEsMultiploDeEste(numero,5) &&  numero < 30 {
		fristDigit = ejercicio45.ejercicio4.ArrayDigitosNumero(numero)[0];
		respuesta = true;
	}
	return respuesta, fristDigit;
}		
func (ejercicio45*Ejercicio45) executeExercise45(numero int)bool{
	respontFirstBool,respontFirst := ejercicio45.ifItIsMultipleOf5AndLessThan30ShowFirstDigitOnScreen(numero);
	if respontFirstBool {
			fmt.Println("El resultado es ",respontFirst);
	}
	respontSecondBool,respontSecond := ejercicio45.ifCousinAndLessThan10ShowOnTheScreenItsLastDigit(numero);
	if respontSecondBool {
			fmt.Println("El resultado es ",respontSecond);
	}
	respontThirdBool,respontThird := ejercicio45.ifItIsCoupleNumberShowSumDigits(numero);
	if respontThirdBool {
			fmt.Println("El resultado es ",respontThird);
	}
	return respontFirstBool  || respontSecondBool || respontThirdBool;
}	
func (ejercicio45*Ejercicio45) ejecutarEjercicio(numeros []int)(bool, int){
	return ejercicio45.executeExercise45(numeros[0]),numeros[0];
}
func  ejecutarEjercicio45(){
	ejercicio45 := Ejercicio45{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarDosDigitos;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio45.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Cumple con las condiciones";
	guardarResultadosMostrar.respuestaNegativa = "No cumple con las condiciones";	
	ejecutarEjercicio("45. Leer un número entero de 2 dígitos y si es par mostrar en pantalla la suma de sus dígitos, si es 	primo y menor que 10 mostrar en pantalla su último dígito y si es múltiplo de 5 y menor que 30		mostrar en pantalla el primer dígito.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
46. Leer un número entero de 2 dígitos y si terminar en 1 mostrar en pantalla su primer dígito, si
termina en 2 mostrar en pantalla la suma de sus dígitos y si termina en 3 mostrar en pantalla el
producto de sus dos dígitos.
*/
type Ejercicio46 struct{	
	ejercicio4 Ejercicio4
}
func (ejercicio46*Ejercicio46) product(number1 int,number int) int {
	 if number1 == 0 { number1 =1; }
	return number1 * number;
}
func (ejercicio46*Ejercicio46) productNumbers(numbers []int) int {
	return  arrayMetodos.ReduceInt(numbers, ejercicio46.product);
}
func (ejercicio46*Ejercicio46) firstDigit(number int)int{
	endOne := ejercicio46.ejercicio4.ArrayDigitosNumero(number);
	return endOne[0];
}
func (ejercicio46*Ejercicio46) numberEndsInOneShowInYourFirstDigit(number int)(bool,int){
	endOne := ejercicio46.ejercicio4.ArrayDigitosNumero(number);
	return endOne[len(endOne)-1]== 1,endOne[0];
}
func (ejercicio46*Ejercicio46) numberEndsInTwoShowInTheSumOfItDigits(number int)(bool,int){
	endOne := ejercicio46.ejercicio4.ArrayDigitosNumero(number);
	return endOne[len(endOne)-1]== 2,endOne[len(endOne)-1] + endOne[0];
}
func (ejercicio46*Ejercicio46) numberEndsInThreeShowInTheProductOfItsTwoDigits(number int)(bool,int){
	endOne := ejercicio46.ejercicio4.ArrayDigitosNumero(number);
	return endOne[len(endOne)-1]== 3,endOne[len(endOne)-1] * endOne[0];
}
func (ejercicio46*Ejercicio46) executeExercise46(numero int)bool{
	response1, number1 :=  ejercicio46.numberEndsInOneShowInYourFirstDigit(numero);
	if response1 {
			fmt.Println("Su ultimo digito es ",number1);

	}
	response2, number2 := ejercicio46.numberEndsInTwoShowInTheSumOfItDigits(numero);
	if response2  {
			fmt.Println("Las suma de sus digitos es ",number2);

	}
	response3, number3 := ejercicio46.numberEndsInThreeShowInTheProductOfItsTwoDigits(numero);
	if response3  {
			fmt.Println("La multiplicacion de sus digitos ",number3);

	}
	return response1 ||	response2 ||	response3;
}
func (ejercicio46*Ejercicio46) ejecutarEjercicio(numeros []int)(bool, int){
	return ejercicio46.executeExercise46(numeros[0]),numeros[0];
}
func  ejecutarEjercicio46(){
	ejercicio46 := Ejercicio46{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarDosDigitos;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio46.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Cumple con las condiciones";
	guardarResultadosMostrar.respuestaNegativa = "No cumple con las condiciones";	
	ejecutarEjercicio("46. Leer un número entero de 2 dígitos y si terminar en 1 mostrar en pantalla su primer dígito, si termina en 2 mostrar en pantalla la suma de sus dígitos y si termina en 3 mostrar en pantalla el producto de sus dos dígitos.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
47. Leer dos números enteros y si la diferencia entre los dos números es par mostrar en pantalla la
suma de los dígitos de los números, si dicha diferencia es un número primo menor que 10
entonces mostrar en pantalla el producto de los dos números y si la diferencia entre ellos
terminar en 4 mostrar en pantalla todos los dígitos por separado.
*/
type Ejercicio47 struct{	
	ejercicio5 Ejercicio5
	ejercicio4 Ejercicio4
	ejercicio41 Ejercicio41
	ejercicio9 Ejercicio9
	ejercicio13 Ejercicio13
	ejercicio46 Ejercicio46
}
func (ejercicio47*Ejercicio47) minorTen(number int) bool {
	return number < 10;
}
func (ejercicio47*Ejercicio47) differenceBetweenTheTwoNumbersIsToShowAddDigit(numbers []int)(bool,int){
	diferenNumber := ejercicio47.ejercicio41.takeTheDifferenceBetweenNumbers(numbers);	 
	arrayPlusDigit := arrayMetodos.MapTransformData(numbers, ejercicio47.ejercicio4.sumarDigitos);
	plusDigit :=  ejercicio47.ejercicio13.sumarArrayDeNumeros(arrayPlusDigit);
	return arrayMetodos.Some(diferenNumber,ejercicio47.ejercicio5.determinarSiPar),plusDigit;
}

func (ejercicio47*Ejercicio47) differenceNumberPrimeTenMinorShowProductDigits(numbers []int)(bool,int){
	primeNumber := ejercicio47.ejercicio41.takeTheDifferenceBetweenNumbersIsPrime(numbers);
	ifPrimeNumber :=  len(primeNumber)>0;
	ifPrimeNumberMinorTen := arrayMetodos.Some(primeNumber,ejercicio47.minorTen);
	productNumber := ejercicio47.ejercicio46.productNumbers(numbers);
	diferentPrimeTenMinor :=  ifPrimeNumber && ifPrimeNumberMinorTen;
	return diferentPrimeTenMinor,productNumber;
}
func (ejercicio47*Ejercicio47) differenceFinishFourShowAlldigitsSeparately(numbers []int)(bool,[]int){
	diferenceNumber := ejercicio47.ejercicio41.subtractionNumbers(numbers);
	differenceTerminateFour := diferenceNumber == 4;
	return differenceTerminateFour,numbers;
}
func (ejercicio47*Ejercicio47) executeExercise47(numero []int)bool{
	response1, number1 :=  ejercicio47.differenceBetweenTheTwoNumbersIsToShowAddDigit(numero);
	if response1 {
			fmt.Println("Diferencia de los dos numeros es par la suma de sus digitos es ",number1);

	}
	response2, number2 := ejercicio47.differenceNumberPrimeTenMinorShowProductDigits(numero);
	if response2  {
			fmt.Println("La diferencia de los numeros es primo menor que 10, el  producto de los dos números",number2);

	}
	response3, number3 := ejercicio47.differenceFinishFourShowAlldigitsSeparately(numero);
	if response3  {
			fmt.Println("La diferencia de los numeros menos que 4 ",number3);

	}
	return response1 || response2 || response3;
}
func (ejercicio47*Ejercicio47) ejecutarEjercicio(numeros []int)(bool, int){
	return ejercicio47.executeExercise47(numeros),numeros[0];
}
func  ejecutarEjercicio47(){
	ejercicio47 := Ejercicio47{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=2;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.IngresarDosDigitos;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio47.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "Cumple con las condiciones";
	guardarResultadosMostrar.respuestaNegativa = "No cumple con las condiciones";	
	ejecutarEjercicio("47. Leer dos números enteros y si la diferencia entre los dos números es par mostrar en pantalla la suma de los dígitos de los números, si dicha diferencia es un número primo menor que 10 entonces mostrar en pantalla el producto de los dos números y si la diferencia entre ellos terminar en 4 mostrar en pantalla todos los dígitos por separado.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
48. Leer un número entero y si es menor que 100 determinar si es primo.
*/
type Ejercicio48 struct{	
	ejercicio6 Ejercicio6
}
func (ejercicio48*Ejercicio48)lessThanOneHundred(number int) bool{
	return number < 100;
}
func (ejercicio48*Ejercicio48)lessThanOneHundredIfPrimeNumber(numbers []int) bool {
	return  arrayMetodos.Some(numbers,ejercicio48.ejercicio6.DeterminarSiPrimoInt ) &&	arrayMetodos.Some(numbers,ejercicio48.lessThanOneHundred );	
}
func (ejercicio48*Ejercicio48) ejecutarEjercicio(numeros []int)(bool, int){
	answer := ejercicio48.lessThanOneHundredIfPrimeNumber(numeros);
	return answer,numeros[0];
}
func  ejecutarEjercicio48(){
	ejercicio48 := Ejercicio48{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio48.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El numero es primo menor que 100";
	guardarResultadosMostrar.respuestaNegativa = "El numero no es primo o no es menor que 100";	
	ejecutarEjercicio("48. Leer un número entero y si es menor que 100 determinar si es primo.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
49. Leer un número entero y si es múltiplo de 4 determinar si su último dígito es primo.
*/

type Ejercicio49 struct{
	ejercicio6 Ejercicio6	
	ejercicio9 Ejercicio9
	ejercicio38 Ejercicio38
}
func (ejercicio49*Ejercicio49)multipleOfFour(number int) bool{
	return ejercicio49.ejercicio9.saberSiEsteNumeroEsMultiploDeEste(number , 4);
}
func (ejercicio49*Ejercicio49)ifMultipleOfFour(numbers []int) bool {
	return arrayMetodos.Some(numbers,ejercicio49.multipleOfFour ) ;	
}
func (ejercicio49*Ejercicio49)ifPrimeNumberLastDigit(numbers []int) bool {	
	numberLastDigito := arrayMetodos.MapTransformData(numbers, ejercicio49.ejercicio38.numberLastDigit);	
	return arrayMetodos.Some(numberLastDigito,ejercicio49.ejercicio6.DeterminarSiPrimoInt )
}
func (ejercicio49*Ejercicio49) ejecutarEjercicio(numeros []int)(bool, int){
	multipleFour := ejercicio49.ifMultipleOfFour(numeros);
	primeNumber := ejercicio49.ifPrimeNumberLastDigit(numeros);
	if multipleFour {
		fmt.Println("Es multiplo de cuatro ");
		if primeNumber{
			fmt.Println("Es numero primo ");
		}
	}
	return multipleFour && primeNumber,numeros[0];
}
func  ejecutarEjercicio49(){
	ejercicio49 := Ejercicio49{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=1;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.AgregarValoreVariableEntera;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio49.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El numero es primo y multiplo de 4";
	guardarResultadosMostrar.respuestaNegativa = "El numero  no es primo o no es multiplo de 4";	
	ejecutarEjercicio("49. Leer un número entero y si es múltiplo de 4 determinar si su último dígito es primo.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}
/*
50. Leer un número entero y si es múltiplo de 4 mostrar en pantalla su mitad, si es múltiplo de 5
mostrar en pantalla su cuadrado y si es múltiplo e 6 mostrar en pantalla su primer dígito.
Asumir que el número no es mayor que 100.
*/
type Ejercicio50 struct{	
	ejercicio9 Ejercicio9
	ejercicio46 Ejercicio46
}
func (ejercicio50*Ejercicio50) multiploOfFour(numero int)bool{
	return ejercicio50.ejercicio9.saberSiEsteNumeroEsMultiploDeEste(numero,4);
}
func (ejercicio50*Ejercicio50) multiploOfFive(numero int)bool{
	return ejercicio50.ejercicio9.saberSiEsteNumeroEsMultiploDeEste(numero,5);
}
func (ejercicio50*Ejercicio50) multiploOfSix(numero int)bool{
	return ejercicio50.ejercicio9.saberSiEsteNumeroEsMultiploDeEste(numero,6);
}
func (ejercicio50*Ejercicio50) showYourHalf(number int)int {
	return number/2;
}

func (ejercicio50*Ejercicio50) showitsSquare(number int)int {
	return number*number;
}
func (ejercicio50*Ejercicio50) ifMultiploOfFourShowYourHalf(numbers []int) []int{
	var answer []int;
	if arrayMetodos.Every(numbers,ejercicio50.multiploOfFour) {
		answer = arrayMetodos.MapTransformData(numbers, ejercicio50.showYourHalf);	

	}
	return answer;
}
func (ejercicio50*Ejercicio50) ifMultiploOfFiveShowitsSquare(numbers []int)[]int{
	var answer []int;
	if arrayMetodos.Every(numbers,ejercicio50.multiploOfFive) {
		answer = arrayMetodos.MapTransformData(numbers, ejercicio50.showitsSquare);	
	}
	return answer;
}
func (ejercicio50*Ejercicio50) ifMultiploOfSixShowFirstDigit(numbers []int)[]int{
	var answer []int;
	if arrayMetodos.Every(numbers,ejercicio50.multiploOfSix) {	
		answer =  arrayMetodos.MapTransformData(numbers,ejercicio50.ejercicio46.firstDigit);
	}
	return answer;
}
func (ejercicio50*Ejercicio50) ejecutarEjercicio(numbers []int)(bool, int){
	multiploOfFourShowYourHalf := ejercicio50.ifMultiploOfFourShowYourHalf(numbers);
	multiploOfFiveShowitsSquare := ejercicio50.ifMultiploOfFiveShowitsSquare(numbers);
	multiploOfSixShowFirstDigit := ejercicio50.ifMultiploOfSixShowFirstDigit(numbers)
	// make([]int,  numeroDatosIngresar);
	if len(multiploOfFourShowYourHalf)>0 {
		fmt.Println("Es multiplo de cuatro y su mitad es ",multiploOfFourShowYourHalf);
	}
	if len(multiploOfFiveShowitsSquare)>0 {
		fmt.Println("Es multiplo de cinco y su cuadrado es ",multiploOfFiveShowitsSquare);
	}
	if len(multiploOfSixShowFirstDigit)>0 {
		fmt.Println("Es multiplo de seis y su primer digito es ",multiploOfSixShowFirstDigit);
	}
	return len(multiploOfFourShowYourHalf)>0 || len(multiploOfFiveShowitsSquare)>0 || len(multiploOfSixShowFirstDigit)>0 , numbers[0];
}
func  ejecutarEjercicio50(){
	ejercicio50 := Ejercicio50{};
	guardarResultadosMostrar := GuardarResultadosMostrar{};
	guardarResultadosMostrar.cantidadDatosIngresar=3;
	guardarResultadosMostrar.cantidadDigitosPermitidos = agregarMostrarValores.NumeroPositivoMenorDe100;
	guardarResultadosMostrar.funcionRecibirArrayInt = ejercicio50.ejecutarEjercicio;
	guardarResultadosMostrar.respuestaPositiva = "El numero cumple con alguno de los requisitos";
	guardarResultadosMostrar.respuestaNegativa = "El numero cumple con ninguno de los requisitos";	
	ejecutarEjercicio("50. Leer un número entero y si es múltiplo de 4 mostrar en pantalla su mitad, si es múltiplo de 5	mostrar en pantalla su cuadrado y si es múltiplo e 6 mostrar en pantalla su primer dígito.	Asumir que el número no es mayor que 100.",guardarResultadosMostrar.ingresarValoresMostrarResultados);	
}


func EjecutarPrograma(){
	controlarFuncionamiento := controlarFuncionamientoPrograma.ConstruirTitulo("condicionales");	
	fmt.Println("Hola, bienvenido al programa de condicionales");
	controlarFuncionamiento.Administrar(ejercicios);
}

type  ObjetosCondicionales  struct {		
	Ejercicio2 Ejercicio2		
	Ejercicio4 Ejercicio4		
	Ejercicio6 Ejercicio6
	Ejercicio17 Ejercicio17	
	Ejercicio21 Ejercicio21	
	}

func ExportarObjetosCondicionales() ObjetosCondicionales{
	return ObjetosCondicionales{}
}
