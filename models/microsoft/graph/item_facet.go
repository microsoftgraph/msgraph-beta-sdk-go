package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemFacet 
type ItemFacet struct {
    Entity
    // The audiences that are able to see the values contained within the associated entity. Possible values are: me, family, contacts, groupMembers, organization, federatedOrganizations, everyone, unknownFutureValue.
    allowedAudiences *AllowedAudiences;
    // 
    createdBy *IdentitySet;
    // Provides the dateTimeOffset for when the entity was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Contains inference detail if the entity is inferred by the creating or modifying application.
    inference *InferenceData;
    // 
    isSearchable *bool;
    // 
    lastModifiedBy *IdentitySet;
    // Provides the dateTimeOffset for when the entity was created.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Where the values within an entity originated if synced from another service.
    source *PersonDataSources;
}
// NewItemFacet instantiates a new itemFacet and sets the default values.
func NewItemFacet()(*ItemFacet) {
    m := &ItemFacet{
        Entity: *NewEntity(),
    }
    return m
}
// GetAllowedAudiences gets the allowedAudiences property value. The audiences that are able to see the values contained within the associated entity. Possible values are: me, family, contacts, groupMembers, organization, federatedOrganizations, everyone, unknownFutureValue.
func (m *ItemFacet) GetAllowedAudiences()(*AllowedAudiences) {
    if m == nil {
        return nil
    } else {
        return m.allowedAudiences
    }
}
// GetCreatedBy gets the createdBy property value. 
func (m *ItemFacet) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetInference gets the inference property value. Contains inference detail if the entity is inferred by the creating or modifying application.
func (m *ItemFacet) GetInference()(*InferenceData) {
    if m == nil {
        return nil
    } else {
        return m.inference
    }
}
// GetIsSearchable gets the isSearchable property value. 
func (m *ItemFacet) GetIsSearchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSearchable
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. 
func (m *ItemFacet) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetSource gets the source property value. Where the values within an entity originated if synced from another service.
func (m *ItemFacet) GetSource()(*PersonDataSources) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemFacet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedAudiences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAllowedAudiences)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedAudiences(val.(*AllowedAudiences))
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["inference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInferenceData() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInference(val.(*InferenceData))
        }
        return nil
    }
    res["isSearchable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSearchable(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonDataSources() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*PersonDataSources))
        }
        return nil
    }
    return res
}
func (m *ItemFacet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ItemFacet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedAudiences() != nil {
        cast := (*m.GetAllowedAudiences()).String()
        err = writer.WriteStringValue("allowedAudiences", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inference", m.GetInference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedAudiences sets the allowedAudiences property value. The audiences that are able to see the values contained within the associated entity. Possible values are: me, family, contacts, groupMembers, organization, federatedOrganizations, everyone, unknownFutureValue.
func (m *ItemFacet) SetAllowedAudiences(value *AllowedAudiences)() {
    if m != nil {
        m.allowedAudiences = value
    }
}
// SetCreatedBy sets the createdBy property value. 
func (m *ItemFacet) SetCreatedBy(value *IdentitySet)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetInference sets the inference property value. Contains inference detail if the entity is inferred by the creating or modifying application.
func (m *ItemFacet) SetInference(value *InferenceData)() {
    if m != nil {
        m.inference = value
    }
}
// SetIsSearchable sets the isSearchable property value. 
func (m *ItemFacet) SetIsSearchable(value *bool)() {
    if m != nil {
        m.isSearchable = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. 
func (m *ItemFacet) SetLastModifiedBy(value *IdentitySet)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetSource sets the source property value. Where the values within an entity originated if synced from another service.
func (m *ItemFacet) SetSource(value *PersonDataSources)() {
    if m != nil {
        m.source = value
    }
}
