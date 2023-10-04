package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzurePermissionsDefinitionAction 
type AzurePermissionsDefinitionAction struct {
    PermissionsDefinitionAction
}
// NewAzurePermissionsDefinitionAction instantiates a new azurePermissionsDefinitionAction and sets the default values.
func NewAzurePermissionsDefinitionAction()(*AzurePermissionsDefinitionAction) {
    m := &AzurePermissionsDefinitionAction{
        PermissionsDefinitionAction: *NewPermissionsDefinitionAction(),
    }
    odataTypeValue := "#microsoft.graph.azurePermissionsDefinitionAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzurePermissionsDefinitionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzurePermissionsDefinitionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.azureActionPermissionsDefinitionAction":
                        return NewAzureActionPermissionsDefinitionAction(), nil
                    case "#microsoft.graph.azureRolePermissionsDefinitionAction":
                        return NewAzureRolePermissionsDefinitionAction(), nil
                }
            }
        }
    }
    return NewAzurePermissionsDefinitionAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzurePermissionsDefinitionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PermissionsDefinitionAction.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AzurePermissionsDefinitionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PermissionsDefinitionAction.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// AzurePermissionsDefinitionActionable 
type AzurePermissionsDefinitionActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PermissionsDefinitionActionable
}
