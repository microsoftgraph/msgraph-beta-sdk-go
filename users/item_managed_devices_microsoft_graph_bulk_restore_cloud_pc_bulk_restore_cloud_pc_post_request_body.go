package users

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody 
type ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The managedDeviceIds property
    managedDeviceIds []string
    // The restorePointDateTime property
    restorePointDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The timeRange property
    timeRange *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestoreTimeRange
}
// NewItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody instantiates a new ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody and sets the default values.
func NewItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody()(*ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody) {
    m := &ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetManagedDeviceIds(res)
        }
        return nil
    }
    res["restorePointDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestorePointDateTime(val)
        }
        return nil
    }
    res["timeRange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseRestoreTimeRange)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeRange(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestoreTimeRange))
        }
        return nil
    }
    return res
}
// GetManagedDeviceIds gets the managedDeviceIds property value. The managedDeviceIds property
func (m *ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody) GetManagedDeviceIds()([]string) {
    return m.managedDeviceIds
}
// GetRestorePointDateTime gets the restorePointDateTime property value. The restorePointDateTime property
func (m *ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody) GetRestorePointDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.restorePointDateTime
}
// GetTimeRange gets the timeRange property value. The timeRange property
func (m *ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody) GetTimeRange()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestoreTimeRange) {
    return m.timeRange
}
// Serialize serializes information the current object
func (m *ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetManagedDeviceIds sets the managedDeviceIds property value. The managedDeviceIds property
func (m *ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody) SetManagedDeviceIds(value []string)() {
    m.managedDeviceIds = value
}
// SetRestorePointDateTime sets the restorePointDateTime property value. The restorePointDateTime property
func (m *ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody) SetRestorePointDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.restorePointDateTime = value
}
// SetTimeRange sets the timeRange property value. The timeRange property
func (m *ItemManagedDevicesMicrosoftGraphBulkRestoreCloudPcBulkRestoreCloudPcPostRequestBody) SetTimeRange(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestoreTimeRange)() {
    m.timeRange = value
}
