package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Directory 
type Directory struct {
    Entity
    // Conceptual container for user and group directory objects.
    administrativeUnits []AdministrativeUnitable
    // Group of related custom security attribute definitions.
    attributeSets []AttributeSetable
    // Schema of a custom security attributes (key-value pairs).
    customSecurityAttributeDefinitions []CustomSecurityAttributeDefinitionable
    // Recently deleted items. Read-only. Nullable.
    deletedItems []DirectoryObjectable
    // The featureRolloutPolicies property
    featureRolloutPolicies []FeatureRolloutPolicyable
    // Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
    federationConfigurations []IdentityProviderBaseable
    // The impactedResources property
    impactedResources []RecommendationResourceable
    // The inboundSharedUserProfiles property
    inboundSharedUserProfiles []InboundSharedUserProfileable
    // The outboundSharedUserProfiles property
    outboundSharedUserProfiles []OutboundSharedUserProfileable
    // The recommendations property
    recommendations []Recommendationable
    // The sharedEmailDomains property
    sharedEmailDomains []SharedEmailDomainable
}
// NewDirectory instantiates a new Directory and sets the default values.
func NewDirectory()(*Directory) {
    m := &Directory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectory(), nil
}
// GetAdministrativeUnits gets the administrativeUnits property value. Conceptual container for user and group directory objects.
func (m *Directory) GetAdministrativeUnits()([]AdministrativeUnitable) {
    if m == nil {
        return nil
    } else {
        return m.administrativeUnits
    }
}
// GetAttributeSets gets the attributeSets property value. Group of related custom security attribute definitions.
func (m *Directory) GetAttributeSets()([]AttributeSetable) {
    if m == nil {
        return nil
    } else {
        return m.attributeSets
    }
}
// GetCustomSecurityAttributeDefinitions gets the customSecurityAttributeDefinitions property value. Schema of a custom security attributes (key-value pairs).
func (m *Directory) GetCustomSecurityAttributeDefinitions()([]CustomSecurityAttributeDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.customSecurityAttributeDefinitions
    }
}
// GetDeletedItems gets the deletedItems property value. Recently deleted items. Read-only. Nullable.
func (m *Directory) GetDeletedItems()([]DirectoryObjectable) {
    if m == nil {
        return nil
    } else {
        return m.deletedItems
    }
}
// GetFeatureRolloutPolicies gets the featureRolloutPolicies property value. The featureRolloutPolicies property
func (m *Directory) GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.featureRolloutPolicies
    }
}
// GetFederationConfigurations gets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) GetFederationConfigurations()([]IdentityProviderBaseable) {
    if m == nil {
        return nil
    } else {
        return m.federationConfigurations
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Directory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAdministrativeUnitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AdministrativeUnitable, len(val))
            for i, v := range val {
                res[i] = v.(AdministrativeUnitable)
            }
            m.SetAdministrativeUnits(res)
        }
        return nil
    }
    res["attributeSets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttributeSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeSetable, len(val))
            for i, v := range val {
                res[i] = v.(AttributeSetable)
            }
            m.SetAttributeSets(res)
        }
        return nil
    }
    res["customSecurityAttributeDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomSecurityAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomSecurityAttributeDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(CustomSecurityAttributeDefinitionable)
            }
            m.SetCustomSecurityAttributeDefinitions(res)
        }
        return nil
    }
    res["deletedItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetDeletedItems(res)
        }
        return nil
    }
    res["featureRolloutPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFeatureRolloutPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FeatureRolloutPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(FeatureRolloutPolicyable)
            }
            m.SetFeatureRolloutPolicies(res)
        }
        return nil
    }
    res["federationConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityProviderBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBaseable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityProviderBaseable)
            }
            m.SetFederationConfigurations(res)
        }
        return nil
    }
    res["impactedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecommendationResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RecommendationResourceable, len(val))
            for i, v := range val {
                res[i] = v.(RecommendationResourceable)
            }
            m.SetImpactedResources(res)
        }
        return nil
    }
    res["inboundSharedUserProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInboundSharedUserProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InboundSharedUserProfileable, len(val))
            for i, v := range val {
                res[i] = v.(InboundSharedUserProfileable)
            }
            m.SetInboundSharedUserProfiles(res)
        }
        return nil
    }
    res["outboundSharedUserProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOutboundSharedUserProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OutboundSharedUserProfileable, len(val))
            for i, v := range val {
                res[i] = v.(OutboundSharedUserProfileable)
            }
            m.SetOutboundSharedUserProfiles(res)
        }
        return nil
    }
    res["recommendations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecommendationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recommendationable, len(val))
            for i, v := range val {
                res[i] = v.(Recommendationable)
            }
            m.SetRecommendations(res)
        }
        return nil
    }
    res["sharedEmailDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSharedEmailDomainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharedEmailDomainable, len(val))
            for i, v := range val {
                res[i] = v.(SharedEmailDomainable)
            }
            m.SetSharedEmailDomains(res)
        }
        return nil
    }
    return res
}
// GetImpactedResources gets the impactedResources property value. The impactedResources property
func (m *Directory) GetImpactedResources()([]RecommendationResourceable) {
    if m == nil {
        return nil
    } else {
        return m.impactedResources
    }
}
// GetInboundSharedUserProfiles gets the inboundSharedUserProfiles property value. The inboundSharedUserProfiles property
func (m *Directory) GetInboundSharedUserProfiles()([]InboundSharedUserProfileable) {
    if m == nil {
        return nil
    } else {
        return m.inboundSharedUserProfiles
    }
}
// GetOutboundSharedUserProfiles gets the outboundSharedUserProfiles property value. The outboundSharedUserProfiles property
func (m *Directory) GetOutboundSharedUserProfiles()([]OutboundSharedUserProfileable) {
    if m == nil {
        return nil
    } else {
        return m.outboundSharedUserProfiles
    }
}
// GetRecommendations gets the recommendations property value. The recommendations property
func (m *Directory) GetRecommendations()([]Recommendationable) {
    if m == nil {
        return nil
    } else {
        return m.recommendations
    }
}
// GetSharedEmailDomains gets the sharedEmailDomains property value. The sharedEmailDomains property
func (m *Directory) GetSharedEmailDomains()([]SharedEmailDomainable) {
    if m == nil {
        return nil
    } else {
        return m.sharedEmailDomains
    }
}
// Serialize serializes information the current object
func (m *Directory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdministrativeUnits() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdministrativeUnits()))
        for i, v := range m.GetAdministrativeUnits() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("administrativeUnits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAttributeSets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttributeSets()))
        for i, v := range m.GetAttributeSets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("attributeSets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomSecurityAttributeDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomSecurityAttributeDefinitions()))
        for i, v := range m.GetCustomSecurityAttributeDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customSecurityAttributeDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedItems()))
        for i, v := range m.GetDeletedItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deletedItems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFeatureRolloutPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFeatureRolloutPolicies()))
        for i, v := range m.GetFeatureRolloutPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFederationConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFederationConfigurations()))
        for i, v := range m.GetFederationConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("federationConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImpactedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImpactedResources()))
        for i, v := range m.GetImpactedResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("impactedResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInboundSharedUserProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInboundSharedUserProfiles()))
        for i, v := range m.GetInboundSharedUserProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("inboundSharedUserProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOutboundSharedUserProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOutboundSharedUserProfiles()))
        for i, v := range m.GetOutboundSharedUserProfiles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("outboundSharedUserProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecommendations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecommendations()))
        for i, v := range m.GetRecommendations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("recommendations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSharedEmailDomains() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharedEmailDomains()))
        for i, v := range m.GetSharedEmailDomains() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sharedEmailDomains", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdministrativeUnits sets the administrativeUnits property value. Conceptual container for user and group directory objects.
func (m *Directory) SetAdministrativeUnits(value []AdministrativeUnitable)() {
    if m != nil {
        m.administrativeUnits = value
    }
}
// SetAttributeSets sets the attributeSets property value. Group of related custom security attribute definitions.
func (m *Directory) SetAttributeSets(value []AttributeSetable)() {
    if m != nil {
        m.attributeSets = value
    }
}
// SetCustomSecurityAttributeDefinitions sets the customSecurityAttributeDefinitions property value. Schema of a custom security attributes (key-value pairs).
func (m *Directory) SetCustomSecurityAttributeDefinitions(value []CustomSecurityAttributeDefinitionable)() {
    if m != nil {
        m.customSecurityAttributeDefinitions = value
    }
}
// SetDeletedItems sets the deletedItems property value. Recently deleted items. Read-only. Nullable.
func (m *Directory) SetDeletedItems(value []DirectoryObjectable)() {
    if m != nil {
        m.deletedItems = value
    }
}
// SetFeatureRolloutPolicies sets the featureRolloutPolicies property value. The featureRolloutPolicies property
func (m *Directory) SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)() {
    if m != nil {
        m.featureRolloutPolicies = value
    }
}
// SetFederationConfigurations sets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) SetFederationConfigurations(value []IdentityProviderBaseable)() {
    if m != nil {
        m.federationConfigurations = value
    }
}
// SetImpactedResources sets the impactedResources property value. The impactedResources property
func (m *Directory) SetImpactedResources(value []RecommendationResourceable)() {
    if m != nil {
        m.impactedResources = value
    }
}
// SetInboundSharedUserProfiles sets the inboundSharedUserProfiles property value. The inboundSharedUserProfiles property
func (m *Directory) SetInboundSharedUserProfiles(value []InboundSharedUserProfileable)() {
    if m != nil {
        m.inboundSharedUserProfiles = value
    }
}
// SetOutboundSharedUserProfiles sets the outboundSharedUserProfiles property value. The outboundSharedUserProfiles property
func (m *Directory) SetOutboundSharedUserProfiles(value []OutboundSharedUserProfileable)() {
    if m != nil {
        m.outboundSharedUserProfiles = value
    }
}
// SetRecommendations sets the recommendations property value. The recommendations property
func (m *Directory) SetRecommendations(value []Recommendationable)() {
    if m != nil {
        m.recommendations = value
    }
}
// SetSharedEmailDomains sets the sharedEmailDomains property value. The sharedEmailDomains property
func (m *Directory) SetSharedEmailDomains(value []SharedEmailDomainable)() {
    if m != nil {
        m.sharedEmailDomains = value
    }
}
