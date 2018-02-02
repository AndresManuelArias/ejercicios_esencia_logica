package main

import (
  "./variablesConstantesOperadores"
  "./controlarFuncionamientoPrograma"
  "./condicionales"
  "./ciclos"
  "./array"
  "./matrices"
  "fmt"
)

func ejerciciosEsenciaLogica(){
	var numeroPregunta string;
	//controlarFuncionamientoPrograma.GetColocarTituloPrograma("esencia logica");
	fmt.Println("Ejercicios de esencia logica");	
	fmt.Println("Escriba el numero de cual capitulo quiere ver los ejercicios:");
	fmt.Println("3. Capitulos 3");	
	fmt.Println("8. Capitulos 8");
	fmt.Println("9. Capitulos 9");		
	fmt.Println("10. Capitulos 10");		
	fmt.Println("11. Capitulos 11");		
	fmt.Scanf("%s\n", &numeroPregunta);		
	switch numeroPregunta {
		case "3":
			variablesConstantesOperadores.EjecutarPrograma();
		break;
		case "8":
			condicionales.EjecutarPrograma();
		break;
		case "9":
			ciclos.EjecutarPrograma();
		break;	
		case "10":
			array.EjecutarPrograma();
		break;
		case "11":
			matrices.EjecutarPrograma();
		break;		
		default:
	}
}


func main() {
	tituloPrograma := controlarFuncionamientoPrograma.ConstruirTitulo("esencia logica");	
	fmt.Println("el titulo es ", tituloPrograma.SetMostrarTituloPrograma());
	tituloPrograma.Administrar(ejerciciosEsenciaLogica);
}