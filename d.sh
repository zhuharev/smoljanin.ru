go build
upx smoljanin.ru
rsync -avzh --exclude data/ --exclude log/ --exclude Godeps/_workspace/pkg . god@simplecloud:/home/god/sites/smoljanin.ru
ssh god@simplecloud "cd /home/god/sites/smoljanin.ru && /home/god/sites/smoljanin.ru/build.sh"