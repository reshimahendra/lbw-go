/*
   Package errors for database
   containing custom error for our app
*/
package errors

const (
    // ErrDatabase is error code for database error
    // msg = "database error" 
    ErrDatabase = iota + 800 

    // ErrDatabaseConfiguration is error code for unproper database configuration
    // msg = "database configuration error"
    ErrDatabaseConfiguration

    // ErrDatabaseTransactionNil is error code for nil database transaction
    // msg = "database transaction nil" 
    ErrDatabaseTransactionNil

    // ErrDatabaseRollback is error code for database rolback fail
    // msg = "database rollback fail" 
    ErrDatabaseRollback

    // ErrDatabasePoolNil is error code for nil database pool
    // msg = "database pool is nil" 
    ErrDatabasePoolNil

    // ErrDataEmpty is error code for empty data result
    // msg = "data is empty"
    ErrDataIsEmpty

    // ErrDataIsInvalid is error code for invalid data
    // msg = "data is invalid"
    ErrDataIsInvalid

    // ErrDataNotFound is error code for not found data 
    // msg = "data not found"
    ErrDataNotFound

    // ErrGettingData is error code for fail to get/ retreive data
    // msg = "could not retreive data"
    ErrGettingData

    // ErrInsertDataFail is error code for 'failling' on inserting data
    // msg = "could not insert data"
    ErrInsertDataFail

    // ErrDataCouldNotUpdate is error code for 'failling' on update data
    // msg = "could not update data"
    ErrUpdateDataFail

    // ErrDeleteData is error code for failing to delete data
    // msg = "could not delete data"
    ErrDeleteDataFail

    // ErrDataAlreadyExist is error code when triying to save data on an already exist data
    // for example 'Primary Key' or 'Unique Constraint' already exist 
    // msg = "data already exist"
    ErrDataAlreadyExist

)

const (
    // ErrDatabaseMsg is error message for database error
    // msg = "database error" 
    ErrDatabaseMsg = "database error" 

    // ErrDatabaseConfigurationMsg is error message for unproper database configuration
    // msg = "database configuration error"
    ErrDatabaseConfigurationMsg = "database configuration error"

    // ErrDatabaseTransactionNilMsg is error message for database transaction nil
    // msg = "database transaction nil" 
    ErrDatabaseTransactionNilMsg = "database transaction nil" 

    // ErrDatabaseRollbackMsg is error message for database fail to rollback
    // msg = "database rollback fail" 
    ErrDatabaseRollbackMsg = "database rollback fail" 

    // ErrDatabasePoolNil is error message for nil database pool
    // msg = "database pool is nil" 
    ErrDatabasePoolNilMsg = "database pool is nil" 

    // ErrDataEmptyMsg is error code for empty data result
    // msg = "data is empty"
    ErrDataIsEmptyMsg = "data is empty"

    // ErrDataIsInvalidMsg is error code for invalid data
    // msg = "data is invalid"
    ErrDataIsInvalidMsg = "data is invalid"

    // ErrDataNotFoundMsg is error code for not found data 
    // msg = "data not found"
    ErrDataNotFoundMsg = "data not found"

    // ErrGettingDataMsg is error message for fail to get/ retreive data
    // msg = "could not retreive data"
    ErrGettingDataMsg = "could not retreive data"

    // ErrInsertDataFailMsg is error code for 'failling' on saving data
    // msg = "could not insert data"
    ErrInsertDataFailMsg = "could not insert data"

    // ErrDataCouldNotUpdateMsg is error code for 'failling' on update data
    // msg = "could not update data"
    ErrUpdateDataFailMsg = "could not update data"

    // ErrDeleteDataMsg is error code for failing to delete data
    // msg = "could not delete data"
    ErrDeleteDataFailMsg = "could not delete data"

    // ErrDataExist is error code when triying to save data on an already exist data
    // msg = "data already exist"
    ErrDataAlreadyExistMsg = "data already exist"
)
