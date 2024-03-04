package inv

var (
	debugMode = true
)

// DebugMode disables any checks for invariants that have a debug prefix.
func NoDebug() {
	debugMode = false
}
