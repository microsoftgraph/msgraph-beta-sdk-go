package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ActionStep 
type ActionStep struct {
    // 
    actionUrl *ActionUrl;
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
// GetActionUrl gets the actionUrl property value. 
func (m *ActionStep) GetActionUrl()(*ActionUrl) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *ActionStep) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewActionUrl() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionUrl(val.(*ActionUrl))
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
func (m *ActionStep) SetActionUrl(value *ActionUrl)() {
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
