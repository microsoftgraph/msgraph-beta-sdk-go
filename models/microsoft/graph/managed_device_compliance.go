package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedDeviceCompliance 
type ManagedDeviceCompliance struct {
    Entity
    // Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager. Optional. Read-only.
    complianceStatus *string;
    // Platform of the device. This property is read-only. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.  Optional. Read-only.
    deviceType *string;
    // The date and time when the grace period will expire. Optional. Read-only.
    inGracePeriodUntilDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The date and time that the device last completed a successful sync with Microsoft Endpoint Manager. Optional. Read-only.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The identifier for the managed device in Microsoft Endpoint Manager. Optional. Read-only.
    managedDeviceId *string;
    // The display name for the managed device. Optional. Read-only.
    managedDeviceName *string;
    // The manufacture for the device. Optional. Read-only.
    manufacturer *string;
    // The model for the device. Optional. Read-only.
    model *string;
    // The description of the operating system for the managed device. Optional. Read-only.
    osDescription *string;
    // The version of the operating system for the managed device. Optional. Read-only.
    osVersion *string;
    // The type of owner for the managed device. Optional. Read-only.
    ownerType *string;
    // The display name for the managed tenant. Optional. Read-only.
    tenantDisplayName *string;
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string;
}
// NewManagedDeviceCompliance instantiates a new managedDeviceCompliance and sets the default values.
func NewManagedDeviceCompliance()(*ManagedDeviceCompliance) {
    m := &ManagedDeviceCompliance{
        Entity: *NewEntity(),
    }
    return m
}
// GetComplianceStatus gets the complianceStatus property value. Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetComplianceStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.complianceStatus
    }
}
// GetDeviceType gets the deviceType property value. Platform of the device. This property is read-only. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.  Optional. Read-only.
func (m *ManagedDeviceCompliance) GetDeviceType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
// GetInGracePeriodUntilDateTime gets the inGracePeriodUntilDateTime property value. The date and time when the grace period will expire. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetInGracePeriodUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.inGracePeriodUntilDateTime
    }
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Microsoft Endpoint Manager. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetManagedDeviceId gets the managedDeviceId property value. The identifier for the managed device in Microsoft Endpoint Manager. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// GetManagedDeviceName gets the managedDeviceName property value. The display name for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetManagedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceName
    }
}
// GetManufacturer gets the manufacturer property value. The manufacture for the device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetModel gets the model property value. The model for the device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetOsDescription gets the osDescription property value. The description of the operating system for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
// GetOsVersion gets the osVersion property value. The version of the operating system for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetOwnerType gets the ownerType property value. The type of owner for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetOwnerType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerType
    }
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceCompliance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["complianceStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceStatus(val)
        }
        return nil
    }
    res["deviceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val)
        }
        return nil
    }
    res["inGracePeriodUntilDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInGracePeriodUntilDateTime(val)
        }
        return nil
    }
    res["lastRefreshedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedDateTime(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["managedDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["osDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
        }
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["ownerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val)
        }
        return nil
    }
    res["tenantDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
func (m *ManagedDeviceCompliance) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceCompliance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("complianceStatus", m.GetComplianceStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceType", m.GetDeviceType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("inGracePeriodUntilDateTime", m.GetInGracePeriodUntilDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osDescription", m.GetOsDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerType", m.GetOwnerType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDisplayName", m.GetTenantDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComplianceStatus sets the complianceStatus property value. Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetComplianceStatus(value *string)() {
    if m != nil {
        m.complianceStatus = value
    }
}
// SetDeviceType sets the deviceType property value. Platform of the device. This property is read-only. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.  Optional. Read-only.
func (m *ManagedDeviceCompliance) SetDeviceType(value *string)() {
    if m != nil {
        m.deviceType = value
    }
}
// SetInGracePeriodUntilDateTime sets the inGracePeriodUntilDateTime property value. The date and time when the grace period will expire. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetInGracePeriodUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.inGracePeriodUntilDateTime = value
    }
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastRefreshedDateTime = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Microsoft Endpoint Manager. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. The identifier for the managed device in Microsoft Endpoint Manager. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetManagedDeviceId(value *string)() {
    if m != nil {
        m.managedDeviceId = value
    }
}
// SetManagedDeviceName sets the managedDeviceName property value. The display name for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetManagedDeviceName(value *string)() {
    if m != nil {
        m.managedDeviceName = value
    }
}
// SetManufacturer sets the manufacturer property value. The manufacture for the device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetModel sets the model property value. The model for the device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetOsDescription sets the osDescription property value. The description of the operating system for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetOsDescription(value *string)() {
    if m != nil {
        m.osDescription = value
    }
}
// SetOsVersion sets the osVersion property value. The version of the operating system for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetOwnerType sets the ownerType property value. The type of owner for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetOwnerType(value *string)() {
    if m != nil {
        m.ownerType = value
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetTenantDisplayName(value *string)() {
    if m != nil {
        m.tenantDisplayName = value
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
