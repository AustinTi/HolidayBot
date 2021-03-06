package events

import (
	"github.com/AustinTi/HolidayBot/db"
	"github.com/AustinTi/HolidayBot/misc"
	"github.com/bwmarrin/discordgo"
)

//GuildDelete event
func GuildDelete(client *discordgo.Session, guild *discordgo.GuildDelete) {
	if guild.Guild.Unavailable {
		return
	}

	dbResult, err := db.GuildFetch(guild.Guild.ID)
	if err != nil {
		return
	}

	deletedGuild := &discordgo.Guild{
		ID:   guild.Guild.ID,
		Name: dbResult.Guildname,
	}

	db.DeleteGuild(guild.Guild)

	misc.Log("", "info", "leave", nil, deletedGuild, "")
}
