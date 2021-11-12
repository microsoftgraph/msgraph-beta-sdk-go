package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedTenant struct {
    Entity
    // Aggregate view of device compliance policies across managed tenants.
    aggregatedPolicyCompliances []AggregatedPolicyCompliance;
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
    // The collection of baseline management templates across managed tenants.
    managementTemplates []ManagementTemplate;
    // The collection of users flagged for risk across managed tenants.
    riskyUsers []RiskyUser;
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
// Instantiates a new managedTenant and sets the default values.
func NewManagedTenant()(*ManagedTenant) {
    m := &ManagedTenant{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the aggregatedPolicyCompliances property value. Aggregate view of device compliance policies across managed tenants.
func (m *ManagedTenant) GetAggregatedPolicyCompliances()([]AggregatedPolicyCompliance) {
    if m == nil {
        return nil
    } else {
        return m.aggregatedPolicyCompliances
    }
}
// Gets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
func (m *ManagedTenant) GetCloudPcConnections()([]CloudPcConnection) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcConnections
    }
}
// Gets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
func (m *ManagedTenant) GetCloudPcDevices()([]CloudPcDevice) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcDevices
    }
}
// Gets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
func (m *ManagedTenant) GetCloudPcsOverview()([]CloudPcOverview) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcsOverview
    }
}
// Gets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
func (m *ManagedTenant) GetConditionalAccessPolicyCoverages()([]ConditionalAccessPolicyCoverage) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessPolicyCoverages
    }
}
// Gets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
func (m *ManagedTenant) GetCredentialUserRegistrationsSummaries()([]CredentialUserRegistrationsSummary) {
    if m == nil {
        return nil
    } else {
        return m.credentialUserRegistrationsSummaries
    }
}
// Gets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenant) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicySettingStateSummaries
    }
}
// Gets the managedDeviceCompliances property value. The collection of compliance for managed devices across managed tenants.
func (m *ManagedTenant) GetManagedDeviceCompliances()([]ManagedDeviceCompliance) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceCompliances
    }
}
// Gets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
func (m *ManagedTenant) GetManagedDeviceComplianceTrends()([]ManagedDeviceComplianceTrend) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceComplianceTrends
    }
}
// Gets the managementActions property value. The collection of baseline management actions across managed tenants.
func (m *ManagedTenant) GetManagementActions()([]ManagementAction) {
    if m == nil {
        return nil
    } else {
        return m.managementActions
    }
}
// Gets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
func (m *ManagedTenant) GetManagementActionTenantDeploymentStatuses()([]ManagementActionTenantDeploymentStatus) {
    if m == nil {
        return nil
    } else {
        return m.managementActionTenantDeploymentStatuses
    }
}
// Gets the managementIntents property value. The collection of baseline management intents across managed tenants.
func (m *ManagedTenant) GetManagementIntents()([]ManagementIntent) {
    if m == nil {
        return nil
    } else {
        return m.managementIntents
    }
}
// Gets the managementTemplates property value. The collection of baseline management templates across managed tenants.
func (m *ManagedTenant) GetManagementTemplates()([]ManagementTemplate) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplates
    }
}
// Gets the riskyUsers property value. The collection of users flagged for risk across managed tenants.
func (m *ManagedTenant) GetRiskyUsers()([]RiskyUser) {
    if m == nil {
        return nil
    } else {
        return m.riskyUsers
    }
}
// Gets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
func (m *ManagedTenant) GetTenantGroups()([]TenantGroup) {
    if m == nil {
        return nil
    } else {
        return m.tenantGroups
    }
}
// Gets the tenants property value. The collection of tenants associated with the managing entity.
func (m *ManagedTenant) GetTenants()([]Tenant) {
    if m == nil {
        return nil
    } else {
        return m.tenants
    }
}
// Gets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
func (m *ManagedTenant) GetTenantsCustomizedInformation()([]TenantCustomizedInformation) {
    if m == nil {
        return nil
    } else {
        return m.tenantsCustomizedInformation
    }
}
// Gets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
func (m *ManagedTenant) GetTenantsDetailedInformation()([]TenantDetailedInformation) {
    if m == nil {
        return nil
    } else {
        return m.tenantsDetailedInformation
    }
}
// Gets the tenantTags property value. The collection of tenant tags across managed tenants.
func (m *ManagedTenant) GetTenantTags()([]TenantTag) {
    if m == nil {
        return nil
    } else {
        return m.tenantTags
    }
}
// Gets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsDeviceMalwareStates()([]WindowsDeviceMalwareState) {
    if m == nil {
        return nil
    } else {
        return m.windowsDeviceMalwareStates
    }
}
// Gets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
func (m *ManagedTenant) GetWindowsProtectionStates()([]WindowsProtectionState) {
    if m == nil {
        return nil
    } else {
        return m.windowsProtectionStates
    }
}
// The deserialization information for the current model
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
    res["riskyUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRiskyUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskyUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*RiskyUser))
            }
            m.SetRiskyUsers(res)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedTenant) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRiskyUsers()))
        for i, v := range m.GetRiskyUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("riskyUsers", cast)
        if err != nil {
            return err
        }
    }
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
// Sets the aggregatedPolicyCompliances property value. Aggregate view of device compliance policies across managed tenants.
// Parameters:
//  - value : Value to set for the aggregatedPolicyCompliances property.
func (m *ManagedTenant) SetAggregatedPolicyCompliances(value []AggregatedPolicyCompliance)() {
    m.aggregatedPolicyCompliances = value
}
// Sets the cloudPcConnections property value. The collection of cloud PC connections across managed tenants.
// Parameters:
//  - value : Value to set for the cloudPcConnections property.
func (m *ManagedTenant) SetCloudPcConnections(value []CloudPcConnection)() {
    m.cloudPcConnections = value
}
// Sets the cloudPcDevices property value. The collection of cloud PC devices across managed tenants.
// Parameters:
//  - value : Value to set for the cloudPcDevices property.
func (m *ManagedTenant) SetCloudPcDevices(value []CloudPcDevice)() {
    m.cloudPcDevices = value
}
// Sets the cloudPcsOverview property value. Overview of cloud PC information across managed tenants.
// Parameters:
//  - value : Value to set for the cloudPcsOverview property.
func (m *ManagedTenant) SetCloudPcsOverview(value []CloudPcOverview)() {
    m.cloudPcsOverview = value
}
// Sets the conditionalAccessPolicyCoverages property value. Aggregate view of conditional access policy coverage across managed tenants.
// Parameters:
//  - value : Value to set for the conditionalAccessPolicyCoverages property.
func (m *ManagedTenant) SetConditionalAccessPolicyCoverages(value []ConditionalAccessPolicyCoverage)() {
    m.conditionalAccessPolicyCoverages = value
}
// Sets the credentialUserRegistrationsSummaries property value. Summary information for user registration for multi-factor authentication and self service password reset across managed tenants.
// Parameters:
//  - value : Value to set for the credentialUserRegistrationsSummaries property.
func (m *ManagedTenant) SetCredentialUserRegistrationsSummaries(value []CredentialUserRegistrationsSummary)() {
    m.credentialUserRegistrationsSummaries = value
}
// Sets the deviceCompliancePolicySettingStateSummaries property value. Summary information for device compliance policy setting states across managed tenants.
// Parameters:
//  - value : Value to set for the deviceCompliancePolicySettingStateSummaries property.
func (m *ManagedTenant) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummary)() {
    m.deviceCompliancePolicySettingStateSummaries = value
}
// Sets the managedDeviceCompliances property value. The collection of compliance for managed devices across managed tenants.
// Parameters:
//  - value : Value to set for the managedDeviceCompliances property.
func (m *ManagedTenant) SetManagedDeviceCompliances(value []ManagedDeviceCompliance)() {
    m.managedDeviceCompliances = value
}
// Sets the managedDeviceComplianceTrends property value. Trend insights for device compliance across managed tenants.
// Parameters:
//  - value : Value to set for the managedDeviceComplianceTrends property.
func (m *ManagedTenant) SetManagedDeviceComplianceTrends(value []ManagedDeviceComplianceTrend)() {
    m.managedDeviceComplianceTrends = value
}
// Sets the managementActions property value. The collection of baseline management actions across managed tenants.
// Parameters:
//  - value : Value to set for the managementActions property.
func (m *ManagedTenant) SetManagementActions(value []ManagementAction)() {
    m.managementActions = value
}
// Sets the managementActionTenantDeploymentStatuses property value. The tenant level status of management actions across managed tenants.
// Parameters:
//  - value : Value to set for the managementActionTenantDeploymentStatuses property.
func (m *ManagedTenant) SetManagementActionTenantDeploymentStatuses(value []ManagementActionTenantDeploymentStatus)() {
    m.managementActionTenantDeploymentStatuses = value
}
// Sets the managementIntents property value. The collection of baseline management intents across managed tenants.
// Parameters:
//  - value : Value to set for the managementIntents property.
func (m *ManagedTenant) SetManagementIntents(value []ManagementIntent)() {
    m.managementIntents = value
}
// Sets the managementTemplates property value. The collection of baseline management templates across managed tenants.
// Parameters:
//  - value : Value to set for the managementTemplates property.
func (m *ManagedTenant) SetManagementTemplates(value []ManagementTemplate)() {
    m.managementTemplates = value
}
// Sets the riskyUsers property value. The collection of users flagged for risk across managed tenants.
// Parameters:
//  - value : Value to set for the riskyUsers property.
func (m *ManagedTenant) SetRiskyUsers(value []RiskyUser)() {
    m.riskyUsers = value
}
// Sets the tenantGroups property value. The collection of a logical grouping of managed tenants used by the multi-tenant management platform.
// Parameters:
//  - value : Value to set for the tenantGroups property.
func (m *ManagedTenant) SetTenantGroups(value []TenantGroup)() {
    m.tenantGroups = value
}
// Sets the tenants property value. The collection of tenants associated with the managing entity.
// Parameters:
//  - value : Value to set for the tenants property.
func (m *ManagedTenant) SetTenants(value []Tenant)() {
    m.tenants = value
}
// Sets the tenantsCustomizedInformation property value. The collection of tenant level customized information across managed tenants.
// Parameters:
//  - value : Value to set for the tenantsCustomizedInformation property.
func (m *ManagedTenant) SetTenantsCustomizedInformation(value []TenantCustomizedInformation)() {
    m.tenantsCustomizedInformation = value
}
// Sets the tenantsDetailedInformation property value. The collection tenant level detailed information across managed tenants.
// Parameters:
//  - value : Value to set for the tenantsDetailedInformation property.
func (m *ManagedTenant) SetTenantsDetailedInformation(value []TenantDetailedInformation)() {
    m.tenantsDetailedInformation = value
}
// Sets the tenantTags property value. The collection of tenant tags across managed tenants.
// Parameters:
//  - value : Value to set for the tenantTags property.
func (m *ManagedTenant) SetTenantTags(value []TenantTag)() {
    m.tenantTags = value
}
// Sets the windowsDeviceMalwareStates property value. The state of malware for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
// Parameters:
//  - value : Value to set for the windowsDeviceMalwareStates property.
func (m *ManagedTenant) SetWindowsDeviceMalwareStates(value []WindowsDeviceMalwareState)() {
    m.windowsDeviceMalwareStates = value
}
// Sets the windowsProtectionStates property value. The protection state for Windows devices, registered with Microsoft Endpoint Manager, across managed tenants.
// Parameters:
//  - value : Value to set for the windowsProtectionStates property.
func (m *ManagedTenant) SetWindowsProtectionStates(value []WindowsProtectionState)() {
    m.windowsProtectionStates = value
}
