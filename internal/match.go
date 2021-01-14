package match

// Match type
type Match struct {
	Strategy strategy
}

// MatchStrings returns float32
func (m *Match) MatchStrings(s1 string, s2 string) uint16 {
	return m.Strategy.matchStrings(s1, s2)
}
