package split

import "strings"

// Split 将s按照sep进行分割，返回一个字符串切片
func Split(s, sep string) (ret []string) {
	//idx := strings.Index(s, sep)
	//for idx > -1 {
	//	ret = append(ret, s[:idx])
	//	s = s[idx+len(sep):]
	//	idx = strings.Index(s, sep)
	//}
	//ret = append(ret, s)
	//return

	// 优化后,
	ret = make([]string, 0, strings.Count(s, sep)+1)
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
