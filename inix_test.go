package inix_test

import (
	. "./inix";
	"fmt";
	"testing";
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
var seccionTest = make(map[string]map[string]string);

func rellenarSeccionTest(){
	seccionTest["default"]=make(map[string]string);
	seccionTest["Seccion 2"]=make(map[string]string);
	seccionTest["default"]["inicio"] = "principio";
	seccionTest["default"]["prueba espacio"] = "prueba espacio";
	seccionTest["default"]["prueba igual"] = "prueba=igual";
	seccionTest["default"]["secciondefault"] = "default";
	seccionTest["Seccion 2"]["otra seccion"] = "una seccion distinta";
	seccionTest["Seccion 2"]["fin"] = "fin";
}

func TestReadAll(t *testing.T){
	inicio := New("inix.ini");		
	if error := inicio.ReadAll();error!=nil {
		fmt.Println(error);
		t.Errorf("No se puede abrir archivo inix.ini");
		t.FailNow();
	}
	if inicio.Seccion != seccionTest {
		fmt.Println(inicio.Seccion, " != ", seccionTest);
		t.Errorf("Datos incorrectos");
	}	
}