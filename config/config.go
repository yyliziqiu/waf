package config

import (
	"reflect"

	"github.com/spf13/viper"

	"github.com/yyliziqiu/waf/ylog"
)

/**
加载配置
*/
func Initialize(path string, c interface{}) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		ylog.FatalE(err)
	}
	if err := viper.Unmarshal(c); err != nil {
		ylog.FatalE(err)
	}
}

/**
创建配置索引
*/
func CreateIndex(c interface{}, i interface{}) {
	cType := reflect.TypeOf(c)
	cValue := reflect.ValueOf(c)
	iType := reflect.TypeOf(i).Elem()
	iValue := reflect.ValueOf(i).Elem()

	// 获取需要索引的 field
	indexKeys := make([]string, 0, iType.NumField())
	for i := 0; i < iType.NumField(); i++ {
		indexKeys = append(indexKeys, iType.Field(i).Name)
	}

	// 建立索引
	for _, key := range indexKeys {
		if _, ok := cType.FieldByName(key); !ok {
			ylog.Fatal("Config field {" + key + "} does not exist")
		}
		cFieldValue := cValue.FieldByName(key)
		if cFieldValue.Kind() != reflect.Slice {
			ylog.Fatal("Config field {" + key + "} is not slice")
		}
		iField, _ := iType.FieldByName(key)
		mapValue := reflect.MakeMapWithSize(iField.Type, cFieldValue.Len())
		for i := 0; i < cFieldValue.Len(); i++ {
			vValue := cFieldValue.Index(i)
			vNameValue := vValue.FieldByName("Name")
			if vNameValue.IsZero() {
				ylog.Fatal(key + " Config field {Name} empty")
			}
			mapValue.SetMapIndex(vNameValue, vValue)
		}
		iValue.FieldByName(key).Set(mapValue)
	}
}
