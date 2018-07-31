package filters

type Filter interface {
	handle() bool
}
