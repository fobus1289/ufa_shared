# ufa_shared

## install service generator

```
go install github.com/fobus1289/ufa_shared/make-service@latest
```

### Example

- Create new service
```cmd
make-service --new user

output: user_service
``` 

- Add new service into existing service
```cmd
cd user_service
make-service --add author

output: user_service/auth
```
