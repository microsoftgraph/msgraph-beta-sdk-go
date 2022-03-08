package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Commandable 
type Commandable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppServiceName()(*string)
    GetError()(*string)
    GetPackageFamilyName()(*string)
    GetPayload()(PayloadRequestable)
    GetPermissionTicket()(*string)
    GetPostBackUri()(*string)
    GetResponsepayload()(PayloadResponseable)
    GetStatus()(*string)
    GetType()(*string)
    SetAppServiceName(value *string)()
    SetError(value *string)()
    SetPackageFamilyName(value *string)()
    SetPayload(value PayloadRequestable)()
    SetPermissionTicket(value *string)()
    SetPostBackUri(value *string)()
    SetResponsepayload(value PayloadResponseable)()
    SetStatus(value *string)()
    SetType(value *string)()
}
