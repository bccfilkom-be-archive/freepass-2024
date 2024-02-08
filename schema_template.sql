-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 08, 2024 at 12:08 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `backend_bcc`
--

-- --------------------------------------------------------

--
-- Table structure for table `comment`
--

CREATE TABLE `comment` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `post_id` int(11) NOT NULL,
  `content` text NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `comment`
--

INSERT INTO `comment` (`id`, `user_id`, `post_id`, `content`, `timestamp`) VALUES
(1, 4, 2, 'I think I\'ll vote you.', '2024-02-07 04:12:44'),
(2, 4, 3, 'well, I hope you keep your promise after being elected.', '2024-02-08 10:56:14'),
(3, 4, 4, 'Yes, Sir!', '2024-02-08 10:57:11'),
(4, 5, 1, 'You look good!', '2024-02-08 10:57:49'),
(5, 7, 1, 'You\'re the best.', '2024-02-08 10:58:41');

-- --------------------------------------------------------

--
-- Table structure for table `election`
--

CREATE TABLE `election` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `election`
--

INSERT INTO `election` (`id`, `name`, `start_date`, `end_date`) VALUES
(1, '1st Election', '2024-02-06 01:00:00', '2025-02-10 13:00:00'),
(2, '2nd Election', '2024-02-13 01:00:00', '2024-02-23 13:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `post`
--

CREATE TABLE `post` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `title` text NOT NULL,
  `content` text NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `post`
--

INSERT INTO `post` (`id`, `user_id`, `title`, `content`, `timestamp`) VALUES
(1, 1, 'My First API', 'Hello world!', '2024-02-08 10:53:19'),
(2, 2, 'The Greatest Middle Earth!', 'I promise to restore our golden era if I am elected!', '2024-02-08 10:53:50'),
(3, 3, 'For Peace and Hogwarts!', 'I\'ll bring glory to our beloved Hogwarts if I\'m elected!', '2024-02-08 10:54:31'),
(4, 1, 'For Democracy!', 'Cast your vote on February 31, 2099! Don\'t forget the date!', '2024-02-08 10:56:55');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `nim` bigint(20) DEFAULT NULL,
  `name` text DEFAULT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `major` text DEFAULT NULL,
  `faculty` text DEFAULT NULL,
  `status` enum('user','candidate','admin') NOT NULL,
  `description` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `nim`, `name`, `username`, `password`, `major`, `faculty`, `status`, `description`) VALUES
(1, 12345678, 'Yoga Raditya Nala', 'yogarn', '$2b$10$7BvVlf.VZQ2PEeVIqYUn7OFlJ69yW.W6bza40GXW028uyLmSKs5jK', 'Computer Science', 'Faculty of Computer Science', 'candidate', 'Make Earth Great Again!'),
(2, 11234567, 'Olorin Gandalf', 'gandalf', '$2b$10$JQmXP1xA7P68C/NEqsrtI.f5IcWwKtvV8k8QieL/b7E46fyQWQrGG', 'Magic', 'Faculty of Witchcraft and Magic', 'candidate', 'Make Middle Earth Great Again!'),
(3, 12234567, 'Albus Percival Wulfric Brian Dumbledore', 'dumbledore', '$2b$10$Bfwom.aD1QkNNG6VeOh/4.4XS.StiKzuQzKOALFSk5POD7KyTT9S2', 'Magic', 'Faculty of Witchcraft and Magic', 'candidate', 'Make Hogwarts Great Again!'),
(4, 21234567, 'Yennefer of Vengerberg', 'yennefer', '$2b$10$XSR11kraucNB1SwTJ0dN/OkFoJRBLhEVC8TFlY0aJvix6Jiv0AzBu', 'Magic', 'Faculty of Witchcraft and Magic', 'user', 'Uhm, what is this?!'),
(5, 22234567, 'Alex Turner', 'am', '$2b$10$dpHggG7E3wAr9Td7rCniauS/pbbJOpJKcw6cbQF/LXpeX5SBUVIL2', 'Music', 'Faculty of Fine Arts and Music', 'user', 'Oh, when you look at me like that my darling, what did you expect?!'),
(6, 23234567, 'Musk Elon', 'elon', '$2b$10$ZJKs27QgyORkRzRHGk4nLuDY9DiVPyuJ7pHVLk1bB3T.oAzCb2KDG', 'Computer Science', 'Faculty of Computer Science', 'user', 'To the Mars!'),
(7, 24234567, 'Turing Alan', 'turing', '$2b$10$tcDMD9.K1qEjpnqBT9hyNuKsREM4i7QTcTb.MUeduevckrvIJ2BwW', 'Electrical Engineering', 'Faculty of Electrical Engineering', 'user', 'For the future!'),
(8, 88888888, 'admin', 'admin', '$2b$10$uJ11w9RQNBVS.lC5lCctROcuRAJNSxvbmcJ0NZfTeXHEtJwzuC2qO', 'Administration', 'Faculty of Administration', 'admin', 'God Mode!');

-- --------------------------------------------------------

--
-- Table structure for table `vote`
--

CREATE TABLE `vote` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `election_id` int(11) NOT NULL,
  `candidate_id` int(11) NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `vote`
--

INSERT INTO `vote` (`id`, `user_id`, `election_id`, `candidate_id`, `timestamp`) VALUES
(1, 6, 1, 1, '2024-02-08 11:04:04'),
(2, 5, 1, 1, '2024-02-08 11:04:16'),
(3, 1, 1, 1, '2024-02-08 11:05:05'),
(4, 4, 1, 2, '2024-02-08 11:06:02'),
(5, 3, 1, 3, '2024-02-08 11:06:28');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `comment`
--
ALTER TABLE `comment`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `post_id` (`post_id`);

--
-- Indexes for table `election`
--
ALTER TABLE `election`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `post`
--
ALTER TABLE `post`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `vote`
--
ALTER TABLE `vote`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`election_id`),
  ADD KEY `candidate_id` (`candidate_id`),
  ADD KEY `election_id` (`election_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `comment`
--
ALTER TABLE `comment`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `election`
--
ALTER TABLE `election`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `post`
--
ALTER TABLE `post`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `vote`
--
ALTER TABLE `vote`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `comment`
--
ALTER TABLE `comment`
  ADD CONSTRAINT `comment_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `comment_ibfk_2` FOREIGN KEY (`post_id`) REFERENCES `post` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `post`
--
ALTER TABLE `post`
  ADD CONSTRAINT `post_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `vote`
--
ALTER TABLE `vote`
  ADD CONSTRAINT `vote_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `vote_ibfk_2` FOREIGN KEY (`candidate_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `vote_ibfk_3` FOREIGN KEY (`election_id`) REFERENCES `election` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
