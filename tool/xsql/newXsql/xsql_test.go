package newXsql

import (
	"github.com/GnezIew/goshell/tool/xsql/newXsql/utils"
	"testing"
)

func TestNewSqlTrans(t *testing.T) {
	s := NewSqlTrans(utils.GlobalObject.FilePath, utils.GlobalObject.DataSourceName)

	//s.SqlTransGoStructs()
	//s.SqlTransProtoMessages()
	s.SqlTableAddColumns()
}
