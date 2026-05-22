package basics

import "testing"

func TestGreetPlayer_ReturnsWelcomeMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Standard name", "Alice", "Welcome to Slither, Alice!"},
		{"Name with spaces", "Snake Master 9000", "Welcome to Slither, Snake Master 9000!"},
		{"Empty name", "", "Welcome to Slither, !"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GreetPlayer(tt.input)
			if got != tt.expected {
				t.Errorf("GreetPlayer(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestClampSpeed_ClampsCorrectly(t *testing.T) {
	tests := []struct {
		name     string
		speed    float64
		min      float64
		max      float64
		expected float64
	}{
		{"Within limits", 5.5, 1.0, 10.0, 5.5},
		{"Below limit", 0.5, 1.0, 10.0, 1.0},
		{"Above limit", 12.3, 1.0, 10.0, 10.0},
		{"Exactly min", 1.0, 1.0, 10.0, 1.0},
		{"Exactly max", 10.0, 1.0, 10.0, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClampSpeed(tt.speed, tt.min, tt.max)
			if got != tt.expected {
				t.Errorf("ClampSpeed(%f, %f, %f) = %f; want %f", tt.speed, tt.min, tt.max, got, tt.expected)
			}
		})
	}
}

func TestSumParticles_ReturnsCorrectSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Multiple particles", []int{10, 20, 5, 15}, 50},
		{"Single particle", []int{42}, 42},
		{"No particles", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumParticles(tt.input)
			if got != tt.expected {
				t.Errorf("SumParticles(%v) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}

func TestCountHighScorers_CountsCorrectly(t *testing.T) {
	tests := []struct {
		name      string
		scores    []int
		threshold int
		expected  int
	}{
		{"Some above threshold", []int{100, 250, 50, 400, 150}, 200, 2},
		{"All above threshold", []int{300, 400, 500}, 200, 3},
		{"None above threshold", []int{50, 100, 150}, 200, 0},
		{"Exactly threshold is counted", []int{200, 150, 300}, 200, 2},
		{"Empty scores list", []int{}, 200, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountHighScorers(tt.scores, tt.threshold)
			if got != tt.expected {
				t.Errorf("CountHighScorers(%v, %d) = %d; want %d", tt.scores, tt.threshold, got, tt.expected)
			}
		})
	}
}

func TestFormatLeaderboard_FormatsCorrectly(t *testing.T) {
	tests := []struct {
		name     string
		rank     int
		playerName string
		score    int
		expected string
	}{
		{"First place", 1, "Alice", 9999, "1. Alice: 9999 pts"},
		{"Tenth place", 10, "Bob", 450, "10. Bob: 450 pts"},
		{"Zero score", 5, "NoobMaster", 0, "5. NoobMaster: 0 pts"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatLeaderboard(tt.rank, tt.playerName, tt.score)
			if got != tt.expected {
				t.Errorf("FormatLeaderboard(%d, %q, %d) = %q; want %q", tt.rank, tt.playerName, tt.score, got, tt.expected)
			}
		})
	}
}
