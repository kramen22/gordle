package state

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/kramen22/gordle/pkg/dictionary"
)

var alphabet = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

type State struct {
	// the target word
	Target string

	// len(Target)
	Width int

	// the list of incorrect letters guessed
	Wrong []string

	// list of correct letters in the wrong spot
	Almost []string

	// the list of correct letters guessed
	Correct map[int]string

	// the list of available letters to guess
	Available []string

	// the map of turns taken, with len(Steps[x]) = Width
	Steps map[int][]int

	// turn we are on, first turn = 0
	Step int

	// dictionary used for the game
	Dict *dictionary.Dictionary
}

func New(dict *dictionary.Dictionary) *State {
	return &State{
		Available: alphabet,
		Dict:      dict,
	}
}

func (s *State) StartGame() {
	s.Available = alphabet
	s.Wrong = nil
	s.Correct = nil
	s.Almost = nil
	s.Step = 0
	s.Steps = make(map[int][]int)

	length := len(s.Dict.Words)
	fmt.Printf("%d \n", length)
	idx := rand.Intn(length)
	for k := range s.Dict.Words {
		switch idx {
		case 0:
			s.Target = k
			s.Width = len(k)
			return
		default:
			idx--
		}
	}
}

func (s *State) GetBoardPrompt() string {
	board := ""

	if s.Step == 0 {
		s.Steps[0] = make([]int, len(s.Target)) // all 0 slice
		for i := 0; i < len(s.Target); i++ {
			board += " _ "
		}
		board += "\n"
	} else {
		for step := 0; step < s.Step; step++ {
			for i := 0; i < len(s.Target); i++ {
				switch s.Steps[step][i] {
				case 0:
					board += " _ "
				case 1:
					board += " ✅ "
				case 2:
					board += " 🟡 "
				}
			}
			board += "\n"
		}
	}

	return board
}

func (s *State) IsValidGuess(guess string) (string, bool) {
	if len(guess) != len(s.Target) {
		return "Guess is the wrong length", false
	}

	if _, ok := s.Dict.Words[guess]; !ok {
		return "Guess is not a valid word", false
	}

	return "", true
}

func (s *State) GuessWord(guess string) bool {
	if guess == s.Target {
		return true
	}

	s.Step++
	s.Steps[s.Step] = make([]int, len(s.Target))

	for idx, toCheck := range guess {
		if string(s.Target[idx]) == string(toCheck) {
			s.Steps[s.Step][idx] = 1
		} else if strings.ContainsRune(s.Target, toCheck) {
			s.Steps[s.Step][idx] = 2
		} else {
			s.Steps[s.Step][idx] = 0
		}
	}

	return false
}
