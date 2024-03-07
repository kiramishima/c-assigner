package models

import (
	"errors"
)

const (
	CreditBase int32 = 100
	Credit300  int32 = 300
	Credit500  int32 = 500
	Credit700  int32 = 700
	Credit1500 int32 = 1500
)

// CreditAssigner interface
type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

var _ CreditAssigner = (*Assigner)(nil)

type Assigner struct {
	Inv300 int32
	Inv500 int32
	Inv700 int32
}

func (c Assigner) Assign(investment int32) (int32, int32, int32, error) {
	// Está garantizado que el monto de inversión es un múltiplo de 100.
	if v := investment % CreditBase; v != 0 {
		return 0, 0, 0, errors.New("investment needs be multiply of 100")
	}

	if investment < Credit300 {
		return 0, 0, 0, errors.New("Can't assign credits with this amount of investment")
	} else if investment > Credit300 && investment < Credit500 {
		return 0, 0, 0, errors.New("Can't assign credits with this amount of investment")
	} else if (investment%Credit300 > 0) && (investment > Credit500) && (investment < Credit700) {
		return 0, 0, 0, errors.New("Can't assign credits with this amount of investment")
	} else {
		// calc max div
		max300 := investment / Credit300
		max500 := investment / Credit500
		max700 := investment / Credit700

		for start300 := int32(0); start300 <= max300; start300++ {
			for start500 := int32(0); start500 <= max500; start500++ {
				for start700 := int32(0); start700 <= max700; start700++ {
					// eval
					amount := start300*Credit300 + start500*Credit500 + start700*Credit700
					if amount == investment {
						c.Inv300 = start300
						c.Inv500 = start500
						c.Inv700 = start700
						break
					}
					if amount > investment {
						break
					}
				}
			}
		}

		// check
		amount := c.Inv300*Credit300 + c.Inv500*Credit500 + c.Inv700*Credit700
		// fmt.Println("obj ", c)
		// fmt.Println("amount ", amount)
		if amount == investment {
			return c.Inv300, c.Inv500, c.Inv700, nil
		} else {
			return 0, 0, 0, errors.New("Las cantidades no coincinden")
		}
	}
}
