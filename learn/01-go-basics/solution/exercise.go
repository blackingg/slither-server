// Package solution provides the reference implementations for the basics exercises.
package solution

import "fmt"

// GreetPlayer returns a welcoming message for a new player.
// We use simple string concatenation here.
func GreetPlayer(name string) string {
	return "Welcome to Slither, " + name + "!"
}

// ClampSpeed limits a snake's movement speed to stay within the server limits.
// Since Go standard library doesn't have a generic math.Clamp or a float64 clamp,
// writing explicit if/else conditions is the standard Go approach.
func ClampSpeed(speed, min, max float64) float64 {
	if speed < min {
		return min
	}
	if speed > max {
		return max
	}
	return speed
}

// SumParticles calculates the total score value of a list of eaten particles.
// We use a standard for-range loop. The blank identifier '_' is used to ignore the
// index because Go will complain of an unused variable if we declare it and don't use it.
func SumParticles(values []int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

// CountHighScorers counts how many snakes in the world have scores
// greater than or equal to the threshold.
func CountHighScorers(scores []int, threshold int) int {
	count := 0
	for _, score := range scores {
		if score >= threshold {
			count++
		}
	}
	return count
}

// FormatLeaderboard creates a formatted entry for the leaderboard.
// We use fmt.Sprintf with '%d' for integers and '%s' for the name string.
func FormatLeaderboard(rank int, name string, score int) string {
	return fmt.Sprintf("%d. %s: %d pts", rank, name, score)
}
