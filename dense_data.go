// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"math"
)

//returns a copy of the row (not a slice)
func (A *DenseMatrix) RowCopy(i int) []float64 {
	row := make([]float64, A.cols)
	for j := 0; j < A.cols; j++ {
		row[j] = A.Get(i, j)
	}
	return row
}

//returns a copy of the column (not a slice)
func (A *DenseMatrix) ColCopy(j int) []float64 {
	col := make([]float64, A.rows)
	for i := 0; i < A.rows; i++ {
		col[i] = A.Get(i, j)
	}
	return col
}

//returns a copy of the diagonal (not a slice)
func (A *DenseMatrix) DiagonalCopy() []float64 {
	span := A.rows
	if A.cols < span {
		span = A.cols
	}
	diag := make([]float64, span)
	for i := 0; i < span; i++ {
		diag[i] = A.Get(i, i)
	}
	return diag
}

func (A *DenseMatrix) BufferRow(i int, buf []float64) {
	for j := 0; j < A.cols; j++ {
		buf[j] = A.Get(i, j)
	}
}

func (A *DenseMatrix) BufferCol(j int, buf []float64) {
	for i := 0; i < A.rows; i++ {
		buf[i] = A.Get(i, j)
	}
}

func (A *DenseMatrix) BufferDiagonal(buf []float64) {
	for i := 0; i < A.rows && i < A.cols; i++ {
		buf[i] = A.Get(i, i)
	}
}

func (A *DenseMatrix) FillRow(i int, buf []float64) {
	for j := 0; j < A.cols; j++ {
		A.Set(i, j, buf[j])
	}
}

func (A *DenseMatrix) FillCol(j int, buf []float64) {
	for i := 0; i < A.rows; i++ {
		A.Set(i, j, buf[i])
	}
}

func (A *DenseMatrix) FillDiagonal(buf []float64) {
	for i := 0; i < A.rows && i < A.cols; i++ {
		A.Set(i, i, buf[i])
	}
}

func (A *DenseMatrix) Fill(val float64) {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.Set(i, j, val)
		}
	}
}

//Returns the max value for each row
func (A *DenseMatrix) MaxRows() []float64 {
	result := make([]float64, A.Rows())

	for i := 0; i < A.rows; i++ {
		maxCol := -1 * math.MaxFloat64
		for j := 0; j < A.cols; j++ {
			if A.elements[i*A.step+j] > maxCol {
				maxCol = A.elements[i*A.step+j]
			}
		}
		result[i] = maxCol
	}

	return result
}

//Returns the max value for each column
func (A *DenseMatrix) MaxCols() []float64 {
	result := make([]float64, A.Cols())

	for j := 0; j < A.cols; j++ {
		maxRow := -1 * math.MaxFloat64
		for i := 0; i < A.rows; i++ {
			if A.elements[i*A.step+j] > maxRow {
				maxRow = A.elements[i*A.step+j]
			}
		}
		result[j] = maxRow
	}

	return result
}

//Returns the index of max value for each column
// (in case of duplicate max values first occurrence is used)
func (A *DenseMatrix) ArgMaxCols() []int {
	result := make([]int, A.Cols())

	for j := 0; j < A.cols; j++ {
		maxRow := -1 * math.MaxFloat64
		maxRowIdx := 0
		for i := 0; i < A.rows; i++ {
			if A.elements[i*A.step+j] > maxRow {
				maxRow = A.elements[i*A.step+j]
				maxRowIdx = i
			}
		}
		result[j] = maxRowIdx
	}

	return result
}

// Returns the index of max value for each row
// (in case of duplicate max values first occurrence is used)
func (A *DenseMatrix) ArgMaxRows() []int {
	result := make([]int, A.Rows())

	for i := 0; i < A.rows; i++ {
		maxCol := -1 * math.MaxFloat64
		maxColIdx := 0
		for j := 0; j < A.cols; j++ {
			if A.elements[i*A.step+j] > maxCol {
				maxCol = A.elements[i*A.step+j]
				maxColIdx = j
			}
		}
		result[i] = maxColIdx
	}

	return result
}
