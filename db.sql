-- --------------------------------------------------------
-- Hôte:                         C:\Users\boucl\Desktop\Ynov\ForumGO\internal\sql\forum.db
-- Version du serveur:           3.45.3
-- SE du serveur:                
-- HeidiSQL Version:             12.7.0.6850
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

-- Listage de la structure de table forum. comments
CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    postid INTEGER,
    content TEXT,
    owner TEXT,
    createat DATETIME, "avatar" VARCHAR(252) NULL DEFAULT 'https://media.discordapp.net/attachments/1254153761211945093/1255220022301954181/image.png?ex=667cfefe&is=667bad7e&hm=45c2c4da5b337bd23b10497100f849013f968bacccd9b84d8155f13ece4f34bd&=&format=webp&quality=lossless&width=643&height=437',
    FOREIGN KEY (postid) REFERENCES topics(id)
);

-- Listage des données de la table forum.comments : -1 rows
/*!40000 ALTER TABLE "comments" DISABLE KEYS */;
INSERT INTO "comments" ("id", "postid", "content", "owner", "createat", "avatar") VALUES
	(45, 36, 'COUCOU TROP BIEN', 'sayzx', '2024-06-27 15:22:37', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.png?rand=NRt4ENPCH0'),
	(46, 36, 'cc', 'sayzx', '2024-06-27 15:59:22', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.png?rand=NRt4ENPCH0'),
	(47, 39, 'cc
', 'Sayzx', '2024-06-28 16:53:51', 'https://avatars.githubusercontent.com/u/74567624?v=4?rand=Y63kEj6NVk'),
	(48, 38, 'Viole
', 'sayzx', '2024-07-03 14:52:18', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.png?rand=gPKqlKTqyH'),
	(49, 37, 'cc
', 'bouclierbleu39@gmail.com', '2024-07-05 13:38:43', 'https://lh3.googleusercontent.com/a-/ALV-UjXqImBEvEDp8FlbA_leJr40mstCQ2RQ9rFvWMKxxbuOWn6_ZNoQ=s96-c?rand=GcckKVEC7L');
/*!40000 ALTER TABLE "comments" ENABLE KEYS */;

-- Listage de la structure de table forum. dislike
CREATE TABLE IF NOT EXISTS "dislike" (
	"id" INTEGER NULL,
	"username" VARCHAR(50) NULL DEFAULT NULL,
	"postid" VARCHAR(50) NULL DEFAULT NULL
);

-- Listage des données de la table forum.dislike : -1 rows
/*!40000 ALTER TABLE "dislike" DISABLE KEYS */;
INSERT INTO "dislike" ("id", "username", "postid") VALUES
	(NULL, 'sayzx', '38');
/*!40000 ALTER TABLE "dislike" ENABLE KEYS */;

-- Listage de la structure de table forum. likes
CREATE TABLE IF NOT EXISTS "likes" (
	"id" INTEGER NULL,
	"username" VARCHAR(50) NULL DEFAULT NULL,
	"postid" INTEGER NULL
);

-- Listage des données de la table forum.likes : -1 rows
/*!40000 ALTER TABLE "likes" DISABLE KEYS */;
INSERT INTO "likes" ("id", "username", "postid") VALUES
	(NULL, 'sayzx', 36),
	(NULL, 'Sayzx', 37),
	(NULL, 'bouclierbleu39@gmail.com', 37);
/*!40000 ALTER TABLE "likes" ENABLE KEYS */;

-- Listage de la structure de table forum. loginlogs
CREATE TABLE IF NOT EXISTS loginlogs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    platform TEXT NOT NULL,
    datetime DATETIME NOT NULL
);

-- Listage des données de la table forum.loginlogs : -1 rows
/*!40000 ALTER TABLE "loginlogs" DISABLE KEYS */;
INSERT INTO "loginlogs" ("id", "username", "platform", "datetime") VALUES
	(12, 'admin@admin.fr', 'Local', '2024-06-27 17:05:35.8983875 +0200 CEST m=+7.039989301'),
	(13, 'sayzx', 'Discord', '2024-06-27 17:14:36.662418 +0200 CEST m=+13.050469401'),
	(14, 'sayzx', 'Discord', '2024-06-27 17:15:57.6665648 +0200 CEST m=+12.019904101'),
	(15, 'sayzx', 'Discord', '2024-06-28 13:41:50.7102953 +0200 CEST m=+14.944832901'),
	(16, 'sayzx', 'Discord', '2024-06-28 13:43:24.8417484 +0200 CEST m=+26.968019301'),
	(17, 'sayzx', 'Discord', '2024-06-28 13:43:55.0577096 +0200 CEST m=+57.183980501'),
	(18, 'Sayzx', 'GitHub', '2024-06-28 13:46:16.224319 +0200 CEST m=+9.646353001'),
	(19, 'bouclierbleu39@gmail.com', 'Google', '2024-06-28 14:16:06.5102372 +0200 CEST m=+14.055313601'),
	(20, 'bouclierbleu39@gmail.com', 'Google', '2024-06-28 14:17:21.4218168 +0200 CEST m=+35.810487501'),
	(21, 'bouclierbleu39@gmail.com', 'Google', '2024-06-28 14:20:23.7335549 +0200 CEST m=+17.525309801'),
	(22, 'bouclierbleu39@gmail.com', 'Google', '2024-06-28 14:22:27.2608171 +0200 CEST m=+8.644252301'),
	(23, 'sayzx', 'Discord', '2024-06-28 16:51:58.3925781 +0200 CEST m=+16.562217701'),
	(24, 'Sayzx', 'GitHub', '2024-06-28 16:52:18.0778849 +0200 CEST m=+36.247524501'),
	(25, 'Sayzx', 'GitHub', '2024-06-28 16:52:26.2390783 +0200 CEST m=+44.408717901'),
	(26, 'sayzx', 'Discord', '2024-07-03 13:43:16.283478 +0200 CEST m=+11.333576901'),
	(27, 'sayzx', 'Discord', '2024-07-04 13:58:09.8980032 +0200 CEST m=+12.728993501'),
	(28, 'Sayzx', 'GitHub', '2024-07-04 14:33:37.9937985 +0200 CEST m=+875.129679801'),
	(29, 'klife@gmail.com', 'Local', '2024-07-04 14:34:07.5912272 +0200 CEST m=+904.727108501'),
	(30, 'sayzx', 'Discord', '2024-07-04 16:18:50.1942868 +0200 CEST m=+20.142587101'),
	(31, 'sayzx', 'Discord', '2024-07-04 16:23:37.6424822 +0200 CEST m=+8.508770201'),
	(32, 'bouclierbleu39@gmail.com', 'Google', '2024-07-04 16:24:03.541959 +0200 CEST m=+34.408247001'),
	(33, 'sayzx', 'Discord', '2024-07-05 13:51:13.9853337+02:00'),
	(34, 'aa@aa.fr', 'Local', '2024-07-05 14:32:51.0058231+02:00'),
	(35, 'aa@aa.fr', 'Local', '2024-07-05 14:33:03.7605766+02:00'),
	(36, 'aa@aa.fr', 'Local', '2024-07-05 14:35:13.0219367+02:00'),
	(37, 'aa@aa.fr', 'Local', '2024-07-05 14:48:27.1000335+02:00'),
	(38, 'sayzx', 'Discord', '2024-07-05 14:50:30.7664555+02:00'),
	(39, 'aa@aa.fr', 'Local', '2024-07-05 14:50:43.1069996+02:00'),
	(40, 'sayzx', 'Discord', '2024-07-05 15:26:37.5801335+02:00'),
	(41, 'Sayzx', 'GitHub', '2024-07-05 15:45:22.5050773+02:00'),
	(42, 'sayzx', 'Discord', '2024-07-05 15:45:46.163174+02:00'),
	(43, 'sayzx', 'Discord', '2024-07-05 15:53:34.9468944+02:00');
/*!40000 ALTER TABLE "loginlogs" ENABLE KEYS */;

-- Listage de la structure de table forum. moderator_wait
CREATE TABLE IF NOT EXISTS "moderator_wait" (
	"id" VARCHAR(50) NULL DEFAULT NULL);

-- Listage des données de la table forum.moderator_wait : -1 rows
/*!40000 ALTER TABLE "moderator_wait" DISABLE KEYS */;
/*!40000 ALTER TABLE "moderator_wait" ENABLE KEYS */;

-- Listage de la structure de table forum. reportspost
CREATE TABLE IF NOT EXISTS "reportspost" (
	"postid" INTEGER NULL, "content" VARCHAR(255) NULL DEFAULT NULL, "title" VARCHAR(255) NULL DEFAULT NULL, "avatar" VARCHAR(255) NULL DEFAULT NULL, "owner" VARCHAR(255) NULL DEFAULT NULL);

-- Listage des données de la table forum.reportspost : -1 rows
/*!40000 ALTER TABLE "reportspost" DISABLE KEYS */;
INSERT INTO "reportspost" ("postid", "content", "title", "avatar", "owner") VALUES
	(38, 'dEV', 'Dev', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.png?rand=ThTf53hyFv', 'sayzx');
/*!40000 ALTER TABLE "reportspost" ENABLE KEYS */;

-- Listage de la structure de table forum. topics
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
	(37, 'DEV', 'Dev', 'https://lh3.googleusercontent.com/a-/ALV-UjXqImBEvEDp8FlbA_leJr40mstCQ2RQ9rFvWMKxxbuOWn6_ZNoQ=s96-c', 1, 'news', 'web\uploads\Capture d''écran 2024-06-03 144323.png', 2, 0, '2024-06-27 13:56:32', 'bouclierbleu39@gmail.com'),
	(38, 'dEV', 'Dev', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.png?rand=NRt4ENPCH0?rand=a0acHin8n5', 1, 'general', 'web\uploads\Capture d''écran 2024-06-03 144323.png', 0, 1, '2024-06-27 14:56:18', 'sayzx'),
	(40, 'devreznfp
eznznf', 'd', 'https://lh3.googleusercontent.com/a-/ALV-UjXqImBEvEDp8FlbA_leJr40mstCQ2RQ9rFvWMKxxbuOWn6_ZNoQ=s96-c?rand=GcckKVEC7L?rand=KC0j3VKxqZ', 1, 'general', '', 0, 0, '2024-07-05 13:41:46', 'bouclierbleu39@gmail.com'),
	(41, 'reportedPosts', 'reportedPosts', 'https://avatars.githubusercontent.com/u/74567624?v=4?rand=KuF94lhiY0?rand=YmBXdhhNgo', 1, 'general', '', 0, 0, '2024-07-05 15:45:32', 'Sayzx');
/*!40000 ALTER TABLE "topics" ENABLE KEYS */;

-- Listage de la structure de table forum. users
CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	, "rank" VARCHAR(50) NULL DEFAULT NULL, "avatar" VARCHAR(256) NULL DEFAULT NULL, "platform" VARCHAR(50) NULL DEFAULT NULL, "userid" VARCHAR(255) NOT NULL DEFAULT '');

-- Listage des données de la table forum.users : -1 rows
/*!40000 ALTER TABLE "users" DISABLE KEYS */;
INSERT INTO "users" ("id", "username", "email", "password", "rank", "avatar", "platform", "userid") VALUES
	(23, 'bouclierbleu39@gmail.com', 'bouclierbleu39@gmail.com', '$2a$10$b2JYPIbhPn3FnFzhs3Y0Y.7fUA1WxsCJ9.rMeHRArcA4xILl5oC3G', 'user', 'https://lh3.googleusercontent.com/a-/ALV-UjXqImBEvEDp8FlbA_leJr40mstCQ2RQ9rFvWMKxxbuOWn6_ZNoQ=s96-c?rand=GcckKVEC7L?rand=RYmDZH14mD?rand=1jBhH75sc2?rand=kUkHUOSlBj?rand=y4kwnL0NW1?rand=vxDnm1hrp8?rand=61na9S0sak?rand=Z4krVVbRaY?rand=4iwdyVUzGQ?rand=eE1JSrrfsQ?rand=kDixCEVZWZ?rand=KrLm9d1Fqx?rand=Fxbn2NLFKe?rand=FgbUQGM5hy?rand=g2m668avBf?rand=pjFAKk7qhe?rand=jxQe8tmlCv?rand=y32XcgKJb4?rand=aBH5rdXWjg?rand=3RNovHkHKf?rand=4tD7xF757V?rand=kNpqR3YVEc?rand=bKaRlP4TOf?rand=yY2UUCTlPK?rand=uSdQQEpTs4?rand=DA8GhPCRu8?rand=y3fdYf7Q0Z?rand=90CsHo4bcE?rand=UcHe63ilkJ?rand=bDyw0eDYcD?rand=TUnROLdffM?rand=Ql3Z2qpI5l?rand=xY2ez61KmT?rand=RxHGsM2gXP?rand=bduMLWdnpG?rand=vkPQ83Ffmi', 'Google', 'e0c906f2-ae13-4044-b8f2-06159855ac16'),
	(27, 'aa', 'aa@aa.fr', '$2a$10$dLvnBaGRFI5XMbaTVLsqjeG94m/mhctX5fZvDcrQ5z5/c/JPKga1C', 'user', 'https://media.discordapp.net/attachments/1224092616426258432/1252742512209301544/1247.png?ex=668913a1&is=6687c221&hm=895af00c0facede320bc213425295dbeae26a1652ae0a217e40a8e80bb418dfe&=&format=webp&quality=lossless&width=640&height=640?rand=aTq2G38oiI?rand=eE1JSrrfsQ?rand=kDixCEVZWZ?rand=KrLm9d1Fqx?rand=Fxbn2NLFKe?rand=FgbUQGM5hy?rand=g2m668avBf?rand=pjFAKk7qhe?rand=jxQe8tmlCv?rand=y32XcgKJb4?rand=NTRFRaPNPh?rand=x34ijBlWLy?rand=iQRT4EdTqR?rand=TjNY85AO24?rand=Hp1TMMIOcE?rand=C52kNnrW01?rand=1nhNVe97Ds?rand=7SuNPPZFPT?rand=RHYg984leC?rand=0wZpuf04Ir?rand=0hu23NxOcr?rand=4yURrXUcLF?rand=x0Oh4m3icm?rand=eXurEpM1BA?rand=6k4e5QwJAv?rand=giFBqtqoKi?rand=6qo5ZSjtX6?rand=SEJaq8v9m7', 'Local', 'a9fe6051-994f-4a23-a8a0-0328a8ba6c7e'),
	(28, 'Sayzx', 'Sayzx', '$2a$10$dBcPEdMJHhwSY7cXeX5k3OdZGUjTH3C.Stst/npRrolJUNr5CGL.2', 'user', 'https://avatars.githubusercontent.com/u/74567624?v=4?rand=KuF94lhiY0?rand=y32XcgKJb4?rand=NTRFRaPNPh?rand=x34ijBlWLy?rand=GPjwurSUhJ?rand=TjNY85AO24?rand=Hp1TMMIOcE?rand=C52kNnrW01?rand=1nhNVe97Ds?rand=7SuNPPZFPT?rand=RHYg984leC?rand=0wZpuf04Ir?rand=0hu23NxOcr?rand=4yURrXUcLF?rand=x0Oh4m3icm?rand=eXurEpM1BA?rand=t57n7XU3Pw?rand=BYsIc4Jxah?rand=6qo5ZSjtX6?rand=SEJaq8v9m7', 'GitHub', '2ccf9462-5677-470c-88b4-f46abc5f7cb8'),
	(29, 'sayzx', 'sayzx', '$2a$10$CyQn035m.bymrPyY.l8fC.DTDa9iBz8X3UutffAcLV6yKxrJpqfC.', 'user', 'https://cdn.discordapp.com/avatars/826826070899949601/d0ee2f29d053e069b0cbd0fc3bdb62fb.png?rand=o0yCAF4yMq?rand=x34ijBlWLy?rand=GPjwurSUhJ?rand=TjNY85AO24?rand=Hp1TMMIOcE?rand=C52kNnrW01?rand=1nhNVe97Ds?rand=7SuNPPZFPT?rand=RHYg984leC?rand=pcaAqbwduj?rand=0hu23NxOcr?rand=4yURrXUcLF?rand=x0Oh4m3icm?rand=eXurEpM1BA?rand=t57n7XU3Pw?rand=BYsIc4Jxah?rand=5B9w1Uuj2T?rand=SEJaq8v9m7', 'Discord', '03cb4d24-36ba-4a75-b5a7-3faa20eb48f8');
/*!40000 ALTER TABLE "users" ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
