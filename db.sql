-- --------------------------------------------------------
-- Hôte:                         C:\Users\aylan\Desktop\Ynov\ForumGO\internal\sql\forum.db
-- Version du serveur:           3.44.0
-- SE du serveur:                
-- HeidiSQL Version:             12.6.0.6765
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES  */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Listage de la structure de la base pour forum
CREATE DATABASE IF NOT EXISTS "forum";
;

-- Listage de la structure de la table forum. comments
CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    postid INTEGER,
    content TEXT,
    owner TEXT,
    createat DATETIME,
    FOREIGN KEY (postid) REFERENCES topics(id)
);

-- Listage des données de la table forum.comments : -1 rows
/*!40000 ALTER TABLE "comments" DISABLE KEYS */;
/*!40000 ALTER TABLE "comments" ENABLE KEYS */;

-- Listage de la structure de la table forum. loginlogs
CREATE TABLE IF NOT EXISTS "loginlogs" (
	"username" VARCHAR(50) NULL DEFAULT NULL,
	"plateform" VARCHAR(50) NULL DEFAULT NULL,
	"datetime" DATETIME NULL DEFAULT NULL
);

-- Listage des données de la table forum.loginlogs : -1 rows
/*!40000 ALTER TABLE "loginlogs" DISABLE KEYS */;
INSERT INTO "loginlogs" ("username", "plateform", "datetime") VALUES
	('flyx_0_', 'Discord', '2024-06-19 16:43:39.9723829 +0200 CEST m=+10.850621901'),
	('hashedPassword@hashedPassword.fr', 'Local', '2024-06-19 16:48:21.8774814 +0200 CEST m=+78.164406001'),
	('aylann.pro@gmail.com', 'Google', '2024-06-20 15:22:10.1283441 +0200 CEST m=+24.794650101'),
	('bouclierbleu39@gmail.com', 'Google', '2024-06-20 15:25:17.8957831 +0200 CEST m=+47.795562101'),
	('flyx_0_', 'Discord', '2024-06-20 15:25:49.4949094 +0200 CEST m=+79.394688401'),
	('topwin.gamerzz@gmail.com', 'Google', '2024-06-24 13:16:02.3991999 +0200 CEST m=+13.362820201'),
	('topwin.gamerzz@gmail.com', 'Google', '2024-06-24 13:48:54.7658288 +0200 CEST m=+23.553901301'),
	('topwin.gamerzz@gmail.com', 'Google', '2024-06-24 13:49:44.8288546 +0200 CEST m=+8.051433201'),
	('topwin.gamerzz@gmail.com', 'Google', '2024-06-24 13:54:44.2155936 +0200 CEST m=+18.791916501'),
	('topwin.gamerzz@gmail.com', 'Google', '2024-06-24 16:14:11.0597231 +0200 CEST m=+11.971507601'),
	('topwin', 'Discord', '2024-06-24 16:14:20.5276411 +0200 CEST m=+21.439425601'),
	('sayzx', 'Discord', '2024-06-25 14:26:57.0459318 +0200 CEST m=+45.441777501'),
	('sayzx', 'Discord', '2024-06-25 14:36:22.2906386 +0200 CEST m=+269.753795401'),
	('sayzx', 'Discord', '2024-06-25 14:44:56.554615 +0200 CEST m=+10.771681101'),
	('sayzx', 'Discord', '2024-06-25 15:38:40.4996069 +0200 CEST m=+1411.298332601'),
	('bouclierbleu39@gmail.com', 'Google', '2024-06-25 15:40:14.0768672 +0200 CEST m=+1504.875592901');
/*!40000 ALTER TABLE "loginlogs" ENABLE KEYS */;

-- Listage de la structure de la table forum. mail
CREATE TABLE IF NOT EXISTS mail (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		email TEXT NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);

-- Listage des données de la table forum.mail : -1 rows
/*!40000 ALTER TABLE "mail" DISABLE KEYS */;
/*!40000 ALTER TABLE "mail" ENABLE KEYS */;

-- Listage de la structure de la table forum. password
CREATE TABLE IF NOT EXISTS password (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		password TEXT NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);

-- Listage des données de la table forum.password : -1 rows
/*!40000 ALTER TABLE "password" DISABLE KEYS */;
/*!40000 ALTER TABLE "password" ENABLE KEYS */;

-- Listage de la structure de la table forum. topics
CREATE TABLE IF NOT EXISTS "topics" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "content" VARCHAR(255) NULL DEFAULT NULL,
    "title" VARCHAR(60) NULL DEFAULT NULL,
    "avatar" VARCHAR(255) NULL DEFAULT NULL,
    "categoryid" INTEGER NULL,
    "tags" VARCHAR(50) NULL DEFAULT NULL,
    "images" VARCHAR(50) NULL DEFAULT NULL,
    "like" INTEGER NULL,
    "dislike" INTEGER NULL,
    "createat" DATETIME NULL
, "owner" VARCHAR(255) NULL DEFAULT NULL);

-- Listage des données de la table forum.topics : -1 rows
/*!40000 ALTER TABLE "topics" DISABLE KEYS */;
INSERT INTO "topics" ("id", "content", "title", "avatar", "categoryid", "tags", "images", "like", "dislike", "createat", "owner") VALUES
	(1, 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent et nibh ac tortor tempus dictum. Fusce pellentesque augue ut purus iaculis, nec porttitor ipsum varius. Sed a velit a tellus consectetur gravida. In ut enim eleifend tellus porta aliquet. Mauris porta orci id velit ultrices, sed dapibus libero interdum. Sed vel finibus libero, sit amet lobortis neque. Sed quis magna at est mattis mattis non sed erat. Fusce facilisis nibh nec mi tristique elementum id in augue. Pellentesque mollis in ex at posuere. Donec tincidunt quam urna, non eleifend nulla tristique quis. Donec vulputate et velit eu aliquam. Donec scelerisque felis dolor, nec consequat mi tempus a. Mauris ut lobortis velit. Vestibulum non turpis lectus.', 'Le Nouveu leak tu turfu wlh', 'https://lh3.googleusercontent.com/a-/ALV-UjXAkIwmWMlij9ywLjGSwTFySqlH56oAE8-KlBm1RzjeYTrouUE=s96-c', 3, 'info', '', 0, 0, NULL, 'topwin.gamerzz@gmail.com'),
	(2, '@xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight ', 'Les noir tah l''epoque', 'https://lh3.googleusercontent.com/a-/ALV-UjXAkIwmWMlij9ywLjGSwTFySqlH56oAE8-KlBm1RzjeYTrouUE=s96-c', 1, 'general', '', 0, 0, NULL, 'topwin.gamerzz@gmail.com'),
	(3, '@xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight ', '@xplknight ', 'https://lh3.googleusercontent.com/a-/ALV-UjXAkIwmWMlij9ywLjGSwTFySqlH56oAE8-KlBm1RzjeYTrouUE=s96-c', 2, 'news', '', 0, 0, NULL, 'topwin.gamerzz@gmail.com'),
	(4, '@xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight ', '@xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight ', 'https://lh3.googleusercontent.com/a-/ALV-UjXAkIwmWMlij9ywLjGSwTFySqlH56oAE8-KlBm1RzjeYTrouUE=s96-c', 4, 'info', '', 0, 0, NULL, 'topwin.gamerzz@gmail.com'),
	(5, '@xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight @xplknight ', 'La canne a sucre des noirs', 'https://lh3.googleusercontent.com/a-/ALV-UjXAkIwmWMlij9ywLjGSwTFySqlH56oAE8-KlBm1RzjeYTrouUE=s96-c', 5, 'general', '', 0, 0, NULL, 'topwin.gamerzz@gmail.com'),
	(20, 'ses mort', 'Manger', 'https://cdn.discordapp.com/avatars/1233487531388047414/6354f8938fdfd057fd830dddd7513182.png', 1, 'general', '', 0, 0, '2024-06-24 16:32:35', 'flyx_0_'),
	(21, 'dev', 'Dev', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.pnga7b3f3c5-009d-48b6-9eed-6d76b718c64d', 2, 'general', '', 0, 0, '2024-06-25 14:27:25', 'sayzx'),
	(22, 'atype CreateTopicData struct {
	LoggedIn bool
	Avatar   string
}type CreateTopicData struct {
	LoggedIn bool
	Avatar   string
}type CreateTopicData struct {
	LoggedIn bool
	Avatar   string
}type CreateTopicData struct {
	LoggedIn bool
	Avatar   string
}', 'aaa', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.pnga7b3f3c5-009d-48b6-9eed-6d76b718c64d', 2, 'bug', '', 0, 0, '2024-06-25 14:27:59', 'sayzx'),
	(23, 'package handler

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"main/internal/api"
	dbsql "main/internal/sql"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const MaxUploadSize = 20 * 1024 * 1024 // 20 MB
const UploadPath = "./web/uploads"

type CreateTopicData struct {
	LoggedIn bool
	Avatar   string
}

func CreateTopicHandler(w http.ResponseWriter, r *http.Request) {
	var data CreateTopicData

	// Tentative de récupération du cookie utilisateur
	cookie, err := r.Cookie("user")
	if err == nil && cookie != nil {
		value, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			log.Println("Error unescaping cookie value:", err)
			http.Error(w, "Error processing cookie", http.StatusBadRequest)
			return
		}

		parts := strings.SplitN(value, ";", 2)
		if len(parts) == 2 {
			data.LoggedIn = true
			data.Avatar = parts[1]
		}
	}

	if !data.LoggedIn {
		// Définir l''avatar par défaut si l''utilisateur n''est pas connecté
		data.Avatar = "https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png"
	}

	// Chargement et exécution du template
	tmpl, err := template.ParseFiles("./web/templates/createtopic.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func AddTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// owner = username by cookie
	username := api.GetUsernameByCookie(r)
	avatar := api.GetAvatarByCookie(r)
	owner := username
	fmt.Println(owner)
	title := r.FormValue("title")
	category := r.FormValue("category")
	tags := r.FormValue("tags")
	content := r.FormValue("content")
	like := 0
	dislike := 0
	createat := api.GetDateAndTime()
	if avatar == "" {
		avatar = "https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png"
	}
	if title == "" || category == "" || tags == "" || content == "" || owner == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Handle image uploads
	err := r.ParseMultipartForm(MaxUploadSize)
	if err != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		return
	}

	var imagePaths []string
	files := r.MultipartForm.File["images"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error opening file", http.StatusInternalServerError)
			log.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		fileType := fileHeader.Header.Get("Content-Type")
		if !strings.HasPrefix(fileType, "image/") {
			http.Error(w, "Invalid file type", http.StatusBadRequest)
			log.Println("Invalid file type:", fileType)
			return
		}

		if fileHeader.Size > MaxUploadSize {
			http.Error(w, "File is too big", http.StatusBadRequest)
			log.Println("File is too big:", fileHeader.Size)
			return
		}

		fileName := filepath.Base(fileHeader.Filename)
		filePath := filepath.Join(UploadPath, fileName)

		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Unable to save the file", http.StatusInternalServerError)
			log.Println("Unable to save the file:", err)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Unable to save the file", http.StatusInternalServerError)
			log.Println("Unable to save the file:", err)
			return
		}

		imagePaths = append(imagePaths, filePath)
		log.Println("File uploaded successfully:", filePath)
	}

	images := strings.Join(imagePaths, ";")
	log.Println("Image paths:", images)

	db, err := dbsql.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO topics (title, categoryid, tags, content, images, owner, like, dislike, avatar, createat) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Database query preparation error", http.StatusInternalServerError)
		log.Println("Database query preparation error:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, category, tags, content, images, owner, like, dislike, avatar, createat)
	if err != nil {
		http.Error(w, "Database query execution error", http.StatusInternalServerError)
		log.Println("Database query execution error:", err)
		return
	}

	log.Println("Topic created successfully")
	http.Redirect(w, r, "/showtopics?id="+category, http.StatusSeeOther)
}
', 'AZAZA
', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.pnga7b3f3c5-009d-48b6-9eed-6d76b718c64d', 4, 'news', '', 0, 0, '2024-06-25 14:32:06', 'sayzx'),
	(24, 'a', 'a', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.pnga7b3f3c5-009d-48b6-9eed-6d76b718c64d', 4, 'info', '', 0, 0, '2024-06-25 14:32:29', 'sayzx'),
	(25, 'aa', 'aa', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.pnga7b3f3c5-009d-48b6-9eed-6d76b718c64d', 1, 'bug', '', 0, 0, '2024-06-25 14:33:28', 'sayzx'),
	(26, 'aa', 'aaa', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.pnge69073e4-1eba-47de-8aeb-14e32dd2d8f8', 2, 'news', '', 0, 0, '2024-06-25 14:36:36', 'sayzx'),
	(27, 'dev', 'Dev', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.png157c17cb-92b0-46cf-8b5b-20904c93615d', 5, 'general', '', 0, 0, '2024-06-25 14:45:33', 'sayzx'),
	(28, 'a', 'a', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.png157c17cb-92b0-46cf-8b5b-20904c93615d', 1, 'general', '', 0, 0, '2024-06-25 14:48:27', 'sayzx'),
	(29, 'a', 'a', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.png157c17cb-92b0-46cf-8b5b-20904c93615d', 3, 'general', 'web\uploads\Capture d''écran 2024-05-24 160817.png', 0, 0, '2024-06-25 14:51:02', 'sayzx'),
	(30, 'de jeuxPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not foundPost not found', 'Leak', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.png38833408-d263-4dd1-95b4-19a48e1999f0', 5, 'info', 'web\uploads\Capture d''écran 2024-06-03 144323.png', 0, 0, '2024-06-25 15:39:00', 'sayzx');
/*!40000 ALTER TABLE "topics" ENABLE KEYS */;

-- Listage de la structure de la table forum. users
CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	, "rank" VARCHAR(50) NULL DEFAULT NULL, "avatar" VARCHAR(256) NULL DEFAULT NULL);

-- Listage des données de la table forum.users : -1 rows
/*!40000 ALTER TABLE "users" DISABLE KEYS */;
INSERT INTO "users" ("id", "username", "email", "password", "rank", "avatar") VALUES
	(1, 'xplit', 'xplit@gmail.com', '$2a$10$k6E99wgEOPgJ.KNJagFPBufnlvoba9AKYL1v9vTXEd.P2jWaUeEne', 'admio,', 'https://media.discordapp.net/attachments/1224092616426258432/1252734375104086137/klife.com.png?ex=66734b4d&is=6671f9cd&hm=1514c7d6cb6e53adb6ab43cfde2ff0dca0cbf22c936663659a3bc89c0b60916e&=&format=webp&quality=lossless&width=749&height=749'),
	(2, 'xplit', 'xplit@gmail.com', '$2a$10$R/Jf3bnQnQ1z4FUMRnOlMeMdLOVL.qiHDTQmhnKDipvitpj89dPwe', NULL, 'https://media.discordapp.net/attachments/1224092616426258432/1252734375104086137/klife.com.png?ex=66734b4d&is=6671f9cd&hm=1514c7d6cb6e53adb6ab43cfde2ff0dca0cbf22c936663659a3bc89c0b60916e&=&format=webp&quality=lossless&width=749&height=749'),
	(3, 'Sayzx', 'sayzx@zdevpro.fr', '$2a$10$sF2wCAS1Cm79rrf9sRJUc.9G9ghRfiFfnbrfFCOQ6GfUyZd.fQTSu', NULL, 'https://media.discordapp.net/attachments/1224092616426258432/1252734375104086137/klife.com.png?ex=66734b4d&is=6671f9cd&hm=1514c7d6cb6e53adb6ab43cfde2ff0dca0cbf22c936663659a3bc89c0b60916e&=&format=webp&quality=lossless&width=749&height=749'),
	(4, 'nicolas', 'nicolas.gouy@epitech.eu', '$2a$10$6aS937YteVGxANDvAGFjIO9mIf1beILLLJA2SIH6fhdd4jy13T.fm', NULL, NULL),
	(10, 'hashedPassword', 'hashedPassword@hashedPassword.fr', '$2a$10$S4kzImqShmYxgMblGOYjpeKf6zr6LZhQC5P2cUoVL.rp9f7qPfNye', NULL, NULL);
/*!40000 ALTER TABLE "users" ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
