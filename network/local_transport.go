package network

import (
	"fmt"
	"sync"
)

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
	t.lock.Lock()
	defer t.lock.Unlock()

	peer, ok := t.peers[netAddr]
	if !ok {
		return fmt.Errorf("%s: couldn't send message to %s", t.addr, netAddr)
	}

	peer.consumeCh <- RPC{
		From:    t.addr,
		Payload: message,
	}
	return nil
}

func (t *LocalTransport) Connect(tr *LocalTransport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr

	return nil
}
