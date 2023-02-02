package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody 
type ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The durationInMinutes property
    durationInMinutes *string
}
// NewComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody instantiates a new ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody()(*ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody) {
    m := &ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDurationInMinutes gets the durationInMinutes property value. The durationInMinutes property
func (m *ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody) GetDurationInMinutes()(*string) {
    return m.durationInMinutes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDurationInMinutes sets the durationInMinutes property value. The durationInMinutes property
func (m *ComanagedDevicesItemMicrosoftGraphPlayLostModeSoundPlayLostModeSoundPostRequestBody) SetDurationInMinutes(value *string)() {
    m.durationInMinutes = value
}
