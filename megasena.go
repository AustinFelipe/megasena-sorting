package main

import (
	"fmt"
	"sort"

	"github.com/tealeg/xlsx"
)

type SenaNumber struct {
	Number int
	Count  int
}

func main() {
	fmt.Printf("Iniciando cálculo\n")
	excelFileName := "C:/Users/austi/go/src/megasena/mega_sena_ate_concurso_2175.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Printf("Erro ao carregar\n%v", err)
	}

	numbers := []SenaNumber{}

	for i := 0; i < 60; i++ {
		numbers = append(numbers, SenaNumber{Number: i, Count: 0})
	}

	// Sum numbers
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				number, _ := cell.Int()

				numbers[number-1].Count++
			}
		}
	}

	// Sort numbers
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i].Count < numbers[j].Count
	})

	fmt.Printf("Maiores números\n")

	for i := 59; i >= 50; i-- {
		fmt.Printf("Número: %d - Sorteado %d vezes\n", numbers[i].Number, numbers[i].Count)
	}
}
