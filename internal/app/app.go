package app

import (
	"log/slog"
	"sync"

	"github.com/pocketbase/pocketbase"
)

var instance *pocketbase.PocketBase

var intanceOnce sync.Once

func GetApp() *pocketbase.PocketBase {
	intanceOnce.Do(func() {
		instance = pocketbase.NewWithConfig(pocketbase.Config{
			HideStartBanner: true,
		})
	})

	return instance
}

func GetLogger() *slog.Logger {
	return GetApp().Logger()
}
