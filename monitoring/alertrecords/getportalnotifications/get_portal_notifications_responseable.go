package getportalnotifications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf "github.com/microsoftgraph/msgraph-beta-sdk-go/models/devicemanagement"
)

// GetPortalNotificationsResponseable 
type GetPortalNotificationsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable)
    SetValue(value []i2edb12705e6a63a8a0fb3f8c7a11f4ab12f4be764e61fa1094f401595fb171bf.PortalNotificationable)()
}
