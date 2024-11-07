package watoken

import (
	"os"
	"testing"
)

func TestEncode(t *testing.T) {
	privkey := os.Getenv("PRIVATEKEY")
	str, _ := EncodeforHours("62881022522920", "Helpdesk Bayar-in", privkey, 43830)
	println(str)
	//atr, _ := DecodeGetId("", str)
	//println(atr)

}
