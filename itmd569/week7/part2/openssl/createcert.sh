mkdir certs
rm certs/*
echo "make server cert"
openssl req -new -nodes -x509 -out certs/server.pem -keyout certs/server.key -days 360 -subj "/C=DE/ST=NRW/L=Earth/O=itmd469569/OU=IT/CN=www.coolwebsite.com/emailAddress=example@gmail.com"
echo "make client cert"
openssl req -new -nodes -x509 -out certs/client.pem -keyout certs/client.key -days 360 -subj "/C=DE/ST=NRW/L=Earth/O=itmd469569/OU=IT/CN=www.coolwebsite.com/emailAddress=example@gmail.com"