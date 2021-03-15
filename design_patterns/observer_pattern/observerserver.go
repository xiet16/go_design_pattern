package observer_pattern

func ObserverTest() {
	mom := &Mom{}
	children := &Children{}
	cat := &Cat{}
	mouse := &Mouse{}

	observer := NewObeserver()
	observer.Add("mom", mom)
	observer.Add("children", children)
	observer.Add("cat", cat)
	observer.Add("mouse", mouse)
	observer.Run()
}
