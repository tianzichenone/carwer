package rpcdaemon

import "fmt"

type RpcDaemonService struct {

}

type Args struct {
	A, B int
}

func (s RpcDaemonService) Divide(args Args, result *float64) error {
	if args.B == 0 {
		return fmt.Errorf("Divide zero error")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}