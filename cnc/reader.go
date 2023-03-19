package main

import "strings"

func (Term *Terminal) ReadLine(Prompt string, Password bool, maxLen int) (string, error) {
	return Term.readLine(Prompt, maxLen, Password, "*", true)
}

func (Term *Terminal) readLine(Prompt string, maxLen int, password bool, passwordChar string, Gradient bool) (string, error) {

	Term.RelayPos = len(Term.BytesReceived)

	var Message []byte

	Term.Conn.Write([]byte{
		255, 251, 1,
		255, 251, 3,
		255, 252, 34,
	})

	if _, error := Term.Conn.Write([]byte("\r")); error != nil {
		return "", error
	}

	if _, error := Term.Conn.Write([]byte(Prompt)); error != nil {
		return "", error
	}

	var CharPos int = 0
	var CharValue int = Term.RelayPos

	for {
		BufLimit := make([]byte, 1)
		ReqLen, error := Term.Conn.Read(BufLimit)
		if error != nil {
			return "", error
		} else if ReqLen > 1 {
			continue
		}

		switch BufLimit[ReqLen-1] {
		case 9:

			if password {
				continue
			}

			if CharValue == 0 && CharPos > 0 {
				CharValue = Term.RelayPos
			}

			if CharValue > 0 {
				CharValue--
			} else {
				continue
			}

			if _, error := Term.Conn.Write([]byte("\r\033[2K\x1b[0m")); error != nil {
				return "", error
			}

			if Gradient {
				var GradientApply string
				for Pos, i := range Term.BytesReceived[CharValue] {
					if Pos > len(Term.GradientCurve) {
						continue
					}
					GradientApply += Term.GradientCurve[Pos] + string(i) + "\x1b[0m"
					continue
				}

				if _, error := Term.Conn.Write([]byte(strings.Split(Prompt, "\r\n")[len(strings.Split(Prompt, "\r\n"))-1] + GradientApply)); error != nil {
					return "", error
				}
			} else {
				if _, error := Term.Conn.Write([]byte(strings.Split(Prompt, "\r\n")[len(strings.Split(Prompt, "\r\n"))-1] + Term.BytesReceived[CharValue])); error != nil {
					return "", error
				}
			}

			Message = []byte(Term.BytesReceived[CharValue])
			CharPos = len(string(Message))

		case 27:
			BufLimit := make([]byte, 9)
			ReqLength, error := Term.Conn.Read(BufLimit)
			if error != nil {
				return "", nil
			} else if ReqLength > 9 {
				continue
			}

			continue

		case 13, 11, 12, 5, 8:
			continue

		case 127:
			if len(Message) == 0 {
				continue
			}
			Term.Conn.Write([]byte{127})
			var StoreMessage []byte
			for Pos, value := range Message {
				if Pos == len(Message)-1 {
					CharPos--
					continue
				} else {
					StoreMessage = append(StoreMessage, value)
				}
			}
			Message = StoreMessage

		case 10:
			Term.BytesReceived = append(Term.BytesReceived, string(Message))
			Term.Conn.Write([]byte("\r\n"))
			return string(Message), nil

		default:

			if Gradient && len(Term.GradientCurve) > len(Message) {
				if password {
					Term.Conn.Write([]byte(Term.GradientCurve[CharPos] + passwordChar + "\x1b[0m"))
				} else {
					Term.Conn.Write([]byte(Term.GradientCurve[CharPos] + string(BufLimit) + "\x1b[0m"))
				}

				Message = append(Message, BufLimit...)
				CharPos++
				continue
			}

			if !Gradient && len(Message) < maxLen {
				if password {
					Term.Conn.Write([]byte(passwordChar))
				} else {
					Term.Conn.Write([]byte(BufLimit))

				}

				Message = append(Message, BufLimit...)
				CharPos++
				continue
			}
		}
	}
}
