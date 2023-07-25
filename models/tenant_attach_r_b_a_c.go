package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantAttachRBAC singleton entity that acts as a container for tenant attach enablement functionality.
type TenantAttachRBAC struct {
    Entity
}
// NewTenantAttachRBAC instantiates a new tenantAttachRBAC and sets the default values.
func NewTenantAttachRBAC()(*TenantAttachRBAC) {
    m := &TenantAttachRBAC{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTenantAttachRBACFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantAttachRBACFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantAttachRBAC(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantAttachRBAC) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TenantAttachRBAC) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// TenantAttachRBACable 
type TenantAttachRBACable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
