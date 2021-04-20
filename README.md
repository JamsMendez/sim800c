Client TCP SIM800c 
========

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/JamsMendez/sim800c"
)

func main() {
  
  	client := sim800c.ClientTCP{
		PortName: "COM6",
		BaudRate: 9600,
		Debug:    true,
		APN:      "internet.itelcel.com",
		User:     "webqprs",
		Password: "webqprs2002",
		IP:       "0.0.0.0",
		Port:     "3000",
		PinSIM:   "1111",
	}

	err := client.Open()
	if err != nil {
		log.Fatal(err)
	}
  
  	defer client.Close()

	err = client.Connect()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CONNECTION SUCCESS")

	err = client.Send("JamsMendez")
	if err != nil {
		fmt.Println("MSG.ERRROR: ", err)
	}

	fmt.Println("Wait ...")

	time.Sleep(time.Second * 5)

	err = client.Disconnect()
	if err != nil {
		fmt.Println(err)
	}
}
```

# Contact

Github: [https://github.com/jamsmendez](https://github.com/jamsmendez/)

Twitter: [https://twitter.com/jamsmendez](https://twitter.com/jamsmendez)