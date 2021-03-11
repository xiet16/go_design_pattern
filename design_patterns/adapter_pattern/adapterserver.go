package adapter_pattern

func AdapterTest() {
	adapter := &AdapterPlay{}
	mp3Audio := &AudioPlay{
		AdvancedPlay: adapter,
	}

	mp3Audio.Play("vlc", "vlc file")
}
