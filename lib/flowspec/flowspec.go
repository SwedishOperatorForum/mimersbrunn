package flowspec;

import (
	gobgpd "github.com/osrg/gobgp/server"
	"fmt"
)

func HandleFlowspec(server *gobgpd.BgpServer) {
	w := server.Watch(gobgpd.WatchBestPath(false))
	for {
		select {
		case ev := <-w.Event():
			switch msg := ev.(type) {
			case *gobgpd.WatchEventBestPath:
				for _, path := range msg.PathList {
					fmt.Printf("%+v\n", path);
				}
			}
		}
	}
}