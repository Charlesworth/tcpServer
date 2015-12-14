package portManager

import "errors"

type portManager struct {
	portChan chan int
}

func (pM *portManager) TakePort() (port int, err error) {
	if len(pM.portChan) == 0 {
		return 0, errors.New("No more ports in port manager")
	}
	port = <-pM.portChan
	return port, nil
}

func (pM *portManager) ReturnPort(port int) {
	pM.portChan <- port
}

func New(firstPort int, lastPort int) (pM *portManager) {
	portAmount := lastPort - firstPort + 1
	pM = &portManager{make(chan int, portAmount)}
	for i := firstPort; i <= lastPort; i++ {
		pM.portChan <- i
	}
	return pM
}
