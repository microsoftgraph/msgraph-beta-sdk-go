package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcProvisioningPolicy provides operations to manage the deviceManagement singleton.
type CloudPcProvisioningPolicy struct {
    Entity
    // The URL of the alternate resource that links to this provisioning policy. Read-only.
    alternateResourceUrl *string
    // A defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Azure AD that have provisioning policy assigned. Returned only on $expand. See an example of getting the assignments relationship.
    assignments []CloudPcProvisioningPolicyAssignmentable
    // The display name of the Cloud PC group that the Cloud PCs reside in. Read-only.
    cloudPcGroupDisplayName *string
    // The provisioning policy description.
    description *string
    // The display name for the provisioning policy.
    displayName *string
    // Specifies how Cloud PCs will join Azure Active Directory.
    domainJoinConfiguration CloudPcDomainJoinConfigurationable
    // The number of hours to wait before reprovisioning/deprovisioning happens. Read-only.
    gracePeriodInHours *int32
    // The display name for the OS image you’re provisioning.
    imageDisplayName *string
    // The ID of the OS image you want to provision on Cloud PCs. The format for a gallery type image is: {publisher_offer_sku}. Supported values for each of the parameters are as follows:publisher: Microsoftwindowsdesktop. offer: windows-ent-cpc. sku: 21h1-ent-cpc-m365, 21h1-ent-cpc-os, 20h2-ent-cpc-m365, 20h2-ent-cpc-os, 20h1-ent-cpc-m365, 20h1-ent-cpc-os, 19h2-ent-cpc-m365 and 19h2-ent-cpc-os.
    imageId *string
    // The type of OS image (custom or gallery) you want to provision on Cloud PCs. Possible values are: gallery, custom.
    imageType *CloudPcProvisioningPolicyImageType
    // Indicates whether the local admin option is enabled. If the local admin option is enabled, the end user can be an admin of the Cloud PC device. Read-only.
    localAdminEnabled *bool
    // Specifies which services manage the Azure network connection. Possible values are: windows365, devBox, unknownFutureValue. Read-only.
    managedBy *CloudPcManagementService
    // The specific settings for the Microsoft Managed Desktop, which enables customers to get a managed device experience for the Cloud PC. Before you can enable Microsoft Managed Desktop, an admin must configure it.
    microsoftManagedDesktop MicrosoftManagedDesktopable
    // The ID of the cloudPcOnPremisesConnection. To ensure that Cloud PCs have network connectivity and that they domain join, choose a connection with a virtual network that’s validated by the Cloud PC service.
    onPremisesConnectionId *string
    // Specific Windows settings to configure while creating Cloud PCs for this provisioning policy.
    windowsSettings CloudPcWindowsSettingsable
}
// NewCloudPcProvisioningPolicy instantiates a new cloudPcProvisioningPolicy and sets the default values.
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
    if m == nil {
        return nil
    } else {
        return m.alternateResourceUrl
    }
}
// GetAssignments gets the assignments property value. A defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Azure AD that have provisioning policy assigned. Returned only on $expand. See an example of getting the assignments relationship.
func (m *CloudPcProvisioningPolicy) GetAssignments()([]CloudPcProvisioningPolicyAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCloudPcGroupDisplayName gets the cloudPcGroupDisplayName property value. The display name of the Cloud PC group that the Cloud PCs reside in. Read-only.
func (m *CloudPcProvisioningPolicy) GetCloudPcGroupDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcGroupDisplayName
    }
}
// GetDescription gets the description property value. The provisioning policy description.
func (m *CloudPcProvisioningPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for the provisioning policy.
func (m *CloudPcProvisioningPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDomainJoinConfiguration gets the domainJoinConfiguration property value. Specifies how Cloud PCs will join Azure Active Directory.
func (m *CloudPcProvisioningPolicy) GetDomainJoinConfiguration()(CloudPcDomainJoinConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.domainJoinConfiguration
    }
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
    if m == nil {
        return nil
    } else {
        return m.gracePeriodInHours
    }
}
// GetImageDisplayName gets the imageDisplayName property value. The display name for the OS image you’re provisioning.
func (m *CloudPcProvisioningPolicy) GetImageDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.imageDisplayName
    }
}
// GetImageId gets the imageId property value. The ID of the OS image you want to provision on Cloud PCs. The format for a gallery type image is: {publisher_offer_sku}. Supported values for each of the parameters are as follows:publisher: Microsoftwindowsdesktop. offer: windows-ent-cpc. sku: 21h1-ent-cpc-m365, 21h1-ent-cpc-os, 20h2-ent-cpc-m365, 20h2-ent-cpc-os, 20h1-ent-cpc-m365, 20h1-ent-cpc-os, 19h2-ent-cpc-m365 and 19h2-ent-cpc-os.
func (m *CloudPcProvisioningPolicy) GetImageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.imageId
    }
}
// GetImageType gets the imageType property value. The type of OS image (custom or gallery) you want to provision on Cloud PCs. Possible values are: gallery, custom.
func (m *CloudPcProvisioningPolicy) GetImageType()(*CloudPcProvisioningPolicyImageType) {
    if m == nil {
        return nil
    } else {
        return m.imageType
    }
}
// GetLocalAdminEnabled gets the localAdminEnabled property value. Indicates whether the local admin option is enabled. If the local admin option is enabled, the end user can be an admin of the Cloud PC device. Read-only.
func (m *CloudPcProvisioningPolicy) GetLocalAdminEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.localAdminEnabled
    }
}
// GetManagedBy gets the managedBy property value. Specifies which services manage the Azure network connection. Possible values are: windows365, devBox, unknownFutureValue. Read-only.
func (m *CloudPcProvisioningPolicy) GetManagedBy()(*CloudPcManagementService) {
    if m == nil {
        return nil
    } else {
        return m.managedBy
    }
}
// GetMicrosoftManagedDesktop gets the microsoftManagedDesktop property value. The specific settings for the Microsoft Managed Desktop, which enables customers to get a managed device experience for the Cloud PC. Before you can enable Microsoft Managed Desktop, an admin must configure it.
func (m *CloudPcProvisioningPolicy) GetMicrosoftManagedDesktop()(MicrosoftManagedDesktopable) {
    if m == nil {
        return nil
    } else {
        return m.microsoftManagedDesktop
    }
}
// GetOnPremisesConnectionId gets the onPremisesConnectionId property value. The ID of the cloudPcOnPremisesConnection. To ensure that Cloud PCs have network connectivity and that they domain join, choose a connection with a virtual network that’s validated by the Cloud PC service.
func (m *CloudPcProvisioningPolicy) GetOnPremisesConnectionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesConnectionId
    }
}
// GetWindowsSettings gets the windowsSettings property value. Specific Windows settings to configure while creating Cloud PCs for this provisioning policy.
func (m *CloudPcProvisioningPolicy) GetWindowsSettings()(CloudPcWindowsSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.windowsSettings
    }
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
    if m != nil {
        m.alternateResourceUrl = value
    }
}
// SetAssignments sets the assignments property value. A defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Azure AD that have provisioning policy assigned. Returned only on $expand. See an example of getting the assignments relationship.
func (m *CloudPcProvisioningPolicy) SetAssignments(value []CloudPcProvisioningPolicyAssignmentable)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCloudPcGroupDisplayName sets the cloudPcGroupDisplayName property value. The display name of the Cloud PC group that the Cloud PCs reside in. Read-only.
func (m *CloudPcProvisioningPolicy) SetCloudPcGroupDisplayName(value *string)() {
    if m != nil {
        m.cloudPcGroupDisplayName = value
    }
}
// SetDescription sets the description property value. The provisioning policy description.
func (m *CloudPcProvisioningPolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the provisioning policy.
func (m *CloudPcProvisioningPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDomainJoinConfiguration sets the domainJoinConfiguration property value. Specifies how Cloud PCs will join Azure Active Directory.
func (m *CloudPcProvisioningPolicy) SetDomainJoinConfiguration(value CloudPcDomainJoinConfigurationable)() {
    if m != nil {
        m.domainJoinConfiguration = value
    }
}
// SetGracePeriodInHours sets the gracePeriodInHours property value. The number of hours to wait before reprovisioning/deprovisioning happens. Read-only.
func (m *CloudPcProvisioningPolicy) SetGracePeriodInHours(value *int32)() {
    if m != nil {
        m.gracePeriodInHours = value
    }
}
// SetImageDisplayName sets the imageDisplayName property value. The display name for the OS image you’re provisioning.
func (m *CloudPcProvisioningPolicy) SetImageDisplayName(value *string)() {
    if m != nil {
        m.imageDisplayName = value
    }
}
// SetImageId sets the imageId property value. The ID of the OS image you want to provision on Cloud PCs. The format for a gallery type image is: {publisher_offer_sku}. Supported values for each of the parameters are as follows:publisher: Microsoftwindowsdesktop. offer: windows-ent-cpc. sku: 21h1-ent-cpc-m365, 21h1-ent-cpc-os, 20h2-ent-cpc-m365, 20h2-ent-cpc-os, 20h1-ent-cpc-m365, 20h1-ent-cpc-os, 19h2-ent-cpc-m365 and 19h2-ent-cpc-os.
func (m *CloudPcProvisioningPolicy) SetImageId(value *string)() {
    if m != nil {
        m.imageId = value
    }
}
// SetImageType sets the imageType property value. The type of OS image (custom or gallery) you want to provision on Cloud PCs. Possible values are: gallery, custom.
func (m *CloudPcProvisioningPolicy) SetImageType(value *CloudPcProvisioningPolicyImageType)() {
    if m != nil {
        m.imageType = value
    }
}
// SetLocalAdminEnabled sets the localAdminEnabled property value. Indicates whether the local admin option is enabled. If the local admin option is enabled, the end user can be an admin of the Cloud PC device. Read-only.
func (m *CloudPcProvisioningPolicy) SetLocalAdminEnabled(value *bool)() {
    if m != nil {
        m.localAdminEnabled = value
    }
}
// SetManagedBy sets the managedBy property value. Specifies which services manage the Azure network connection. Possible values are: windows365, devBox, unknownFutureValue. Read-only.
func (m *CloudPcProvisioningPolicy) SetManagedBy(value *CloudPcManagementService)() {
    if m != nil {
        m.managedBy = value
    }
}
// SetMicrosoftManagedDesktop sets the microsoftManagedDesktop property value. The specific settings for the Microsoft Managed Desktop, which enables customers to get a managed device experience for the Cloud PC. Before you can enable Microsoft Managed Desktop, an admin must configure it.
func (m *CloudPcProvisioningPolicy) SetMicrosoftManagedDesktop(value MicrosoftManagedDesktopable)() {
    if m != nil {
        m.microsoftManagedDesktop = value
    }
}
// SetOnPremisesConnectionId sets the onPremisesConnectionId property value. The ID of the cloudPcOnPremisesConnection. To ensure that Cloud PCs have network connectivity and that they domain join, choose a connection with a virtual network that’s validated by the Cloud PC service.
func (m *CloudPcProvisioningPolicy) SetOnPremisesConnectionId(value *string)() {
    if m != nil {
        m.onPremisesConnectionId = value
    }
}
// SetWindowsSettings sets the windowsSettings property value. Specific Windows settings to configure while creating Cloud PCs for this provisioning policy.
func (m *CloudPcProvisioningPolicy) SetWindowsSettings(value CloudPcWindowsSettingsable)() {
    if m != nil {
        m.windowsSettings = value
    }
}
