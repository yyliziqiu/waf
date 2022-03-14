package mysqldb

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
)

/**
判断 err 是否为 记录不存在
*/
func IsNoRowsError(err error) bool {
	return err == sql.ErrNoRows
}

/**
判断 err 是否为 索引冲突
*/
func IsDuplicateEntryError(err error) bool {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return false
	}
	return mysqlErr.Number == 1062
}

/**
防止 sql 注入
*/
func AddSlashes(str string) string {
	chars := []rune(str)
	temp := make([]rune, 0, len(chars))
	for _, c := range chars {
		if c == '\\' || c == '"' || c == '\'' {
			temp = append(temp, '\\')
			temp = append(temp, c)
		} else {
			temp = append(temp, c)
		}
	}
	return string(temp)
}

/**
拼接 IN 条件
*/
func InConditionStringAddSlashes(items []string) string {
	length := len(items)
	for i := 0; i < length; i++ {
		items[i] = AddSlashes(items[i])
	}
	return InConditionString(items)
}

/**
拼接 IN 条件
*/
func InConditionString(items []string) string {
	length := len(items)
	if length == 0 {
		return "()"
	}
	return "('" + strings.Join(items, "', '") + "')"
}

/**
拼接 IN 条件
*/
func InConditionInt(items []int) string {
	length := len(items)
	if length == 0 {
		return "()"
	}
	sb := strings.Builder{}
	sb.WriteString("(")
	for i := 0; i < length-1; i++ {
		sb.WriteString(strconv.Itoa(items[i]))
		sb.WriteString(", ")
	}
	sb.WriteString(strconv.Itoa(items[length-1]))
	sb.WriteString(")")
	return sb.String()
}
