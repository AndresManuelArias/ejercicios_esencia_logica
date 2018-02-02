// ejercicios de la pagina 42 en golang
package variablesConstantesOperadores

import ( "fmt"
		"../controlarFuncionamientoPrograma"
		"../agregarMostrarValores"
)

// alineacion pagina 42
func ejercicio11()float64  {
	fmt.Println("pregunta 11")	
	a ,b , c := agregarMostrarValores.AgregarValoresVariablesABC();
	return (a+b/c)/(a/b+c);
}
func ejercicio12()float64  {
	fmt.Println("pregunta 12")	
	a ,b , c := agregarMostrarValores.AgregarValoresVariablesABC();
	return (a+b +a/b)/c;
}
func ejercicio13()float64  {
	fmt.Println("pregunta 13")	
	a , b  := agregarMostrarValores.AgregarValoresVariablesAB();
	return (a/(a+b))/(a/(a-b));
}
func ejercicio14()float64  {
	fmt.Println("pregunta 14")	
	a , b ,c := agregarMostrarValores.AgregarValoresVariablesABC();
	return (a+b/(a+b+b/c))/(a+b/(c+a));
}
func ejercicio15()float64  {
	fmt.Println("pregunta 15")	
	a , b ,c := agregarMostrarValores.AgregarValoresVariablesABC();
	return a+b+c/(a+b/c);
}
func ejercicio16()float64  {
	fmt.Println("pregunta 16")	
	a , b ,c,d := agregarMostrarValores.AgregarValoresVariablesABCD();
	return (a+b+c/(d*a))/(c/d*b+a);
}
func ejercicio17()float64  {
	fmt.Println("pregunta 17")	
	a , b ,c,d := agregarMostrarValores.AgregarValoresVariablesABCD();
	return (a+b/c+d)/a;
}
func ejercicio18()float64  {
	fmt.Println("pregunta 18")	
	a , b ,c := agregarMostrarValores.AgregarValoresVariablesABC();
	return (a/b + b/c)/(a/b - b/c);
}
func ejercicio19()float64  {
	fmt.Println("pregunta 19")	
	a , b ,c,d := agregarMostrarValores.AgregarValoresVariablesABCD();
	return a+(a+(a+b/c+d))/(a+a/b);
}
func ejercicio20()float64  {
	fmt.Println("pregunta 20")	
	a , b ,c,d := agregarMostrarValores.AgregarValoresVariablesABCD();
	return a+b+c/d+(a/(b-c)/(a/(b+c)));
}

func ejercicios(){
	var numeroPregunta string
	fmt.Println("Escoja el número de la pregunta; que son el 11,12,13,14,15,16,17,18,19 o 20")	
	fmt.Scanf("%s\n", &numeroPregunta)	
	switch numeroPregunta {
		case "11":
			agregarMostrarValores.RespuestaEjercicio(ejercicio11())
		break;
		case "12":
			agregarMostrarValores.RespuestaEjercicio(ejercicio12())
		break;
		case "13":
			agregarMostrarValores.RespuestaEjercicio(ejercicio13())
		break;
		case "14":
			agregarMostrarValores.RespuestaEjercicio(ejercicio14())
		break;
		case "15":
			agregarMostrarValores.RespuestaEjercicio(ejercicio15())
		break;
		case "16":
			agregarMostrarValores.RespuestaEjercicio(ejercicio16())
		break;
		case "17":
			agregarMostrarValores.RespuestaEjercicio(ejercicio17())
		break;
		case "18":
			agregarMostrarValores.RespuestaEjercicio(ejercicio18())
		break;
		case "19":
			agregarMostrarValores.RespuestaEjercicio(ejercicio19())
		break;
		case "20":
			agregarMostrarValores.RespuestaEjercicio(ejercicio20())
		break;		
		default:
	}
}
func EjecutarPrograma(){
	controlarFuncionamiento := controlarFuncionamientoPrograma.ConstruirTitulo("variables constantes y operadores");	
	fmt.Println("Hola, bienvenido al programa de variables, constantes y operadores, donde estan los ejercicios del 11 al 20");
	controlarFuncionamiento.Administrar(ejercicios);
}
