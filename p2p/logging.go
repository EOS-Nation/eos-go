package p2p

import (
	"fmt"
	"go.uber.org/zap"

	"github.com/streamingfast/logging"
)

var zlog *zap.Logger
var tracer = &Tracer{}

func init() {
	logging.Register("github.com/eoscanada/eos-go/p2p", &zlog)
}

type Tracer struct{}

func (t *Tracer) Enabled() bool {
	return logging.IsTraceEnabled("eos-go", "github.com/eoscanada/eos-go/p2p")
}

// SyncLogger sync logger, should `defer SyncLogger()` when use p2p package
func SyncLogger() {
	err := zlog.Sync()
	if err != nil {
		fmt.Printf("unable to sync p2p logger: %s\n", err)
	}
}
