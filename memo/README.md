# git-plusのメモ
## 使う前に
git configを設定しておく.  
最低限user.email, user.name  
`git config --global user.email "hoge@fuga"`  
`git config --global user.name "myname"`  
ローカルブランチとマッチするリモートブランチにのみpush  
`git config --global push.default matching`  
httpsのリポジトリにpushするときssh経由に置き換える  
`git config --global "url.git@github.com:.pushinsteadof" "https://github.com/"`  
## コマンド
#### Git Plus:Run
git~ から続くコマンドを実行できる
#### Git Plus:Add All
修正保存したファイルすべてをgit addできる
#### Git Plus:Commit All
Addしたファイルをすべてコミットできる
