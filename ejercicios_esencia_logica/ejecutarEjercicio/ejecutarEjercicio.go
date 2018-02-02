package ejecutarEjercicio

import ( "fmt"
	"../agregarMostrarValores"
)
/*
arreglar de tal manera que el resultado de un solo digito se resiva por un array 
que al final pueda mostrar como resultado una impresion en consola
*/

type EjecucionEjercicioMatrices  struct{
	titulo string
	columnas int
	filas int
	ejercicioMatrix func([][][]int)string
	cantidadMatricesIngresar  int
}
func(ejecucionEjercicioMatrices*EjecucionEjercicioMatrices)EjecutarEjercicioMatrices(){
	fmt.Println(ejecucionEjercicioMatrices.titulo);
	var colectionMatrices [][][]int;
	for matricesIngresadas := 0;matricesIngresadas < ejecucionEjercicioMatrices.cantidadMatricesIngresar;matricesIngresadas++{
		fmt.Println("Agregando datos a matrix numero",matricesIngresadas+1);		
		datosMatrix := agregarMostrarValores.AgregarValoreVariableEnteraMatrix(ejecucionEjercicioMatrices.filas,ejecucionEjercicioMatrices.columnas);
		colectionMatrices = append(colectionMatrices,datosMatrix);
	}
	fmt.Println("Numeros ingresados ",colectionMatrices);		
	fmt.Println(ejecucionEjercicioMatrices.ejercicioMatrix(colectionMatrices)	)		
}
func NewEjecucionEjercicioMatrices(	titulo string,
									filasColumna [2]int,													
									ejercicioMatrix func([][][]int)string,
									cantidadMatricesIngresar  int){
	ejecucionEjercicioMatrices := EjecucionEjercicioMatrices{titulo:titulo,
										filas:filasColumna[0],
										columnas:filasColumna[1],
										ejercicioMatrix:ejercicioMatrix,
										cantidadMatricesIngresar:cantidadMatricesIngresar};
	ejecucionEjercicioMatrices.EjecutarEjercicioMatrices();
}

type EjecucionEjercicio  struct{
	titulo string
	cantidadDatosIngresar int
	cantidadDigitosPermitidos func() int
	numeroDatosIngresar func([]int) bool
}
	
func (ejecucionEjercicio * EjecucionEjercicio) EjecutarEjercicio( funcion func([]int)string){
	imprimirResultado := agregarMostrarValores.ClaseEjercicio(ejecucionEjercicio.titulo);
	imprimirResultado.MostrarTituloEjercicio();
	numerosIngresados := agregarMostrarValores.DeterminarCuantosDatosIngresar(ejecucionEjercicio.cantidadDatosIngresar,ejecucionEjercicio.cantidadDigitosPermitidos);
	respuestaTexto := funcion(numerosIngresados);
	fmt.Println(respuestaTexto);
}

func (ejecucionEjercicio * EjecucionEjercicio) EjecutarEjercicioSinData( funcion func()string){
	imprimirResultado := agregarMostrarValores.ClaseEjercicio(ejecucionEjercicio.titulo);
	imprimirResultado.MostrarTituloEjercicio();
	respuestaTexto := funcion();
	fmt.Println(respuestaTexto);
}

func (ejecucionEjercicio * EjecucionEjercicio) EjecutarEjercicioF( funcion func([]int)string){
	imprimirResultado := agregarMostrarValores.ClaseEjercicio(ejecucionEjercicio.titulo);
	imprimirResultado.MostrarTituloEjercicio();
	numerosIngresados := agregarMostrarValores.DatosIngresar(ejecucionEjercicio.numeroDatosIngresar,ejecucionEjercicio.cantidadDigitosPermitidos);
	respuestaTexto := funcion(numerosIngresados);
	fmt.Println(respuestaTexto);
}

func HeredarObjeto(titulo string,	cantidadDatosIngresar int,	cantidadDigitosPermitidos func() int ) EjecucionEjercicio {
	return EjecucionEjercicio{titulo:titulo,cantidadDatosIngresar:cantidadDatosIngresar,cantidadDigitosPermitidos:cantidadDigitosPermitidos};
}
func HeredarObjetoNoIngresarDatos(titulo string ) EjecucionEjercicio {
	return EjecucionEjercicio{titulo:titulo};
}
func HeredarObjetoF(titulo string,	numeroDatosIngresar func([]int) bool,	cantidadDigitosPermitidos func() int ) EjecucionEjercicio {
	return EjecucionEjercicio{titulo:titulo,numeroDatosIngresar:numeroDatosIngresar,cantidadDigitosPermitidos:cantidadDigitosPermitidos};
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