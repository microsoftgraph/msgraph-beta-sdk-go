package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailFolderable 
type MailFolderable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildFolderCount()(*int32)
    GetChildFolders()([]MailFolderable)
    GetDisplayName()(*string)
    GetIsHidden()(*bool)
    GetMessageRules()([]MessageRuleable)
    GetMessages()([]Messageable)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetParentFolderId()(*string)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    GetTotalItemCount()(*int32)
    GetType()(*string)
    GetUnreadItemCount()(*int32)
    GetUserConfigurations()([]UserConfigurationable)
    GetWellKnownName()(*string)
    SetChildFolderCount(value *int32)()
    SetChildFolders(value []MailFolderable)()
    SetDisplayName(value *string)()
    SetIsHidden(value *bool)()
    SetMessageRules(value []MessageRuleable)()
    SetMessages(value []Messageable)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetParentFolderId(value *string)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
    SetTotalItemCount(value *int32)()
    SetType(value *string)()
    SetUnreadItemCount(value *int32)()
    SetUserConfigurations(value []UserConfigurationable)()
    SetWellKnownName(value *string)()
}
