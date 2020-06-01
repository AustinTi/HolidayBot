package commands

import "github.com/AustinTi/HolidayBot/db"

func setDaily(p Params) bool {
	if len(p.Args) < 3 {
		p.Args = append(p.Args, "")
	}

	var newDaily bool
	switch p.Args[2] {
	case "on", "true":
		newDaily = true
	case "off", "false":
		newDaily = false
	default:
		Errors(p.Client, p.Message, p.Guild, "SYN", SetDaily)
		return false
	}

	go db.UpdateDaily(p.Guild.ID, newDaily)

	return true
}
