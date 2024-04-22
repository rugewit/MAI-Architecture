CREATE TABLE Users (
                       Id VARCHAR(255) PRIMARY KEY,
                       Name VARCHAR(255),
                       Lastname VARCHAR(255),
                       Password VARCHAR(255),
                       CreationDate TIMESTAMP
);

INSERT INTO Users (Id, Name, Lastname, Password, CreationDate)
VALUES
('6c6f5595-bacb-4e13-9a35-e880602a7200', 'Giovanny', 'Champlin', '$2a$10$XTnJFSzVsN8wlRkZSsuYiuCpscsRlD492rl.gyihGEEIDnnhjookW', '2024-04-17 07:38:23'),
('26fc1fde-676c-46ac-a155-412778f2673c', 'Ashlynn', 'Jast', '$2a$10$7yW8gRy3XF0AegcRAwSHKecqbd9tOFrUfx5qbg2eWV8hz0WiUcF1C', '2024-04-17 06:38:23'),
('6a118da2-8794-4c18-b39e-2220bd7261cc', 'Velma', 'Bashirian', '$2a$10$SmA5kbfripaxwvEXokH7qeM24rNLzzlngGjfFGijZuru8hsbg8Zb.', '2024-04-16 23:38:23'),
('37d4897b-752e-4229-acaa-69365b9263de', 'Isidro', 'Emard', '$2a$10$J7oDLad6QJFb4buGMg/gVOT7xlUsYexKsQVxBDeSsK3nuc19tYNV.', '2024-04-17 06:38:23'),
('5eaabb0b-ecac-4fe3-9c45-5a96d1b768f4', 'Destinee', 'Grady', '$2a$10$aRA/ZTiShCRxlCFkc9ecHu5oxWF.YS0icsSKRGcpz2HxsBrtbqUA.', '2024-04-17 09:38:24'),
('ce07a7ac-ebd2-453c-b955-ca766345913b', 'Stacey', 'Reilly', '$2a$10$kWBF4L2JbTtvrfrbqjGlTe0Gf9OZouP5AIMW2qKe0l6FSzUgWpmnq', '2024-04-17 02:38:24'),
('d49e1d69-e19e-4f74-8f7c-2a09cb6ca1c6', 'Celine', 'Fadel', '$2a$10$ha3eYfcT5aTeNHi72Twr8eDD3B.MkzCv/pGeX/8wuTqMXa3QCs2qS', '2024-04-17 05:38:24'),
('90058aa7-0ac3-48e3-9ddf-0da06b808d79', 'Destinee', 'Brown', '$2a$10$wFk5HrthYLs0E4ggar5/weKdDHsZROxtoKtYEqgKgFolUaEquBvm.', '2024-04-17 06:38:24'),
('4aaec7ee-64d2-4ceb-b956-e00fd1003110', 'Eleonore', 'Leuschke', '$2a$10$7xgYm0tHfwIosDjDEonUJ.kpVGfY.WHJH9/zGQ4YZWmwNvhZW.ype', '2024-04-17 00:38:24'),
('90148d40-5ee5-410c-af04-2c7cddcbd896', 'Russ', 'Shields', '$2a$10$5ch2KGJSGfdKRrkZTLSGWOBtGtTVyEG5eg5u8vC8t92Mb2gv.tX42', '2024-04-17 05:38:24');

