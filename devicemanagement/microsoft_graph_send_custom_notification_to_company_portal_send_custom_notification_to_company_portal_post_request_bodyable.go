package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBodyable 
type MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupsToNotify()([]string)
    GetNotificationBody()(*string)
    GetNotificationTitle()(*string)
    SetGroupsToNotify(value []string)()
    SetNotificationBody(value *string)()
    SetNotificationTitle(value *string)()
}
