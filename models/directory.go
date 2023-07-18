package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Directory 
type Directory struct {
    Entity
}
// NewDirectory instantiates a new directory and sets the default values.
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
    val, err := m.GetBackingStore().Get("administrativeUnits")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AdministrativeUnitable)
    }
    return nil
}
// GetAttributeSets gets the attributeSets property value. Group of related custom security attribute definitions.
func (m *Directory) GetAttributeSets()([]AttributeSetable) {
    val, err := m.GetBackingStore().Get("attributeSets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AttributeSetable)
    }
    return nil
}
// GetCertificateAuthorities gets the certificateAuthorities property value. The certificateAuthorities property
func (m *Directory) GetCertificateAuthorities()(CertificateAuthorityPathable) {
    val, err := m.GetBackingStore().Get("certificateAuthorities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CertificateAuthorityPathable)
    }
    return nil
}
// GetCustomSecurityAttributeDefinitions gets the customSecurityAttributeDefinitions property value. Schema of a custom security attributes (key-value pairs).
func (m *Directory) GetCustomSecurityAttributeDefinitions()([]CustomSecurityAttributeDefinitionable) {
    val, err := m.GetBackingStore().Get("customSecurityAttributeDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomSecurityAttributeDefinitionable)
    }
    return nil
}
// GetDeletedItems gets the deletedItems property value. The deletedItems property
func (m *Directory) GetDeletedItems()([]DirectoryObjectable) {
    val, err := m.GetBackingStore().Get("deletedItems")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DirectoryObjectable)
    }
    return nil
}
// GetFeatureRolloutPolicies gets the featureRolloutPolicies property value. The featureRolloutPolicies property
func (m *Directory) GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable) {
    val, err := m.GetBackingStore().Get("featureRolloutPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]FeatureRolloutPolicyable)
    }
    return nil
}
// GetFederationConfigurations gets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) GetFederationConfigurations()([]IdentityProviderBaseable) {
    val, err := m.GetBackingStore().Get("federationConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IdentityProviderBaseable)
    }
    return nil
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
                if v != nil {
                    res[i] = v.(AdministrativeUnitable)
                }
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
                if v != nil {
                    res[i] = v.(AttributeSetable)
                }
            }
            m.SetAttributeSets(res)
        }
        return nil
    }
    res["certificateAuthorities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificateAuthorityPathFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateAuthorities(val.(CertificateAuthorityPathable))
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
                if v != nil {
                    res[i] = v.(CustomSecurityAttributeDefinitionable)
                }
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
                if v != nil {
                    res[i] = v.(DirectoryObjectable)
                }
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
                if v != nil {
                    res[i] = v.(FeatureRolloutPolicyable)
                }
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
                if v != nil {
                    res[i] = v.(IdentityProviderBaseable)
                }
            }
            m.SetFederationConfigurations(res)
        }
        return nil
    }
    res["impactedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImpactedResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImpactedResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ImpactedResourceable)
                }
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
                if v != nil {
                    res[i] = v.(InboundSharedUserProfileable)
                }
            }
            m.SetInboundSharedUserProfiles(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["onPremisesSynchronization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesDirectorySynchronizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesDirectorySynchronizationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OnPremisesDirectorySynchronizationable)
                }
            }
            m.SetOnPremisesSynchronization(res)
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
                if v != nil {
                    res[i] = v.(OutboundSharedUserProfileable)
                }
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
                if v != nil {
                    res[i] = v.(Recommendationable)
                }
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
                if v != nil {
                    res[i] = v.(SharedEmailDomainable)
                }
            }
            m.SetSharedEmailDomains(res)
        }
        return nil
    }
    res["subscriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCompanySubscriptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CompanySubscriptionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CompanySubscriptionable)
                }
            }
            m.SetSubscriptions(res)
        }
        return nil
    }
    return res
}
// GetImpactedResources gets the impactedResources property value. The impactedResources property
func (m *Directory) GetImpactedResources()([]ImpactedResourceable) {
    val, err := m.GetBackingStore().Get("impactedResources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ImpactedResourceable)
    }
    return nil
}
// GetInboundSharedUserProfiles gets the inboundSharedUserProfiles property value. A collection of external Azure AD users whose profile data has been shared with the Azure AD tenant. Nullable.
func (m *Directory) GetInboundSharedUserProfiles()([]InboundSharedUserProfileable) {
    val, err := m.GetBackingStore().Get("inboundSharedUserProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]InboundSharedUserProfileable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Directory) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnPremisesSynchronization gets the onPremisesSynchronization property value. A container for on-premises directory synchronization functionalities that are available for the organization.
func (m *Directory) GetOnPremisesSynchronization()([]OnPremisesDirectorySynchronizationable) {
    val, err := m.GetBackingStore().Get("onPremisesSynchronization")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OnPremisesDirectorySynchronizationable)
    }
    return nil
}
// GetOutboundSharedUserProfiles gets the outboundSharedUserProfiles property value. The outboundSharedUserProfiles property
func (m *Directory) GetOutboundSharedUserProfiles()([]OutboundSharedUserProfileable) {
    val, err := m.GetBackingStore().Get("outboundSharedUserProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OutboundSharedUserProfileable)
    }
    return nil
}
// GetRecommendations gets the recommendations property value. List of recommended improvements to improve tenant posture.
func (m *Directory) GetRecommendations()([]Recommendationable) {
    val, err := m.GetBackingStore().Get("recommendations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Recommendationable)
    }
    return nil
}
// GetSharedEmailDomains gets the sharedEmailDomains property value. The sharedEmailDomains property
func (m *Directory) GetSharedEmailDomains()([]SharedEmailDomainable) {
    val, err := m.GetBackingStore().Get("sharedEmailDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SharedEmailDomainable)
    }
    return nil
}
// GetSubscriptions gets the subscriptions property value. The subscriptions property
func (m *Directory) GetSubscriptions()([]CompanySubscriptionable) {
    val, err := m.GetBackingStore().Get("subscriptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CompanySubscriptionable)
    }
    return nil
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("administrativeUnits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAttributeSets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttributeSets()))
        for i, v := range m.GetAttributeSets() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("attributeSets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("certificateAuthorities", m.GetCertificateAuthorities())
        if err != nil {
            return err
        }
    }
    if m.GetCustomSecurityAttributeDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomSecurityAttributeDefinitions()))
        for i, v := range m.GetCustomSecurityAttributeDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("customSecurityAttributeDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedItems()))
        for i, v := range m.GetDeletedItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deletedItems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFeatureRolloutPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFeatureRolloutPolicies()))
        for i, v := range m.GetFeatureRolloutPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("featureRolloutPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFederationConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFederationConfigurations()))
        for i, v := range m.GetFederationConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("federationConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImpactedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImpactedResources()))
        for i, v := range m.GetImpactedResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("impactedResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInboundSharedUserProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInboundSharedUserProfiles()))
        for i, v := range m.GetInboundSharedUserProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("inboundSharedUserProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOnPremisesSynchronization() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnPremisesSynchronization()))
        for i, v := range m.GetOnPremisesSynchronization() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("onPremisesSynchronization", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOutboundSharedUserProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOutboundSharedUserProfiles()))
        for i, v := range m.GetOutboundSharedUserProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("outboundSharedUserProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecommendations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecommendations()))
        for i, v := range m.GetRecommendations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("recommendations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSharedEmailDomains() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharedEmailDomains()))
        for i, v := range m.GetSharedEmailDomains() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sharedEmailDomains", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubscriptions()))
        for i, v := range m.GetSubscriptions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("subscriptions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdministrativeUnits sets the administrativeUnits property value. Conceptual container for user and group directory objects.
func (m *Directory) SetAdministrativeUnits(value []AdministrativeUnitable)() {
    err := m.GetBackingStore().Set("administrativeUnits", value)
    if err != nil {
        panic(err)
    }
}
// SetAttributeSets sets the attributeSets property value. Group of related custom security attribute definitions.
func (m *Directory) SetAttributeSets(value []AttributeSetable)() {
    err := m.GetBackingStore().Set("attributeSets", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateAuthorities sets the certificateAuthorities property value. The certificateAuthorities property
func (m *Directory) SetCertificateAuthorities(value CertificateAuthorityPathable)() {
    err := m.GetBackingStore().Set("certificateAuthorities", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomSecurityAttributeDefinitions sets the customSecurityAttributeDefinitions property value. Schema of a custom security attributes (key-value pairs).
func (m *Directory) SetCustomSecurityAttributeDefinitions(value []CustomSecurityAttributeDefinitionable)() {
    err := m.GetBackingStore().Set("customSecurityAttributeDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetDeletedItems sets the deletedItems property value. The deletedItems property
func (m *Directory) SetDeletedItems(value []DirectoryObjectable)() {
    err := m.GetBackingStore().Set("deletedItems", value)
    if err != nil {
        panic(err)
    }
}
// SetFeatureRolloutPolicies sets the featureRolloutPolicies property value. The featureRolloutPolicies property
func (m *Directory) SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)() {
    err := m.GetBackingStore().Set("featureRolloutPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetFederationConfigurations sets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) SetFederationConfigurations(value []IdentityProviderBaseable)() {
    err := m.GetBackingStore().Set("federationConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetImpactedResources sets the impactedResources property value. The impactedResources property
func (m *Directory) SetImpactedResources(value []ImpactedResourceable)() {
    err := m.GetBackingStore().Set("impactedResources", value)
    if err != nil {
        panic(err)
    }
}
// SetInboundSharedUserProfiles sets the inboundSharedUserProfiles property value. A collection of external Azure AD users whose profile data has been shared with the Azure AD tenant. Nullable.
func (m *Directory) SetInboundSharedUserProfiles(value []InboundSharedUserProfileable)() {
    err := m.GetBackingStore().Set("inboundSharedUserProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Directory) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremisesSynchronization sets the onPremisesSynchronization property value. A container for on-premises directory synchronization functionalities that are available for the organization.
func (m *Directory) SetOnPremisesSynchronization(value []OnPremisesDirectorySynchronizationable)() {
    err := m.GetBackingStore().Set("onPremisesSynchronization", value)
    if err != nil {
        panic(err)
    }
}
// SetOutboundSharedUserProfiles sets the outboundSharedUserProfiles property value. The outboundSharedUserProfiles property
func (m *Directory) SetOutboundSharedUserProfiles(value []OutboundSharedUserProfileable)() {
    err := m.GetBackingStore().Set("outboundSharedUserProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetRecommendations sets the recommendations property value. List of recommended improvements to improve tenant posture.
func (m *Directory) SetRecommendations(value []Recommendationable)() {
    err := m.GetBackingStore().Set("recommendations", value)
    if err != nil {
        panic(err)
    }
}
// SetSharedEmailDomains sets the sharedEmailDomains property value. The sharedEmailDomains property
func (m *Directory) SetSharedEmailDomains(value []SharedEmailDomainable)() {
    err := m.GetBackingStore().Set("sharedEmailDomains", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriptions sets the subscriptions property value. The subscriptions property
func (m *Directory) SetSubscriptions(value []CompanySubscriptionable)() {
    err := m.GetBackingStore().Set("subscriptions", value)
    if err != nil {
        panic(err)
    }
}
// Directoryable 
type Directoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdministrativeUnits()([]AdministrativeUnitable)
    GetAttributeSets()([]AttributeSetable)
    GetCertificateAuthorities()(CertificateAuthorityPathable)
    GetCustomSecurityAttributeDefinitions()([]CustomSecurityAttributeDefinitionable)
    GetDeletedItems()([]DirectoryObjectable)
    GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable)
    GetFederationConfigurations()([]IdentityProviderBaseable)
    GetImpactedResources()([]ImpactedResourceable)
    GetInboundSharedUserProfiles()([]InboundSharedUserProfileable)
    GetOdataType()(*string)
    GetOnPremisesSynchronization()([]OnPremisesDirectorySynchronizationable)
    GetOutboundSharedUserProfiles()([]OutboundSharedUserProfileable)
    GetRecommendations()([]Recommendationable)
    GetSharedEmailDomains()([]SharedEmailDomainable)
    GetSubscriptions()([]CompanySubscriptionable)
    SetAdministrativeUnits(value []AdministrativeUnitable)()
    SetAttributeSets(value []AttributeSetable)()
    SetCertificateAuthorities(value CertificateAuthorityPathable)()
    SetCustomSecurityAttributeDefinitions(value []CustomSecurityAttributeDefinitionable)()
    SetDeletedItems(value []DirectoryObjectable)()
    SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)()
    SetFederationConfigurations(value []IdentityProviderBaseable)()
    SetImpactedResources(value []ImpactedResourceable)()
    SetInboundSharedUserProfiles(value []InboundSharedUserProfileable)()
    SetOdataType(value *string)()
    SetOnPremisesSynchronization(value []OnPremisesDirectorySynchronizationable)()
    SetOutboundSharedUserProfiles(value []OutboundSharedUserProfileable)()
    SetRecommendations(value []Recommendationable)()
    SetSharedEmailDomains(value []SharedEmailDomainable)()
    SetSubscriptions(value []CompanySubscriptionable)()
}
