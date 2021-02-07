package magicbell

var api IAPI

// Init initializes the global MagicBell API. This allows for shorthand
// access to all the API methods instead of instantiating and managing your own
// IAPI instance.
func Init(config Config) {
	api = New(config)
}
