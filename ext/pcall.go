package ext

func PCall(f func()) {
	defer func() {
		if err := recover(); err != nil {
			LogError(err.(error))
		}
	}()

	f()
}
