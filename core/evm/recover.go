package evm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AlexRipoll/go-bridge/core/event"
	"github.com/AlexRipoll/go-bridge/sys/storage"
	log "github.com/sirupsen/logrus"
)

func RecoverFromShutDown(ctx context.Context, db storage.Db, releaser Releaser) error {
	log.Infof("initializing recovery...")
	keys := db.GetAllKeys()
	fmt.Println("UNPROCESSED KEYS: ", keys)

	if len(keys) == 0 {
		log.Infof("nothing to recover...")
	}
	for _, key := range keys {
		buf, err := db.Get(key)
		if err != nil {
			return err
		}

		var rx event.Rx
		if err := json.Unmarshal(buf, &rx); err != nil {
			return err
		}

		log.Infof("recovering %v", string(key))
		go releaser.releaseToken(ctx, rx)
	}

	return nil
}
