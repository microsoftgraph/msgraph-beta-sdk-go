package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagedTenant 
type ManagedTenant struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewManagedTenant instantiates a new ManagedTenant and sets the default values.
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
    val, err := m.GetBackingStore().Get("aggregatedPolicyCompliances")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AggregatedPolicyComplianceable)
    }
    return nil
}
// GetAuditEvents gets the auditEvents property value. The collection of audit events across managed tenants.
func (m *ManagedTenant) GetAuditEvents()([]AuditEventable) {
    val, err := m.GetBackingStore().Get("auditEvents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuditEventable)
    }
    return nil
}
// GetCloudPcConnections gets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
func (m *ManagedTenant) GetCloudPcConnections()([]CloudPcConnectionable) {
    val, err := m.GetBackingStore().Get("cloudPcConnections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcConnectionable)
    }
    return nil
}
// GetCloudPcDevices gets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
func (m *ManagedTenant) GetCloudPcDevices()([]CloudPcDeviceable) {
    val, err := m.GetBackingStore().Get("cloudPcDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcDeviceable)
    }
    return nil
}
// GetCloudPcsOverview gets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
func (m *ManagedTenant) GetCloudPcsOverview()([]CloudPcOverviewable) {
    val, err := m.GetBackingStore().Get("cloudPcsOverview")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcOverviewable)
    }
    return nil
}
// GetConditionalAccessPolicyCoverages gets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
func (m *ManagedTenant) GetConditionalAccessPolicyCoverages()([]ConditionalAccessPolicyCoverageable) {
    val, err := m.GetBackingStore().Get("conditionalAccessPolicyCoverages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ConditionalAccessPolicyCoverageable)
    }
    return nil
}
// GetCredentialUserRegistrationsSummaries gets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *ManagedTenant) GetCredentialUserRegistrationsSummaries()([]CredentialUserRegistrationsSummaryable) {
    val, err := m.GetBackingStore().Get("credentialUserRegistrationsSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CredentialUserRegistrationsSummaryable)
    }
    return nil
}
// GetDeviceCompliancePolicySettingStateSummaries gets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenant) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable) {
    val, err := m.GetBackingStore().Get("deviceCompliancePolicySettingStateSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceCompliancePolicySettingStateSummaryable)
    }
    return nil
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
    res["managedTenantAlertLogs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertLogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertLogable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantAlertLogable)
            }
            m.SetManagedTenantAlertLogs(res)
        }
        return nil
    }
    res["managedTenantAlertRuleDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertRuleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertRuleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantAlertRuleDefinitionable)
            }
            m.SetManagedTenantAlertRuleDefinitions(res)
        }
        return nil
    }
    res["managedTenantAlertRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertRuleable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantAlertRuleable)
            }
            m.SetManagedTenantAlertRules(res)
        }
        return nil
    }
    res["managedTenantAlerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantAlertable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantAlertable)
            }
            m.SetManagedTenantAlerts(res)
        }
        return nil
    }
    res["managedTenantApiNotifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantApiNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantApiNotificationable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantApiNotificationable)
            }
            m.SetManagedTenantApiNotifications(res)
        }
        return nil
    }
    res["managedTenantEmailNotifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantEmailNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantEmailNotificationable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantEmailNotificationable)
            }
            m.SetManagedTenantEmailNotifications(res)
        }
        return nil
    }
    res["managedTenantTicketingEndpoints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedTenantTicketingEndpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedTenantTicketingEndpointable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedTenantTicketingEndpointable)
            }
            m.SetManagedTenantTicketingEndpoints(res)
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
    res["managementTemplateCollectionTenantSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateCollectionTenantSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateCollectionTenantSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateCollectionTenantSummaryable)
            }
            m.SetManagementTemplateCollectionTenantSummaries(res)
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
    res["managementTemplateStepTenantSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepTenantSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepTenantSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateStepTenantSummaryable)
            }
            m.SetManagementTemplateStepTenantSummaries(res)
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
    val, err := m.GetBackingStore().Get("managedDeviceCompliances")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceComplianceable)
    }
    return nil
}
// GetManagedDeviceComplianceTrends gets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
func (m *ManagedTenant) GetManagedDeviceComplianceTrends()([]ManagedDeviceComplianceTrendable) {
    val, err := m.GetBackingStore().Get("managedDeviceComplianceTrends")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceComplianceTrendable)
    }
    return nil
}
// GetManagedTenantAlertLogs gets the managedTenantAlertLogs property value. The managedTenantAlertLogs property
func (m *ManagedTenant) GetManagedTenantAlertLogs()([]ManagedTenantAlertLogable) {
    val, err := m.GetBackingStore().Get("managedTenantAlertLogs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedTenantAlertLogable)
    }
    return nil
}
// GetManagedTenantAlertRuleDefinitions gets the managedTenantAlertRuleDefinitions property value. The managedTenantAlertRuleDefinitions property
func (m *ManagedTenant) GetManagedTenantAlertRuleDefinitions()([]ManagedTenantAlertRuleDefinitionable) {
    val, err := m.GetBackingStore().Get("managedTenantAlertRuleDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedTenantAlertRuleDefinitionable)
    }
    return nil
}
// GetManagedTenantAlertRules gets the managedTenantAlertRules property value. The managedTenantAlertRules property
func (m *ManagedTenant) GetManagedTenantAlertRules()([]ManagedTenantAlertRuleable) {
    val, err := m.GetBackingStore().Get("managedTenantAlertRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedTenantAlertRuleable)
    }
    return nil
}
// GetManagedTenantAlerts gets the managedTenantAlerts property value. The managedTenantAlerts property
func (m *ManagedTenant) GetManagedTenantAlerts()([]ManagedTenantAlertable) {
    val, err := m.GetBackingStore().Get("managedTenantAlerts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedTenantAlertable)
    }
    return nil
}
// GetManagedTenantApiNotifications gets the managedTenantApiNotifications property value. The managedTenantApiNotifications property
func (m *ManagedTenant) GetManagedTenantApiNotifications()([]ManagedTenantApiNotificationable) {
    val, err := m.GetBackingStore().Get("managedTenantApiNotifications")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedTenantApiNotificationable)
    }
    return nil
}
// GetManagedTenantEmailNotifications gets the managedTenantEmailNotifications property value. The managedTenantEmailNotifications property
func (m *ManagedTenant) GetManagedTenantEmailNotifications()([]ManagedTenantEmailNotificationable) {
    val, err := m.GetBackingStore().Get("managedTenantEmailNotifications")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedTenantEmailNotificationable)
    }
    return nil
}
// GetManagedTenantTicketingEndpoints gets the managedTenantTicketingEndpoints property value. The managedTenantTicketingEndpoints property
func (m *ManagedTenant) GetManagedTenantTicketingEndpoints()([]ManagedTenantTicketingEndpointable) {
    val, err := m.GetBackingStore().Get("managedTenantTicketingEndpoints")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedTenantTicketingEndpointable)
    }
    return nil
}
// GetManagementActions gets the managementActions property value. The collection of baseline management actions across managed tenants.
func (m *ManagedTenant) GetManagementActions()([]ManagementActionable) {
    val, err := m.GetBackingStore().Get("managementActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementActionable)
    }
    return nil
}
// GetManagementActionTenantDeploymentStatuses gets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
func (m *ManagedTenant) GetManagementActionTenantDeploymentStatuses()([]ManagementActionTenantDeploymentStatusable) {
    val, err := m.GetBackingStore().Get("managementActionTenantDeploymentStatuses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementActionTenantDeploymentStatusable)
    }
    return nil
}
// GetManagementIntents gets the managementIntents property value. The collection of baseline management intents across managed tenants.
func (m *ManagedTenant) GetManagementIntents()([]ManagementIntentable) {
    val, err := m.GetBackingStore().Get("managementIntents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementIntentable)
    }
    return nil
}
// GetManagementTemplateCollections gets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagedTenant) GetManagementTemplateCollections()([]ManagementTemplateCollectionable) {
    val, err := m.GetBackingStore().Get("managementTemplateCollections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementTemplateCollectionable)
    }
    return nil
}
// GetManagementTemplateCollectionTenantSummaries gets the managementTemplateCollectionTenantSummaries property value. The managementTemplateCollectionTenantSummaries property
func (m *ManagedTenant) GetManagementTemplateCollectionTenantSummaries()([]ManagementTemplateCollectionTenantSummaryable) {
    val, err := m.GetBackingStore().Get("managementTemplateCollectionTenantSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementTemplateCollectionTenantSummaryable)
    }
    return nil
}
// GetManagementTemplates gets the managementTemplates property value. The collection of baseline management templates across managed tenants.
func (m *ManagedTenant) GetManagementTemplates()([]ManagementTemplateable) {
    val, err := m.GetBackingStore().Get("managementTemplates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementTemplateable)
    }
    return nil
}
// GetManagementTemplateSteps gets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagedTenant) GetManagementTemplateSteps()([]ManagementTemplateStepable) {
    val, err := m.GetBackingStore().Get("managementTemplateSteps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementTemplateStepable)
    }
    return nil
}
// GetManagementTemplateStepTenantSummaries gets the managementTemplateStepTenantSummaries property value. The managementTemplateStepTenantSummaries property
func (m *ManagedTenant) GetManagementTemplateStepTenantSummaries()([]ManagementTemplateStepTenantSummaryable) {
    val, err := m.GetBackingStore().Get("managementTemplateStepTenantSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementTemplateStepTenantSummaryable)
    }
    return nil
}
// GetManagementTemplateStepVersions gets the managementTemplateStepVersions property value. The managementTemplateStepVersions property
func (m *ManagedTenant) GetManagementTemplateStepVersions()([]ManagementTemplateStepVersionable) {
    val, err := m.GetBackingStore().Get("managementTemplateStepVersions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagementTemplateStepVersionable)
    }
    return nil
}
// GetMyRoles gets the myRoles property value. The collection of role assignments to a signed-in user for a managed tenant.
func (m *ManagedTenant) GetMyRoles()([]MyRoleable) {
    val, err := m.GetBackingStore().Get("myRoles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MyRoleable)
    }
    return nil
}
// GetTenantGroups gets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *ManagedTenant) GetTenantGroups()([]TenantGroupable) {
    val, err := m.GetBackingStore().Get("tenantGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TenantGroupable)
    }
    return nil
}
// GetTenants gets the tenants property value. The collection of tenants associated with the managing entity.
func (m *ManagedTenant) GetTenants()([]Tenantable) {
    val, err := m.GetBackingStore().Get("tenants")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Tenantable)
    }
    return nil
}
// GetTenantsCustomizedInformation gets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
func (m *ManagedTenant) GetTenantsCustomizedInformation()([]TenantCustomizedInformationable) {
    val, err := m.GetBackingStore().Get("tenantsCustomizedInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TenantCustomizedInformationable)
    }
    return nil
}
// GetTenantsDetailedInformation gets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
func (m *ManagedTenant) GetTenantsDetailedInformation()([]TenantDetailedInformationable) {
    val, err := m.GetBackingStore().Get("tenantsDetailedInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TenantDetailedInformationable)
    }
    return nil
}
// GetTenantTags gets the tenantTags property value. The collection of tenant tags across managed tenants.
func (m *ManagedTenant) GetTenantTags()([]TenantTagable) {
    val, err := m.GetBackingStore().Get("tenantTags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TenantTagable)
    }
    return nil
}
// GetWindowsDeviceMalwareStates gets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsDeviceMalwareStates()([]WindowsDeviceMalwareStateable) {
    val, err := m.GetBackingStore().Get("windowsDeviceMalwareStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsDeviceMalwareStateable)
    }
    return nil
}
// GetWindowsProtectionStates gets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsProtectionStates()([]WindowsProtectionStateable) {
    val, err := m.GetBackingStore().Get("windowsProtectionStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsProtectionStateable)
    }
    return nil
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
    if m.GetManagedTenantAlertLogs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantAlertLogs()))
        for i, v := range m.GetManagedTenantAlertLogs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantAlertLogs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantAlertRuleDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantAlertRuleDefinitions()))
        for i, v := range m.GetManagedTenantAlertRuleDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantAlertRuleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantAlertRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantAlertRules()))
        for i, v := range m.GetManagedTenantAlertRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantAlertRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantAlerts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantAlerts()))
        for i, v := range m.GetManagedTenantAlerts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantAlerts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantApiNotifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantApiNotifications()))
        for i, v := range m.GetManagedTenantApiNotifications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantApiNotifications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantEmailNotifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantEmailNotifications()))
        for i, v := range m.GetManagedTenantEmailNotifications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantEmailNotifications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedTenantTicketingEndpoints() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedTenantTicketingEndpoints()))
        for i, v := range m.GetManagedTenantTicketingEndpoints() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedTenantTicketingEndpoints", cast)
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
    if m.GetManagementTemplateCollectionTenantSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateCollectionTenantSummaries()))
        for i, v := range m.GetManagementTemplateCollectionTenantSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateCollectionTenantSummaries", cast)
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
    if m.GetManagementTemplateStepTenantSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateStepTenantSummaries()))
        for i, v := range m.GetManagementTemplateStepTenantSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateStepTenantSummaries", cast)
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
    err := m.GetBackingStore().Set("aggregatedPolicyCompliances", value)
    if err != nil {
        panic(err)
    }
}
// SetAuditEvents sets the auditEvents property value. The collection of audit events across managed tenants.
func (m *ManagedTenant) SetAuditEvents(value []AuditEventable)() {
    err := m.GetBackingStore().Set("auditEvents", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcConnections sets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
func (m *ManagedTenant) SetCloudPcConnections(value []CloudPcConnectionable)() {
    err := m.GetBackingStore().Set("cloudPcConnections", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcDevices sets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
func (m *ManagedTenant) SetCloudPcDevices(value []CloudPcDeviceable)() {
    err := m.GetBackingStore().Set("cloudPcDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcsOverview sets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
func (m *ManagedTenant) SetCloudPcsOverview(value []CloudPcOverviewable)() {
    err := m.GetBackingStore().Set("cloudPcsOverview", value)
    if err != nil {
        panic(err)
    }
}
// SetConditionalAccessPolicyCoverages sets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
func (m *ManagedTenant) SetConditionalAccessPolicyCoverages(value []ConditionalAccessPolicyCoverageable)() {
    err := m.GetBackingStore().Set("conditionalAccessPolicyCoverages", value)
    if err != nil {
        panic(err)
    }
}
// SetCredentialUserRegistrationsSummaries sets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *ManagedTenant) SetCredentialUserRegistrationsSummaries(value []CredentialUserRegistrationsSummaryable)() {
    err := m.GetBackingStore().Set("credentialUserRegistrationsSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCompliancePolicySettingStateSummaries sets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenant) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)() {
    err := m.GetBackingStore().Set("deviceCompliancePolicySettingStateSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceCompliances sets the managedDeviceCompliances property value. The collection of compliance for managed devices across managed tenants.
func (m *ManagedTenant) SetManagedDeviceCompliances(value []ManagedDeviceComplianceable)() {
    err := m.GetBackingStore().Set("managedDeviceCompliances", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceComplianceTrends sets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
func (m *ManagedTenant) SetManagedDeviceComplianceTrends(value []ManagedDeviceComplianceTrendable)() {
    err := m.GetBackingStore().Set("managedDeviceComplianceTrends", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedTenantAlertLogs sets the managedTenantAlertLogs property value. The managedTenantAlertLogs property
func (m *ManagedTenant) SetManagedTenantAlertLogs(value []ManagedTenantAlertLogable)() {
    err := m.GetBackingStore().Set("managedTenantAlertLogs", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedTenantAlertRuleDefinitions sets the managedTenantAlertRuleDefinitions property value. The managedTenantAlertRuleDefinitions property
func (m *ManagedTenant) SetManagedTenantAlertRuleDefinitions(value []ManagedTenantAlertRuleDefinitionable)() {
    err := m.GetBackingStore().Set("managedTenantAlertRuleDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedTenantAlertRules sets the managedTenantAlertRules property value. The managedTenantAlertRules property
func (m *ManagedTenant) SetManagedTenantAlertRules(value []ManagedTenantAlertRuleable)() {
    err := m.GetBackingStore().Set("managedTenantAlertRules", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedTenantAlerts sets the managedTenantAlerts property value. The managedTenantAlerts property
func (m *ManagedTenant) SetManagedTenantAlerts(value []ManagedTenantAlertable)() {
    err := m.GetBackingStore().Set("managedTenantAlerts", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedTenantApiNotifications sets the managedTenantApiNotifications property value. The managedTenantApiNotifications property
func (m *ManagedTenant) SetManagedTenantApiNotifications(value []ManagedTenantApiNotificationable)() {
    err := m.GetBackingStore().Set("managedTenantApiNotifications", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedTenantEmailNotifications sets the managedTenantEmailNotifications property value. The managedTenantEmailNotifications property
func (m *ManagedTenant) SetManagedTenantEmailNotifications(value []ManagedTenantEmailNotificationable)() {
    err := m.GetBackingStore().Set("managedTenantEmailNotifications", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedTenantTicketingEndpoints sets the managedTenantTicketingEndpoints property value. The managedTenantTicketingEndpoints property
func (m *ManagedTenant) SetManagedTenantTicketingEndpoints(value []ManagedTenantTicketingEndpointable)() {
    err := m.GetBackingStore().Set("managedTenantTicketingEndpoints", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementActions sets the managementActions property value. The collection of baseline management actions across managed tenants.
func (m *ManagedTenant) SetManagementActions(value []ManagementActionable)() {
    err := m.GetBackingStore().Set("managementActions", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementActionTenantDeploymentStatuses sets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
func (m *ManagedTenant) SetManagementActionTenantDeploymentStatuses(value []ManagementActionTenantDeploymentStatusable)() {
    err := m.GetBackingStore().Set("managementActionTenantDeploymentStatuses", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementIntents sets the managementIntents property value. The collection of baseline management intents across managed tenants.
func (m *ManagedTenant) SetManagementIntents(value []ManagementIntentable)() {
    err := m.GetBackingStore().Set("managementIntents", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateCollections sets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagedTenant) SetManagementTemplateCollections(value []ManagementTemplateCollectionable)() {
    err := m.GetBackingStore().Set("managementTemplateCollections", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateCollectionTenantSummaries sets the managementTemplateCollectionTenantSummaries property value. The managementTemplateCollectionTenantSummaries property
func (m *ManagedTenant) SetManagementTemplateCollectionTenantSummaries(value []ManagementTemplateCollectionTenantSummaryable)() {
    err := m.GetBackingStore().Set("managementTemplateCollectionTenantSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplates sets the managementTemplates property value. The collection of baseline management templates across managed tenants.
func (m *ManagedTenant) SetManagementTemplates(value []ManagementTemplateable)() {
    err := m.GetBackingStore().Set("managementTemplates", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateSteps sets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagedTenant) SetManagementTemplateSteps(value []ManagementTemplateStepable)() {
    err := m.GetBackingStore().Set("managementTemplateSteps", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateStepTenantSummaries sets the managementTemplateStepTenantSummaries property value. The managementTemplateStepTenantSummaries property
func (m *ManagedTenant) SetManagementTemplateStepTenantSummaries(value []ManagementTemplateStepTenantSummaryable)() {
    err := m.GetBackingStore().Set("managementTemplateStepTenantSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementTemplateStepVersions sets the managementTemplateStepVersions property value. The managementTemplateStepVersions property
func (m *ManagedTenant) SetManagementTemplateStepVersions(value []ManagementTemplateStepVersionable)() {
    err := m.GetBackingStore().Set("managementTemplateStepVersions", value)
    if err != nil {
        panic(err)
    }
}
// SetMyRoles sets the myRoles property value. The collection of role assignments to a signed-in user for a managed tenant.
func (m *ManagedTenant) SetMyRoles(value []MyRoleable)() {
    err := m.GetBackingStore().Set("myRoles", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantGroups sets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *ManagedTenant) SetTenantGroups(value []TenantGroupable)() {
    err := m.GetBackingStore().Set("tenantGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetTenants sets the tenants property value. The collection of tenants associated with the managing entity.
func (m *ManagedTenant) SetTenants(value []Tenantable)() {
    err := m.GetBackingStore().Set("tenants", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantsCustomizedInformation sets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
func (m *ManagedTenant) SetTenantsCustomizedInformation(value []TenantCustomizedInformationable)() {
    err := m.GetBackingStore().Set("tenantsCustomizedInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantsDetailedInformation sets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
func (m *ManagedTenant) SetTenantsDetailedInformation(value []TenantDetailedInformationable)() {
    err := m.GetBackingStore().Set("tenantsDetailedInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantTags sets the tenantTags property value. The collection of tenant tags across managed tenants.
func (m *ManagedTenant) SetTenantTags(value []TenantTagable)() {
    err := m.GetBackingStore().Set("tenantTags", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsDeviceMalwareStates sets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) SetWindowsDeviceMalwareStates(value []WindowsDeviceMalwareStateable)() {
    err := m.GetBackingStore().Set("windowsDeviceMalwareStates", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsProtectionStates sets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) SetWindowsProtectionStates(value []WindowsProtectionStateable)() {
    err := m.GetBackingStore().Set("windowsProtectionStates", value)
    if err != nil {
        panic(err)
    }
}
// ManagedTenantable 
type ManagedTenantable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAggregatedPolicyCompliances()([]AggregatedPolicyComplianceable)
    GetAuditEvents()([]AuditEventable)
    GetCloudPcConnections()([]CloudPcConnectionable)
    GetCloudPcDevices()([]CloudPcDeviceable)
    GetCloudPcsOverview()([]CloudPcOverviewable)
    GetConditionalAccessPolicyCoverages()([]ConditionalAccessPolicyCoverageable)
    GetCredentialUserRegistrationsSummaries()([]CredentialUserRegistrationsSummaryable)
    GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable)
    GetManagedDeviceCompliances()([]ManagedDeviceComplianceable)
    GetManagedDeviceComplianceTrends()([]ManagedDeviceComplianceTrendable)
    GetManagedTenantAlertLogs()([]ManagedTenantAlertLogable)
    GetManagedTenantAlertRuleDefinitions()([]ManagedTenantAlertRuleDefinitionable)
    GetManagedTenantAlertRules()([]ManagedTenantAlertRuleable)
    GetManagedTenantAlerts()([]ManagedTenantAlertable)
    GetManagedTenantApiNotifications()([]ManagedTenantApiNotificationable)
    GetManagedTenantEmailNotifications()([]ManagedTenantEmailNotificationable)
    GetManagedTenantTicketingEndpoints()([]ManagedTenantTicketingEndpointable)
    GetManagementActions()([]ManagementActionable)
    GetManagementActionTenantDeploymentStatuses()([]ManagementActionTenantDeploymentStatusable)
    GetManagementIntents()([]ManagementIntentable)
    GetManagementTemplateCollections()([]ManagementTemplateCollectionable)
    GetManagementTemplateCollectionTenantSummaries()([]ManagementTemplateCollectionTenantSummaryable)
    GetManagementTemplates()([]ManagementTemplateable)
    GetManagementTemplateSteps()([]ManagementTemplateStepable)
    GetManagementTemplateStepTenantSummaries()([]ManagementTemplateStepTenantSummaryable)
    GetManagementTemplateStepVersions()([]ManagementTemplateStepVersionable)
    GetMyRoles()([]MyRoleable)
    GetTenantGroups()([]TenantGroupable)
    GetTenants()([]Tenantable)
    GetTenantsCustomizedInformation()([]TenantCustomizedInformationable)
    GetTenantsDetailedInformation()([]TenantDetailedInformationable)
    GetTenantTags()([]TenantTagable)
    GetWindowsDeviceMalwareStates()([]WindowsDeviceMalwareStateable)
    GetWindowsProtectionStates()([]WindowsProtectionStateable)
    SetAggregatedPolicyCompliances(value []AggregatedPolicyComplianceable)()
    SetAuditEvents(value []AuditEventable)()
    SetCloudPcConnections(value []CloudPcConnectionable)()
    SetCloudPcDevices(value []CloudPcDeviceable)()
    SetCloudPcsOverview(value []CloudPcOverviewable)()
    SetConditionalAccessPolicyCoverages(value []ConditionalAccessPolicyCoverageable)()
    SetCredentialUserRegistrationsSummaries(value []CredentialUserRegistrationsSummaryable)()
    SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)()
    SetManagedDeviceCompliances(value []ManagedDeviceComplianceable)()
    SetManagedDeviceComplianceTrends(value []ManagedDeviceComplianceTrendable)()
    SetManagedTenantAlertLogs(value []ManagedTenantAlertLogable)()
    SetManagedTenantAlertRuleDefinitions(value []ManagedTenantAlertRuleDefinitionable)()
    SetManagedTenantAlertRules(value []ManagedTenantAlertRuleable)()
    SetManagedTenantAlerts(value []ManagedTenantAlertable)()
    SetManagedTenantApiNotifications(value []ManagedTenantApiNotificationable)()
    SetManagedTenantEmailNotifications(value []ManagedTenantEmailNotificationable)()
    SetManagedTenantTicketingEndpoints(value []ManagedTenantTicketingEndpointable)()
    SetManagementActions(value []ManagementActionable)()
    SetManagementActionTenantDeploymentStatuses(value []ManagementActionTenantDeploymentStatusable)()
    SetManagementIntents(value []ManagementIntentable)()
    SetManagementTemplateCollections(value []ManagementTemplateCollectionable)()
    SetManagementTemplateCollectionTenantSummaries(value []ManagementTemplateCollectionTenantSummaryable)()
    SetManagementTemplates(value []ManagementTemplateable)()
    SetManagementTemplateSteps(value []ManagementTemplateStepable)()
    SetManagementTemplateStepTenantSummaries(value []ManagementTemplateStepTenantSummaryable)()
    SetManagementTemplateStepVersions(value []ManagementTemplateStepVersionable)()
    SetMyRoles(value []MyRoleable)()
    SetTenantGroups(value []TenantGroupable)()
    SetTenants(value []Tenantable)()
    SetTenantsCustomizedInformation(value []TenantCustomizedInformationable)()
    SetTenantsDetailedInformation(value []TenantDetailedInformationable)()
    SetTenantTags(value []TenantTagable)()
    SetWindowsDeviceMalwareStates(value []WindowsDeviceMalwareStateable)()
    SetWindowsProtectionStates(value []WindowsProtectionStateable)()
}
