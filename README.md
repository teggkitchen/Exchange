# 作業 外幣API



## 創建資料庫

這裡使用Docker MySQL

docker指令：

```cmd
docker run --name buyapi_mysql -p 3307:3306 --net=aappii -v /Users/tingk/DockerProject/buyapi-mysql/mysql-data:/var/lib/mysql -v /Users/tingk/DockerProject/buyapi-mysql/mysql-config:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=12345600 -d mysql:8.0.12 --character-set-server=utf8 --collation-server=utf8_unicode_ci --init-connect='SET NAMES UTF8;'
```

## 資料庫結構

```cmd
moneys Table
|---id
|---name
|---created_at
|---updated_at
```

```cmd
current_markets Table
|---id 
|---money_id 
|---buy
|---sell
|---created_at
|---updated_at
```

```cmd
historical_markets Table
|---id 
|---money_id 
|---buy
|---sell
|---created_at
|---updated_at
```


## 創建資料庫步驟

創建資料庫：
```cmd
create database MONEY character set utf8;
```

使用資料庫：
```cmd
use MONEY;
```

創建moneys資料表：
```cmd
CREATE TABLE moneys(
id  INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(50),
created_at DATETIME DEFAULT NULL,
updated_at DATETIME DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```


創建current_markets資料表：
```cmd
CREATE TABLE current_markets(
id  INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
money_id INT,
buy decimal(10,6),
sell decimal(10,6),
created_at DATETIME DEFAULT NULL,
updated_at DATETIME DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```


創建historical_markets資料表：
```cmd
CREATE TABLE historical_markets(
id  INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
money_id INT,
buy decimal(10,6),
sell decimal(10,6),
created_at DATETIME DEFAULT NULL,
updated_at DATETIME DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
```



## 運行Go

```cmd
go run main.go
```

## 使用postman
<!-- [Postman檔案](https://github.com/teggkitchen/buyapi/blob/master/postman/BuyApi.postman_collection.json) -->
<a href="https://github.com/teggkitchen/buyapi/blob/master/postman/BuyApi.postman_collection.json" download="postman.json">Postman檔案
</a>

## 備註

因一些環境參數的因素

1. 建議使用go run main.go</br>
2. docker不使用network以免影響sql的連線

