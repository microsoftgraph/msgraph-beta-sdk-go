package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GovernancePolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    decisionMakerCriteria []GovernanceCriteria;
    // 
    notificationPolicy *GovernanceNotificationPolicy;
}
// Instantiates a new governancePolicy and sets the default values.
func NewGovernancePolicy()(*GovernancePolicy) {
    m := &GovernancePolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernancePolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the decisionMakerCriteria property value. 
func (m *GovernancePolicy) GetDecisionMakerCriteria()([]GovernanceCriteria) {
    if m == nil {
        return nil
    } else {
        return m.decisionMakerCriteria
    }
}
// Gets the notificationPolicy property value. 
func (m *GovernancePolicy) GetNotificationPolicy()(*GovernanceNotificationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.notificationPolicy
    }
}
// The deserialization information for the current model
func (m *GovernancePolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["decisionMakerCriteria"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceCriteria() })
        if err != nil {
            return err
        }
        res := make([]GovernanceCriteria, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernanceCriteria))
        }
        m.SetDecisionMakerCriteria(res)
        return nil
    }
    res["notificationPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceNotificationPolicy() })
        if err != nil {
            return err
        }
        m.SetNotificationPolicy(val.(*GovernanceNotificationPolicy))
        return nil
    }
    return res
}
func (m *GovernancePolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GovernancePolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *GovernancePolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the decisionMakerCriteria property value. 
// Parameters:
//  - value : Value to set for the decisionMakerCriteria property.
func (m *GovernancePolicy) SetDecisionMakerCriteria(value []GovernanceCriteria)() {
    m.decisionMakerCriteria = value
}
// Sets the notificationPolicy property value. 
// Parameters:
//  - value : Value to set for the notificationPolicy property.
func (m *GovernancePolicy) SetNotificationPolicy(value *GovernanceNotificationPolicy)() {
    m.notificationPolicy = value
}
