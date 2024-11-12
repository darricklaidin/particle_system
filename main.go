package main

import (
	"fmt"
	"particle_system/particles"
	"time"
)

var c = `               .........:--===++++======================++++===--:.........
              ...:-=*#%%%%%%%%%%###***+++++++++++++****##%%%%%%%%%%#*+=::...
              :::.:--=+**##%%%%%%%###***************####%%%%%%%##**+=-:..:::
              :::::::..........::::::::::::::::::::::::::::.........::::::::
              ::::........................     ...................:::::::::.
              :===:......................        .................:::::--==..
               :---==-:::-=-:::..........        ..........::--=====+*+=--:......
               ::::..::::-==+++******++++++++++++++++******++======-::::::.   ....
               .:::........................::::..................:::::::::     ....
                ::::............................................:::::::::.     ....
                 :::......::..................................:::::::::::      ....
                 .:::.....::::..............................::::::::::::.    .....
                  ::::.....:::::::......................:::::::::::::::::::......
                   ::::....::::::::::::::::....::::::::::::::::::::::::::::....
                  .:::::...:::::::::::::::::::::::::::::::::::::::::::..
             ..::---:::::...::::::::::::::::::::::::::::::::::::::::::---:::.
         .::---------::::::..::::::::::::::::::::::::::::::::::::::::---------::.
       ::-----------::::::::..:::::::::::::::::::::::::::::::::::::::------------::.
     .:-------::::::::::::::::.::::::::::::::::::::::::::::::::::::::::::::---------:
    .-----:::::::::::::----:::::::::::::::::::::::::::::::::::------:::::::::::------.
    .--:::::::::::::::-----===-::::::::::::::::::::::::::::-===-----::::::::::::::---.
    ..::::::::::::::::::---=+***+=-::::::::::::::::::::-=+**+==---::::::::::::::::::..
     ..::::::::::::::::::::--==++***++==---------===++**+++=--:::::::::::::::::::::..
     ..:-::::::::::::::::::::::::-----==============-----::::::::::::::::::::::::-:..
       ..:==-::.::::::::::::::::....:::::::::::::::::::::::::::::::::::::::.:-==:..
         ...:-=+==-::...:::::::::::::...............:::::::::::::::..::--=+==::..
            ...::-=++**++==---:::.........................:::--==++**++=-::...
                 ....::--==++***#####*****************####***++==--::.....
                       ......::::::----============----::::::......
                                 ........................


`

func main() {
	coffee := particles.NewCoffee(61, 8, 9.0)
	coffee.Start()

	timer := time.NewTicker(100 * time.Millisecond)
	for {
		<-timer.C
		fmt.Print("\033[H\033[2J")
		coffee.Update()
		steam := coffee.Display()
		for _, row := range steam {
			fmt.Printf("              %s\n", row)
		}
		fmt.Println(c)
	}
}
