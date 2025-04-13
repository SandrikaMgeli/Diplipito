package network

import "sync"

type LocalTransport struct {
	addr      NetAddr
	lock      sync.RWMutex
	consumeCh chan RPC
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}

func (t *LocalTransport) SendMessage(netAddr NetAddr, message []byte) error {
	return nil
}

func (t *LocalTransport) Connect(transport Transport) error {
	return nil
}
