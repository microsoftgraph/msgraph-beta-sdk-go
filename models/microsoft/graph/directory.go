package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Directory struct {
    Entity
    // Conceptual container for user and group directory objects.
    administrativeUnits []AdministrativeUnit;
    // Recently deleted items. Read-only. Nullable.
    deletedItems []DirectoryObject;
    // Nullable.
    featureRolloutPolicies []FeatureRolloutPolicy;
    // Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
    federationConfigurations []IdentityProviderBase;
    // 
    sharedEmailDomains []SharedEmailDomain;
}
// Instantiates a new directory and sets the default values.
func NewDirectory()(*Directory) {
    m := &Directory{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the administrativeUnits property value. Conceptual container for user and group directory objects.
func (m *Directory) GetAdministrativeUnits()([]AdministrativeUnit) {
    if m == nil {
        return nil
    } else {
        return m.administrativeUnits
    }
}
// Gets the deletedItems property value. Recently deleted items. Read-only. Nullable.
func (m *Directory) GetDeletedItems()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.deletedItems
    }
}
// Gets the featureRolloutPolicies property value. Nullable.
func (m *Directory) GetFeatureRolloutPolicies()([]FeatureRolloutPolicy) {
    if m == nil {
        return nil
    } else {
        return m.featureRolloutPolicies
    }
}
// Gets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) GetFederationConfigurations()([]IdentityProviderBase) {
    if m == nil {
        return nil
    } else {
        return m.federationConfigurations
    }
}
// Gets the sharedEmailDomains property value. 
func (m *Directory) GetSharedEmailDomains()([]SharedEmailDomain) {
    if m == nil {
        return nil
    } else {
        return m.sharedEmailDomains
    }
}
// The deserialization information for the current model
func (m *Directory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdministrativeUnit() })
        if err != nil {
            return err
        }
        res := make([]AdministrativeUnit, len(val))
        for i, v := range val {
            res[i] = *(v.(*AdministrativeUnit))
        }
        m.SetAdministrativeUnits(res)
        return nil
    }
    res["deletedItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetDeletedItems(res)
        return nil
    }
    res["featureRolloutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFeatureRolloutPolicy() })
        if err != nil {
            return err
        }
        res := make([]FeatureRolloutPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*FeatureRolloutPolicy))
        }
        m.SetFeatureRolloutPolicies(res)
        return nil
    }
    res["federationConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityProviderBase() })
        if err != nil {
            return err
        }
        res := make([]IdentityProviderBase, len(val))
        for i, v := range val {
            res[i] = *(v.(*IdentityProviderBase))
        }
        m.SetFederationConfigurations(res)
        return nil
    }
    res["sharedEmailDomains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharedEmailDomain() })
        if err != nil {
            return err
        }
        res := make([]SharedEmailDomain, len(val))
        for i, v := range val {
            res[i] = *(v.(*SharedEmailDomain))
        }
        m.SetSharedEmailDomains(res)
        return nil
    }
    return res
}
func (m *Directory) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Directory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdministrativeUnits()))
        for i, v := range m.GetAdministrativeUnits() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("administrativeUnits", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeletedItems()))
        for i, v := range m.GetDeletedItems() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deletedItems", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFeatureRolloutPolicies()))
        for i, v := range m.GetFeatureRolloutPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFederationConfigurations()))
        for i, v := range m.GetFederationConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("federationConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSharedEmailDomains()))
        for i, v := range m.GetSharedEmailDomains() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sharedEmailDomains", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the administrativeUnits property value. Conceptual container for user and group directory objects.
// Parameters:
//  - value : Value to set for the administrativeUnits property.
func (m *Directory) SetAdministrativeUnits(value []AdministrativeUnit)() {
    m.administrativeUnits = value
}
// Sets the deletedItems property value. Recently deleted items. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the deletedItems property.
func (m *Directory) SetDeletedItems(value []DirectoryObject)() {
    m.deletedItems = value
}
// Sets the featureRolloutPolicies property value. Nullable.
// Parameters:
//  - value : Value to set for the featureRolloutPolicies property.
func (m *Directory) SetFeatureRolloutPolicies(value []FeatureRolloutPolicy)() {
    m.featureRolloutPolicies = value
}
// Sets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
// Parameters:
//  - value : Value to set for the federationConfigurations property.
func (m *Directory) SetFederationConfigurations(value []IdentityProviderBase)() {
    m.federationConfigurations = value
}
// Sets the sharedEmailDomains property value. 
// Parameters:
//  - value : Value to set for the sharedEmailDomains property.
func (m *Directory) SetSharedEmailDomains(value []SharedEmailDomain)() {
    m.sharedEmailDomains = value
}
