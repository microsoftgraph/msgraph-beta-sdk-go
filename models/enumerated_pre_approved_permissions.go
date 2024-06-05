package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EnumeratedPreApprovedPermissions struct {
    PreApprovedPermissions
}
// NewEnumeratedPreApprovedPermissions instantiates a new EnumeratedPreApprovedPermissions and sets the default values.
func NewEnumeratedPreApprovedPermissions()(*EnumeratedPreApprovedPermissions) {
    m := &EnumeratedPreApprovedPermissions{
        PreApprovedPermissions: *NewPreApprovedPermissions(),
    }
    odataTypeValue := "#microsoft.graph.enumeratedPreApprovedPermissions"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEnumeratedPreApprovedPermissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnumeratedPreApprovedPermissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnumeratedPreApprovedPermissions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnumeratedPreApprovedPermissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PreApprovedPermissions.GetFieldDeserializers()
    res["permissionIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPermissionIds(res)
        }
        return nil
    }
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
// GetPermissionIds gets the permissionIds property value. The list of id values for the specific resource-specific application permissions to match with. See the resourceSpecificApplicationPermissions property of the servicePrincipal object for the list of permissions.
// returns a []string when successful
func (m *EnumeratedPreApprovedPermissions) GetPermissionIds()([]string) {
    val, err := m.GetBackingStore().Get("permissionIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetResourceApplicationId gets the resourceApplicationId property value. The appId of the resource application (the API). Required.
// returns a *string when successful
func (m *EnumeratedPreApprovedPermissions) GetResourceApplicationId()(*string) {
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
func (m *EnumeratedPreApprovedPermissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PreApprovedPermissions.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPermissionIds() != nil {
        err = writer.WriteCollectionOfStringValues("permissionIds", m.GetPermissionIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceApplicationId", m.GetResourceApplicationId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPermissionIds sets the permissionIds property value. The list of id values for the specific resource-specific application permissions to match with. See the resourceSpecificApplicationPermissions property of the servicePrincipal object for the list of permissions.
func (m *EnumeratedPreApprovedPermissions) SetPermissionIds(value []string)() {
    err := m.GetBackingStore().Set("permissionIds", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceApplicationId sets the resourceApplicationId property value. The appId of the resource application (the API). Required.
func (m *EnumeratedPreApprovedPermissions) SetResourceApplicationId(value *string)() {
    err := m.GetBackingStore().Set("resourceApplicationId", value)
    if err != nil {
        panic(err)
    }
}
type EnumeratedPreApprovedPermissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PreApprovedPermissionsable
    GetPermissionIds()([]string)
    GetResourceApplicationId()(*string)
    SetPermissionIds(value []string)()
    SetResourceApplicationId(value *string)()
}
