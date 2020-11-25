package controller_adaptor

import (
	"bytes"
	"encoding/binary"
	"github.com/eden-framework/reverse-proxy/worker"
	"github.com/sirupsen/logrus"
)

const prefix = "proxy"
const prefixLength = 5

type ProxyAdaptor struct {
	name       string
	remotePort int
	worker     *worker.Worker
	route      *worker.Route
}

func NewProxyAdaptor(worker *worker.Worker, remotePort int) *ProxyAdaptor {
	return &ProxyAdaptor{
		name:       "ProxyAdaptor",
		remotePort: remotePort,
		worker:     worker,
	}
}

func (p *ProxyAdaptor) Name() string {
	return p.name
}

func (p *ProxyAdaptor) SetName(n string) {
	p.name = n
}

func (p *ProxyAdaptor) Connect() error {
	p.route = p.worker.AddRoute(p.remotePort, worker.Handler{
		HandleFunc: p.handler,
		PackFunc:   p.pack,
		UnpackFunc: p.unpack,
	})
	go p.worker.Start(nil)
	return nil
}

func (p *ProxyAdaptor) Finalize() error {
	p.worker.Stop()
	return nil
}

func (p *ProxyAdaptor) AddHandler() {

}

func (p *ProxyAdaptor) Send(payload []byte) error {
	if p.route != nil {
		payload, err := p.pack(payload)
		if err != nil {
			return err
		}

		return p.route.Send(payload)
	}
	return nil
}

func (p *ProxyAdaptor) handler(payload []byte) (response []byte, err error) {
	logrus.Infof("received: %s", payload)
	return append(payload, []byte(" response")...), nil
}

func (p *ProxyAdaptor) pack(payload []byte) (response []byte, err error) {
	packetLengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(packetLengthBytes, uint32(len(payload)))

	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(prefix)
	buf.Write(packetLengthBytes)
	buf.Write(payload)

	return buf.Bytes(), nil
}

func (p *ProxyAdaptor) unpack(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 校验数据有效位必须大于前缀+4个字节（数据包长度）
	if atEOF || len(data) <= prefixLength+4 {
		return
	}

	// 读取实际数据长度
	var length uint32
	err = binary.Read(bytes.NewReader(data[prefixLength:prefixLength+4]), binary.BigEndian, &length)
	if err != nil {
		return
	}

	availableLength := int(length) + prefixLength + 4
	if availableLength <= len(data) {
		return availableLength, data[prefixLength+4 : availableLength], nil
	}
	return
}
