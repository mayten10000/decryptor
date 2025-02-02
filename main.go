package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func caesarCipher(text string, shift int) string {

	result := ""

	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			result += string((char-'A'+rune(shift))%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			result += string((char-'a'+rune(shift))%26 + 'a')
		} else {
			result += string(char)
		}
	}
	return result
}

func atbashCipher(text string) string {
	result := ""
	for _, char := range text {
		if char >= 'A' && char <= 'z' {
			result += string('Z' - (char - 'A'))
		} else if char >= 'a' && char <= 'z' {
			result += string('z' - (char - 'a'))
		} else {
			result += string(char)
		}
	}
	return result
}

var morseCode = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
	'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---",
	'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---",
	'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
	'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
	'Z': "--..", '1': ".----", '2': "..---", '3': "...--", '4': "....-",
	'5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.",
	'0': "-----", ' ': "/",
}

func MorseCipher(text string) string {
	result := ""

	for _, char := range strings.ToUpper(text) {
		result += morseCode[char] + " "
	}
	return result
}

func vigenereCipher(text, key string) string {
	result := ""
	key = strings.ToUpper(key)
	keyIndex := 0

	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			result += string((char-'A'+rune(key[keyIndex]-'A'))%26 + 'A')
			keyIndex = (keyIndex + 1) % len(key)
		} else if char >= 'a' && char <= 'z' {
			result += string((char-'a'+rune(key[keyIndex]-'A'))%26 + 'a')
			keyIndex = (keyIndex + 1) % len(key)
		} else {
			result += string(char)
		}
	}
	return result
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var text, key string
	var shift, choice int

	fmt.Println("Выберите шифр:")
	fmt.Println("1. Шифр Цезаря")
	fmt.Println("2. Шифр Атбаша")
	fmt.Println("3. Шифр Морзе")
	fmt.Println("4. Шифр Виженера")
	fmt.Scan(&choice)
	fmt.Scanln() // Очистка буфера

	switch choice {
	case 1:
		fmt.Print("Введите текст: ")
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)

		fmt.Print("Введите сдвиг: ")
		fmt.Scan(&shift)

		fmt.Println("Результат:", caesarCipher(text, shift))
	case 2:
		fmt.Print("Введите текст: ")
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)

		fmt.Println("Результат:", atbashCipher(text))
	case 3:
		fmt.Print("Введите текст: ")
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)

		fmt.Println("Результат:", MorseCipher(text))
	case 4:
		fmt.Print("Введите текст: ")
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)

		fmt.Print("Введите ключ: ")
		key, _ = reader.ReadString('\n')
		key = strings.TrimSpace(key)

		if len(key) == 0 {
			fmt.Println("Ошибка: ключ не должен быть пустым")
			return
		}

		fmt.Println("Результат:", vigenereCipher(text, key))
	default:
		fmt.Println("Неверный выбор.")
	}

}
