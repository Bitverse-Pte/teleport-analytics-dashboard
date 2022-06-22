package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"time"
)

func GetTelegramGroupMemberTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Chat ID", "groupId", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("User ID", "userId", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Join At", "joinAt", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02 15:04:05", value.Value)
		if err != nil {
			return ""
		}

		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	})
	info.AddField("Message Count", "messageCount", db.Int)
	info.AddField("Last Seen", "lastSeen", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02 15:04:05", value.Value)
		if err != nil {
			return ""
		}

		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	})
	info.AddField("Active Days", "activeDays", db.Int)
	info.SetTable("TelegramChatMember").SetTitle("Telegram Member Manager").SetDescription("")
	return
}
