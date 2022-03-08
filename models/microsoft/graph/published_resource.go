package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PublishedResource provides operations to manage the collection of onPremisesPublishingProfile entities.
type PublishedResource struct {
    Entity
    // List of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
    agentGroups []OnPremisesAgentGroupable;
    // Display Name of the publishedResource.
    displayName *string;
    // Possible values are: applicationProxy, exchangeOnline, authentication, provisioning, adAdministration.
    publishingType *OnPremisesPublishingType;
    // Name of the publishedResource.
    resourceName *string;
}
// NewPublishedResource instantiates a new publishedResource and sets the default values.
func NewPublishedResource()(*PublishedResource) {
    m := &PublishedResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePublishedResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePublishedResourceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPublishedResource(), nil
}
// GetAgentGroups gets the agentGroups property value. List of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *PublishedResource) GetAgentGroups()([]OnPremisesAgentGroupable) {
    if m == nil {
        return nil
    } else {
        return m.agentGroups
    }
}
// GetDisplayName gets the displayName property value. Display Name of the publishedResource.
func (m *PublishedResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PublishedResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agentGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesAgentGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesAgentGroupable, len(val))
            for i, v := range val {
                res[i] = v.(OnPremisesAgentGroupable)
            }
            m.SetAgentGroups(res)
        }
        return nil
    }
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
    res["publishingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnPremisesPublishingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingType(val.(*OnPremisesPublishingType))
        }
        return nil
    }
    res["resourceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
        }
        return nil
    }
    return res
}
// GetPublishingType gets the publishingType property value. Possible values are: applicationProxy, exchangeOnline, authentication, provisioning, adAdministration.
func (m *PublishedResource) GetPublishingType()(*OnPremisesPublishingType) {
    if m == nil {
        return nil
    } else {
        return m.publishingType
    }
}
// GetResourceName gets the resourceName property value. Name of the publishedResource.
func (m *PublishedResource) GetResourceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceName
    }
}
func (m *PublishedResource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PublishedResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAgentGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgentGroups()))
        for i, v := range m.GetAgentGroups() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
        cast := (*m.GetPublishingType()).String()
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
// SetAgentGroups sets the agentGroups property value. List of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *PublishedResource) SetAgentGroups(value []OnPremisesAgentGroupable)() {
    if m != nil {
        m.agentGroups = value
    }
}
// SetDisplayName sets the displayName property value. Display Name of the publishedResource.
func (m *PublishedResource) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetPublishingType sets the publishingType property value. Possible values are: applicationProxy, exchangeOnline, authentication, provisioning, adAdministration.
func (m *PublishedResource) SetPublishingType(value *OnPremisesPublishingType)() {
    if m != nil {
        m.publishingType = value
    }
}
// SetResourceName sets the resourceName property value. Name of the publishedResource.
func (m *PublishedResource) SetResourceName(value *string)() {
    if m != nil {
        m.resourceName = value
    }
}
