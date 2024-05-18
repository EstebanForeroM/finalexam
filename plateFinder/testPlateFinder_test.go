package platefinder

import "testing"

func TestIsDigit(t *testing.T) {
    err := comproveCharacterIsDigit('1')
    
    if err != nil {
        t.Error("Character should be a digit")
    }

    err = comproveCharacterIsDigit('a')

    if err == nil {
        t.Error("Character should not be a digit")
    }
}

func TestIdDigitComplete(t *testing.T) {
    for i := 0; i < 10; i++ {
        err := comproveCharacterIsDigit(byte(i) + '0')

        if err != nil {
            t.Error("Character should be a digit")
        }
    }
}
