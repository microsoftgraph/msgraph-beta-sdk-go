package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DowngradeJustification 
type DowngradeJustification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether the downgrade is or is not justified.
    isDowngradeJustified *bool;
    // Message that indicates why a downgrade is justified. The message will appear in administrative logs.
    justificationMessage *string;
}
// NewDowngradeJustification instantiates a new downgradeJustification and sets the default values.
func NewDowngradeJustification()(*DowngradeJustification) {
    m := &DowngradeJustification{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DowngradeJustification) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsDowngradeJustified gets the isDowngradeJustified property value. Indicates whether the downgrade is or is not justified.
func (m *DowngradeJustification) GetIsDowngradeJustified()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDowngradeJustified
    }
}
// GetJustificationMessage gets the justificationMessage property value. Message that indicates why a downgrade is justified. The message will appear in administrative logs.
func (m *DowngradeJustification) GetJustificationMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justificationMessage
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DowngradeJustification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isDowngradeJustified"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDowngradeJustified(val)
        }
        return nil
    }
    res["justificationMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustificationMessage(val)
        }
        return nil
    }
    return res
}
func (m *DowngradeJustification) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DowngradeJustification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDowngradeJustified", m.GetIsDowngradeJustified())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("justificationMessage", m.GetJustificationMessage())
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
func (m *DowngradeJustification) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsDowngradeJustified sets the isDowngradeJustified property value. Indicates whether the downgrade is or is not justified.
func (m *DowngradeJustification) SetIsDowngradeJustified(value *bool)() {
    if m != nil {
        m.isDowngradeJustified = value
    }
}
// SetJustificationMessage sets the justificationMessage property value. Message that indicates why a downgrade is justified. The message will appear in administrative logs.
func (m *DowngradeJustification) SetJustificationMessage(value *string)() {
    if m != nil {
        m.justificationMessage = value
    }
}
