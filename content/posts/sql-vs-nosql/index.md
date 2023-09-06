---
weight: 10
title: "SQL vs NoSQL 資料庫的選擇"
description: "compare sql vs nosql database"
subtitle: ""
date: 2023-08-31
lastmod: 2023-08-31
draft: false

images: []

tags: ['sql', 'nosql', '淺談', 'DBMS']

featuredImage: ""
featuredImagePreview: "banner.gif"
hiddenFromHomePage: false

hiddenFromSearch: false
twemoji: false
lightgallery: true
ruby: true
fraction: true
fontawesome: true
linkToMarkdown: true
rssFullText: false
categories: ["Documents"]

toc:
  enable: true
  auto: true
code:
  copy: true
  maxShownLines: 50
share:
  enable: true
  # ...
comment:
  enable: true
  # ...
library:
  css:
    # someCSS = "some.css"
    # located in "assets/"
    # Or
    # someCSS = "https://cdn.example.com/some.css"
  js:
    # someJS = "some.js"
    # located in "assets/"
    # Or
    # someJS = "https://cdn.example.com/some.js"
seo:
  images: []
  # ...
---

在後端工程領域，我們經常需要存儲和處理數據。

NoSQL和DBMS是兩種不同的數據存儲和管理方法，各有其優點和缺點。今天讓我們深入瞭解它們之間的區別。

此文章的重點會比較放在**差異**上。

<!--more-->

{{<figure src="banner.gif" title="banner" >}}


## Common Databases

SQL是一種語法，NoSQL也就是不使用語法的資料庫。

而常見的`DBMS - 關聯型數據庫管理系統`通常都會聽到**MySQL**和**postgres**。

而NoSQL則是**Mongodb**了，小弟我的第一個專案使用的就是**Mongodb**

大家通常也常拿這兩家database比較。

### SQL Feature


### sql syntax

我們常用的語法不外乎就是以下幾種：
{{<admonition>}}
* **INSERT**

* **UPDATE**

* **SELECT**

* **DELETE**

* **DROP**

* **CREATE**

* **TRUNCATE**

{{</admonition>}}

方便的是語法的結構大致是不變的，你可以理解成`文法`大部分是相同的，至少比學英文簡單多了。

* 假設我需要獲得使用者id=1的名字跟email：

```sql
select name, email from customers where id =1;
```

* 更新特定users的電話號碼：

```sql
update customers set phone_number = '123-456-7890' where id = 1;
```

## Data structure 的不同

當你了解這兩種資料結構的不同後，你就可以發現它們各自的優勢，也並沒有誰是真正所謂**比較好**的資料庫


### DBMS {#dbms}

這些資料都會是以`row` and `column`的方式儲存在資料中

而這些data會利用**`key`**與彼此有關聯，像是`foreign`, `primary key`讓這些資料串接起來已協助我們搜尋資料。

於是我們就可以利用`where`,`join on`, `count`, `sum`...等等方式搭配`index`得到我們想要的資料。

這也是為什麼我們會稱之為[`關聯型資料庫`](#dbms)。

### Unstructured Data

非結構化數據 - `unstructured data`，沒有明確的數據結構形式。

你可以依照json格式塞入任何你想要的形式，且也不用每個資料有著一樣column或row，這給了noSQL極大的**靈活性**。

### Key-Value

相對來看NoSql就簡單多了，他並沒有所謂的index, column以上雜七雜八的東西。

事實上，他就是一種資料結構的database，你可以把它當作`key-value`的儲存方式，而不是table的形式，當然這只是一種，但也是最常見的一種。


### Redis Example

 **redis** 是一個key-value的non-sql資料庫，通常拿來當作`in-memory database`，拿來儲存一些你會需要快速可以獲取的資料，或是你為了減少server請求的來儲存資料，同時這些資料就算丟失也無傷大雅的內容。
 
當我載入一個圖片很多的網站，如果我沒有把資料塞到**CACHE**中，那每次登入此頁面的時候就會要重新拿取一次，試想你每次開啟ig都要等個5秒，那使用者還不揍你。

<br>

於是我們就會需要將圖片儲存起來如果10分鐘內有人同樣對同個網頁發送請求，那我就只需要：

```php

$redis = new Redis();
$key = 'unique_key_about_images';
$image = $redis->get($key);

// or 

if(isset($keyValueStorage[$key])){
  //直接透過keyvlue回傳
    $image = $keyValueStorage[$key];
}else{
  //去往server端發送請求拿圖片
    $image = $server->getImage();
}

```

那你可能會好奇那我就都用key-value的方式儲存就好，反正這麼好用，我下面會舉個SQL優勢的例子：

### Query與資料結構的關聯性

如果今天要處理我想要一次對多個table做資料搜集該怎麼做。

如果是key-value的資料結構，我需要把所有table都拿需要的資料出來再整理。

但sql可以透過table關聯的index去簡化整個流程。

### SQL Example


{{<admonition>}}
如果我今天要處理每個部門的平均薪資：

*假設有下面幾個表:*
- **員工表(Employees)**:記錄員工信息
- **部門表(Departments)**:記錄部門信息
- **工資表(Salaries)**:記錄員工工資信息

查詢每個部門的平均工資,可以這樣寫SQL語句:

```sql

sql
SELECT
  d.dept_name,
  AVG(s.salary) AS avg_salary 
FROM
  Employees e
JOIN
  Departments d ON e.dept_id = d.dept_id
JOIN
  Salaries s ON e.emp_id = s.emp_id
GROUP BY
  d.dept_name

```
{{</admonition>}}
這條SQL語句實現了以下功能：

1. 從員工表、部門表和薪資表中get數據

2. 通過員工表e的`dept_id`欄位和部門表d的`dept_id`欄位join兩個表

3. 同時通過員工表的`emp_id`和工資表的`emp_id`聯結

4. 從部門表d中選擇`dept_name`欄位

5. 從工資表中選擇工資欄位計算`mysql，起別名為avg_salary`

6. 根據部門名稱`dept_name`對結果分組

那麼這條語句得到的結果是：

**每個部門的平均工資，總共會返回部門數量總計條記錄，每條記錄包括部門名稱和該部門的平均工資。**

## Compare to Different Statements

剛剛只是**資料搜尋**的部分

所以你真的很難去說sql比較好或是甚至是新手比較適合什麼等等

所以我想要再提幾個有關sql的點

**Structured Data**的好處可以強制資料型態為何！！

這點很重要因為當我們在一個團隊裡做事的時候，常會遇到習慣不同的同事 :laughing:

**不要小看這一點小時間，在協作過程中這可以節省極大的時間來溝通。**

當我們去閱讀前人寫的code的時候，使用DBMS可以`明確`的知道每個column該放什麼樣的資料，以及每個資料之間的關聯性為何，也可以創建更嚴謹的資料。

**但同時也會需要更複雜的query，但有可能更快**

### Writing And Reading

但如果我們今天的需求會需要`快速開發`：

no-sql query依照資料型態的**複雜度**，可能會讓我們的query變慢<br><br>

但同時可以讓工程師**設計**以及**coding的時間縮短**

但反之**reading的部分就會相對來說是未來有可能的雷區**。

json的名稱大小寫，value該存array還是int

如果有需要儲存`時間`，那他的資料格式會不會有insert的時候不同的情況。

...等等

**(當然也有可能寫得很好，一切都是相對)**

--- 
{{<admonition warning>}}
* sql 可能會在設計或query的寫法比較複雜，但時間大部分來說會比較快。

* no-sql在大部分來說設計的時候會相對快，query寫法簡單但效率較慢，閱讀起來也說不定會比較困難一些。
{{</admonition>}}

### Connection 問題

**用戶連接時**

### No-SQL - Connect
no-sql比較好handle用戶連接數高的情況

**不管是用戶寫入或讀取資料在no-sql這方便是一大優勢。**

### SQL - Connect
sql相對來說就比較需要去避免用戶對db的操作

更不用說有可能會有**大量數據的訪問**或是**高複雜的Query**

都有可能會造成**DB崩潰**或是卡死。

所以可能會需要`連線池`或是`load balancer`之類的功能協助減少**對DB的過載**。

### Across Multiple Different Servers

許多No-sql Database也較擅長在**不同的server中進行擴展**。

如果你需要在不同的Server中同步你的資料，那no-sql是相對來說比較有優勢的。

{{<admonition tip>}}
* sql在建立table時還會需要去建立`Index`，所以sql也會更適合垂直擴展。<br>
(當然水平擴展也可以，只是方法不同，並沒有一定也沒有哪個一定比較好。)
* sql資料結構的複雜性以及一致性，會導致相比no-sql複製資料來說比較有優勢。
{{</admonition>}}

**But** 如果你有比較特殊的需求或是query那sql至少可以不會***因為query***造成你的`database loading`

## 後話

雖然我都拿json來舉例，但事實上你也可以在column中存取json

也可以利用sql語法對json中的內容做操作

但當然要考慮`資料庫正規化`的內容

甚至你可以兩種sql and no-sql 都用也可以

## Simple Compare

{{<admonition note>}}
:bucket: **NoSQL數據庫：**

✔️ 靈活性：NoSQL數據庫以鍵值對、文檔、列和圖形等形式存儲數據，適用於各種數據結構。

✔️ 擴展性：易於水平擴展，可以處理大量數據和高負載。

❌ 一致性：在某些情況下，一致性可能較弱，取決於所選擇的NoSQL數據庫類型。

📖 **DBMS（關聯型數據庫管理系統）：**

✔️ 一致性：嚴格的一致性，確保數據完整性和一致性。

✔️ 複雜查詢：支持複雜的SQL查詢，適用於需要多表關聯的場景。

❌ 靈活性：對於半結構化數據，可能需要複雜的模式更改。

{{</admonition>}}

:tv: **總結：**
NoSQL適用於需要高度靈活性和可擴展性的場景，如大數據和實時數據處理。然而，DBMS在需要強一致性和複雜查詢的傳統業務應用中表現出色。

無論您選擇哪種數據存儲方法，都需要根據項目需求和性能目標進行明智的選擇。

## Thinking

許多事情跟標準都是人們訂出來的，我認為要有能夠**動態思考**且不被框架限制住的思維，才更為重要。

很多觀念或是人們**現在**認為是 __『用這個很好』__ 的經驗，都有可能在明天因為某個大神寫的新的工具而改變。

**Thanks for watching**
