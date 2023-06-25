# ERRLOC

---

## In SHORT 
small lib to add error location in your application

---

## why choose it

* don't use reflect
* less code
* less errors
* standardized output
* std lib compatibility

---

## Examples:

### `AddLoc(currentPackage string, handlingError error) error`
#### Example:
```go
func(r *Repository) GetUsers(...) ([]User, error) { 
	users, err := r.db.Get(...)
		if err != nil {
			return []Users{}, errloc.AddLoc(
			"users/repository",
			fmt.Errorf("cannot get users: %w", err)
			)
			//or
			return []Users{}, errloc.AddLoc(
			"users/repository Repository.GetUsers()",
			fmt.Errorf("cannot get users: %w", err)
			)
			}
	...
{
```

#### Output:

    users/repository: cannot get users: some error
    or
    users/repository Repository.GetUsers(): cannot get users: some error

---
### `AppendLoc(currentPackage string, handlingError error) error`
#### Example:
```go
func(s *Service) GetUsers(...) ([]User, error) {
	users, err := r.db.Get(...)
	if err != nil {
		return []Users{}, errloc.AppendLoc("accounts", err)
	}
	...
{
```

#### Output:

    accounts/users/repository: cannot get users: some error


## Welcome to any suggestions
### write it in issues



