package mws

type Printer interface {
	Printf(format string, args ...interface{})
}
