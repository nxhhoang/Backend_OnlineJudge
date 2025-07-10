package utils

import (
	"encoding/json"
	"judge/models"
	"os"

	"github.com/antchfx/xmlquery"
)

/*
ParseProblem() - Read problem.xml

Current problems
- Only english statements are allowed

Workflow:
- Read problem.xml to get Name, ShortName, Tags and create a Problem
*/
func ParseProblemStruct(xml *os.File) (models.Problem, error) {
	var problem models.Problem

	doc, err := xmlquery.Parse(xml)
	if err != nil {
		return problem, err
	}

	problem.Name = xmlquery.Find(doc, "//problem/names/name[@language='english']/@value")[0].InnerText()
	problem.ShortName = xmlquery.Find(doc, "//problem/@short-name")[0].InnerText()
	tags := xmlquery.Find(doc, "//problem/tags/tag")
	for _, tag := range tags {
		problem.Tags = append(problem.Tags, tag.SelectAttr(("value")))
	}

	return problem, nil
}

/*
SaveProblemToJson - Save Problem to a json file with specified filepath
*/
func SaveProblemToJson(problem models.Problem, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(problem); err != nil {
		return err
	}

	return nil
}
