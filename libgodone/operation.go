package libgodone

type Operation int

const (
	HELP    Operation = iota
	VERSION Operation = iota
	NEW     Operation = iota
	INIT    Operation = iota
	RUN     Operation = iota
	BUILD   Operation = iota
	CLEAN   Operation = iota
)
