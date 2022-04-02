package logs

import (
	"errors"
	"testing"
)

func Test_Logs(t *testing.T) {
	Warn("warn")
	Warn(errors.New("error warn"))
	Warnf("%#v %v", Config{}, []string{"1", "2"})
	Warnf("%v %v", 10, []string{"1", "2"})
	Warnf("%+v %v", 10, []string{"1", "2"})
	Warnf("%#v %v", 10, []string{"1", "2"})

	rf := RequestField{XRequestId: "123456"}
	rf.Warn("warn")
	rf.Warn(errors.New("error warn"))
	rf.Warnf("%#v %v", Config{}, []string{"1", "2"})
}
