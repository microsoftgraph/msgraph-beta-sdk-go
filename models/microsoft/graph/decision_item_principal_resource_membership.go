package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DecisionItemPrincipalResourceMembership 
type DecisionItemPrincipalResourceMembership struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    membershipType *DecisionItemPrincipalResourceMembershipType;
}
// NewDecisionItemPrincipalResourceMembership instantiates a new decisionItemPrincipalResourceMembership and sets the default values.
func NewDecisionItemPrincipalResourceMembership()(*DecisionItemPrincipalResourceMembership) {
    m := &DecisionItemPrincipalResourceMembership{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DecisionItemPrincipalResourceMembership) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMembershipType gets the membershipType property value. 
func (m *DecisionItemPrincipalResourceMembership) GetMembershipType()(*DecisionItemPrincipalResourceMembershipType) {
    if m == nil {
        return nil
    } else {
        return m.membershipType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DecisionItemPrincipalResourceMembership) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["membershipType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDecisionItemPrincipalResourceMembershipType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipType(val.(*DecisionItemPrincipalResourceMembershipType))
        }
        return nil
    }
    return res
}
func (m *DecisionItemPrincipalResourceMembership) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DecisionItemPrincipalResourceMembership) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetMembershipType() != nil {
        cast := (*m.GetMembershipType()).String()
        err := writer.WriteStringValue("membershipType", &cast)
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
func (m *DecisionItemPrincipalResourceMembership) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMembershipType sets the membershipType property value. 
func (m *DecisionItemPrincipalResourceMembership) SetMembershipType(value *DecisionItemPrincipalResourceMembershipType)() {
    if m != nil {
        m.membershipType = value
    }
}
