package newXsql

type ISqlTrans interface {
	SqlTransGoStructs()
	SqlTransProtoMessages()
	SqlTableAddColumns()
}
