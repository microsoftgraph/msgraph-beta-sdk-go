package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
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
    // The deletedItems property
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
    odataTypeValue := "#microsoft.graph.directory";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectory(), nil
}
// GetAdministrativeUnits gets the administrativeUnits property value. Conceptual container for user and group directory objects.
func (m *Directory) GetAdministrativeUnits()([]AdministrativeUnitable) {
    return m.administrativeUnits
}
// GetAttributeSets gets the attributeSets property value. Group of related custom security attribute definitions.
func (m *Directory) GetAttributeSets()([]AttributeSetable) {
    return m.attributeSets
}
// GetCustomSecurityAttributeDefinitions gets the customSecurityAttributeDefinitions property value. Schema of a custom security attributes (key-value pairs).
func (m *Directory) GetCustomSecurityAttributeDefinitions()([]CustomSecurityAttributeDefinitionable) {
    return m.customSecurityAttributeDefinitions
}
// GetDeletedItems gets the deletedItems property value. The deletedItems property
func (m *Directory) GetDeletedItems()([]DirectoryObjectable) {
    return m.deletedItems
}
// GetFeatureRolloutPolicies gets the featureRolloutPolicies property value. The featureRolloutPolicies property
func (m *Directory) GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable) {
    return m.featureRolloutPolicies
}
// GetFederationConfigurations gets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) GetFederationConfigurations()([]IdentityProviderBaseable) {
    return m.federationConfigurations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Directory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnits"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAdministrativeUnitFromDiscriminatorValue , m.SetAdministrativeUnits)
    res["attributeSets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAttributeSetFromDiscriminatorValue , m.SetAttributeSets)
    res["customSecurityAttributeDefinitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomSecurityAttributeDefinitionFromDiscriminatorValue , m.SetCustomSecurityAttributeDefinitions)
    res["deletedItems"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue , m.SetDeletedItems)
    res["featureRolloutPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateFeatureRolloutPolicyFromDiscriminatorValue , m.SetFeatureRolloutPolicies)
    res["federationConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIdentityProviderBaseFromDiscriminatorValue , m.SetFederationConfigurations)
    res["impactedResources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRecommendationResourceFromDiscriminatorValue , m.SetImpactedResources)
    res["inboundSharedUserProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateInboundSharedUserProfileFromDiscriminatorValue , m.SetInboundSharedUserProfiles)
    res["outboundSharedUserProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOutboundSharedUserProfileFromDiscriminatorValue , m.SetOutboundSharedUserProfiles)
    res["recommendations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRecommendationFromDiscriminatorValue , m.SetRecommendations)
    res["sharedEmailDomains"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSharedEmailDomainFromDiscriminatorValue , m.SetSharedEmailDomains)
    return res
}
// GetImpactedResources gets the impactedResources property value. The impactedResources property
func (m *Directory) GetImpactedResources()([]RecommendationResourceable) {
    return m.impactedResources
}
// GetInboundSharedUserProfiles gets the inboundSharedUserProfiles property value. The inboundSharedUserProfiles property
func (m *Directory) GetInboundSharedUserProfiles()([]InboundSharedUserProfileable) {
    return m.inboundSharedUserProfiles
}
// GetOutboundSharedUserProfiles gets the outboundSharedUserProfiles property value. The outboundSharedUserProfiles property
func (m *Directory) GetOutboundSharedUserProfiles()([]OutboundSharedUserProfileable) {
    return m.outboundSharedUserProfiles
}
// GetRecommendations gets the recommendations property value. The recommendations property
func (m *Directory) GetRecommendations()([]Recommendationable) {
    return m.recommendations
}
// GetSharedEmailDomains gets the sharedEmailDomains property value. The sharedEmailDomains property
func (m *Directory) GetSharedEmailDomains()([]SharedEmailDomainable) {
    return m.sharedEmailDomains
}
// Serialize serializes information the current object
func (m *Directory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdministrativeUnits() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAdministrativeUnits())
        err = writer.WriteCollectionOfObjectValues("administrativeUnits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAttributeSets() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAttributeSets())
        err = writer.WriteCollectionOfObjectValues("attributeSets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomSecurityAttributeDefinitions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomSecurityAttributeDefinitions())
        err = writer.WriteCollectionOfObjectValues("customSecurityAttributeDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedItems() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeletedItems())
        err = writer.WriteCollectionOfObjectValues("deletedItems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFeatureRolloutPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetFeatureRolloutPolicies())
        err = writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFederationConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetFederationConfigurations())
        err = writer.WriteCollectionOfObjectValues("federationConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImpactedResources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetImpactedResources())
        err = writer.WriteCollectionOfObjectValues("impactedResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInboundSharedUserProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetInboundSharedUserProfiles())
        err = writer.WriteCollectionOfObjectValues("inboundSharedUserProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOutboundSharedUserProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOutboundSharedUserProfiles())
        err = writer.WriteCollectionOfObjectValues("outboundSharedUserProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecommendations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRecommendations())
        err = writer.WriteCollectionOfObjectValues("recommendations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSharedEmailDomains() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSharedEmailDomains())
        err = writer.WriteCollectionOfObjectValues("sharedEmailDomains", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdministrativeUnits sets the administrativeUnits property value. Conceptual container for user and group directory objects.
func (m *Directory) SetAdministrativeUnits(value []AdministrativeUnitable)() {
    m.administrativeUnits = value
}
// SetAttributeSets sets the attributeSets property value. Group of related custom security attribute definitions.
func (m *Directory) SetAttributeSets(value []AttributeSetable)() {
    m.attributeSets = value
}
// SetCustomSecurityAttributeDefinitions sets the customSecurityAttributeDefinitions property value. Schema of a custom security attributes (key-value pairs).
func (m *Directory) SetCustomSecurityAttributeDefinitions(value []CustomSecurityAttributeDefinitionable)() {
    m.customSecurityAttributeDefinitions = value
}
// SetDeletedItems sets the deletedItems property value. The deletedItems property
func (m *Directory) SetDeletedItems(value []DirectoryObjectable)() {
    m.deletedItems = value
}
// SetFeatureRolloutPolicies sets the featureRolloutPolicies property value. The featureRolloutPolicies property
func (m *Directory) SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)() {
    m.featureRolloutPolicies = value
}
// SetFederationConfigurations sets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) SetFederationConfigurations(value []IdentityProviderBaseable)() {
    m.federationConfigurations = value
}
// SetImpactedResources sets the impactedResources property value. The impactedResources property
func (m *Directory) SetImpactedResources(value []RecommendationResourceable)() {
    m.impactedResources = value
}
// SetInboundSharedUserProfiles sets the inboundSharedUserProfiles property value. The inboundSharedUserProfiles property
func (m *Directory) SetInboundSharedUserProfiles(value []InboundSharedUserProfileable)() {
    m.inboundSharedUserProfiles = value
}
// SetOutboundSharedUserProfiles sets the outboundSharedUserProfiles property value. The outboundSharedUserProfiles property
func (m *Directory) SetOutboundSharedUserProfiles(value []OutboundSharedUserProfileable)() {
    m.outboundSharedUserProfiles = value
}
// SetRecommendations sets the recommendations property value. The recommendations property
func (m *Directory) SetRecommendations(value []Recommendationable)() {
    m.recommendations = value
}
// SetSharedEmailDomains sets the sharedEmailDomains property value. The sharedEmailDomains property
func (m *Directory) SetSharedEmailDomains(value []SharedEmailDomainable)() {
    m.sharedEmailDomains = value
}
