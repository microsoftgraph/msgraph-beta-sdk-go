package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new cloudAppSecurityProfile and sets the default values.
func NewCloudAppSecurityProfile()(*CloudAppSecurityProfile) {
    m := &CloudAppSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the azureSubscriptionId property value. 
func (m *CloudAppSecurityProfile) GetAzureSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureSubscriptionId
    }
}
// Gets the azureTenantId property value. 
func (m *CloudAppSecurityProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// Gets the createdDateTime property value. 
func (m *CloudAppSecurityProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the deploymentPackageUrl property value. 
func (m *CloudAppSecurityProfile) GetDeploymentPackageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deploymentPackageUrl
    }
}
// Gets the destinationServiceName property value. 
func (m *CloudAppSecurityProfile) GetDestinationServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationServiceName
    }
}
// Gets the isSigned property value. 
func (m *CloudAppSecurityProfile) GetIsSigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSigned
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *CloudAppSecurityProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the manifest property value. 
func (m *CloudAppSecurityProfile) GetManifest()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manifest
    }
}
// Gets the name property value. 
func (m *CloudAppSecurityProfile) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the permissionsRequired property value. 
func (m *CloudAppSecurityProfile) GetPermissionsRequired()(*ApplicationPermissionsRequired) {
    if m == nil {
        return nil
    } else {
        return m.permissionsRequired
    }
}
// Gets the platform property value. 
func (m *CloudAppSecurityProfile) GetPlatform()(*string) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// Gets the policyName property value. 
func (m *CloudAppSecurityProfile) GetPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyName
    }
}
// Gets the publisher property value. 
func (m *CloudAppSecurityProfile) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// Gets the riskScore property value. 
func (m *CloudAppSecurityProfile) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// Gets the tags property value. 
func (m *CloudAppSecurityProfile) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// Gets the type_escaped property value. 
func (m *CloudAppSecurityProfile) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the vendorInformation property value. 
func (m *CloudAppSecurityProfile) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// The deserialization information for the current model
func (m *CloudAppSecurityProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureSubscriptionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureSubscriptionId(val)
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureTenantId(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["deploymentPackageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeploymentPackageUrl(val)
        return nil
    }
    res["destinationServiceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDestinationServiceName(val)
        return nil
    }
    res["isSigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSigned(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["manifest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManifest(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["permissionsRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationPermissionsRequired)
        if err != nil {
            return err
        }
        cast := val.(ApplicationPermissionsRequired)
        m.SetPermissionsRequired(&cast)
        return nil
    }
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPlatform(val)
        return nil
    }
    res["policyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyName(val)
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublisher(val)
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRiskScore(val)
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetTags(res)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityVendorInformation() })
        if err != nil {
            return err
        }
        m.SetVendorInformation(val.(*SecurityVendorInformation))
        return nil
    }
    return res
}
func (m *CloudAppSecurityProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        cast := m.GetPermissionsRequired().String()
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
    {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
// Sets the azureSubscriptionId property value. 
// Parameters:
//  - value : Value to set for the azureSubscriptionId property.
func (m *CloudAppSecurityProfile) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
// Sets the azureTenantId property value. 
// Parameters:
//  - value : Value to set for the azureTenantId property.
func (m *CloudAppSecurityProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// Sets the createdDateTime property value. 
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *CloudAppSecurityProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the deploymentPackageUrl property value. 
// Parameters:
//  - value : Value to set for the deploymentPackageUrl property.
func (m *CloudAppSecurityProfile) SetDeploymentPackageUrl(value *string)() {
    m.deploymentPackageUrl = value
}
// Sets the destinationServiceName property value. 
// Parameters:
//  - value : Value to set for the destinationServiceName property.
func (m *CloudAppSecurityProfile) SetDestinationServiceName(value *string)() {
    m.destinationServiceName = value
}
// Sets the isSigned property value. 
// Parameters:
//  - value : Value to set for the isSigned property.
func (m *CloudAppSecurityProfile) SetIsSigned(value *bool)() {
    m.isSigned = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *CloudAppSecurityProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the manifest property value. 
// Parameters:
//  - value : Value to set for the manifest property.
func (m *CloudAppSecurityProfile) SetManifest(value *string)() {
    m.manifest = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *CloudAppSecurityProfile) SetName(value *string)() {
    m.name = value
}
// Sets the permissionsRequired property value. 
// Parameters:
//  - value : Value to set for the permissionsRequired property.
func (m *CloudAppSecurityProfile) SetPermissionsRequired(value *ApplicationPermissionsRequired)() {
    m.permissionsRequired = value
}
// Sets the platform property value. 
// Parameters:
//  - value : Value to set for the platform property.
func (m *CloudAppSecurityProfile) SetPlatform(value *string)() {
    m.platform = value
}
// Sets the policyName property value. 
// Parameters:
//  - value : Value to set for the policyName property.
func (m *CloudAppSecurityProfile) SetPolicyName(value *string)() {
    m.policyName = value
}
// Sets the publisher property value. 
// Parameters:
//  - value : Value to set for the publisher property.
func (m *CloudAppSecurityProfile) SetPublisher(value *string)() {
    m.publisher = value
}
// Sets the riskScore property value. 
// Parameters:
//  - value : Value to set for the riskScore property.
func (m *CloudAppSecurityProfile) SetRiskScore(value *string)() {
    m.riskScore = value
}
// Sets the tags property value. 
// Parameters:
//  - value : Value to set for the tags property.
func (m *CloudAppSecurityProfile) SetTags(value []string)() {
    m.tags = value
}
// Sets the type_escaped property value. 
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *CloudAppSecurityProfile) SetType_escaped(value *string)() {
    m.type_escaped = value
}
// Sets the vendorInformation property value. 
// Parameters:
//  - value : Value to set for the vendorInformation property.
func (m *CloudAppSecurityProfile) SetVendorInformation(value *SecurityVendorInformation)() {
    m.vendorInformation = value
}
