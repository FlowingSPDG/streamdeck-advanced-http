# Advanced HTTP Client（Stream Deck プラグイン）

Stream Deck のボタンやダイヤル操作で、任意の HTTP リクエストを送信できるプラグインです。  
OBS・配信ツール・家電・自作 API など、「URL を叩けば動くもの」を Stream Deck から操作できます。

---

## ダウンロード（ここから入手してください）

最新版は GitHub Releases からダウンロードできます。

| リンク | 説明 |
| --- | --- |
| **[最新版をダウンロード（Recommended）](https://github.com/FlowingSPDG/streamdeck-advanced-http/releases/latest)** | 常に最新リリースを開きます |
| **[Releases 一覧](https://github.com/FlowingSPDG/streamdeck-advanced-http/releases)** | 過去バージョンも含めて確認できます |

### ダウンロードするファイル

Releases ページにある、次の拡張子のファイルをダウンロードしてください。

- **`*.streamDeckPlugin`**（例: `dev.flowingspdg.advancedhttp.streamDeckPlugin`）

> **注意:** ソースコード（`Source code (zip)` / `Source code (tar.gz)`）は開発者向けです。通常の利用ではダウンロード不要です。

---

## インストール手順（はじめての方向け）

1. 上記の **最新版ダウンロード** から `.streamDeckPlugin` ファイルを保存します。
2. 保存したファイルを **ダブルクリック** します。
3. Stream Deck ソフトウェアが起動し、「プラグインをインストールしますか？」のような確認が出たら **インストール / Install** を選びます。
4. Stream Deck ソフトウェアを開き、右側のアクション一覧に **「Advanced HTTP Client [FlowingSPDG]」** カテゴリが増えていることを確認します。

### うまくいかないとき

- Stream Deck ソフトウェア（Elgato Stream Deck）がインストールされているか確認してください。
- ダウンロードしたファイルが `.streamDeckPlugin` であることを確認してください（`.zip` のままではインストールできません）。
- ブラウザによってはダウンロード時に警告が出ることがあります。問題なければ「保存して続行」してください。
- インストール後にカテゴリが見えない場合は、Stream Deck ソフトウェアを一度終了して再起動してください。

---

## かんたん使い方（ボタンアクション）

1. Stream Deck ソフトウェア左側のデバイス画面で、空いているボタンを選びます。
2. 右側のアクション一覧から **HTTP REQUEST** をボタンへドラッグ＆ドロップします。
3. ボタンを選択したまま、下側（または右側）に表示される **Property Inspector（設定パネル）** で URL などを入力します。
4. Stream Deck 本体のボタンを押して動作を確認します。

### 設定画面の見かた（ボタン）

ボタン用アクションを選ぶと、次のような設定画面が表示されます。

![ボタン用 Property Inspector](docs/images/pi-key.png)

#### 各項目の意味

| 項目 | いつ使うか / 動作 |
| --- | --- |
| **METHOD** | HTTP メソッド（`GET` / `POST` / `PUT` / `DELETE` など）。相手の API 指定に合わせて選びます。分からなければまず `GET` を試してください。 |
| **URL (Press)** | **ボタンを押し込んだ瞬間** にリクエストする URL。空欄の場合、押し込み時は何もしません。 |
| **URL (Release)** | **ボタンを離した瞬間** にリクエストする URL。空欄の場合、離した時は何もしません。 |
| **BODY** | リクエスト本文（主に `POST` / `PUT` / `PATCH` で使用）。不要なら空欄のままで問題ありません。 |
| **Basic AUTH ID** | Basic 認証のユーザー名。相手サーバーが Basic 認証を要求する場合に入力します。 |
| **Basic AUTH Password** | Basic 認証のパスワード。ID とセットで入力してください。 |
| **Authorization Header** | Authorization ヘッダー用の値を入力する項目です。 |
| **Enable "OK" Notification** | オンにすると、リクエスト成功時に Stream Deck 上へ成功（OK）表示が出ます。 |
| **Enable "WARNING" Notification** | オンにすると、リクエスト失敗時に Stream Deck 上へ警告表示が出ます。 |

#### よくある設定例

- **押したときだけ送信したい**  
  `URL (Press)` だけ入力し、`URL (Release)` は空欄にします。
- **押したときと離したときで別 URL を呼びたい**  
  例: 押し込みで電源 ON、離したときに電源 OFF。  
  `URL (Press)` と `URL (Release)` の両方を入力します。
- **離したときだけ送信したい**  
  `URL (Release)` だけ入力し、`URL (Press)` は空欄にします。

---

## ダイヤル（Stream Deck + など）の使い方

Stream Deck + などのダイヤル対応機種では、**HTTP REQUEST(DIAL)** アクションを使えます。

1. ダイヤル枠へ **HTTP REQUEST(DIAL)** を配置します。
2. Property Inspector で、回転・押し込み・タッチなどの URL を設定します。

### 設定画面の見かた（ダイヤル）

![ダイヤル用 Property Inspector](docs/images/pi-dial.png)

#### 各項目の意味

| 項目 | いつリクエストするか |
| --- | --- |
| **METHOD** | 送信に使う HTTP メソッド（全操作で共通） |
| **URL(left)** | ダイヤルを **左に回した** とき |
| **URL(right)** | ダイヤルを **右に回した** とき |
| **URL(push)** | ダイヤルを **押し込んだ** とき |
| **URL(release)** | ダイヤルを **離した** とき（空欄なら何もしません） |
| **URL(touch)** | タッチ画面を **タップした** とき |
| **BODY** | リクエスト本文（不要なら空欄） |
| **Basic AUTH ID / Password** | Basic 認証が必要な場合に入力 |
| **Authorization Header** | Authorization ヘッダー用の値 |
| **Enable "OK" / "WARNING" Notification** | 成功・失敗時の画面表示 ON/OFF |

> **補足:** ダイヤルを速く回すと、回転量に応じて同じ方向のリクエストが複数回送られることがあります。

---

## 動作のポイント（非技術者向けまとめ）

- **URL が空欄の操作は実行されません。** 必要な操作の URL だけ入力すれば OK です。
- **METHOD と BODY は、そのアクション内の全操作で共通**です。
- **成功/失敗の表示**は通知チェックで切り替えられます。表示がうるさい場合は OFF にしてください。
- 相手サービス側で「この URL を叩けば動く」ことが確認できていれば、その URL をそのまま設定できます。

---

## 動作環境

- Windows 10 以降 / macOS 10.11 以降
- Elgato Stream Deck ソフトウェア **6.9 以降**
- 対応アクション:
  - 通常ボタン: **HTTP REQUEST**
  - ダイヤル（Encoder）: **HTTP REQUEST(DIAL)**

---

## 開発者向け（ビルド）

通常の利用者は、このセクションを読む必要はありません。

```bash
# 依存関係・ビルド手順は Source/code および Makefile を参照
make
```

---

## サポート・不具合報告

- Issues: https://github.com/FlowingSPDG/streamdeck-advanced-http/issues
- Repository: https://github.com/FlowingSPDG/streamdeck-advanced-http

不具合を報告する際は、次があると調査しやすいです。

- 使用している Stream Deck 本体（例: Stream Deck MK.2 / Stream Deck +）
- OS（Windows / macOS）
- プラグインのバージョン（Releases のタグ名）
- 設定した METHOD / URL（秘密情報が含まれる場合は伏せてください）

---

## License

See [LICENSE](LICENSE).
