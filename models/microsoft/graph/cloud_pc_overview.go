package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CloudPcOverview struct {
    Entity
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    numberOfCloudPcConnectionStatusFailed *int32;
    numberOfCloudPcConnectionStatusPassed *int32;
    numberOfCloudPcConnectionStatusPending *int32;
    numberOfCloudPcConnectionStatusRunning *int32;
    numberOfCloudPcConnectionStatusUnkownFutureValue *int32;
    numberOfCloudPcStatusDeprovisioning *int32;
    numberOfCloudPcStatusFailed *int32;
    numberOfCloudPcStatusInGracePeriod *int32;
    numberOfCloudPcStatusNotProvisioned *int32;
    numberOfCloudPcStatusProvisioned *int32;
    numberOfCloudPcStatusProvisioning *int32;
    numberOfCloudPcStatusUnknown *int32;
    numberOfCloudPcStatusUpgrading *int32;
    tenantDisplayName *string;
    tenantId *string;
    totalCloudPcConnectionStatus *int32;
    totalCloudPcStatus *int32;
}
func NewCloudPcOverview()(*CloudPcOverview) {
    m := &CloudPcOverview{
        Entity: *NewEntity(),
    }
    return m
}
func (m *CloudPcOverview) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusFailed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusFailed
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPassed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusPassed
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPending()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusPending
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusRunning()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusRunning
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusUnkownFutureValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcConnectionStatusUnkownFutureValue
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcStatusDeprovisioning()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusDeprovisioning
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcStatusFailed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusFailed
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcStatusInGracePeriod()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusInGracePeriod
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcStatusNotProvisioned()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusNotProvisioned
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioned()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusProvisioned
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioning()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusProvisioning
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUnknown()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusUnknown
    }
}
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUpgrading()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfCloudPcStatusUpgrading
    }
}
func (m *CloudPcOverview) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
func (m *CloudPcOverview) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
func (m *CloudPcOverview) GetTotalCloudPcConnectionStatus()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalCloudPcConnectionStatus
    }
}
func (m *CloudPcOverview) GetTotalCloudPcStatus()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalCloudPcStatus
    }
}
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
func (m *CloudPcOverview) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusFailed(value *int32)() {
    m.numberOfCloudPcConnectionStatusFailed = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPassed(value *int32)() {
    m.numberOfCloudPcConnectionStatusPassed = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPending(value *int32)() {
    m.numberOfCloudPcConnectionStatusPending = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusRunning(value *int32)() {
    m.numberOfCloudPcConnectionStatusRunning = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusUnkownFutureValue(value *int32)() {
    m.numberOfCloudPcConnectionStatusUnkownFutureValue = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcStatusDeprovisioning(value *int32)() {
    m.numberOfCloudPcStatusDeprovisioning = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcStatusFailed(value *int32)() {
    m.numberOfCloudPcStatusFailed = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcStatusInGracePeriod(value *int32)() {
    m.numberOfCloudPcStatusInGracePeriod = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcStatusNotProvisioned(value *int32)() {
    m.numberOfCloudPcStatusNotProvisioned = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioned(value *int32)() {
    m.numberOfCloudPcStatusProvisioned = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioning(value *int32)() {
    m.numberOfCloudPcStatusProvisioning = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUnknown(value *int32)() {
    m.numberOfCloudPcStatusUnknown = value
}
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUpgrading(value *int32)() {
    m.numberOfCloudPcStatusUpgrading = value
}
func (m *CloudPcOverview) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
func (m *CloudPcOverview) SetTenantId(value *string)() {
    m.tenantId = value
}
func (m *CloudPcOverview) SetTotalCloudPcConnectionStatus(value *int32)() {
    m.totalCloudPcConnectionStatus = value
}
func (m *CloudPcOverview) SetTotalCloudPcStatus(value *int32)() {
    m.totalCloudPcStatus = value
}
