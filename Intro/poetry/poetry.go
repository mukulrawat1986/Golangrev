package poetry

// Line is an alias for a string
type Line string

// Stanza is an alias for a slice of lines
type Stanza []Line

// Poem is an alias for a slice of Stanza
type Poem []Stanza

// Stats
func (p Poem) Stats() (numVowels, numConsonants int) {
	for _, s := range p {
		for _, l := range s {
			for _, r := range l {
				switch r {
				case 'a', 'e', 'i', 'o', 'u':
					numVowels++
				default:
					numConsonants++
				}
			}
		}
	}
}
