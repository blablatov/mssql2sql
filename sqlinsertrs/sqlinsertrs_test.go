package sqlinsertrs

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		insertIntegrTableSql string
	}{
		{"    :, , :"},
		{"wertyuittt_90/,_=="},
		{"__+_8_0000//,_+=;;;.?/-@"},
		{"QQVNMBVFRERTYuiop,asdfghU56789543211IJNM<>L><M???_"},
	}

	var previnsertIntegrTableSql string
	for _, test := range tests {
		if test.insertIntegrTableSql != previnsertIntegrTableSql {
			fmt.Printf("\n%s\n", test.insertIntegrTableSql)
			previnsertIntegrTableSql = test.insertIntegrTableSql
		}
	}
}
