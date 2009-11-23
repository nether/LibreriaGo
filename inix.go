package inix

import (
	"os";
	"io";
	"strings";
	"fmt";
//	"container/vector";
)

type contenido map[string]string;
type seccion map[string]contenido;

type InitFile struct {
	Nombre string;
	Seccion seccion;
}

const(
	comentario byte = 0x23; // Ascii 23 = #
	corchete byte = 0x5B;
	cierracorchete = 0x5D;
)

func New(nombre string) *InitFile {
	return &InitFile{ nombre,make(seccion,1) };
}

func (iF *InitFile) ReadAll() (error os.Error){
	
	micontenido := make(contenido);
	miseccion := "default";
	if mifichero,error := io.ReadFile(iF.Nombre);error!=nil{
		return error
	} else {
		cadena := string(mifichero);
		lineas := strings.Split(cadena,"\n",0);
		for _,linea := range lineas {
//			fmt.Printf("Analizando línea: %s", linea);
			if len(linea) == 0 {
				continue
			}
			if linea[0]==comentario{
				continue
			}
			if linea[0]==corchete && linea[len(linea)-1]==cierracorchete {
//				fmt.Printf(" - Nueva seccion\n");
				iF.Seccion[miseccion] = micontenido;
				miseccion = string(linea[1:len(linea)-1]);
				micontenido = make(contenido);
				continue;
			}
			valores := strings.Split(linea,"=",2);
			if len(valores) != 2 {
				fmt.Printf("Linea incomprensible: %s",linea);
				continue;
			}		
//			fmt.Printf(" - Añadido al mapa (seccion %s).\n",miseccion);
			micontenido[valores[0]]=valores[1];
		}
	}
	if _,ok := iF.Seccion[miseccion];!ok{
//		fmt.Println("\nAñadiendo contenido final.");
		iF.Seccion[miseccion]=micontenido;
	}
	return nil;
}

func ReadInit(file string) (mapa map[string]string, error os.Error) {
	mapa = make(map[string]string);
	if contenido,error := io.ReadFile(file);error!=nil{
		return nil,error
	} else {
		cadena := string(contenido);
		lineas := strings.Split(cadena,"\n",0);
		for _,linea := range lineas {
			if len(linea) == 0 {
				continue
			}
			if linea[0]==comentario{
				continue
			}
			valores := strings.Split(linea,"=",2);
			if len(valores) != 2 {
				continue;
			}		
			mapa[valores[0]]=valores[1];
		}
	}
	return mapa,nil;
}

func WriteInit(mapa map[string]string, file string, perm int) (error os.Error) {
	f, e := os.Open(file, perm, os.O_CREAT | os.O_WRONLY);
	if e!= nil {
		return e;
	}
	for key,value := range mapa {
		if _,err := fmt.Fprintf(f,"%s=%s\n",key,value); err!=nil{
			return err;
		}
	}
	return nil;
}