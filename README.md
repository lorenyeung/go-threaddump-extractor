# thread dump extractor go script

## Purpose
Strip thread dumps that were taken using kill -3 from a log file

## Installation
### Standalone Binary
See the releases section :) 
Those can be run with `./td-<DISTRO>-<ARCH>`

### Source code method
Find your go home (`go env`) 
then install under `$GO_HOME/src` (do not create another folder)
`$ git clone https://github.com/lorenyeung/go-threaddump-extractor.git`
then run
`$ go run $GO_HOME/src/go-threaddump-extractor/main.go`

Happy extracting! :)

## Usage
### Commands
* begin
    - Description:
    	- Beginning line of td (default "thread dump")

* end
    - Description:
    	- Ending line of td (default "VM Periodic Task Thread") - Sometimes this is "Metaspace"

* file
    - Description:
    	- File path to strip for thread dumps

* log
    - Description:
    	- Debug log level. Order of Severity: TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC (default "INFO")

* prefix
    - Description:
    	- td file prefix (default "tdfile")

* v	
    - Description:
        - Print the current version and exit

## Dependencies
```
github.com/sirupsen/logrus
```
