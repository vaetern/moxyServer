package communicationBodyService

type ComCachedService struct {
	store StoreService
}

func NewComCachedService() (cs *ComCachedService) {
	cs = &ComCachedService{}
	cs.store = *NewStoreService()

	return cs
}

func (cs ComCachedService) GetCachedBodyFor(hashBody *ComHashedBody, target string)(foundBody string, err error){

	foundBody, err = cs.store.GetBodyByKeyAndTarget(hashBody.Output, target)

	return foundBody, err
}