# CHATInput
VRCHATのChatBox用入力補助プログラムです、TTS（読み上げ）機能付属しております。

## 使い方
- 基本利用
  1. Boothで実行可能なバイナリーファイルをダウンロードする、あるいは自分でソースコードからビルドしてバイナリーファイルを取得する。
  2. `ChatInput.exe`を実行する。
  3. オプションタブのオプションをチェックする。
  4. Inputタブでなにか入力してSend押すかキーボードのEnter押せば正常に送信されるはずです。
- TTS（読み上げ）
  - Depsの準備
    1. [VB-Cable](https://vb-audio.com/Cable/)をインストールする。
    2. [VoiceVox Engine](https://github.com/VOICEVOX/voicevox_engine/releases)をダウンロード＆アンパックする。
  - 使用
    1. VoiceVox Engineのフォルダーを開いて、`run.exe`を実行する。`Uvicorn running on`なんちゃらが出ればそれでおｋです、何も出ないあるいはエラーが出る場合自力で解決してください。
    2. ChatInputのオプションタブのTTS項目をOnにして、適当になにか入力して送信してWindowsの音量ミキサーで`ChatInput.exe`の出力デバイスを`CABLE Input`に設定する。
    3. VRCHATでマイクデバイスを`CABLE Output`に設定する。
    4. ほかは基本利用とおなじです
## 問題
- UI関連
  - 入力候補のいちが悪い（この問題は`go-glfw`がIMEにサポートしていないからです、修正するには他のUI ToolkitでUIをさい構築する必要があります。）
- オプション関連
  - オプションの`Realtime Send`と`Voice Control`まだ作業中なので、今のところ利用できません。

## TODO 
- バグの修正
- UIの再構築
- `OSC message forward`機能の追加
- `VoiceVox Engine`管理機能の追加
- 出力デバイス自動と手動的に選択する機能の追加
- メッセージ履歴機能の追加
- DiscordからのメッセージをChatBoxに転送する機能の追加

## 注意事項
- これは私が仕事の合間に作ったおもちゃ程度のプロジェクトなので、色々大雑把です。
- バグの報告以外はIssuesを使わないでください。
- 問題の修正はこちらが暇なときにだけさせていただきます。
- 私が使いやすいように仕込んでるので、パソコンあんまわかんない人には使いづらいかもしれません。
- めんどくさいから、UIは英語のみです。

## ライセンス
本プロジェクトは`GPL3.0`を基づき発行しております、ユーザーには使用、二次開発、二次配布などの自由があります。ただし二次開発する場合など本プロジェクトを利用していることの声明と二次開発後のソースコードを同じ`GPL3.0`でオプンソースすることが必要です。具体的なライセンスの内容は[GNU一般公衆ライセンス](https://www.gnu.org/licenses/gpl-3.0.ja.html)をご覧ください。

```
Copyright (C) 2025  Sayuki

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```

# Thanks
- [Fyne Ui toolkit](https://fyne.io/)
- [VoiceVox Engine](https://github.com/VOICEVOX/voicevox_engine)