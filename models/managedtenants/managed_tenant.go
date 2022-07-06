package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedTenant provides operations to manage the collection of accessReviewDecision entities.
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
    return m
}
// CreateManagedTenantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenant(), nil
}
// GetAggregatedPolicyCompliances gets the aggregatedPolicyCompliances property value. Aggregate view of device compliance policies across managed tenants.
func (m *ManagedTenant) GetAggregatedPolicyCompliances()([]AggregatedPolicyComplianceable) {
    if m == nil {
        return nil
    } else {
        return m.aggregatedPolicyCompliances
    }
}
// GetAuditEvents gets the auditEvents property value. The collection of audit events across managed tenants.
func (m *ManagedTenant) GetAuditEvents()([]AuditEventable) {
    if m == nil {
        return nil
    } else {
        return m.auditEvents
    }
}
// GetCloudPcConnections gets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
func (m *ManagedTenant) GetCloudPcConnections()([]CloudPcConnectionable) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcConnections
    }
}
// GetCloudPcDevices gets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
func (m *ManagedTenant) GetCloudPcDevices()([]CloudPcDeviceable) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcDevices
    }
}
// GetCloudPcsOverview gets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
func (m *ManagedTenant) GetCloudPcsOverview()([]CloudPcOverviewable) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcsOverview
    }
}
// GetConditionalAccessPolicyCoverages gets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
func (m *ManagedTenant) GetConditionalAccessPolicyCoverages()([]ConditionalAccessPolicyCoverageable) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessPolicyCoverages
    }
}
// GetCredentialUserRegistrationsSummaries gets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *ManagedTenant) GetCredentialUserRegistrationsSummaries()([]CredentialUserRegistrationsSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.credentialUserRegistrationsSummaries
    }
}
// GetDeviceCompliancePolicySettingStateSummaries gets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenant) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicySettingStateSummaries
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aggregatedPolicyCompliances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAggregatedPolicyComplianceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AggregatedPolicyComplianceable, len(val))
            for i, v := range val {
                res[i] = v.(AggregatedPolicyComplianceable)
            }
            m.SetAggregatedPolicyCompliances(res)
        }
        return nil
    }
    res["auditEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditEventable, len(val))
            for i, v := range val {
                res[i] = v.(AuditEventable)
            }
            m.SetAuditEvents(res)
        }
        return nil
    }
    res["cloudPcConnections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcConnectionable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcConnectionable)
            }
            m.SetCloudPcConnections(res)
        }
        return nil
    }
    res["cloudPcDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcDeviceable)
            }
            m.SetCloudPcDevices(res)
        }
        return nil
    }
    res["cloudPcsOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcOverviewable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcOverviewable)
            }
            m.SetCloudPcsOverview(res)
        }
        return nil
    }
    res["conditionalAccessPolicyCoverages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionalAccessPolicyCoverageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessPolicyCoverageable, len(val))
            for i, v := range val {
                res[i] = v.(ConditionalAccessPolicyCoverageable)
            }
            m.SetConditionalAccessPolicyCoverages(res)
        }
        return nil
    }
    res["credentialUserRegistrationsSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCredentialUserRegistrationsSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CredentialUserRegistrationsSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(CredentialUserRegistrationsSummaryable)
            }
            m.SetCredentialUserRegistrationsSummaries(res)
        }
        return nil
    }
    res["deviceCompliancePolicySettingStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicySettingStateSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceCompliancePolicySettingStateSummaryable)
            }
            m.SetDeviceCompliancePolicySettingStateSummaries(res)
        }
        return nil
    }
    res["managedDeviceCompliances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceComplianceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceComplianceable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceComplianceable)
            }
            m.SetManagedDeviceCompliances(res)
        }
        return nil
    }
    res["managedDeviceComplianceTrends"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceComplianceTrendFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceComplianceTrendable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceComplianceTrendable)
            }
            m.SetManagedDeviceComplianceTrends(res)
        }
        return nil
    }
    res["managementActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementActionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementActionable)
            }
            m.SetManagementActions(res)
        }
        return nil
    }
    res["managementActionTenantDeploymentStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementActionTenantDeploymentStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementActionTenantDeploymentStatusable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementActionTenantDeploymentStatusable)
            }
            m.SetManagementActionTenantDeploymentStatuses(res)
        }
        return nil
    }
    res["managementIntents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementIntentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementIntentable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementIntentable)
            }
            m.SetManagementIntents(res)
        }
        return nil
    }
    res["managementTemplateCollections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateCollectionable)
            }
            m.SetManagementTemplateCollections(res)
        }
        return nil
    }
    res["managementTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateable)
            }
            m.SetManagementTemplates(res)
        }
        return nil
    }
    res["managementTemplateSteps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateStepable)
            }
            m.SetManagementTemplateSteps(res)
        }
        return nil
    }
    res["managementTemplateStepVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepVersionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateStepVersionable)
            }
            m.SetManagementTemplateStepVersions(res)
        }
        return nil
    }
    res["myRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMyRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MyRoleable, len(val))
            for i, v := range val {
                res[i] = v.(MyRoleable)
            }
            m.SetMyRoles(res)
        }
        return nil
    }
    res["tenantGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantGroupable, len(val))
            for i, v := range val {
                res[i] = v.(TenantGroupable)
            }
            m.SetTenantGroups(res)
        }
        return nil
    }
    res["tenants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Tenantable, len(val))
            for i, v := range val {
                res[i] = v.(Tenantable)
            }
            m.SetTenants(res)
        }
        return nil
    }
    res["tenantsCustomizedInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantCustomizedInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantCustomizedInformationable, len(val))
            for i, v := range val {
                res[i] = v.(TenantCustomizedInformationable)
            }
            m.SetTenantsCustomizedInformation(res)
        }
        return nil
    }
    res["tenantsDetailedInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantDetailedInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantDetailedInformationable, len(val))
            for i, v := range val {
                res[i] = v.(TenantDetailedInformationable)
            }
            m.SetTenantsDetailedInformation(res)
        }
        return nil
    }
    res["tenantTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantTagable, len(val))
            for i, v := range val {
                res[i] = v.(TenantTagable)
            }
            m.SetTenantTags(res)
        }
        return nil
    }
    res["windowsDeviceMalwareStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDeviceMalwareStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDeviceMalwareStateable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsDeviceMalwareStateable)
            }
            m.SetWindowsDeviceMalwareStates(res)
        }
        return nil
    }
    res["windowsProtectionStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsProtectionStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsProtectionStateable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsProtectionStateable)
            }
            m.SetWindowsProtectionStates(res)
        }
        return nil
    }
    return res
}
// GetManagedDeviceCompliances gets the managedDeviceCompliances property value. The collection of compliance for managed devices across managed tenants.
func (m *ManagedTenant) GetManagedDeviceCompliances()([]ManagedDeviceComplianceable) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceCompliances
    }
}
// GetManagedDeviceComplianceTrends gets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
func (m *ManagedTenant) GetManagedDeviceComplianceTrends()([]ManagedDeviceComplianceTrendable) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceComplianceTrends
    }
}
// GetManagementActions gets the managementActions property value. The collection of baseline management actions across managed tenants.
func (m *ManagedTenant) GetManagementActions()([]ManagementActionable) {
    if m == nil {
        return nil
    } else {
        return m.managementActions
    }
}
// GetManagementActionTenantDeploymentStatuses gets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
func (m *ManagedTenant) GetManagementActionTenantDeploymentStatuses()([]ManagementActionTenantDeploymentStatusable) {
    if m == nil {
        return nil
    } else {
        return m.managementActionTenantDeploymentStatuses
    }
}
// GetManagementIntents gets the managementIntents property value. The collection of baseline management intents across managed tenants.
func (m *ManagedTenant) GetManagementIntents()([]ManagementIntentable) {
    if m == nil {
        return nil
    } else {
        return m.managementIntents
    }
}
// GetManagementTemplateCollections gets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagedTenant) GetManagementTemplateCollections()([]ManagementTemplateCollectionable) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateCollections
    }
}
// GetManagementTemplates gets the managementTemplates property value. The collection of baseline management templates across managed tenants.
func (m *ManagedTenant) GetManagementTemplates()([]ManagementTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplates
    }
}
// GetManagementTemplateSteps gets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagedTenant) GetManagementTemplateSteps()([]ManagementTemplateStepable) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateSteps
    }
}
// GetManagementTemplateStepVersions gets the managementTemplateStepVersions property value. The managementTemplateStepVersions property
func (m *ManagedTenant) GetManagementTemplateStepVersions()([]ManagementTemplateStepVersionable) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateStepVersions
    }
}
// GetMyRoles gets the myRoles property value. The collection of role assignments to a signed-in user for a managed tenant.
func (m *ManagedTenant) GetMyRoles()([]MyRoleable) {
    if m == nil {
        return nil
    } else {
        return m.myRoles
    }
}
// GetTenantGroups gets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *ManagedTenant) GetTenantGroups()([]TenantGroupable) {
    if m == nil {
        return nil
    } else {
        return m.tenantGroups
    }
}
// GetTenants gets the tenants property value. The collection of tenants associated with the managing entity.
func (m *ManagedTenant) GetTenants()([]Tenantable) {
    if m == nil {
        return nil
    } else {
        return m.tenants
    }
}
// GetTenantsCustomizedInformation gets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
func (m *ManagedTenant) GetTenantsCustomizedInformation()([]TenantCustomizedInformationable) {
    if m == nil {
        return nil
    } else {
        return m.tenantsCustomizedInformation
    }
}
// GetTenantsDetailedInformation gets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
func (m *ManagedTenant) GetTenantsDetailedInformation()([]TenantDetailedInformationable) {
    if m == nil {
        return nil
    } else {
        return m.tenantsDetailedInformation
    }
}
// GetTenantTags gets the tenantTags property value. The collection of tenant tags across managed tenants.
func (m *ManagedTenant) GetTenantTags()([]TenantTagable) {
    if m == nil {
        return nil
    } else {
        return m.tenantTags
    }
}
// GetWindowsDeviceMalwareStates gets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsDeviceMalwareStates()([]WindowsDeviceMalwareStateable) {
    if m == nil {
        return nil
    } else {
        return m.windowsDeviceMalwareStates
    }
}
// GetWindowsProtectionStates gets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsProtectionStates()([]WindowsProtectionStateable) {
    if m == nil {
        return nil
    } else {
        return m.windowsProtectionStates
    }
}
// Serialize serializes information the current object
func (m *ManagedTenant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAggregatedPolicyCompliances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAggregatedPolicyCompliances()))
        for i, v := range m.GetAggregatedPolicyCompliances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("aggregatedPolicyCompliances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuditEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuditEvents()))
        for i, v := range m.GetAuditEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcConnections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPcConnections()))
        for i, v := range m.GetCloudPcConnections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPcConnections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPcDevices()))
        for i, v := range m.GetCloudPcDevices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPcDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcsOverview() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPcsOverview()))
        for i, v := range m.GetCloudPcsOverview() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPcsOverview", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessPolicyCoverages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConditionalAccessPolicyCoverages()))
        for i, v := range m.GetConditionalAccessPolicyCoverages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("conditionalAccessPolicyCoverages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCredentialUserRegistrationsSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCredentialUserRegistrationsSummaries()))
        for i, v := range m.GetCredentialUserRegistrationsSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("credentialUserRegistrationsSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicySettingStateSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCompliancePolicySettingStateSummaries()))
        for i, v := range m.GetDeviceCompliancePolicySettingStateSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicySettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceCompliances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceCompliances()))
        for i, v := range m.GetManagedDeviceCompliances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceCompliances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceComplianceTrends() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceComplianceTrends()))
        for i, v := range m.GetManagedDeviceComplianceTrends() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceComplianceTrends", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementActions()))
        for i, v := range m.GetManagementActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementActionTenantDeploymentStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementActionTenantDeploymentStatuses()))
        for i, v := range m.GetManagementActionTenantDeploymentStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementActionTenantDeploymentStatuses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementIntents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementIntents()))
        for i, v := range m.GetManagementIntents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementIntents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateCollections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateCollections()))
        for i, v := range m.GetManagementTemplateCollections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateCollections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplates()))
        for i, v := range m.GetManagementTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateSteps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateSteps()))
        for i, v := range m.GetManagementTemplateSteps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateSteps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateStepVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateStepVersions()))
        for i, v := range m.GetManagementTemplateStepVersions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateStepVersions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMyRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMyRoles()))
        for i, v := range m.GetMyRoles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("myRoles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantGroups()))
        for i, v := range m.GetTenantGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenantGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenants()))
        for i, v := range m.GetTenants() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenants", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantsCustomizedInformation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantsCustomizedInformation()))
        for i, v := range m.GetTenantsCustomizedInformation() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenantsCustomizedInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantsDetailedInformation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantsDetailedInformation()))
        for i, v := range m.GetTenantsDetailedInformation() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenantsDetailedInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantTags()))
        for i, v := range m.GetTenantTags() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenantTags", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsDeviceMalwareStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsDeviceMalwareStates()))
        for i, v := range m.GetWindowsDeviceMalwareStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsDeviceMalwareStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsProtectionStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsProtectionStates()))
        for i, v := range m.GetWindowsProtectionStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsProtectionStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAggregatedPolicyCompliances sets the aggregatedPolicyCompliances property value. Aggregate view of device compliance policies across managed tenants.
func (m *ManagedTenant) SetAggregatedPolicyCompliances(value []AggregatedPolicyComplianceable)() {
    if m != nil {
        m.aggregatedPolicyCompliances = value
    }
}
// SetAuditEvents sets the auditEvents property value. The collection of audit events across managed tenants.
func (m *ManagedTenant) SetAuditEvents(value []AuditEventable)() {
    if m != nil {
        m.auditEvents = value
    }
}
// SetCloudPcConnections sets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
func (m *ManagedTenant) SetCloudPcConnections(value []CloudPcConnectionable)() {
    if m != nil {
        m.cloudPcConnections = value
    }
}
// SetCloudPcDevices sets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
func (m *ManagedTenant) SetCloudPcDevices(value []CloudPcDeviceable)() {
    if m != nil {
        m.cloudPcDevices = value
    }
}
// SetCloudPcsOverview sets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
func (m *ManagedTenant) SetCloudPcsOverview(value []CloudPcOverviewable)() {
    if m != nil {
        m.cloudPcsOverview = value
    }
}
// SetConditionalAccessPolicyCoverages sets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
func (m *ManagedTenant) SetConditionalAccessPolicyCoverages(value []ConditionalAccessPolicyCoverageable)() {
    if m != nil {
        m.conditionalAccessPolicyCoverages = value
    }
}
// SetCredentialUserRegistrationsSummaries sets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *ManagedTenant) SetCredentialUserRegistrationsSummaries(value []CredentialUserRegistrationsSummaryable)() {
    if m != nil {
        m.credentialUserRegistrationsSummaries = value
    }
}
// SetDeviceCompliancePolicySettingStateSummaries sets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenant) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)() {
    if m != nil {
        m.deviceCompliancePolicySettingStateSummaries = value
    }
}
// SetManagedDeviceCompliances sets the managedDeviceCompliances property value. The collection of compliance for managed devices across managed tenants.
func (m *ManagedTenant) SetManagedDeviceCompliances(value []ManagedDeviceComplianceable)() {
    if m != nil {
        m.managedDeviceCompliances = value
    }
}
// SetManagedDeviceComplianceTrends sets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
func (m *ManagedTenant) SetManagedDeviceComplianceTrends(value []ManagedDeviceComplianceTrendable)() {
    if m != nil {
        m.managedDeviceComplianceTrends = value
    }
}
// SetManagementActions sets the managementActions property value. The collection of baseline management actions across managed tenants.
func (m *ManagedTenant) SetManagementActions(value []ManagementActionable)() {
    if m != nil {
        m.managementActions = value
    }
}
// SetManagementActionTenantDeploymentStatuses sets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
func (m *ManagedTenant) SetManagementActionTenantDeploymentStatuses(value []ManagementActionTenantDeploymentStatusable)() {
    if m != nil {
        m.managementActionTenantDeploymentStatuses = value
    }
}
// SetManagementIntents sets the managementIntents property value. The collection of baseline management intents across managed tenants.
func (m *ManagedTenant) SetManagementIntents(value []ManagementIntentable)() {
    if m != nil {
        m.managementIntents = value
    }
}
// SetManagementTemplateCollections sets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagedTenant) SetManagementTemplateCollections(value []ManagementTemplateCollectionable)() {
    if m != nil {
        m.managementTemplateCollections = value
    }
}
// SetManagementTemplates sets the managementTemplates property value. The collection of baseline management templates across managed tenants.
func (m *ManagedTenant) SetManagementTemplates(value []ManagementTemplateable)() {
    if m != nil {
        m.managementTemplates = value
    }
}
// SetManagementTemplateSteps sets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagedTenant) SetManagementTemplateSteps(value []ManagementTemplateStepable)() {
    if m != nil {
        m.managementTemplateSteps = value
    }
}
// SetManagementTemplateStepVersions sets the managementTemplateStepVersions property value. The managementTemplateStepVersions property
func (m *ManagedTenant) SetManagementTemplateStepVersions(value []ManagementTemplateStepVersionable)() {
    if m != nil {
        m.managementTemplateStepVersions = value
    }
}
// SetMyRoles sets the myRoles property value. The collection of role assignments to a signed-in user for a managed tenant.
func (m *ManagedTenant) SetMyRoles(value []MyRoleable)() {
    if m != nil {
        m.myRoles = value
    }
}
// SetTenantGroups sets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *ManagedTenant) SetTenantGroups(value []TenantGroupable)() {
    if m != nil {
        m.tenantGroups = value
    }
}
// SetTenants sets the tenants property value. The collection of tenants associated with the managing entity.
func (m *ManagedTenant) SetTenants(value []Tenantable)() {
    if m != nil {
        m.tenants = value
    }
}
// SetTenantsCustomizedInformation sets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
func (m *ManagedTenant) SetTenantsCustomizedInformation(value []TenantCustomizedInformationable)() {
    if m != nil {
        m.tenantsCustomizedInformation = value
    }
}
// SetTenantsDetailedInformation sets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
func (m *ManagedTenant) SetTenantsDetailedInformation(value []TenantDetailedInformationable)() {
    if m != nil {
        m.tenantsDetailedInformation = value
    }
}
// SetTenantTags sets the tenantTags property value. The collection of tenant tags across managed tenants.
func (m *ManagedTenant) SetTenantTags(value []TenantTagable)() {
    if m != nil {
        m.tenantTags = value
    }
}
// SetWindowsDeviceMalwareStates sets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) SetWindowsDeviceMalwareStates(value []WindowsDeviceMalwareStateable)() {
    if m != nil {
        m.windowsDeviceMalwareStates = value
    }
}
// SetWindowsProtectionStates sets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) SetWindowsProtectionStates(value []WindowsProtectionStateable)() {
    if m != nil {
        m.windowsProtectionStates = value
    }
}
