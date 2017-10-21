package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"govwa/util"
	"govwa/util/middleware"
	"govwa/user"
)

//index and set cookie

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie := util.SetCookieLevel(w, r)
	data := make(map[string]interface{})
	data["level"] = cookie
	data["title"] = "Index"
	data["weburl"] = util.Fullurl

	fmt.Println(r.FormValue("govwa_session"))
	util.SafeRender(w,"template.index", data)
}

func main() {
	
	mw := middleware.New()
	router := httprouter.New()
	userObj := user.New()

	router.ServeFiles("/public/*filepath", http.Dir("public/"))
	router.GET("/", mw.LoggingMiddleware(mw.AuthCheck(indexHandler)))

	userObj.SetRouter(router)

	s := http.Server{
		Addr : ":8082",
		Handler : router,
	}
	fmt.Printf("Server running at port %s\n", s.Addr)
	s.ListenAndServe()

}