package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ContinuousAccessEvaluationSessionControl struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
    mode *ContinuousAccessEvaluationMode;
}
// Instantiates a new continuousAccessEvaluationSessionControl and sets the default values.
func NewContinuousAccessEvaluationSessionControl()(*ContinuousAccessEvaluationSessionControl) {
    m := &ContinuousAccessEvaluationSessionControl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContinuousAccessEvaluationSessionControl) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the mode property value. Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
func (m *ContinuousAccessEvaluationSessionControl) GetMode()(*ContinuousAccessEvaluationMode) {
    if m == nil {
        return nil
    } else {
        return m.mode
    }
}
// The deserialization information for the current model
func (m *ContinuousAccessEvaluationSessionControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["mode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseContinuousAccessEvaluationMode)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ContinuousAccessEvaluationMode)
            m.SetMode(&cast)
        }
        return nil
    }
    return res
}
func (m *ContinuousAccessEvaluationSessionControl) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ContinuousAccessEvaluationSessionControl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetMode() != nil {
        cast := m.GetMode().String()
        err := writer.WriteStringValue("mode", &cast)
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
func (m *ContinuousAccessEvaluationSessionControl) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the mode property value. Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
// Parameters:
//  - value : Value to set for the mode property.
func (m *ContinuousAccessEvaluationSessionControl) SetMode(value *ContinuousAccessEvaluationMode)() {
    m.mode = value
}
