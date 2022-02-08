package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcProvisioningPolicy 
type CloudPcProvisioningPolicy struct {
    Entity
    // A defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Azure AD that have provisioning policy assigned. Returned only on $expand. See an example of getting the assignments relationship.
    assignments []CloudPcProvisioningPolicyAssignment;
    // The provisioning policy description.
    description *string;
    // The display name for the provisioning policy.
    displayName *string;
    // Specifies how Cloud PCs will join Azure Active Directory.
    domainJoinConfiguration *CloudPcDomainJoinConfiguration;
    // The display name for the OS image you’re provisioning.
    imageDisplayName *string;
    // The ID of the OS image you want to provision on Cloud PCs. The format for a gallery type image is: {publisher_offer_sku}. Supported values for each of the parameters are as follows:publisher: Microsoftwindowsdesktop. offer: windows-ent-cpc. sku: 21h1-ent-cpc-m365, 21h1-ent-cpc-os, 20h2-ent-cpc-m365, 20h2-ent-cpc-os, 20h1-ent-cpc-m365, 20h1-ent-cpc-os, 19h2-ent-cpc-m365 and 19h2-ent-cpc-os.
    imageId *string;
    // The type of OS image (custom or gallery) you want to provision on Cloud PCs. Possible values are: gallery, custom.
    imageType *CloudPcProvisioningPolicyImageType;
    // The specific settings for the Microsoft Managed Desktop, which enables customers to get a managed device experience for the Cloud PC. Before you can enable Microsoft Managed Desktop, an admin must configure it.
    microsoftManagedDesktop *MicrosoftManagedDesktop;
    // The ID of the cloudPcOnPremisesConnection. To ensure that Cloud PCs have network connectivity and that they domain join, choose a connection with a virtual network that’s validated by the Cloud PC service.
    onPremisesConnectionId *string;
}
// NewCloudPcProvisioningPolicy instantiates a new cloudPcProvisioningPolicy and sets the default values.
func NewCloudPcProvisioningPolicy()(*CloudPcProvisioningPolicy) {
    m := &CloudPcProvisioningPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. A defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Azure AD that have provisioning policy assigned. Returned only on $expand. See an example of getting the assignments relationship.
func (m *CloudPcProvisioningPolicy) GetAssignments()([]CloudPcProvisioningPolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
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
func (m *CloudPcProvisioningPolicy) GetDomainJoinConfiguration()(*CloudPcDomainJoinConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.domainJoinConfiguration
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
// GetMicrosoftManagedDesktop gets the microsoftManagedDesktop property value. The specific settings for the Microsoft Managed Desktop, which enables customers to get a managed device experience for the Cloud PC. Before you can enable Microsoft Managed Desktop, an admin must configure it.
func (m *CloudPcProvisioningPolicy) GetMicrosoftManagedDesktop()(*MicrosoftManagedDesktop) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcProvisioningPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcProvisioningPolicyAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcProvisioningPolicyAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcProvisioningPolicyAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["domainJoinConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcDomainJoinConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainJoinConfiguration(val.(*CloudPcDomainJoinConfiguration))
        }
        return nil
    }
    res["imageDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageDisplayName(val)
        }
        return nil
    }
    res["imageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageId(val)
        }
        return nil
    }
    res["imageType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcProvisioningPolicyImageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageType(val.(*CloudPcProvisioningPolicyImageType))
        }
        return nil
    }
    res["microsoftManagedDesktop"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftManagedDesktop() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftManagedDesktop(val.(*MicrosoftManagedDesktop))
        }
        return nil
    }
    res["onPremisesConnectionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesConnectionId(val)
        }
        return nil
    }
    return res
}
func (m *CloudPcProvisioningPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudPcProvisioningPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
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
    return nil
}
// SetAssignments sets the assignments property value. A defined collection of provisioning policy assignments. Represents the set of Microsoft 365 groups and security groups in Azure AD that have provisioning policy assigned. Returned only on $expand. See an example of getting the assignments relationship.
func (m *CloudPcProvisioningPolicy) SetAssignments(value []CloudPcProvisioningPolicyAssignment)() {
    if m != nil {
        m.assignments = value
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
func (m *CloudPcProvisioningPolicy) SetDomainJoinConfiguration(value *CloudPcDomainJoinConfiguration)() {
    if m != nil {
        m.domainJoinConfiguration = value
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
// SetMicrosoftManagedDesktop sets the microsoftManagedDesktop property value. The specific settings for the Microsoft Managed Desktop, which enables customers to get a managed device experience for the Cloud PC. Before you can enable Microsoft Managed Desktop, an admin must configure it.
func (m *CloudPcProvisioningPolicy) SetMicrosoftManagedDesktop(value *MicrosoftManagedDesktop)() {
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
