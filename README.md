# Advanced HTTP Client（Stream Deckプラグイン）

Stream Deckのボタン／ダイヤル操作から任意のHTTPリクエストを送信するプラグインです。  
インターネット上のHTTP APIやREST APIなどに対して指定のHTTPリクエストを送信可能です。

---

## ダウンロード

**[ダウンロード](https://github.com/FlowingSPDG/streamdeck-advanced-http/releases/latest)**

Releasesページの `*.streamDeckPlugin`（例: `dev.flowingspdg.advancedhttp.streamDeckPlugin`）を取得してください。  
`Source code (zip/tar.gz)`はソース配布用です。

---

## インストール

1. ダウンロードした `.streamDeckPlugin`をダブルクリックする
2. Stream Deckソフトウェアの確認ダイアログでインストールを実行する
3. アクション一覧に **Advanced HTTP Client [FlowingSPDG]** が表示されれば完了

---

## 使い方

1. Stream Deckソフトウェアで空きボタンを選択する
2. アクション一覧から **HTTP REQUEST** をドラッグ＆ドロップする
3. Property Inspectorで METHOD / URLなどを設定する
4. 本体のボタンで動作を確認する

### 設定項目（ボタン）

![Property Inspector（ボタン）](docs/images/pi-key.png)

| 項目 | 説明 |
| --- | --- |
| **METHOD** | HTTPメソッド（`GET` / `HEAD` / `QUERY` / `POST` / `PUT` / `DELETE` / `PATCH`） |
| **URL (Press)** | ボタン押下時に送信するURL。空欄なら送信しない |
| **URL (Release)** | ボタン解放時に送信するURL。空欄なら送信しない |
| **BODY** | リクエスト本文（`QUERY` / `POST` / `PUT` / `PATCH`などで使用） |
| **Basic AUTH ID** | Basic認証のユーザー名 |
| **Basic AUTH Password** | Basic認証のパスワード |
| **Authorization Header** | Authorizationヘッダーに設定する値 |
| **Enable "OK" Notification** | 成功時にOK表示を出す |
| **Enable "WARNING" Notification** | 失敗時に警告表示を出す |

Press / Releaseは別URLを指定できます。片方のみ使う場合は、使わない側を空欄にしてください。

---

## ダイヤル（Stream Deck +など）

ダイヤル対応機種では **HTTP REQUEST(DIAL)** を利用できます。

1. ダイヤル枠に **HTTP REQUEST(DIAL)** を配置する
2. Property Inspectorで各操作のURLを設定する

### 設定項目（ダイヤル）

![Property Inspector（ダイヤル）](docs/images/pi-dial.png)

| 項目 | 説明 |
| --- | --- |
| **METHOD** | HTTPメソッド（全操作で共通） |
| **URL(left)** | 左回転時 |
| **URL(right)** | 右回転時 |
| **URL(push)** | 押し込み時 |
| **URL(release)** | 解放時（空欄なら送信しない） |
| **URL(touch)** | タッチパネルのタップ時 |
| **BODY** | リクエスト本文 |
| **Basic AUTH ID / Password** | Basic認証 |
| **Authorization Header** | Authorizationヘッダーに設定する値 |
| **Enable "OK" / "WARNING" Notification** | 成功／失敗時の表示 |

回転量に応じて、同一方向のリクエストが複数回送信される場合があります。

---

## 動作環境

- Windows 10以降 / macOS 10.11以降
- Elgato Stream Deckソフトウェア 6.9以降
- アクション: **HTTP REQUEST**（ボタン）、**HTTP REQUEST(DIAL)**（ダイヤル）

---

## サポート

- Issues: https://github.com/FlowingSPDG/streamdeck-advanced-http/issues
- Repository: https://github.com/FlowingSPDG/streamdeck-advanced-http

不具合報告時は、本体機種・OS・プラグインバージョン・METHOD / URLと共にIssueまでご報告ください。

---

## License

See [LICENSE](LICENSE).
