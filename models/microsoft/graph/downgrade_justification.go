package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DowngradeJustification struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether the downgrade is or is not justified.
    isDowngradeJustified *bool;
    // Message that indicates why a downgrade is justified. The message will appear in administrative logs.
    justificationMessage *string;
}
// Instantiates a new downgradeJustification and sets the default values.
func NewDowngradeJustification()(*DowngradeJustification) {
    m := &DowngradeJustification{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DowngradeJustification) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the isDowngradeJustified property value. Indicates whether the downgrade is or is not justified.
func (m *DowngradeJustification) GetIsDowngradeJustified()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDowngradeJustified
    }
}
// Gets the justificationMessage property value. Message that indicates why a downgrade is justified. The message will appear in administrative logs.
func (m *DowngradeJustification) GetJustificationMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justificationMessage
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DowngradeJustification) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the isDowngradeJustified property value. Indicates whether the downgrade is or is not justified.
// Parameters:
//  - value : Value to set for the isDowngradeJustified property.
func (m *DowngradeJustification) SetIsDowngradeJustified(value *bool)() {
    m.isDowngradeJustified = value
}
// Sets the justificationMessage property value. Message that indicates why a downgrade is justified. The message will appear in administrative logs.
// Parameters:
//  - value : Value to set for the justificationMessage property.
func (m *DowngradeJustification) SetJustificationMessage(value *string)() {
    m.justificationMessage = value
}
