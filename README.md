# FTP File Inspector (spectr)

Given a list of targets, or singular target, attempt to login & gather *sensitive* files stored on the server.

*Sensitive* files are gathered by matching regexes specified within `utilities/ExamineContent.go`.

`ftpspectr` can operate in 2 different modes:

1. `anonymous` - anonymous login, no creds supplied
2. `spray` - attempt to login using creds supplied in username & password file

## Usage

`go run main.go -mode=spray -target=<singular target or target file> -username=users.txt -password=passwords.txt`

or 

`go run main.go -mode=anonymous -target=<singular target or target file>`

Important data will be logged to a file named `ftpspectr.log`

![](https://i.ibb.co/DfnT2tct/2025-06-29-16-55.png)

## Future Todo

- [ ] Implement modularity for user-supplied regexes to define what's sensitive
- [ ] Implement channels for handling multiple hosts at once for large-scale inspecting 
- [x] Log each file's contents, even if not sensitive
