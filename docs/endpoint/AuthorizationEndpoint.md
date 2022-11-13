# 認可エンドポイントの仕様

## メソッド
|メソッド|必須要件|
|---| ---|
| GET  | MUST |
| POST | MAY  |

## リクエスト
### 共通
- Content-Type：application / x-www-form-urlencoded
- 値なしパラメータは省略されたように扱わなければならない
- パラメータは複数回含めることはできない

### リクエストパラメータ：  
| 項目           | 必須要件 | 備考 |
| ------------- | ------- | ----|
| response_type | 必須     | 値は "code"，"token"，"code token"のいずれか |
| client_id     | 必須     | |
| redirect_uri  | オプション| |
| scope         | オプション| |
| state         | 推奨     | CSRF対策 |
#### リクエスト例（認可コードフロー）：
GET /authorize?response_type=code&client_id=s6BhdRkqt3&state=xyz
        &redirect_uri=https%3A%2F%2Fclient%2Eexample%2Ecom%2Fcb

## レスポンス
### 共通
- Content-Type：application / x-www-form-urlencoded

### 認可コードフロー
| 項目           | 必須要件 | 備考 |
| ------------- | ------- | ----|
| code | 必須    | 認可コードの最長有効期限は10分が推奨 |
| state| 必須    | リクエストにstateが含まれる場合は必須 |

### インプリシットフロー
- 認可サーバはリフレッシュトークンを払い出してはいけない

| 項目           | 必須要件  | 備考 |
| ------------- | -------  | ----|
| access_token | 必須       |  |
| token_type     | 必須     | 特に決まりはないのでbearerにしておく． |
| expires_in  | オプション   | アクセストークン有効期限．単位は秒． |
| scope         | オプション | リクエストと同一のスコープでない場合は必須なので，常に返しておく． |
| state         | 必須      | リクエストにstateが含まれる場合は必須 |

## 例外
### レスポンスパラメータ
| 項目           | 必須要件 | 備考 |
| ------------- | ------- | ----|
| error | 必須   |  |
| state | 必須   | リクエストにstateパラメータが存在する場合は必須 |
| error_description| オプション | |
| error_uri | オプション ||

### 例外レスポンス一覧
errorには以下のコードを用いる．
| 項目           | 意味 |
| ------------- | ------- |
| invalid_request | リクエストに必須パラメーターが欠落しているか、無効なパラメーター値が含まれているか、パラメーターが2回以上含まれているか、またはその他の形式が誤っています。   |
| unauthorized_client | クライアントは、このメソッドを使用して認証コードをリクエストすることを許可されていません。|
| access_denied | リソース所有者または許可サーバーが要求を拒否しました。|
| unsupported_response_type | サーバーは、このメソッドを使用した認可コードの取得をサポートしていません。|
| invalid_scope | リクエストされたスコープは、無効、不明、または不正な形式です。|
| server_error | 許可サーバーが予期しない条件を検出したため、要求を実行できませんでした。|
| temporarily_unavailable | サーバーの一時的な過負荷またはメンテナンスのため、現在、承認サーバーはリクエストを処理できません。|

### 例外レスポンス例
HTTP/1.1 302 Found
Location: https://client.example.com/cb?error=access_denied&state=xyz

