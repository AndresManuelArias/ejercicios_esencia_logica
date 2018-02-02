package controlarFuncionamientoPrograma
import 	"fmt"
  

type ControlarPrograma struct{
	titulo string
	seguir bool
}

func (controlarPrograma ControlarPrograma) getdefinirSegir(verdaderoOFalso bool){
	controlarPrograma.seguir = verdaderoOFalso;
}
func (controlarPrograma ControlarPrograma) setMostrarDefinirSeguir()bool{
	return controlarPrograma.seguir;
}
func ConstruirTitulo( titulo string)ControlarPrograma{
	return ControlarPrograma{titulo:titulo}
}
func(controlarPrograma ControlarPrograma) SetMostrarTituloPrograma()string {
	return controlarPrograma.titulo;
}
func (t ControlarPrograma) getColocarTituloPrograma(titulo string){
	t.titulo = titulo;
}
func (controlarPrograma ControlarPrograma) terminarPrograma() bool {
	var respuesta string;
	fmt.Println("Si quiere salirse del programa << "+ controlarPrograma.SetMostrarTituloPrograma()+" >> oprima Y\nSi no oprima cualquier tecla " );			
	fmt.Scanf("%s\n", &respuesta) // add "\n"
	if respuesta == "Y" || respuesta == "y" {
		controlarPrograma.seguir = false;
	}
	return controlarPrograma.seguir;
}
func  (controlarPrograma ControlarPrograma)  Administrar(f func()){
	controlarPrograma.seguir = true;
	for controlarPrograma.seguir {
		f()
		controlarPrograma.seguir = controlarPrograma.terminarPrograma();
	}

}
