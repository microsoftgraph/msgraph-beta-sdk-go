package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OcrSettings 
type OcrSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether or not OCR is enabled for the case.
    isEnabled *bool;
    // Maximum image size that will be processed in KB).
    maxImageSize *int32;
    // The timeout duration for the OCR engine. A longer timeout may increase success of OCR, but may add to the total processing time.
    timeout *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
}
// NewOcrSettings instantiates a new ocrSettings and sets the default values.
func NewOcrSettings()(*OcrSettings) {
    m := &OcrSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OcrSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsEnabled gets the isEnabled property value. Indicates whether or not OCR is enabled for the case.
func (m *OcrSettings) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetMaxImageSize gets the maxImageSize property value. Maximum image size that will be processed in KB).
func (m *OcrSettings) GetMaxImageSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxImageSize
    }
}
// GetTimeout gets the timeout property value. The timeout duration for the OCR engine. A longer timeout may increase success of OCR, but may add to the total processing time.
func (m *OcrSettings) GetTimeout()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.timeout
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OcrSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["maxImageSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxImageSize(val)
        }
        return nil
    }
    res["timeout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeout(val)
        }
        return nil
    }
    return res
}
func (m *OcrSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OcrSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxImageSize", m.GetMaxImageSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("timeout", m.GetTimeout())
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
func (m *OcrSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsEnabled sets the isEnabled property value. Indicates whether or not OCR is enabled for the case.
func (m *OcrSettings) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetMaxImageSize sets the maxImageSize property value. Maximum image size that will be processed in KB).
func (m *OcrSettings) SetMaxImageSize(value *int32)() {
    if m != nil {
        m.maxImageSize = value
    }
}
// SetTimeout sets the timeout property value. The timeout duration for the OCR engine. A longer timeout may increase success of OCR, but may add to the total processing time.
func (m *OcrSettings) SetTimeout(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.timeout = value
    }
}
