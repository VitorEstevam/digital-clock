package clock

type digitalNumber [5]string

var numbers = map[int]digitalNumber{
	0: {
		"█████",
		"█   █",
		"█   █",
		"█   █",
		"█████",
	},
	1: {
		"    █",
		"    █",
		"    █",
		"    █",
		"    █",
	},
	2: {
		"█████",
		"    █",
		"█████",
		"█    ",
		"█████",
	},
	3: {
		"█████",
		"    █",
		"█████",
		"    █",
		"█████",
	},
	4: {
		"█   █",
		"█   █",
		"█████",
		"    █",
		"    █",
	},
	5: {
		"█████",
		"█    ",
		"█████",
		"    █",
		"█████",
	},
	6: {
		"█████",
		"█    ",
		"█████",
		"█   █",
		"█████",
	},
	7: {
		"█████",
		"    █",
		"    █",
		"    █",
		"    █",
	},
	8: {
		"█████",
		"█   █",
		"█████",
		"█   █",
		"█████",
	},
	9: {
		"█████",
		"█   █",
		"█████",
		"    █",
		"█████",
	},
}

var simbols = map[string]digitalNumber{
	":": {
		"     ",
		"  █  ",
		"     ",
		"  █  ",
		"     ",
	},
	" ": {
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	},
}
