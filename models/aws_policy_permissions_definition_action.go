package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AwsPolicyPermissionsDefinitionAction 
type AwsPolicyPermissionsDefinitionAction struct {
    AwsPermissionsDefinitionAction
}
// NewAwsPolicyPermissionsDefinitionAction instantiates a new awsPolicyPermissionsDefinitionAction and sets the default values.
func NewAwsPolicyPermissionsDefinitionAction()(*AwsPolicyPermissionsDefinitionAction) {
    m := &AwsPolicyPermissionsDefinitionAction{
        AwsPermissionsDefinitionAction: *NewAwsPermissionsDefinitionAction(),
    }
    odataTypeValue := "#microsoft.graph.awsPolicyPermissionsDefinitionAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsPolicyPermissionsDefinitionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsPolicyPermissionsDefinitionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsPolicyPermissionsDefinitionAction(), nil
}
// GetAssignToRoleId gets the assignToRoleId property value. The assignToRoleId property
func (m *AwsPolicyPermissionsDefinitionAction) GetAssignToRoleId()(*string) {
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
func (m *AwsPolicyPermissionsDefinitionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["policies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionsDefinitionAwsPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionsDefinitionAwsPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PermissionsDefinitionAwsPolicyable)
                }
            }
            m.SetPolicies(res)
        }
        return nil
    }
    return res
}
// GetPolicies gets the policies property value. The policies property
func (m *AwsPolicyPermissionsDefinitionAction) GetPolicies()([]PermissionsDefinitionAwsPolicyable) {
    val, err := m.GetBackingStore().Get("policies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PermissionsDefinitionAwsPolicyable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsPolicyPermissionsDefinitionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPolicies()))
        for i, v := range m.GetPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("policies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignToRoleId sets the assignToRoleId property value. The assignToRoleId property
func (m *AwsPolicyPermissionsDefinitionAction) SetAssignToRoleId(value *string)() {
    err := m.GetBackingStore().Set("assignToRoleId", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicies sets the policies property value. The policies property
func (m *AwsPolicyPermissionsDefinitionAction) SetPolicies(value []PermissionsDefinitionAwsPolicyable)() {
    err := m.GetBackingStore().Set("policies", value)
    if err != nil {
        panic(err)
    }
}
// AwsPolicyPermissionsDefinitionActionable 
type AwsPolicyPermissionsDefinitionActionable interface {
    AwsPermissionsDefinitionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignToRoleId()(*string)
    GetPolicies()([]PermissionsDefinitionAwsPolicyable)
    SetAssignToRoleId(value *string)()
    SetPolicies(value []PermissionsDefinitionAwsPolicyable)()
}
