package main

import "fmt"
import "time"
import "os"
import "math/big"

func MyTimer(duration time.Duration) {

	start_time := time.Now()
	end_time := start_time.Add(duration)
	ticker := time.NewTicker(time.Nanosecond)
	go func() {
		for _ = range ticker.C {
			h := big.NewInt(int64(time.Until(end_time).Hours()))
			m := big.NewInt(int64(time.Until(end_time).Minutes()))
			s := big.NewInt(int64(time.Until(end_time).Seconds()))
			s = s.Mod(s, big.NewInt(60))
			m = m.Mod(m, big.NewInt(60))
			fmt.Printf("\r %02d:%02d:%02d", h, m, s)
		}
	}()
	time.Sleep(duration / time.Nanosecond)
	ticker.Stop()
	fmt.Printf("\n Timer completed! \n")
}

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please enter a duration")
		os.Exit(2)
	} else if len(args) == 1 {
		arg1 := args[0]
		duration, err := time.ParseDuration(arg1)
		if err != nil {
			panic(err)
			os.Exit(2)
		}
		MyTimer(duration)
		os.Exit(0)
	}
}
