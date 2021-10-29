package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcOverview struct {
    Entity
    // Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The number of cloud PC connections that have a status of failed. Optional. Read-only.
    numberOfCloudPcConnectionStatusFailed *int32;
    // The number of cloud PC connections that have a status of passed. Optional. Read-only.
    numberOfCloudPcConnectionStatusPassed *int32;
    // The number of cloud PC connections that have a status of pending. Optional. Read-only.
    numberOfCloudPcConnectionStatusPending *int32;
    // The number of cloud PC connections that have a status of running. Optional. Read-only.
    numberOfCloudPcConnectionStatusRunning *int32;
    // The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
    numberOfCloudPcConnectionStatusUnkownFutureValue *int32;
    // The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
    numberOfCloudPcStatusDeprovisioning *int32;
    // The number of cloud PCs that have a status of failed. Optional. Read-only.
    numberOfCloudPcStatusFailed *int32;
    // The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
    numberOfCloudPcStatusInGracePeriod *int32;
    // The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
    numberOfCloudPcStatusNotProvisioned *int32;
    // The number of cloud PCs that have a status of provisioned. Optional. Read-only.
    numberOfCloudPcStatusProvisioned *int32;
    // The number of cloud PCs that have a status of provisioning. Optional. Read-only.
    numberOfCloudPcStatusProvisioning *int32;
    // The number of cloud PCs that have a status of unknown. Optional. Read-only.
    numberOfCloudPcStatusUnknown *int32;
    // The number of cloud PCs that have a status of upgrading. Optional. Read-only.
    numberOfCloudPcStatusUpgrading *int32;
    // The display name for the managed tenant. Optional. Read-only.
    tenantDisplayName *string;
    // 
    tenantId *string;
    // The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
    totalCloudPcConnectionStatus *int32;
    // The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
    totalCloudPcStatus *int32;
}
// Instantiates a new cloudPcOverview and sets the default values.
func NewCloudPcOverview()(*CloudPcOverview) {
    m := &CloudPcOverview{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CloudPcOverview) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
// Gets the numberOfCloudPcConnectionStatusFailed property value. The number of cloud PC connections that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusFailed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusFailed
    }
}
// Gets the numberOfCloudPcConnectionStatusPassed property value. The number of cloud PC connections that have a status of passed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPassed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusPassed
    }
}
// Gets the numberOfCloudPcConnectionStatusPending property value. The number of cloud PC connections that have a status of pending. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPending()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusPending
    }
}
// Gets the numberOfCloudPcConnectionStatusRunning property value. The number of cloud PC connections that have a status of running. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusRunning()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusRunning
    }
}
// Gets the numberOfCloudPcConnectionStatusUnkownFutureValue property value. The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusUnkownFutureValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusUnkownFutureValue
    }
}
// Gets the numberOfCloudPcStatusDeprovisioning property value. The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusDeprovisioning()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusDeprovisioning
    }
}
// Gets the numberOfCloudPcStatusFailed property value. The number of cloud PCs that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusFailed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusFailed
    }
}
// Gets the numberOfCloudPcStatusInGracePeriod property value. The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusInGracePeriod()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusInGracePeriod
    }
}
// Gets the numberOfCloudPcStatusNotProvisioned property value. The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusNotProvisioned()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusNotProvisioned
    }
}
// Gets the numberOfCloudPcStatusProvisioned property value. The number of cloud PCs that have a status of provisioned. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioned()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusProvisioned
    }
}
// Gets the numberOfCloudPcStatusProvisioning property value. The number of cloud PCs that have a status of provisioning. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioning()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusProvisioning
    }
}
// Gets the numberOfCloudPcStatusUnknown property value. The number of cloud PCs that have a status of unknown. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUnknown()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusUnknown
    }
}
// Gets the numberOfCloudPcStatusUpgrading property value. The number of cloud PCs that have a status of upgrading. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUpgrading()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusUpgrading
    }
}
// Gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
// Gets the tenantId property value. 
func (m *CloudPcOverview) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Gets the totalCloudPcConnectionStatus property value. The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTotalCloudPcConnectionStatus()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalCloudPcConnectionStatus
    }
}
// Gets the totalCloudPcStatus property value. The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTotalCloudPcStatus()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalCloudPcStatus
    }
}
// The deserialization information for the current model
func (m *CloudPcOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastRefreshedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastRefreshedDateTime(val)
        return nil
    }
    res["numberOfCloudPcConnectionStatusFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcConnectionStatusFailed(val)
        return nil
    }
    res["numberOfCloudPcConnectionStatusPassed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcConnectionStatusPassed(val)
        return nil
    }
    res["numberOfCloudPcConnectionStatusPending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcConnectionStatusPending(val)
        return nil
    }
    res["numberOfCloudPcConnectionStatusRunning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcConnectionStatusRunning(val)
        return nil
    }
    res["numberOfCloudPcConnectionStatusUnkownFutureValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcConnectionStatusUnkownFutureValue(val)
        return nil
    }
    res["numberOfCloudPcStatusDeprovisioning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcStatusDeprovisioning(val)
        return nil
    }
    res["numberOfCloudPcStatusFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcStatusFailed(val)
        return nil
    }
    res["numberOfCloudPcStatusInGracePeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcStatusInGracePeriod(val)
        return nil
    }
    res["numberOfCloudPcStatusNotProvisioned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcStatusNotProvisioned(val)
        return nil
    }
    res["numberOfCloudPcStatusProvisioned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcStatusProvisioned(val)
        return nil
    }
    res["numberOfCloudPcStatusProvisioning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcStatusProvisioning(val)
        return nil
    }
    res["numberOfCloudPcStatusUnknown"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcStatusUnknown(val)
        return nil
    }
    res["numberOfCloudPcStatusUpgrading"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfCloudPcStatusUpgrading(val)
        return nil
    }
    res["tenantDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantDisplayName(val)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    res["totalCloudPcConnectionStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalCloudPcConnectionStatus(val)
        return nil
    }
    res["totalCloudPcStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalCloudPcStatus(val)
        return nil
    }
    return res
}
func (m *CloudPcOverview) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CloudPcOverview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcConnectionStatusFailed", m.GetNumberOfCloudPcConnectionStatusFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcConnectionStatusPassed", m.GetNumberOfCloudPcConnectionStatusPassed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcConnectionStatusPending", m.GetNumberOfCloudPcConnectionStatusPending())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcConnectionStatusRunning", m.GetNumberOfCloudPcConnectionStatusRunning())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcConnectionStatusUnkownFutureValue", m.GetNumberOfCloudPcConnectionStatusUnkownFutureValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusDeprovisioning", m.GetNumberOfCloudPcStatusDeprovisioning())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusFailed", m.GetNumberOfCloudPcStatusFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusInGracePeriod", m.GetNumberOfCloudPcStatusInGracePeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusNotProvisioned", m.GetNumberOfCloudPcStatusNotProvisioned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusProvisioned", m.GetNumberOfCloudPcStatusProvisioned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusProvisioning", m.GetNumberOfCloudPcStatusProvisioning())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusUnknown", m.GetNumberOfCloudPcStatusUnknown())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusUpgrading", m.GetNumberOfCloudPcStatusUpgrading())
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
    {
        err = writer.WriteInt32Value("totalCloudPcConnectionStatus", m.GetTotalCloudPcConnectionStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalCloudPcStatus", m.GetTotalCloudPcStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
// Parameters:
//  - value : Value to set for the lastRefreshedDateTime property.
func (m *CloudPcOverview) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
// Sets the numberOfCloudPcConnectionStatusFailed property value. The number of cloud PC connections that have a status of failed. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcConnectionStatusFailed property.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusFailed(value *int32)() {
    m.numberOfCloudPcConnectionStatusFailed = value
}
// Sets the numberOfCloudPcConnectionStatusPassed property value. The number of cloud PC connections that have a status of passed. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcConnectionStatusPassed property.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPassed(value *int32)() {
    m.numberOfCloudPcConnectionStatusPassed = value
}
// Sets the numberOfCloudPcConnectionStatusPending property value. The number of cloud PC connections that have a status of pending. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcConnectionStatusPending property.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPending(value *int32)() {
    m.numberOfCloudPcConnectionStatusPending = value
}
// Sets the numberOfCloudPcConnectionStatusRunning property value. The number of cloud PC connections that have a status of running. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcConnectionStatusRunning property.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusRunning(value *int32)() {
    m.numberOfCloudPcConnectionStatusRunning = value
}
// Sets the numberOfCloudPcConnectionStatusUnkownFutureValue property value. The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcConnectionStatusUnkownFutureValue property.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusUnkownFutureValue(value *int32)() {
    m.numberOfCloudPcConnectionStatusUnkownFutureValue = value
}
// Sets the numberOfCloudPcStatusDeprovisioning property value. The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcStatusDeprovisioning property.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusDeprovisioning(value *int32)() {
    m.numberOfCloudPcStatusDeprovisioning = value
}
// Sets the numberOfCloudPcStatusFailed property value. The number of cloud PCs that have a status of failed. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcStatusFailed property.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusFailed(value *int32)() {
    m.numberOfCloudPcStatusFailed = value
}
// Sets the numberOfCloudPcStatusInGracePeriod property value. The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcStatusInGracePeriod property.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusInGracePeriod(value *int32)() {
    m.numberOfCloudPcStatusInGracePeriod = value
}
// Sets the numberOfCloudPcStatusNotProvisioned property value. The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcStatusNotProvisioned property.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusNotProvisioned(value *int32)() {
    m.numberOfCloudPcStatusNotProvisioned = value
}
// Sets the numberOfCloudPcStatusProvisioned property value. The number of cloud PCs that have a status of provisioned. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcStatusProvisioned property.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioned(value *int32)() {
    m.numberOfCloudPcStatusProvisioned = value
}
// Sets the numberOfCloudPcStatusProvisioning property value. The number of cloud PCs that have a status of provisioning. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcStatusProvisioning property.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioning(value *int32)() {
    m.numberOfCloudPcStatusProvisioning = value
}
// Sets the numberOfCloudPcStatusUnknown property value. The number of cloud PCs that have a status of unknown. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcStatusUnknown property.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUnknown(value *int32)() {
    m.numberOfCloudPcStatusUnknown = value
}
// Sets the numberOfCloudPcStatusUpgrading property value. The number of cloud PCs that have a status of upgrading. Optional. Read-only.
// Parameters:
//  - value : Value to set for the numberOfCloudPcStatusUpgrading property.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUpgrading(value *int32)() {
    m.numberOfCloudPcStatusUpgrading = value
}
// Sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the tenantDisplayName property.
func (m *CloudPcOverview) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// Sets the tenantId property value. 
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *CloudPcOverview) SetTenantId(value *string)() {
    m.tenantId = value
}
// Sets the totalCloudPcConnectionStatus property value. The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the totalCloudPcConnectionStatus property.
func (m *CloudPcOverview) SetTotalCloudPcConnectionStatus(value *int32)() {
    m.totalCloudPcConnectionStatus = value
}
// Sets the totalCloudPcStatus property value. The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the totalCloudPcStatus property.
func (m *CloudPcOverview) SetTotalCloudPcStatus(value *int32)() {
    m.totalCloudPcStatus = value
}
