package main

import (
	"fmt"	
	"path/filepath"
	"strings"
)

func fileName(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext("head_first_izuchaem_go_2020_makgavren.pdf"))]
}

func fileExt()  {
    myfile := "head_first_izuchaem_go_2020_makgavren.pdf"
    extension := strings.LastIndex(myfile, ".")
    if extension == -1 {
       fmt.Println("The file has no extension")
       return
    }
    ext := myfile[extension:]
    fmt.Println("The extension of", myfile, "is", ext) 
    
}

func main() {
	fmt.Println(fileName("head_first_izuchaem_go_2020_makgavren.pdf")) 
	fileExt()
}