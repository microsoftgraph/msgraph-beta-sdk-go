package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudAppSecurityProfile 
type CloudAppSecurityProfile struct {
    Entity
    // 
    azureSubscriptionId *string;
    // 
    azureTenantId *string;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    deploymentPackageUrl *string;
    // 
    destinationServiceName *string;
    // 
    isSigned *bool;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    manifest *string;
    // 
    name *string;
    // 
    permissionsRequired *ApplicationPermissionsRequired;
    // 
    platform *string;
    // 
    policyName *string;
    // 
    publisher *string;
    // 
    riskScore *string;
    // 
    tags []string;
    // 
    type_escaped *string;
    // 
    vendorInformation *SecurityVendorInformation;
}
// NewCloudAppSecurityProfile instantiates a new cloudAppSecurityProfile and sets the default values.
func NewCloudAppSecurityProfile()(*CloudAppSecurityProfile) {
    m := &CloudAppSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. 
func (m *CloudAppSecurityProfile) GetAzureSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureSubscriptionId
    }
}
// GetAzureTenantId gets the azureTenantId property value. 
func (m *CloudAppSecurityProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *CloudAppSecurityProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDeploymentPackageUrl gets the deploymentPackageUrl property value. 
func (m *CloudAppSecurityProfile) GetDeploymentPackageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deploymentPackageUrl
    }
}
// GetDestinationServiceName gets the destinationServiceName property value. 
func (m *CloudAppSecurityProfile) GetDestinationServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationServiceName
    }
}
// GetIsSigned gets the isSigned property value. 
func (m *CloudAppSecurityProfile) GetIsSigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSigned
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *CloudAppSecurityProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetManifest gets the manifest property value. 
func (m *CloudAppSecurityProfile) GetManifest()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manifest
    }
}
// GetName gets the name property value. 
func (m *CloudAppSecurityProfile) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPermissionsRequired gets the permissionsRequired property value. 
func (m *CloudAppSecurityProfile) GetPermissionsRequired()(*ApplicationPermissionsRequired) {
    if m == nil {
        return nil
    } else {
        return m.permissionsRequired
    }
}
// GetPlatform gets the platform property value. 
func (m *CloudAppSecurityProfile) GetPlatform()(*string) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// GetPolicyName gets the policyName property value. 
func (m *CloudAppSecurityProfile) GetPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyName
    }
}
// GetPublisher gets the publisher property value. 
func (m *CloudAppSecurityProfile) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetRiskScore gets the riskScore property value. 
func (m *CloudAppSecurityProfile) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// GetTags gets the tags property value. 
func (m *CloudAppSecurityProfile) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// GetType gets the type property value. 
func (m *CloudAppSecurityProfile) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetVendorInformation gets the vendorInformation property value. 
func (m *CloudAppSecurityProfile) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudAppSecurityProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureSubscriptionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureSubscriptionId(val)
        }
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deploymentPackageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentPackageUrl(val)
        }
        return nil
    }
    res["destinationServiceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationServiceName(val)
        }
        return nil
    }
    res["isSigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSigned(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["manifest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManifest(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["permissionsRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationPermissionsRequired)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionsRequired(val.(*ApplicationPermissionsRequired))
        }
        return nil
    }
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val)
        }
        return nil
    }
    res["policyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityVendorInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInformation(val.(*SecurityVendorInformation))
        }
        return nil
    }
    return res
}
func (m *CloudAppSecurityProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudAppSecurityProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureSubscriptionId", m.GetAzureSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deploymentPackageUrl", m.GetDeploymentPackageUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("destinationServiceName", m.GetDestinationServiceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSigned", m.GetIsSigned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manifest", m.GetManifest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionsRequired() != nil {
        cast := (*m.GetPermissionsRequired()).String()
        err = writer.WriteStringValue("permissionsRequired", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("platform", m.GetPlatform())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyName", m.GetPolicyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("riskScore", m.GetRiskScore())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureSubscriptionId sets the azureSubscriptionId property value. 
func (m *CloudAppSecurityProfile) SetAzureSubscriptionId(value *string)() {
    if m != nil {
        m.azureSubscriptionId = value
    }
}
// SetAzureTenantId sets the azureTenantId property value. 
func (m *CloudAppSecurityProfile) SetAzureTenantId(value *string)() {
    if m != nil {
        m.azureTenantId = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *CloudAppSecurityProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDeploymentPackageUrl sets the deploymentPackageUrl property value. 
func (m *CloudAppSecurityProfile) SetDeploymentPackageUrl(value *string)() {
    if m != nil {
        m.deploymentPackageUrl = value
    }
}
// SetDestinationServiceName sets the destinationServiceName property value. 
func (m *CloudAppSecurityProfile) SetDestinationServiceName(value *string)() {
    if m != nil {
        m.destinationServiceName = value
    }
}
// SetIsSigned sets the isSigned property value. 
func (m *CloudAppSecurityProfile) SetIsSigned(value *bool)() {
    if m != nil {
        m.isSigned = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *CloudAppSecurityProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetManifest sets the manifest property value. 
func (m *CloudAppSecurityProfile) SetManifest(value *string)() {
    if m != nil {
        m.manifest = value
    }
}
// SetName sets the name property value. 
func (m *CloudAppSecurityProfile) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPermissionsRequired sets the permissionsRequired property value. 
func (m *CloudAppSecurityProfile) SetPermissionsRequired(value *ApplicationPermissionsRequired)() {
    if m != nil {
        m.permissionsRequired = value
    }
}
// SetPlatform sets the platform property value. 
func (m *CloudAppSecurityProfile) SetPlatform(value *string)() {
    if m != nil {
        m.platform = value
    }
}
// SetPolicyName sets the policyName property value. 
func (m *CloudAppSecurityProfile) SetPolicyName(value *string)() {
    if m != nil {
        m.policyName = value
    }
}
// SetPublisher sets the publisher property value. 
func (m *CloudAppSecurityProfile) SetPublisher(value *string)() {
    if m != nil {
        m.publisher = value
    }
}
// SetRiskScore sets the riskScore property value. 
func (m *CloudAppSecurityProfile) SetRiskScore(value *string)() {
    if m != nil {
        m.riskScore = value
    }
}
// SetTags sets the tags property value. 
func (m *CloudAppSecurityProfile) SetTags(value []string)() {
    if m != nil {
        m.tags = value
    }
}
// SetType sets the type property value. 
func (m *CloudAppSecurityProfile) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetVendorInformation sets the vendorInformation property value. 
func (m *CloudAppSecurityProfile) SetVendorInformation(value *SecurityVendorInformation)() {
    if m != nil {
        m.vendorInformation = value
    }
}
