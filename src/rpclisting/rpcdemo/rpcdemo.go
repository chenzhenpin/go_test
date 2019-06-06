package rpcdemo

import "errors"

// 服务
type DemoService struct {
}

// 参数
type Args struct {
    A, B int
}

// jsonrpc 的服务需要定义为参入参数和传出参数的格式
// 传出参数必须为指针格式
func (e DemoService) DIV(args Args, result *float64) error {
    if args.B == 0 {
        return errors.New("division by zero")
    }

    *result = float64(args.A) / float64(args.B)
    return nil
}