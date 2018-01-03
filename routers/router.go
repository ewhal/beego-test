// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/ewhal/pantsu/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/comments",
			beego.NSInclude(
				&controllers.CommentsController{},
			),
		),

		beego.NSNamespace("/user_follows",
			beego.NSInclude(
				&controllers.UserFollowsController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),

		beego.NSNamespace("/user_uploads_old",
			beego.NSInclude(
				&controllers.UserUploadsOldController{},
			),
		),

		beego.NSNamespace("/hydra_oauth2_access",
			beego.NSInclude(
				&controllers.HydraOauth2AccessController{},
			),
		),

		beego.NSNamespace("/sukebei_comments",
			beego.NSInclude(
				&controllers.SukebeiCommentsController{},
			),
		),

		beego.NSNamespace("/sukebei_files",
			beego.NSInclude(
				&controllers.SukebeiFilesController{},
			),
		),

		beego.NSNamespace("/sukebei_notifications",
			beego.NSInclude(
				&controllers.SukebeiNotificationsController{},
			),
		),

		beego.NSNamespace("/hydra_oauth2_code",
			beego.NSInclude(
				&controllers.HydraOauth2CodeController{},
			),
		),

		beego.NSNamespace("/hydra_oauth2_oidc",
			beego.NSInclude(
				&controllers.HydraOauth2OidcController{},
			),
		),

		beego.NSNamespace("/hydra_oauth2_refresh",
			beego.NSInclude(
				&controllers.HydraOauth2RefreshController{},
			),
		),

		beego.NSNamespace("/accounts",
			beego.NSInclude(
				&controllers.AccountsController{},
			),
		),

		beego.NSNamespace("/files",
			beego.NSInclude(
				&controllers.FilesController{},
			),
		),

		beego.NSNamespace("/notifications",
			beego.NSInclude(
				&controllers.NotificationsController{},
			),
		),

		beego.NSNamespace("/scrape",
			beego.NSInclude(
				&controllers.ScrapeController{},
			),
		),

		beego.NSNamespace("/sukebei_activities",
			beego.NSInclude(
				&controllers.SukebeiActivitiesController{},
			),
		),

		beego.NSNamespace("/sukebei_scrape",
			beego.NSInclude(
				&controllers.SukebeiScrapeController{},
			),
		),

		beego.NSNamespace("/activities",
			beego.NSInclude(
				&controllers.ActivitiesController{},
			),
		),

		beego.NSNamespace("/sukebei_torrent_reports",
			beego.NSInclude(
				&controllers.SukebeiTorrentReportsController{},
			),
		),

		beego.NSNamespace("/hydra_client",
			beego.NSInclude(
				&controllers.HydraClientController{},
			),
		),

		beego.NSNamespace("/torrent_reports",
			beego.NSInclude(
				&controllers.TorrentReportsController{},
			),
		),

		beego.NSNamespace("/torrents",
			beego.NSInclude(
				&controllers.TorrentsController{},
			),
		),

		beego.NSNamespace("/sukebei_torrents",
			beego.NSInclude(
				&controllers.SukebeiTorrentsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
