package match

// Combined returns struct
type Combined struct{}

func (Combined) matchStrings(s1 string, s2 string) float32 {
	var output float32

	output = Simple.matchStrings(Simple{}, s1, s2)

	if output < 0.85 {
		output = DeepDive.matchStrings(DeepDive{}, s1, s2)
	}

	return output
}
