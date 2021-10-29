package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PublishedResource struct {
    Entity
    // List of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
    agentGroups []OnPremisesAgentGroup;
    // Display Name of the publishedResource.
    displayName *string;
    // Possible values are: applicationProxy, exchangeOnline, authentication, provisioning, adAdministration.
    publishingType *OnPremisesPublishingType;
    // Name of the publishedResource.
    resourceName *string;
}
// Instantiates a new publishedResource and sets the default values.
func NewPublishedResource()(*PublishedResource) {
    m := &PublishedResource{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the agentGroups property value. List of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *PublishedResource) GetAgentGroups()([]OnPremisesAgentGroup) {
    if m == nil {
        return nil
    } else {
        return m.agentGroups
    }
}
// Gets the displayName property value. Display Name of the publishedResource.
func (m *PublishedResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the publishingType property value. Possible values are: applicationProxy, exchangeOnline, authentication, provisioning, adAdministration.
func (m *PublishedResource) GetPublishingType()(*OnPremisesPublishingType) {
    if m == nil {
        return nil
    } else {
        return m.publishingType
    }
}
// Gets the resourceName property value. Name of the publishedResource.
func (m *PublishedResource) GetResourceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceName
    }
}
// The deserialization information for the current model
func (m *PublishedResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesAgentGroup() })
        if err != nil {
            return err
        }
        res := make([]OnPremisesAgentGroup, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnPremisesAgentGroup))
        }
        m.SetAgentGroups(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["publishingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnPremisesPublishingType)
        if err != nil {
            return err
        }
        cast := val.(OnPremisesPublishingType)
        m.SetPublishingType(&cast)
        return nil
    }
    res["resourceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceName(val)
        return nil
    }
    return res
}
func (m *PublishedResource) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PublishedResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgentGroups()))
        for i, v := range m.GetAgentGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("agentGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetPublishingType() != nil {
        cast := m.GetPublishingType().String()
        err = writer.WriteStringValue("publishingType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the agentGroups property value. List of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the agentGroups property.
func (m *PublishedResource) SetAgentGroups(value []OnPremisesAgentGroup)() {
    m.agentGroups = value
}
// Sets the displayName property value. Display Name of the publishedResource.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PublishedResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the publishingType property value. Possible values are: applicationProxy, exchangeOnline, authentication, provisioning, adAdministration.
// Parameters:
//  - value : Value to set for the publishingType property.
func (m *PublishedResource) SetPublishingType(value *OnPremisesPublishingType)() {
    m.publishingType = value
}
// Sets the resourceName property value. Name of the publishedResource.
// Parameters:
//  - value : Value to set for the resourceName property.
func (m *PublishedResource) SetResourceName(value *string)() {
    m.resourceName = value
}
