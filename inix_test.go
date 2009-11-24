package inix_test

import (
	. "inix";
	"fmt";
	"testing";
	"os";
	"io";
)

/*var seccionTest seccion = map[string]map[string]string {
	"default": map[string]string{ 
		"inicio":"principio"; "prueba espacio":"prueba espacio";
		"prueba igual":"prueba=igual" ;"secciondefault":"default"
	};
	"Seccion 2":map[string]string{
		"otra seccion":"una seccion distinta";"fin":"fin"
	};
}
*/

var textoInix = `
#Archivo para realizar test iniciales de la libreria inix.
#Las líneas que comienzan con # son ignoradas.

#Las primeras igualdades que no van en una seccion 
# se guardan en la seccion default.
inicio=principio
prueba espacio=prueba espacio
prueba igual=prueba=igual

#Las secciones comienzan con [ y terminan en ]. 
#Cualquier otra cosa es ignorada
[aaaaaa
]bbbbbb
[cccccc]c
]dddddd]
[eeeeee[

#Los espacio antes de la seccion son validos
 [Seccion 3]

#Los espacios al final de la seccion son válidos
[Seccion 4] 
secciondefault=default
#Descomentar siguiente línea para probar el test.
#probando=test 

[Seccion 2]
otra seccion=una seccion distinta
fin=fin

#Puesto para comprobar el test
#[Seccion 3]
#pruebatest=el test debe fallar

#Archivo creado desde inix_test.go
`
func crearArchivoIni() os.Error{
	file,err := os.Open("inix.ini", os.O_CREAT | os.O_WRONLY, 0600);
	if err != nil { return err; }
	io.WriteString(file,textoInix);
	return nil;
}
var seccionTest = make(map[string]map[string]string);

func rellenarSeccionTest(){
	seccionTest["default"]=make(map[string]string);
	seccionTest["Seccion 2"]=make(map[string]string);
	seccionTest["default"]["inicio"] = "principio";
	seccionTest["default"]["prueba espacio"] = "prueba espacio";
	seccionTest["default"]["prueba igual"] = "prueba=igual";
	
	seccionTest["Seccion 2"]["otra seccion"] = "una seccion distinta";
	seccionTest["Seccion 2"]["fin"] = "fin";
//	seccionTest["Seccion 2"]["unomas"] = "uno";
	seccionTest["Seccion 3"]=make(map[string]string);
	seccionTest["Seccion 4"]=make(map[string]string);
	seccionTest["Seccion 4"]["secciondefault"] = "default";
}

func TestReadAll(t *testing.T){
	if err := crearArchivoIni(); err != nil {
		t.Errorf(fmt.Sprint("Error al crear archivo .ini",err));
	}
	inicio := New("inix.ini");
	rellenarSeccionTest();
	if error := inicio.ReadAll();error!=nil {
		fmt.Println(error);
		t.Errorf("No se puede abrir archivo inix.ini");
		t.FailNow();
	}
	//t.Println(inicio.Seccion, " \n", seccionTest);
	if cad,equal := comparemaps(inicio.Seccion,seccionTest);!equal {
		t.Errorf(cad);
	}

		
	
}


func comparemaps(map1,map2 map[string]map[string]string) (string, bool){
	if len(map1) != len(map2) { 
		return fmt.Sprint("Longitud diferente:\n\t",map1,"\n\t",map2),false;
	}
	for k,map1v1 := range map1 {
		if map2v1,ok := map2[k]; ok {
			if len(map1v1) != len(map2v1){
				return fmt.Sprint("Longitud de subelemento diferente:\n\t",map1v1,"\n\t",map2v1),false;
			}			
			for k2,map1v2 := range map1v1 {
				if map2v2,exist := map2v1[k2];exist{
					if map1v2 != map2v2 {
						return fmt.Sprint("Valores direfentes para clave ",k2,":\n",map1v2,"\t",map2v2),false;
					}
				} else {
					return fmt.Sprint("No existe la clave",k2,"en el segundo mapa"),false; 
				}
			}		
		} else { 
			return fmt.Sprint("No existe la clave",k,"en el segundo mapa"),false; 
		}
	}

	return "",true;
}