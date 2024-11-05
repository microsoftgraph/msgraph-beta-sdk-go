package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceAndAppManagementAssignedRoleDefinition complex type to represent the role that is assigned to the user. This type contains the ID, the display name, and the permissions of the role.
type DeviceAndAppManagementAssignedRoleDefinition struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceAndAppManagementAssignedRoleDefinition instantiates a new DeviceAndAppManagementAssignedRoleDefinition and sets the default values.
func NewDeviceAndAppManagementAssignedRoleDefinition()(*DeviceAndAppManagementAssignedRoleDefinition) {
    m := &DeviceAndAppManagementAssignedRoleDefinition{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceAndAppManagementAssignedRoleDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceAndAppManagementAssignedRoleDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAndAppManagementAssignedRoleDefinition(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeviceAndAppManagementAssignedRoleDefinition) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DeviceAndAppManagementAssignedRoleDefinition) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceAndAppManagementAssignedRoleDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPermissions(res)
        }
        return nil
    }
    res["roleDefinitionDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionDisplayName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DeviceAndAppManagementAssignedRoleDefinition) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPermissions gets the permissions property value. A list of permissions based on its associated role. Each permission defines the specific actions the user can perform on Intune resources, such as managing devices, applications, or configurations. Some possible values are: Microsoft.Intune/MobileApps/Read, Microsoft.Intune/DeviceConfigurations/Write, Microsoft.Intune/ManagedDevices/Retire, and Microsoft.Intune/DeviceCompliancePolicies/Assign. This Permissions property offers a comprehensive view of the user's effective access rights, ensuring that they can only perform actions relevant to their assigned roles. This property is read-only.
// returns a []string when successful
func (m *DeviceAndAppManagementAssignedRoleDefinition) GetPermissions()([]string) {
    val, err := m.GetBackingStore().Get("permissions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRoleDefinitionDisplayName gets the roleDefinitionDisplayName property value. The RoleDefinitionDisplayName property represents the human-readable name of a specific role definition in Microsoft Intune. This property provides a clear and descriptive name that indicates the purpose or scope of the role, helping administrators identify and assign appropriate roles to users or groups.Some example values for RoleDefinitionDisplayName might include: "Helpdesk Operator," "Application Manager," or "Policy Administrator." This display name is primarily used in the Intune console or Graph API to present roles in a user-friendly manner, making it easier for administrators to manage role-based access control (RBAC) efficiently. This property is read-only.
// returns a *string when successful
func (m *DeviceAndAppManagementAssignedRoleDefinition) GetRoleDefinitionDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("roleDefinitionDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementAssignedRoleDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAndAppManagementAssignedRoleDefinition) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceAndAppManagementAssignedRoleDefinition) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceAndAppManagementAssignedRoleDefinition) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissions sets the permissions property value. A list of permissions based on its associated role. Each permission defines the specific actions the user can perform on Intune resources, such as managing devices, applications, or configurations. Some possible values are: Microsoft.Intune/MobileApps/Read, Microsoft.Intune/DeviceConfigurations/Write, Microsoft.Intune/ManagedDevices/Retire, and Microsoft.Intune/DeviceCompliancePolicies/Assign. This Permissions property offers a comprehensive view of the user's effective access rights, ensuring that they can only perform actions relevant to their assigned roles. This property is read-only.
func (m *DeviceAndAppManagementAssignedRoleDefinition) SetPermissions(value []string)() {
    err := m.GetBackingStore().Set("permissions", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitionDisplayName sets the roleDefinitionDisplayName property value. The RoleDefinitionDisplayName property represents the human-readable name of a specific role definition in Microsoft Intune. This property provides a clear and descriptive name that indicates the purpose or scope of the role, helping administrators identify and assign appropriate roles to users or groups.Some example values for RoleDefinitionDisplayName might include: "Helpdesk Operator," "Application Manager," or "Policy Administrator." This display name is primarily used in the Intune console or Graph API to present roles in a user-friendly manner, making it easier for administrators to manage role-based access control (RBAC) efficiently. This property is read-only.
func (m *DeviceAndAppManagementAssignedRoleDefinition) SetRoleDefinitionDisplayName(value *string)() {
    err := m.GetBackingStore().Set("roleDefinitionDisplayName", value)
    if err != nil {
        panic(err)
    }
}
type DeviceAndAppManagementAssignedRoleDefinitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetPermissions()([]string)
    GetRoleDefinitionDisplayName()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetPermissions(value []string)()
    SetRoleDefinitionDisplayName(value *string)()
}
