package debeziumSimulation

import "errors"

func (e *DebeziumSimulation) Populate(createTotal int, createUpdate, createDelete float64) (err error) {
	if e.realDataPointer == nil {
		err = errors.New("use SetData() function first")
		return
	}

	for i := 0; i != createTotal; i += 1 {
		_, _, err = e.GetCreate()
		if err != nil {
			return
		}

		if e.getRandFloat64() > createUpdate {
			_, _, _, err = e.GetUpdate()
			if err != nil {
				return
			}
		}

		if e.getRandFloat64() > createDelete {
			_, _, err = e.GetDelete()
			if err != nil {
				return
			}
		}
	}

	return
}
