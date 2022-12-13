package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemManagedDevicesItemPlayLostModeSoundPostRequestBody provides operations to call the playLostModeSound method.
type ItemManagedDevicesItemPlayLostModeSoundPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The durationInMinutes property
    durationInMinutes *string
}
// NewItemManagedDevicesItemPlayLostModeSoundPostRequestBody instantiates a new ItemManagedDevicesItemPlayLostModeSoundPostRequestBody and sets the default values.
func NewItemManagedDevicesItemPlayLostModeSoundPostRequestBody()(*ItemManagedDevicesItemPlayLostModeSoundPostRequestBody) {
    m := &ItemManagedDevicesItemPlayLostModeSoundPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemManagedDevicesItemPlayLostModeSoundPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemPlayLostModeSoundPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemPlayLostModeSoundPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesItemPlayLostModeSoundPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDurationInMinutes gets the durationInMinutes property value. The durationInMinutes property
func (m *ItemManagedDevicesItemPlayLostModeSoundPostRequestBody) GetDurationInMinutes()(*string) {
    return m.durationInMinutes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesItemPlayLostModeSoundPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["durationInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInMinutes(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemManagedDevicesItemPlayLostModeSoundPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("durationInMinutes", m.GetDurationInMinutes())
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
func (m *ItemManagedDevicesItemPlayLostModeSoundPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDurationInMinutes sets the durationInMinutes property value. The durationInMinutes property
func (m *ItemManagedDevicesItemPlayLostModeSoundPostRequestBody) SetDurationInMinutes(value *string)() {
    m.durationInMinutes = value
}
