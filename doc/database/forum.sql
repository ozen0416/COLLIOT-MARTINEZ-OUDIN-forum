-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : mar. 20 juin 2023 à 19:43
-- Version du serveur : 8.0.31
-- Version de PHP : 8.0.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `forum2`
--

-- --------------------------------------------------------

--
-- Structure de la table `category`
--

DROP TABLE IF EXISTS `category`;
CREATE TABLE IF NOT EXISTS `category` (
  `id_category` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
  PRIMARY KEY (`id_category`),
  UNIQUE KEY `title` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `category`
--

INSERT INTO `category` (`id_category`, `title`) VALUES
(2, 'beauty'),
(4, 'cinema'),
(1, 'funny'),
(5, 'music'),
(3, 'news');

-- --------------------------------------------------------

--
-- Structure de la table `message`
--

DROP TABLE IF EXISTS `message`;
CREATE TABLE IF NOT EXISTS `message` (
  `id_message` int NOT NULL AUTO_INCREMENT,
  `content` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
  `Date_Post` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `id_user` int NOT NULL,
  `id_topic` int NOT NULL,
  PRIMARY KEY (`id_message`),
  KEY `id_user` (`id_user`),
  KEY `id_topic` (`id_topic`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `message`
--

INSERT INTO `message` (`id_message`, `content`, `Date_Post`, `id_user`, `id_topic`) VALUES
(4, 'ELLE EST TROP BEEEELLE', '2023-06-20 21:36:16', 23, 3),
(5, 'Ouiii, je l\'ai vu l\'année dernière à Paris...', '2023-06-20 21:39:35', 24, 3),
(6, 'Hooo la chance', '2023-06-20 21:40:19', 23, 3);

-- --------------------------------------------------------

--
-- Structure de la table `picture`
--

DROP TABLE IF EXISTS `picture`;
CREATE TABLE IF NOT EXISTS `picture` (
  `id_picture` int NOT NULL AUTO_INCREMENT,
  `path_picture` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
  `extension_pic` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
  `name_pic` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
  PRIMARY KEY (`id_picture`),
  UNIQUE KEY `path_picture` (`path_picture`),
  UNIQUE KEY `extension_pic` (`extension_pic`),
  UNIQUE KEY `name_pic` (`name_pic`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

-- --------------------------------------------------------

--
-- Structure de la table `topic`
--

DROP TABLE IF EXISTS `topic`;
CREATE TABLE IF NOT EXISTS `topic` (
  `id_topic` int NOT NULL AUTO_INCREMENT,
  `content` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
  `publication_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `id_category` int NOT NULL,
  `id_user` int NOT NULL,
  PRIMARY KEY (`id_topic`),
  KEY `id_category` (`id_category`),
  KEY `id_user` (`id_user`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `topic`
--

INSERT INTO `topic` (`id_topic`, `content`, `publication_date`, `id_category`, `id_user`) VALUES
(2, 'test', '2023-06-16 00:24:19', 1, 2),
(3, 'ELLE EST TROP BEEEELLE', '2023-06-20 21:36:16', 1, 23);

-- --------------------------------------------------------

--
-- Structure de la table `upvote`
--

DROP TABLE IF EXISTS `upvote`;
CREATE TABLE IF NOT EXISTS `upvote` (
  `id_user` int NOT NULL,
  `id_message` int NOT NULL,
  PRIMARY KEY (`id_user`,`id_message`),
  KEY `id_message` (`id_message`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

-- --------------------------------------------------------

--
-- Structure de la table `users`
--

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `id_user` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
  `email` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
  `passwd` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
  `birth_date` date NOT NULL,
  `signin_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `id_picture` int DEFAULT NULL,
  PRIMARY KEY (`id_user`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  KEY `id_picture` (`id_picture`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `users`
--

INSERT INTO `users` (`id_user`, `username`, `email`, `passwd`, `birth_date`, `signin_date`, `id_picture`) VALUES
(2, 'test', 'test', '$2a$14$Lisxf21GTgQrafJ8FfJYNOV3rniatDJRcMaN2BigMsfQ/kEE9o9Zu', '2023-06-07', '2023-06-16 00:22:15', NULL),
(4, 'tesdty', 'tesxtd', '$2a$14$V5GB2UerqqlMXaEK2D8/mO.lU2y84f4FoCz/JzmG9fQ15XU6nC63C', '2023-06-21', '2023-06-19 14:50:37', NULL),
(6, 'ee', 'ee', '$2a$14$HnBFFqPktXtwqhiE5CQdw.712uZJEcqLEY77QVYOtfhmdUB1rmS8.', '2023-05-31', '2023-06-19 14:52:39', NULL),
(7, 'eee', 'eee', '$2a$14$Yy8qVyfjpWkE3DYNq8DjQ.zSUvTfIGTrTGp2mSI0DykQrg1bV6Vyq', '2023-06-01', '2023-06-19 14:58:56', NULL),
(9, '', '', '$2a$14$Vli8sXkTu/vB.7ukX93PQOp...7Vi0CInsiaWy17YEYtIbITEUPpC', '0000-00-00', '2023-06-19 15:29:24', NULL),
(23, 'DUALOVER', 'dua.lover@dua.dua', '$2a$14$oQfj6R725miQFZvhGLNOZe9PZ.KYm72gQ1jp5dsdKEQNTR.7LQvhO', '2023-06-15', '2023-06-20 21:34:57', NULL),
(24, 'iloveDUA', 'dua@dua.dua', '$2a$14$PzAnNC/j0LIrQxu491Kf.OHZOt43yphmG0xg8InhVQZB4v7HI3TYO', '2023-06-14', '2023-06-20 21:37:28', NULL);

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `message`
--
ALTER TABLE `message`
  ADD CONSTRAINT `message_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id_user`),
  ADD CONSTRAINT `message_ibfk_2` FOREIGN KEY (`id_topic`) REFERENCES `topic` (`id_topic`);

--
-- Contraintes pour la table `topic`
--
ALTER TABLE `topic`
  ADD CONSTRAINT `topic_ibfk_1` FOREIGN KEY (`id_category`) REFERENCES `category` (`id_category`),
  ADD CONSTRAINT `topic_ibfk_2` FOREIGN KEY (`id_user`) REFERENCES `users` (`id_user`);

--
-- Contraintes pour la table `upvote`
--
ALTER TABLE `upvote`
  ADD CONSTRAINT `upvote_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id_user`),
  ADD CONSTRAINT `upvote_ibfk_2` FOREIGN KEY (`id_message`) REFERENCES `message` (`id_message`);

--
-- Contraintes pour la table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `users_ibfk_1` FOREIGN KEY (`id_picture`) REFERENCES `picture` (`id_picture`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
