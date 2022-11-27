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

func printZipStruct(file *os.File, size int, suffix string, rangeType string) {
	fmt.Fprintf(file, "type zip%s%sResult[", tupleNames[size], suffix)

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s any] struct {\n", tupleTypeNames[size])

	for i := 1; i <= size; i++ {
		fmt.Fprintf(file, "\t%s %s[%s]\n", tupleArgNames[i], rangeType, tupleTypeNames[i])
	}

	fmt.Fprintln(file, "}")
}

func printZipMethodHead(file *os.File, size int, suffix string) {
	fmt.Fprintf(file, "func (z *zip%s%sResult[", tupleNames[size], suffix)

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s]) ", tupleTypeNames[size])
}

func printZipEmptyMethod(file *os.File, size int, suffix string) {
	printZipMethodHead(file, size, suffix)

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

func printZipFrontMethod(file *os.File, size int, suffix string) {
	printZipMethodHead(file, size, suffix)

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

func printZipPopFrontMethod(file *os.File, size int, suffix string) {
	printZipMethodHead(file, size, suffix)

	fmt.Fprintln(file, "PopFront() {")

	for i := 1; i <= size; i++ {
		fmt.Fprintf(file, "\tz.%s.PopFront()\n", tupleArgNames[i])
	}

	fmt.Fprintln(file, "}")
}

func printZipSaveMethod(file *os.File, size int, suffix string) {
	printZipMethodHead(file, size, suffix)

	fmt.Fprintf(file, "Save(")
	printZipReturnType(file, size, suffix, "ForwardRange")
	fmt.Fprintf(file, "\treturn &zip%s%sResult[", tupleNames[size], suffix)

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s]{\n", tupleTypeNames[size])

	for i := 1; i <= size; i++ {
		fmt.Fprintf(file, "\t\tz.%s.Save(),\n", tupleArgNames[i])
	}

	fmt.Fprintln(file, "\t}")
	fmt.Fprintln(file, "}")
}

func printZipReturnType(file *os.File, size int, suffix string, rangeType string) {
	fmt.Fprintf(file, ") %s[%s[", rangeType, tupleNames[size])

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s]] {\n", tupleTypeNames[size])
}

func printZipFunc(file *os.File, size int, suffix string, rangeType string) {
	fmt.Fprintf(file, "// Zip%d%s produces items from %d ranges in parallel.\n", size, suffix, size)
	fmt.Fprintln(file, "// The range will be empty when any of the ranges are empty.")
	fmt.Fprintf(file, "func Zip%d%s[", size, suffix)

	for i := 1; i < size; i++ {
		fmt.Fprintf(file, "%s, ", tupleTypeNames[i])
	}

	fmt.Fprintf(file, "%s any](\n", tupleTypeNames[size])

	for i := 1; i <= size; i++ {
		fmt.Fprintf(file, "\t%s %s[%s],\n", tupleArgNames[i], rangeType, tupleTypeNames[i])
	}

	printZipReturnType(file, size, suffix, rangeType)
	fmt.Fprintf(file, "\treturn &zip%s%sResult[", tupleNames[size], suffix)

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

type zipType struct {
	suffix    string
	rangeType string
}

func writeZipFile(file *os.File) {
	writeGenerateFileHeader(file)

	for size := 2; size <= maxArgs; size++ {
		for _, item := range []zipType{{"", "InputRange"}, {"F", "ForwardRange"}} {
			fmt.Fprintln(file)
			printZipStruct(file, size, item.suffix, item.rangeType)
			fmt.Fprintln(file)
			printZipEmptyMethod(file, size, item.suffix)
			fmt.Fprintln(file)
			printZipFrontMethod(file, size, item.suffix)
			fmt.Fprintln(file)
			printZipPopFrontMethod(file, size, item.suffix)

			if item.suffix == "F" {
				fmt.Fprintln(file)
				printZipSaveMethod(file, size, item.suffix)
			}

			fmt.Fprintln(file)
			printZipFunc(file, size, item.suffix, item.rangeType)
		}
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

	createFile(&wg, "ranges/tuple.go", writeTupleFile)
	createFile(&wg, "ranges/zip.go", writeZipFile)

	wg.Wait()
}
