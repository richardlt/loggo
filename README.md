loggo
=====

import (
	log "github.com/richardlt/loggo"
)

func main() {    

	log.SetLevel(log.DebugLevel)

    log.Debug("my debug message")
    log.Info("my info message")
    
}
