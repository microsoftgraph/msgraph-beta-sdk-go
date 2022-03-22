package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EvaluateLabelJobResultGroup 
type EvaluateLabelJobResultGroup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    automatic EvaluateLabelJobResultable;
    // 
    recommended EvaluateLabelJobResultable;
}
// NewEvaluateLabelJobResultGroup instantiates a new evaluateLabelJobResultGroup and sets the default values.
func NewEvaluateLabelJobResultGroup()(*EvaluateLabelJobResultGroup) {
    m := &EvaluateLabelJobResultGroup{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEvaluateLabelJobResultGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEvaluateLabelJobResultGroupFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEvaluateLabelJobResultGroup(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateLabelJobResultGroup) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAutomatic gets the automatic property value. 
func (m *EvaluateLabelJobResultGroup) GetAutomatic()(EvaluateLabelJobResultable) {
    if m == nil {
        return nil
    } else {
        return m.automatic
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateLabelJobResultGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["automatic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEvaluateLabelJobResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomatic(val.(EvaluateLabelJobResultable))
        }
        return nil
    }
    res["recommended"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEvaluateLabelJobResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommended(val.(EvaluateLabelJobResultable))
        }
        return nil
    }
    return res
}
// GetRecommended gets the recommended property value. 
func (m *EvaluateLabelJobResultGroup) GetRecommended()(EvaluateLabelJobResultable) {
    if m == nil {
        return nil
    } else {
        return m.recommended
    }
}
// Serialize serializes information the current object
func (m *EvaluateLabelJobResultGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("automatic", m.GetAutomatic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recommended", m.GetRecommended())
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
func (m *EvaluateLabelJobResultGroup) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAutomatic sets the automatic property value. 
func (m *EvaluateLabelJobResultGroup) SetAutomatic(value EvaluateLabelJobResultable)() {
    if m != nil {
        m.automatic = value
    }
}
// SetRecommended sets the recommended property value. 
func (m *EvaluateLabelJobResultGroup) SetRecommended(value EvaluateLabelJobResultable)() {
    if m != nil {
        m.recommended = value
    }
}
