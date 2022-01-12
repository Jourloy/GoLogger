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
import "github.com/Jourloy/GoLogger"
```

### Init in your file

```golang
var logger GoLogger.Logger

func init() {
    logger = GoLogger.New(1, "Main")
}
```

Logger have 5 levels:

-   1 - Error
-   2 - Warn
-   3 - Log
-   4 - Info
-   5 - Debug

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
