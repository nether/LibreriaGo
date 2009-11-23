package logx

import "log";
//import "fmt";
import "os";
import "io";
import "runtime";
import "strings";
import "./inifile";

const (	
      TRACE uint8 = iota;	
      DEBUG;
      WARN;
      INFO;
)     

const (
    // Flags
    Lok     = iota;
    Lexit;  // terminate execution when written
    Lcrash; // crash (panic) when written
    // Bits or'ed together to control what's printed. There is no control over the
    // order they appear (the order listed here) or the format they present (as
    // described in the comments).  A colon appears after these items:
    //             2009/0123 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate          = 1 << iota; // the date: 2009/0123
    Ltime;                      // the time: 01:23:23
    Lmicroseconds;              // microsecond resolution: 01:23:23.123123.  assumes Ltime.
    Llongfile;                  // full file name and line number: /a/b/c/d.go:23
    Lshortfile;                 // final file name element and line number: d.go:23. overrides Llongfile
)

var (
    flags = map[string]uint8 {"Ldate":Ldate,"Ltime":Ltime,"Lmicroseconds":Lmicroseconds,"Llongfile":Llongfile,
    "Lshortfile":Lshortfile,"Lok":Lok,"Lexit":Lexit,"Lcrash":Lcrash};
    levels = map[string]uint8 {"DEBUG":DEBUG,"TRACE":TRACE,"WARN":WARN,"INFO":INFO};
)

type Log struct {
     Level uint8;
     *log.Logger;
}

func New(out0 io.Writer,out1 io.Writer, prefix string, flag int,level uint8) *Log{
     return &Log{ level, log.New(out0,out1,prefix,flag)};
}

func NewFrom(file string) *Log{
     
     initial := inifile.New(file);
     initial.ReadAll();

     _,program,_,ok := runtime.Caller(1);
     l := log.New(os.Stdout,nil,"",log.Lshortfile);
     if !ok{       
          l.Log("Cagada");
	  return nil;
     }
     path := strings.Split(program,"/",0);
     index := len(path)-1;
     nombre := path[index];

     var out0,out1 io.Writer;
     var level;
     var flag uint8;

     if v,ok := initial.Seccion["default"][nombre];ok {
     	      variable := strings.ToUpper(v);
	      if v,ok := levels[variable];ok{
	      	 level=v;
		 continue;
	      }
	      if v,ok := flags[variable];ok{
		 flag |= v;
		 continue,
	      }
	      if variable == "FILE1" {

	      	 out0
	      } else if variable == "FILE2" {
	      }
     } else {
       return nil
     }
     l.Log(path[len(path)-1]);
     return nil;
}

func (l *Log) Debug(message string){
     if l.Level <= DEBUG {
          l.Log(message); 
     }
}

func (l Log) Write(b []byte) os.Error{
     return nil;
}