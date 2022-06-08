package main

import (
	"fmt"
	"os"
	"sync"
)

// The maximum arguments for
const maxArgs = 10

var tupleNames map[int]string
var tupleTypeNames map[int]string
var tupleArgNames map[int]string

func writeGenerateFileHeader(file *os.File) {
	fmt.Fprintln(file, "package ranges")
	fmt.Fprintln(file)
	fmt.Fprintln(file, "// Generated with: go run internal/cmd/generatecode/main.go")
}

func printTuple(file *os.File, size int) {
	fmt.Fprintf(file, "// %s holds %d items\n", tupleNames[size], size)
	fmt.Fprintf(file, "type %s[", tupleNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s any] struct {\n", tupleTypeNames[size])

	for i := 1; i <= size; i++ {
		fmt.Fprintf(file, "\t%s %s\n", tupleTypeNames[i], tupleTypeNames[i])
	}

	fmt.Fprintln(file, "}")
	fmt.Fprintln(file)

	fmt.Fprintf(file, "// Make%s creates a %s\n", tupleNames[size], tupleNames[size])
	fmt.Fprintf(file, "func Make%s[", tupleNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s any](", tupleTypeNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s %s, ", tupleArgNames[i], tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s %s) %s[", tupleArgNames[size], tupleTypeNames[size], tupleNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s] {\n\treturn %s[", tupleTypeNames[size], tupleNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s]{", tupleTypeNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleArgNames[i])
	}

	fmt.Fprintf(file, "%s}\n", tupleArgNames[size])
	fmt.Fprintln(file, "}")
}

func printZipStruct(file *os.File, size int) {
	fmt.Fprintf(file, "type zip%sResult[", tupleNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s any] struct {\n", tupleTypeNames[size])

	for i := 1; i <= size; i++ {
		fmt.Fprintf(file, "\t%s InputRange[%s]\n", tupleArgNames[i], tupleTypeNames[i])
	}

	fmt.Fprintln(file, "}")
}

func printZipMethodHead(file *os.File, size int) {
	fmt.Fprintf(file, "func (z *zip%sResult[", tupleNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s]) ", tupleTypeNames[size])
}

func printZipEmptyMethod(file *os.File, size int) {
	printZipMethodHead(file, size)

	fmt.Fprintln(file, "Empty() bool {")
	fmt.Fprintln(file, "\treturn z.a.Empty() ||")

	for i := 2; i <= size; i++ {
		fmt.Fprintf(file, "\t\tz.%s.Empty()", tupleArgNames[i])

		if i < size {
			fmt.Fprintf(file, " ||")
		}

		fmt.Fprintln(file)
	}

	fmt.Fprintln(file, "}")
}

func printZipFrontMethod(file *os.File, size int) {
	printZipMethodHead(file, size)

	fmt.Fprintf(file, "Front() %s[", tupleNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s] {\n", tupleTypeNames[size])
	fmt.Fprintf(file, "\treturn %s[", tupleNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s]{\n", tupleTypeNames[size])

	for i := 1; i <= size; i++ {
		fmt.Fprintf(file, "\t\tz.%s.Front(),\n", tupleArgNames[i])
	}

	fmt.Fprintln(file, "\t}")
	fmt.Fprintln(file, "}")
}

func printZipPopFrontMethod(file *os.File, size int) {
	printZipMethodHead(file, size)

	fmt.Fprintln(file, "PopFront() {")

	for i := 1; i <= size; i++ {
		fmt.Fprintf(file, "\tz.%s.PopFront()\n", tupleArgNames[i])
	}

	fmt.Fprintln(file, "}")
}

func printZipFunc(file *os.File, size int) {
	fmt.Fprintf(file, "// Zip%d produces items from %d ranges in parallel.\n", size, size)
	fmt.Fprintln(file, "// The range will be empty when any of the ranges are empty.")
	fmt.Fprintf(file, "func Zip%d[", size)

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s any](\n", tupleTypeNames[size])

	for i := 1; i <= size; i++ {
		fmt.Fprintf(file, "\t%s InputRange[%s],\n", tupleArgNames[i], tupleTypeNames[i])
	}

	fmt.Fprintf(file, ") InputRange[%s[", tupleNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s]] {\n", tupleTypeNames[size])
	fmt.Fprintf(file, "\treturn &zip%sResult[", tupleNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s]{", tupleTypeNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleArgNames[i])
	}

	fmt.Fprintf(file, "%s}\n", tupleArgNames[size])
	fmt.Fprintln(file, "}")
}

func writeTupleFile(file *os.File) {
	writeGenerateFileHeader(file)

	for size := 2; size <= maxArgs; size++ {
		fmt.Fprintln(file)
		printTuple(file, size)
	}
}

func writeZipFile(file *os.File) {
	writeGenerateFileHeader(file)

	for size := 2; size <= maxArgs; size++ {
		fmt.Fprintln(file)
		printZipStruct(file, size)
		fmt.Fprintln(file)
		printZipEmptyMethod(file, size)
		fmt.Fprintln(file)
		printZipFrontMethod(file, size)
		fmt.Fprintln(file)
		printZipPopFrontMethod(file, size)
		fmt.Fprintln(file)
		printZipFunc(file, size)
	}
}

func createFile(wg *sync.WaitGroup, path string, cb func(*os.File)) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		file, err := os.Create(path)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open %s \n", path)
			os.Exit(1)
		}

		defer file.Close()

		cb(file)
	}()
}

func main() {
	tupleNames = map[int]string{
		1:  "Unit",
		2:  "Pair",
		3:  "Triplet",
		4:  "Quartet",
		5:  "Quintet",
		6:  "Sextet",
		7:  "Septet",
		8:  "Octet",
		9:  "Ennead",
		10: "Decade",
	}
	tupleTypeNames = map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E", 6: "F", 7: "G", 8: "H", 9: "I", 10: "J"}
	tupleArgNames = map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f", 7: "g", 8: "h", 9: "i", 10: "j"}

	var wg sync.WaitGroup

	go createFile(&wg, "pkg/ranges/tuple.go", writeTupleFile)
	go createFile(&wg, "pkg/ranges/zip.go", writeZipFile)

	wg.Wait()
}
