package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityConfigurationTaskCollectionResponse 
type SecurityConfigurationTaskCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewSecurityConfigurationTaskCollectionResponse instantiates a new SecurityConfigurationTaskCollectionResponse and sets the default values.
func NewSecurityConfigurationTaskCollectionResponse()(*SecurityConfigurationTaskCollectionResponse) {
    m := &SecurityConfigurationTaskCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSecurityConfigurationTaskCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityConfigurationTaskCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityConfigurationTaskCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityConfigurationTaskCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityConfigurationTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityConfigurationTaskable, len(val))
            for i, v := range val {
                res[i] = v.(SecurityConfigurationTaskable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SecurityConfigurationTaskCollectionResponse) GetValue()([]SecurityConfigurationTaskable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityConfigurationTaskable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SecurityConfigurationTaskCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SecurityConfigurationTaskCollectionResponse) SetValue(value []SecurityConfigurationTaskable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// SecurityConfigurationTaskCollectionResponseable 
type SecurityConfigurationTaskCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]SecurityConfigurationTaskable)
    SetValue(value []SecurityConfigurationTaskable)()
}
