.PHONY: play
play:
	go run golang.org/x/tools/cmd/present -base './' -content './slides' -orighost 'localhost'
