package tables

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"time"
)

func GetDiscordGuildTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("Guild ID", "id", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			fmt.Println("debug joy", value.Value)
			return value.Value
		})
	info.SetTable("DiscordGuild").SetTitle("DiscordGuild Manager").SetDescription("")
	return
}

func GetDiscordGuildDailyTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Site", "discordGuildId", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value != "" {
			return template.Default().Link().
				SetURL(fmt.Sprintf("https://discord.com/channels/%s", value.Value)).
				SetContent(template.HTML("link")).
				SetAttributes("target=_blank").
				GetContent()
		}
		return "-"
	})
	info.AddField("Date", "date", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02T15:04:05Z", value.Value)
		if err != nil {
			return ""
		}

		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	})
	info.AddField("Start Total Member", "startTotalMemberCount", db.Int)
	info.AddField("Start Online Member", "startOnlineMemberCount", db.Int)
	info.AddField("End Total Member", "endTotalMemberCount", db.Int)
	info.AddField("End Online Member", "endOnlineMemberCount", db.Int)
	info.AddField("Highest Online Member", "highestOnlineMemberCount", db.Int)
	info.AddField("Lowest Online Member", "lowestOnlineMemberCount", db.Int)
	info.SetTable("DiscordGuildDailyStat").SetTitle("Guild Daily Manager").SetDescription("")
	return
}

func GetDiscordGuildRealTimeTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Site", "discordGuildId", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value != "" {
			return template.Default().Link().
				SetURL(fmt.Sprintf("https://discord.com/channels/%s", value.Value)).
				SetContent(template.HTML("link")).
				SetAttributes("target=_blank").
				GetContent()
		}
		return "-"
	})

	info.AddField("Created At", "createdAt", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02T15:04:05Z", value.Value)
		if err != nil {
			return ""
		}

		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	})
	info.AddField("Total Member", "totalMemberCount", db.Int)
	info.AddField("Online Member", "onlineMemberCount", db.Int)
	info.SetTable("DiscordGuildStat").SetTitle("Guild Real Time Manager").SetDescription("")
	return
}
