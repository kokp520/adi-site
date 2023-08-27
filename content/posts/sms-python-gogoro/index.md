---
weight: 300
title: "First project - automator tool"
date: 2023-08-27
# lastmod: 2023-08-26T15:58
draft: false
author: "Adi"
authorLink: "https://kokp520.github.io/"
description: ""
images: []

featuredImagePreview: "crawler-image.png"

# tags: ["content", ""]

categories: ["project"]
lightgallery: true

toc:
  auto: false
math:
  enable: true
---

Send sms notification by crawler tool with python.

<!--more-->

## 前言

此專案是我在進入軟體工程前製作的小工具，當時我只有對於program只有基本的邏輯判斷寫法，經驗只有大學用arduino自動化設備，like: `熱感應燈`、`熱感應風扇`。
<br>人們經過時會自動亮燈或是溫度太高開啟風扇的自動化設備。

<br>
使專案是我在gogoro實習時已經上手且想實現更多價值，於是試著思考如果`創造實際收益`或是`減少成本耗損`，提高廠內運轉速度以及工位周轉率，開發自動化發送簡訊的輔助工具。

![schematic diagram](crawler-image.png "illustative diagram")

## Project

### Profit

* **節省30-40萬/年**

{{< admonition info >}}
* 已該店實際工程師營收時薪計算，以及工作經驗實際平均評估，
有公定的收費時薪計算方式，以此分析。
{{< /admonition >}}

### Environment

* python3.11.4
* Selenium

### How it works
---
依照輸入客戶電話也可以自動根據隔日預約名單，利用簡易`爬蟲`讀取第三方call api已達到發送簡訊內容。<br>
內容有基本的模板，也可以由客戶端維修工程師或櫃檯服務人員自定義。

---

By inputting the customer's phone number, the system can automatically retrieve the next day's appointment list. It utilizes a simple web `crawler` to call third-party API, enabling the sending of SMS notifications.

The messages follow a basic template but can also be customized by the client, maintenance engineers, or front desk personnel.

### Github source code

[source code link](https://github.com/kokp520/automator-sms-python)

