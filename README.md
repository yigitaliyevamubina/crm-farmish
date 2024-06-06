# Animal Management API

Bu API hayvonlarni boshqarish tizimini yaratish uchun ishlatiladi. Quyida mavjud endpointlar va ularning tavsifi keltirilgan.

## Bazaviy URL


## Animal Types

### <span style="color: green;">POST</span> Create Animal Type
- **URL:** `/animal-type`
- **Method:** `POST`
- **Description:** Yangi hayvon turini yaratish.

### <span style="color: blue;">GET</span> Get Animal Type
- **URL:** `/animal-type/get`
- **Method:** `GET`
- **Description:** Berilgan ID bo'yicha hayvon turini olish.

### <span style="color: blue;">GET</span> List Animal Types
- **URL:** `/animal-type`
- **Method:** `GET`
- **Description:** Barcha hayvon turlarini ro'yxatini olish.

### <span style="color: orange;">PUT</span> Update Animal Type
- **URL:** `/animal-type`
- **Method:** `PUT`
- **Description:** Hayvon turini yangilash.

### <span style="color: red;">DELETE</span> Delete Animal Type
- **URL:** `/animal-type`
- **Method:** `DELETE`
- **Description:** Hayvon turini o'chirish.

## Animals

### <span style="color: green;">POST</span> Create Animal
- **URL:** `/animals`
- **Method:** `POST`
- **Description:** Yangi hayvon yaratish.

### <span style="color: blue;">GET</span> Get Animal
- **URL:** `/animals/get`
- **Method:** `GET`
- **Description:** Berilgan ID bo'yicha hayvonni olish.

### <span style="color: blue;">GET</span> List Animals
- **URL:** `/animals`
- **Method:** `GET`
- **Description:** Barcha hayvonlarni ro'yxatini olish.

### <span style="color: orange;">PUT</span> Update Animal
- **URL:** `/animals`
- **Method:** `PUT`
- **Description:** Hayvonni malumotlarini yangilash.

### <span style="color: red;">DELETE</span> Delete Animal
- **URL:** `/animals`
- **Method:** `DELETE`
- **Description:** Hayvonni o'chirish.

## Food Warehouse

### <span style="color: green;">POST</span> Create Food Warehouse
- **URL:** `/food`
- **Method:** `POST`
- **Description:** Yangi oziq-ovqat ombori yaratish.

### <span style="color: blue;">GET</span> Get Food Warehouse
- **URL:** `/food/get`
- **Method:** `GET`
- **Description:** Berilgan ID bo'yicha oziq-ovqat omborini olish.

### <span style="color: blue;">GET</span> List Food Warehouse
- **URL:** `/food`
- **Method:** `GET`
- **Description:** Barcha oziq-ovqat omborlarini ro'yxatini olish.

### <span style="color: orange;">PUT</span> Update Food Warehouse
- **URL:** `/food`
- **Method:** `PUT`
- **Description:** Oziq-ovqat omborini yangilash.

### <span style="color: red;">DELETE</span> Delete Food Warehouse
- **URL:** `/food`
- **Method:** `DELETE`
- **Description:** Oziq-ovqat omborini o'chirish.

## Medicine Warehouse

### <span style="color: green;">POST</span> Create Medicine Warehouse
- **URL:** `/medicine`
- **Method:** `POST`
- **Description:** Yangi dori ombori yaratish.

### <span style="color: blue;">GET</span> Get Medicine Warehouse
- **URL:** `/medicine/get`
- **Method:** `GET`
- **Description:** Berilgan ID bo'yicha dori omborini olish.

### <span style="color: blue;">GET</span> List Medicine Warehouse
- **URL:** `/medicine`
- **Method:** `GET`
- **Description:** Barcha dori omborlarini ro'yxatini olish.

### <span style="color: orange;">PUT</span> Update Medicine Warehouse
- **URL:** `/medicine`
- **Method:** `PUT`
- **Description:** Dori omborini yangilash.

### <span style="color: red;">DELETE</span> Delete Medicine Warehouse
- **URL:** `/medicine`
- **Method:** `DELETE`
- **Description:** Dori omborini o'chirish.

## Feeding

### <span style="color: green;">POST</span> Create Feeding
- **URL:** `/feeding`
- **Method:** `POST`
- **Description:** Bir hayvonni ovqatlantirish.

### <span style="color: blue;">GET</span> Get Feeding
- **URL:** `/feeding/get`
- **Method:** `GET`
- **Description:** Berilgan ID bo'yicha ovqatlantirilganligi haqida malumot olish.

### <span style="color: blue;">GET</span> List Feeding
- **URL:** `/feeding`
- **Method:** `GET`
- **Description:** Barcha ovqatlantirish yozuvlarini ro'yxatini olish.

### <span style="color: blue;">GET</span> Not Feeding
- **URL:** `/feeding/not-feeding`
- **Method:** `GET`
- **Description:** Ovqatlantirilmagan hayvonlarni olish.

### <span style="color: blue;">GET</span> List Feeding by Animal ID
- **URL:** `/feeding/animal-id`
- **Method:** `GET`
- **Description:** Hayvon ID bo'yicha ovqatlantirilganligi haqida malumot olish.

## Watering

### <span style="color: green;">POST</span> Create Watering
- **URL:** `/watering`
- **Method:** `POST`
- **Description:** Bir hayvonni ID boyicha sug'orish.

### <span style="color: blue;">GET</span> Get Watering
- **URL:** `/watering/get`
- **Method:** `GET`
- **Description:** Berilgan ID bo'yicha sug'orishlarni olish.

### <span style="color: blue;">GET</span> Not Watering
- **URL:** `/watering/not-watering`
- **Method:** `GET`
- **Description:** Sug'orilmagan hayvonlarni olish.

## Treatment

### <span style="color: green;">POST</span> Create Treatment
- **URL:** `/treatment`
- **Method:** `POST`
- **Description:** Bir hayvonni IDsi boyicha davolash yaratish.

### <span style="color: blue;">GET</span> Get Treatment
- **URL:** `/treatment/get`
- **Method:** `GET`
- **Description:** Berilgan ID bo'yicha davolash royhatini olish.

### <span style="color: blue;">GET</span> List Treatment
- **URL:** `/treatment/`
- **Method:** `GET`
- **Description:** Barcha davolash ro'yxatini olish.

### <span style="color: blue;">GET</span> List Treatment by Animal ID
- **URL:** `/treatment/animal-id`
- **Method:** `GET`
- **Description:** Hayvon ID bo'yicha davolash royhatini olish.

### <span style="color: blue;">GET</span> List Treatment by Medicine ID
- **URL:** `/treatment/medicine-id`
- **Method:** `GET`
- **Description:** Dori ID bo'yicha davolash malumotlarini olish.

## API dan foydalanish bo'yicha qo'shimcha ma'lumotlar

database nomi -> farmish

parol -> 20030505

Project

host = localhost    
port = 9050
