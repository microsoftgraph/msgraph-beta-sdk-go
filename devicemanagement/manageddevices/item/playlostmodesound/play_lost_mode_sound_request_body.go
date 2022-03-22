package playlostmodesound

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlayLostModeSoundRequestBody provides operations to call the playLostModeSound method.
type PlayLostModeSoundRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
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
func CreatePlayLostModeSoundRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
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
// GetDurationInMinutes gets the durationInMinutes property value. 
func (m *PlayLostModeSoundRequestBody) GetDurationInMinutes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.durationInMinutes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlayLostModeSoundRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["durationInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *PlayLostModeSoundRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetDurationInMinutes sets the durationInMinutes property value. 
func (m *PlayLostModeSoundRequestBody) SetDurationInMinutes(value *string)() {
    if m != nil {
        m.durationInMinutes = value
    }
}
