package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AwsActionsPermissionsDefinitionAction 
type AwsActionsPermissionsDefinitionAction struct {
    AwsPermissionsDefinitionAction
}
// NewAwsActionsPermissionsDefinitionAction instantiates a new awsActionsPermissionsDefinitionAction and sets the default values.
func NewAwsActionsPermissionsDefinitionAction()(*AwsActionsPermissionsDefinitionAction) {
    m := &AwsActionsPermissionsDefinitionAction{
        AwsPermissionsDefinitionAction: *NewAwsPermissionsDefinitionAction(),
    }
    odataTypeValue := "#microsoft.graph.awsActionsPermissionsDefinitionAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsActionsPermissionsDefinitionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsActionsPermissionsDefinitionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsActionsPermissionsDefinitionAction(), nil
}
// GetAssignToRoleId gets the assignToRoleId property value. The assignToRoleId property
func (m *AwsActionsPermissionsDefinitionAction) GetAssignToRoleId()(*string) {
    val, err := m.GetBackingStore().Get("assignToRoleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AwsActionsPermissionsDefinitionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AwsPermissionsDefinitionAction.GetFieldDeserializers()
    res["assignToRoleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignToRoleId(val)
        }
        return nil
    }
    res["statements"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAwsStatementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AwsStatementable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AwsStatementable)
                }
            }
            m.SetStatements(res)
        }
        return nil
    }
    return res
}
// GetStatements gets the statements property value. The statements property
func (m *AwsActionsPermissionsDefinitionAction) GetStatements()([]AwsStatementable) {
    val, err := m.GetBackingStore().Get("statements")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AwsStatementable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsActionsPermissionsDefinitionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AwsPermissionsDefinitionAction.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignToRoleId", m.GetAssignToRoleId())
        if err != nil {
            return err
        }
    }
    if m.GetStatements() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStatements()))
        for i, v := range m.GetStatements() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("statements", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignToRoleId sets the assignToRoleId property value. The assignToRoleId property
func (m *AwsActionsPermissionsDefinitionAction) SetAssignToRoleId(value *string)() {
    err := m.GetBackingStore().Set("assignToRoleId", value)
    if err != nil {
        panic(err)
    }
}
// SetStatements sets the statements property value. The statements property
func (m *AwsActionsPermissionsDefinitionAction) SetStatements(value []AwsStatementable)() {
    err := m.GetBackingStore().Set("statements", value)
    if err != nil {
        panic(err)
    }
}
// AwsActionsPermissionsDefinitionActionable 
type AwsActionsPermissionsDefinitionActionable interface {
    AwsPermissionsDefinitionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignToRoleId()(*string)
    GetStatements()([]AwsStatementable)
    SetAssignToRoleId(value *string)()
    SetStatements(value []AwsStatementable)()
}
