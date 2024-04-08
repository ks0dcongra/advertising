<h1 align='center'><b>Simple advertising system/b></h1>
Implement a simplified advertising service

<br>
<br>

# Table of Contents
- [Features](#features)
- [API Document](#api-document)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installing](#installing)
  - [Run Test](#run-test)
  - [Run Docker](#run-docker)
- [Design Concept](#design-concept)
- [提供 Public API 能超過 10,000 Requests Per Second 的設計](#design-concept)

<br>
<br>

# Features
A few things you can do on Simple advertising system:
1. creating ads
2. listing ads

# API Document
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

## **Run Test**
1. use go to test
```
go test ./...
```
<br/> 

## **Run Docker**
1. Prepare your `.env` file. Please copy to `.env.example` and rename it to `.env` ,final modify `.env` to your local postgres configuration.

<br/>

2. Open `Docker` and run `Container`

```
docker compose up -d
```
3. Call API

    refer to [API Document](https://advertising4.docs.apiary.io/#reference/0/advertising-system/public-api) and call API


4. Stop `Container`
```
docker compose down
```

<br/>

#  **Design Concept**
1. 使用三層式架構，controller, service, repository。controller用來辨識輸入的資料有無合法，service為商業邏輯，repository負責與DB互動。
2. 基本上都是依照貴司所提供的**API 範例**進行設計的
3. 單元測試目前只有針對商業邏輯層(service)去撰寫

# **提供 Public API 能超過 10,000 Requests Per Second 的設計** 
1. 運用k8s與Nginx
    * k8s於server使用率高的尖峰時段自動擴展出多台虛擬機，每台虛擬機皆代表一組廣告投放服務
    * 利用Nginx負載平衡的特性，將傳入的請求分發到多台已啟動廣告投放服務的虛擬機上
2. 運用Redis快取伺服器
    * 當Request第一次時將其存入Redis設置2秒的過期時間，之後2秒內若有相同request則由Redis取出資料，不必透過postgres下SQL撈取資料
3. 建立索引
    * 由於Public API每次的查詢條件必定會有`start_at`與`end_at`，因此可以為他們兩者建立複合索引，加快搜尋效率。
---
