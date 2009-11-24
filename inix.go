package inix

import (
	"os";
	"io";
	"strings";
)

const(
	comentario byte = 0x23; // Ascii 23 = #
	corchete byte = 0x5B; //Ascii 0x5B = [
	cierracorchete = 0x5D; // Ascii 0x5D = ]
)

type contenido map[string]string;
type seccion map[string]contenido;

// Estructura que define un archivo '.ini'.
// El elemento Seccion es un mapa de mapas.
// Las secciones en el .ini deben tener la forma [NombreSeccion]. 
// Los elementos válidos tienen la forma clave=valor.
// Para recuperar un elemento se puede hacer Seccion[NombreSeccion][clave]
type InitFile struct {
	Nombre string;
	Seccion seccion;
}

// Crea una estructura vacía representando un archivo '.ini'.
func New(nombre string) *InitFile {
	return &InitFile{ nombre,make(seccion,1) };
}

// ReadAll lee el archivo .ini especificado en la estructura.
// Crea un mapa de mapas representando la estructura del archivo,
// separando secciones y elementos dentro de la seccion y lo almacena
// en la variable Seccion de la estructura InitFile.
func (iF *InitFile) ReadAll() (error os.Error){
	micontenido := make(contenido);
	miseccion := "default";
	if mifichero,error := io.ReadFile(iF.Nombre);error!=nil{
		return error
	} else {
		cadena := string(mifichero);
		lineas := strings.Split(cadena,"\n",0);
		for _,linea := range lineas {
			expresion := strings.TrimSpace(linea);
			if len(expresion) == 0 {
				continue
			}
			if expresion[0]==comentario{
				continue
			}
			if expresion[0]==corchete && expresion[len(expresion)-1]==cierracorchete {
				iF.Seccion[miseccion] = micontenido;
				miseccion = string(expresion[1:len(expresion)-1]);
				micontenido = make(contenido);
				continue;
			}
			valores := strings.Split(expresion,"=",2);
			if len(valores) != 2 {
				continue;
			}		
			micontenido[valores[0]]=valores[1];
		}
	}
	if _,ok := iF.Seccion[miseccion];!ok{
		iF.Seccion[miseccion]=micontenido;
	}
	return nil;
}

