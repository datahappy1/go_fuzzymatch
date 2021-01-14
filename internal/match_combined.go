package match

// Combined returns struct
type Combined struct{}

func (Combined) matchStrings(s1 string, s2 string) uint16 {
	var output uint16

	output = Simple.matchStrings(Simple{}, s1, s2)

	if output < 85 {
		output = DeepDive.matchStrings(DeepDive{}, s1, s2)
	}

	return output
}
