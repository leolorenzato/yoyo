package utils

func AssertErr(err error) {
	if err == nil {
		panic("expected error but got nil")
	}
}
