package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnPremisesAgentGroup 
type OnPremisesAgentGroup struct {
    Entity
    // List of onPremisesAgent that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
    agents []OnPremisesAgentable;
    // Display name of the onPremisesAgentGroup.
    displayName *string;
    // Indicates if the onPremisesAgentGroup is the default agent group. Only a single agent group can be the default onPremisesAgentGroup and is set by the system.
    isDefault *bool;
    // List of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
    publishedResources []PublishedResourceable;
    // Possible values are: applicationProxy, exchangeOnline, authentication, provisioning, adAdministration.
    publishingType *OnPremisesPublishingType;
}
// NewOnPremisesAgentGroup instantiates a new onPremisesAgentGroup and sets the default values.
func NewOnPremisesAgentGroup()(*OnPremisesAgentGroup) {
    m := &OnPremisesAgentGroup{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOnPremisesAgentGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesAgentGroupFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOnPremisesAgentGroup(), nil
}
// GetAgents gets the agents property value. List of onPremisesAgent that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesAgentGroup) GetAgents()([]OnPremisesAgentable) {
    if m == nil {
        return nil
    } else {
        return m.agents
    }
}
// GetDisplayName gets the displayName property value. Display name of the onPremisesAgentGroup.
func (m *OnPremisesAgentGroup) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesAgentGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesAgentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesAgentable, len(val))
            for i, v := range val {
                res[i] = v.(OnPremisesAgentable)
            }
            m.SetAgents(res)
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
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["publishedResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePublishedResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PublishedResourceable, len(val))
            for i, v := range val {
                res[i] = v.(PublishedResourceable)
            }
            m.SetPublishedResources(res)
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
    return res
}
// GetIsDefault gets the isDefault property value. Indicates if the onPremisesAgentGroup is the default agent group. Only a single agent group can be the default onPremisesAgentGroup and is set by the system.
func (m *OnPremisesAgentGroup) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetPublishedResources gets the publishedResources property value. List of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesAgentGroup) GetPublishedResources()([]PublishedResourceable) {
    if m == nil {
        return nil
    } else {
        return m.publishedResources
    }
}
// GetPublishingType gets the publishingType property value. Possible values are: applicationProxy, exchangeOnline, authentication, provisioning, adAdministration.
func (m *OnPremisesAgentGroup) GetPublishingType()(*OnPremisesPublishingType) {
    if m == nil {
        return nil
    } else {
        return m.publishingType
    }
}
// Serialize serializes information the current object
func (m *OnPremisesAgentGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAgents() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgents()))
        for i, v := range m.GetAgents() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("agents", cast)
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
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    if m.GetPublishedResources() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPublishedResources()))
        for i, v := range m.GetPublishedResources() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("publishedResources", cast)
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
    return nil
}
// SetAgents sets the agents property value. List of onPremisesAgent that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesAgentGroup) SetAgents(value []OnPremisesAgentable)() {
    if m != nil {
        m.agents = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the onPremisesAgentGroup.
func (m *OnPremisesAgentGroup) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsDefault sets the isDefault property value. Indicates if the onPremisesAgentGroup is the default agent group. Only a single agent group can be the default onPremisesAgentGroup and is set by the system.
func (m *OnPremisesAgentGroup) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetPublishedResources sets the publishedResources property value. List of publishedResource that are assigned to an onPremisesAgentGroup. Read-only. Nullable.
func (m *OnPremisesAgentGroup) SetPublishedResources(value []PublishedResourceable)() {
    if m != nil {
        m.publishedResources = value
    }
}
// SetPublishingType sets the publishingType property value. Possible values are: applicationProxy, exchangeOnline, authentication, provisioning, adAdministration.
func (m *OnPremisesAgentGroup) SetPublishingType(value *OnPremisesPublishingType)() {
    if m != nil {
        m.publishingType = value
    }
}
