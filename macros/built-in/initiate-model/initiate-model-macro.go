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

func GetNextQuestion() (nextQuestion model.MacroQuestion, err error) {
	return model.NoMoreQuestions(), nil
}

func ApplyAnswer(questionID string, answer ...string) (message string, validResult bool, err error) {
	return "Answer processed", true, nil
}

func GoBack() (message string, validResult bool, err error) {
	return "Cannot go back further", false, nil
}

func GetFinalChangeImpact(modelInput *model.ModelInput) (changes []string, message string, validResult bool, err error) {
	return []string{"creating the initial model file"}, "Changeset valid", true, err
}

func Execute(modelInput *model.ModelInput) (message string, validResult bool, err error) {
	return "Initial Threagile model created successful", true, nil
}
