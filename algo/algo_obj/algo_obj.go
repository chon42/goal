package algo_obj

import "fmt"

type AlgoObj struct {
	IAlgoObj
}
type IAlgoObj interface {
	Printf(fmt string, v ...any)
}

func (AlgoObj) Printf(format string, v ...any) (int, error) {
	return fmt.Printf(format, v...)
}
