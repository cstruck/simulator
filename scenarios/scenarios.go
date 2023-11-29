package scenarios

import (
	"embed"
	"errors"
	"log/slog"
	"sort"

	"gopkg.in/yaml.v2"
)

//go:embed scenarios.yaml
var config embed.FS

type Scenario struct {
	ID          string `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Category    string `yaml:"category"`
	Difficulty  string `yaml:"difficulty"`
}

func List() ([]Scenario, error) {
	var scenarios []Scenario

	bytes, err := config.ReadFile("scenarios.yaml")
	if err != nil {
		slog.Error("failed to load scenarios file")
		return nil, errors.Join(errors.New("failed to list scenarios"), err)
	}

	err = yaml.Unmarshal(bytes, &scenarios)
	if err != nil {
		slog.Error("failed to unmarshall scenarios")
		return nil, errors.Join(errors.New("failed to list scenarios"), err)
	}

	sort.Slice(scenarios, func(i, j int) bool {
		iDifficulty := scenarios[i].Difficulty
		jDifficulty := scenarios[j].Difficulty

		switch iDifficulty {
		case "Complex":
			return jDifficulty != "Complex"
		case "Medium":
			return jDifficulty == "Easy"
		default:
			return false
		}
	})

	return scenarios, nil
}

func Find(scenarioID string) (Scenario, error) {
	var scenario Scenario

	scenarios, err := List()
	if err != nil {
		return scenario, errors.Join(errors.New("failed to find scenario"), err)
	}

	for _, scenario = range scenarios {
		if scenario.ID == scenarioID {
			return scenario, nil
		}
	}

	return scenario, errors.New("unable to find scenario")
}
