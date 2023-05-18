package algo_data

import (
	"github.com/chon42/goal/goalbase"
)

type AlgoData struct {
	IAlgoData
}

type IAlgoData interface {
	goalbase.IPrintf
}
