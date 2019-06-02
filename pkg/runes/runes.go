package runes

import (
	"unicode/utf8"
)

// ToByteSlice converts rune slice to byte slice
func ToByteSlice(runeSlice []rune) ([]byte) {
	byteSliceSize := 0
	
	for _, runeChar := range runeSlice {
        byteSliceSize += utf8.RuneLen(runeChar)
    }

    byteSlice := make([]byte, byteSliceSize)
    byteIndex := 0
	
	for _, runeChar := range runeSlice {
        byteIndex += utf8.EncodeRune(byteSlice[byteIndex:], runeChar)
    }

    return byteSlice
}