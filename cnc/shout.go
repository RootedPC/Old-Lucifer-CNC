package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func getLetter(str string) []string {

	A := make([]string, 8)
	A[0] = `     /\     `
	A[1] = `    /  \    `
	A[2] = `   / /\ \   `
	A[3] = `  / ____ \  `
	A[4] = ` /_/    \_\ `
	A[5] = `            `

	B := make([]string, 8)
	B[0] = ` ____    `
	B[1] = ` |  _ \  `
	B[2] = ` | |_) | `
	B[3] = ` |  _ <  `
	B[4] = ` | |_) | `
	B[5] = ` |____/  `

	C := make([]string, 8)
	C[0] = `   _____  `
	C[1] = `  / ____| `
	C[2] = ` | |      `
	C[3] = ` | |      `
	C[4] = ` | |____  `
	C[5] = `  \_____| `

	D := make([]string, 8)
	D[0] = `  _____   `
	D[1] = ` |  __ \  `
	D[2] = ` | |  | | `
	D[3] = ` | |  | | `
	D[4] = ` | |__| | `
	D[5] = ` |_____/  `

	E := make([]string, 8)
	E[0] = `  ______  `
	E[1] = ` |  ____| `
	E[2] = ` | |__    `
	E[3] = ` |  __|   `
	E[4] = ` | |____  `
	E[5] = ` |______| `

	F := make([]string, 8)
	F[0] = `  ______  `
	F[1] = ` |  ____| `
	F[2] = ` | |__    `
	F[3] = ` |  __|   `
	F[4] = ` | |      `
	F[5] = ` |_|      `

	G := make([]string, 8)
	G[0] = `   _____  `
	G[1] = `  / ____| `
	G[2] = ` | |  __  `
	G[3] = ` | | |_ | `
	G[4] = ` | |__| | `
	G[5] = `  \_____| `

	H := make([]string, 8)
	H[0] = `  _    _  `
	H[1] = ` | |  | | `
	H[2] = ` | |__| | `
	H[3] = ` |  __  | `
	H[4] = ` | |  | | `
	H[5] = ` |_|  |_| `

	I := make([]string, 8)
	I[0] = `  _____  `
	I[1] = ` |_   _| `
	I[2] = `   | |   `
	I[3] = `   | |   `
	I[4] = `  _| |_  `
	I[5] = ` |_____| `

	J := make([]string, 8)
	J[0] = `       _  `
	J[1] = `      | | `
	J[2] = `      | | `
	J[3] = `  _   | | `
	J[4] = ` | |__| | `
	J[5] = `  \____/  `

	K := make([]string, 8)
	K[0] = `  _   _ `
	K[1] = ` | |/ / `
	K[2] = ` | ' /  `
	K[3] = ` |  <   `
	K[4] = ` | . \  `
	K[5] = ` |_|\_\ `

	L := make([]string, 8)
	L[0] = `  __      `
	L[1] = ` | |      `
	L[2] = ` | |      `
	L[3] = ` | |      `
	L[4] = ` | |____  `
	L[5] = ` |______| `

	M := make([]string, 8)
	M[0] = ` __    __ `
	M[1] = ` |  \/  | `
	M[2] = ` | \  / | `
	M[3] = ` | |\/| | `
	M[4] = ` | |  | | `
	M[5] = ` |_|  |_| `

	N := make([]string, 8)
	N[0] = `  _   _  `
	N[1] = ` | \ | | `
	N[2] = ` |  \| | `
	N[3] = ` | .   | `
	N[4] = ` | |\  | `
	N[5] = ` |_| \_| `

	O := make([]string, 8)
	O[0] = `   ____   `
	O[1] = `  / __ \  `
	O[2] = ` | |  | | `
	O[3] = ` | |  | | `
	O[4] = ` | |__| | `
	O[5] = `  \____/  `

	P := make([]string, 8)
	P[0] = ` ______   `
	P[1] = ` |  __ \  `
	P[2] = ` | |__) | `
	P[3] = ` |  ___/  `
	P[4] = ` | |      `
	P[5] = ` |_|      `

	Q := make([]string, 8)
	Q[0] = `   ____   `
	Q[1] = `  / __ \  `
	Q[2] = ` | |  | | `
	Q[3] = ` | |  | | `
	Q[4] = ` | |__| | `
	Q[5] = `  \___\_\ `

	R := make([]string, 8)
	R[0] = ` ______   `
	R[1] = ` |  __ \  `
	R[2] = ` | |__) | `
	R[3] = ` |  _  /  `
	R[4] = ` | | \ \  `
	R[5] = ` |_|  \_\ `

	S := make([]string, 8)
	S[0] = `   _____  `
	S[1] = `  / ____| `
	S[2] = ` | (___   `
	S[3] = `  \___ \  `
	S[4] = `  ____) | `
	S[5] = ` |_____/  `

	T := make([]string, 8)
	T[0] = `  _______  `
	T[1] = ` |__   __| `
	T[2] = `    | |    `
	T[3] = `    | |    `
	T[4] = `    | |    `
	T[5] = `    |_|    `

	U := make([]string, 8)
	U[0] = `  _    _  `
	U[1] = ` | |  | | `
	U[2] = ` | |  | | `
	U[3] = ` | |  | | `
	U[4] = ` | |__| | `
	U[5] = `  \____/  `

	V := make([]string, 8)
	V[0] = `  _      _  `
	V[1] = ` \ \    / / `
	V[2] = `  \ \  / /  `
	V[3] = `   \ \/ /   `
	V[4] = `    \  /    `
	V[5] = `     \/     `

	W := make([]string, 8)
	W[0] = ` __          __ `
	W[1] = ` \ \        / / `
	W[2] = `  \ \  /\  / /  `
	W[3] = `   \ \/  \/ /   `
	W[4] = `    \  /\  /    `
	W[5] = `     \/  \/     `

	X := make([]string, 8)
	X[0] = ` _     _  `
	X[1] = ` \ \ / /  `
	X[2] = `  \ V /   `
	X[3] = `   > <    `
	X[4] = `  / . \   `
	X[5] = ` /_/ \_\  `

	Y := make([]string, 8)
	Y[0] = `  __    __  `
	Y[1] = ` \ \   / /  `
	Y[2] = `  \ \_/ /   `
	Y[3] = `   \   /    `
	Y[4] = `    | |     `
	Y[5] = `    |_|     `

	Z := make([]string, 8)
	Z[0] = `  ______  `
	Z[1] = ` |___  /  `
	Z[2] = `    / /   `
	Z[3] = `   / /    `
	Z[4] = `  / /__   `
	Z[5] = ` /_____|  `

	BANG := make([]string, 8)
	BANG[0] = `  _    `
	BANG[1] = ` |\_\  `
	BANG[2] = ` | | | `
	BANG[3] = ` | | | `
	BANG[4] = `  \|_| `
	BANG[5] = ` |\_\  `
	BANG[6] = `  \|_| `
	BANG[7] = `       `

	PERIOD := make([]string, 8)
	PERIOD[0] = `         `
	PERIOD[1] = `         `
	PERIOD[2] = `         `
	PERIOD[3] = `         `
	PERIOD[4] = `  _      `
	PERIOD[5] = ` |\_\    `
	PERIOD[6] = `  \|_|   `
	PERIOD[7] = `         `

	DASH := make([]string, 8)
	DASH[0] = `             `
	DASH[1] = `             `
	DASH[2] = `  ______     `
	DASH[3] = ` |\______\   `
	DASH[4] = `  \|______|  `
	DASH[5] = `             `
	DASH[6] = `             `
	DASH[7] = `             `

	UNDERSCORE := make([]string, 8)
	UNDERSCORE[0] = `             `
	UNDERSCORE[1] = `             `
	UNDERSCORE[2] = `             `
	UNDERSCORE[3] = `             `
	UNDERSCORE[4] = `  ______     `
	UNDERSCORE[5] = ` |\______\   `
	UNDERSCORE[6] = `  \|______|  `
	UNDERSCORE[7] = `             `

	QUESTIONMARK := make([]string, 8)
	QUESTIONMARK[0] = `  ______    `
	QUESTIONMARK[1] = ` |\______\  `
	QUESTIONMARK[2] = ` | |  __  | `
	QUESTIONMARK[3] = `  \|_| | |  `
	QUESTIONMARK[4] = `     \/_/   `
	QUESTIONMARK[5] = `    |\_\    `
	QUESTIONMARK[6] = `     \|_|   `
	QUESTIONMARK[7] = `            `

	SPACE := make([]string, 8)
	SPACE[0] = `        `
	SPACE[1] = `        `
	SPACE[2] = `        `
	SPACE[3] = `        `
	SPACE[4] = `        `
	SPACE[5] = `        `
	SPACE[6] = `        `
	SPACE[7] = `        `

	BLANK := make([]string, 8)
	BLANK[0] = `   `
	BLANK[1] = `   `
	BLANK[2] = `   `
	BLANK[3] = `   `
	BLANK[4] = `   `
	BLANK[5] = `   `
	BLANK[6] = `   `
	BLANK[7] = `   `

	letterMap := make(map[string][]string)

	letterMap["A"] = A
	letterMap["B"] = B
	letterMap["C"] = C
	letterMap["D"] = D
	letterMap["E"] = E
	letterMap["F"] = F
	letterMap["G"] = G
	letterMap["H"] = H
	letterMap["I"] = I
	letterMap["J"] = J
	letterMap["K"] = K
	letterMap["L"] = L
	letterMap["M"] = M
	letterMap["N"] = N
	letterMap["O"] = O
	letterMap["P"] = P
	letterMap["Q"] = Q
	letterMap["R"] = R
	letterMap["S"] = S
	letterMap["T"] = T
	letterMap["U"] = U
	letterMap["V"] = V
	letterMap["W"] = W
	letterMap["X"] = X
	letterMap["Y"] = Y
	letterMap["Z"] = Z

	letterMap["-"] = DASH
	letterMap["_"] = UNDERSCORE
	letterMap["!"] = BANG
	letterMap["."] = PERIOD
	letterMap[" "] = SPACE
	letterMap["?"] = QUESTIONMARK
	if letterMap[str] != nil {
		return letterMap[str]
	} else {
		ltr := BLANK
		ltr[6] = fmt.Sprintf(" %s ", str)
		return ltr
	}
}

func (this *Admin) GetWidth() (int, error) {
	cmd := exec.Command("tput", "cols")
	w, _ := cmd.Output()
	buf := bytes.NewBuffer(w)
	return strconv.Atoi(strings.TrimSpace(buf.String()))
}

func Blockify(str string, w int) string {
	strs := strings.Split(strings.ToUpper(str), "")
	res := make([]string, 8)
	for _, s := range strs {
		var ltr []string
		if s == "\r\n" {
			r := strings.Join(res, "\r\n")
			res = make([]string, 8)
			res[0] = fmt.Sprintf("%s\r\n", r)
			continue
		}
		ltr = getLetter(s)
		if len(res[1])+len(ltr[1]) > w {
			r := strings.Join(res, "\r\n")
			res = make([]string, 8)
			res[0] = fmt.Sprintf("%s\r\n", r)
		}
		for i := 0; i < 8; i++ {
			res[i] = fmt.Sprintf("%s%s", res[i], ltr[i])
		}
	}
	return strings.Join(res, "\r\n")
}

func (this *Admin) PrintBlocks(str string) {
	w, err := this.GetWidth()
	if err != nil {
		w = 80
	}
	this.conn.Write([]byte(fmt.Sprintf(Blockify(str, w))))
}

func (this *Admin) MegaFail() {
	this.PrintBlocks("FAIL!")
}

func (this *Admin) NoGo() {
	this.PrintBlocks("no go!")
}

func (this *Admin) BadGo() {
	this.PrintBlocks("Bad go!")
}
