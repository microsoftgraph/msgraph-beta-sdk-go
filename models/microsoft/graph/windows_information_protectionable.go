package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsInformationProtectionable 
type WindowsInformationProtectionable interface {
    ManagedAppPolicyable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]TargetedManagedAppPolicyAssignmentable)
    GetAzureRightsManagementServicesAllowed()(*bool)
    GetDataRecoveryCertificate()(WindowsInformationProtectionDataRecoveryCertificateable)
    GetEnforcementLevel()(*WindowsInformationProtectionEnforcementLevel)
    GetEnterpriseDomain()(*string)
    GetEnterpriseInternalProxyServers()([]WindowsInformationProtectionResourceCollectionable)
    GetEnterpriseIPRanges()([]WindowsInformationProtectionIPRangeCollectionable)
    GetEnterpriseIPRangesAreAuthoritative()(*bool)
    GetEnterpriseNetworkDomainNames()([]WindowsInformationProtectionResourceCollectionable)
    GetEnterpriseProtectedDomainNames()([]WindowsInformationProtectionResourceCollectionable)
    GetEnterpriseProxiedDomains()([]WindowsInformationProtectionProxiedDomainCollectionable)
    GetEnterpriseProxyServers()([]WindowsInformationProtectionResourceCollectionable)
    GetEnterpriseProxyServersAreAuthoritative()(*bool)
    GetExemptAppLockerFiles()([]WindowsInformationProtectionAppLockerFileable)
    GetExemptApps()([]WindowsInformationProtectionAppable)
    GetIconsVisible()(*bool)
    GetIndexingEncryptedStoresOrItemsBlocked()(*bool)
    GetIsAssigned()(*bool)
    GetNeutralDomainResources()([]WindowsInformationProtectionResourceCollectionable)
    GetProtectedAppLockerFiles()([]WindowsInformationProtectionAppLockerFileable)
    GetProtectedApps()([]WindowsInformationProtectionAppable)
    GetProtectionUnderLockConfigRequired()(*bool)
    GetRevokeOnUnenrollDisabled()(*bool)
    GetRightsManagementServicesTemplateId()(*string)
    GetSmbAutoEncryptedFileExtensions()([]WindowsInformationProtectionResourceCollectionable)
    SetAssignments(value []TargetedManagedAppPolicyAssignmentable)()
    SetAzureRightsManagementServicesAllowed(value *bool)()
    SetDataRecoveryCertificate(value WindowsInformationProtectionDataRecoveryCertificateable)()
    SetEnforcementLevel(value *WindowsInformationProtectionEnforcementLevel)()
    SetEnterpriseDomain(value *string)()
    SetEnterpriseInternalProxyServers(value []WindowsInformationProtectionResourceCollectionable)()
    SetEnterpriseIPRanges(value []WindowsInformationProtectionIPRangeCollectionable)()
    SetEnterpriseIPRangesAreAuthoritative(value *bool)()
    SetEnterpriseNetworkDomainNames(value []WindowsInformationProtectionResourceCollectionable)()
    SetEnterpriseProtectedDomainNames(value []WindowsInformationProtectionResourceCollectionable)()
    SetEnterpriseProxiedDomains(value []WindowsInformationProtectionProxiedDomainCollectionable)()
    SetEnterpriseProxyServers(value []WindowsInformationProtectionResourceCollectionable)()
    SetEnterpriseProxyServersAreAuthoritative(value *bool)()
    SetExemptAppLockerFiles(value []WindowsInformationProtectionAppLockerFileable)()
    SetExemptApps(value []WindowsInformationProtectionAppable)()
    SetIconsVisible(value *bool)()
    SetIndexingEncryptedStoresOrItemsBlocked(value *bool)()
    SetIsAssigned(value *bool)()
    SetNeutralDomainResources(value []WindowsInformationProtectionResourceCollectionable)()
    SetProtectedAppLockerFiles(value []WindowsInformationProtectionAppLockerFileable)()
    SetProtectedApps(value []WindowsInformationProtectionAppable)()
    SetProtectionUnderLockConfigRequired(value *bool)()
    SetRevokeOnUnenrollDisabled(value *bool)()
    SetRightsManagementServicesTemplateId(value *string)()
    SetSmbAutoEncryptedFileExtensions(value []WindowsInformationProtectionResourceCollectionable)()
}