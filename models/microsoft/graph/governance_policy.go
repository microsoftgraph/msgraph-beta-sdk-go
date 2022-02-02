package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernancePolicy 
type GovernancePolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    decisionMakerCriteria []GovernanceCriteria;
    // 
    notificationPolicy *GovernanceNotificationPolicy;
}
// NewGovernancePolicy instantiates a new governancePolicy and sets the default values.
func NewGovernancePolicy()(*GovernancePolicy) {
    m := &GovernancePolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernancePolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDecisionMakerCriteria gets the decisionMakerCriteria property value. 
func (m *GovernancePolicy) GetDecisionMakerCriteria()([]GovernanceCriteria) {
    if m == nil {
        return nil
    } else {
        return m.decisionMakerCriteria
    }
}
// GetNotificationPolicy gets the notificationPolicy property value. 
func (m *GovernancePolicy) GetNotificationPolicy()(*GovernanceNotificationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.notificationPolicy
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernancePolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["decisionMakerCriteria"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceCriteria() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceCriteria, len(val))
            for i, v := range val {
                res[i] = *(v.(*GovernanceCriteria))
            }
            m.SetDecisionMakerCriteria(res)
        }
        return nil
    }
    res["notificationPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceNotificationPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationPolicy(val.(*GovernanceNotificationPolicy))
        }
        return nil
    }
    return res
}
func (m *GovernancePolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GovernancePolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDecisionMakerCriteria() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDecisionMakerCriteria()))
        for i, v := range m.GetDecisionMakerCriteria() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("decisionMakerCriteria", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notificationPolicy", m.GetNotificationPolicy())
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
func (m *GovernancePolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDecisionMakerCriteria sets the decisionMakerCriteria property value. 
func (m *GovernancePolicy) SetDecisionMakerCriteria(value []GovernanceCriteria)() {
    if m != nil {
        m.decisionMakerCriteria = value
    }
}
// SetNotificationPolicy sets the notificationPolicy property value. 
func (m *GovernancePolicy) SetNotificationPolicy(value *GovernanceNotificationPolicy)() {
    if m != nil {
        m.notificationPolicy = value
    }
}
