package decorator_pattern

func DecoratorTest() {

	jacket := &DecoratorComponent{
		Name:  "jacket",
		Entry: &DecoratorEntry{},
	}

	suit := &DecoratorComponent{
		Name:  "suit",
		Entry: &DecoratorEntry{},
	}

	people := &DecoratorComponent{
		Name: "xiet",
		md:   []*DecoratorEntry{},
	}

	people.Add(jacket.Name, jacket.Entry)
	people.Add(suit.Name, suit.Entry)
	people.Show()
}
