//programma 15sm ne prigovor

/*dlinnii komment
i tyt
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Vvedite dlinu: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Naiden Error1   ")
		log.Fatal(err)

	}

	input = strings.TrimSpace(input)
	size, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Print("Naiden Error2   ")
		log.Fatal(err)

	}

	var status string
	if size < 15 {
		status = "U men9 dl9 teb9 est' klassnii krem"
	} else if size == 15 {
		status = "15sm ne prigovor"
	} else {
		status = "ti slavnii maliy"
	}

	fmt.Println(status)
}
