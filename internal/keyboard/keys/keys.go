package keys

type KeyBoardKey string

type AppShortcut struct {
	AppName     string      `yaml:"app_name" json:"app_name"`
	Key         KeyBoardKey `yaml:"key" json:"key"`
	AppFilePath string      `yaml:"app_file_path" json:"app_file_path"`
}

const (
	F1          KeyBoardKey = "f1"
	F2          KeyBoardKey = "f2"
	F3          KeyBoardKey = "f3"
	F4          KeyBoardKey = "f4"
	F5          KeyBoardKey = "f5"
	F6          KeyBoardKey = "f6"
	F7          KeyBoardKey = "f7"
	F8          KeyBoardKey = "f8"
	F9          KeyBoardKey = "f9"
	F10         KeyBoardKey = "f10"
	F11         KeyBoardKey = "f11"
	F12         KeyBoardKey = "f12"
	A           KeyBoardKey = "a"
	B           KeyBoardKey = "b"
	C           KeyBoardKey = "c"
	D           KeyBoardKey = "d"
	E           KeyBoardKey = "e"
	F           KeyBoardKey = "f" //reserved for full screen
	G           KeyBoardKey = "g"
	H           KeyBoardKey = "h"
	I           KeyBoardKey = "i"
	J           KeyBoardKey = "j"
	K           KeyBoardKey = "k"
	L           KeyBoardKey = "l"
	M           KeyBoardKey = "m"
	N           KeyBoardKey = "n"
	O           KeyBoardKey = "o"
	P           KeyBoardKey = "p"
	Q           KeyBoardKey = "q"
	R           KeyBoardKey = "r"
	S           KeyBoardKey = "s"
	T           KeyBoardKey = "t"
	U           KeyBoardKey = "u"
	V           KeyBoardKey = "v"
	W           KeyBoardKey = "w"
	X           KeyBoardKey = "x"
	Y           KeyBoardKey = "y"
	Z           KeyBoardKey = "z"
	NumberZero  KeyBoardKey = "0"
	NumberOne   KeyBoardKey = "1"
	NumberTwo   KeyBoardKey = "2"
	NumberThree KeyBoardKey = "3"
	NumberFour  KeyBoardKey = "4"
	NumberFive  KeyBoardKey = "5"
	NumberSix   KeyBoardKey = "6"
	NumberSeven KeyBoardKey = "7"
	NumberEight KeyBoardKey = "8"
	NumberNine  KeyBoardKey = "9"
	NumberTen   KeyBoardKey = "10"
	SemiColon   KeyBoardKey = ";" //does not work as hot key with karabiner
)
