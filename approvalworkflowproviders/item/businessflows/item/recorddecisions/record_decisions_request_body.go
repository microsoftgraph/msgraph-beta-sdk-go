package recorddecisions

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RecordDecisionsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    justification *string;
    // 
    reviewResult *string;
}
// Instantiates a new recordDecisionsRequestBody and sets the default values.
func NewRecordDecisionsRequestBody()(*RecordDecisionsRequestBody) {
    m := &RecordDecisionsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordDecisionsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the justification property value. 
func (m *RecordDecisionsRequestBody) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// Gets the reviewResult property value. 
func (m *RecordDecisionsRequestBody) GetReviewResult()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reviewResult
    }
}
// The deserialization information for the current model
func (m *RecordDecisionsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJustification(val)
        return nil
    }
    res["reviewResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReviewResult(val)
        return nil
    }
    return res
}
func (m *RecordDecisionsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RecordDecisionsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reviewResult", m.GetReviewResult())
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
func (m *RecordDecisionsRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the justification property value. 
// Parameters:
//  - value : Value to set for the justification property.
func (m *RecordDecisionsRequestBody) SetJustification(value *string)() {
    m.justification = value
}
// Sets the reviewResult property value. 
// Parameters:
//  - value : Value to set for the reviewResult property.
func (m *RecordDecisionsRequestBody) SetReviewResult(value *string)() {
    m.reviewResult = value
}
