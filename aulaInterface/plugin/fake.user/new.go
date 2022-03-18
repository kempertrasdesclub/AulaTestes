package main

func (e *FakeUser) New() (referenceInitialized interface{}, err error) {
	return e, nil
}
