package main

import "fmt"
import "time"
import "os"
import "math/big"

func MyTimer(duration time.Duration, note string) {
	start_time := time.Now()
	end_time := start_time.Add(duration)
	h := big.NewInt(int64(time.Until(end_time).Hours()))
	m := big.NewInt(int64(time.Until(end_time).Minutes()))
	s := big.NewInt(int64(time.Until(end_time).Seconds()))
	s = s.Mod(s, big.NewInt(60))
	m = m.Mod(m, big.NewInt(60))
	fmt.Printf("Start:\t %s \n", start_time.Format(time.ANSIC))
	fmt.Printf("End:\t %s\n", end_time.Format(time.ANSIC))
	if note != "" {

		fmt.Printf("Note:\t %s\n", note)
	}
	fmt.Printf("\rTimer:\t %02d:%02d:%02d", h, m, s)
	i := int64(1)
	ticker1 := time.NewTicker(1 * time.Second)
	for _ = range ticker1.C {
		i++
		h = big.NewInt(int64(time.Until(end_time).Hours()))
		m = big.NewInt(int64(time.Until(end_time).Minutes()))
		s = big.NewInt(int64(time.Until(end_time).Seconds()))
		s = s.Mod(s, big.NewInt(60))
		m = m.Mod(m, big.NewInt(60))
		fmt.Printf("\rTimer:\t %02d:%02d:%02d", h, m, s)
		if i > int64(duration/time.Second) {
			ticker1.Stop()
			break
		}
	}
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
		MyTimer(duration, "")
		os.Exit(0)
	} else if len(args) == 2 {
		arg1 := args[0]
		arg2 := args[1]
		duration, err := time.ParseDuration(arg1)
		if err != nil {
			panic(err)
			os.Exit(2)
		}
		MyTimer(duration, arg2)
		os.Exit(0)
	}
}
