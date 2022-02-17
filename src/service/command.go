package service

type Command struct {
	Format int
	Symbol byte
	Digit  byte
}

const (
	Esc        byte = 27  // ESC
	Tare       byte = 'T' // тарирование или установка на ноль
	Print      byte = 'P' // печать
	Func       byte = 'F' // кнопки
	Underscore byte = '_'
)

func (c Command) message() (message []byte) {
	if c.Format == 1 {
		message = []byte{Esc, c.Symbol}
	} else {
		message = []byte{Esc, c.Symbol, c.Digit, Underscore}
	}

	return message
}
