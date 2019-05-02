package kit

import "strconv"

var DbKit = dbKit{}

type dbKit struct {
}

func (this dbKit) Set0At(field string, pos int) string {
	return " " + field + "=(" + field + " & ~(1 << " + strconv.Itoa(pos) + ")" + ") "
}

func (this dbKit) Set1At(field string, pos int) string {
	return " " + field + "=(" + field + " | (1 << " + strconv.Itoa(pos) + ")" + ") "
}

func ss() {
	_ := "update account set " + DbKit.Set1At("Status", 1)
}
