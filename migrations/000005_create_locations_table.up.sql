CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    image VARCHAR(255)
);

INSERT INTO locations (name, image)VALUES
('Jakarta','https://pict.sindonews.net/dyn/850/pena/news/2021/02/26/173/348284/beragam-nama-jakarta-sejak-tahun-397-sampai-sekarang-txm.jpg'),
('Bandung','https://images.contentstack.io/v3/assets/blt00454ccee8f8fe6b/blt7f86b6cf7e72ffe7/61bc491f85b59c201581b414/US_Bandung_ID_Header.jpg'),
('Bali','https://awsimages.detik.net.id/community/media/visual/2023/07/19/ilustrasi-pulau-bali_169.jpeg?w=1200'),
('Aceh','https://www.indonesia.travel/content/dam/indtravelrevamp/id-id/destinasi/banda-aceh/image-1.jpg'),
('Solo','https://akcdn.detik.net.id/visual/2024/03/28/wisata-solo_43.jpeg?w=720&q=90'),
('Yogyakarta','https://nnc-media.netralnews.com/IMG-Netral-News-Admin-19-PI7OABWM6U.jpg'),
('Semarang','https://upload.wikimedia.org/wikipedia/commons/thumb/7/7e/Lawang_Sewu_in_Semarang_City.jpg/1200px-Lawang_Sewu_in_Semarang_City.jpg');
