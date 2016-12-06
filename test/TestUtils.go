package test

func PanicOnError(err error) {

	if err != nil {
		panic(err)
	}

}
