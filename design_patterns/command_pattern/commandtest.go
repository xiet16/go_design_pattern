package command_pattern

func CommandTest() {
	middle := NewACommandMiddleware()
	user := NewUser()
	user.SetCommand(middle)
	user.Invoke()
}
