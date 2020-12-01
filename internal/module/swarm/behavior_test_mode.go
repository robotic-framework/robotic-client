package swarm

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func TestWorkingBehavior(r *Swarm) func() {
	return func() {
		fmt.Println("get pid...")
		pid, err := r.downstreamAdaptor.GetPID()
		if err != nil {
			logrus.Error(err)
			return
		}
		fmt.Printf("%+v\n", pid)
	}
}
