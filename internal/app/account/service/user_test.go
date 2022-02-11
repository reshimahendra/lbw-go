package service

import (
	"testing"
	"time"

	"github.com/google/uuid"
	d "github.com/reshimahendra/lbw-go/internal/domain"
	E "github.com/reshimahendra/lbw-go/internal/pkg/errors"
	"github.com/reshimahendra/lbw-go/internal/pkg/helper"
	"github.com/stretchr/testify/assert"
)

var (
    uHeader = []string{"id","username","first_name","last_name","email",
        "status_id","role_id","created_at","updated_at"}

    u = []*d.User{
        {
            ID        : uuid.New(),
            Username  : "leonard",
            Firstname : "Leo",
            Lastname  : "Singa",
            Email     : "leo@gmail.com",
            PassKey   : "secret",
            StatusID  : 1,
            RoleID    : 1,
            CreatedAt : time.Now(),
            UpdatedAt : time.Now(),
        },
        {
            ID        : uuid.New(),
            Username  : "jennydoe",
            Firstname : "Jenny",
            Lastname  : "Doe",
            Email     : "jenny@gmail.com",
            PassKey   : "secret",
            StatusID  : 1,
            RoleID    : 0,
            CreatedAt : time.Now(),
            UpdatedAt : time.Now(),
        },
    }
)

// mockUserService is wrapper to IDatabase interface
type mockUserService struct {
    t *testing.T
}

// NewMockUserService is new instance of NewMockUserService
func NewMockUserService(t *testing.T) *mockUserService{
    return &mockUserService{t}
}

// Create is mocked Create method to satisfy IUserStore interface
func (m *mockUserService) Create(input d.User) (*d.User, error) {
    if !input.IsValid() {
        return nil, E.New(E.ErrDataIsInvalid)
    }
    if !helper.EmailIsValid(input.Email) {
        return nil, E.New(E.ErrEmailIsInvalid)
    }
    if wantErr {
        return nil, E.New(E.ErrInsertDataFail)
    }
    return &input, nil
}

// Get is mocked get method to satisfy IUserStore interface
func (m *mockUserService) Get(id uuid.UUID) (*d.User, error) {
    if wantErr {
        return nil, E.New(E.ErrDataIsEmpty)
    }

    return u[0], nil
}

// Gets is mocked Gets method to satisfy IUserStore interface
func (m *mockUserService) Gets() ([]*d.User, error) {
    if wantErr {
        return nil, E.New(E.ErrDataIsEmpty)
    }

    return u, nil
}

// Update is mocked Update method to satisfy IUserStore interface
func (m *mockUserService) Update(id uuid.UUID, input d.User) (*d.User, error) {
    if !input.IsValid() {
        return nil, E.New(E.ErrDataIsInvalid)
    }

    if !helper.EmailIsValid(input.Email) {
        return nil, E.New(E.ErrEmailIsInvalid)
    }

    if wantErr {
        return nil, E.New(E.ErrUpdateDataFail)
    }

    return &input, nil
}

// Delete is mocked Delete method to satisfy IUserStore interface
func (m *mockUserService) Delete(id uuid.UUID) (*d.User, error) {
    if wantErr {
        return nil, E.New(E.ErrDeleteDataFail)
    }

    return u[0], nil
}

// GetByEmail is mocked GetCredential method to satisfy IUserStore interface
func (m *mockUserService) GetByEmail(email string) (*d.UserCredential, error) {
    if wantErr {
        return nil, E.New(E.ErrDataIsEmpty)
    }

    return &d.UserCredential{
        ID : u[0].ID,
        Username : u[0].Username,
        PassKey : u[0].PassKey,
        StatusID : u[0].StatusID,
    }, nil
}

// GetCredential is mocked GetCredential method to satisfy IUserStore interface
func (m *mockUserService) GetCredential(username,passkey string) (*d.UserCredential, error) {
    if wantErr {
        return nil, E.New(E.ErrDataIsEmpty)
    }

    return &d.UserCredential{
        ID : u[0].ID,
        Username : u[0].Username,
        PassKey : u[0].PassKey,
        StatusID : u[0].StatusID,
    }, nil
}

// IsUserExist is mocked IsUserExist method to satisfy IUserStore interface
func (m *mockUserService) IsUserExist(username,email string) (bool, error) {
    if wantErr {
        return false, E.New(E.ErrDatabase)
    }

    return true, nil
}

// TestParseUUID will test the helper function parseUUID
func TestParseUUID(t *testing.T) {
    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call
        uuid := ParseUUID(u[0].ID.String())

        assert.NotNil(t, uuid)
        assert.Equal(t, len(uuid.String()), 36)
    })

    // EXPECT FAIL cannot parse string to UUID. simulated by giving non uuid string input
    t.Run("EXPECT FAIL cannot parse string to uuid", func(t *testing.T){
        // actual method call
        uuid := ParseUUID("test")

        assert.Nil(t, uuid)
    })

}

// TestUserServiceCreate will test "Create" method for user service
func TestUserServiceCreate(t *testing.T) {
    // prepare mock and service
    mock := NewMockUserService(t)
    service := NewUserService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call
        got, err := service.Create(*convertToRequest(*u[0]))

        // test validation and verification
        assert.NoError(t, err)
        assert.NotNil(t, got)
        assert.Equal(t, u[0].Username, got.Username)
        assert.Equal(t, u[0].Firstname, got.Firstname)
        assert.Equal(t, u[0].Email, got.Email)
    })

    // EXPECT FAIL invalid data error. Simulated by removing some required data
    t.Run("EXPECT FAIL invalid data error", func (t *testing.T) {
        // prepare invalid data
        invalidUser := convertToRequest(*u[0])
        invalidUser.Email = ""
        invalidUser.Firstname = ""

        // actual method call
        got, err := service.Create(*invalidUser)

        // test validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL invalid email input. Simulation done by giving email with invalid format
    t.Run("EXPECT FAIL invalid email imput error", func (t *testing.T) {
        // prepare invalid data
        invalidUser := convertToRequest(*u[0])
        invalidUser.Email = "testmailerror.com"

        // actual method call
        got, err := service.Create(*invalidUser)

        // test validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL generate hash password error. Simulated by override helper.HashPassword function
    t.Run("EXPECT FAIL generate hash pass error", func (t *testing.T) {
        // prepare to override HashPassword
        generateHashPass := generateHashPassFunc
        generateHashPassFunc = func(password string) (string, error) {
            return "", E.New(E.ErrPasswordTooShort)
        }
        defer func() {
            generateHashPassFunc = generateHashPass
        }()

        // actual method call
        got, err := service.Create(*convertToRequest(*u[0]))

        // test validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)
        // t.Logf("ERROR: %v\n", err)
    })

    // EXPECT FAIL insert data error. Simulated by forcing Error result with wantErr set to true 
    t.Run("EXPECT FAIL insert data error", func (t *testing.T) {
        // actual method call (method to test)
        // trigger error from the mocked interface
        wantErr = true
        got, err := service.Create(*convertToRequest(*u[0]))
        wantErr = false

        // test validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)
    })
}

// TestUserServiceGet will test behaviour of Get method of user service layer
func TestUserServiceGet(t *testing.T) {
    // prepare mock and service
    mock := NewMockUserService(t)
    service := NewUserService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call
        got, err := service.Get(u[0].ID.String())

        // test validation and verification
        assert.NoError(t, err)
        assert.NotNil(t, got)
        assert.Equal(t, u[0].ConvertToResponse(), got)
    })

    // EXPECT FAIL data is empty error. Simulated by forcing to return error (set wantErr=true)
    t.Run("EXPECT FAIL data is empty error", func (t *testing.T) {
        // actual method call / test
        // set wantErr to true to force error return
        wantErr = true
        got, err := service.Get(u[0].ID.String())
        wantErr = false

        assert.Error(t, err)
        assert.Nil(t, got)
    })

}

// TestUserServiceGets will test behaviour of 'gets' method of 'user' service
func TestUserServiceGets(t *testing.T) {
    // prepare mock and service
    mock := NewMockUserService(t)
    service := NewUserService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call (method to test)
        got, err := service.Gets()

        wantUsers := make([]*d.UserResponse, 0)
        for _, user := range u {
            wantUsers = append(wantUsers, user.ConvertToResponse())
        }

        assert.NoError(t, err)
        assert.Equal(t, wantUsers, got)
    })

    // EXPECT FAIL data is empty error will simulated fail getting record
    // error return is ErrDataIsEmpty. Simulation done by triggering 'wantError' variable 
    // so the mocked interface return an error 
    t.Run("EXPECT FAIL database error", func (t *testing.T) {
        // actual method call (method to test)
        // trigger error from the mocked interface
        wantErr = true
        got, err := service.Gets()
        wantErr = false

        // test validation and verification
        assert.Error(t, err)
        assert.Nil(t, got)
    })
}

// TestUserServiceUpdate will test Update method behaviour of User service
func TestUserServiceUpdate(t *testing.T) {
    // prepare mock and service
    mock := NewMockUserService(t)
    service := NewUserService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call (method to test)
        got, err := service.Update(u[0].ID.String(), *convertToRequest(*u[0]))

        assert.NoError(t, err)
        assert.Equal(t, u[0].ID, got.ID)
        assert.Equal(t, u[0].Username, got.Username)
        assert.Equal(t, u[0].Firstname, got.Firstname)
        assert.Equal(t, u[0].Email, got.Email)
    })

    // EXPECT SUCCESS with new hashed pass generated. 
    // Simulated by override helper.HashPassword and helper.CheckPasswordHash
    t.Run("EXPECT SUCCESS new hashed password", func(t *testing.T){
        // mock checkPassHashFunc (instance func of helper.CheckPasswordHash)
        checkPassHash := checkPassHashFunc
        checkPassHashFunc = func(password string, hash string) bool {
            return true 
        }
        // mock generateHashPassFunc (instance func of helper.HashPassword)
        hashPass := generateHashPassFunc
        generateHashPassFunc = func(password string) (string,error) {
            return u[0].PassKey, nil
        }
        defer func() {
            generateHashPassFunc = hashPass
            checkPassHashFunc = checkPassHash
        }()

        // actual method call (method to test)
        got, err := service.Update(u[0].ID.String(), *convertToRequest(*u[0]))

        assert.NoError(t, err)
        assert.NotNil(t, got)
    })

    // EXPECT SUCCESS with old hashed pass.
    // Simulated by override helper.HashPassword and helper.CheckPasswordHash
    t.Run("EXPECT SUCCESS old hashed password", func(t *testing.T){
        // mock checkPassHashFunc (instance func of helper.CheckPasswordHash)
        checkPassHash := checkPassHashFunc
        checkPassHashFunc = func(password string, hash string) bool {
            return false
        }
        // mock generateHashPassFunc (instance func of helper.HashPassword)
        hashPass := generateHashPassFunc
        generateHashPassFunc = func(password string) (string,error) {
            return u[0].PassKey, nil
        }
        defer func() {
            generateHashPassFunc = hashPass
            checkPassHashFunc = checkPassHash
        }()

        // actual method call (method to test)
        got, err := service.Update(u[0].ID.String(), *convertToRequest(*u[0]))

        assert.NoError(t, err)
        assert.NotNil(t, got)
    })

    // EXPECT FAIL with hash pass generated fail.
    // Simulated by override helper.HashPassword and helper.CheckPasswordHash
    t.Run("EXPECT FAIL hashed password generate error", func(t *testing.T){
        // mock checkPassHashFunc (instance func of helper.CheckPasswordHash)
        checkPassHash := checkPassHashFunc
        checkPassHashFunc = func(password string, hash string) bool {
            return false
        }
        // mock generateHashPassFunc (instance func of helper.HashPassword)
        hashPass := generateHashPassFunc
        generateHashPassFunc = func(password string) (string,error) {
            return "", E.New(E.ErrPasswordTooShort)
        }
        defer func() {
            generateHashPassFunc = hashPass
            checkPassHashFunc = checkPassHash
        }()

        // actual method call (method to test)
        got, err := service.Update(u[0].ID.String(), *convertToRequest(*u[0]))

        assert.Error(t, err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL invalid data error. Simulated by giving invalid input data
    t.Run("EXPECT FAIL invalid data error", func(t *testing.T){
        // prepare invalid user data
        invalidUser := convertToRequest(*u[0])
        invalidUser.Firstname = ""

        // actual method call (method to test)
        got, err := service.Update(u[0].ID.String(), *invalidUser)

        assert.Error(t, err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL invalid email error. Simulated by giving invalid email data
    t.Run("EXPECT FAIL invalid email error", func(t *testing.T){
        // prepare invalid user data
        invalidUser := convertToRequest(*u[0])
        invalidUser.Email = "john.doe.com"

        // actual method call (method to test)
        got, err := service.Update(u[0].ID.String(), *invalidUser)

        assert.Error(t, err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL get user record for comparison. Simulated by forcing to return error
    // by setting wantErr=true
    t.Run("EXPECT FAIL get user record error", func(t *testing.T){
        // prepare invalid user data
        invalidUser := convertToRequest(*u[0])
        invalidUser.RoleID = 5

        // actual method call (method to test)
        wantErr = true
        got, err := service.Update(u[0].ID.String(), *convertToRequest(*u[0]))
        wantErr = false

        assert.Error(t, err)
        assert.Nil(t, got)
    })

    // EXPECT FAIL update record error. Simulated by forcing to return error
    // by setting wantErr=true
    t.Run("EXPECT FAIL update record error", func(t *testing.T){
        // prepare invalid user data
        invalidUser := convertToRequest(*u[0])

        // actual method call (method to test)
        wantErr = true
        got, err := service.Update(uuid.NewString(), *invalidUser)
        wantErr = false

        assert.Error(t, err)
        assert.Nil(t, got)
    })
}

// TestUserServiceDelete will test Delete method behaviour of User service
func TestUserServiceDelete(t *testing.T) {
    // prepare mock and service
    mock := NewMockUserService(t)
    service := NewUserService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call
        got, err := service.Delete(u[0].ID.String())

        // test verification and validation
        assert.NoError(t, err)
        assert.NotNil(t, got)

    })

    // EXPECT FAIL delete record error. Simulated by forcing to return error
    // by setting wantErr=true
    t.Run("EXPECT FAIL delete record error", func(t *testing.T){
        // actual method call (method to test)
        wantErr = true
        got, err := service.Delete(u[0].ID.String())
        wantErr = false

        assert.Error(t, err)
        assert.Nil(t, got)
    })
}

// TestUserServiceGetByEmail will test GetByEmail method behaviour of user service
func TestUserServiceGetByEmail(t *testing.T) {
    // prepare mock and service
    mock := NewMockUserService(t)
    service := NewUserService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call
        got, err := service.GetByEmail(u[0].Email)

        // test verification and validation
        assert.NoError(t, err)
        assert.NotNil(t, got)
        assert.Equal(t, u[0].ConvertToCredential(), got)
    })

    // EXPECT FAIL get record error. Simulated by forcing to return error
    // by setting wantErr=true
    t.Run("EXPECT FAIL get record error", func(t *testing.T){
        // actual method call (method to test)
        wantErr = true
        got, err := service.GetByEmail(u[0].Email)
        wantErr = false

        assert.Error(t, err)
        assert.Nil(t, got)
    })
}

// TestUserServiceGetCredential will test GetCredential method behaviour of user service
func TestUserServiceGetCredential(t *testing.T) {
    // prepare mock and service
    mock := NewMockUserService(t)
    service := NewUserService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call
        got, err := service.GetCredential(u[0].Username, u[0].PassKey)

        // test verification and validation
        assert.NoError(t, err)
        assert.NotNil(t, got)
        assert.Equal(t, u[0].ConvertToCredential(), got)
    })

    // EXPECT FAIL get record error. Simulated by forcing to return error
    // by setting wantErr=true
    t.Run("EXPECT FAIL get record error", func(t *testing.T){
        // actual method call (method to test)
        wantErr = true
        got, err := service.GetCredential(u[0].Username, u[0].PassKey)
        wantErr = false

        assert.Error(t, err)
        assert.Nil(t, got)
    })
}

// TestIsUserExist will test IsUserExist method behaviour of user service
func TestIsUserExist(t *testing.T) {
    // prepare mock and service
    mock := NewMockUserService(t)
    service := NewUserService(mock)

    // EXPECT SUCCESS will simulated normal operation with no error return
    // this simulation expect all process goes as expected
    t.Run("EXPECT SUCCESS", func(t *testing.T){
        // actual method call
        got := service.IsUserExist(u[0].Username,u[0].Email)

        // test verification and validation
        assert.Equal(t, true, got)
    })

    // EXPECT SUCCESS record not found. Simulated by forcing to return error
    // by setting wantErr=true so the result will be false
    t.Run("EXPECT SUCCESS data not found", func(t *testing.T){
        // actual method call (method to test)
        wantErr = true
        got := service.IsUserExist(u[0].Username,u[0].Email)
        wantErr = false

        assert.Equal(t, false, got)
    })

    // EXPECT FAIL email invalid. Simulated by inserting invalid mail 
    t.Run("EXPECT SUCCESS data not found", func(t *testing.T){
        // actual method call (method to test)
        got := service.IsUserExist(u[0].Username, "aaa.com")

        assert.Equal(t, false, got)
    })
}

// convertToRequest is test helper function to convert from user to request dto
func convertToRequest(u d.User) *d.UserRequest{
    return &d.UserRequest{
        ID        : u.ID,
        Username  : u.Username,
        Firstname : u.Firstname,
        Lastname  : u.Lastname,
        Email     : u.Email,
        PassKey   : u.PassKey,
        StatusID  : u.StatusID,
        RoleID    : u.RoleID,
    }
}
