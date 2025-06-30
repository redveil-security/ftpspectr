package utilities

import (
	"github.com/jlaffaye/ftp"
	"github.com/rs/zerolog/log"
	"fmt"
	"time"
	"io/ioutil"
)

func ListFiles(ipAddress string, username string, password string) {
	cHandle, err := ftp.Dial(ipAddress + ":21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		fmt.Println(err)
	}
	err = cHandle.Login(username, password)
	if err != nil {
		fmt.Println(err)
		log.Info().Str("module", "FTP Failed Login").Msg(err.Error())
	}
	logData := fmt.Sprintf("Login to %s with %s:%s", ipAddress, username, password)
	log.Info().Str("module", "FTP Successful Login").Msg(logData)
	// Recursively grab files & examine contents
	wHandle := cHandle.Walk("/")
	for wHandle.Next() {
		entry := wHandle.Stat()
		path := wHandle.Path()
		if entry.Type == ftp.EntryTypeFile {
			//fmt.Printf("Path: %s\n", path)
			r, err := cHandle.Retr(path)
			if err != nil {
				fmt.Println(err)
			}

			buf, err := ioutil.ReadAll(r)
			if err != nil {
				fmt.Println(err)
			}
			// Log file content even if there's no match
			fileData := fmt.Sprintf("File Found - Target: %s Path: %s | Content %s", ipAddress, path, string(buf))
			log.Info().Str("module", "FTP File Data").Msg(fileData)

			// Print match if one is found
			// TODO: Maybe download file if match is found & log it to a file too?
			matchFound, content := ExamineContents(string(buf))
			if matchFound == true {
				matchLogData := fmt.Sprintf("Sensitive File Match Found - Target: %s Path: %s | Content: %s", ipAddress, path, content)
				log.Info().Str("module", "FTP Sensitive File Found").Msg(matchLogData)
				fmt.Printf("[+] Match found: Target: %s Path: %s | Content: %s\n", ipAddress, path, content)
			}
			r.Close()
		}
	}
	
	cHandle.Quit()

}
