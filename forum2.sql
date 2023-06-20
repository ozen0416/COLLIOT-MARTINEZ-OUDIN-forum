-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : lun. 12 juin 2023 à 09:59
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
-- Base de données : `test`
--

-- --------------------------------------------------------

--
-- Structure de la table `category`
--

DROP TABLE IF EXISTS `category`;
CREATE TABLE IF NOT EXISTS `category` (
  `id_category` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE latin1_bin NOT NULL,
  PRIMARY KEY (`id_category`),
  UNIQUE KEY `title` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

-- --------------------------------------------------------

--
-- Structure de la table `message`
--

DROP TABLE IF EXISTS `message`;
CREATE TABLE IF NOT EXISTS `message` (
  `id_message` int NOT NULL AUTO_INCREMENT,
  `content` varchar(255) COLLATE latin1_bin NOT NULL,
  `Date_Post` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  `id_user` int NOT NULL,
  `id_topic` int NOT NULL,
  PRIMARY KEY (`id_message`),
  KEY `id_user` (`id_user`),
  KEY `id_topic` (`id_topic`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

-- --------------------------------------------------------

--
-- Structure de la table `picture`
--

DROP TABLE IF EXISTS `picture`;
CREATE TABLE IF NOT EXISTS `picture` (
  `id_picture` int NOT NULL AUTO_INCREMENT,
  `path_picture` varchar(255) COLLATE latin1_bin NOT NULL,
  `extension_pic` varchar(255) COLLATE latin1_bin NOT NULL,
  `name_pic` varchar(255) COLLATE latin1_bin NOT NULL,
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
  `content` varchar(255) COLLATE latin1_bin NOT NULL,
  `publication_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  `id_category` int NOT NULL,
  `id_user` int NOT NULL,
  PRIMARY KEY (`id_topic`),
  KEY `id_category` (`id_category`),
  KEY `id_user` (`id_user`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

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
  `username` varchar(255) COLLATE latin1_bin NOT NULL,
  `email` varchar(255) COLLATE latin1_bin NOT NULL,
  `passwd` varchar(255) COLLATE latin1_bin NOT NULL,
  `birth_date` date NOT NULL,
  `signin_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  `id_picture` int,
  PRIMARY KEY (`id_user`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  KEY `id_picture` (`id_picture`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

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
