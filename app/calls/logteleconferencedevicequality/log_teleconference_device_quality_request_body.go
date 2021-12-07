package logteleconferencedevicequality

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// LogTeleconferenceDeviceQualityRequestBody 
type LogTeleconferenceDeviceQualityRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    quality *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeleconferenceDeviceQuality;
}
// NewLogTeleconferenceDeviceQualityRequestBody instantiates a new logTeleconferenceDeviceQualityRequestBody and sets the default values.
func NewLogTeleconferenceDeviceQualityRequestBody()(*LogTeleconferenceDeviceQualityRequestBody) {
    m := &LogTeleconferenceDeviceQualityRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LogTeleconferenceDeviceQualityRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetQuality gets the quality property value. 
func (m *LogTeleconferenceDeviceQualityRequestBody) GetQuality()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeleconferenceDeviceQuality) {
    if m == nil {
        return nil
    } else {
        return m.quality
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LogTeleconferenceDeviceQualityRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["quality"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTeleconferenceDeviceQuality() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuality(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeleconferenceDeviceQuality))
        }
        return nil
    }
    return res
}
func (m *LogTeleconferenceDeviceQualityRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LogTeleconferenceDeviceQualityRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *LogTeleconferenceDeviceQualityRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetQuality sets the quality property value. 
func (m *LogTeleconferenceDeviceQualityRequestBody) SetQuality(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeleconferenceDeviceQuality)() {
    if m != nil {
        m.quality = value
    }
}
