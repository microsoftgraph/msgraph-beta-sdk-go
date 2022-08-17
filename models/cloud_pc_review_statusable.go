package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcReviewStatusable 
type CloudPcReviewStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureStorageAccountId()(*string)
    GetAzureStorageAccountName()(*string)
    GetAzureStorageContainerName()(*string)
    GetInReview()(*bool)
    GetOdataType()(*string)
    GetRestorePointDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReviewStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubscriptionId()(*string)
    GetSubscriptionName()(*string)
    GetUserAccessLevel()(*CloudPcUserAccessLevel)
    SetAzureStorageAccountId(value *string)()
    SetAzureStorageAccountName(value *string)()
    SetAzureStorageContainerName(value *string)()
    SetInReview(value *bool)()
    SetOdataType(value *string)()
    SetRestorePointDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReviewStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubscriptionId(value *string)()
    SetSubscriptionName(value *string)()
    SetUserAccessLevel(value *CloudPcUserAccessLevel)()
}
