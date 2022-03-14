package auth

import (
	"encoding/json"
	"errors"
)

func ParseTokenInfo(tokenInfo string) (Identify, error) {
	if tokenInfo == "" {
		return Identify{}, errors.New("tokenInfo 为空")
	}

	identify := Identify{}
	err := json.Unmarshal([]byte(tokenInfo), &identify)
	if err != nil {
		return Identify{}, errors.New("tokenInfo 解析错误：" + err.Error())
	}

	return identify, nil
}
