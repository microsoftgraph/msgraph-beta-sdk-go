package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CloudPcOverview 
type CloudPcOverview struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewCloudPcOverview instantiates a new cloudPcOverview and sets the default values.
func NewCloudPcOverview()(*CloudPcOverview) {
    m := &CloudPcOverview{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateCloudPcOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcOverview(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastRefreshedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedDateTime(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusFailed(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusPassed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusPassed(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusPending"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusPending(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusRunning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusRunning(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusUnkownFutureValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusUnkownFutureValue(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusDeprovisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusDeprovisioning(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusFailed(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusInGracePeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusInGracePeriod(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusNotProvisioned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusNotProvisioned(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusProvisioned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusProvisioned(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusProvisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusProvisioning(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusUnknown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusUnknown(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusUpgrading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusUpgrading(val)
        }
        return nil
    }
    res["tenantDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["totalBusinessLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalBusinessLicenses(val)
        }
        return nil
    }
    res["totalCloudPcConnectionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCloudPcConnectionStatus(val)
        }
        return nil
    }
    res["totalCloudPcStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCloudPcStatus(val)
        }
        return nil
    }
    res["totalEnterpriseLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalEnterpriseLicenses(val)
        }
        return nil
    }
    return res
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CloudPcOverview) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastRefreshedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNumberOfCloudPcConnectionStatusFailed gets the numberOfCloudPcConnectionStatusFailed property value. The number of cloud PC connections that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusFailed()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcConnectionStatusFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcConnectionStatusPassed gets the numberOfCloudPcConnectionStatusPassed property value. The number of cloud PC connections that have a status of passed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPassed()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcConnectionStatusPassed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcConnectionStatusPending gets the numberOfCloudPcConnectionStatusPending property value. The number of cloud PC connections that have a status of pending. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPending()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcConnectionStatusPending")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcConnectionStatusRunning gets the numberOfCloudPcConnectionStatusRunning property value. The number of cloud PC connections that have a status of running. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusRunning()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcConnectionStatusRunning")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcConnectionStatusUnkownFutureValue gets the numberOfCloudPcConnectionStatusUnkownFutureValue property value. The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusUnkownFutureValue()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcConnectionStatusUnkownFutureValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcStatusDeprovisioning gets the numberOfCloudPcStatusDeprovisioning property value. The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusDeprovisioning()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcStatusDeprovisioning")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcStatusFailed gets the numberOfCloudPcStatusFailed property value. The number of cloud PCs that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusFailed()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcStatusFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcStatusInGracePeriod gets the numberOfCloudPcStatusInGracePeriod property value. The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusInGracePeriod()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcStatusInGracePeriod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcStatusNotProvisioned gets the numberOfCloudPcStatusNotProvisioned property value. The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusNotProvisioned()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcStatusNotProvisioned")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcStatusProvisioned gets the numberOfCloudPcStatusProvisioned property value. The number of cloud PCs that have a status of provisioned. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioned()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcStatusProvisioned")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcStatusProvisioning gets the numberOfCloudPcStatusProvisioning property value. The number of cloud PCs that have a status of provisioning. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioning()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcStatusProvisioning")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcStatusUnknown gets the numberOfCloudPcStatusUnknown property value. The number of cloud PCs that have a status of unknown. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUnknown()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcStatusUnknown")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNumberOfCloudPcStatusUpgrading gets the numberOfCloudPcStatusUpgrading property value. The number of cloud PCs that have a status of upgrading. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUpgrading()(*int32) {
    val, err := m.GetBackingStore().Get("numberOfCloudPcStatusUpgrading")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTenantDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("tenantDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *CloudPcOverview) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalBusinessLicenses gets the totalBusinessLicenses property value. The total number of cloud PC devices that have the Business SKU. Optional. Read-only.
func (m *CloudPcOverview) GetTotalBusinessLicenses()(*int32) {
    val, err := m.GetBackingStore().Get("totalBusinessLicenses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalCloudPcConnectionStatus gets the totalCloudPcConnectionStatus property value. The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTotalCloudPcConnectionStatus()(*int32) {
    val, err := m.GetBackingStore().Get("totalCloudPcConnectionStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalCloudPcStatus gets the totalCloudPcStatus property value. The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTotalCloudPcStatus()(*int32) {
    val, err := m.GetBackingStore().Get("totalCloudPcStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalEnterpriseLicenses gets the totalEnterpriseLicenses property value. The total number of cloud PC devices that have the Enterprise SKU. Optional. Read-only.
func (m *CloudPcOverview) GetTotalEnterpriseLicenses()(*int32) {
    val, err := m.GetBackingStore().Get("totalEnterpriseLicenses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("totalBusinessLicenses", m.GetTotalBusinessLicenses())
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
    {
        err = writer.WriteInt32Value("totalEnterpriseLicenses", m.GetTotalEnterpriseLicenses())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CloudPcOverview) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastRefreshedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcConnectionStatusFailed sets the numberOfCloudPcConnectionStatusFailed property value. The number of cloud PC connections that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusFailed(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcConnectionStatusFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcConnectionStatusPassed sets the numberOfCloudPcConnectionStatusPassed property value. The number of cloud PC connections that have a status of passed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPassed(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcConnectionStatusPassed", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcConnectionStatusPending sets the numberOfCloudPcConnectionStatusPending property value. The number of cloud PC connections that have a status of pending. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPending(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcConnectionStatusPending", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcConnectionStatusRunning sets the numberOfCloudPcConnectionStatusRunning property value. The number of cloud PC connections that have a status of running. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusRunning(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcConnectionStatusRunning", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcConnectionStatusUnkownFutureValue sets the numberOfCloudPcConnectionStatusUnkownFutureValue property value. The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusUnkownFutureValue(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcConnectionStatusUnkownFutureValue", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcStatusDeprovisioning sets the numberOfCloudPcStatusDeprovisioning property value. The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusDeprovisioning(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcStatusDeprovisioning", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcStatusFailed sets the numberOfCloudPcStatusFailed property value. The number of cloud PCs that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusFailed(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcStatusFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcStatusInGracePeriod sets the numberOfCloudPcStatusInGracePeriod property value. The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusInGracePeriod(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcStatusInGracePeriod", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcStatusNotProvisioned sets the numberOfCloudPcStatusNotProvisioned property value. The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusNotProvisioned(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcStatusNotProvisioned", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcStatusProvisioned sets the numberOfCloudPcStatusProvisioned property value. The number of cloud PCs that have a status of provisioned. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioned(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcStatusProvisioned", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcStatusProvisioning sets the numberOfCloudPcStatusProvisioning property value. The number of cloud PCs that have a status of provisioning. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioning(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcStatusProvisioning", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcStatusUnknown sets the numberOfCloudPcStatusUnknown property value. The number of cloud PCs that have a status of unknown. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUnknown(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcStatusUnknown", value)
    if err != nil {
        panic(err)
    }
}
// SetNumberOfCloudPcStatusUpgrading sets the numberOfCloudPcStatusUpgrading property value. The number of cloud PCs that have a status of upgrading. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUpgrading(value *int32)() {
    err := m.GetBackingStore().Set("numberOfCloudPcStatusUpgrading", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTenantDisplayName(value *string)() {
    err := m.GetBackingStore().Set("tenantDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *CloudPcOverview) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalBusinessLicenses sets the totalBusinessLicenses property value. The total number of cloud PC devices that have the Business SKU. Optional. Read-only.
func (m *CloudPcOverview) SetTotalBusinessLicenses(value *int32)() {
    err := m.GetBackingStore().Set("totalBusinessLicenses", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalCloudPcConnectionStatus sets the totalCloudPcConnectionStatus property value. The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTotalCloudPcConnectionStatus(value *int32)() {
    err := m.GetBackingStore().Set("totalCloudPcConnectionStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalCloudPcStatus sets the totalCloudPcStatus property value. The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTotalCloudPcStatus(value *int32)() {
    err := m.GetBackingStore().Set("totalCloudPcStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalEnterpriseLicenses sets the totalEnterpriseLicenses property value. The total number of cloud PC devices that have the Enterprise SKU. Optional. Read-only.
func (m *CloudPcOverview) SetTotalEnterpriseLicenses(value *int32)() {
    err := m.GetBackingStore().Set("totalEnterpriseLicenses", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcOverviewable 
type CloudPcOverviewable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumberOfCloudPcConnectionStatusFailed()(*int32)
    GetNumberOfCloudPcConnectionStatusPassed()(*int32)
    GetNumberOfCloudPcConnectionStatusPending()(*int32)
    GetNumberOfCloudPcConnectionStatusRunning()(*int32)
    GetNumberOfCloudPcConnectionStatusUnkownFutureValue()(*int32)
    GetNumberOfCloudPcStatusDeprovisioning()(*int32)
    GetNumberOfCloudPcStatusFailed()(*int32)
    GetNumberOfCloudPcStatusInGracePeriod()(*int32)
    GetNumberOfCloudPcStatusNotProvisioned()(*int32)
    GetNumberOfCloudPcStatusProvisioned()(*int32)
    GetNumberOfCloudPcStatusProvisioning()(*int32)
    GetNumberOfCloudPcStatusUnknown()(*int32)
    GetNumberOfCloudPcStatusUpgrading()(*int32)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetTotalBusinessLicenses()(*int32)
    GetTotalCloudPcConnectionStatus()(*int32)
    GetTotalCloudPcStatus()(*int32)
    GetTotalEnterpriseLicenses()(*int32)
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumberOfCloudPcConnectionStatusFailed(value *int32)()
    SetNumberOfCloudPcConnectionStatusPassed(value *int32)()
    SetNumberOfCloudPcConnectionStatusPending(value *int32)()
    SetNumberOfCloudPcConnectionStatusRunning(value *int32)()
    SetNumberOfCloudPcConnectionStatusUnkownFutureValue(value *int32)()
    SetNumberOfCloudPcStatusDeprovisioning(value *int32)()
    SetNumberOfCloudPcStatusFailed(value *int32)()
    SetNumberOfCloudPcStatusInGracePeriod(value *int32)()
    SetNumberOfCloudPcStatusNotProvisioned(value *int32)()
    SetNumberOfCloudPcStatusProvisioned(value *int32)()
    SetNumberOfCloudPcStatusProvisioning(value *int32)()
    SetNumberOfCloudPcStatusUnknown(value *int32)()
    SetNumberOfCloudPcStatusUpgrading(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetTotalBusinessLicenses(value *int32)()
    SetTotalCloudPcConnectionStatus(value *int32)()
    SetTotalCloudPcStatus(value *int32)()
    SetTotalEnterpriseLicenses(value *int32)()
}
