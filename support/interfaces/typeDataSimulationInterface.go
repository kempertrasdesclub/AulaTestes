package interfaces

type DataToSimulateInterface interface {
	Populate() (err error)
	Update() (err error)
	Get() (data interface{})
	GetID() (ID interface{}, err error)
}
