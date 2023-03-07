package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineDeviceStateCollectionResponse 
type SecurityBaselineDeviceStateCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewSecurityBaselineDeviceStateCollectionResponse instantiates a new SecurityBaselineDeviceStateCollectionResponse and sets the default values.
func NewSecurityBaselineDeviceStateCollectionResponse()(*SecurityBaselineDeviceStateCollectionResponse) {
    m := &SecurityBaselineDeviceStateCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSecurityBaselineDeviceStateCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityBaselineDeviceStateCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityBaselineDeviceStateCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityBaselineDeviceStateCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityBaselineDeviceStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineDeviceStateable, len(val))
            for i, v := range val {
                res[i] = v.(SecurityBaselineDeviceStateable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SecurityBaselineDeviceStateCollectionResponse) GetValue()([]SecurityBaselineDeviceStateable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityBaselineDeviceStateable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SecurityBaselineDeviceStateCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *SecurityBaselineDeviceStateCollectionResponse) SetValue(value []SecurityBaselineDeviceStateable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// SecurityBaselineDeviceStateCollectionResponseable 
type SecurityBaselineDeviceStateCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]SecurityBaselineDeviceStateable)
    SetValue(value []SecurityBaselineDeviceStateable)()
}
