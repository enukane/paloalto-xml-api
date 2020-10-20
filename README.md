## WTF is this

PaloAlto XML API test for golang


## usage

```
#login
% go run go-paloalto-loginout-user.go go-paloalto-api.go -ipaddr 8.8.8.8 -user pauser -timeout 10

#logout
% go run go-paloalto-loginout-user.go go-paloalto-api.go -ipaddr 8.8.8.8 -user pauser -logout

#query
% go run go-paloalto-query-userid.go go-paloalto-api.go
```


using pit(https://github.com/typester/go-pit) for acquiring key and target host.
```
pa-xml-api:
  key: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  host: 192.168.1.1
```

access https://${host}/api/?type=keygen&user=${adminuser}&password=${password} for API key
