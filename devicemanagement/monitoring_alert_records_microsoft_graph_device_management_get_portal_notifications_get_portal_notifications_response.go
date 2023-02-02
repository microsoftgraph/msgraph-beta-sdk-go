package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf "github.com/microsoftgraph/msgraph-beta-sdk-go/models/devicemanagement"
)

// MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse 
type MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
    // The value property
    value []i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable
}
// NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse instantiates a new MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse and sets the default values.
func NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse()(*MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse) {
    m := &MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.CreatePortalNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable, len(val))
            for i, v := range val {
                res[i] = v.(i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse) GetValue()([]i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MonitoringAlertRecordsMicrosoftGraphDeviceManagementGetPortalNotificationsGetPortalNotificationsResponse) SetValue(value []i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable)() {
    m.value = value
}
