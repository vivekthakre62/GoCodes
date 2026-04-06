package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// f, err := os.Open("demo.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()
	// fmt.Println("file Name:", fileInfo.Name())
	// fmt.Println("file directory is folder:", fileInfo.IsDir())
	// fmt.Println("Last update of file:", fileInfo.ModTime())
	// fmt.Println("File Size:", fileInfo.Size())
	//Read file*******

	// f, err := os.Open("demo.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// fileInfo, err := f.Stat()

	// buf := make([]byte, fileInfo.Size())
	// d, er := f.Read(buf)
	// if er != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	fmt.Println(d, string(buf[i]))
	// }
	//easy way to read file
	d, err := os.ReadFile("demo.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))

	//for read the folder
	dir, error := os.Open(".")
	if error != nil {
		panic(error)
	}
	defer dir.Close()
	fileInfo, err := dir.ReadDir(2)
	for i := 0; i < len(fileInfo); i++ {
		fmt.Println(fileInfo[i].Name())
	}

	//Create the file
	f1, err1 := os.Create("demo2.txt")
	if err1 != nil {
		panic(err1)
	}
	f1.WriteString("Hey It's ")
	f1.WriteString("Vivek!")

	//file data transfer by using stream style

	sourceFile, errs := os.Open("demo.txt")
	if errs != nil {
		panic(errs)
	}
	defer sourceFile.Close()

	destFile, errd := os.Create("demo3.txt")
	if errd != nil {
		panic(errd)
	}
	defer destFile.Close()
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		r, eReader := reader.ReadByte()
		if eReader != nil {
			if eReader.Error() != "EOF" {
				panic(eReader)
			}
			break
		}
		eWriter := writer.WriteByte(r)
		if eWriter != nil {
			panic(eWriter)
		}
	}
	writer.Flush()
	fmt.Println("Succesully written to file")

	//delete file
	eD := os.Remove("filSample.txt")
	if eD != nil {
		panic(eD)
	}

}
