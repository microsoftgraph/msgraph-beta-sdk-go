package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CloudPcOverview provides operations to manage the collection of activityStatistics entities.
type CloudPcOverview struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of cloud PC connections that have a status of failed. Optional. Read-only.
    numberOfCloudPcConnectionStatusFailed *int32
    // The number of cloud PC connections that have a status of passed. Optional. Read-only.
    numberOfCloudPcConnectionStatusPassed *int32
    // The number of cloud PC connections that have a status of pending. Optional. Read-only.
    numberOfCloudPcConnectionStatusPending *int32
    // The number of cloud PC connections that have a status of running. Optional. Read-only.
    numberOfCloudPcConnectionStatusRunning *int32
    // The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
    numberOfCloudPcConnectionStatusUnkownFutureValue *int32
    // The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
    numberOfCloudPcStatusDeprovisioning *int32
    // The number of cloud PCs that have a status of failed. Optional. Read-only.
    numberOfCloudPcStatusFailed *int32
    // The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
    numberOfCloudPcStatusInGracePeriod *int32
    // The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
    numberOfCloudPcStatusNotProvisioned *int32
    // The number of cloud PCs that have a status of provisioned. Optional. Read-only.
    numberOfCloudPcStatusProvisioned *int32
    // The number of cloud PCs that have a status of provisioning. Optional. Read-only.
    numberOfCloudPcStatusProvisioning *int32
    // The number of cloud PCs that have a status of unknown. Optional. Read-only.
    numberOfCloudPcStatusUnknown *int32
    // The number of cloud PCs that have a status of upgrading. Optional. Read-only.
    numberOfCloudPcStatusUpgrading *int32
    // The display name for the managed tenant. Optional. Read-only.
    tenantDisplayName *string
    // The tenantId property
    tenantId *string
    // The totalBusinessLicenses property
    totalBusinessLicenses *int32
    // The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
    totalCloudPcConnectionStatus *int32
    // The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
    totalCloudPcStatus *int32
    // The totalEnterpriseLicenses property
    totalEnterpriseLicenses *int32
}
// NewCloudPcOverview instantiates a new cloudPcOverview and sets the default values.
func NewCloudPcOverview()(*CloudPcOverview) {
    m := &CloudPcOverview{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.cloudPcOverview";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcOverview(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastRefreshedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastRefreshedDateTime)
    res["numberOfCloudPcConnectionStatusFailed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcConnectionStatusFailed)
    res["numberOfCloudPcConnectionStatusPassed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcConnectionStatusPassed)
    res["numberOfCloudPcConnectionStatusPending"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcConnectionStatusPending)
    res["numberOfCloudPcConnectionStatusRunning"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcConnectionStatusRunning)
    res["numberOfCloudPcConnectionStatusUnkownFutureValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcConnectionStatusUnkownFutureValue)
    res["numberOfCloudPcStatusDeprovisioning"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcStatusDeprovisioning)
    res["numberOfCloudPcStatusFailed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcStatusFailed)
    res["numberOfCloudPcStatusInGracePeriod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcStatusInGracePeriod)
    res["numberOfCloudPcStatusNotProvisioned"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcStatusNotProvisioned)
    res["numberOfCloudPcStatusProvisioned"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcStatusProvisioned)
    res["numberOfCloudPcStatusProvisioning"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcStatusProvisioning)
    res["numberOfCloudPcStatusUnknown"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcStatusUnknown)
    res["numberOfCloudPcStatusUpgrading"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNumberOfCloudPcStatusUpgrading)
    res["tenantDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantDisplayName)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    res["totalBusinessLicenses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalBusinessLicenses)
    res["totalCloudPcConnectionStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalCloudPcConnectionStatus)
    res["totalCloudPcStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalCloudPcStatus)
    res["totalEnterpriseLicenses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalEnterpriseLicenses)
    return res
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CloudPcOverview) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRefreshedDateTime
}
// GetNumberOfCloudPcConnectionStatusFailed gets the numberOfCloudPcConnectionStatusFailed property value. The number of cloud PC connections that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusFailed()(*int32) {
    return m.numberOfCloudPcConnectionStatusFailed
}
// GetNumberOfCloudPcConnectionStatusPassed gets the numberOfCloudPcConnectionStatusPassed property value. The number of cloud PC connections that have a status of passed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPassed()(*int32) {
    return m.numberOfCloudPcConnectionStatusPassed
}
// GetNumberOfCloudPcConnectionStatusPending gets the numberOfCloudPcConnectionStatusPending property value. The number of cloud PC connections that have a status of pending. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPending()(*int32) {
    return m.numberOfCloudPcConnectionStatusPending
}
// GetNumberOfCloudPcConnectionStatusRunning gets the numberOfCloudPcConnectionStatusRunning property value. The number of cloud PC connections that have a status of running. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusRunning()(*int32) {
    return m.numberOfCloudPcConnectionStatusRunning
}
// GetNumberOfCloudPcConnectionStatusUnkownFutureValue gets the numberOfCloudPcConnectionStatusUnkownFutureValue property value. The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusUnkownFutureValue()(*int32) {
    return m.numberOfCloudPcConnectionStatusUnkownFutureValue
}
// GetNumberOfCloudPcStatusDeprovisioning gets the numberOfCloudPcStatusDeprovisioning property value. The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusDeprovisioning()(*int32) {
    return m.numberOfCloudPcStatusDeprovisioning
}
// GetNumberOfCloudPcStatusFailed gets the numberOfCloudPcStatusFailed property value. The number of cloud PCs that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusFailed()(*int32) {
    return m.numberOfCloudPcStatusFailed
}
// GetNumberOfCloudPcStatusInGracePeriod gets the numberOfCloudPcStatusInGracePeriod property value. The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusInGracePeriod()(*int32) {
    return m.numberOfCloudPcStatusInGracePeriod
}
// GetNumberOfCloudPcStatusNotProvisioned gets the numberOfCloudPcStatusNotProvisioned property value. The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusNotProvisioned()(*int32) {
    return m.numberOfCloudPcStatusNotProvisioned
}
// GetNumberOfCloudPcStatusProvisioned gets the numberOfCloudPcStatusProvisioned property value. The number of cloud PCs that have a status of provisioned. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioned()(*int32) {
    return m.numberOfCloudPcStatusProvisioned
}
// GetNumberOfCloudPcStatusProvisioning gets the numberOfCloudPcStatusProvisioning property value. The number of cloud PCs that have a status of provisioning. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioning()(*int32) {
    return m.numberOfCloudPcStatusProvisioning
}
// GetNumberOfCloudPcStatusUnknown gets the numberOfCloudPcStatusUnknown property value. The number of cloud PCs that have a status of unknown. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUnknown()(*int32) {
    return m.numberOfCloudPcStatusUnknown
}
// GetNumberOfCloudPcStatusUpgrading gets the numberOfCloudPcStatusUpgrading property value. The number of cloud PCs that have a status of upgrading. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUpgrading()(*int32) {
    return m.numberOfCloudPcStatusUpgrading
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTenantDisplayName()(*string) {
    return m.tenantDisplayName
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *CloudPcOverview) GetTenantId()(*string) {
    return m.tenantId
}
// GetTotalBusinessLicenses gets the totalBusinessLicenses property value. The totalBusinessLicenses property
func (m *CloudPcOverview) GetTotalBusinessLicenses()(*int32) {
    return m.totalBusinessLicenses
}
// GetTotalCloudPcConnectionStatus gets the totalCloudPcConnectionStatus property value. The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTotalCloudPcConnectionStatus()(*int32) {
    return m.totalCloudPcConnectionStatus
}
// GetTotalCloudPcStatus gets the totalCloudPcStatus property value. The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTotalCloudPcStatus()(*int32) {
    return m.totalCloudPcStatus
}
// GetTotalEnterpriseLicenses gets the totalEnterpriseLicenses property value. The totalEnterpriseLicenses property
func (m *CloudPcOverview) GetTotalEnterpriseLicenses()(*int32) {
    return m.totalEnterpriseLicenses
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
    m.lastRefreshedDateTime = value
}
// SetNumberOfCloudPcConnectionStatusFailed sets the numberOfCloudPcConnectionStatusFailed property value. The number of cloud PC connections that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusFailed(value *int32)() {
    m.numberOfCloudPcConnectionStatusFailed = value
}
// SetNumberOfCloudPcConnectionStatusPassed sets the numberOfCloudPcConnectionStatusPassed property value. The number of cloud PC connections that have a status of passed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPassed(value *int32)() {
    m.numberOfCloudPcConnectionStatusPassed = value
}
// SetNumberOfCloudPcConnectionStatusPending sets the numberOfCloudPcConnectionStatusPending property value. The number of cloud PC connections that have a status of pending. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPending(value *int32)() {
    m.numberOfCloudPcConnectionStatusPending = value
}
// SetNumberOfCloudPcConnectionStatusRunning sets the numberOfCloudPcConnectionStatusRunning property value. The number of cloud PC connections that have a status of running. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusRunning(value *int32)() {
    m.numberOfCloudPcConnectionStatusRunning = value
}
// SetNumberOfCloudPcConnectionStatusUnkownFutureValue sets the numberOfCloudPcConnectionStatusUnkownFutureValue property value. The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusUnkownFutureValue(value *int32)() {
    m.numberOfCloudPcConnectionStatusUnkownFutureValue = value
}
// SetNumberOfCloudPcStatusDeprovisioning sets the numberOfCloudPcStatusDeprovisioning property value. The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusDeprovisioning(value *int32)() {
    m.numberOfCloudPcStatusDeprovisioning = value
}
// SetNumberOfCloudPcStatusFailed sets the numberOfCloudPcStatusFailed property value. The number of cloud PCs that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusFailed(value *int32)() {
    m.numberOfCloudPcStatusFailed = value
}
// SetNumberOfCloudPcStatusInGracePeriod sets the numberOfCloudPcStatusInGracePeriod property value. The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusInGracePeriod(value *int32)() {
    m.numberOfCloudPcStatusInGracePeriod = value
}
// SetNumberOfCloudPcStatusNotProvisioned sets the numberOfCloudPcStatusNotProvisioned property value. The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusNotProvisioned(value *int32)() {
    m.numberOfCloudPcStatusNotProvisioned = value
}
// SetNumberOfCloudPcStatusProvisioned sets the numberOfCloudPcStatusProvisioned property value. The number of cloud PCs that have a status of provisioned. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioned(value *int32)() {
    m.numberOfCloudPcStatusProvisioned = value
}
// SetNumberOfCloudPcStatusProvisioning sets the numberOfCloudPcStatusProvisioning property value. The number of cloud PCs that have a status of provisioning. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioning(value *int32)() {
    m.numberOfCloudPcStatusProvisioning = value
}
// SetNumberOfCloudPcStatusUnknown sets the numberOfCloudPcStatusUnknown property value. The number of cloud PCs that have a status of unknown. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUnknown(value *int32)() {
    m.numberOfCloudPcStatusUnknown = value
}
// SetNumberOfCloudPcStatusUpgrading sets the numberOfCloudPcStatusUpgrading property value. The number of cloud PCs that have a status of upgrading. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUpgrading(value *int32)() {
    m.numberOfCloudPcStatusUpgrading = value
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *CloudPcOverview) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetTotalBusinessLicenses sets the totalBusinessLicenses property value. The totalBusinessLicenses property
func (m *CloudPcOverview) SetTotalBusinessLicenses(value *int32)() {
    m.totalBusinessLicenses = value
}
// SetTotalCloudPcConnectionStatus sets the totalCloudPcConnectionStatus property value. The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTotalCloudPcConnectionStatus(value *int32)() {
    m.totalCloudPcConnectionStatus = value
}
// SetTotalCloudPcStatus sets the totalCloudPcStatus property value. The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTotalCloudPcStatus(value *int32)() {
    m.totalCloudPcStatus = value
}
// SetTotalEnterpriseLicenses sets the totalEnterpriseLicenses property value. The totalEnterpriseLicenses property
func (m *CloudPcOverview) SetTotalEnterpriseLicenses(value *int32)() {
    m.totalEnterpriseLicenses = value
}
