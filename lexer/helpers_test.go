package lexer

import "testing"

func TestIsLetterMinusLettersAreLetters(t *testing.T) {
	minusLetters := []byte{
		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',
	}

	for _, letter := range minusLetters {
		if !isLetter(letter) {
			t.Errorf("expected %q to be a letter", letter)
		}
	}
}

func TestIsLetterMayusLettersAreLetters(t *testing.T) {
	mayusLetters := []byte{
		'A',
		'B',
		'C',
		'D',
		'E',
		'F',
		'G',
		'H',
		'I',
		'J',
		'K',
		'L',
		'M',
		'N',
		'O',
		'P',
		'Q',
		'R',
		'S',
		'T',
		'U',
		'V',
		'W',
		'X',
		'Y',
		'Z',
	}

	for _, letter := range mayusLetters {
		if !isLetter(letter) {
			t.Errorf("expected %q to be a letter", letter)
		}
	}
}

func TestIsLetterUnderscoreIsALetter(t *testing.T) {
	if !isLetter('_') {
		t.Errorf("expected %q to be a letter", '-')
	}
}

func TestIsLetterNoneOfTheseCharactersAreLetters(t *testing.T) {
	symbols := []byte{
		'°',
		'|',
		'¬',
		'!',
		'"',
		'#',
		'$',
		'%',
		'&',
		'/',
		'(',
		')',
		'=',
		'?',
		'\'',
		'\\',
		'¡',
		'¿',
		'¨',
		'´',
		'*',
		'+',
		'~',
		']',
		'}',
		'[',
		'}',
		'^',
		';',
		',',
		':',
		'.',
		'-',
		'<',
		'>',
		'@',
		0,
	}

	for _, symbol := range symbols {
		if isLetter(symbol) {
			t.Errorf("expected %q to not be a letter", symbol)
		}
	}
}

func TestIsDigitTheseCharactersAreDigits(t *testing.T) {
	digits := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for _, d := range digits {
		if !isDigit(d) {
			t.Errorf("expected %q to be a digit", d)
		}
	}
}

func TestIsDigitTheseCharacterAreNotDigits(t *testing.T) {
	characters := []byte{
		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',
		'A',
		'B',
		'C',
		'D',
		'E',
		'F',
		'G',
		'H',
		'I',
		'J',
		'K',
		'L',
		'M',
		'N',
		'O',
		'P',
		'Q',
		'R',
		'S',
		'T',
		'U',
		'V',
		'W',
		'X',
		'Y',
		'°',
		'|',
		'¬',
		'!',
		'"',
		'#',
		'$',
		'%',
		'&',
		'/',
		'(',
		')',
		'=',
		'?',
		'\'',
		'\\',
		'¡',
		'¿',
		'¨',
		'´',
		'*',
		'+',
		'~',
		']',
		'}',
		'[',
		'}',
		'^',
		';',
		',',
		':',
		'.',
		'-',
		'<',
		'>',
		'@',
		0,
		'Z',
	}

	for _, ch := range characters {
		if isDigit(ch) {
			t.Errorf("expected %q to not be a digit", ch)
		}
	}
}
