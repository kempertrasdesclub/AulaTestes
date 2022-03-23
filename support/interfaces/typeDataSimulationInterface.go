package interfaces

type DataToSimulateInterface interface {
	Populate() (err error)
	Update() (err error)
	Get() (data interface{})
	GetID() (ID string, err error)
}
