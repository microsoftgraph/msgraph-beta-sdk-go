package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppCallsLogTeleconferenceDeviceQualityPostRequestBody provides operations to call the logTeleconferenceDeviceQuality method.
type AppCallsLogTeleconferenceDeviceQualityPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The quality property
    quality TeleconferenceDeviceQualityable
}
// NewAppCallsLogTeleconferenceDeviceQualityPostRequestBody instantiates a new AppCallsLogTeleconferenceDeviceQualityPostRequestBody and sets the default values.
func NewAppCallsLogTeleconferenceDeviceQualityPostRequestBody()(*AppCallsLogTeleconferenceDeviceQualityPostRequestBody) {
    m := &AppCallsLogTeleconferenceDeviceQualityPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppCallsLogTeleconferenceDeviceQualityPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppCallsLogTeleconferenceDeviceQualityPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppCallsLogTeleconferenceDeviceQualityPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppCallsLogTeleconferenceDeviceQualityPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppCallsLogTeleconferenceDeviceQualityPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["quality"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeleconferenceDeviceQualityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuality(val.(TeleconferenceDeviceQualityable))
        }
        return nil
    }
    return res
}
// GetQuality gets the quality property value. The quality property
func (m *AppCallsLogTeleconferenceDeviceQualityPostRequestBody) GetQuality()(TeleconferenceDeviceQualityable) {
    return m.quality
}
// Serialize serializes information the current object
func (m *AppCallsLogTeleconferenceDeviceQualityPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("quality", m.GetQuality())
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
func (m *AppCallsLogTeleconferenceDeviceQualityPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetQuality sets the quality property value. The quality property
func (m *AppCallsLogTeleconferenceDeviceQualityPostRequestBody) SetQuality(value TeleconferenceDeviceQualityable)() {
    m.quality = value
}
