package tables

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"html/template"
	"strings"
	"time"
)

func GetTelegramGroupTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Chat ID", "chatId", db.Decimal).FieldHide()
	info.AddField("Photo", "photo", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "" {
			return ""
		} else {
			return template2.Default().Image().
				SetSrc(template.HTML(value.Value)).
				SetWidth("100").
				SetHeight("100").
				WithModal().GetContent()
		}
	})
	info.AddField("Group Name", "title", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "" {
			return "-"
		}
		return template2.Default().Link().
			SetURL(fmt.Sprintf("tg://resolve?domain=%s", strings.ReplaceAll(value.Value, " ", ""))).
			SetContent(template.HTML(value.Value)).
			SetAttributes("target=_blank").
			GetContent()
	})

	info.AddField("Invite Link", "invite_link", db.Varchar)
	info.SetTable("TelegramGroup").SetTitle("Telegram Chat Group Manager").SetDescription("")
	return
}

func GetTelegramGroupDailyTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Chat ID", "groupId", db.Decimal).FieldHide().FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Group Name", "title", db.Varchar).FieldJoin(types.Join{
		Table:     "TelegramGroup",
		Field:     "groupId",
		JoinField: "chatId",
	}).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "" {
			return "-"
		}
		return template2.Default().Link().
			SetURL(fmt.Sprintf("tg://resolve?domain=%s", strings.ReplaceAll(value.Value, " ", ""))).
			SetContent(template.HTML(value.Value)).
			SetAttributes("target=_blank").
			GetContent()
	})
	info.AddField("Date", "date", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02T15:04:05Z", value.Value)
		if err != nil {
			return ""
		}

		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	})
	info.AddField("New Member", "newMemberCount", db.Int)
	info.AddField("Message Count", "messageCount", db.Int)
	info.AddField("Total Member Count", "totalMemberCount", db.Int)
	info.AddField("Active Member Count", "activeMemberCount", db.Int)
	info.SetTable("TelegramGroupDailyStat").SetTitle("Telegram Chat Group Daily Manager").SetDescription("")
	return
}

func GetTelegramGroupRealTimeTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Chat ID", "groupId", db.Decimal).FieldHide()
	info.AddField("Group Name", "title", db.Varchar).FieldJoin(types.Join{
		Table:     "TelegramGroup",
		Field:     "groupId",
		JoinField: "chatId",
	}).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "" {
				return "-"
			}
			return template2.Default().Link().
				SetURL(fmt.Sprintf("tg://resolve?domain=%s", strings.ReplaceAll(value.Value, " ", ""))).
				SetContent(template.HTML(value.Value)).
				SetAttributes("target=_blank").
				GetContent()
		})
	info.AddField("Date", "date", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02T15:04:05Z", value.Value)
		if err != nil {
			return ""
		}

		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	})
	info.AddField("New Member", "newMemberCount", db.Int)
	info.AddField("Message Count", "messageCount", db.Int)
	info.AddField("Total Member Count", "totalMemberCount", db.Int)
	info.AddField("Active Member Count", "activeMemberCount", db.Int)
	info.SetTable("TelegramGroupStats").SetTitle("Real Time Manager").SetDescription("")
	return
}
