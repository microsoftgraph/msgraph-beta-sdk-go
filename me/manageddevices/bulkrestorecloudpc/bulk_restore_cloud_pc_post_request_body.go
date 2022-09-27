package bulkrestorecloudpc

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// BulkRestoreCloudPcPostRequestBody provides operations to call the bulkRestoreCloudPc method.
type BulkRestoreCloudPcPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The managedDeviceIds property
    managedDeviceIds []string
    // The restorePointDateTime property
    restorePointDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The timeRange property
    timeRange *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestoreTimeRange
}
// NewBulkRestoreCloudPcPostRequestBody instantiates a new bulkRestoreCloudPcPostRequestBody and sets the default values.
func NewBulkRestoreCloudPcPostRequestBody()(*BulkRestoreCloudPcPostRequestBody) {
    m := &BulkRestoreCloudPcPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBulkRestoreCloudPcPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBulkRestoreCloudPcPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBulkRestoreCloudPcPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BulkRestoreCloudPcPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BulkRestoreCloudPcPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedDeviceIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetManagedDeviceIds)
    res["restorePointDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetRestorePointDateTime)
    res["timeRange"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseRestoreTimeRange , m.SetTimeRange)
    return res
}
// GetManagedDeviceIds gets the managedDeviceIds property value. The managedDeviceIds property
func (m *BulkRestoreCloudPcPostRequestBody) GetManagedDeviceIds()([]string) {
    return m.managedDeviceIds
}
// GetRestorePointDateTime gets the restorePointDateTime property value. The restorePointDateTime property
func (m *BulkRestoreCloudPcPostRequestBody) GetRestorePointDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.restorePointDateTime
}
// GetTimeRange gets the timeRange property value. The timeRange property
func (m *BulkRestoreCloudPcPostRequestBody) GetTimeRange()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestoreTimeRange) {
    return m.timeRange
}
// Serialize serializes information the current object
func (m *BulkRestoreCloudPcPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManagedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("managedDeviceIds", m.GetManagedDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("restorePointDateTime", m.GetRestorePointDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetTimeRange() != nil {
        cast := (*m.GetTimeRange()).String()
        err := writer.WriteStringValue("timeRange", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BulkRestoreCloudPcPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManagedDeviceIds sets the managedDeviceIds property value. The managedDeviceIds property
func (m *BulkRestoreCloudPcPostRequestBody) SetManagedDeviceIds(value []string)() {
    m.managedDeviceIds = value
}
// SetRestorePointDateTime sets the restorePointDateTime property value. The restorePointDateTime property
func (m *BulkRestoreCloudPcPostRequestBody) SetRestorePointDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.restorePointDateTime = value
}
// SetTimeRange sets the timeRange property value. The timeRange property
func (m *BulkRestoreCloudPcPostRequestBody) SetTimeRange(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestoreTimeRange)() {
    m.timeRange = value
}
