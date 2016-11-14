package numgo

import (
	"fmt"
	"math"
)

func Array(start, end, step float64) []float64 {
	m := make([]float64, int(1+((end-start)/step)))
	for i := 0; i < len(m); i++ {
		m[i] = start + (float64(i) * step)
	}
	return m
}

func ArrayOfInt(start, end int) []int {
	m := make([]int, int(1+((end-start)/1)))
	for i := 0; i < len(m); i++ {
		m[i] = start + (i * 1)
	}
	return m
}

func ArrayOf(number float64, lenght int) []float64 {
	m := make([]float64, lenght)
	for i := 0; i < len(m); i++ {
		m[i] = number
	}
	return m
}

func Mat(length, width int, number float64) [][]float64 {
	m := make([][]float64, length)
	for i := 0; i < length; i++ {
		m[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			m[i][j] = number
		}
	}
	return m
}

func AaddNum(arr []float64, num float64) []float64 {
	m := make([]float64, len(arr))
	for i, v := range arr {
		m[i] = v + num
	}
	return m
}

func AaddA(arr1, arr2 []float64) []float64 {
	m := make([]float64, len(arr1))
	for i, v := range arr1 {
		m[i] = v + arr2[i]
	}
	return m
}

func AmulA(arr1, arr2 []float64) []float64 {
	m := make([]float64, len(arr1))
	for i, v := range arr1 {
		m[i] = v * arr2[i]
	}
	return m
}

func AmulNum(arr []float64, num float64) []float64 {
	m := make([]float64, len(arr))
	for i, v := range arr {
		m[i] = v * num
	}
	return m
}

func AdivNum(arr []float64, num float64) []float64 {
	m := make([]float64, len(arr))
	for i, v := range arr {
		m[i] = v / num
	}
	return m
}

func AdivA(arr1, arr2 []float64) []float64 {
	m := make([]float64, len(arr1))
	for i, v := range arr1 {
		m[i] = v / arr2[i]
	}
	return m
}

func Atm(arr []float64) [][]float64 {
	m := make([][]float64, 1)
	m[0] = arr
	return m
}

func ModIndexes(arrToBeModified, modifierArr []float64, indexArr []int) []float64 {
	for i, v := range indexArr {
		arrToBeModified[v] = modifierArr[i]
	}

	return arrToBeModified
}

func ExpNegArr(arr []float64) []float64 {
	for i, v := range arr {
		arr[i] = math.Exp(-v)
	}

	return arr
}

func ExpArr(arr []float64) []float64 {
	for i, v := range arr {
		arr[i] = math.Exp(v)
	}

	return arr
}

func Copy2dArr(src [][]float64) [][]float64 {
	dest := make([][]float64, len(src))
	for i := range src {
		dest[i] = make([]float64, len(src[i]))
		copy(dest[i], src[i])
	}

	return dest
}

func CopyArr(src []float64) []float64 {
	dest := make([]float64, len(src))
	for i := range src {
		dest[i] = src[i]

	}

	return dest
}

func ArrCompare(arr1, arr2 [][]float64) {

	for i := 0; i < len(arr1); i++ {
		for y := 0; y < len(arr1[i]); y++ {
			if arr1[i][y] != arr2[i][y] {
				fmt.Println("Not equal:", i, y)
			}
		}

	}
}
