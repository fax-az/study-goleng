package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("9 zagadau 4islo ot 1 do 100")
	fmt.Println("smozeIII ugadat'?")

	second := time.Now().Unix()

	rand.Seed(second)
	target := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)
	succes := false

	for i := 1; i <= 10; i++ {
		fmt.Println("Popitik ostalos'", 11-i, " \nVvedite celoe 4islo: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Naiden Error1   ")
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		otvet, err := strconv.Atoi(input)

		if err != nil {
			fmt.Print("Naiden Error2   ")
			log.Fatal(err)

		}

		if otvet < target {
			fmt.Println("VaIIIe 4islo men'IIIe zagadannogo")
		} else if otvet > target {
			fmt.Println("VaIIIe 4islo bol'IIIe zagadannogo")
		} else {
			succes = true
			fmt.Println("Vi pobedili")
			fmt.Println("Zagadannoe 4islo=", target)
			break
		}
	}

	if !succes {
		fmt.Println("K sogeleniu vi proigrali. Zagadannoe 4islo", target)
	}

}
