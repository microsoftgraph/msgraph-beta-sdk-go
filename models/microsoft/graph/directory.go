package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Directory 
type Directory struct {
    Entity
    // Conceptual container for user and group directory objects.
    administrativeUnits []AdministrativeUnit;
    // Group of related custom security attribute definitions.
    attributeSets []AttributeSet;
    // Schema of a custom security attributes (key-value pairs).
    customSecurityAttributeDefinitions []CustomSecurityAttributeDefinition;
    // Recently deleted items. Read-only. Nullable.
    deletedItems []DirectoryObject;
    // Nullable.
    featureRolloutPolicies []FeatureRolloutPolicy;
    // Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
    federationConfigurations []IdentityProviderBase;
    // 
    impactedResources []RecommendationResource;
    // 
    inboundSharedUserProfiles []InboundSharedUserProfile;
    // 
    outboundSharedUserProfiles []OutboundSharedUserProfile;
    // 
    recommendations []Recommendation;
    // 
    sharedEmailDomains []SharedEmailDomain;
}
// NewDirectory instantiates a new directory and sets the default values.
func NewDirectory()(*Directory) {
    m := &Directory{
        Entity: *NewEntity(),
    }
    return m
}
// GetAdministrativeUnits gets the administrativeUnits property value. Conceptual container for user and group directory objects.
func (m *Directory) GetAdministrativeUnits()([]AdministrativeUnit) {
    if m == nil {
        return nil
    } else {
        return m.administrativeUnits
    }
}
// GetAttributeSets gets the attributeSets property value. Group of related custom security attribute definitions.
func (m *Directory) GetAttributeSets()([]AttributeSet) {
    if m == nil {
        return nil
    } else {
        return m.attributeSets
    }
}
// GetCustomSecurityAttributeDefinitions gets the customSecurityAttributeDefinitions property value. Schema of a custom security attributes (key-value pairs).
func (m *Directory) GetCustomSecurityAttributeDefinitions()([]CustomSecurityAttributeDefinition) {
    if m == nil {
        return nil
    } else {
        return m.customSecurityAttributeDefinitions
    }
}
// GetDeletedItems gets the deletedItems property value. Recently deleted items. Read-only. Nullable.
func (m *Directory) GetDeletedItems()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.deletedItems
    }
}
// GetFeatureRolloutPolicies gets the featureRolloutPolicies property value. Nullable.
func (m *Directory) GetFeatureRolloutPolicies()([]FeatureRolloutPolicy) {
    if m == nil {
        return nil
    } else {
        return m.featureRolloutPolicies
    }
}
// GetFederationConfigurations gets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) GetFederationConfigurations()([]IdentityProviderBase) {
    if m == nil {
        return nil
    } else {
        return m.federationConfigurations
    }
}
// GetImpactedResources gets the impactedResources property value. 
func (m *Directory) GetImpactedResources()([]RecommendationResource) {
    if m == nil {
        return nil
    } else {
        return m.impactedResources
    }
}
// GetInboundSharedUserProfiles gets the inboundSharedUserProfiles property value. 
func (m *Directory) GetInboundSharedUserProfiles()([]InboundSharedUserProfile) {
    if m == nil {
        return nil
    } else {
        return m.inboundSharedUserProfiles
    }
}
// GetOutboundSharedUserProfiles gets the outboundSharedUserProfiles property value. 
func (m *Directory) GetOutboundSharedUserProfiles()([]OutboundSharedUserProfile) {
    if m == nil {
        return nil
    } else {
        return m.outboundSharedUserProfiles
    }
}
// GetRecommendations gets the recommendations property value. 
func (m *Directory) GetRecommendations()([]Recommendation) {
    if m == nil {
        return nil
    } else {
        return m.recommendations
    }
}
// GetSharedEmailDomains gets the sharedEmailDomains property value. 
func (m *Directory) GetSharedEmailDomains()([]SharedEmailDomain) {
    if m == nil {
        return nil
    } else {
        return m.sharedEmailDomains
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Directory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdministrativeUnit() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AdministrativeUnit, len(val))
            for i, v := range val {
                res[i] = *(v.(*AdministrativeUnit))
            }
            m.SetAdministrativeUnits(res)
        }
        return nil
    }
    res["attributeSets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttributeSet() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeSet, len(val))
            for i, v := range val {
                res[i] = *(v.(*AttributeSet))
            }
            m.SetAttributeSets(res)
        }
        return nil
    }
    res["customSecurityAttributeDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomSecurityAttributeDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomSecurityAttributeDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*CustomSecurityAttributeDefinition))
            }
            m.SetCustomSecurityAttributeDefinitions(res)
        }
        return nil
    }
    res["deletedItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetDeletedItems(res)
        }
        return nil
    }
    res["featureRolloutPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFeatureRolloutPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FeatureRolloutPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*FeatureRolloutPolicy))
            }
            m.SetFeatureRolloutPolicies(res)
        }
        return nil
    }
    res["federationConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityProviderBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentityProviderBase))
            }
            m.SetFederationConfigurations(res)
        }
        return nil
    }
    res["impactedResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecommendationResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RecommendationResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*RecommendationResource))
            }
            m.SetImpactedResources(res)
        }
        return nil
    }
    res["inboundSharedUserProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInboundSharedUserProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InboundSharedUserProfile, len(val))
            for i, v := range val {
                res[i] = *(v.(*InboundSharedUserProfile))
            }
            m.SetInboundSharedUserProfiles(res)
        }
        return nil
    }
    res["outboundSharedUserProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutboundSharedUserProfile() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutboundSharedUserProfile, len(val))
            for i, v := range val {
                res[i] = *(v.(*OutboundSharedUserProfile))
            }
            m.SetOutboundSharedUserProfiles(res)
        }
        return nil
    }
    res["recommendations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecommendation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recommendation, len(val))
            for i, v := range val {
                res[i] = *(v.(*Recommendation))
            }
            m.SetRecommendations(res)
        }
        return nil
    }
    res["sharedEmailDomains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharedEmailDomain() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharedEmailDomain, len(val))
            for i, v := range val {
                res[i] = *(v.(*SharedEmailDomain))
            }
            m.SetSharedEmailDomains(res)
        }
        return nil
    }
    return res
}
func (m *Directory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Directory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdministrativeUnits() != nil {
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
    if m.GetAttributeSets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttributeSets()))
        for i, v := range m.GetAttributeSets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("attributeSets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomSecurityAttributeDefinitions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomSecurityAttributeDefinitions()))
        for i, v := range m.GetCustomSecurityAttributeDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customSecurityAttributeDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedItems() != nil {
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
    if m.GetFeatureRolloutPolicies() != nil {
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
    if m.GetFederationConfigurations() != nil {
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
    if m.GetImpactedResources() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetImpactedResources()))
        for i, v := range m.GetImpactedResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("impactedResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInboundSharedUserProfiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInboundSharedUserProfiles()))
        for i, v := range m.GetInboundSharedUserProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("inboundSharedUserProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOutboundSharedUserProfiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOutboundSharedUserProfiles()))
        for i, v := range m.GetOutboundSharedUserProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("outboundSharedUserProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecommendations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecommendations()))
        for i, v := range m.GetRecommendations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("recommendations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSharedEmailDomains() != nil {
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
// SetAdministrativeUnits sets the administrativeUnits property value. Conceptual container for user and group directory objects.
func (m *Directory) SetAdministrativeUnits(value []AdministrativeUnit)() {
    if m != nil {
        m.administrativeUnits = value
    }
}
// SetAttributeSets sets the attributeSets property value. Group of related custom security attribute definitions.
func (m *Directory) SetAttributeSets(value []AttributeSet)() {
    if m != nil {
        m.attributeSets = value
    }
}
// SetCustomSecurityAttributeDefinitions sets the customSecurityAttributeDefinitions property value. Schema of a custom security attributes (key-value pairs).
func (m *Directory) SetCustomSecurityAttributeDefinitions(value []CustomSecurityAttributeDefinition)() {
    if m != nil {
        m.customSecurityAttributeDefinitions = value
    }
}
// SetDeletedItems sets the deletedItems property value. Recently deleted items. Read-only. Nullable.
func (m *Directory) SetDeletedItems(value []DirectoryObject)() {
    if m != nil {
        m.deletedItems = value
    }
}
// SetFeatureRolloutPolicies sets the featureRolloutPolicies property value. Nullable.
func (m *Directory) SetFeatureRolloutPolicies(value []FeatureRolloutPolicy)() {
    if m != nil {
        m.featureRolloutPolicies = value
    }
}
// SetFederationConfigurations sets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) SetFederationConfigurations(value []IdentityProviderBase)() {
    if m != nil {
        m.federationConfigurations = value
    }
}
// SetImpactedResources sets the impactedResources property value. 
func (m *Directory) SetImpactedResources(value []RecommendationResource)() {
    if m != nil {
        m.impactedResources = value
    }
}
// SetInboundSharedUserProfiles sets the inboundSharedUserProfiles property value. 
func (m *Directory) SetInboundSharedUserProfiles(value []InboundSharedUserProfile)() {
    if m != nil {
        m.inboundSharedUserProfiles = value
    }
}
// SetOutboundSharedUserProfiles sets the outboundSharedUserProfiles property value. 
func (m *Directory) SetOutboundSharedUserProfiles(value []OutboundSharedUserProfile)() {
    if m != nil {
        m.outboundSharedUserProfiles = value
    }
}
// SetRecommendations sets the recommendations property value. 
func (m *Directory) SetRecommendations(value []Recommendation)() {
    if m != nil {
        m.recommendations = value
    }
}
// SetSharedEmailDomains sets the sharedEmailDomains property value. 
func (m *Directory) SetSharedEmailDomains(value []SharedEmailDomain)() {
    if m != nil {
        m.sharedEmailDomains = value
    }
}
