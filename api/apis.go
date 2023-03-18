package api

import (
	"github.com/go-chi/chi/v5"
)

func Route() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/v0/api", Apis())
	return r
}

func Apis() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/account/login", ApiAccountLogin)
	r.Mount("/account/logout", ApiAccountLogout)
	r.Mount("/account/register", ApiAccountRegister)
	r.Mount("/account/register/check-email", ApiAccountRegisterCheckEmail)
	r.Mount("/account/register/check-name", ApiAccountRegisterCheckName)

	r.Mount("/article/add", ApiAddArticle)
	r.Mount("/article/set", ApiUpdateArticle)
	r.Mount("/article/del", ApiDeleteArticle)
	r.Mount("/article/get", ApiGetArticle)

	r.Mount("/articles/get", ApiGetArticles)

	r.Mount("/reviews/get", ApiGetReviews)
	r.Mount("/review/get", ApiGetReview)
	r.Mount("/review/set", ApiUpdateReview)

	r.Mount("/profile/get", ApiGetUser)
	r.Mount("/profile/auth", ApiAuthenticate)

	r.Mount("/admin/dashboard", ApiAdminDashboard)
	return r
}
