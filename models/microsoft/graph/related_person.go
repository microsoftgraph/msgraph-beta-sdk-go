package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RelatedPerson struct {
    additionalData map[string]interface{};
    displayName *string;
    relationship *PersonRelationship;
    userPrincipalName *string;
}
func NewRelatedPerson()(*RelatedPerson) {
    m := &RelatedPerson{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RelatedPerson) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RelatedPerson) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *RelatedPerson) GetRelationship()(*PersonRelationship) {
    if m == nil {
        return nil
    } else {
        return m.relationship
    }
}
func (m *RelatedPerson) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *RelatedPerson) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["relationship"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePersonRelationship)
        if err != nil {
            return err
        }
        cast := val.(PersonRelationship)
        m.SetRelationship(&cast)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *RelatedPerson) IsNil()(bool) {
    return m == nil
}
func (m *RelatedPerson) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetRelationship() != nil {
        cast := m.GetRelationship().String()
        err := writer.WriteStringValue("relationship", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *RelatedPerson) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RelatedPerson) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *RelatedPerson) SetRelationship(value *PersonRelationship)() {
    m.relationship = value
}
func (m *RelatedPerson) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
