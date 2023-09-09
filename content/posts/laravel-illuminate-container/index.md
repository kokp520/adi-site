---
weight: 260
title: "Laravel Singleton and Binding"
description: "Laravel Singleton and Binding by illuminate/container"
subtitle: ""
date: 2023-09-09
lastmod: 2023-09-09
draft: false

images: []

tags: ["依賴注入", "DI", "singleton", "bind", DependencyManager]

featuredImage: ""
featuredImagePreview: "banner.png"

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
---

Laravel Dependency Injecttion MainFunction, talkeing about why so important about DI and what is different between Not using DI project...

<!--more-->

## illuminate Container

{{<figure src="banner.png" title="" >}}

### Simple Talk

[illuminate\Container - github](https://github.com/illuminate/container)

今天要講個再設計 Laravel 系統時常用的 Component - **Container**

他負責的工作是：

- 管理服務容器 **Server Container**
- 依賴注入**Dependency Injection**

都是 Laravel 非常強大的功能。

### Functional

你可以想像成一個科技大櫃子 :u6709:

你可以把你需要的功能以及使用方法用 container 共同管理。

當你下次要使用什麼 **功能** 時，直接開櫃子去使用，科技的地方在於當你他能夠依照你的環境去給你適合東西。

什麼意思呢...

### Example

假設，你今天在學校有個櫃子，裡面有很多個課本。

```php
/*
文件上標準寫法是serviceProvider，所以這邊用provider示範
也可以改成Illuminate\Container\Container 綁定interface
 */
class Cabinet extends Illuminate\Support\ServiceProvider
{
    public function register() {
        $this->bind(BookInterface::class, function ($app) {
          /* 依照架構可以寫在產品底下的App用於區分，就可以不用 if 判斷 */
            if (app()->isLocale('en')) {
                return new EnglishBook();
            } else {
                return new MathBook();
            }
        });
    }
}

interface BookInterface
{
    public function getBook();
}

class MathBook implements BookInterface
{
    public function getBook()
    {
        return $this->book = 'MathBook';
    }
}

class EnglishBook implements BookInterface
{
    public function getBook()
    {
        return $this->book = 'EnglishBook';
    }
}

```

### Instance

**從底層去切 -**

當你今天在上數學課，你從抽體內拿出"課本"，他會自動給你國文課本。

```php
/* 建立instance -> make book */
$cabinet = \Cabinet::getInstance();
$book = $cabinet->make(BookInterface::class);
$book->getBook(); // MathBook
```

當你今天要上英文課，你要抽屜拿出“課本”，他會拿英文課本給你。

```php
$cabinet = \Cabinet::getInstance();
$book = $cabinet->make(BookInterface::class);
$book->getBook(); // EnglishBook
```

好處就是如果你以後有換教室(**產品/底層**)的需求，不需要改變作法，只需變更環境換新的櫃子就好。

只要你定義清楚，你以後可以依照 Interface，去拿到該拿的物件

### Container 功能

{{<admonition note>}}

- 依賴注入：它允許您定義類之間的依賴關係，並自動解析和注入這些依賴。這有助於實現鬆散耦合的代碼，並提高測試性。

- 服務容器：服務容器是一個管理類實例的容器，可以用來管理應用程序中的各種服務，例如數據庫連接、郵件服務、緩存等。容器允許您輕鬆地獲取這些服務的實例。

- 綁定：您可以使用容器來定義抽象類或接口與其具體實現之間的綁定。這就是所謂的 Binding，它確保當您需要一個特定的類時，容器將返回相對應的實例。

- 單例（Singleton）：容器還支持將類設置為 Singleton，確保該類的實例只創建一次，並在後續請求中返回相同的實例。

{{</admonition>}}

### Singleton

- **單例模式 singleton Pattern**

在我們實際創建一個 Object 之後，我們其實會希望可以使用同一個實例，既可以改變參數，也不用擔心重複創建 Object 的問題。

於是

```php
/* 當我們創建一個櫃子服務器 */
$cabinet = new Cabinet();
$cabinet->getBook();
$cabinet->addBook('ChineseBook'); // 新的書

/* 如果你需要從其他地方call到chinesebook你要怎麼確定這個物件裡面是有chinesebook的物件 */
$cabinet = new Cabinet();
$cabinet->getBook('ChineseBook'); // 找不到此書
//因為是新物件 他身上是沒有chinesebook這本書的
```

```php

/* singleton - 將Cabinet綁定就是此物件 */
class App
{
  $this->singleton('Cabinet', function (){
      return new Cabinet();
    });
}

/* 其他地方要使用Cabinet */
class AnotherPlace
{
  $app = \App->getInstance();
  $app->make('Cabinet');
  $app->getBook('ChineseBook'); //可以得到這本書
}

```

**單例模式來確保只有一個 Cabinet 實例存在並在不同地方共享使用。**
