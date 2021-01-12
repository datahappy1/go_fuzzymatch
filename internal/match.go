package match

type Match struct {
	Strategy strategy
}

// MatchStrings returns float32
func (m *Match) MatchStrings(s1 string, s2 string) float32 {
	return m.Strategy.matchStrings(s1, s2)
}
