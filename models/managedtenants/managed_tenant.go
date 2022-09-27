package managedtenants

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedTenant 
type ManagedTenant struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Aggregate view of device compliance policies across managed tenants.
    aggregatedPolicyCompliances []AggregatedPolicyComplianceable
    // The collection of audit events across managed tenants.
    auditEvents []AuditEventable
    // The collection of cloud PC connections across managed tenants.
    cloudPcConnections []CloudPcConnectionable
    // The collection of cloud PC devices across managed tenants.
    cloudPcDevices []CloudPcDeviceable
    // Overview of cloud PC information across managed tenants.
    cloudPcsOverview []CloudPcOverviewable
    // Aggregate view of conditional access policy coverage across managed tenants.
    conditionalAccessPolicyCoverages []ConditionalAccessPolicyCoverageable
    // Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
    credentialUserRegistrationsSummaries []CredentialUserRegistrationsSummaryable
    // Summary information for device compliance policy setting states across managed tenants.
    deviceCompliancePolicySettingStateSummaries []DeviceCompliancePolicySettingStateSummaryable
    // The collection of compliance for managed devices across managed tenants.
    managedDeviceCompliances []ManagedDeviceComplianceable
    // Trend insights for device compliance across managed tenants.
    managedDeviceComplianceTrends []ManagedDeviceComplianceTrendable
    // The collection of baseline management actions across managed tenants.
    managementActions []ManagementActionable
    // The tenant level status of management actions across managed tenants.
    managementActionTenantDeploymentStatuses []ManagementActionTenantDeploymentStatusable
    // The collection of baseline management intents across managed tenants.
    managementIntents []ManagementIntentable
    // The managementTemplateCollections property
    managementTemplateCollections []ManagementTemplateCollectionable
    // The collection of baseline management templates across managed tenants.
    managementTemplates []ManagementTemplateable
    // The managementTemplateSteps property
    managementTemplateSteps []ManagementTemplateStepable
    // The managementTemplateStepVersions property
    managementTemplateStepVersions []ManagementTemplateStepVersionable
    // The collection of role assignments to a signed-in user for a managed tenant.
    myRoles []MyRoleable
    // The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
    tenantGroups []TenantGroupable
    // The collection of tenants associated with the managing entity.
    tenants []Tenantable
    // The collection of tenant level customized information across managed tenants.
    tenantsCustomizedInformation []TenantCustomizedInformationable
    // The collection tenant level detailed information across managed tenants.
    tenantsDetailedInformation []TenantDetailedInformationable
    // The collection of tenant tags across managed tenants.
    tenantTags []TenantTagable
    // The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
    windowsDeviceMalwareStates []WindowsDeviceMalwareStateable
    // The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
    windowsProtectionStates []WindowsProtectionStateable
}
// NewManagedTenant instantiates a new managedTenant and sets the default values.
func NewManagedTenant()(*ManagedTenant) {
    m := &ManagedTenant{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.managedTenant";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagedTenantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenant(), nil
}
// GetAggregatedPolicyCompliances gets the aggregatedPolicyCompliances property value. Aggregate view of device compliance policies across managed tenants.
func (m *ManagedTenant) GetAggregatedPolicyCompliances()([]AggregatedPolicyComplianceable) {
    return m.aggregatedPolicyCompliances
}
// GetAuditEvents gets the auditEvents property value. The collection of audit events across managed tenants.
func (m *ManagedTenant) GetAuditEvents()([]AuditEventable) {
    return m.auditEvents
}
// GetCloudPcConnections gets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
func (m *ManagedTenant) GetCloudPcConnections()([]CloudPcConnectionable) {
    return m.cloudPcConnections
}
// GetCloudPcDevices gets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
func (m *ManagedTenant) GetCloudPcDevices()([]CloudPcDeviceable) {
    return m.cloudPcDevices
}
// GetCloudPcsOverview gets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
func (m *ManagedTenant) GetCloudPcsOverview()([]CloudPcOverviewable) {
    return m.cloudPcsOverview
}
// GetConditionalAccessPolicyCoverages gets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
func (m *ManagedTenant) GetConditionalAccessPolicyCoverages()([]ConditionalAccessPolicyCoverageable) {
    return m.conditionalAccessPolicyCoverages
}
// GetCredentialUserRegistrationsSummaries gets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *ManagedTenant) GetCredentialUserRegistrationsSummaries()([]CredentialUserRegistrationsSummaryable) {
    return m.credentialUserRegistrationsSummaries
}
// GetDeviceCompliancePolicySettingStateSummaries gets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenant) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable) {
    return m.deviceCompliancePolicySettingStateSummaries
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aggregatedPolicyCompliances"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAggregatedPolicyComplianceFromDiscriminatorValue , m.SetAggregatedPolicyCompliances)
    res["auditEvents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAuditEventFromDiscriminatorValue , m.SetAuditEvents)
    res["cloudPcConnections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcConnectionFromDiscriminatorValue , m.SetCloudPcConnections)
    res["cloudPcDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcDeviceFromDiscriminatorValue , m.SetCloudPcDevices)
    res["cloudPcsOverview"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcOverviewFromDiscriminatorValue , m.SetCloudPcsOverview)
    res["conditionalAccessPolicyCoverages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateConditionalAccessPolicyCoverageFromDiscriminatorValue , m.SetConditionalAccessPolicyCoverages)
    res["credentialUserRegistrationsSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCredentialUserRegistrationsSummaryFromDiscriminatorValue , m.SetCredentialUserRegistrationsSummaries)
    res["deviceCompliancePolicySettingStateSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue , m.SetDeviceCompliancePolicySettingStateSummaries)
    res["managedDeviceCompliances"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceComplianceFromDiscriminatorValue , m.SetManagedDeviceCompliances)
    res["managedDeviceComplianceTrends"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceComplianceTrendFromDiscriminatorValue , m.SetManagedDeviceComplianceTrends)
    res["managementActions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementActionFromDiscriminatorValue , m.SetManagementActions)
    res["managementActionTenantDeploymentStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementActionTenantDeploymentStatusFromDiscriminatorValue , m.SetManagementActionTenantDeploymentStatuses)
    res["managementIntents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementIntentFromDiscriminatorValue , m.SetManagementIntents)
    res["managementTemplateCollections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementTemplateCollectionFromDiscriminatorValue , m.SetManagementTemplateCollections)
    res["managementTemplates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementTemplateFromDiscriminatorValue , m.SetManagementTemplates)
    res["managementTemplateSteps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementTemplateStepFromDiscriminatorValue , m.SetManagementTemplateSteps)
    res["managementTemplateStepVersions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementTemplateStepVersionFromDiscriminatorValue , m.SetManagementTemplateStepVersions)
    res["myRoles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMyRoleFromDiscriminatorValue , m.SetMyRoles)
    res["tenantGroups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTenantGroupFromDiscriminatorValue , m.SetTenantGroups)
    res["tenants"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTenantFromDiscriminatorValue , m.SetTenants)
    res["tenantsCustomizedInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTenantCustomizedInformationFromDiscriminatorValue , m.SetTenantsCustomizedInformation)
    res["tenantsDetailedInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTenantDetailedInformationFromDiscriminatorValue , m.SetTenantsDetailedInformation)
    res["tenantTags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTenantTagFromDiscriminatorValue , m.SetTenantTags)
    res["windowsDeviceMalwareStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsDeviceMalwareStateFromDiscriminatorValue , m.SetWindowsDeviceMalwareStates)
    res["windowsProtectionStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsProtectionStateFromDiscriminatorValue , m.SetWindowsProtectionStates)
    return res
}
// GetManagedDeviceCompliances gets the managedDeviceCompliances property value. The collection of compliance for managed devices across managed tenants.
func (m *ManagedTenant) GetManagedDeviceCompliances()([]ManagedDeviceComplianceable) {
    return m.managedDeviceCompliances
}
// GetManagedDeviceComplianceTrends gets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
func (m *ManagedTenant) GetManagedDeviceComplianceTrends()([]ManagedDeviceComplianceTrendable) {
    return m.managedDeviceComplianceTrends
}
// GetManagementActions gets the managementActions property value. The collection of baseline management actions across managed tenants.
func (m *ManagedTenant) GetManagementActions()([]ManagementActionable) {
    return m.managementActions
}
// GetManagementActionTenantDeploymentStatuses gets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
func (m *ManagedTenant) GetManagementActionTenantDeploymentStatuses()([]ManagementActionTenantDeploymentStatusable) {
    return m.managementActionTenantDeploymentStatuses
}
// GetManagementIntents gets the managementIntents property value. The collection of baseline management intents across managed tenants.
func (m *ManagedTenant) GetManagementIntents()([]ManagementIntentable) {
    return m.managementIntents
}
// GetManagementTemplateCollections gets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagedTenant) GetManagementTemplateCollections()([]ManagementTemplateCollectionable) {
    return m.managementTemplateCollections
}
// GetManagementTemplates gets the managementTemplates property value. The collection of baseline management templates across managed tenants.
func (m *ManagedTenant) GetManagementTemplates()([]ManagementTemplateable) {
    return m.managementTemplates
}
// GetManagementTemplateSteps gets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagedTenant) GetManagementTemplateSteps()([]ManagementTemplateStepable) {
    return m.managementTemplateSteps
}
// GetManagementTemplateStepVersions gets the managementTemplateStepVersions property value. The managementTemplateStepVersions property
func (m *ManagedTenant) GetManagementTemplateStepVersions()([]ManagementTemplateStepVersionable) {
    return m.managementTemplateStepVersions
}
// GetMyRoles gets the myRoles property value. The collection of role assignments to a signed-in user for a managed tenant.
func (m *ManagedTenant) GetMyRoles()([]MyRoleable) {
    return m.myRoles
}
// GetTenantGroups gets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *ManagedTenant) GetTenantGroups()([]TenantGroupable) {
    return m.tenantGroups
}
// GetTenants gets the tenants property value. The collection of tenants associated with the managing entity.
func (m *ManagedTenant) GetTenants()([]Tenantable) {
    return m.tenants
}
// GetTenantsCustomizedInformation gets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
func (m *ManagedTenant) GetTenantsCustomizedInformation()([]TenantCustomizedInformationable) {
    return m.tenantsCustomizedInformation
}
// GetTenantsDetailedInformation gets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
func (m *ManagedTenant) GetTenantsDetailedInformation()([]TenantDetailedInformationable) {
    return m.tenantsDetailedInformation
}
// GetTenantTags gets the tenantTags property value. The collection of tenant tags across managed tenants.
func (m *ManagedTenant) GetTenantTags()([]TenantTagable) {
    return m.tenantTags
}
// GetWindowsDeviceMalwareStates gets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsDeviceMalwareStates()([]WindowsDeviceMalwareStateable) {
    return m.windowsDeviceMalwareStates
}
// GetWindowsProtectionStates gets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsProtectionStates()([]WindowsProtectionStateable) {
    return m.windowsProtectionStates
}
// Serialize serializes information the current object
func (m *ManagedTenant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAggregatedPolicyCompliances() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAggregatedPolicyCompliances())
        err = writer.WriteCollectionOfObjectValues("aggregatedPolicyCompliances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuditEvents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAuditEvents())
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcConnections() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCloudPcConnections())
        err = writer.WriteCollectionOfObjectValues("cloudPcConnections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcDevices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCloudPcDevices())
        err = writer.WriteCollectionOfObjectValues("cloudPcDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcsOverview() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCloudPcsOverview())
        err = writer.WriteCollectionOfObjectValues("cloudPcsOverview", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessPolicyCoverages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetConditionalAccessPolicyCoverages())
        err = writer.WriteCollectionOfObjectValues("conditionalAccessPolicyCoverages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCredentialUserRegistrationsSummaries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCredentialUserRegistrationsSummaries())
        err = writer.WriteCollectionOfObjectValues("credentialUserRegistrationsSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicySettingStateSummaries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceCompliancePolicySettingStateSummaries())
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicySettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceCompliances() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedDeviceCompliances())
        err = writer.WriteCollectionOfObjectValues("managedDeviceCompliances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceComplianceTrends() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedDeviceComplianceTrends())
        err = writer.WriteCollectionOfObjectValues("managedDeviceComplianceTrends", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementActions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementActions())
        err = writer.WriteCollectionOfObjectValues("managementActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementActionTenantDeploymentStatuses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementActionTenantDeploymentStatuses())
        err = writer.WriteCollectionOfObjectValues("managementActionTenantDeploymentStatuses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementIntents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementIntents())
        err = writer.WriteCollectionOfObjectValues("managementIntents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateCollections() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementTemplateCollections())
        err = writer.WriteCollectionOfObjectValues("managementTemplateCollections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementTemplates())
        err = writer.WriteCollectionOfObjectValues("managementTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateSteps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementTemplateSteps())
        err = writer.WriteCollectionOfObjectValues("managementTemplateSteps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateStepVersions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementTemplateStepVersions())
        err = writer.WriteCollectionOfObjectValues("managementTemplateStepVersions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMyRoles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMyRoles())
        err = writer.WriteCollectionOfObjectValues("myRoles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantGroups() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTenantGroups())
        err = writer.WriteCollectionOfObjectValues("tenantGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenants() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTenants())
        err = writer.WriteCollectionOfObjectValues("tenants", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantsCustomizedInformation() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTenantsCustomizedInformation())
        err = writer.WriteCollectionOfObjectValues("tenantsCustomizedInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantsDetailedInformation() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTenantsDetailedInformation())
        err = writer.WriteCollectionOfObjectValues("tenantsDetailedInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantTags() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTenantTags())
        err = writer.WriteCollectionOfObjectValues("tenantTags", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsDeviceMalwareStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsDeviceMalwareStates())
        err = writer.WriteCollectionOfObjectValues("windowsDeviceMalwareStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsProtectionStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsProtectionStates())
        err = writer.WriteCollectionOfObjectValues("windowsProtectionStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAggregatedPolicyCompliances sets the aggregatedPolicyCompliances property value. Aggregate view of device compliance policies across managed tenants.
func (m *ManagedTenant) SetAggregatedPolicyCompliances(value []AggregatedPolicyComplianceable)() {
    m.aggregatedPolicyCompliances = value
}
// SetAuditEvents sets the auditEvents property value. The collection of audit events across managed tenants.
func (m *ManagedTenant) SetAuditEvents(value []AuditEventable)() {
    m.auditEvents = value
}
// SetCloudPcConnections sets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
func (m *ManagedTenant) SetCloudPcConnections(value []CloudPcConnectionable)() {
    m.cloudPcConnections = value
}
// SetCloudPcDevices sets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
func (m *ManagedTenant) SetCloudPcDevices(value []CloudPcDeviceable)() {
    m.cloudPcDevices = value
}
// SetCloudPcsOverview sets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
func (m *ManagedTenant) SetCloudPcsOverview(value []CloudPcOverviewable)() {
    m.cloudPcsOverview = value
}
// SetConditionalAccessPolicyCoverages sets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
func (m *ManagedTenant) SetConditionalAccessPolicyCoverages(value []ConditionalAccessPolicyCoverageable)() {
    m.conditionalAccessPolicyCoverages = value
}
// SetCredentialUserRegistrationsSummaries sets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *ManagedTenant) SetCredentialUserRegistrationsSummaries(value []CredentialUserRegistrationsSummaryable)() {
    m.credentialUserRegistrationsSummaries = value
}
// SetDeviceCompliancePolicySettingStateSummaries sets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenant) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)() {
    m.deviceCompliancePolicySettingStateSummaries = value
}
// SetManagedDeviceCompliances sets the managedDeviceCompliances property value. The collection of compliance for managed devices across managed tenants.
func (m *ManagedTenant) SetManagedDeviceCompliances(value []ManagedDeviceComplianceable)() {
    m.managedDeviceCompliances = value
}
// SetManagedDeviceComplianceTrends sets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
func (m *ManagedTenant) SetManagedDeviceComplianceTrends(value []ManagedDeviceComplianceTrendable)() {
    m.managedDeviceComplianceTrends = value
}
// SetManagementActions sets the managementActions property value. The collection of baseline management actions across managed tenants.
func (m *ManagedTenant) SetManagementActions(value []ManagementActionable)() {
    m.managementActions = value
}
// SetManagementActionTenantDeploymentStatuses sets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
func (m *ManagedTenant) SetManagementActionTenantDeploymentStatuses(value []ManagementActionTenantDeploymentStatusable)() {
    m.managementActionTenantDeploymentStatuses = value
}
// SetManagementIntents sets the managementIntents property value. The collection of baseline management intents across managed tenants.
func (m *ManagedTenant) SetManagementIntents(value []ManagementIntentable)() {
    m.managementIntents = value
}
// SetManagementTemplateCollections sets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagedTenant) SetManagementTemplateCollections(value []ManagementTemplateCollectionable)() {
    m.managementTemplateCollections = value
}
// SetManagementTemplates sets the managementTemplates property value. The collection of baseline management templates across managed tenants.
func (m *ManagedTenant) SetManagementTemplates(value []ManagementTemplateable)() {
    m.managementTemplates = value
}
// SetManagementTemplateSteps sets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagedTenant) SetManagementTemplateSteps(value []ManagementTemplateStepable)() {
    m.managementTemplateSteps = value
}
// SetManagementTemplateStepVersions sets the managementTemplateStepVersions property value. The managementTemplateStepVersions property
func (m *ManagedTenant) SetManagementTemplateStepVersions(value []ManagementTemplateStepVersionable)() {
    m.managementTemplateStepVersions = value
}
// SetMyRoles sets the myRoles property value. The collection of role assignments to a signed-in user for a managed tenant.
func (m *ManagedTenant) SetMyRoles(value []MyRoleable)() {
    m.myRoles = value
}
// SetTenantGroups sets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *ManagedTenant) SetTenantGroups(value []TenantGroupable)() {
    m.tenantGroups = value
}
// SetTenants sets the tenants property value. The collection of tenants associated with the managing entity.
func (m *ManagedTenant) SetTenants(value []Tenantable)() {
    m.tenants = value
}
// SetTenantsCustomizedInformation sets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
func (m *ManagedTenant) SetTenantsCustomizedInformation(value []TenantCustomizedInformationable)() {
    m.tenantsCustomizedInformation = value
}
// SetTenantsDetailedInformation sets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
func (m *ManagedTenant) SetTenantsDetailedInformation(value []TenantDetailedInformationable)() {
    m.tenantsDetailedInformation = value
}
// SetTenantTags sets the tenantTags property value. The collection of tenant tags across managed tenants.
func (m *ManagedTenant) SetTenantTags(value []TenantTagable)() {
    m.tenantTags = value
}
// SetWindowsDeviceMalwareStates sets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) SetWindowsDeviceMalwareStates(value []WindowsDeviceMalwareStateable)() {
    m.windowsDeviceMalwareStates = value
}
// SetWindowsProtectionStates sets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) SetWindowsProtectionStates(value []WindowsProtectionStateable)() {
    m.windowsProtectionStates = value
}
