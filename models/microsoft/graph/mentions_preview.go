package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MentionsPreview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // True if the signed-in user is mentioned in the parent resource instance. Read-only. Supports filter.
    isMentioned *bool;
}
// Instantiates a new mentionsPreview and sets the default values.
func NewMentionsPreview()(*MentionsPreview) {
    m := &MentionsPreview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MentionsPreview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the isMentioned property value. True if the signed-in user is mentioned in the parent resource instance. Read-only. Supports filter.
func (m *MentionsPreview) GetIsMentioned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMentioned
    }
}
// The deserialization information for the current model
func (m *MentionsPreview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isMentioned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMentioned(val)
        return nil
    }
    return res
}
func (m *MentionsPreview) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MentionsPreview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isMentioned", m.GetIsMentioned())
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
func (m *MentionsPreview) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the isMentioned property value. True if the signed-in user is mentioned in the parent resource instance. Read-only. Supports filter.
// Parameters:
//  - value : Value to set for the isMentioned property.
func (m *MentionsPreview) SetIsMentioned(value *bool)() {
    m.isMentioned = value
}
