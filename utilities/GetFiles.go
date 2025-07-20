package utilities

import (
	"github.com/jlaffaye/ftp"
	"github.com/rs/zerolog/log"
	"fmt"
	"time"
	"io/ioutil"
	"strings"
	"bytes"
	"code.sajari.com/docconv/v2"
)

func ListFiles(ipAddress string, username string, password string, configFile string) {
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
			
			/*
			FUTURE TODO:
			- [ ] TEST FILE EXTS & RUN THEM THROUGH RESPECTIVE PARSERS TO EXTRACT DATA 
			*/
			pathSplit := strings.Split(path, ".")
			fileExt := pathSplit[len(pathSplit)-1] // last index to grab file ext
						

			r, err := cHandle.Retr(path)
			if err != nil {
				fmt.Println(err)
			}

			buf, err := ioutil.ReadAll(r)
			if err != nil {
				fmt.Println(err)
			}

			// Extract PDF text accordingly
			if fileExt == "pdf" {
				reader := bytes.NewReader(buf)
				data := ParsePDF(reader)
				matchFound, content := ExamineContents(data, configFile)
				if matchFound == true {
					matchLogData := fmt.Sprintf("Sensitive File Match found - Target: %s Path: %s | Content: %s", ipAddress, path, content)
					log.Info().Str("module", "FTP Sensitive File Found").Msg(matchLogData)
					fmt.Printf("[+] Match found: Target: %s Path: %s | Content: %s\n", ipAddress, path, content)
				}
				r.Close()
			} else {
				// Print match if one is found
				// TODO: Maybe download file if match is found & log it to a file too?
				matchFound, content := ExamineContents(string(buf), configFile)
				if matchFound == true {
					matchLogData := fmt.Sprintf("Sensitive File Match found - Target: %s Path: %s | Content: %s", ipAddress, path, content)
					log.Info().Str("module", "FTP Sensitive File Found").Msg(matchLogData)
					fmt.Printf("[+] Match found: Target: %s Path: %s | Content: %s\n", ipAddress, path, content)
				}
				r.Close()
			}
		}
	}
	
	cHandle.Quit()

}

func ParsePDF(buffer *bytes.Reader) string {
	res, _, err := docconv.ConvertPDF(buffer)
	if err != nil {
		fmt.Println(err)
	}
	return res
}