package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"problem/models"
	"problem/utils/polygon"
)

/*
AddProblem(ProblemId, PackageId)

Workflow:
- Download the packages using DownloadPackage()

FUTURE:
- Automatically get the latest package of the problem
*/

func AddProblem(ProblemId uint64, PackageId uint64) error {
	err := polygon.DownloadPackage(ProblemId, PackageId)
	if err != nil {
		return err
	}

	file, err := os.Open(fmt.Sprintf("%s/%d/problem.json", os.Getenv("PROBLEM_STORAGE_DIR"), ProblemId))
	if err != nil {
		return err
	}
	defer file.Close()

	var problem models.Problem
	if err := json.NewDecoder(file).Decode(&problem); err != nil {
		return err
	}

	// if _, err := db.Database("Problems").Collection("Problems").InsertOne(ctx, problem); err != nil {
	// 	return err
	// }

	return nil
}
