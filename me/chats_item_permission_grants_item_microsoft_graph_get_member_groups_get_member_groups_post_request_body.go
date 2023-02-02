package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody 
type ChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The securityEnabledOnly property
    securityEnabledOnly *bool
}
// NewChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody instantiates a new ChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody and sets the default values.
func NewChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody()(*ChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) {
    m := &ChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["securityEnabledOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEnabledOnly(val)
        }
        return nil
    }
    return res
}
// GetSecurityEnabledOnly gets the securityEnabledOnly property value. The securityEnabledOnly property
func (m *ChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) GetSecurityEnabledOnly()(*bool) {
    return m.securityEnabledOnly
}
// Serialize serializes information the current object
func (m *ChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("securityEnabledOnly", m.GetSecurityEnabledOnly())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSecurityEnabledOnly sets the securityEnabledOnly property value. The securityEnabledOnly property
func (m *ChatsItemPermissionGrantsItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) SetSecurityEnabledOnly(value *bool)() {
    m.securityEnabledOnly = value
}
