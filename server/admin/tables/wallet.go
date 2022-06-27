package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetWalletDailyDataTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Date", "date", db.Varchar).FieldFilterable(types.FilterType{FormType: form.DateRange})
	info.AddField("Download(Total)", "download_count", db.Int)
	info.AddField("Download(New)", "download_count_new", db.Int)
	info.SetTable("WalletDaily").SetTitle("Wallet Daily Data Manager").SetDescription("")

	formList := t.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldMust().FieldHideWhenCreate()
	formList.AddField("Date", "date", db.Time, form.Datetime)
	formList.AddField("Download(Total)", "download_count", db.Int, form.Number)
	formList.AddField("Download(New)", "download_new", db.Int, form.Number)
	formList.SetTable("WalletDaily").SetTitle("Twitter Manager").SetDescription("")
	return
}
