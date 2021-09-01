package initiate_model

import (
	"github.com/threagile/threagile/model"
)

func GetMacroDetails() model.MacroDetails {
	return model.MacroDetails{
		ID:    "initiate-model",
		Title: "Initiate a Threat Model",
		Description: "This model macro creates a new model and associated file with a some initial application " +
			"information, create a data component, some technical assets, and network trust boundaries to hold " +
			"the technical assets.",
	}
}
