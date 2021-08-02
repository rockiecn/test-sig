package utils

func Str2Byte(str string) []byte {
	var ret []byte = []byte(str)
	return ret
}

func Byte2Str(data []byte) string {
	var str string = string(data[:len(data)])
	return str
}
