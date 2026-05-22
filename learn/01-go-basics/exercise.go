// Package basics introduces variables, loops, conditionals, and formatting.
// You can run the tests for this package using:
//   go test -v ./learn/01-go-basics/...
package basics

// GreetPlayer returns a welcoming message for a new player.
// It should return "Welcome to Slither, <name>!"
//
// Example: GreetPlayer("Alice") -> "Welcome to Slither, Alice!"
func GreetPlayer(name string) string {
	// TODO: Return the greeting string using string concatenation.
	return ""
}

// ClampSpeed limits a snake's movement speed to stay within the server limits.
// If speed is below min, return min.
// If speed is above max, return max.
// Otherwise, return speed.
func ClampSpeed(speed, min, max float64) float64 {
	// TODO: Implement the clamp logic using if/else statements.
	return 0.0
}

// SumParticles calculates the total score value of a list of eaten particles.
// Use a for loop with 'range' to sum the values in the slice.
func SumParticles(values []int) int {
	// TODO: Iterate over the values slice using range and return the total sum.
	return 0
}

// CountHighScorers counts how many snakes in the world have scores
// greater than or equal to the threshold.
func CountHighScorers(scores []int, threshold int) int {
	// TODO: Iterate over scores and count how many meet or exceed threshold.
	return 0
}

// FormatLeaderboard creates a formatted entry for the leaderboard.
// It should return a string formatted as "Rank. Name: Score pts"
// e.g., "1. Black: 1500 pts"
// Hint: Use fmt.Sprintf.
func FormatLeaderboard(rank int, name string, score int) string {
	// TODO: Format the string using fmt.Sprintf and return it.
	return ""
}
