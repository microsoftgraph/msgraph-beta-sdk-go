package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantRelationshipAccessPolicyBaseable 
type TenantRelationshipAccessPolicyBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PolicyBaseable
    GetDefinition()([]string)
    GetType()(*string)
    SetDefinition(value []string)()
    SetType(value *string)()
}
