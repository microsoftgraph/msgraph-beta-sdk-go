package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcOverview 
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
// NewCloudPcOverview instantiates a new cloudPcOverview and sets the default values.
func NewCloudPcOverview()(*CloudPcOverview) {
    m := &CloudPcOverview{
        Entity: *NewEntity(),
    }
    return m
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CloudPcOverview) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
// GetNumberOfCloudPcConnectionStatusFailed gets the numberOfCloudPcConnectionStatusFailed property value. The number of cloud PC connections that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusFailed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusFailed
    }
}
// GetNumberOfCloudPcConnectionStatusPassed gets the numberOfCloudPcConnectionStatusPassed property value. The number of cloud PC connections that have a status of passed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPassed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusPassed
    }
}
// GetNumberOfCloudPcConnectionStatusPending gets the numberOfCloudPcConnectionStatusPending property value. The number of cloud PC connections that have a status of pending. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPending()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusPending
    }
}
// GetNumberOfCloudPcConnectionStatusRunning gets the numberOfCloudPcConnectionStatusRunning property value. The number of cloud PC connections that have a status of running. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusRunning()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusRunning
    }
}
// GetNumberOfCloudPcConnectionStatusUnkownFutureValue gets the numberOfCloudPcConnectionStatusUnkownFutureValue property value. The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusUnkownFutureValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusUnkownFutureValue
    }
}
// GetNumberOfCloudPcStatusDeprovisioning gets the numberOfCloudPcStatusDeprovisioning property value. The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusDeprovisioning()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusDeprovisioning
    }
}
// GetNumberOfCloudPcStatusFailed gets the numberOfCloudPcStatusFailed property value. The number of cloud PCs that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusFailed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusFailed
    }
}
// GetNumberOfCloudPcStatusInGracePeriod gets the numberOfCloudPcStatusInGracePeriod property value. The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusInGracePeriod()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusInGracePeriod
    }
}
// GetNumberOfCloudPcStatusNotProvisioned gets the numberOfCloudPcStatusNotProvisioned property value. The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusNotProvisioned()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusNotProvisioned
    }
}
// GetNumberOfCloudPcStatusProvisioned gets the numberOfCloudPcStatusProvisioned property value. The number of cloud PCs that have a status of provisioned. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioned()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusProvisioned
    }
}
// GetNumberOfCloudPcStatusProvisioning gets the numberOfCloudPcStatusProvisioning property value. The number of cloud PCs that have a status of provisioning. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioning()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusProvisioning
    }
}
// GetNumberOfCloudPcStatusUnknown gets the numberOfCloudPcStatusUnknown property value. The number of cloud PCs that have a status of unknown. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUnknown()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusUnknown
    }
}
// GetNumberOfCloudPcStatusUpgrading gets the numberOfCloudPcStatusUpgrading property value. The number of cloud PCs that have a status of upgrading. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUpgrading()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusUpgrading
    }
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
// GetTenantId gets the tenantId property value. 
func (m *CloudPcOverview) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetTotalCloudPcConnectionStatus gets the totalCloudPcConnectionStatus property value. The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTotalCloudPcConnectionStatus()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalCloudPcConnectionStatus
    }
}
// GetTotalCloudPcStatus gets the totalCloudPcStatus property value. The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTotalCloudPcStatus()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalCloudPcStatus
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["numberOfCloudPcConnectionStatusFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusFailed(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusPassed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusPassed(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusPending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusPending(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusRunning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusRunning(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusUnkownFutureValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusUnkownFutureValue(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusDeprovisioning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusDeprovisioning(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusFailed(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusInGracePeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusInGracePeriod(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusNotProvisioned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusNotProvisioned(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusProvisioned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusProvisioned(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusProvisioning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusProvisioning(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusUnknown"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusUnknown(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusUpgrading"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusUpgrading(val)
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
    res["totalCloudPcConnectionStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCloudPcConnectionStatus(val)
        }
        return nil
    }
    res["totalCloudPcStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCloudPcStatus(val)
        }
        return nil
    }
    return res
}
func (m *CloudPcOverview) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CloudPcOverview) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastRefreshedDateTime = value
    }
}
// SetNumberOfCloudPcConnectionStatusFailed sets the numberOfCloudPcConnectionStatusFailed property value. The number of cloud PC connections that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusFailed(value *int32)() {
    if m != nil {
        m.numberOfCloudPcConnectionStatusFailed = value
    }
}
// SetNumberOfCloudPcConnectionStatusPassed sets the numberOfCloudPcConnectionStatusPassed property value. The number of cloud PC connections that have a status of passed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPassed(value *int32)() {
    if m != nil {
        m.numberOfCloudPcConnectionStatusPassed = value
    }
}
// SetNumberOfCloudPcConnectionStatusPending sets the numberOfCloudPcConnectionStatusPending property value. The number of cloud PC connections that have a status of pending. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPending(value *int32)() {
    if m != nil {
        m.numberOfCloudPcConnectionStatusPending = value
    }
}
// SetNumberOfCloudPcConnectionStatusRunning sets the numberOfCloudPcConnectionStatusRunning property value. The number of cloud PC connections that have a status of running. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusRunning(value *int32)() {
    if m != nil {
        m.numberOfCloudPcConnectionStatusRunning = value
    }
}
// SetNumberOfCloudPcConnectionStatusUnkownFutureValue sets the numberOfCloudPcConnectionStatusUnkownFutureValue property value. The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusUnkownFutureValue(value *int32)() {
    if m != nil {
        m.numberOfCloudPcConnectionStatusUnkownFutureValue = value
    }
}
// SetNumberOfCloudPcStatusDeprovisioning sets the numberOfCloudPcStatusDeprovisioning property value. The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusDeprovisioning(value *int32)() {
    if m != nil {
        m.numberOfCloudPcStatusDeprovisioning = value
    }
}
// SetNumberOfCloudPcStatusFailed sets the numberOfCloudPcStatusFailed property value. The number of cloud PCs that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusFailed(value *int32)() {
    if m != nil {
        m.numberOfCloudPcStatusFailed = value
    }
}
// SetNumberOfCloudPcStatusInGracePeriod sets the numberOfCloudPcStatusInGracePeriod property value. The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusInGracePeriod(value *int32)() {
    if m != nil {
        m.numberOfCloudPcStatusInGracePeriod = value
    }
}
// SetNumberOfCloudPcStatusNotProvisioned sets the numberOfCloudPcStatusNotProvisioned property value. The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusNotProvisioned(value *int32)() {
    if m != nil {
        m.numberOfCloudPcStatusNotProvisioned = value
    }
}
// SetNumberOfCloudPcStatusProvisioned sets the numberOfCloudPcStatusProvisioned property value. The number of cloud PCs that have a status of provisioned. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioned(value *int32)() {
    if m != nil {
        m.numberOfCloudPcStatusProvisioned = value
    }
}
// SetNumberOfCloudPcStatusProvisioning sets the numberOfCloudPcStatusProvisioning property value. The number of cloud PCs that have a status of provisioning. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioning(value *int32)() {
    if m != nil {
        m.numberOfCloudPcStatusProvisioning = value
    }
}
// SetNumberOfCloudPcStatusUnknown sets the numberOfCloudPcStatusUnknown property value. The number of cloud PCs that have a status of unknown. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUnknown(value *int32)() {
    if m != nil {
        m.numberOfCloudPcStatusUnknown = value
    }
}
// SetNumberOfCloudPcStatusUpgrading sets the numberOfCloudPcStatusUpgrading property value. The number of cloud PCs that have a status of upgrading. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUpgrading(value *int32)() {
    if m != nil {
        m.numberOfCloudPcStatusUpgrading = value
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTenantDisplayName(value *string)() {
    if m != nil {
        m.tenantDisplayName = value
    }
}
// SetTenantId sets the tenantId property value. 
func (m *CloudPcOverview) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetTotalCloudPcConnectionStatus sets the totalCloudPcConnectionStatus property value. The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTotalCloudPcConnectionStatus(value *int32)() {
    if m != nil {
        m.totalCloudPcConnectionStatus = value
    }
}
// SetTotalCloudPcStatus sets the totalCloudPcStatus property value. The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTotalCloudPcStatus(value *int32)() {
    if m != nil {
        m.totalCloudPcStatus = value
    }
}
