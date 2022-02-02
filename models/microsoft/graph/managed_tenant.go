package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedTenant 
type ManagedTenant struct {
    Entity
    // Aggregate view of device compliance policies across managed tenants.
    aggregatedPolicyCompliances []AggregatedPolicyCompliance;
    // 
    auditEvents []AuditEvent;
    // The collection of cloud PC connections across managed tenants.
    cloudPcConnections []CloudPcConnection;
    // The collection of cloud PC devices across managed tenants.
    cloudPcDevices []CloudPcDevice;
    // Overview of cloud PC information across managed tenants.
    cloudPcsOverview []CloudPcOverview;
    // Aggregate view of conditional access policy coverage across managed tenants.
    conditionalAccessPolicyCoverages []ConditionalAccessPolicyCoverage;
    // Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
    credentialUserRegistrationsSummaries []CredentialUserRegistrationsSummary;
    // Summary information for device compliance policy setting states across managed tenants.
    deviceCompliancePolicySettingStateSummaries []DeviceCompliancePolicySettingStateSummary;
    // The collection of compliance for managed devices across managed tenants.
    managedDeviceCompliances []ManagedDeviceCompliance;
    // Trend insights for device compliance across managed tenants.
    managedDeviceComplianceTrends []ManagedDeviceComplianceTrend;
    // The collection of baseline management actions across managed tenants.
    managementActions []ManagementAction;
    // The tenant level status of management actions across managed tenants.
    managementActionTenantDeploymentStatuses []ManagementActionTenantDeploymentStatus;
    // The collection of baseline management intents across managed tenants.
    managementIntents []ManagementIntent;
    // 
    managementTemplateCollections []ManagementTemplateCollection;
    // The collection of baseline management templates across managed tenants.
    managementTemplates []ManagementTemplate;
    // 
    managementTemplateSteps []ManagementTemplateStep;
    // 
    managementTemplateStepVersions []ManagementTemplateStepVersion;
    // The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
    tenantGroups []TenantGroup;
    // The collection of tenants associated with the managing entity.
    tenants []Tenant;
    // The collection of tenant level customized information across managed tenants.
    tenantsCustomizedInformation []TenantCustomizedInformation;
    // The collection tenant level detailed information across managed tenants.
    tenantsDetailedInformation []TenantDetailedInformation;
    // The collection of tenant tags across managed tenants.
    tenantTags []TenantTag;
    // The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
    windowsDeviceMalwareStates []WindowsDeviceMalwareState;
    // The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
    windowsProtectionStates []WindowsProtectionState;
}
// NewManagedTenant instantiates a new managedTenant and sets the default values.
func NewManagedTenant()(*ManagedTenant) {
    m := &ManagedTenant{
        Entity: *NewEntity(),
    }
    return m
}
// GetAggregatedPolicyCompliances gets the aggregatedPolicyCompliances property value. Aggregate view of device compliance policies across managed tenants.
func (m *ManagedTenant) GetAggregatedPolicyCompliances()([]AggregatedPolicyCompliance) {
    if m == nil {
        return nil
    } else {
        return m.aggregatedPolicyCompliances
    }
}
// GetAuditEvents gets the auditEvents property value. 
func (m *ManagedTenant) GetAuditEvents()([]AuditEvent) {
    if m == nil {
        return nil
    } else {
        return m.auditEvents
    }
}
// GetCloudPcConnections gets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
func (m *ManagedTenant) GetCloudPcConnections()([]CloudPcConnection) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcConnections
    }
}
// GetCloudPcDevices gets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
func (m *ManagedTenant) GetCloudPcDevices()([]CloudPcDevice) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcDevices
    }
}
// GetCloudPcsOverview gets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
func (m *ManagedTenant) GetCloudPcsOverview()([]CloudPcOverview) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcsOverview
    }
}
// GetConditionalAccessPolicyCoverages gets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
func (m *ManagedTenant) GetConditionalAccessPolicyCoverages()([]ConditionalAccessPolicyCoverage) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessPolicyCoverages
    }
}
// GetCredentialUserRegistrationsSummaries gets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *ManagedTenant) GetCredentialUserRegistrationsSummaries()([]CredentialUserRegistrationsSummary) {
    if m == nil {
        return nil
    } else {
        return m.credentialUserRegistrationsSummaries
    }
}
// GetDeviceCompliancePolicySettingStateSummaries gets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenant) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicySettingStateSummaries
    }
}
// GetManagedDeviceCompliances gets the managedDeviceCompliances property value. The collection of compliance for managed devices across managed tenants.
func (m *ManagedTenant) GetManagedDeviceCompliances()([]ManagedDeviceCompliance) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceCompliances
    }
}
// GetManagedDeviceComplianceTrends gets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
func (m *ManagedTenant) GetManagedDeviceComplianceTrends()([]ManagedDeviceComplianceTrend) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceComplianceTrends
    }
}
// GetManagementActions gets the managementActions property value. The collection of baseline management actions across managed tenants.
func (m *ManagedTenant) GetManagementActions()([]ManagementAction) {
    if m == nil {
        return nil
    } else {
        return m.managementActions
    }
}
// GetManagementActionTenantDeploymentStatuses gets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
func (m *ManagedTenant) GetManagementActionTenantDeploymentStatuses()([]ManagementActionTenantDeploymentStatus) {
    if m == nil {
        return nil
    } else {
        return m.managementActionTenantDeploymentStatuses
    }
}
// GetManagementIntents gets the managementIntents property value. The collection of baseline management intents across managed tenants.
func (m *ManagedTenant) GetManagementIntents()([]ManagementIntent) {
    if m == nil {
        return nil
    } else {
        return m.managementIntents
    }
}
// GetManagementTemplateCollections gets the managementTemplateCollections property value. 
func (m *ManagedTenant) GetManagementTemplateCollections()([]ManagementTemplateCollection) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateCollections
    }
}
// GetManagementTemplates gets the managementTemplates property value. The collection of baseline management templates across managed tenants.
func (m *ManagedTenant) GetManagementTemplates()([]ManagementTemplate) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplates
    }
}
// GetManagementTemplateSteps gets the managementTemplateSteps property value. 
func (m *ManagedTenant) GetManagementTemplateSteps()([]ManagementTemplateStep) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateSteps
    }
}
// GetManagementTemplateStepVersions gets the managementTemplateStepVersions property value. 
func (m *ManagedTenant) GetManagementTemplateStepVersions()([]ManagementTemplateStepVersion) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateStepVersions
    }
}
// GetTenantGroups gets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *ManagedTenant) GetTenantGroups()([]TenantGroup) {
    if m == nil {
        return nil
    } else {
        return m.tenantGroups
    }
}
// GetTenants gets the tenants property value. The collection of tenants associated with the managing entity.
func (m *ManagedTenant) GetTenants()([]Tenant) {
    if m == nil {
        return nil
    } else {
        return m.tenants
    }
}
// GetTenantsCustomizedInformation gets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
func (m *ManagedTenant) GetTenantsCustomizedInformation()([]TenantCustomizedInformation) {
    if m == nil {
        return nil
    } else {
        return m.tenantsCustomizedInformation
    }
}
// GetTenantsDetailedInformation gets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
func (m *ManagedTenant) GetTenantsDetailedInformation()([]TenantDetailedInformation) {
    if m == nil {
        return nil
    } else {
        return m.tenantsDetailedInformation
    }
}
// GetTenantTags gets the tenantTags property value. The collection of tenant tags across managed tenants.
func (m *ManagedTenant) GetTenantTags()([]TenantTag) {
    if m == nil {
        return nil
    } else {
        return m.tenantTags
    }
}
// GetWindowsDeviceMalwareStates gets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsDeviceMalwareStates()([]WindowsDeviceMalwareState) {
    if m == nil {
        return nil
    } else {
        return m.windowsDeviceMalwareStates
    }
}
// GetWindowsProtectionStates gets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsProtectionStates()([]WindowsProtectionState) {
    if m == nil {
        return nil
    } else {
        return m.windowsProtectionStates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aggregatedPolicyCompliances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAggregatedPolicyCompliance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AggregatedPolicyCompliance, len(val))
            for i, v := range val {
                res[i] = *(v.(*AggregatedPolicyCompliance))
            }
            m.SetAggregatedPolicyCompliances(res)
        }
        return nil
    }
    res["auditEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuditEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditEvent, len(val))
            for i, v := range val {
                res[i] = *(v.(*AuditEvent))
            }
            m.SetAuditEvents(res)
        }
        return nil
    }
    res["cloudPcConnections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcConnection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcConnection, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcConnection))
            }
            m.SetCloudPcConnections(res)
        }
        return nil
    }
    res["cloudPcDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcDevice() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcDevice, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcDevice))
            }
            m.SetCloudPcDevices(res)
        }
        return nil
    }
    res["cloudPcsOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcOverview() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcOverview, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcOverview))
            }
            m.SetCloudPcsOverview(res)
        }
        return nil
    }
    res["conditionalAccessPolicyCoverages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessPolicyCoverage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessPolicyCoverage, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConditionalAccessPolicyCoverage))
            }
            m.SetConditionalAccessPolicyCoverages(res)
        }
        return nil
    }
    res["credentialUserRegistrationsSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCredentialUserRegistrationsSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CredentialUserRegistrationsSummary, len(val))
            for i, v := range val {
                res[i] = *(v.(*CredentialUserRegistrationsSummary))
            }
            m.SetCredentialUserRegistrationsSummaries(res)
        }
        return nil
    }
    res["deviceCompliancePolicySettingStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicySettingStateSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicySettingStateSummary, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceCompliancePolicySettingStateSummary))
            }
            m.SetDeviceCompliancePolicySettingStateSummaries(res)
        }
        return nil
    }
    res["managedDeviceCompliances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceCompliance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceCompliance, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedDeviceCompliance))
            }
            m.SetManagedDeviceCompliances(res)
        }
        return nil
    }
    res["managedDeviceComplianceTrends"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceComplianceTrend() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceComplianceTrend, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedDeviceComplianceTrend))
            }
            m.SetManagedDeviceComplianceTrends(res)
        }
        return nil
    }
    res["managementActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementAction() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementAction, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementAction))
            }
            m.SetManagementActions(res)
        }
        return nil
    }
    res["managementActionTenantDeploymentStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementActionTenantDeploymentStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementActionTenantDeploymentStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementActionTenantDeploymentStatus))
            }
            m.SetManagementActionTenantDeploymentStatuses(res)
        }
        return nil
    }
    res["managementIntents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementIntent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementIntent, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementIntent))
            }
            m.SetManagementIntents(res)
        }
        return nil
    }
    res["managementTemplateCollections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateCollection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateCollection, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplateCollection))
            }
            m.SetManagementTemplateCollections(res)
        }
        return nil
    }
    res["managementTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplate() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplate, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplate))
            }
            m.SetManagementTemplates(res)
        }
        return nil
    }
    res["managementTemplateSteps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateStep() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStep, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplateStep))
            }
            m.SetManagementTemplateSteps(res)
        }
        return nil
    }
    res["managementTemplateStepVersions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateStepVersion() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepVersion, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagementTemplateStepVersion))
            }
            m.SetManagementTemplateStepVersions(res)
        }
        return nil
    }
    res["tenantGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTenantGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*TenantGroup))
            }
            m.SetTenantGroups(res)
        }
        return nil
    }
    res["tenants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTenant() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Tenant, len(val))
            for i, v := range val {
                res[i] = *(v.(*Tenant))
            }
            m.SetTenants(res)
        }
        return nil
    }
    res["tenantsCustomizedInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTenantCustomizedInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantCustomizedInformation, len(val))
            for i, v := range val {
                res[i] = *(v.(*TenantCustomizedInformation))
            }
            m.SetTenantsCustomizedInformation(res)
        }
        return nil
    }
    res["tenantsDetailedInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTenantDetailedInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantDetailedInformation, len(val))
            for i, v := range val {
                res[i] = *(v.(*TenantDetailedInformation))
            }
            m.SetTenantsDetailedInformation(res)
        }
        return nil
    }
    res["tenantTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTenantTag() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantTag, len(val))
            for i, v := range val {
                res[i] = *(v.(*TenantTag))
            }
            m.SetTenantTags(res)
        }
        return nil
    }
    res["windowsDeviceMalwareStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDeviceMalwareState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDeviceMalwareState, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsDeviceMalwareState))
            }
            m.SetWindowsDeviceMalwareStates(res)
        }
        return nil
    }
    res["windowsProtectionStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsProtectionState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsProtectionState, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsProtectionState))
            }
            m.SetWindowsProtectionStates(res)
        }
        return nil
    }
    return res
}
func (m *ManagedTenant) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedTenant) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAggregatedPolicyCompliances() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAggregatedPolicyCompliances()))
        for i, v := range m.GetAggregatedPolicyCompliances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("aggregatedPolicyCompliances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuditEvents() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAuditEvents()))
        for i, v := range m.GetAuditEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcConnections() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCloudPcConnections()))
        for i, v := range m.GetCloudPcConnections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPcConnections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcDevices() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCloudPcDevices()))
        for i, v := range m.GetCloudPcDevices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPcDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcsOverview() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCloudPcsOverview()))
        for i, v := range m.GetCloudPcsOverview() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPcsOverview", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessPolicyCoverages() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConditionalAccessPolicyCoverages()))
        for i, v := range m.GetConditionalAccessPolicyCoverages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("conditionalAccessPolicyCoverages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCredentialUserRegistrationsSummaries() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCredentialUserRegistrationsSummaries()))
        for i, v := range m.GetCredentialUserRegistrationsSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("credentialUserRegistrationsSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicySettingStateSummaries() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceCompliancePolicySettingStateSummaries()))
        for i, v := range m.GetDeviceCompliancePolicySettingStateSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicySettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceCompliances() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedDeviceCompliances()))
        for i, v := range m.GetManagedDeviceCompliances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceCompliances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceComplianceTrends() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedDeviceComplianceTrends()))
        for i, v := range m.GetManagedDeviceComplianceTrends() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceComplianceTrends", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementActions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementActions()))
        for i, v := range m.GetManagementActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementActionTenantDeploymentStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementActionTenantDeploymentStatuses()))
        for i, v := range m.GetManagementActionTenantDeploymentStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementActionTenantDeploymentStatuses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementIntents() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementIntents()))
        for i, v := range m.GetManagementIntents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementIntents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateCollections() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplateCollections()))
        for i, v := range m.GetManagementTemplateCollections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateCollections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplates()))
        for i, v := range m.GetManagementTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateSteps() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplateSteps()))
        for i, v := range m.GetManagementTemplateSteps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateSteps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateStepVersions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplateStepVersions()))
        for i, v := range m.GetManagementTemplateStepVersions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateStepVersions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTenantGroups()))
        for i, v := range m.GetTenantGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tenantGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenants() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTenants()))
        for i, v := range m.GetTenants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tenants", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantsCustomizedInformation() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTenantsCustomizedInformation()))
        for i, v := range m.GetTenantsCustomizedInformation() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tenantsCustomizedInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantsDetailedInformation() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTenantsDetailedInformation()))
        for i, v := range m.GetTenantsDetailedInformation() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tenantsDetailedInformation", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantTags() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTenantTags()))
        for i, v := range m.GetTenantTags() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tenantTags", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsDeviceMalwareStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsDeviceMalwareStates()))
        for i, v := range m.GetWindowsDeviceMalwareStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsDeviceMalwareStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsProtectionStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsProtectionStates()))
        for i, v := range m.GetWindowsProtectionStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsProtectionStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAggregatedPolicyCompliances sets the aggregatedPolicyCompliances property value. Aggregate view of device compliance policies across managed tenants.
func (m *ManagedTenant) SetAggregatedPolicyCompliances(value []AggregatedPolicyCompliance)() {
    if m != nil {
        m.aggregatedPolicyCompliances = value
    }
}
// SetAuditEvents sets the auditEvents property value. 
func (m *ManagedTenant) SetAuditEvents(value []AuditEvent)() {
    if m != nil {
        m.auditEvents = value
    }
}
// SetCloudPcConnections sets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
func (m *ManagedTenant) SetCloudPcConnections(value []CloudPcConnection)() {
    if m != nil {
        m.cloudPcConnections = value
    }
}
// SetCloudPcDevices sets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
func (m *ManagedTenant) SetCloudPcDevices(value []CloudPcDevice)() {
    if m != nil {
        m.cloudPcDevices = value
    }
}
// SetCloudPcsOverview sets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
func (m *ManagedTenant) SetCloudPcsOverview(value []CloudPcOverview)() {
    if m != nil {
        m.cloudPcsOverview = value
    }
}
// SetConditionalAccessPolicyCoverages sets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
func (m *ManagedTenant) SetConditionalAccessPolicyCoverages(value []ConditionalAccessPolicyCoverage)() {
    if m != nil {
        m.conditionalAccessPolicyCoverages = value
    }
}
// SetCredentialUserRegistrationsSummaries sets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *ManagedTenant) SetCredentialUserRegistrationsSummaries(value []CredentialUserRegistrationsSummary)() {
    if m != nil {
        m.credentialUserRegistrationsSummaries = value
    }
}
// SetDeviceCompliancePolicySettingStateSummaries sets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenant) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummary)() {
    if m != nil {
        m.deviceCompliancePolicySettingStateSummaries = value
    }
}
// SetManagedDeviceCompliances sets the managedDeviceCompliances property value. The collection of compliance for managed devices across managed tenants.
func (m *ManagedTenant) SetManagedDeviceCompliances(value []ManagedDeviceCompliance)() {
    if m != nil {
        m.managedDeviceCompliances = value
    }
}
// SetManagedDeviceComplianceTrends sets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
func (m *ManagedTenant) SetManagedDeviceComplianceTrends(value []ManagedDeviceComplianceTrend)() {
    if m != nil {
        m.managedDeviceComplianceTrends = value
    }
}
// SetManagementActions sets the managementActions property value. The collection of baseline management actions across managed tenants.
func (m *ManagedTenant) SetManagementActions(value []ManagementAction)() {
    if m != nil {
        m.managementActions = value
    }
}
// SetManagementActionTenantDeploymentStatuses sets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
func (m *ManagedTenant) SetManagementActionTenantDeploymentStatuses(value []ManagementActionTenantDeploymentStatus)() {
    if m != nil {
        m.managementActionTenantDeploymentStatuses = value
    }
}
// SetManagementIntents sets the managementIntents property value. The collection of baseline management intents across managed tenants.
func (m *ManagedTenant) SetManagementIntents(value []ManagementIntent)() {
    if m != nil {
        m.managementIntents = value
    }
}
// SetManagementTemplateCollections sets the managementTemplateCollections property value. 
func (m *ManagedTenant) SetManagementTemplateCollections(value []ManagementTemplateCollection)() {
    if m != nil {
        m.managementTemplateCollections = value
    }
}
// SetManagementTemplates sets the managementTemplates property value. The collection of baseline management templates across managed tenants.
func (m *ManagedTenant) SetManagementTemplates(value []ManagementTemplate)() {
    if m != nil {
        m.managementTemplates = value
    }
}
// SetManagementTemplateSteps sets the managementTemplateSteps property value. 
func (m *ManagedTenant) SetManagementTemplateSteps(value []ManagementTemplateStep)() {
    if m != nil {
        m.managementTemplateSteps = value
    }
}
// SetManagementTemplateStepVersions sets the managementTemplateStepVersions property value. 
func (m *ManagedTenant) SetManagementTemplateStepVersions(value []ManagementTemplateStepVersion)() {
    if m != nil {
        m.managementTemplateStepVersions = value
    }
}
// SetTenantGroups sets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *ManagedTenant) SetTenantGroups(value []TenantGroup)() {
    if m != nil {
        m.tenantGroups = value
    }
}
// SetTenants sets the tenants property value. The collection of tenants associated with the managing entity.
func (m *ManagedTenant) SetTenants(value []Tenant)() {
    if m != nil {
        m.tenants = value
    }
}
// SetTenantsCustomizedInformation sets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
func (m *ManagedTenant) SetTenantsCustomizedInformation(value []TenantCustomizedInformation)() {
    if m != nil {
        m.tenantsCustomizedInformation = value
    }
}
// SetTenantsDetailedInformation sets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
func (m *ManagedTenant) SetTenantsDetailedInformation(value []TenantDetailedInformation)() {
    if m != nil {
        m.tenantsDetailedInformation = value
    }
}
// SetTenantTags sets the tenantTags property value. The collection of tenant tags across managed tenants.
func (m *ManagedTenant) SetTenantTags(value []TenantTag)() {
    if m != nil {
        m.tenantTags = value
    }
}
// SetWindowsDeviceMalwareStates sets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) SetWindowsDeviceMalwareStates(value []WindowsDeviceMalwareState)() {
    if m != nil {
        m.windowsDeviceMalwareStates = value
    }
}
// SetWindowsProtectionStates sets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) SetWindowsProtectionStates(value []WindowsProtectionState)() {
    if m != nil {
        m.windowsProtectionStates = value
    }
}
