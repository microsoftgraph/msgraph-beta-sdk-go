package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponse provides operations to call the getPortalNotifications method.
type DeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewDeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponse instantiates a new DeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponse and sets the default values.
func NewDeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponse()(*DeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponse) {
    m := &DeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementMonitoringAlertRecordsGetPortalNotificationsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
