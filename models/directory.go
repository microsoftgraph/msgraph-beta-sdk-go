package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Directory struct {
    Entity
}
// NewDirectory instantiates a new Directory and sets the default values.
func NewDirectory()(*Directory) {
    m := &Directory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectory(), nil
}
// GetAdministrativeUnits gets the administrativeUnits property value. Conceptual container for user and group directory objects.
// returns a []AdministrativeUnitable when successful
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
// returns a []AttributeSetable when successful
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
// GetAuthenticationMethodDevices gets the authenticationMethodDevices property value. Exposes the hardware OATH method in the directory.
// returns a AuthenticationMethodDeviceable when successful
func (m *Directory) GetAuthenticationMethodDevices()(AuthenticationMethodDeviceable) {
    val, err := m.GetBackingStore().Get("authenticationMethodDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationMethodDeviceable)
    }
    return nil
}
// GetCertificateAuthorities gets the certificateAuthorities property value. The certificateAuthorities property
// returns a CertificateAuthorityPathable when successful
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
// returns a []CustomSecurityAttributeDefinitionable when successful
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
// GetDeletedItems gets the deletedItems property value. Recently deleted items. Read-only. Nullable.
// returns a []DirectoryObjectable when successful
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
// GetDeviceLocalCredentials gets the deviceLocalCredentials property value. The credentials of the device's local administrator account backed up to Microsoft Entra ID.
// returns a []DeviceLocalCredentialInfoable when successful
func (m *Directory) GetDeviceLocalCredentials()([]DeviceLocalCredentialInfoable) {
    val, err := m.GetBackingStore().Get("deviceLocalCredentials")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceLocalCredentialInfoable)
    }
    return nil
}
// GetExternalUserProfiles gets the externalUserProfiles property value. Collection of external user profiles that represent collaborators in the directory.
// returns a []ExternalUserProfileable when successful
func (m *Directory) GetExternalUserProfiles()([]ExternalUserProfileable) {
    val, err := m.GetBackingStore().Get("externalUserProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExternalUserProfileable)
    }
    return nil
}
// GetFeatureRolloutPolicies gets the featureRolloutPolicies property value. The featureRolloutPolicies property
// returns a []FeatureRolloutPolicyable when successful
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
// returns a []IdentityProviderBaseable when successful
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
    res["authenticationMethodDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethodDevices(val.(AuthenticationMethodDeviceable))
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
    res["deviceLocalCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceLocalCredentialInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceLocalCredentialInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceLocalCredentialInfoable)
                }
            }
            m.SetDeviceLocalCredentials(res)
        }
        return nil
    }
    res["externalUserProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalUserProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalUserProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExternalUserProfileable)
                }
            }
            m.SetExternalUserProfiles(res)
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
    res["pendingExternalUserProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePendingExternalUserProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PendingExternalUserProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PendingExternalUserProfileable)
                }
            }
            m.SetPendingExternalUserProfiles(res)
        }
        return nil
    }
    res["publicKeyInfrastructure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicKeyInfrastructureRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicKeyInfrastructure(val.(PublicKeyInfrastructureRootable))
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
    res["templates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplates(val.(Templateable))
        }
        return nil
    }
    return res
}
// GetImpactedResources gets the impactedResources property value. The impactedResources property
// returns a []ImpactedResourceable when successful
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
// GetInboundSharedUserProfiles gets the inboundSharedUserProfiles property value. A collection of external users whose profile data is shared with the Microsoft Entra tenant. Nullable.
// returns a []InboundSharedUserProfileable when successful
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
// GetOnPremisesSynchronization gets the onPremisesSynchronization property value. A container for on-premises directory synchronization functionalities that are available for the organization.
// returns a []OnPremisesDirectorySynchronizationable when successful
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
// returns a []OutboundSharedUserProfileable when successful
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
// GetPendingExternalUserProfiles gets the pendingExternalUserProfiles property value. Collection of pending external user profiles representing collaborators in the directory that are unredeemed.
// returns a []PendingExternalUserProfileable when successful
func (m *Directory) GetPendingExternalUserProfiles()([]PendingExternalUserProfileable) {
    val, err := m.GetBackingStore().Get("pendingExternalUserProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PendingExternalUserProfileable)
    }
    return nil
}
// GetPublicKeyInfrastructure gets the publicKeyInfrastructure property value. The collection of public key infrastructure instances for the certificate-based authentication feature for users in a Microsoft Entra tenant.
// returns a PublicKeyInfrastructureRootable when successful
func (m *Directory) GetPublicKeyInfrastructure()(PublicKeyInfrastructureRootable) {
    val, err := m.GetBackingStore().Get("publicKeyInfrastructure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PublicKeyInfrastructureRootable)
    }
    return nil
}
// GetRecommendations gets the recommendations property value. List of recommended improvements to improve tenant posture.
// returns a []Recommendationable when successful
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
// returns a []SharedEmailDomainable when successful
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
// GetSubscriptions gets the subscriptions property value. List of commercial subscriptions that an organization has.
// returns a []CompanySubscriptionable when successful
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
// GetTemplates gets the templates property value. The templates property
// returns a Templateable when successful
func (m *Directory) GetTemplates()(Templateable) {
    val, err := m.GetBackingStore().Get("templates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Templateable)
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
        err = writer.WriteObjectValue("authenticationMethodDevices", m.GetAuthenticationMethodDevices())
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
    if m.GetDeviceLocalCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceLocalCredentials()))
        for i, v := range m.GetDeviceLocalCredentials() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceLocalCredentials", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExternalUserProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExternalUserProfiles()))
        for i, v := range m.GetExternalUserProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("externalUserProfiles", cast)
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
    if m.GetPendingExternalUserProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPendingExternalUserProfiles()))
        for i, v := range m.GetPendingExternalUserProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("pendingExternalUserProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publicKeyInfrastructure", m.GetPublicKeyInfrastructure())
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
    {
        err = writer.WriteObjectValue("templates", m.GetTemplates())
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
// SetAuthenticationMethodDevices sets the authenticationMethodDevices property value. Exposes the hardware OATH method in the directory.
func (m *Directory) SetAuthenticationMethodDevices(value AuthenticationMethodDeviceable)() {
    err := m.GetBackingStore().Set("authenticationMethodDevices", value)
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
// SetDeletedItems sets the deletedItems property value. Recently deleted items. Read-only. Nullable.
func (m *Directory) SetDeletedItems(value []DirectoryObjectable)() {
    err := m.GetBackingStore().Set("deletedItems", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceLocalCredentials sets the deviceLocalCredentials property value. The credentials of the device's local administrator account backed up to Microsoft Entra ID.
func (m *Directory) SetDeviceLocalCredentials(value []DeviceLocalCredentialInfoable)() {
    err := m.GetBackingStore().Set("deviceLocalCredentials", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalUserProfiles sets the externalUserProfiles property value. Collection of external user profiles that represent collaborators in the directory.
func (m *Directory) SetExternalUserProfiles(value []ExternalUserProfileable)() {
    err := m.GetBackingStore().Set("externalUserProfiles", value)
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
// SetInboundSharedUserProfiles sets the inboundSharedUserProfiles property value. A collection of external users whose profile data is shared with the Microsoft Entra tenant. Nullable.
func (m *Directory) SetInboundSharedUserProfiles(value []InboundSharedUserProfileable)() {
    err := m.GetBackingStore().Set("inboundSharedUserProfiles", value)
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
// SetPendingExternalUserProfiles sets the pendingExternalUserProfiles property value. Collection of pending external user profiles representing collaborators in the directory that are unredeemed.
func (m *Directory) SetPendingExternalUserProfiles(value []PendingExternalUserProfileable)() {
    err := m.GetBackingStore().Set("pendingExternalUserProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetPublicKeyInfrastructure sets the publicKeyInfrastructure property value. The collection of public key infrastructure instances for the certificate-based authentication feature for users in a Microsoft Entra tenant.
func (m *Directory) SetPublicKeyInfrastructure(value PublicKeyInfrastructureRootable)() {
    err := m.GetBackingStore().Set("publicKeyInfrastructure", value)
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
// SetSubscriptions sets the subscriptions property value. List of commercial subscriptions that an organization has.
func (m *Directory) SetSubscriptions(value []CompanySubscriptionable)() {
    err := m.GetBackingStore().Set("subscriptions", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplates sets the templates property value. The templates property
func (m *Directory) SetTemplates(value Templateable)() {
    err := m.GetBackingStore().Set("templates", value)
    if err != nil {
        panic(err)
    }
}
type Directoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdministrativeUnits()([]AdministrativeUnitable)
    GetAttributeSets()([]AttributeSetable)
    GetAuthenticationMethodDevices()(AuthenticationMethodDeviceable)
    GetCertificateAuthorities()(CertificateAuthorityPathable)
    GetCustomSecurityAttributeDefinitions()([]CustomSecurityAttributeDefinitionable)
    GetDeletedItems()([]DirectoryObjectable)
    GetDeviceLocalCredentials()([]DeviceLocalCredentialInfoable)
    GetExternalUserProfiles()([]ExternalUserProfileable)
    GetFeatureRolloutPolicies()([]FeatureRolloutPolicyable)
    GetFederationConfigurations()([]IdentityProviderBaseable)
    GetImpactedResources()([]ImpactedResourceable)
    GetInboundSharedUserProfiles()([]InboundSharedUserProfileable)
    GetOnPremisesSynchronization()([]OnPremisesDirectorySynchronizationable)
    GetOutboundSharedUserProfiles()([]OutboundSharedUserProfileable)
    GetPendingExternalUserProfiles()([]PendingExternalUserProfileable)
    GetPublicKeyInfrastructure()(PublicKeyInfrastructureRootable)
    GetRecommendations()([]Recommendationable)
    GetSharedEmailDomains()([]SharedEmailDomainable)
    GetSubscriptions()([]CompanySubscriptionable)
    GetTemplates()(Templateable)
    SetAdministrativeUnits(value []AdministrativeUnitable)()
    SetAttributeSets(value []AttributeSetable)()
    SetAuthenticationMethodDevices(value AuthenticationMethodDeviceable)()
    SetCertificateAuthorities(value CertificateAuthorityPathable)()
    SetCustomSecurityAttributeDefinitions(value []CustomSecurityAttributeDefinitionable)()
    SetDeletedItems(value []DirectoryObjectable)()
    SetDeviceLocalCredentials(value []DeviceLocalCredentialInfoable)()
    SetExternalUserProfiles(value []ExternalUserProfileable)()
    SetFeatureRolloutPolicies(value []FeatureRolloutPolicyable)()
    SetFederationConfigurations(value []IdentityProviderBaseable)()
    SetImpactedResources(value []ImpactedResourceable)()
    SetInboundSharedUserProfiles(value []InboundSharedUserProfileable)()
    SetOnPremisesSynchronization(value []OnPremisesDirectorySynchronizationable)()
    SetOutboundSharedUserProfiles(value []OutboundSharedUserProfileable)()
    SetPendingExternalUserProfiles(value []PendingExternalUserProfileable)()
    SetPublicKeyInfrastructure(value PublicKeyInfrastructureRootable)()
    SetRecommendations(value []Recommendationable)()
    SetSharedEmailDomains(value []SharedEmailDomainable)()
    SetSubscriptions(value []CompanySubscriptionable)()
    SetTemplates(value Templateable)()
}
