package coinpaymentslibs

import (
	"log"
)

type StepEnum string

const (
	NO_STEP StepEnum = "NO_STEP"

	STEP_ONE   StepEnum = "WORLD Challenge"
	STEP_TWO   StepEnum = "Verificación"
	STEP_THREE StepEnum = "WORLD Trader"
)

func (step StepEnum) NextPhase(challengeName string) StepEnum {
	log.Println("step.NextPhase")
	log.Println(challengeName)
	log.Println(string(step))
	if step == STEP_THREE {
		return NO_STEP
	}

	if challengeName == NORMAL {
		if step == STEP_ONE {
			return STEP_TWO
		} else {
			return STEP_THREE
		}
	}

	if challengeName == OMICRON {
		return STEP_THREE
	}

	return NO_STEP
}
