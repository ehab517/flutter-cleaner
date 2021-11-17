// This App was developed to help me automate cleaning up the flutter development projects folders
// it simply reads all the folder entries on the searchPath navigates to each directory and runs the command
// searchpath/directory> flutter clean
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
)

var searchPath *string

func main() {
	searchPath = flag.String("s", ".", "specify the search path, if not present it wil default to current directory")
	flag.Parse()
	// Read all the dir entries on the search path
	fileList, err := os.ReadDir(*searchPath)
	if err != nil {
		panic(err)
	}
	// for every entry on the fileList, change to that directory and run the command and print output and success
	for _, f := range fileList {
		if f.IsDir() { //the entry is a directory
			os.Chdir(path.Join(*searchPath, f.Name()))                           //change to that directory path
			fmt.Println("Cleaning ", f.Name(), " ...")                           //print work progress message
			cmd := exec.CommandContext(context.Background(), "flutter", "clean") //prepare the command
			output, err := cmd.Output()                                          //execute the command and get the output from stdout
			fmt.Println(string(output))
			if err != nil {
				fmt.Println(err.Error()) //print the error and continue to the next entry
			} else {
				fmt.Println("** success. **") //print success and continue to the next entry
			}
		}
	}

}
