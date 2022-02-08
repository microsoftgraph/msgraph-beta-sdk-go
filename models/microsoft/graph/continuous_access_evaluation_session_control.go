package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContinuousAccessEvaluationSessionControl 
type ContinuousAccessEvaluationSessionControl struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
    mode *ContinuousAccessEvaluationMode;
}
// NewContinuousAccessEvaluationSessionControl instantiates a new continuousAccessEvaluationSessionControl and sets the default values.
func NewContinuousAccessEvaluationSessionControl()(*ContinuousAccessEvaluationSessionControl) {
    m := &ContinuousAccessEvaluationSessionControl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContinuousAccessEvaluationSessionControl) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMode gets the mode property value. Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
func (m *ContinuousAccessEvaluationSessionControl) GetMode()(*ContinuousAccessEvaluationMode) {
    if m == nil {
        return nil
    } else {
        return m.mode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContinuousAccessEvaluationSessionControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["mode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseContinuousAccessEvaluationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMode(val.(*ContinuousAccessEvaluationMode))
        }
        return nil
    }
    return res
}
func (m *ContinuousAccessEvaluationSessionControl) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ContinuousAccessEvaluationSessionControl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetMode() != nil {
        cast := (*m.GetMode()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContinuousAccessEvaluationSessionControl) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMode sets the mode property value. Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
func (m *ContinuousAccessEvaluationSessionControl) SetMode(value *ContinuousAccessEvaluationMode)() {
    if m != nil {
        m.mode = value
    }
}
