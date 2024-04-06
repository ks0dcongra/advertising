<h1 align='center'><b>Simple advertising system/b></h1>
Implement a simplified ad serving service

<br>
<br>
<br>

# Table of Contetns
- [Features](#features)
- [API Document](#api)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installing](#installing)
  - [Run Server](#run-server)
- [Demo](#demo)

<br>
<br>

# Features
A few things you can do on Simple advertising system:
1. creating ads
2. listing ads

# Api Document
https://advertising4.docs.apiary.io/#reference/0/advertising-system/public-api

<br>
<br>

# Getting Started
## **Prerequisites**
Make sure you already have `Docker` and `Golang`.

<br>

## **Installing**
1. Clone the project and go to the project directory
```
 git clone https://github.com/ks0dcongra/advertising.git

 cd advertising
```

<br/>

3. Prepare your `.env` file. Please copy to `.env.example` and rename it to `.env` ,final modify `.env` to your local postgres configuration.

<br/>

4. Open `Docker` and run `Container`

```
docker compose up -d
```
5. Call API

refer to [API Document](#api) and call it


6. Stop `Container`
```
docker compose down
```

<br/>

7. Install dependencies
```
go mod tidy
go mod vendor
```

<br/>

8. Run Test
```
go test ./...
```
<br/>

##  **Design Concept**
1. 使用三層式架構，controller, service, repository。controller用來辨識輸入的資料有無合法，service為商業邏輯，repository負責與DB互動。
2. 基本上都是依照貴司所提供的**API 範例**進行設計的

## ** 提供 Public API 能超過 10,000 Requests Per Second 的設計** 
1. 運用k8s與Nginx
    * k8s於server使用率高的尖峰時段自動擴展出多台虛擬機，每台虛擬機皆代表一組廣告投放服務
    * 利用Nginx負載平衡的特性，將傳入的請求分發到多台已啟動廣告投放服務的虛擬機上
2. 運用Redis快取伺服器
    * 當Request第一次時將其存入Redis設置2秒的過期時間，之後2秒內若有相同request則由Redis取出資料，不必透過postgres下SQL撈取資料
---