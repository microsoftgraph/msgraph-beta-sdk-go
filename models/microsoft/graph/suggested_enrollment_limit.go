package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SuggestedEnrollmentLimit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The suggested enrollment limit within a day
    suggestedDailyLimit *int32;
}
// Instantiates a new suggestedEnrollmentLimit and sets the default values.
func NewSuggestedEnrollmentLimit()(*SuggestedEnrollmentLimit) {
    m := &SuggestedEnrollmentLimit{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SuggestedEnrollmentLimit) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the suggestedDailyLimit property value. The suggested enrollment limit within a day
func (m *SuggestedEnrollmentLimit) GetSuggestedDailyLimit()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.suggestedDailyLimit
    }
}
// The deserialization information for the current model
func (m *SuggestedEnrollmentLimit) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["suggestedDailyLimit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSuggestedDailyLimit(val)
        return nil
    }
    return res
}
func (m *SuggestedEnrollmentLimit) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SuggestedEnrollmentLimit) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("suggestedDailyLimit", m.GetSuggestedDailyLimit())
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
func (m *SuggestedEnrollmentLimit) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the suggestedDailyLimit property value. The suggested enrollment limit within a day
// Parameters:
//  - value : Value to set for the suggestedDailyLimit property.
func (m *SuggestedEnrollmentLimit) SetSuggestedDailyLimit(value *int32)() {
    m.suggestedDailyLimit = value
}
