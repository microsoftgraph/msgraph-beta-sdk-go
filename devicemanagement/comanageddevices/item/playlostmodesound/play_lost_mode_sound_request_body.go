package playlostmodesound

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlayLostModeSoundRequestBody provides operations to call the playLostModeSound method.
type PlayLostModeSoundRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The durationInMinutes property
    durationInMinutes *string;
}
// NewPlayLostModeSoundRequestBody instantiates a new playLostModeSoundRequestBody and sets the default values.
func NewPlayLostModeSoundRequestBody()(*PlayLostModeSoundRequestBody) {
    m := &PlayLostModeSoundRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlayLostModeSoundRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlayLostModeSoundRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlayLostModeSoundRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlayLostModeSoundRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDurationInMinutes gets the durationInMinutes property value. The durationInMinutes property
func (m *PlayLostModeSoundRequestBody) GetDurationInMinutes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.durationInMinutes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlayLostModeSoundRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["durationInMinutes"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *PlayLostModeSoundRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PlayLostModeSoundRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDurationInMinutes sets the durationInMinutes property value. The durationInMinutes property
func (m *PlayLostModeSoundRequestBody) SetDurationInMinutes(value *string)() {
    if m != nil {
        m.durationInMinutes = value
    }
}
