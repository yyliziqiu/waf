package util

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/**
IP 转整型数字
*/
func IPToInt(ip string) (uint32, error) {
	var (
		n uint32 = 0
		s        = fmt.Sprintf("%s IP 格式错误", ip)
	)

	segments := strings.Split(strings.TrimSpace(ip), ".")
	if len(segments) != 4 {
		return 0, errors.New(s)
	}

	for _, segment := range segments {
		i, err := strconv.Atoi(segment)
		if err != nil {
			return 0, errors.New(s)
		}
		if i < 0 || i > 255 {
			return 0, errors.New(s)
		}
		n = n<<8 | uint32(i)
	}

	return n, nil
}

/**
匹配 IP
*/
func IPMatch(ip, p string) (bool, error) {
	if strings.Index(p, "/") == -1 {
		return ip == p, nil
	}

	var (
		m uint32 = 0
		s        = fmt.Sprintf("%s 匹配格式错误", p)
	)

	segments := strings.Split(strings.TrimSpace(p), "/")
	if len(segments) != 2 {
		return false, errors.New(s)
	}

	b, err := strconv.Atoi(segments[1])
	if err != nil {
		return false, errors.New(s)
	}
	if b < 0 || b > 32 {
		return false, errors.New(s)
	}
	m = (1<<b - 1) << (32 - b)

	i1, err := IPToInt(ip)
	if err != nil {
		return false, err
	}
	i2, err := IPToInt(segments[0])
	if err != nil {
		return false, err
	}

	return i1&m == i2&m, nil
}

/**
匹配 IP 列表。列表中有一个 IP 匹配则返回 true
*/
func IPMMatch(ip string, ps []string) (bool, error) {
	for _, p := range ps {
		if match, err := IPMatch(ip, p); err != nil {
			return false, err
		} else if match {
			return true, nil
		}
	}
	return false, nil
}
