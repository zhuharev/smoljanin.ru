SERVER=simplecloud
USER=god
APPNAME=smoljanin.ru

go build -o $APPNAME
upx smoljanin.ru

ssh $USER@$SERVER "mkdir -p /home/$USER/apps/$APPNAME"
rsync -avzh --exclude scripts/install.sh . $USER@$SERVER:/home/$USER/apps/$APPNAME
ssh $USER@$SERVER << EOF
  cd /home/$USER/apps/$APPNAME
  sudo mv scripts/$APPNAME.conf /etc/init
  sudo mv scripts/$APPNAME_caddy.conf /etc/caddy/sites
  sudo stop caddy
  sudo start caddy
  sudo stop $APPNAME
  sudo start $APPNAME
EOF