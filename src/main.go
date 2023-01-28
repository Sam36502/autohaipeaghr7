package main

import (
	"autohaipeaghr7/src/model"
	"autohaipeaghr7/src/util"
	"fmt"
	"os"

	fs "github.com/Sam36502/funcscript"
)

func main() {
	err := fs.Initialise()
	if err != nil {
		fmt.Printf("[Error]: Failed to initialise FuncScript: %v; Aborting...\n", err)
		os.Exit(1)
	}

	err = model.LoadDataFile(util.DATA_FILE_NAME)
	if err != nil {
		fmt.Printf(
			"[Error]: Failed to load data file '%s': %v; Aborting...\n",
			util.DATA_FILE_NAME, err)
		os.Exit(1)
	}

	A := model.NewTeamOfIDs("Sussy Balls", []string{
		"furry",
		"weeb",
		"",
		"weeb",
	})

	fmt.Println(A.ToString())

	A.Members[1].TakeDamage(1000000)
	A.CheckFainted()

	fmt.Println(A.ToString())

}
