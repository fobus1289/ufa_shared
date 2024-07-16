# ufa_shared

## install service generator

```
go install github.com/fobus1289/ufa_shared/make-service@latest
```

### Example

- Create new service
```cmd
make-service --new
Enter project name: user
Enter project mod path: github.com/beccoder/user_service


output: user_service
``` 

- Add new service into existing service
```cmd
cd user_service
make-service --add
Enter project name: user_profile
Enter project mod path: github.com/beccoder/user_service

output: user_service/auth
```
