package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Здесь должна быть справка по использованию утилиты")
		return
	}
	pathProgram := os.Args[0]
	pathSrcFile := os.Args[1]
	pathDstFile := os.Args[2]

        srcFile, err := os.Open(pathSrcFile)
	checkError("Не удалось открыть файл", err)
        defer srcFile.Close()

        dstFile, err := os.Create(pathDstFile)
        checkError("Cannot create file:", err)
        defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
        checkError("Не удалось скопировать файл:", err)
}

func checkError(message string, err error) {
    if err != nil {
	fmt.Println(message, err)
	return
    }
}
