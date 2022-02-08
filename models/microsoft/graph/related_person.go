package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RelatedPerson 
type RelatedPerson struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the person.
    displayName *string;
    // Possible values are: manager, colleague, directReport, dotLineReport, assistant, dotLineManager, alternateContact, friend, spouse, sibling, child, parent, sponsor, emergencyContact, other, unknownFutureValue.
    relationship *PersonRelationship;
    // Email address or reference to person within organization.
    userPrincipalName *string;
}
// NewRelatedPerson instantiates a new relatedPerson and sets the default values.
func NewRelatedPerson()(*RelatedPerson) {
    m := &RelatedPerson{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RelatedPerson) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Name of the person.
func (m *RelatedPerson) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetRelationship gets the relationship property value. Possible values are: manager, colleague, directReport, dotLineReport, assistant, dotLineManager, alternateContact, friend, spouse, sibling, child, parent, sponsor, emergencyContact, other, unknownFutureValue.
func (m *RelatedPerson) GetRelationship()(*PersonRelationship) {
    if m == nil {
        return nil
    } else {
        return m.relationship
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. Email address or reference to person within organization.
func (m *RelatedPerson) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RelatedPerson) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["relationship"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePersonRelationship)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelationship(val.(*PersonRelationship))
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *RelatedPerson) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RelatedPerson) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetRelationship() != nil {
        cast := (*m.GetRelationship()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RelatedPerson) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Name of the person.
func (m *RelatedPerson) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetRelationship sets the relationship property value. Possible values are: manager, colleague, directReport, dotLineReport, assistant, dotLineManager, alternateContact, friend, spouse, sibling, child, parent, sponsor, emergencyContact, other, unknownFutureValue.
func (m *RelatedPerson) SetRelationship(value *PersonRelationship)() {
    if m != nil {
        m.relationship = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. Email address or reference to person within organization.
func (m *RelatedPerson) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
