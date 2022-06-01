package pool

import "sync"

// UOW is an implementation of UnitOfWork, with generic request/response type support.
type UOW[REQ, RESP any] struct {
	req  REQ
	resp RESP
	err  error

	wg *sync.WaitGroup
}

// GetRequest returns the request object from the unit of work.
func (uow *UOW[REQ, RESP]) GetRequest() REQ { return uow.req }

// GetResponse returns the response object from the unit of work.
func (uow *UOW[REQ, RESP]) GetResponse() RESP { return uow.resp }

// SetResponse sets the response of the unit of work.
func (uow *UOW[REQ, RESP]) SetResponse(resp RESP) { uow.resp = resp }

// GetError returns an error if an error was encountered while processing the unit of work, otherwise nil.
func (uow *UOW[REQ, RESP]) GetError() error { return uow.err }

// SetError sets the error encountered while processing a unit of work.
func (uow *UOW[REQ, RESP]) SetError(err error) { uow.err = err }

// Done marks the unit of work as processed.
func (uow *UOW[REQ, RESP]) Done() { uow.wg.Done() }

// NewUnitOfWork returns a UnitOfWork instantiated with the given request/response types.
func NewUnitOfWork[REQ, RESP any](req REQ, wg *sync.WaitGroup) *UOW[REQ, RESP] {
	if wg == nil {
		wg = &sync.WaitGroup{}
	}
	return &UOW[REQ, RESP]{
		req: req,
		wg:  wg,
	}
}
