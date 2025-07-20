# FTP File Inspector (spectr)

Given a list of targets, or singular target, attempt to login & gather *sensitive* files stored on the server.

*Sensitive* files are gathered by matching regexes specified within `utilities/ExamineContent.go`.

`ftpspectr` can operate in 2 different modes:

1. `anonymous` - anonymous login, no creds supplied
2. `spray` - attempt to login using creds supplied in username & password file

If you have regex patterns that you want to search files for, you can specify them with the `-config` flag. All regex patterns must be in a YAML file and be in the following file format:

```yml
patterns:
  - '\b\d{3}-\d{2}-\d{4}\b'
  - '\b[\w.-]+:[^\s:@]{1,100}\b'
```
If you do not specify a regex pattern file, then it will default to hardcoded regex patterns.

## Dependencies

* Uses https://github.com/sajari/docconv for document conversion, must install either via Apt: `poppler-utils wv unrtf tidy` packages: `sudo apt-get install poppler-utils wv unrtf tidy` or via Brew `brew install poppler-qt5 wv unrtf tidy-html5`

## Usage

`go run main.go -mode=spray -target=<singular target or target file> -username=users.txt -password=passwords.txt`

or 

`go run main.go -mode=anonymous -target=<singular target or target file>`


Important data will be logged to a file named `ftpspectr.log`


With a patterns file passed:

![](https://i.ibb.co/LXPYzDhn/2025-07-12-08-27.png)

With no patterns file passed:

![](https://i.ibb.co/gFZMCbpF/2025-07-12-08-29.png)

[![asciicast](https://asciinema.org/a/M0PXkaPCC4BKwd14g9oLfvpKp.svg)](https://asciinema.org/a/M0PXkaPCC4BKwd14g9oLfvpKp)

## Future Todo

- [x] Implement modularity for user-supplied regexes to define what's sensitive
- [ ] Implement channels for handling multiple hosts at once for large-scale inspecting 
- [x] Log each file's contents, even if not sensitive
- [ ] Implement way to detect *hidden* files & directories (EX: `/home/ftp/.aws/credentials`)
- [ ] Implement support for office docs (~~PDF~~, DOCX, ~~ODT~~, etc)
