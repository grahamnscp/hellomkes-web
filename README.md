# Build an image to create basic SSL webserver

Generate the server certificate and key pair to use for your fqdn:
```
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

Build the image using multi-stage build Dockerfile:
```
docker build -t hellomkes-web:build .  
```

Tag as required and push:
```
docker tag hellomkes-web:build grahamh/hellomkes-web:2.0
docker push grahamh/hellomkes-web:2.0
```

Start a local container instance to test:
```
docker run -d --rm -p 8443:443 grahamh/hellomkes-web:2.0
```

Test:
```
$ curl --inscure localhost:8443
[06-28-2019 08:35:18.47 304bc6a574f7] Hello from hellomkes HTTPS web server.
```

View Certificate using openssl s_client:
```
openssl s_client -showcerts -connect localhost:8443 | grep hello
depth=0 C = UK, ST = London, L = London, O = mesoslab.io, OU = labs, CN = hellomkes.mesoslab.io, emailAddress = root@hellomkes.mesoslab.io
verify error:num=18:self signed certificate
verify return:1
depth=0 C = UK, ST = London, L = London, O = mesoslab.io, OU = labs, CN = hellomkes.mesoslab.io, emailAddress = root@hellomkes.mesoslab.io
verify return:1
 0 s:/C=UK/ST=London/L=London/O=mesoslab.io/OU=labs/CN=hellomkes.mesoslab.io/emailAddress=root@hellomkes.mesoslab.io
   i:/C=UK/ST=London/L=London/O=mesoslab.io/OU=labs/CN=hellomkes.mesoslab.io/emailAddress=root@hellomkes.mesoslab.io
subject=/C=UK/ST=London/L=London/O=mesoslab.io/OU=labs/CN=hellomkes.mesoslab.io/emailAddress=root@hellomkes.mesoslab.io
issuer=/C=UK/ST=London/L=London/O=mesoslab.io/OU=labs/CN=hellomkes.mesoslab.io/emailAddress=root@hellomkes.mesoslab.io
^C
```
