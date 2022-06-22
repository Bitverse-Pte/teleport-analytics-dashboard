package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"time"
)

func GetTelegramGroupTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Chat ID", "chatId", db.Decimal).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.SetTable("TelegramGroup").SetTitle("Telegram Chat Group Manager").SetDescription("")
	return
}

func GetTelegramGroupDailyTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Chat ID", "groupId", db.Decimal).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Date", "date", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02 15:04:05", value.Value)
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
	info.AddField("Chat ID", "groupId", db.Decimal).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Date", "date", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02 15:04:05", value.Value)
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
