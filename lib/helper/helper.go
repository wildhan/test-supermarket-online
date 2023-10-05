package helper

func EmptyStringToNull(input string) interface{} {
	if input != "" {
		return input
	} else {
		return nil
	}

}
