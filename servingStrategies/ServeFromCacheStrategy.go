package servingStrategies

type ServeFromCacheStrategy struct {
}

func NewServeFromCacheStrategy() ServeFromCacheStrategy {
	return ServeFromCacheStrategy{}
}

func (s ServeFromCacheStrategy) Start(operationPort *string, verbose *bool) {

}
