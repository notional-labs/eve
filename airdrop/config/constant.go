package config

import "time"

const (
	LimitPerPage = 100000000
	EveAirdrop   = "1000000000" // 1,000,000,000
	Badkids      = "stars19jq6mj84cnt9p7sagjxqf8hxtczwc8wlpuwe4sh62w45aheseues57n420"
	Cryptonium   = "stars1g2ptrqnky5pu70r3g584zpk76cwqplyc63e8apwayau6l3jr8c0sp9q45u"
	APICoingecko = "https://api.coingecko.com/api/v3/simple/price?ids="
	MaxRetries   = 5
	BackOff      = 200 * time.Millisecond
)
