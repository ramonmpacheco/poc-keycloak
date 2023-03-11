# poc-keycloak
Testing keycloak

doc:
https://www.keycloak.org/getting-started/getting-started-docker

> Para acessar a pagina do admin console: [http://localhost:8080/realms/master/protocol/openid-connect/auth?client_id=security-admin-console&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fadmin%2Fmaster%2Fconsole%2F&state=eacecd96-66ef-4f0a-a59b-739e572c8054&response_mode=fragment&response_type=code&scope=openid&nonce=63e17a3f-b417-48be-944d-926ee434f794&code_challenge=OvhgVA-z24jQbnR-5Qa88caznxvt3wfUY8QIH-EwQqg&code_challenge_method=S256 ](http://localhost:8080/admin/)

> Logar com admin, admin

> Subir o projeto usando o docker-compose up

> Após subir a app com o comando ```go run main.go```, acessar o endereço localhost:8081, isso irá redirecionar para a página de login no :8080

> Para ver as informações do token: [jwt.io](https://jwt.io/)
>
> Endpoints disponíveis no realm: http://localhost:8080/realms/myrealm/.well-known/openid-configuration