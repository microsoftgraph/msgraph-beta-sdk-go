package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse 
type DevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
    // The value property
    value []string
}
// NewDevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse instantiates a new DevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse and sets the default values.
func NewDevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse()(*DevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse) {
    m := &DevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse) GetValue()([]string) {
    return m.value
}
// Serialize serializes information the current object
func (m *DevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        err = writer.WriteCollectionOfStringValues("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DevicesItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsResponse) SetValue(value []string)() {
    m.value = value
}
