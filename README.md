# 目次
- [API一覧](#API一覧)
- [APIリクエスト・レスポンス詳細](#APIリクエスト・レスポンス詳細)

# API一覧

## 一覧系列（ログイン未定）
- [GET-全ユーザーのTODO一覧を表示](#GET-全ユーザーのTODO一覧を表示)；非優先：フォローに限定
- [GET-該当ユーザーのTODO一覧を表示](#GET-該当ユーザーのTODO一覧を表示)
- [GET-全ユーザーのゴールしたTODOリスト](#GET-全ユーザーのゴールTODOリスト)
- [GET-該当ユーザーのゴールTODOリスト](#GET-該当ユーザーのゴールTODOリスト)

## ログイン必須系列
- [GET-本人のTODO一覧を表示](#GET-本人のTODO一覧を表示)
- [POST-TODOを登録](#POST-TODOを登録)
- [DELETE-TODOを削除（論理削除）](#DELETE-TODOを削除（論理削除）)
- [POST-当日TODO完了](#POST-当日TODO完了)
- [DELETE-当日TODO完了取消](#DELETE-当日TODO完了取消)
- [PATCH-TODOをゴールに変更](#PATCH-TODOをゴールに変更)
- [GET-該当ユーザーの月別TODO達成状況取得](#GET-該当ユーザーの月別TODO達成状況取得) - 非優先：グラフで可視化

## ユーザー情報詳細系列
- [GET-本人情報詳細表示](#GET-本人情報詳細表示)
- [PATCH-ユーザー情報の更新](#PATCH-ユーザー情報の変更)
- [GET-ユーザー情報詳細表示](#GET-ユーザー情報詳細表示)


## ユーザー登録・大会
- [POST-ユーザー登録](#POST-ユーザー登録)
- [GET-ユーザーログイン](#GET-ユーザーログイン)
- [DELETE-退会（論理削除）](#DELETE-退会（論理削除）)


## ユーザー秘匿情報系列 - ＜＜非優先＞＞
- [GET-ユーザー秘匿情報表示](#GET-ユーザー秘匿情報表示)
- [PATCH-メールアドレス更新](#PATCH-メールアドレス更新)

## フォロー系列 - ＜＜非優先＞＞
- [GET-フォロー一覧](#GET-フォロー一覧)
- [DELETE-フォロー削除（物理削除）](#DELETE-フォロー削除（物理削除）)

※ゴール……完全に自分の達成したい目標に達成し、そのTODOをやる必要がなくなったこと

-----

# APIリクエスト・レスポンス詳細

## GET-全ユーザーのTODO一覧を表示
### URI
```
GET /todo{?page,limit,order}
```
### 処理概要
- 全ユーザーのTODO一覧を表示する。
- limitで各ページ上限、pageでページ数を指定できる。
デフォルトはlimit:100、page:1。クエリパラメータで取得する。
- orderは最終達成日順（last_achieved）、達成回数順（achieved_times）、最近設定された順（set）にできる。デフォルトは最終達成日順。
- ゴールは含まない


### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| page | numeric | ページ数 |  o |
| limit | numeric | ページ内表示Todo数 | o |
| order | string | 順序 |  o |


### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| TodoArray | array | 全todoのリスト| 
| TodoObj[TodoArray] | object | todo内容 |
| TodoID{TodoObj} | numeric | todoのID |
| IsDeleted{TodoObj} | boolean | 削除されたか（falseのみ表示） |
| Content{TodoObj} | string | todoの詳細 |
| CreatedAt{TodoObj} | string | todo登録日 | 
| LastAchieved{TodoObj} | string | 最終達成日（n日前） |
| User[TodoArray] | list | 所有user情報|
| UserId{User} | numeric | 所有userのID |
| UserName{User} | string | 所有ユーザー名 |
| UserHN{User} | string | 所有ユーザーのハンドルネーム |
| UserImg{User} | string | 所有ユーザーのアイコン；非優先 |
| limit | numeric | ページ内表示Todo数 |
| page | numeric | ページ数 | 
| order | string | 順序 | 

### 正常レスポンス
- ステータス：200
```json
HTTP/1.1 200 OK
{
    "TodoArray" :[
        {
            "TodoObj":{
                "TodoID": 1,
                "IsDeleted": false,
                "Content": "プログラミング",
                "CreatedAt": "2020-10-31",
                "LastAchieved": "2日前"
            },
            "User":{
                "UserId": 1,
                "UserName": "gopher0120",
                "UserHN": "Gopherくん",
                "UserImg": "cutiegopher.jpg",
            },
        },
    ],
    "limit": 100,
    "page": 1,
    "order": "Achieved_times"
}
```

### 異常レスポンス
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```

## GET-該当ユーザーのTODO一覧を表示

### URI
```
GET /:name{?order}
```
### 処理概要
- キーで取得したユーザーのTODO一覧を表示する。
- orderは最終達成日順（last_Achieved）、達成回数順（Achieved_times）、最近設定された順（set）にできる。デフォルトは最終達成日順。

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| name | string | ユーザー名 | x |
| order | string | 表示順 | o |

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| User | list | ユーザー情報 | 
| UserID{User} | numeric | ユーザーID | 
| UserName{User} | string | ユーザー名 |
| UserHN{User} | string | ユーザーのハンドルネーム |
| UserImg{User} | string | ユーザー画像 |
| TodoArray | Array | todo内容 |
| TodoID[TodoArray] | numeric | todoのID |
| IsDeleted[TodoArray] | boolean | 削除されたか（falseのみ表示） |
| Content[TodoArray] | string | todoの詳細 |
| CreatedAt[TodoArray] | string | todo登録日 |
| LastAchieved[TodoArray] | string | 最終達成日（n日前） | 
| TodayAchieved[TodoArray] | boolean | 本日達成したか |
| order | string | 表示順序 |

### 正常レスポンス
```json
HTTP/1.1 200 OK
{
    "Todo": {
        "User": {
            "UserID": 1,
            "UserName": "gopher0120",
            "UserHN": "Gopherくん",
            "UserImg": "cutiegopher.jpg"
        },
        "TodoArray": [
            {
                "TodoID" : 1,
                "IsDeleted": false,
                "Content": "プログラミング",
                "CreatedAt": "2020-10-30",
                "LastAchieved": "達成日はありません",
                "TodayAchieved": false
            }
        ]
    },
    "order": "last_achieved"
}
```

### 異常レスポンス
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```

## GET-全ユーザーのゴールTODOリスト
### URI
```
GET /goal{?page,limit}
```
### 処理概要
- 全ユーザーのゴールしたTODO一覧を表示する。
- limitで各ページ上限、pageでページ数、orderで表示順序を指定できる。
デフォルトはlimit:100、page:1、order:"last_achieved"。クエリパラメータで取得する。
- ゴールのみ表示


### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| page | numeric | ページ数 |  o |
| limit | numeric | ページ内表示Todo数 | o |
| order | string | 表示順序 | o |

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| TodoArray | array | 全ゴールリスト| 
| TodoObj[TodoArray] | object | ゴール内容 |
| TodoID{TodoObj} | numeric | ゴールしたtodoのID |
| IsDeleted{TodoObj} | boolean | 削除されたか（falseのみ表示） |
| Content{TodoObj} | string | ゴールしたtodoの詳細 |
| GoaledAt{TodoObj} | numeric | ゴール日 |
| User[TodoArray] | list | 所有user情報|
| UserId{User} | numeric | 所有userのID |
| UserName{User} | string | 所有ユーザーの名前 |
| UserHN{User} | string | 所有ユーザーのハンドルネーム |
| UserImg{User} | string | 所有ユーザーのアイコン；非優先 |
| limit | numeric | ページ内表示Todo数 |
| page | numeric | ページ数 | 
| order | string | 表示順序 |

### 正常レスポンス
- ステータス：200
```json
HTTP/1.1 200 OK
{
    "TodoArray" :[
        {
            "TodoObj":{
                "TodoID": 1,
                "IsDeleted": false,
                "Content": "プログラミング",
                "GoaledAt": "20201101",
            },
            "User":{
                "UserId": 1,
                "UserName": "gopher0120",
                "UserHN": "Gopherくん",
                "UserImg": "cutiegopher.jpg",
            },
        },
    ],
    "limit": 100,
    "page": 1,
    "order": "last_achieved"
}
```

### 異常レスポンス
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```

## GET-該当ユーザーのゴールTODOリスト

### URI
```
GET /goal/:name
```
### 処理概要
- キーで取得したユーザーのTODO一覧を表示する。
- ゴール日順に表示する。
- orderで表示順序を指定できる。デフォルトは"last_achieved"。

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| name | string | ユーザー名 | x |
| order | string | 表示順序 | o |

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| User | list | ユーザー情報 | 
| UserID{User} | numeric | ユーザーID | 
| UserName{User} | string | ユーザー名 |
| UserHN{User} | string | ユーザーのハンドルネーム |
| UserImg{User} | string | ユーザー画像 |
| TodoObj | array | todo取得 |
| TodoID[TodoArray] | numeric | todoのID |
| Content[TodoArray] | string | todoの詳細 |
| CreatedAt[TodoArray] | string | todo登録日 |
| GoaledAt[TodoArray] | string | ゴール日 | 
| order | string | 表示順序 |

### 正常レスポンス
```json
/* status: 200 */
{
    "Goal":{
        "User": {
            "UserID": 1,
            "UserName": "gopher0120",
            "UserHN": "Gopherくん",
            "UserImg": "cutiegopher.jpg",
        },
        "TodoArray": [
            {
                "TodoID" : 1,
                "Content": "プログラミング",
                "CreatedAt": "2020-10-30",
                "GoaledAt": "2020-11-10"
            }
        ]
    },
    "order": "last_achieved
}
```

## GET-本人のTODO一覧を表示

GET /mypage{?order}

該当ユーザーのTODO一覧と同じ
ただしname = 本人UserName

## POST-TODOを登録

### URI
```
POST /mypage
```
### 処理概要
- Todoリストに内容を登録する

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| Content | string | Todoの詳細 | x |

### 入力例
```json
{
    "Content": "プログラミング"
}
```

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| TodoObj | object | todo内容 |
| TodoId{TodoObj} | numeric | todoのID |
| Content{TodoObj} | string | todoの詳細 |
| CreatedAt{TodoObj} | string | todo登録日 | 
| LastAchieved{TodoObj} | string | 最終達成日（0日前） |

### 正常レスポンス
```json
HTTP/1.1 201 Created
{
    
    "TodoObj": 
        {
            "TodoID" : 1,
            "Content": "プログラミング",
            "CreatedAt": "2020-10-31",
            "LastAchieved": "達成した日はありません"
        }
}
```

### 異常レスポンス
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```

## DELETE-TODOを削除（論理削除）

### URI
```
DELETE /mypage/:id
```
### 処理概要
- 取得したキーのTodo項目をTodoリストから削除する
- ゴールしたTodoは退会しない限りDeleteできない

### リクエストパラメータ

| key | type | content |  null |
| ---- | ---- | ---- | ---- |
| TodoID | numeric | TodoのID | x |

### ステータスコード

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- | 
| TodoID | numeric | TodoのID |
| DeletedTodo | boolean | Todoが削除されたか |

### 正常レスポンス
```json
HTTP/1.1 200 OK
{
    "TodoID": 1,
    "Is_Deleted": true,
}
```

### 異常レスポンス
```json
/* status: 400 */
{
    "Error": "Bad Request."
}
```
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```

## POST-当日TODO完了
### URI
```
POST /:id/today
```
### 処理概要
- Todoリストに当日のtodo達成を登録する

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| id | string | TodoのID | x |

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| TodoObj | list | todo |
| TodoId | numeric | todoのID |
| IsDeleted | boolean | 削除されたか（falseのみ表示） |
| Content | string | todo内容 |
| CreatedAt | string | todo作成日 |
| LastAchieved | string | 達成日（今日） |
| TodayAchieved | boolean | 本日達成しているか |

### 正常レスポンス
```json
HTTP/1.1 200 OK
{
    "TodoObj": {
        "ID": 1,
        "IsDeleted": false,
        "Content": "プログラミング",
        "CreatedAt": "2020-11-17",
        "LastAchieved": "今日",
        "TodayAchieved": true
    }
    
}
```

### 異常レスポンス
```json
/* status: 400 */
{
    "Error": "Bad Request."
}
```
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```



## DELETE-当日TODO完了取消
### URI
```
POST /:id/today
```
### 処理概要
- Todoリストの当日のtodo達成を取り消す

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| id | string | TodoのID | x |

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| TodoObj | list | todo |
| TodoId | numeric | todoのID |
| IsDeleted | boolean | 削除されたか（falseのみ表示） |
| Content | string | todo内容 |
| CreatedAt | string | todo作成日 |
| LastAchieved | string | n日前に達成 |
| TodayAchieved | boolean | 本日達成しているか |

### 正常レスポンス
```json
HTTP/1.1 200 OK
{
    "TodoObj": {
        "ID": 1,
        "IsDeleted": false,
        "Content": "プログラミング",
        "CreatedAt": "2020-11-17",
        "LastAchieved": "4日前",
        "TodayAchieved": false
    }
    
}
```
### 異常レスポンス
```json
/* status: 400 */
{
    "Error": "Bad Request."
}
```
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```


## PATCH-TODOをゴールに変更

### URI
```
PATCH /:id/goal
```

### 処理概要
- Todoリストの達成状況を「ゴール」に変更する

### リクエストパラメータ

| key | type | content |  null |
| ---- | ---- | ---- | ---- |
| id | int | TodoのID | x |

### 入力例
```json
{
    "id": 1
}
```

### レスポンスパラメータ
| key | type | content | 
| ---- | ---- | ---- |
| TodoID | numeric | TodoのID |
| Goaled | boolean | Todoがゴールか |
| GoaledAt | datetime | ゴールした日 |

### 正常レスポンス
```json
HTTP/1.1 201 Created
```

### 異常レスポンス
```json
/* status: 400 */
{
    "Error": "Bad Request."
}
```
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```









## GET-該当ユーザーの月別TODO達成状況取得
※未実装
### URI
```
GET /mypage/achieved
```
### 処理概要
- キーで取得したユーザーのTODO達成状況を確認する

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| name | string | ユーザー名 | x |


### 入力例
```json
{
    "name": "gopher0120"
}
```

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| User | list | ユーザー情報 | 
| UserID{User} | numeric | ユーザーID | 
| UserName{User} | string | ユーザー名 |
| UserHN{User} | string | ユーザーのハンドルネーム |
| UserImg{User} | string | ユーザー画像 |
| TodoArray | array | todo取得 |
| TodoID[TodoArray] | numeric | todoのID |
| Content[TodoArray] | string | todoの詳細 |
| CreatedAt[TodoArray] | string | todo登録日 |
| AchievedAll[TodoArray] | numeric | 達成回数 |
| Achieved[TodoArray] | Object | 達成詳細 |
| ByYear{Achieved} | list | 年ごとの達成 |
| Year[ByYear] | numeric | 達成年 |
| TimesByYear[ByYear]
| ByMonth[ByYear] | array | 月ごとの達成 |
| Month[ByMonth] | numeric | 達成月 |
| TimesByMonth[ByMonth] | numeric | 該当月の達成回数 |


### 正常レスポンス
```json
/* status: 200 */
{
    "User": {
        "UserID": 1,
        "UserName": "gopher0120",
        "UserHN": "Gopherくん",
        "UserImg": "cutiegopher.jpg",
    },
    "TodoArray": [
        {
            "TodoID" : 1,
            "Content": "プログラミング",
            "CreatedAt": "2020-10-30",
            "AchievedAll": 30,
            "Achieved": {
                "ByYear": [
                    {
                        "Year": 2020,
                        "TimesByYear": 20,
                        "ByMonth": [                    
                            {
                                "Month": 11,
                                "TimesByMonth": 5
                            },
                            {
                                "Month": 12,
                                "TimesByMonth": 15
                            }
                        ]
                    },
                    {
                        "Year": 2021,
                        "TimesByMonth": 10,
                        "ByMonth" :[
                            {
                                "Month": 1,
                                "TimesByMonth": 10
                            }
                        ]
                    }
                ]
            }
        }
    ]
}
```

### 異常レスポンス
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```





## GET-本人情報詳細表示

[GET-ユーザー情報詳細表示](#GET-ユーザー情報詳細表示)と同じ
ただしname = UserName

## PATCH-ユーザー情報の更新

### URI
```
GET /profile
```
### 処理概要
- ユーザー情報を更新する（秘匿情報以外）

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| Name | string | ユーザーの名前 | x |
| HN | string | ユーザーのハンドルネーム | x |
| Img | string | ユーザーのアイコン；非優先 | o |
| FinalGoal | string | ユーザーの目標 | o |
| Profile | string | ユーザーのプロフィール（自由記述） | o |
| Twitter | string | ユーザーのTwitterアカウント | o |
| Instagram | string | ユーザーのInstagramアカウント | o |
| Facebook | string | ユーザーのFacebookアカウント | o |
| GitHub | string | ユーザーのGitHubアカウント | o |
| URL | string | その他ユーザーが載せたいURL | o |

### 入力例
```json
HTTP/1.1 201 Created
```

### レスポンスパラメータ
- なし

### 正常レスポンス
```json
/* status: 204 */
```

### 異常レスポンス
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```



## GET-ユーザー情報詳細表示

### URI
```
GET /profile
```
### 処理概要
- ユーザー情報の詳細を取得する

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| name | string | ユーザー名 | x |

### 入力例
```json
{
    "name": "gopher0120"
}
```

### レスポンスパラメータ

| key | type | content | 
| ---- | ---- | ---- |
| ID | numeric | ユーザーID |
| Name | string | ユーザーの名前 |
| HN | string | ユーザーのハンドルネーム |
| Img | string | ユーザーのアイコン；非優先 |
| FinalGoal | string | ユーザーの目標 |
| Profile | string | ユーザーのプロフィール（自由記述） |
| Twitter | string | ユーザーのTwitterアカウント |
| Instagram | string | ユーザーのInstagramアカウント |
| Facebook | string | ユーザーのFacebookアカウント |
| GitHub | string | ユーザーのGitHubアカウント |
| URL | string | その他ユーザーが載せたいURL |

### 正常レスポンス
```json
HTTP/1.1 200 OK
{
    "ID": 1,
    "Name": "gopher0120",
    "HN": "Gopherくん",
    "Img": "cutiegopher.jpg",
    "FinalGoal": "Golangの神になりたい！！",
    "Profile": "僕はGopher。Golangが大好き！最近Goで参加する競技プログラミングのYouTubeチャンネル始めました。Golangがもっと広まると嬉しいな！",
    "Twitter": "go",
    "Instagram": "go",
    "Facebook": "go",
    "Github": "go",
    "URL": "http://www.cutiegophergogogo.com/"
}
```

### 異常レスポンス
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```



## POST-ユーザー登録
### URI
```
POST /register
```
### 処理概要
- ユーザー登録をする。
- 固有のユーザー名、メールアドレス、パスワードが必須。
- HNがない場合は、ユーザー名がそのまま使用される。
- 詳細情報は空になる。

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| Name | string | ユーザー名 | x |
| HN | string | ハンドルネーム | o |
| MailAddress | string | メールアドレス | o |
| Password | string | パスワード | x | 

### 入力例
```json
{
    "Name": "gopher0120",
    "HN": "Gopherくん",
    "MailAddress": "cutegopher@gophergogo.com",
    "Password": "golanggggggg"
}
```

### レスポンスパラメータ

```json
HTTP/1.1 302 redirect GET / 
```

### 異常レスポンス
```json
/* status 400 */
{
    "Error": "Bad Request."
}
```
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}

```






## GET-ユーザーログイン

### URI
```
GET /login
```

### 処理概要
- ユーザーのログイン
- nameとpasswordで認証する
- クッキーで情報を保存する
- セッション中はログインをキープする

### リクエストパラメータ

| key | type | content | null |
| --- | --- | --- | --- | 
| name | string | ユーザーネーム | x |
| password | string | パスワード | x |

### 正常レスポンス

```json
HTTP/1.1 302 redirect 
GET /mypage
```

### 異常レスポンス

```json
HTTP/1.1 〜〜〜 
```



## GET-ユーザーログアウト

### URI
```
DELETE /logout
```

### 処理概要
- ログアウトする。
- cookieからログイン情報を削除する

### レスポンスパラメータ
```
HTTP/1.1 204 No Content
```

### 異常レスポンス
```
HTTP/1.1 404 Not Found
```

## DELETE-退会（論理削除）
### URI
```
DELETE /delete
```
### 処理概要
- ユーザーを削除する（論理削除）

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| Password | string | パスワード | x |

### 正常レスポンス
```json
HTTP/1.1 302 Redirect
GET /
```
### 異常レスポンス
```json
/* status 400 */
{
    "Error": "Bad Request."
}
```
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}
```

## GET-ユーザー秘匿情報表示
※未実装
### URI
```
GET /secret
```
### 処理概要
- 秘匿情報を表示する（本人のみ）

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| Password | string | パスワード | x |

### 入力例
```json
{
    "Password": "golanggggggg"
}
```

### レスポンスパラメータ

| key | type | content |
| ---- | ---- | ---- |
| ID | numeric | ユーザーID |
| Name | string | ユーザー名 |
| HN | string | ハンドルネーム |
| MailAddress | string | メールアドレス |

### 正常レスポンス
```json
/* status 200 */
```json
{
    "ID": 1,
    "Name": "gopher0120",
    "HN": "Gopherくん",
    "MailAddress": "cutegopher@gophergogo.com",
}
```
### 異常レスポンス
```json
/* status 400 */
{
    "Error": "Bad Request."
}
```
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}
```

## PATCH-メールアドレス更新
※未実装
### URI
```
PATCH /secret
```
### 処理概要
- メールアドレス を更新する（本人のみ）

### リクエストパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| MailAddress | string | メールアドレス | x |
| Password | string | パスワード | x |

### 入力例
```json
{
    "MailAddress": "go@go.com",
    "Password": "golanggggggg"
}
```

### レスポンスパラメータ

| key | type | content | null |
| ---- | ---- | ---- | ---- |
| ID | numeric | ユーザーID |
| Name | string | ユーザー名 |
| HN | string | ハンドルネーム |
| MailAddress | string | メールアドレス |

### 正常レスポンス
```json
/* status 200 */
```json
{
    "ID": 1,
    "Name": "gopher0120",
    "HN": "Gopherくん",
    "MailAddress": "cutegopher@gophergogo.com",
}
```
### 異常レスポンス
```json
/* status 400 */
{
    "Error": "Bad Request."
}
```
```json
/* status: 404 */
{
    "Error": "Not Found."
}
```
```json
/* status: 500 */
{
    "error": "Server Error."
}
```


## GET-フォロー一覧
※未実装

## DELETE-フォロー削除（物理削除）
※未実装

