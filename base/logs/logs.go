package logs

import (
	"log"
	"os"
)

func CreateLogs() (*os.File, *log.Logger, error){
	// os.0_APPEND append data to the file when writing.
	// os.0_CREATE create a new file if none exists.
	// os.0_WRONLY open the file write-only
	f,err := os.OpenFile("text.log",os.O_APPEND | os.O_CREATE |os.O_WRONLY,0644)
	if err!=nil{
		log.Println(err)
	}
	lg:=setLogger(f, "LOGS:",log.LstdFlags)
	return f,lg,nil
}


func setLogger(f *os.File,pr string, fl int)*log.Logger{
	logger:=log.New(f,pr,fl)
	return logger
}
