package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AllPreApprovedPermissionsOnResourceApp struct {
    PreApprovedPermissions
}
// NewAllPreApprovedPermissionsOnResourceApp instantiates a new AllPreApprovedPermissionsOnResourceApp and sets the default values.
func NewAllPreApprovedPermissionsOnResourceApp()(*AllPreApprovedPermissionsOnResourceApp) {
    m := &AllPreApprovedPermissionsOnResourceApp{
        PreApprovedPermissions: *NewPreApprovedPermissions(),
    }
    odataTypeValue := "#microsoft.graph.allPreApprovedPermissionsOnResourceApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAllPreApprovedPermissionsOnResourceAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAllPreApprovedPermissionsOnResourceAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllPreApprovedPermissionsOnResourceApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AllPreApprovedPermissionsOnResourceApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PreApprovedPermissions.GetFieldDeserializers()
    res["resourceApplicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceApplicationId(val)
        }
        return nil
    }
    return res
}
// GetResourceApplicationId gets the resourceApplicationId property value. The appId of the resource application (the API). Required.
// returns a *string when successful
func (m *AllPreApprovedPermissionsOnResourceApp) GetResourceApplicationId()(*string) {
    val, err := m.GetBackingStore().Get("resourceApplicationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AllPreApprovedPermissionsOnResourceApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PreApprovedPermissions.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("resourceApplicationId", m.GetResourceApplicationId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResourceApplicationId sets the resourceApplicationId property value. The appId of the resource application (the API). Required.
func (m *AllPreApprovedPermissionsOnResourceApp) SetResourceApplicationId(value *string)() {
    err := m.GetBackingStore().Set("resourceApplicationId", value)
    if err != nil {
        panic(err)
    }
}
type AllPreApprovedPermissionsOnResourceAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PreApprovedPermissionsable
    GetResourceApplicationId()(*string)
    SetResourceApplicationId(value *string)()
}
