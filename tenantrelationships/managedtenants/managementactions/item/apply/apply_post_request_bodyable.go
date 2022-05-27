package apply

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplyPostRequestBodyable 
type ApplyPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludeGroups()([]string)
    GetIncludeAllUsers()(*bool)
    GetIncludeGroups()([]string)
    GetManagementTemplateId()(*string)
    GetTenantGroupId()(*string)
    GetTenantId()(*string)
    SetExcludeGroups(value []string)()
    SetIncludeAllUsers(value *bool)()
    SetIncludeGroups(value []string)()
    SetManagementTemplateId(value *string)()
    SetTenantGroupId(value *string)()
    SetTenantId(value *string)()
}
