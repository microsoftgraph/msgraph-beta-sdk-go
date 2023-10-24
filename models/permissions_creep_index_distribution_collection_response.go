package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PermissionsCreepIndexDistributionCollectionResponse 
type PermissionsCreepIndexDistributionCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewPermissionsCreepIndexDistributionCollectionResponse instantiates a new permissionsCreepIndexDistributionCollectionResponse and sets the default values.
func NewPermissionsCreepIndexDistributionCollectionResponse()(*PermissionsCreepIndexDistributionCollectionResponse) {
    m := &PermissionsCreepIndexDistributionCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePermissionsCreepIndexDistributionCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePermissionsCreepIndexDistributionCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermissionsCreepIndexDistributionCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PermissionsCreepIndexDistributionCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionsCreepIndexDistributionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionsCreepIndexDistributionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PermissionsCreepIndexDistributionable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *PermissionsCreepIndexDistributionCollectionResponse) GetValue()([]PermissionsCreepIndexDistributionable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PermissionsCreepIndexDistributionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PermissionsCreepIndexDistributionCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *PermissionsCreepIndexDistributionCollectionResponse) SetValue(value []PermissionsCreepIndexDistributionable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// PermissionsCreepIndexDistributionCollectionResponseable 
type PermissionsCreepIndexDistributionCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]PermissionsCreepIndexDistributionable)
    SetValue(value []PermissionsCreepIndexDistributionable)()
}
