package src

import (
	"errors"
	"leetCodeGo/src/basic"
	"sync"
)

var BLOCK_SIZE = 16

func initMatrixA(input [][]float32) []float32 {
	matrix := make([]float32, len(input)*len(input[0]))
	row := len(input)
	col := len(input[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			index := i*col + j
			matrix[index] = input[i][j]
		}
	}
	return matrix
}

func initMatrixB(input [][]float32) []float32 {
	matrix := make([]float32, len(input)*len(input[0]))
	row := len(input)
	col := len(input[0])
	for j := 0; j < col; j++ {
		for i := 0; i < row; i++ {
			index := j*row + i
			matrix[index] = input[i][j]
		}
	}
	return matrix
}

func multiplyMatrix(matrixA [][]float32, matrixB [][]float32) ([]float32, error) {
	if len(matrixA[0]) != len(matrixB) {
		return []float32{}, errors.New("乘数矩阵的列和被乘矩阵的行数不同，无法相乘")
	}

	rowA := len(matrixA)
	colA := len(matrixA[0])
	colB := len(matrixB[0])
	arrC := make([]float32, colA*colB)
	wg := sync.WaitGroup{}

	arrA := initMatrixA(matrixA)
	arrB := initMatrixB(matrixB)

	for i := 0; i < rowA; i += BLOCK_SIZE {
		go func(startRow int) {
			defer wg.Done()
			for ii := startRow; ii < basic.Min(startRow+BLOCK_SIZE, rowA); ii++ {
				for j := 0; j < colB; j += BLOCK_SIZE {
					for m := 0; m < colB; m += BLOCK_SIZE {
						for jj := j; jj < basic.Min(j+BLOCK_SIZE, colB); jj++ {
							var dot float32 = 0.0
							for mm := m; mm < basic.Min(m+BLOCK_SIZE, colA); mm++ {
								indexA := ii*colA + mm
								indexB := jj*colA + mm
								dot += arrA[indexA] * arrB[indexB]
								indexC := ii*colB + jj
								arrC[indexC] = dot
							}
						}
					}
				}
			}
		}(i)
	}
	return arrC, nil
}
