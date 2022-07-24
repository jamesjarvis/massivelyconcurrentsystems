package pool

import (
	"math"
	"runtime"
	"time"
)

// Config defines params for Dispatcher.
type Config struct {
	bufferSize int

	batchSize     int
	batchInterval time.Duration

	numConsumers int // TODO(jamesjarvis): Make this a min and max num consumers so it can scale based on backpressure.
}

// SetBufferSize defines the inner queue's buffer size.
func SetBufferSize(bufferSize int) Opt {
	return func(c *Config) {
		c.bufferSize = bufferSize
	}
}

// SetBatchSize defines the max batchsize each worker will operate on.
func SetBatchSize(batchSize int) Opt {
	return func(c *Config) {
		c.batchSize = batchSize
	}
}

// SetBatchInterval defines the max duration each worker will wait for to retrieve items.
func SetBatchInterval(batchInterval time.Duration) Opt {
	return func(c *Config) {
		c.batchInterval = batchInterval
	}
}

// SetNumConsumers defines the number of consumers available to the dispatcher.
func SetNumConsumers(numConsumers int) Opt {
	return func(c *Config) {
		c.numConsumers = numConsumers
	}
}

// Opt constructs a config.
type Opt func(c *Config)

// NewConfig returns a pool Config with defaults that can be overridden.
func NewConfig(opts ...Opt) Config {
	conf := Config{}

	for _, opt := range opts {
		opt(&conf)
	}

	if conf.bufferSize == 0 {
		conf.bufferSize = 1000
	}

	if conf.batchSize == 0 {
		conf.batchSize = 100
	}

	if conf.batchInterval == 0 {
		conf.batchInterval = time.Millisecond
	}

	if conf.numConsumers == 0 {
		numConsumers := math.Max(float64(runtime.NumCPU()-2), 2)
		conf.numConsumers = int(numConsumers)
	}

	return conf
}
