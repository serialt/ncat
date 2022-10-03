package pkg

import "os"

func Env(key, def string) string {
	if x := os.Getenv(key); x != "" {
		return x
	}
	return def
}

// 判断文件目录否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
