/*
   packgage interfaces
   - implementing 'User' interface
*/
package interfaces

import "github.com/reshimahendra/lbw-go/internal/domain"

type IUser interface {
    Create(input domain.User) (err error)
    Update(id string, input domain.User) (err error)
    Get(id int) (user *domain.User, err error)
}
