// perf2Seconds reduces a perf .csv file to a by-seconds perf .csv file
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	log.SetFlags(log.Lshortfile) // show file:line in logs
	flag.Parse()
	if flag.NArg() == 0 || flag.Arg(0) == "-" || flag.Arg(0) == "" {
		fmt.Fprint(os.Stderr, "Usage: perf2seconds file.csv\n")
		os.Exit(1)
	}
	file2Seconds(flag.Arg(0))
}

// file2Seconds sorts transactions and passes them to XXX, to roll up into 1-second samples
func file2Seconds(filename string) {

	// get stdin or a file and sort it
	cmd := exec.Command("sort", "-k2.1,2.2nb", "-k2.5,2.6nb", "-k2.8,2.12nb",
		"--temporary-directory=/var/tmp", filename)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	//out, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Printf("unable to run sort, %v", err)
	//}
	//fmt.Printf("%s\n", out)

}

//	}()
//
//	err = cmd.Start()
//	if err != nil {
//		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
//		os.Exit(1)
//	}
//
//	err = cmd.Wait()
//	if err != nil {
//		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
//		os.Exit(1)
//	}

//next bit:
//	file, err := os.Open("/path/to/file.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		fmt.Println(scanner.Text())
//	}
//
//	if err := scanner.Err(); err != nil {
//		log.Fatal(err)
//	}
//
//}

//
//func perf2Seconds() {
//	// initialize the time to 0
//	// 	cat $name |\
//	//
//	//awk '
//	//NR == 1 {
//	//	# this assume no leading comments
//	//	if ($1 == "#yyy-mm-dd") {
//	//		getline
//	//	}
//	//	date = $1
//	//	sub("\\.[0-9]*", "", $2)
//	//	time = $2
//	//	print "#date time latency xfertime thinktime bytes transactions"
//	//}
//	// /^#/ { echo $0; next } # This does comments: contradiction
//	// /.*/ {
//	//	sub("\\.[0-9]*", "", $2)
//	//	if (time != $2) {
//	//		report(date, time, latency, xfertime, thinktime,
//	//		bytes, transactions)
//	//		date = $1
//	//		time = $2
//	//		latency = $3
//	//		xfertime = $4
//	//		thinktime = $5
//	//		bytes = $6
//	//		transactions = 0
//	//	}
//	//	else {
//	//		latency += $3
//	//		xfertime += $4
//	//		thinktime += $5
//	//		bytes += $6
//	//		transactions++
//	//	}
//	//}
//	//END {
//	//	report(date, time, latency, xfertime, thinktime, bytes, transacyions)
//	//}
//}
//
////func report(date, time, latency, xfertime, thinktime, bytes, xacts) {
////	if xacts > 0 {
////		printf("%s %s %f %f %f %d %d\n",
////			date, time, latency/xacts, xfertime/xacts,
////			thinktime/xacts, bytes/xacts, xacts)
////	}
