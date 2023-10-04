package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureRolePermissionsDefinitionAction 
type AzureRolePermissionsDefinitionAction struct {
    AzurePermissionsDefinitionAction
}
// NewAzureRolePermissionsDefinitionAction instantiates a new azureRolePermissionsDefinitionAction and sets the default values.
func NewAzureRolePermissionsDefinitionAction()(*AzureRolePermissionsDefinitionAction) {
    m := &AzureRolePermissionsDefinitionAction{
        AzurePermissionsDefinitionAction: *NewAzurePermissionsDefinitionAction(),
    }
    odataTypeValue := "#microsoft.graph.azureRolePermissionsDefinitionAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureRolePermissionsDefinitionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureRolePermissionsDefinitionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureRolePermissionsDefinitionAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureRolePermissionsDefinitionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AzurePermissionsDefinitionAction.GetFieldDeserializers()
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionsDefinitionAzureRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionsDefinitionAzureRoleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PermissionsDefinitionAzureRoleable)
                }
            }
            m.SetRoles(res)
        }
        return nil
    }
    return res
}
// GetRoles gets the roles property value. The roles property
func (m *AzureRolePermissionsDefinitionAction) GetRoles()([]PermissionsDefinitionAzureRoleable) {
    val, err := m.GetBackingStore().Get("roles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PermissionsDefinitionAzureRoleable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AzureRolePermissionsDefinitionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AzurePermissionsDefinitionAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoles()))
        for i, v := range m.GetRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roles", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRoles sets the roles property value. The roles property
func (m *AzureRolePermissionsDefinitionAction) SetRoles(value []PermissionsDefinitionAzureRoleable)() {
    err := m.GetBackingStore().Set("roles", value)
    if err != nil {
        panic(err)
    }
}
// AzureRolePermissionsDefinitionActionable 
type AzureRolePermissionsDefinitionActionable interface {
    AzurePermissionsDefinitionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRoles()([]PermissionsDefinitionAzureRoleable)
    SetRoles(value []PermissionsDefinitionAzureRoleable)()
}
