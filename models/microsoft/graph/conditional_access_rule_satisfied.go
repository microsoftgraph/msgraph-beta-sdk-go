package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessRuleSatisfied 
type ConditionalAccessRuleSatisfied struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    conditionalAccessCondition *ConditionalAccessConditions;
    // 
    ruleSatisfied *ConditionalAccessRule;
}
// NewConditionalAccessRuleSatisfied instantiates a new conditionalAccessRuleSatisfied and sets the default values.
func NewConditionalAccessRuleSatisfied()(*ConditionalAccessRuleSatisfied) {
    m := &ConditionalAccessRuleSatisfied{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessRuleSatisfied) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConditionalAccessCondition gets the conditionalAccessCondition property value. 
func (m *ConditionalAccessRuleSatisfied) GetConditionalAccessCondition()(*ConditionalAccessConditions) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessCondition
    }
}
// GetRuleSatisfied gets the ruleSatisfied property value. 
func (m *ConditionalAccessRuleSatisfied) GetRuleSatisfied()(*ConditionalAccessRule) {
    if m == nil {
        return nil
    } else {
        return m.ruleSatisfied
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessRuleSatisfied) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["conditionalAccessCondition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessConditions)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ConditionalAccessConditions)
            m.SetConditionalAccessCondition(&cast)
        }
        return nil
    }
    res["ruleSatisfied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessRule)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ConditionalAccessRule)
            m.SetRuleSatisfied(&cast)
        }
        return nil
    }
    return res
}
func (m *ConditionalAccessRuleSatisfied) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConditionalAccessRuleSatisfied) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetConditionalAccessCondition() != nil {
        cast := m.GetConditionalAccessCondition().String()
        err := writer.WriteStringValue("conditionalAccessCondition", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRuleSatisfied() != nil {
        cast := m.GetRuleSatisfied().String()
        err := writer.WriteStringValue("ruleSatisfied", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessRuleSatisfied) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConditionalAccessCondition sets the conditionalAccessCondition property value. 
func (m *ConditionalAccessRuleSatisfied) SetConditionalAccessCondition(value *ConditionalAccessConditions)() {
    if m != nil {
        m.conditionalAccessCondition = value
    }
}
// SetRuleSatisfied sets the ruleSatisfied property value. 
func (m *ConditionalAccessRuleSatisfied) SetRuleSatisfied(value *ConditionalAccessRule)() {
    if m != nil {
        m.ruleSatisfied = value
    }
}
