package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedEmailDomainCollectionResponse provides operations to manage the sharedEmailDomains property of the microsoft.graph.directory entity.
type SharedEmailDomainCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []SharedEmailDomainable
}
// NewSharedEmailDomainCollectionResponse instantiates a new SharedEmailDomainCollectionResponse and sets the default values.
func NewSharedEmailDomainCollectionResponse()(*SharedEmailDomainCollectionResponse) {
    m := &SharedEmailDomainCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSharedEmailDomainCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedEmailDomainCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharedEmailDomainCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedEmailDomainCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSharedEmailDomainFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *SharedEmailDomainCollectionResponse) GetValue()([]SharedEmailDomainable) {
    return m.value
}
// Serialize serializes information the current object
func (m *SharedEmailDomainCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *SharedEmailDomainCollectionResponse) SetValue(value []SharedEmailDomainable)() {
    m.value = value
}
