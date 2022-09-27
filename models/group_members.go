package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupMembers 
type GroupMembers struct {
    UserSet
    // The name of the group in Azure AD. Read only.
    description *string
    // The ID of the group in Azure AD.
    id *string
}
// NewGroupMembers instantiates a new GroupMembers and sets the default values.
func NewGroupMembers()(*GroupMembers) {
    m := &GroupMembers{
        UserSet: *NewUserSet(),
    }
    odataTypeValue := "#microsoft.graph.groupMembers";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGroupMembersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupMembersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupMembers(), nil
}
// GetDescription gets the description property value. The name of the group in Azure AD. Read only.
func (m *GroupMembers) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupMembers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserSet.GetFieldDeserializers()
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    return res
}
// GetId gets the id property value. The ID of the group in Azure AD.
func (m *GroupMembers) GetId()(*string) {
    return m.id
}
// Serialize serializes information the current object
func (m *GroupMembers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserSet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The name of the group in Azure AD. Read only.
func (m *GroupMembers) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. The ID of the group in Azure AD.
func (m *GroupMembers) SetId(value *string)() {
    m.id = value
}
