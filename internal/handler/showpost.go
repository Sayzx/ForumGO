package handler

// func ShowPostsHandler(w http.ResponseWriter, r *http.Request) {
// 	// le lien final c'est locahost/showtopics?id=1/showposts?idpost=ID DU POST
// 	// on récupère l'id du post
// 	postid := r.URL.Query().Get("idpost")
// 	// on récupère le post
// 	post := api.GetPostById(postid)
// 	if post == nil {
// 		http.Error(w, "Could not fetch post", http.StatusInternalServerError)
// 		return
// 	}
// }
