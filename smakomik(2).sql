-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 24 Feb 2024 pada 23.00
-- Versi server: 10.4.13-MariaDB
-- Versi PHP: 7.4.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `smakomik`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `comic_books`
--

CREATE TABLE `comic_books` (
  `id_comic` int(11) NOT NULL,
  `name_c` varchar(50) DEFAULT NULL,
  `genre` varchar(50) DEFAULT NULL,
  `image` varchar(100) DEFAULT NULL,
  `country_c` varchar(50) DEFAULT NULL,
  `access_u` int(100) DEFAULT NULL,
  `view_u` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `comic_books`
--

INSERT INTO `comic_books` (`id_comic`, `name_c`, `genre`, `image`, `country_c`, `access_u`, `view_u`) VALUES
(103, 'Junji Ito: Cat Diary', 'Horror', 'd80bd95b-c195-49ab-86cc-a11d38045aec.jpg', 'Jepang', 1, ''),
(104, 'Junji Ito: Tomie Take Over', 'Horror', 'd31b9614-719c-403f-998b-e034aed6f64c.jpg', 'Jepang', 0, ''),
(105, 'Junji Ito: Zona Hantu', 'Horror', '3a132fb6-a7dc-4317-89e5-5fb7e254abef.jpeg', 'Jepang', 1, '');

-- --------------------------------------------------------

--
-- Struktur dari tabel `comic_details`
--

CREATE TABLE `comic_details` (
  `id_det_comic` int(11) NOT NULL,
  `comic_id` int(11) NOT NULL,
  `status` varchar(50) NOT NULL,
  `reader_age` int(11) NOT NULL,
  `how_read` varchar(50) NOT NULL,
  `comic_artist` varchar(50) NOT NULL,
  `description` varchar(300) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `comic_details`
--

INSERT INTO `comic_details` (`id_det_comic`, `comic_id`, `status`, `reader_age`, `how_read`, `comic_artist`, `description`) VALUES
(21, 103, 'Ongoing', 18, 'Kanan', 'Junji Ito', 'Cerita tentang kucing'),
(22, 104, 'Ongoing', 18, 'Kanan', 'Junji Ito', 'Tentang kisah tomie'),
(23, 105, 'Onoging', 18, 'Kanan', 'Junji Ito', 'Tentang zona hantu');

-- --------------------------------------------------------

--
-- Struktur dari tabel `comic_subbtns`
--

CREATE TABLE `comic_subbtns` (
  `id_sub` int(11) NOT NULL,
  `comic_id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `name_sub` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `comic_subbtns`
--

INSERT INTO `comic_subbtns` (`id_sub`, `comic_id`, `name`, `name_sub`) VALUES
(27, 103, 'Chapter 1', 'Kemunculan mu'),
(28, 104, 'Chapter 1', 'Wanita menangis'),
(29, 105, 'Chapter 1', 'Wanita menangis');

-- --------------------------------------------------------

--
-- Struktur dari tabel `subbtns_images`
--

CREATE TABLE `subbtns_images` (
  `id_image` int(11) NOT NULL,
  `sub_id` int(11) NOT NULL,
  `images` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `subbtns_images`
--

INSERT INTO `subbtns_images` (`id_image`, `sub_id`, `images`) VALUES
(101, 27, '7a71b08e-9e9f-4bb8-af4f-8ea6476690a9.jpg'),
(102, 27, '6484756f-c7fe-411e-9012-cec5aae01406.jpg'),
(103, 27, 'd508c75c-0087-466d-b018-0f04b12b9de8.jpg'),
(104, 27, 'b9459cc1-16a2-4018-ab9d-ed72c074377a.jpg'),
(105, 27, '2d10c044-96ad-4e99-98c1-7d2e94368886.jpg'),
(106, 27, 'f3ce1631-e265-4d9b-90ad-4b9e95b7319c.jpg'),
(107, 27, 'd331d240-2864-4cf7-9193-b8f9e5833b2f.jpg'),
(108, 27, '3e3a7e15-6256-49b2-8fa0-8a07a02a2a19.jpg'),
(109, 27, '38367d52-02a3-4141-9db9-08b3f4496ede.jpg'),
(110, 27, 'e66700be-7331-46e8-8b86-e22939af5125.jpg'),
(111, 27, '2a49f64f-b41f-474c-b79a-619a9c47b043.jpg'),
(112, 27, 'ee51c4d5-b4ff-4263-809e-fb9a1909a23d.jpg'),
(113, 28, '184bd8e7-2e21-4e3f-a6f8-63df7e48f9d1.jpg'),
(114, 28, '8cdf4640-a858-460e-8095-f28d2dbfdbdf.jpg'),
(115, 28, '7285d76e-d143-48a7-9c26-b2fdbe7dee9b.jpg'),
(116, 28, '010afbb6-ade0-4891-a367-7972009aec08.jpg'),
(119, 29, 'b2337e6c-c071-4f0e-9ce3-7ae5494281d5.jpg'),
(120, 29, '12144431-b535-480b-a5b3-2ad83c368dea.jpg'),
(121, 29, 'c356e4b1-39d9-48d2-8430-b0a9798154b0.jpg'),
(122, 29, 'db01e651-27df-4abe-a1e4-11e932e8b1f4.jpg'),
(123, 29, 'b98adf88-0416-4682-a964-2fb9eff4f7bc.jpg'),
(124, 29, '36aba4ab-f8e3-4770-9aaf-01e06e4e4d73.jpg'),
(125, 29, '5e36ad31-6a93-4f75-b380-8695df6cd2aa.jpg'),
(126, 29, '24910996-1278-47d2-9a69-ea854ec682da.jpg'),
(127, 29, '48403e90-a8e3-42a1-b63c-b8b405faba1b.jpg'),
(128, 29, 'deded23f-ee8d-477e-84bd-adc04e3de8a5.jpg'),
(129, 29, 'af4b1cdd-6302-4baa-8476-20b75f717f34.jpeg');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id_users` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(150) NOT NULL,
  `access_u` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id_users`, `username`, `password`, `access_u`) VALUES
(7, 'admin', '$2a$10$c1cXNja4GioprxrOMIsFpO17jjUbApqk0Mf6CeZ/MH2iu.k.dOEKy', '');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `comic_books`
--
ALTER TABLE `comic_books`
  ADD PRIMARY KEY (`id_comic`);

--
-- Indeks untuk tabel `comic_details`
--
ALTER TABLE `comic_details`
  ADD PRIMARY KEY (`id_det_comic`),
  ADD UNIQUE KEY `comic_id` (`comic_id`);

--
-- Indeks untuk tabel `comic_subbtns`
--
ALTER TABLE `comic_subbtns`
  ADD PRIMARY KEY (`id_sub`),
  ADD KEY `fkid` (`comic_id`);

--
-- Indeks untuk tabel `subbtns_images`
--
ALTER TABLE `subbtns_images`
  ADD PRIMARY KEY (`id_image`),
  ADD KEY `sub_id` (`sub_id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_users`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `comic_books`
--
ALTER TABLE `comic_books`
  MODIFY `id_comic` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=106;

--
-- AUTO_INCREMENT untuk tabel `comic_details`
--
ALTER TABLE `comic_details`
  MODIFY `id_det_comic` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- AUTO_INCREMENT untuk tabel `comic_subbtns`
--
ALTER TABLE `comic_subbtns`
  MODIFY `id_sub` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=30;

--
-- AUTO_INCREMENT untuk tabel `subbtns_images`
--
ALTER TABLE `subbtns_images`
  MODIFY `id_image` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=130;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id_users` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `comic_details`
--
ALTER TABLE `comic_details`
  ADD CONSTRAINT `comic_details_ibfk_1` FOREIGN KEY (`comic_id`) REFERENCES `comic_books` (`id_comic`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `comic_subbtns`
--
ALTER TABLE `comic_subbtns`
  ADD CONSTRAINT `fkid` FOREIGN KEY (`comic_id`) REFERENCES `comic_books` (`id_comic`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `subbtns_images`
--
ALTER TABLE `subbtns_images`
  ADD CONSTRAINT `subbtns_images_ibfk_1` FOREIGN KEY (`sub_id`) REFERENCES `comic_subbtns` (`id_sub`) ON DELETE CASCADE ON UPDATE NO ACTION;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
