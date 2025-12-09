package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	// "slices"
	// "strconv"
	"strings"
)

func main() {
	file := "06.input.txt"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	var operatorsWithPaddings []string
	operatorRegex := regexp.MustCompile(` ([+\*])`)
	fmt.Printf("theLine:\n\"%v\"\n", lines[len(lines)-1])
	fmt.Printf("theReplacedLine:\n\"%v\"\n", string(operatorRegex.ReplaceAll([]byte(string(lines[len(lines)-1])), []byte("|$1"))))
	operatorsWithPaddings = strings.Split(
		string(operatorRegex.ReplaceAll([]byte(string(lines[len(lines)-1])), []byte("|$1"))),
		"|",
	)

	fmt.Println("operatorsWithPaddings: ", operatorsWithPaddings)

	var operands [][]string
	for _, line := range lines[:len(lines)-1] {
		// operandRegex := regexp.MustCompile(` (\d)`)
		// operandColumns := strings.Split(
		// 	string(operandRegex.ReplaceAll([]byte(line), []byte("|$1"))),
		// 	"|",
		// )

		var operandColumns []string
		marker := 0
		for _, operatorWithPadding := range operatorsWithPaddings {
			operandColumns = append(operandColumns, line[marker:marker+len(operatorWithPadding)])
			marker += len(operatorWithPadding) + 1
		}

		operands = append(operands, operandColumns)
	}

	operators := strings.Fields(lines[len(lines)-1])
	fmt.Println("operands: ", operands)
	fmt.Println("operators: ", operators)

	grandTotal := 0
	for columnIndex, operator := range operators {
		operandTotal := 0
		columnWidth := len(operands[0][columnIndex])
		fmt.Printf("the columnWidth:\n\"%v\"\n\n", columnWidth)
		for numberColumnIndex := range columnWidth {

			columnNumberString := ""

			for _, operandLine := range operands {

				// fmt.Printf("operandlinezero:\n\"%v\"\n\n", operandLine[0])
				// fmt.Printf("operandlinefull:\n\"%v\" type: %T len: %v\n\n", operandLine, operandLine, len(operandLine))

				// for _, theStrings := range operandLine {
				// 	fmt.Printf("theStrings:\n\"%v\"\n\n", theStrings)
				// }

				// fmt.Printf("operandLine:\n\"%v\"\n\ncolumnIndex:\n\"%v\"\n\nnumberColumnIndex:\n\"%v\"\n\noperandLine[columnIndex]:\n\"%v\"\n\n", operandLine, columnIndex, numberColumnIndex, operandLine[columnIndex])
				columnNumberString += string(operandLine[columnIndex][numberColumnIndex])
				// fmt.Printf("current columnNumberString: \"%v\"\n", columnNumberString)
			}

			columnNumber, err := strconv.Atoi(strings.TrimSpace(string(columnNumberString)))
			if err != nil {
				fmt.Println("Error converting columns number string to integer", err)
				return
			}

			switch operator {
			case "+":
				operandTotal += columnNumber
				// fmt.Printf(" + %v", columnNumber)
			case "*":
				if operandTotal == 0 {
					operandTotal = 1
				}
				operandTotal *= columnNumber
				// fmt.Printf(" * %v", columnNumber)
			}

		}

		// fmt.Printf("\nOperand total: %v\n\n", operandTotal)

		// fmt.Println("columnIndex: ", columnIndex)
		// fmt.Println("operands[1:]: ", operands[1:])

		grandTotal += operandTotal
	}

	fmt.Println("The grand total is: ", grandTotal)
}

/*
123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
*/
