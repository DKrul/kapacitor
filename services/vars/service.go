package vars

import (
	"fmt"
	"log"
	"sync"

	"github.com/influxdata/kapacitor/tick/stateful"
	"github.com/pkg/errors"
)

type Service struct {
	mu     sync.Mutex
	vars   Config
	logger *log.Logger

	context *stateful.FunctionContext
}

func NewService(c Config, l *log.Logger) (*Service, error) {
	if c == nil {
		return nil, errors.New("must pass non nil config")
	}
	s := &Service{
		vars:    c,
		logger:  l,
		context: stateful.NewFunctionContext(),
	}
	return s, nil
}

func (s *Service) Open() error {
	// Register the "config" function
	err := s.context.Register("config", func() stateful.Func { return s })
	return errors.Wrap(err, "registering config function")
}
func (s *Service) Close() error {
	return nil
}

func (s *Service) FunctionContext() *stateful.FunctionContext {
	return s.context
}

func (s *Service) Get(key string) string {
	s.mu.Lock()
	str := s.vars[key]
	s.mu.Unlock()
	return str
}

func (s *Service) Reset() {}

func (s *Service) Call(args ...interface{}) (interface{}, error) {
	if len(args) != 1 {
		return nil, errors.New("config expects exactly one argument")
	}
	key, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("cannot pass %T to config, must be string", args[0])
	}
	return s.Get(key), nil
}
