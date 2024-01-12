package xsql

import (
	"testing"
)

func TestSqlTransGoStruct(t *testing.T) {
	SqlTransGoStruct()
	SqlTransProtoMessage()

}

func hasIntersection(dataRanges [][]int) bool {
	if len(dataRanges) < 2 {
		return false
	}

	preMax := dataRanges[0][1]
	for i := 1; i < len(dataRanges); i++ {
		if preMax > dataRanges[i][0] {
			return true
		}
		preMax = dataRanges[i][1]
	}

	return false

}
