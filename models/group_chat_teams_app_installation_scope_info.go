package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupChatTeamsAppInstallationScopeInfo struct {
    TeamsAppInstallationScopeInfo
}
// NewGroupChatTeamsAppInstallationScopeInfo instantiates a new GroupChatTeamsAppInstallationScopeInfo and sets the default values.
func NewGroupChatTeamsAppInstallationScopeInfo()(*GroupChatTeamsAppInstallationScopeInfo) {
    m := &GroupChatTeamsAppInstallationScopeInfo{
        TeamsAppInstallationScopeInfo: *NewTeamsAppInstallationScopeInfo(),
    }
    odataTypeValue := "#microsoft.graph.groupChatTeamsAppInstallationScopeInfo"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGroupChatTeamsAppInstallationScopeInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupChatTeamsAppInstallationScopeInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupChatTeamsAppInstallationScopeInfo(), nil
}
// GetChatId gets the chatId property value. The chatId property
// returns a *string when successful
func (m *GroupChatTeamsAppInstallationScopeInfo) GetChatId()(*string) {
    val, err := m.GetBackingStore().Get("chatId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupChatTeamsAppInstallationScopeInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeamsAppInstallationScopeInfo.GetFieldDeserializers()
    res["chatId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GroupChatTeamsAppInstallationScopeInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeamsAppInstallationScopeInfo.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("chatId", m.GetChatId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChatId sets the chatId property value. The chatId property
func (m *GroupChatTeamsAppInstallationScopeInfo) SetChatId(value *string)() {
    err := m.GetBackingStore().Set("chatId", value)
    if err != nil {
        panic(err)
    }
}
type GroupChatTeamsAppInstallationScopeInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TeamsAppInstallationScopeInfoable
    GetChatId()(*string)
    SetChatId(value *string)()
}
