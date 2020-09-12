package models

// {
// 	"access_token": {access-token},
// 	"token_type": {type},
// 	"expires_in":  {seconds-til-expiration}
// }

type Authorization struct {
	AccessToken string `json:"access_token"`
	Token       string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

// {
// 	"data": [
// 	  {
// 		"access_token": "EAAJjmJ...",
// 		"category": "App Page",
// 		"category_list": [
// 		  {
// 			"id": "2301",
// 			"name": "App Page"
// 		  }
// 		],
// 		"name": "Metricsaurus",
// 		"id": "134895793791914",  // capture the Page ID
// 		"tasks": [
// 		  "ANALYZE",
// 		  "ADVERTISE",
// 		  "MODERATE",
// 		  "CREATE_CONTENT",
// 		  "MANAGE"
// 		]
// 	  }
// 	]
// }

type FBAccount struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}

// {
// 	"instagram_business_account": {
// 	  "id": "17841405822304914"  // Connected IG User ID
// 	},
// 	"id": "134895793791914"  // Facebook Page ID
//   }

type IGAccount struct {
	IGBAccount struct {
		ID string `json:"id"`
	} `json:"instagram_business_account"`
}
