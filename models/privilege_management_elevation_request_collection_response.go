package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegeManagementElevationRequestCollectionResponse 
type PrivilegeManagementElevationRequestCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewPrivilegeManagementElevationRequestCollectionResponse instantiates a new privilegeManagementElevationRequestCollectionResponse and sets the default values.
func NewPrivilegeManagementElevationRequestCollectionResponse()(*PrivilegeManagementElevationRequestCollectionResponse) {
    m := &PrivilegeManagementElevationRequestCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePrivilegeManagementElevationRequestCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegeManagementElevationRequestCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegeManagementElevationRequestCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegeManagementElevationRequestCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegeManagementElevationRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegeManagementElevationRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrivilegeManagementElevationRequestable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *PrivilegeManagementElevationRequestCollectionResponse) GetValue()([]PrivilegeManagementElevationRequestable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrivilegeManagementElevationRequestable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivilegeManagementElevationRequestCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PrivilegeManagementElevationRequestCollectionResponse) SetValue(value []PrivilegeManagementElevationRequestable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// PrivilegeManagementElevationRequestCollectionResponseable 
type PrivilegeManagementElevationRequestCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]PrivilegeManagementElevationRequestable)
    SetValue(value []PrivilegeManagementElevationRequestable)()
}
