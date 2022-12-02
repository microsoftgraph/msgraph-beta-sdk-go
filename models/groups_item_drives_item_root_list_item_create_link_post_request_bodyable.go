package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemDrivesItemRootListItemCreateLinkPostRequestBodyable 
type GroupsItemDrivesItemRootListItemCreateLinkPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPassword()(*string)
    GetRecipients()([]DriveRecipientable)
    GetRetainInheritedPermissions()(*bool)
    GetScope()(*string)
    GetType()(*string)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPassword(value *string)()
    SetRecipients(value []DriveRecipientable)()
    SetRetainInheritedPermissions(value *bool)()
    SetScope(value *string)()
    SetType(value *string)()
}
