/* Alta3 Research | RZFeeser
   Executing system commands  */

package main

import (
	"log"
	"os/exec"
	"fmt"
)

func main() {

	// prepares a "cmd" struct
	cmd, err := exec.Command("ls").Output()
  fmt.Println(cmd)
	//err := cmd.Run()

	if err != nil {
		log.Fatal(err)
		fmt.Println(cmd)
	}
}
