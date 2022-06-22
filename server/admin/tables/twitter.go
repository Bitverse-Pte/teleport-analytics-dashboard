package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"time"
)

func GetTwitterAccountTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Name", "name", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Followers Count", "followersCount", db.Int)
	info.AddField("Tweet Count", "tweetCount", db.Int)
	info.AddField("Access Expired", "expiresAt", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02 15:04:05", value.Value)
		if err != nil {
			return ""
		}
		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	})
	info.SetTable("TwitterAccount").SetTitle("Twitter Account Manager").SetDescription("")
	return
}

func GetTwitterAccountRealTime(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().HideEditButton()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Name", "name", db.Varchar).
		FieldJoin(types.Join{
			Table:     "TwitterAccount",
			Field:     "twitterAccountId",
			JoinField: "accountId",
		}).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Date", "date", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02 15:04:05", value.Value)
		if err != nil {
			return ""
		}
		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	})
	info.AddField("Followers Count", "followersCount", db.Int)
	info.AddField("Following Count", "followingCount", db.Int)
	info.AddField("Tweet Count", "tweetCount", db.Int)
	info.AddField("List Count", "listedCount", db.Int)
	info.SetTable("TwitterAccountRealTimeStat").SetTitle("Twitter Account Real Time Manager").SetDescription("")
	return
}

func GetTwitterAccountDailyLog(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Name", "name", db.Varchar).
		FieldJoin(types.Join{
			Table:     "TwitterAccount",
			Field:     "twitterAccountId",
			JoinField: "accountId",
		}).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Date", "date", db.Time)
	info.AddField("Followers Count", "followersCount", db.Int)
	info.AddField("New Followers", "newFollowersCount", db.Int)
	info.AddField("Tweet Count", "tweetCount", db.Int)
	info.AddField("New Tweet Count", "newTweetCount", db.Int)
	info.SetTable("TwitterAccountDailyStat").SetTitle("Twitter Manager").SetDescription("")
	formList := t.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldMust().FieldHideWhenCreate()
	formList.AddField("Twitter Account ID", "twitterAccountId", db.Varchar, form.Text).FieldDisplayButCanNotEditWhenUpdate()
	formList.AddField("Followers Count", "followersCount", db.Int, form.Number)
	formList.AddField("New Followers", "newFollowersCount", db.Int, form.Number)
	formList.AddField("Tweet Count", "tweetCount", db.Int, form.Number)
	formList.AddField("New Tweet Count", "newTweetCount", db.Int, form.Number)
	formList.SetTable("TwitterAccountDailyStat").SetTitle("Twitter Manager").SetDescription("")
	return
}
