package util

func HandleError(e error) {
	if e != nil {
		panic(e)
	}
}
