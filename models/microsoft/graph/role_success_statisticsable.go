package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleSuccessStatisticsable 
type RoleSuccessStatisticsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetPermanentFail()(*int64)
    GetPermanentSuccess()(*int64)
    GetRemoveFail()(*int64)
    GetRemoveSuccess()(*int64)
    GetRoleId()(*string)
    GetRoleName()(*string)
    GetTemporaryFail()(*int64)
    GetTemporarySuccess()(*int64)
    GetUnknownFail()(*int64)
    SetPermanentFail(value *int64)()
    SetPermanentSuccess(value *int64)()
    SetRemoveFail(value *int64)()
    SetRemoveSuccess(value *int64)()
    SetRoleId(value *string)()
    SetRoleName(value *string)()
    SetTemporaryFail(value *int64)()
    SetTemporarySuccess(value *int64)()
    SetUnknownFail(value *int64)()
}
