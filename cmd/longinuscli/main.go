package main

import (
	"fmt"
	"longinus/internal/cwe22"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: longinuscli <Java source file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fmt.Printf("Auditing Java file: %s\n", filePath)

	// CWE-22 Java审计逻辑调用
	cwe22.AuditJavaFile(filePath)
}
