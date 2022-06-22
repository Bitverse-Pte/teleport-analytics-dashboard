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

func GetDiscordChannelTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Varchar)
	info.AddField("Site", "discordGuildId", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		fmt.Println("debug joy", value.Row["id"])
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
		t, err := time.Parse("2006-01-02 15:04:05", value.Value)
		if err != nil {
			return ""
		}

		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	})
	info.AddField("Name", "name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Type", "type", db.Varchar).FieldFilterable(types.FilterType{FormType: form.SelectSingle, Placeholder: "Select Type"}).FieldFilterOptions(types.FieldOptions{})
	info.SetTable("DiscordChannel").SetTitle("Channel Manager").SetDescription("")
	return
}

func GetDiscordChannelRealTimeTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Name", "name", db.Varchar).
		FieldJoin(types.Join{
			Table:     "DiscordChannel",
			Field:     "discordChannelId",
			JoinField: "id",
		}).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Created At", "createdAt", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02 15:04:05", value.Value)
		if err != nil {
			return ""
		}

		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	})
	info.AddField("Total Member", "totalMemberCount", db.Int)
	info.AddField("Online Member", "onlineMemberCount", db.Int)
	info.SetTable("DiscordGuildChannelStat").SetTitle("Guild Real Time Manager").SetDescription("")
	return
}

func GetDiscordGuildChannelDailyTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Name", "name", db.Varchar).
		FieldJoin(types.Join{
			Table:     "DiscordChannel",
			Field:     "discordChannelId",
			JoinField: "id",
		}).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Date", "date", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02 15:04:05", value.Value)
		if err != nil {
			return ""
		}

		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	})
	info.AddField("Start Total Member", "startTotalMemberCount", db.Int)
	info.AddField("Start Online Member", "startOnlineMemberCount", db.Int)
	info.AddField("End Total Member", "endTotalMemberCount", db.Int)
	info.AddField("End Online Member", "endOnlineMemberCount", db.Int)
	info.AddField("Highest Total Member", "highestTotalMemberCount", db.Int)
	info.AddField("Lowest Total Member", "lowestTotalMemberCount", db.Int)
	info.SetTable("DiscordGuildChannelDailyStat").SetTitle("Channel Daily Manager").SetDescription("")
	return
}
