-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Mar 01, 2024 at 09:46 AM
-- Server version: 5.7.34
-- PHP Version: 7.4.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `steradian`
--

-- --------------------------------------------------------

--
-- Table structure for table `car`
--

CREATE TABLE `car` (
  `car_id` int(11) NOT NULL,
  `car_name` varchar(50) NOT NULL,
  `day_rate` double NOT NULL,
  `month_rate` double NOT NULL,
  `image_car` varchar(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `car`
--

INSERT INTO `car` (`car_id`, `car_name`, `day_rate`, `month_rate`, `image_car`) VALUES
(3, 'Mustang', 120000, 300000, 'app/upload/7.jpeg'),
(4, 'GTR - 500', 160000, 300000, 'app/upload/7.jpeg'),
(5, 'GTR - 200', 160000, 300000, 'app/upload/7.jpeg'),
(6, 'GTR - 200', 160000, 300000, 'app/upload/7.jpeg'),
(7, 'GTR - 200', 160000, 300000, 'app/upload/7.jpeg');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `order_id` int(11) NOT NULL,
  `car_id` int(11) NOT NULL,
  `order_date` date NOT NULL,
  `pickup_date` date NOT NULL,
  `dropoff_date` date NOT NULL,
  `pickup_location` varchar(50) NOT NULL,
  `dropoff_location` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`order_id`, `car_id`, `order_date`, `pickup_date`, `dropoff_date`, `pickup_location`, `dropoff_location`) VALUES
(1, 3, '2024-03-04', '2024-05-04', '2024-06-08', 'dago', 'petamburan'),
(3, 3, '2024-03-04', '2024-05-04', '2024-06-08', 'cikopo', 'petamburan'),
(4, 3, '2024-03-04', '2024-05-04', '2024-06-08', 'cikopo', 'petamburan'),
(5, 3, '2024-03-04', '2024-05-04', '2024-06-08', 'cikopo', 'petamburan'),
(6, 3, '2024-03-04', '2024-05-04', '2024-06-08', 'cikopo', 'petamburan'),
(7, 3, '2024-03-04', '2024-05-04', '2024-06-08', 'cikopo', 'petamburan'),
(8, 3, '2024-03-04', '2024-05-04', '2024-06-08', 'cikopo', 'petamburan'),
(9, 3, '2024-03-04', '2024-05-04', '2024-06-08', 'cikopo', 'petamburan');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `car`
--
ALTER TABLE `car`
  ADD PRIMARY KEY (`car_id`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`order_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `car`
--
ALTER TABLE `car`
  MODIFY `car_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `order_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
