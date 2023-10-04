package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GcpPermissionsDefinitionAction 
type GcpPermissionsDefinitionAction struct {
    PermissionsDefinitionAction
}
// NewGcpPermissionsDefinitionAction instantiates a new gcpPermissionsDefinitionAction and sets the default values.
func NewGcpPermissionsDefinitionAction()(*GcpPermissionsDefinitionAction) {
    m := &GcpPermissionsDefinitionAction{
        PermissionsDefinitionAction: *NewPermissionsDefinitionAction(),
    }
    odataTypeValue := "#microsoft.graph.gcpPermissionsDefinitionAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGcpPermissionsDefinitionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGcpPermissionsDefinitionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.gcpActionPermissionsDefinitionAction":
                        return NewGcpActionPermissionsDefinitionAction(), nil
                    case "#microsoft.graph.gcpRolePermissionsDefinitionAction":
                        return NewGcpRolePermissionsDefinitionAction(), nil
                }
            }
        }
    }
    return NewGcpPermissionsDefinitionAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GcpPermissionsDefinitionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PermissionsDefinitionAction.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *GcpPermissionsDefinitionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PermissionsDefinitionAction.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// GcpPermissionsDefinitionActionable 
type GcpPermissionsDefinitionActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PermissionsDefinitionActionable
}
