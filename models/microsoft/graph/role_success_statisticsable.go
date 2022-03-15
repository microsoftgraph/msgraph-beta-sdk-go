package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleSuccessStatisticsable 
type RoleSuccessStatisticsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetPermanentFail()(*int32)
    GetPermanentSuccess()(*int32)
    GetRemoveFail()(*int32)
    GetRemoveSuccess()(*int32)
    GetRoleId()(*string)
    GetRoleName()(*string)
    GetTemporaryFail()(*int32)
    GetTemporarySuccess()(*int32)
    GetUnknownFail()(*int32)
    SetPermanentFail(value *int32)()
    SetPermanentSuccess(value *int32)()
    SetRemoveFail(value *int32)()
    SetRemoveSuccess(value *int32)()
    SetRoleId(value *string)()
    SetRoleName(value *string)()
    SetTemporaryFail(value *int32)()
    SetTemporarySuccess(value *int32)()
    SetUnknownFail(value *int32)()
}
