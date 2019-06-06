
package exception

type UserError interface {
    error // 内嵌类型
    Message() string
}
