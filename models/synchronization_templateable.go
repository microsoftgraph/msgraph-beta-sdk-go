package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationTemplateable 
type SynchronizationTemplateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationId()(*string)
    GetDefault()(*bool)
    GetDescription()(*string)
    GetDiscoverable()(*bool)
    GetFactoryTag()(*string)
    GetMetadata()([]MetadataEntryable)
    GetSchema()(SynchronizationSchemaable)
    SetApplicationId(value *string)()
    SetDefault(value *bool)()
    SetDescription(value *string)()
    SetDiscoverable(value *bool)()
    SetFactoryTag(value *string)()
    SetMetadata(value []MetadataEntryable)()
    SetSchema(value SynchronizationSchemaable)()
}
