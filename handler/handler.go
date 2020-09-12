package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"insta_graph/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var loginTmpl = template.Must(template.ParseFiles("templates/base.html", "templates/login.html"))

func Login(w http.ResponseWriter, r *http.Request) {
	loginTmpl.Execute(w, map[string]interface{}{
		"app_id":   os.Getenv("FACEBOOK_APP_ID"),
		"base_url": os.Getenv("BASE_URL"),
		//"app_secret": os.Getenv("FACEBOOK_APP_SECRET")
	})
}

func Callback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	// 	https://graph.facebook.com/v8.0/oauth/access_token?
	//    client_id={app-id}
	//    &redirect_uri={redirect-uri}
	//    &client_secret={app-secret}
	//    &code={code-parameter}
	resp, err := http.PostForm("https://graph.facebook.com/v8.0/oauth/access_token",
		url.Values{
			"client_id":     {os.Getenv("FACEBOOK_APP_ID")},
			"client_secret": {os.Getenv("FACEBOOK_APP_SECRET")},
			"redirect_uri":  {os.Getenv("BASE_URL") + "/callback"},
			"code":          {code},
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(body))
	var auth models.Authorization
	err = json.Unmarshal(body, &auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Get the User's Pages
	fbr, err := http.Get("https://graph.facebook.com/v8.0/me/accounts?access_token=" + auth.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fbrBody, err := ioutil.ReadAll(fbr.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var fba models.FBAccount
	err = json.Unmarshal(fbrBody, &fba)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Get the Page's Instagram Business Account
	igr, err := http.Get("https://graph.facebook.com/v8.0/" + fba.Data[0].ID + "?fields=instagram_business_account&access_token=" + auth.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	igrBody, err := ioutil.ReadAll(igr.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var iga models.IGAccount
	err = json.Unmarshal(igrBody, &iga)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(iga)

	http.SetCookie(w, &http.Cookie{
		Name:   "authtoken",
		Value:  auth.AccessToken + "@" + iga.IGBAccount.ID,
		Path:   "/",
		MaxAge: 60 * 60 * 24,
	})

	http.Redirect(w, r, "/home", http.StatusFound)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "yee")
}
