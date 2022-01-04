# GoLogger

Make logs more informative

This tool add time before log, make logs differents colors and put name of file right in log message

## How to use

### Download package

```bash
$ go get github.com/Jourloy/GoLogger
```

### Import package

```golang
import logger "github.com/Jourloy/GoLogger"
```

### Init in your file

```golang
func init() {
    logger.New("all", "File name")
}
```

Logger have two levels:
- `all` - this is all output 
- `prod` - this is only LOG, WARN and ERROR output

### Usage

#### Log

```golang
logger.Log("Success")
```

#### Warning

```golang
logger.Warn("This password is short")
```

#### Error

```golang
logger.Error("File not found")
```

#### Info

```golang
logger.Info("File don't have a name")
```

#### Debug

```golang
logger.Debug("1")
```
