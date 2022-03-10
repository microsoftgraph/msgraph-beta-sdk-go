package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Domainable 
type Domainable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAuthenticationType()(*string)
    GetAvailabilityStatus()(*string)
    GetDomainNameReferences()([]DirectoryObjectable)
    GetFederationConfiguration()([]InternalDomainFederationable)
    GetIsAdminManaged()(*bool)
    GetIsDefault()(*bool)
    GetIsInitial()(*bool)
    GetIsRoot()(*bool)
    GetIsVerified()(*bool)
    GetPasswordNotificationWindowInDays()(*int32)
    GetPasswordValidityPeriodInDays()(*int32)
    GetServiceConfigurationRecords()([]DomainDnsRecordable)
    GetSharedEmailDomainInvitations()([]SharedEmailDomainInvitationable)
    GetState()(DomainStateable)
    GetSupportedServices()([]string)
    GetVerificationDnsRecords()([]DomainDnsRecordable)
    SetAuthenticationType(value *string)()
    SetAvailabilityStatus(value *string)()
    SetDomainNameReferences(value []DirectoryObjectable)()
    SetFederationConfiguration(value []InternalDomainFederationable)()
    SetIsAdminManaged(value *bool)()
    SetIsDefault(value *bool)()
    SetIsInitial(value *bool)()
    SetIsRoot(value *bool)()
    SetIsVerified(value *bool)()
    SetPasswordNotificationWindowInDays(value *int32)()
    SetPasswordValidityPeriodInDays(value *int32)()
    SetServiceConfigurationRecords(value []DomainDnsRecordable)()
    SetSharedEmailDomainInvitations(value []SharedEmailDomainInvitationable)()
    SetState(value DomainStateable)()
    SetSupportedServices(value []string)()
    SetVerificationDnsRecords(value []DomainDnsRecordable)()
}
