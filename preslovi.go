package preslovi

import "strings"

var replacer = strings.NewReplacer(
      "а", "a",
      "А", "A",
      "б", "b",
      "Б", "B",
      "в", "v",
      "В", "V",
      "г", "g",
      "Г", "G",
      "д", "d",
      "Д", "D",
      "ђ", "đ",
      "Ђ", "Đ",
      "е", "e",
      "Е", "E",
      "ж", "ž",
      "Ж", "Ž",
      "з", "z",
      "З", "Z",
      "и", "i",
      "И", "I",
      "ј", "j",
      "Ј", "J",
      "к", "k",
      "К", "K",
      "л", "l",
      "Л", "L",
      "љ", "lj",
      "Љ", "LJ",
      "м", "m",
      "М", "M",
      "н", "n",
      "Н", "N",
      "њ", "nj",
      "Њ", "NJ",
      "о", "o",
      "О", "O",
      "п", "p",
      "П", "P",
      "р", "r",
      "Р", "R",
      "с", "s",
      "С", "S",
      "т", "t",
      "Т", "T",
      "ћ", "ć",
      "Ћ", "Ć",
      "ц", "c",
      "Ц", "C",
      "у", "u",
      "У", "U",
      "ф", "f",
      "Ф", "F",
      "х", "h",
      "Х", "H",
      "ч", "č",
      "Ч", "Č",
      "џ", "dž",
      "Џ", "DŽ",
      "ш", "š",
      "Ш", "Š",
)

func Latinicom(s string) string {
	return replacer.Replace(s)
}