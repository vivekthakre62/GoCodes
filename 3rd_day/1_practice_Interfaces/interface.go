package main

import "fmt"

type logger interface {
	info(message string)
	error(message string)
}

type consolLogger struct {
}

func (console consolLogger) info(message string) {
	fmt.Println("info: ", message)
}
func (console consolLogger) error(message string) {
	fmt.Println("Error: ", message)
}

type fileLogger struct {
	fileName string
}

func (file fileLogger) info(message string) {
	fmt.Println("info ", "About ", file.fileName, message)
}
func (file fileLogger) error(message string) {
	fmt.Println("Error ", "In ", file.fileName, message)
}

type DBLogger struct {
	dbName string
}

func (db DBLogger) info(message string) {
	fmt.Println("info about the", db.dbName, " db", message)
}
func (db DBLogger) error(message string) {
	fmt.Println("Error in ", db.dbName, " db", message)
}
func showMessage(l logger) {
	l.info("System start")
	l.error("something went wrong")
}

func main() {
	c := consolLogger{}
	f := fileLogger{
		fileName: "vivekthakre.pdf",
	}
	db := DBLogger{
		dbName: "VivekThakre",
	}
	showMessage(c)
	fmt.Println("________________________")
	showMessage(f)
	fmt.Println("_________________________")
	showMessage(db)

}
