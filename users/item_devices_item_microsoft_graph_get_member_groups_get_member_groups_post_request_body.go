package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody 
type ItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The securityEnabledOnly property
    securityEnabledOnly *bool
}
// NewItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody instantiates a new ItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody and sets the default values.
func NewItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody()(*ItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) {
    m := &ItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) GetSecurityEnabledOnly()(*bool) {
    return m.securityEnabledOnly
}
// Serialize serializes information the current object
func (m *ItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSecurityEnabledOnly sets the securityEnabledOnly property value. The securityEnabledOnly property
func (m *ItemDevicesItemMicrosoftGraphGetMemberGroupsGetMemberGroupsPostRequestBody) SetSecurityEnabledOnly(value *bool)() {
    m.securityEnabledOnly = value
}
