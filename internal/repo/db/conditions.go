package db

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Where conditions
type Condition func() (sql string, args []interface{}, err error)

/*
	 Can be used with:
		ListInstruments
		ListPresets
*/
func ByKitId(kit int64) Condition {
	return Eq("kit", kit)
}

func ById(id int64) Condition {
	return Eq("id", id)
}

func ByUid(uuid string) Condition {
	return Eq("uid", uuid)
}

func Eq(field string, inargs ...interface{}) Condition {
	return func() (sql string, args []interface{}, err error) {
		return fmt.Sprintf("%s = ?", field), inargs, nil
	}
}

func In(field string, inargs ...interface{}) Condition {
	return func() (sql string, args []interface{}, err error) {
		return sqlx.In(fmt.Sprintf("%s in (?)", field), inargs)
	}
}

func buildConditions(conds ...Condition) (sql string, args []interface{}, err error) {
	sqls := make([]string, len(conds))
	for i, cond := range conds {
		s, a, err := cond()
		if err != nil {
			return "", nil, fmt.Errorf("failed construct sql condition: %w", err)
		}
		sqls[i] = s
		args = append(args, a...)
	}
	if len(sqls) > 0 {
		sql = "where " + strings.Join(sqls, " and ")
	}
	return
}
