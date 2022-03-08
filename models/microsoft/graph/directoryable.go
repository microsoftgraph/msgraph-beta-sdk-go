package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Directoryable 
type Directoryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdministrativeUnits()([]AdministrativeUnitable)
    GetAttributeSets()([]AttributeSetable)
    GetCustomSecurityAttributeDefinitions()([]CustomSecurityAttributeDefinitionable)
    GetDeletedItems()([]DirectoryObjectable)
    GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable)
    GetFederationConfigurations()([]IdentityProviderBaseable)
    GetImpactedResources()([]RecommendationResourceable)
    GetInboundSharedUserProfiles()([]InboundSharedUserProfileable)
    GetOutboundSharedUserProfiles()([]OutboundSharedUserProfileable)
    GetRecommendations()([]Recommendationable)
    GetSharedEmailDomains()([]SharedEmailDomainable)
    SetAdministrativeUnits(value []AdministrativeUnitable)()
    SetAttributeSets(value []AttributeSetable)()
    SetCustomSecurityAttributeDefinitions(value []CustomSecurityAttributeDefinitionable)()
    SetDeletedItems(value []DirectoryObjectable)()
    SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)()
    SetFederationConfigurations(value []IdentityProviderBaseable)()
    SetImpactedResources(value []RecommendationResourceable)()
    SetInboundSharedUserProfiles(value []InboundSharedUserProfileable)()
    SetOutboundSharedUserProfiles(value []OutboundSharedUserProfileable)()
    SetRecommendations(value []Recommendationable)()
    SetSharedEmailDomains(value []SharedEmailDomainable)()
}
