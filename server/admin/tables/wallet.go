package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/teleport-network/teleport-analytics-dashboard/model"
	"github.com/teleport-network/teleport-analytics-dashboard/server/admin/models"
	"strconv"
	"strings"
	"time"
)

func GetWalletDailyDataTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().SetSortField("date")
	info.AddActionIconButton("Init Data", action.Ajax("/admin/initWalletData", func(ctx *context.Context) (success bool, msg string, data interface{}) {
		return true, "success", nil
	}))
	info.AddButton("Init Data", icon.Database, action.Ajax("/admin/audit",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			initData()
			return true, "Success", ""
		}))
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Date", "date", db.Time).
		FieldDisplay(func(value types.FieldModel) interface{} {
			t, err := time.Parse("2006-01-02T15:04:05Z", value.Value)
			if err != nil {
				return ""
			}
			return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
		}).FieldSortable().SetSortDesc()
	info.AddField("Download(Total)", "download_count", db.Int)
	info.AddField("Download(New)", "new_download_count", db.Int)
	info.AddField("Wallet Type", "type", db.Varchar).
		FieldFilterable(types.FilterType{FormType: form.SelectSingle}).
		FieldFilterOptions(getWalletTypeOptions())
	info.SetTable("WalletDaily").SetTitle("Wallet Daily Data Manager").SetDescription("")

	formList := t.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldMust().FieldHideWhenCreate()
	formList.AddField("Date", "date", db.Time, form.Datetime)
	//formList.AddField("Download(Total)", "download_count", db.Int, form.Number)
	formList.AddField("Download(New)", "new_download_count", db.Int, form.Number)
	formList.AddField("Wallet Type", "type", db.Varchar, form.SelectSingle).
		FieldOptions(getWalletTypeOptions()).
		FieldDefault(model.WalletTypeExtension.String()).FieldMust()
	formList.SetTable("WalletDaily").SetTitle("Twitter Manager").SetDescription("")
	formList.SetPreProcessFn(func(values form2.Values) form2.Values {
		var list []model.WalletDaily
		if err := models.GetDB().Where("date < ?", values.Get("date")).Find(&list).Error; err != nil {
			return values
		}
		var total int
		for _, one := range list {
			total += one.NewDownloadCount
		}
		new, _ := strconv.Atoi(values.Get("new_download_count"))
		total += new
		values.Add("download_count", strconv.Itoa(total))
		return values
	})
	return
}

func getWalletTypeOptions() types.FieldOptions {
	return types.FieldOptions{
		{Text: model.WalletTypeExtension.String(), Value: model.WalletTypeExtension.String()},
	}
}

func initData() {
	models.GetDB().Unscoped().Delete(&model.WalletDaily{})
	items := strings.Split(data, "\n")
	var total int
	for _, one := range items {
		keys := strings.Split(one, "\t")
		t, _ := time.Parse("2006/1/2", keys[0])
		install, _ := strconv.Atoi(keys[1])
		total += install
		err := models.GetDB().Create(&model.WalletDaily{
			Date:             t,
			Type:             model.WalletTypeExtension,
			DownloadCount:    total,
			NewDownloadCount: install,
		}).Error
		if err != nil {
			logger.Warn(err.Error())
		}
	}
}

var data = `2022/3/23	8
2022/3/24	12
2022/3/25	5
2022/3/26	2
2022/3/27	3
2022/3/28	4
2022/3/29	6
2022/3/30	3
2022/3/31	0
2022/4/1	3
2022/4/2	5
2022/4/3	3
2022/4/4	5
2022/4/5	29
2022/4/6	45
2022/4/7	22
2022/4/8	7
2022/4/9	3
2022/4/10	8
2022/4/11	9
2022/4/12	5
2022/4/13	6
2022/4/14	277
2022/4/15	53
2022/4/16	45
2022/4/17	16
2022/4/18	16
2022/4/19	9
2022/4/20	6
2022/4/21	134
2022/4/22	48
2022/4/23	13
2022/4/24	10
2022/4/25	16
2022/4/26	10
2022/4/27	17
2022/4/28	0
2022/4/29	5
2022/4/30	0
2022/5/1	2
2022/5/2	2
2022/5/3	2
2022/5/4	0
2022/5/5	4
2022/5/6	0
2022/5/7	2
2022/5/8	1
2022/5/9	2
2022/5/10	1
2022/5/11	1
2022/5/12	2
2022/5/13	2
2022/5/14	1
2022/5/15	3
2022/5/16	1
2022/5/17	2
2022/5/18	2
2022/5/19	9
2022/5/20	10
2022/5/21	121
2022/5/22	22
2022/5/23	21
2022/5/24	12
2022/5/25	9
2022/5/26	9
2022/5/27	8
2022/5/28	5
2022/5/29	6
2022/5/30	3
2022/5/31	4
2022/6/1	6
2022/6/2	50
2022/6/3	30
2022/6/4	10
2022/6/5	9
2022/6/6	5
2022/6/7	8
2022/6/8	9
2022/6/9	11
2022/6/10	11
2022/6/11	6
2022/6/12	1
2022/6/13	2
2022/6/14	3
2022/6/15	6
2022/6/16	1
2022/6/17	21
2022/6/18	4533
2022/6/19	1526
2022/6/20	791
2022/6/21	132
2022/6/22	0`
