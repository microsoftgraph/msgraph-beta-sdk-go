package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceAndAppManagementAssignedRoleDetail the DeviceAndAppManagementAssignedRoleDetail is a complex type in Microsoft Intune used to represent the Role Definitions and Permissions that are assigned to a specific user. This type provides a detailed view of the roles a user holds, along with the associated permissions that determine the specific actions the user can perform within Intune environment.
type DeviceAndAppManagementAssignedRoleDetail struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceAndAppManagementAssignedRoleDetail instantiates a new DeviceAndAppManagementAssignedRoleDetail and sets the default values.
func NewDeviceAndAppManagementAssignedRoleDetail()(*DeviceAndAppManagementAssignedRoleDetail) {
    m := &DeviceAndAppManagementAssignedRoleDetail{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceAndAppManagementAssignedRoleDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceAndAppManagementAssignedRoleDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAndAppManagementAssignedRoleDetail(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeviceAndAppManagementAssignedRoleDetail) GetAdditionalData()(map[string]any) {
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
func (m *DeviceAndAppManagementAssignedRoleDetail) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceAndAppManagementAssignedRoleDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["roleDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceAndAppManagementAssignedRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAndAppManagementAssignedRoleDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceAndAppManagementAssignedRoleDefinitionable)
                }
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DeviceAndAppManagementAssignedRoleDetail) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPermissions gets the permissions property value. The list of permissions assigned to a specific user based on their associated role definitions. Each permission defines the specific actions the user can perform on Intune resources, such as managing devices, applications, or configurations. Some possible values are: Microsoft.Intune/MobileApps/Read, Microsoft.Intune/DeviceConfigurations/Write, Microsoft.Intune/ManagedDevices/Retire, and Microsoft.Intune/DeviceCompliancePolicies/Assign. This Permissions property provides a comprehensive view of the user's effective access rights, ensuring that they can only perform actions relevant to their assigned roles. This property is read-only.
// returns a []string when successful
func (m *DeviceAndAppManagementAssignedRoleDetail) GetPermissions()([]string) {
    val, err := m.GetBackingStore().Get("permissions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRoleDefinitions gets the roleDefinitions property value. A collection of RoleDefinitions represents the various administrative roles that define permissions and access levels within Microsoft Intune. Each RoleDefinition outlines a set of permissions that determine the actions an admin or user can perform in the Intune environment. These permissions can include actions like reading or writing to specific resources, managing device configurations, deploying policies, or handling user data. RoleDefinitions are critical for enforcing role-based access control (RBAC), ensuring that administrators can only interact with the features and data relevant to their responsibilities. RoleDefinitions in Intune can either be built-in roles provided by Microsoft or custom roles created by an organization to tailor access based on specific needs. These definitions are referenced when assigning roles to users or groups, effectively controlling the scope of their administrative privileges. The collection of RoleDefinitions is managed through the Intune console or the Graph API, allowing for scalable role management across large environments. This property is read-only.
// returns a []DeviceAndAppManagementAssignedRoleDefinitionable when successful
func (m *DeviceAndAppManagementAssignedRoleDetail) GetRoleDefinitions()([]DeviceAndAppManagementAssignedRoleDefinitionable) {
    val, err := m.GetBackingStore().Get("roleDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceAndAppManagementAssignedRoleDefinitionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementAssignedRoleDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceAndAppManagementAssignedRoleDetail) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceAndAppManagementAssignedRoleDetail) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceAndAppManagementAssignedRoleDetail) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissions sets the permissions property value. The list of permissions assigned to a specific user based on their associated role definitions. Each permission defines the specific actions the user can perform on Intune resources, such as managing devices, applications, or configurations. Some possible values are: Microsoft.Intune/MobileApps/Read, Microsoft.Intune/DeviceConfigurations/Write, Microsoft.Intune/ManagedDevices/Retire, and Microsoft.Intune/DeviceCompliancePolicies/Assign. This Permissions property provides a comprehensive view of the user's effective access rights, ensuring that they can only perform actions relevant to their assigned roles. This property is read-only.
func (m *DeviceAndAppManagementAssignedRoleDetail) SetPermissions(value []string)() {
    err := m.GetBackingStore().Set("permissions", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. A collection of RoleDefinitions represents the various administrative roles that define permissions and access levels within Microsoft Intune. Each RoleDefinition outlines a set of permissions that determine the actions an admin or user can perform in the Intune environment. These permissions can include actions like reading or writing to specific resources, managing device configurations, deploying policies, or handling user data. RoleDefinitions are critical for enforcing role-based access control (RBAC), ensuring that administrators can only interact with the features and data relevant to their responsibilities. RoleDefinitions in Intune can either be built-in roles provided by Microsoft or custom roles created by an organization to tailor access based on specific needs. These definitions are referenced when assigning roles to users or groups, effectively controlling the scope of their administrative privileges. The collection of RoleDefinitions is managed through the Intune console or the Graph API, allowing for scalable role management across large environments. This property is read-only.
func (m *DeviceAndAppManagementAssignedRoleDetail) SetRoleDefinitions(value []DeviceAndAppManagementAssignedRoleDefinitionable)() {
    err := m.GetBackingStore().Set("roleDefinitions", value)
    if err != nil {
        panic(err)
    }
}
type DeviceAndAppManagementAssignedRoleDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetPermissions()([]string)
    GetRoleDefinitions()([]DeviceAndAppManagementAssignedRoleDefinitionable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetPermissions(value []string)()
    SetRoleDefinitions(value []DeviceAndAppManagementAssignedRoleDefinitionable)()
}
