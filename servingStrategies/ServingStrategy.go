package servingStrategies

type ServingStrategy interface {
	Start(operationPort *string, verbose *bool)
}