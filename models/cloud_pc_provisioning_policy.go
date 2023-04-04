package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcProvisioningPolicy 
type CloudPcProvisioningPolicy struct {
    Entity
}
// NewCloudPcProvisioningPolicy instantiates a new CloudPcProvisioningPolicy and sets the default values.
func NewCloudPcProvisioningPolicy()(*CloudPcProvisioningPolicy) {
    m := &CloudPcProvisioningPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcProvisioningPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcProvisioningPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcProvisioningPolicy(), nil
}
// GetAlternateResourceUrl gets the alternateResourceUrl property value. The URL of the alternate resource that links to this provisioning policy. Read-only.
func (m *CloudPcProvisioningPolicy) GetAlternateResourceUrl()(*string) {
    val, err := m.GetBackingStore().Get("alternateResourceUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignments gets the assignments property value. A defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Azure AD that have provisioning policy assigned. Returned only on $expand. For an example about how to get the assignments relationship, see Get cloudPcProvisioningPolicy.
func (m *CloudPcProvisioningPolicy) GetAssignments()([]CloudPcProvisioningPolicyAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcProvisioningPolicyAssignmentable)
    }
    return nil
}
// GetCloudPcGroupDisplayName gets the cloudPcGroupDisplayName property value. The display name of the Cloud PC group that the Cloud PCs reside in. Read-only.
func (m *CloudPcProvisioningPolicy) GetCloudPcGroupDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("cloudPcGroupDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCloudPcNamingTemplate gets the cloudPcNamingTemplate property value. The cloudPcNamingTemplate property
func (m *CloudPcProvisioningPolicy) GetCloudPcNamingTemplate()(*string) {
    val, err := m.GetBackingStore().Get("cloudPcNamingTemplate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The provisioning policy description.
func (m *CloudPcProvisioningPolicy) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the provisioning policy.
func (m *CloudPcProvisioningPolicy) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDomainJoinConfiguration gets the domainJoinConfiguration property value. Specifies how Cloud PCs will join Azure Active Directory.
func (m *CloudPcProvisioningPolicy) GetDomainJoinConfiguration()(CloudPcDomainJoinConfigurationable) {
    val, err := m.GetBackingStore().Get("domainJoinConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcDomainJoinConfigurationable)
    }
    return nil
}
// GetDomainJoinConfigurations gets the domainJoinConfigurations property value. The domainJoinConfigurations property
func (m *CloudPcProvisioningPolicy) GetDomainJoinConfigurations()([]CloudPcDomainJoinConfigurationable) {
    val, err := m.GetBackingStore().Get("domainJoinConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcDomainJoinConfigurationable)
    }
    return nil
}
// GetEnableSingleSignOn gets the enableSingleSignOn property value. True if the provisioned Cloud PC can be accessed by single sign-on. False indicates that the provisioned Cloud PC doesn't support this feature. Default value is false. Windows 365 users can use single sign-on to authenticate to Azure Active Directory (Azure AD) with passwordless options (for example, FIDO keys) to access their Cloud PC. Optional.
func (m *CloudPcProvisioningPolicy) GetEnableSingleSignOn()(*bool) {
    val, err := m.GetBackingStore().Get("enableSingleSignOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcProvisioningPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alternateResourceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateResourceUrl(val)
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcProvisioningPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcProvisioningPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcProvisioningPolicyAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["cloudPcGroupDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcGroupDisplayName(val)
        }
        return nil
    }
    res["cloudPcNamingTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcNamingTemplate(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["domainJoinConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcDomainJoinConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainJoinConfiguration(val.(CloudPcDomainJoinConfigurationable))
        }
        return nil
    }
    res["domainJoinConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcDomainJoinConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcDomainJoinConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcDomainJoinConfigurationable)
            }
            m.SetDomainJoinConfigurations(res)
        }
        return nil
    }
    res["enableSingleSignOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSingleSignOn(val)
        }
        return nil
    }
    res["gracePeriodInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodInHours(val)
        }
        return nil
    }
    res["imageDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageDisplayName(val)
        }
        return nil
    }
    res["imageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageId(val)
        }
        return nil
    }
    res["imageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcProvisioningPolicyImageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageType(val.(*CloudPcProvisioningPolicyImageType))
        }
        return nil
    }
    res["localAdminEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalAdminEnabled(val)
        }
        return nil
    }
    res["managedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcManagementService)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBy(val.(*CloudPcManagementService))
        }
        return nil
    }
    res["microsoftManagedDesktop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMicrosoftManagedDesktopFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftManagedDesktop(val.(MicrosoftManagedDesktopable))
        }
        return nil
    }
    res["onPremisesConnectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesConnectionId(val)
        }
        return nil
    }
    res["provisioningType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcProvisioningType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningType(val.(*CloudPcProvisioningType))
        }
        return nil
    }
    res["windowsSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcWindowsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSettings(val.(CloudPcWindowsSettingsable))
        }
        return nil
    }
    return res
}
// GetGracePeriodInHours gets the gracePeriodInHours property value. The number of hours to wait before reprovisioning/deprovisioning happens. Read-only.
func (m *CloudPcProvisioningPolicy) GetGracePeriodInHours()(*int32) {
    val, err := m.GetBackingStore().Get("gracePeriodInHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetImageDisplayName gets the imageDisplayName property value. The display name for the OS image you’re provisioning.
func (m *CloudPcProvisioningPolicy) GetImageDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("imageDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetImageId gets the imageId property value. The ID of the OS image you want to provision on Cloud PCs. The format for a gallery type image is: {publisher_offer_sku}. Supported values for each of the parameters are as follows:publisher: Microsoftwindowsdesktop. offer: windows-ent-cpc. sku: 21h1-ent-cpc-m365, 21h1-ent-cpc-os, 20h2-ent-cpc-m365, 20h2-ent-cpc-os, 20h1-ent-cpc-m365, 20h1-ent-cpc-os, 19h2-ent-cpc-m365 and 19h2-ent-cpc-os.
func (m *CloudPcProvisioningPolicy) GetImageId()(*string) {
    val, err := m.GetBackingStore().Get("imageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetImageType gets the imageType property value. The imageType property
func (m *CloudPcProvisioningPolicy) GetImageType()(*CloudPcProvisioningPolicyImageType) {
    val, err := m.GetBackingStore().Get("imageType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcProvisioningPolicyImageType)
    }
    return nil
}
// GetLocalAdminEnabled gets the localAdminEnabled property value. Indicates whether the local admin option is enabled. If the local admin option is enabled, the end user can be an admin of the Cloud PC device. Read-only.
func (m *CloudPcProvisioningPolicy) GetLocalAdminEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("localAdminEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManagedBy gets the managedBy property value. The managedBy property
func (m *CloudPcProvisioningPolicy) GetManagedBy()(*CloudPcManagementService) {
    val, err := m.GetBackingStore().Get("managedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcManagementService)
    }
    return nil
}
// GetMicrosoftManagedDesktop gets the microsoftManagedDesktop property value. The specific settings for the Microsoft Managed Desktop, which enables customers to get a managed device experience for the Cloud PC. Before you can enable Microsoft Managed Desktop, an admin must configure it.
func (m *CloudPcProvisioningPolicy) GetMicrosoftManagedDesktop()(MicrosoftManagedDesktopable) {
    val, err := m.GetBackingStore().Get("microsoftManagedDesktop")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MicrosoftManagedDesktopable)
    }
    return nil
}
// GetOnPremisesConnectionId gets the onPremisesConnectionId property value. The ID of the cloudPcOnPremisesConnection. To ensure that Cloud PCs have network connectivity and that they domain join, choose a connection with a virtual network that’s validated by the Cloud PC service.
func (m *CloudPcProvisioningPolicy) GetOnPremisesConnectionId()(*string) {
    val, err := m.GetBackingStore().Get("onPremisesConnectionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProvisioningType gets the provisioningType property value. Specifies the type of license used when provisioning Cloud PCs using this policy. By default, the license type is dedicated if the provisioningType isn't specified when you create the cloudPcProvisioningPolicy. You can't change this property after the cloudPcProvisioningPolicy was created. Possible values are: dedicated, shared, unknownFutureValue.
func (m *CloudPcProvisioningPolicy) GetProvisioningType()(*CloudPcProvisioningType) {
    val, err := m.GetBackingStore().Get("provisioningType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcProvisioningType)
    }
    return nil
}
// GetWindowsSettings gets the windowsSettings property value. Specific Windows settings to configure while creating Cloud PCs for this provisioning policy.
func (m *CloudPcProvisioningPolicy) GetWindowsSettings()(CloudPcWindowsSettingsable) {
    val, err := m.GetBackingStore().Get("windowsSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcWindowsSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcProvisioningPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("alternateResourceUrl", m.GetAlternateResourceUrl())
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cloudPcGroupDisplayName", m.GetCloudPcGroupDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cloudPcNamingTemplate", m.GetCloudPcNamingTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("domainJoinConfiguration", m.GetDomainJoinConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetDomainJoinConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDomainJoinConfigurations()))
        for i, v := range m.GetDomainJoinConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("domainJoinConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSingleSignOn", m.GetEnableSingleSignOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("gracePeriodInHours", m.GetGracePeriodInHours())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("imageDisplayName", m.GetImageDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("imageId", m.GetImageId())
        if err != nil {
            return err
        }
    }
    if m.GetImageType() != nil {
        cast := (*m.GetImageType()).String()
        err = writer.WriteStringValue("imageType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("localAdminEnabled", m.GetLocalAdminEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetManagedBy() != nil {
        cast := (*m.GetManagedBy()).String()
        err = writer.WriteStringValue("managedBy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("microsoftManagedDesktop", m.GetMicrosoftManagedDesktop())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesConnectionId", m.GetOnPremisesConnectionId())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningType() != nil {
        cast := (*m.GetProvisioningType()).String()
        err = writer.WriteStringValue("provisioningType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsSettings", m.GetWindowsSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlternateResourceUrl sets the alternateResourceUrl property value. The URL of the alternate resource that links to this provisioning policy. Read-only.
func (m *CloudPcProvisioningPolicy) SetAlternateResourceUrl(value *string)() {
    err := m.GetBackingStore().Set("alternateResourceUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignments sets the assignments property value. A defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Azure AD that have provisioning policy assigned. Returned only on $expand. For an example about how to get the assignments relationship, see Get cloudPcProvisioningPolicy.
func (m *CloudPcProvisioningPolicy) SetAssignments(value []CloudPcProvisioningPolicyAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcGroupDisplayName sets the cloudPcGroupDisplayName property value. The display name of the Cloud PC group that the Cloud PCs reside in. Read-only.
func (m *CloudPcProvisioningPolicy) SetCloudPcGroupDisplayName(value *string)() {
    err := m.GetBackingStore().Set("cloudPcGroupDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcNamingTemplate sets the cloudPcNamingTemplate property value. The cloudPcNamingTemplate property
func (m *CloudPcProvisioningPolicy) SetCloudPcNamingTemplate(value *string)() {
    err := m.GetBackingStore().Set("cloudPcNamingTemplate", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The provisioning policy description.
func (m *CloudPcProvisioningPolicy) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the provisioning policy.
func (m *CloudPcProvisioningPolicy) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainJoinConfiguration sets the domainJoinConfiguration property value. Specifies how Cloud PCs will join Azure Active Directory.
func (m *CloudPcProvisioningPolicy) SetDomainJoinConfiguration(value CloudPcDomainJoinConfigurationable)() {
    err := m.GetBackingStore().Set("domainJoinConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainJoinConfigurations sets the domainJoinConfigurations property value. The domainJoinConfigurations property
func (m *CloudPcProvisioningPolicy) SetDomainJoinConfigurations(value []CloudPcDomainJoinConfigurationable)() {
    err := m.GetBackingStore().Set("domainJoinConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSingleSignOn sets the enableSingleSignOn property value. True if the provisioned Cloud PC can be accessed by single sign-on. False indicates that the provisioned Cloud PC doesn't support this feature. Default value is false. Windows 365 users can use single sign-on to authenticate to Azure Active Directory (Azure AD) with passwordless options (for example, FIDO keys) to access their Cloud PC. Optional.
func (m *CloudPcProvisioningPolicy) SetEnableSingleSignOn(value *bool)() {
    err := m.GetBackingStore().Set("enableSingleSignOn", value)
    if err != nil {
        panic(err)
    }
}
// SetGracePeriodInHours sets the gracePeriodInHours property value. The number of hours to wait before reprovisioning/deprovisioning happens. Read-only.
func (m *CloudPcProvisioningPolicy) SetGracePeriodInHours(value *int32)() {
    err := m.GetBackingStore().Set("gracePeriodInHours", value)
    if err != nil {
        panic(err)
    }
}
// SetImageDisplayName sets the imageDisplayName property value. The display name for the OS image you’re provisioning.
func (m *CloudPcProvisioningPolicy) SetImageDisplayName(value *string)() {
    err := m.GetBackingStore().Set("imageDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetImageId sets the imageId property value. The ID of the OS image you want to provision on Cloud PCs. The format for a gallery type image is: {publisher_offer_sku}. Supported values for each of the parameters are as follows:publisher: Microsoftwindowsdesktop. offer: windows-ent-cpc. sku: 21h1-ent-cpc-m365, 21h1-ent-cpc-os, 20h2-ent-cpc-m365, 20h2-ent-cpc-os, 20h1-ent-cpc-m365, 20h1-ent-cpc-os, 19h2-ent-cpc-m365 and 19h2-ent-cpc-os.
func (m *CloudPcProvisioningPolicy) SetImageId(value *string)() {
    err := m.GetBackingStore().Set("imageId", value)
    if err != nil {
        panic(err)
    }
}
// SetImageType sets the imageType property value. The imageType property
func (m *CloudPcProvisioningPolicy) SetImageType(value *CloudPcProvisioningPolicyImageType)() {
    err := m.GetBackingStore().Set("imageType", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalAdminEnabled sets the localAdminEnabled property value. Indicates whether the local admin option is enabled. If the local admin option is enabled, the end user can be an admin of the Cloud PC device. Read-only.
func (m *CloudPcProvisioningPolicy) SetLocalAdminEnabled(value *bool)() {
    err := m.GetBackingStore().Set("localAdminEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedBy sets the managedBy property value. The managedBy property
func (m *CloudPcProvisioningPolicy) SetManagedBy(value *CloudPcManagementService)() {
    err := m.GetBackingStore().Set("managedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftManagedDesktop sets the microsoftManagedDesktop property value. The specific settings for the Microsoft Managed Desktop, which enables customers to get a managed device experience for the Cloud PC. Before you can enable Microsoft Managed Desktop, an admin must configure it.
func (m *CloudPcProvisioningPolicy) SetMicrosoftManagedDesktop(value MicrosoftManagedDesktopable)() {
    err := m.GetBackingStore().Set("microsoftManagedDesktop", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremisesConnectionId sets the onPremisesConnectionId property value. The ID of the cloudPcOnPremisesConnection. To ensure that Cloud PCs have network connectivity and that they domain join, choose a connection with a virtual network that’s validated by the Cloud PC service.
func (m *CloudPcProvisioningPolicy) SetOnPremisesConnectionId(value *string)() {
    err := m.GetBackingStore().Set("onPremisesConnectionId", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisioningType sets the provisioningType property value. Specifies the type of license used when provisioning Cloud PCs using this policy. By default, the license type is dedicated if the provisioningType isn't specified when you create the cloudPcProvisioningPolicy. You can't change this property after the cloudPcProvisioningPolicy was created. Possible values are: dedicated, shared, unknownFutureValue.
func (m *CloudPcProvisioningPolicy) SetProvisioningType(value *CloudPcProvisioningType)() {
    err := m.GetBackingStore().Set("provisioningType", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsSettings sets the windowsSettings property value. Specific Windows settings to configure while creating Cloud PCs for this provisioning policy.
func (m *CloudPcProvisioningPolicy) SetWindowsSettings(value CloudPcWindowsSettingsable)() {
    err := m.GetBackingStore().Set("windowsSettings", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcProvisioningPolicyable 
type CloudPcProvisioningPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlternateResourceUrl()(*string)
    GetAssignments()([]CloudPcProvisioningPolicyAssignmentable)
    GetCloudPcGroupDisplayName()(*string)
    GetCloudPcNamingTemplate()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDomainJoinConfiguration()(CloudPcDomainJoinConfigurationable)
    GetDomainJoinConfigurations()([]CloudPcDomainJoinConfigurationable)
    GetEnableSingleSignOn()(*bool)
    GetGracePeriodInHours()(*int32)
    GetImageDisplayName()(*string)
    GetImageId()(*string)
    GetImageType()(*CloudPcProvisioningPolicyImageType)
    GetLocalAdminEnabled()(*bool)
    GetManagedBy()(*CloudPcManagementService)
    GetMicrosoftManagedDesktop()(MicrosoftManagedDesktopable)
    GetOnPremisesConnectionId()(*string)
    GetProvisioningType()(*CloudPcProvisioningType)
    GetWindowsSettings()(CloudPcWindowsSettingsable)
    SetAlternateResourceUrl(value *string)()
    SetAssignments(value []CloudPcProvisioningPolicyAssignmentable)()
    SetCloudPcGroupDisplayName(value *string)()
    SetCloudPcNamingTemplate(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDomainJoinConfiguration(value CloudPcDomainJoinConfigurationable)()
    SetDomainJoinConfigurations(value []CloudPcDomainJoinConfigurationable)()
    SetEnableSingleSignOn(value *bool)()
    SetGracePeriodInHours(value *int32)()
    SetImageDisplayName(value *string)()
    SetImageId(value *string)()
    SetImageType(value *CloudPcProvisioningPolicyImageType)()
    SetLocalAdminEnabled(value *bool)()
    SetManagedBy(value *CloudPcManagementService)()
    SetMicrosoftManagedDesktop(value MicrosoftManagedDesktopable)()
    SetOnPremisesConnectionId(value *string)()
    SetProvisioningType(value *CloudPcProvisioningType)()
    SetWindowsSettings(value CloudPcWindowsSettingsable)()
}
