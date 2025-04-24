package internal

import (
	"fmt"
	"strconv"
)

const (
	// NOTE: These are all ansi codes
	// FG: Foreground, BG: background
	ANSI_RESET_COLOR = "\033[0m"
	FG_BLACK         = 30
	FG_RED           = 31
	FG_GREEN         = 32
	FG_YELLOW        = 33
	FG_BLUE          = 34
	FG_MAGENTA       = 35
	FG_CYAN          = 36
	FG_LIGHTGRAY     = 37
	FG_DARKGRAY      = 90
	FG_LIGHTRED      = 91
	FG_LIGHTGREEN    = 92
	FG_LIGHTYELLOW   = 93
	FG_LIGHTBLUE     = 94
	FG_LIGHTMAGENTA  = 95
	FG_LIGHTCYAN     = 96
	FG_WHITE         = 97
	BG_BLACK         = 40
	BG_RED           = 41
	BG_GREEN         = 42
	BG_YELLOW        = 43
	BG_BLUE          = 44
	BG_MAGENTA       = 45
	BG_CYAN          = 46
	BG_LIGHTGRAY     = 47
	BG_DARKGRAY      = 100
	BG_LIGHTRED      = 101
	BG_LIGHTGREEN    = 102
	BG_LIGHTYELLOW   = 103
	BG_LIGHTBLUE     = 104
	BG_LIGHTMAGENTA  = 105
	BG_LIGHTCYAN     = 106
	BG_WHITE         = 107
)

type color struct {
	Fg int // Foreground
	Bg int // Background
}

func ColorPrefix(s string) string {
	c := color{Fg: FG_BLACK, Bg: BG_LIGHTGREEN}
	return colorString(c, s)
}

func ColorString(bgColor, fgColor, s string) string {
	c := color{Fg: getFgColor(fgColor), Bg: getBgColor(bgColor)}
	return colorString(c, s)
}

func colorString(c color, v string) string {
	if (c.Fg < 30 || c.Fg > 97) || (c.Fg > 37 && c.Fg < 90) {
		c.Fg = 39
	}
	if (c.Bg < 40 || c.Bg > 107) || (c.Bg > 47 && c.Bg < 100) {
		c.Bg = 49
	}
	return fmt.Sprintf("\033[%s;%sm%s%s", strconv.Itoa(c.Fg), strconv.Itoa(c.Bg), v, ANSI_RESET_COLOR)
}

func getFgColor(c string) int {
	switch c {
	case "black":
		return FG_BLACK
	case "red":
		return FG_RED
	case "green":
		return FG_GREEN
	case "yellow":
		return FG_YELLOW
	case "blue":
		return FG_BLUE
	case "magenta":
		return FG_MAGENTA
	case "cyan":
		return FG_CYAN
	case "light gray":
		return FG_LIGHTGRAY
	case "dark gray":
		return FG_DARKGRAY
	case "light red":
		return FG_LIGHTRED
	case "light green":
		return FG_LIGHTGREEN
	case "light yellow":
		return FG_LIGHTYELLOW
	case "light blue":
		return FG_LIGHTBLUE
	case "light magenta":
		return FG_LIGHTMAGENTA
	case "light cyan":
		return FG_LIGHTCYAN
	case "white":
		return FG_WHITE
	default:
		return 39
	}
}

func getBgColor(c string) int {
	switch c {
	case "black":
		return BG_BLACK
	case "red":
		return BG_RED
	case "green":
		return BG_GREEN
	case "yellow":
		return BG_YELLOW
	case "blue":
		return BG_BLUE
	case "magenta":
		return BG_MAGENTA
	case "cyan":
		return BG_CYAN
	case "light gray":
		return BG_LIGHTGRAY
	case "dark gray":
		return BG_DARKGRAY
	case "light red":
		return BG_LIGHTRED
	case "light green":
		return BG_LIGHTGREEN
	case "light yellow":
		return BG_LIGHTYELLOW
	case "light blue":
		return BG_LIGHTBLUE
	case "light magenta":
		return BG_LIGHTMAGENTA
	case "light cyan":
		return BG_LIGHTCYAN
	case "white":
		return BG_WHITE
	default:
		return 49
	}
}
