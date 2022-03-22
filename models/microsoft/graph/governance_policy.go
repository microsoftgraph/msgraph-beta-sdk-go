package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernancePolicy 
type GovernancePolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    decisionMakerCriteria []GovernanceCriteriaable;
    // 
    notificationPolicy GovernanceNotificationPolicyable;
}
// NewGovernancePolicy instantiates a new governancePolicy and sets the default values.
func NewGovernancePolicy()(*GovernancePolicy) {
    m := &GovernancePolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGovernancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernancePolicyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGovernancePolicy(), nil
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
func (m *GovernancePolicy) GetDecisionMakerCriteria()([]GovernanceCriteriaable) {
    if m == nil {
        return nil
    } else {
        return m.decisionMakerCriteria
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernancePolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["decisionMakerCriteria"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceCriteriaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceCriteriaable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceCriteriaable)
            }
            m.SetDecisionMakerCriteria(res)
        }
        return nil
    }
    res["notificationPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceNotificationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationPolicy(val.(GovernanceNotificationPolicyable))
        }
        return nil
    }
    return res
}
// GetNotificationPolicy gets the notificationPolicy property value. 
func (m *GovernancePolicy) GetNotificationPolicy()(GovernanceNotificationPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.notificationPolicy
    }
}
// Serialize serializes information the current object
func (m *GovernancePolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDecisionMakerCriteria() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDecisionMakerCriteria()))
        for i, v := range m.GetDecisionMakerCriteria() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *GovernancePolicy) SetDecisionMakerCriteria(value []GovernanceCriteriaable)() {
    if m != nil {
        m.decisionMakerCriteria = value
    }
}
// SetNotificationPolicy sets the notificationPolicy property value. 
func (m *GovernancePolicy) SetNotificationPolicy(value GovernanceNotificationPolicyable)() {
    if m != nil {
        m.notificationPolicy = value
    }
}
