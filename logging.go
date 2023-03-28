package eos

import (
	"fmt"

	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

var zlog *zap.Logger
var tracer = &Tracer{}

func init() {
	logging.Register("github.com/eoscanada/eos-go", &zlog)
}

type Tracer struct{}

func (t *Tracer) Enabled() bool {
	return logging.IsTraceEnabled("eos-go", "github.com/eoscanada/eos-go")
}

type logStringerFunc func() string

func (f logStringerFunc) String() string { return f() }

func typeField(field string, v interface{}) zap.Field {
	return zap.Stringer(field, logStringerFunc(func() string {
		return fmt.Sprintf("%T", v)
	}))
}
