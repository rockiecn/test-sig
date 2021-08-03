package utils

func Str2Byte(str string) []byte {
	var ret []byte = []byte(str)
	return ret
}

func Byte2Str(data []byte) string {
	var str string = string(data[:len(data)])
	return str
}

func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}
