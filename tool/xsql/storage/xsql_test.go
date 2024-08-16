package storage

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestSqlTransGoStruct(t *testing.T) {
	//SqlTransGoStruct()
	//SqlTransProtoMessage()
	//TableAddColumn("./question_bank.sql")

	db, err := sql.Open("mysql", "liubei:gsmfts43mgsGSf@tcp(bj-cdb-jxn3gh0g.sql.tencentcdb.com:63801)/learn?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	query := " SELECT '' as paper_submit_id,qp.question_paper_id,'' as question_student_id,0 as score,DATETIME as create_time,'' as paper_answer_name, qp.paper_name FROM question_paper qp LEFT JOIN question_paper_submit qps ON qp.question_paper_id = qps.question_paper_id WHERE qps.id IS NULL and qp.create_paper_uid='4vd2rk0k-queteauid' limit 0,20"
	DbRes, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	typesList, _ := DbRes.ColumnTypes()
	for _, v := range typesList {
		fmt.Println(v)
	}
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

func TestTableAddColumn(t *testing.T) {
	TableAddColumn2("./question_bank.sql")
}
