package logic

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)
