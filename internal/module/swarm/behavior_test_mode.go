package swarm

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func TestWorkingBehavior(r *Swarm) func() {
	return func() {

		for {
			data, err := r.downstreamAdaptor.GetMotor()
			if err != nil {
				logrus.Error(err)
			}
			fmt.Printf("\r%+v", data)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("\r%s", strings.Repeat(" ", 100))
		}

	}
}
