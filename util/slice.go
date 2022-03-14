package util

/**
检查切片中是否包含某个字符串
*/
func ContainsString(slice []string, target string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

/**
检查切片中是否包含某个数字
*/
func ContainsInt(slice []int, target int) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}
