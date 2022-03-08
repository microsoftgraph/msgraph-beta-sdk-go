package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ActionStep provides operations to manage the directory singleton.
type ActionStep struct {
    // 
    actionUrl ActionUrlable;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    stepNumber *int64;
    // 
    text *string;
}
// NewActionStep instantiates a new actionStep and sets the default values.
func NewActionStep()(*ActionStep) {
    m := &ActionStep{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateActionStepFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActionStepFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewActionStep(), nil
}
// GetActionUrl gets the actionUrl property value. 
func (m *ActionStep) GetActionUrl()(ActionUrlable) {
    if m == nil {
        return nil
    } else {
        return m.actionUrl
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActionStep) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActionStep) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateActionUrlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionUrl(val.(ActionUrlable))
        }
        return nil
    }
    res["stepNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStepNumber(val)
        }
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
// GetStepNumber gets the stepNumber property value. 
func (m *ActionStep) GetStepNumber()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.stepNumber
    }
}
// GetText gets the text property value. 
func (m *ActionStep) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
func (m *ActionStep) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ActionStep) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("actionUrl", m.GetActionUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("stepNumber", m.GetStepNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
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
// SetActionUrl sets the actionUrl property value. 
func (m *ActionStep) SetActionUrl(value ActionUrlable)() {
    if m != nil {
        m.actionUrl = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActionStep) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetStepNumber sets the stepNumber property value. 
func (m *ActionStep) SetStepNumber(value *int64)() {
    if m != nil {
        m.stepNumber = value
    }
}
// SetText sets the text property value. 
func (m *ActionStep) SetText(value *string)() {
    if m != nil {
        m.text = value
    }
}
