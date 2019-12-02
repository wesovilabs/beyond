package helper

// CheckError checkec error an returns panic if it exists
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
