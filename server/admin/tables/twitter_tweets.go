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

func GetTweetsTable(ctx *context.Context) (t table.Table) {
	t = table.NewDefaultTable(table.DefaultConfig())
	info := t.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue().
		HideEditButton().
		SetSortField("createdAt").
		SetSortDesc()
	info.AddField("ID", "id", db.Int).FieldSortable().FieldHide()
	info.AddField("Created At", "createdAt", db.Time).FieldDisplay(func(value types.FieldModel) interface{} {
		t, err := time.Parse("2006-01-02 15:04:05", value.Value)
		if err != nil {
			return ""
		}

		return t.In(time.FixedZone("GMT", 8*3600)).Format("2006-01-02 15:04:05")
	}).FieldSortable()
	info.AddField("Link", "tweetId", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value != "" {
			// TODO: username
			return template.Default().Link().
				SetURL(fmt.Sprintf("https://twitter.com/TeleportChain/status/%s", value.Value)).
				SetContent(template.HTML("link")).
				SetAttributes("target=_blank").
				GetContent()
		}
		return "-"
	})
	info.AddField("Text", "text", db.Varchar)
	info.AddField("Impressions", "impressions", db.Int).FieldSortable()
	info.AddField("Engagement", "engagement", db.Int).FieldSortable()
	info.AddField("Retweets", "retweets", db.Int).FieldSortable()
	info.AddField("QuoteTweets", "quoteTweets", db.Int).FieldSortable()
	info.AddField("Likes", "likes", db.Int).FieldSortable()
	info.AddField("Replies", "replies", db.Int).FieldSortable()
	info.AddField("UserProfileClicks", "userProfileClicks", db.Int).FieldSortable()
	info.AddField("UrlLinkClicks", "urlLinkClicks", db.Int).FieldSortable()
	info.AddField("Media Views", "mediaViews", db.Int).FieldSortable()
	info.SetTable("Tweet").SetTitle("Tweet Manager").SetDescription("")
	return
}
