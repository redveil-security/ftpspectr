package main

import (
	"fmt"
	"time"
	"flag"
	"ftpspectr/utilities"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	target := flag.String("target", "", "Single target or new-line separated list of targets")
	username := flag.String("username", "", "Single username or new-line separated list of usernames")
	password := flag.String("password", "", "Single password or new-line separated list of passwords")
	// Operating modes: anonymous, spray (username & password file)
	mode := flag.String("mode", "anonymous", "Mode to operate in: Anonymous login or user-supplied creds")
	flag.Parse()

	//Initialize logging
	file, err := os.OpenFile("ftpspectr.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = zerolog.New(file).With().Timestamp().Logger()	
	// Ensure target & mode are passed
	if (*target != "") && (*mode != "") {
		// Check if target input is a file or singular input
		if utilities.DoesExist(*target) {
			ips := utilities.ParseInputFile(*target)
			if *mode == "anonymous" {
				fmt.Println("[+] Anonymous mode against target list:", ips)
				// TODO: Implement channels instead
				// anon login on each specified IP
				for _, ip := range ips {
					utilities.ListFiles(ip, "anonymous", "anonymous")
				}

			} else if *mode == "spray" {
				fmt.Println("[+] Pass spray mode against:", ips)
				// Ensure username & input passed
				if (*username != "") && (*password != "") {
					// Ensure both inputs are files
					if (utilities.DoesExist(*username)) && (utilities.DoesExist(*password)) {
						users := utilities.ParseInputFile(*username)
						passwords := utilities.ParseInputFile(*password)
						// 3 for loops to loop through each IP, username, & pass
						for _, ip := range ips {
							for _, user := range users {
								for _, password := range passwords {
									// Ensure username & password being used aren't empty
									if (len(user) != 0) && (len(password) != 0) && (len(ip) != 0) {
										utilities.ListFiles(ip, user, password)
									}
								}
							}
						}
					}
				}
			}
		} else {
			// singular target passed, anon or pass spray mode 
			if *mode == "anonymous" {
				fmt.Println("[+] Anonymous mode against:", *target)
				utilities.ListFiles(*target, "anonymous", "anonymous")
			} else if *mode == "spray" {
				fmt.Println("[+] Pass spray mode against:", *target)
				// Ensure username & password input passed
				if (*username != "") && (*password != "") {
					// Check if inputs are both files
					if (utilities.DoesExist(*username)) && (utilities.DoesExist(*password)) {
						users := utilities.ParseInputFile(*username)
						passwords := utilities.ParseInputFile(*password)
						// 3 loops to loop through each username, & pass to attack single IP
						for _, user := range users {
							for _, password := range passwords {
								if (len(user) != 0) && (len(password) != 0) {
									utilities.ListFiles(*target, user, password)
								}
							}
						}
						
					}
				}
			}
		}
	} else {
		fmt.Println("[!] No args passed! Must pass target & mode to operate in! Run with -h for more info.")
	}
}
