# wordbook
# git stash test ok

# 画面設計メモ
1.top画面 - ボタン押下でログイン画面へ
2.ログイン画面 - pw, user_name(unique)でログイン + create&sign up
3.home画面 - 単語帳一覧&選択で単語帳リスト画面へ
4.単語帳リスト画面 - home画面で選択した単語帳表示
    ・問題→・解答　繰り返し
    ・終了時ポップアップ表示　okボタン押下でhome画面へ

# テーブル
○user_tb
・user_id(pk) ・user_name ・user_pw

○wordbook_tb
・wordbook_id(pk) ・wordbook_name ・wordbook_kind(num)

○wordbook_kind_tb ※wordbook_tb.wordbook_kindの対応表

○wordbook_words ※単語帳の総数は COUNTで取得する。
・wordbook_words_id(pk) ・wordbook_id(fk) ・word ・description

○user_wordbook_correct
・user_wordbook_correct_id(pk) ・user_id(fk) ・wordbook_id(fk) ・correct