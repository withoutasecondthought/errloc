package errloc

import "fmt"

// AddLoc add new error location. Use only in the place, that returns error.
// Also, you can add object with method or  function that you called for that use space bar
//
// Example
//
//		func(r *Repository) GetUsers(...) ([]User, error) {
//			users, err := r.db.Get(...)
//			if err != nil {
//				return []Users{}, errloc.AddLoc(
//				"users/repository",
//				fmt.Errorf("cannot get users: %w", err)
//				)
//				//or
//				return []Users{}, errloc.AddLoc(
//				"users/repository Repository.GetUsers()",
//				fmt.Errorf("cannot get users: %w", err)
//				)
//			}
//	    ...
//	 {
//	Output
//
// users/repository: cannot get users: some error
//
// users/repository Repository.GetUsers(): cannot get users: some error
func AddLoc(currentPackage string, handlingError error) error {
	return fmt.Errorf("%s: %w", currentPackage, handlingError)
}

// AppendLoc add new intermediate location. Use in the places where you just return error from you another
// already wrapped function
//
// Example
//
//		func(s *Service) GetUsers(...) ([]User, error) {
//			users, err := r.db.Get(...)
//			if err != nil {
//				return []Users{}, errloc.AppendLoc("accounts", err)
//			}
//	    ...
//	 {
//	Output
//
// accounts/users/repository: cannot get users: some error
func AppendLoc(currentPackage string, handlingError error) error {
	return fmt.Errorf("%s/%w", currentPackage, handlingError)
}
