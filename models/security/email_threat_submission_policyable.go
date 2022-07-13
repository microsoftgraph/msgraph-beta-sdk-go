package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// EmailThreatSubmissionPolicyable 
type EmailThreatSubmissionPolicyable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomizedNotificationSenderEmailAddress()(*string)
    GetCustomizedReportRecipientEmailAddress()(*string)
    GetIsAlwaysReportEnabledForUsers()(*bool)
    GetIsAskMeEnabledForUsers()(*bool)
    GetIsCustomizedMessageEnabled()(*bool)
    GetIsCustomizedMessageEnabledForPhishing()(*bool)
    GetIsCustomizedNotificationSenderEnabled()(*bool)
    GetIsNeverReportEnabledForUsers()(*bool)
    GetIsOrganizationBrandingEnabled()(*bool)
    GetIsReportFromQuarantineEnabled()(*bool)
    GetIsReportToCustomizedEmailAddressEnabled()(*bool)
    GetIsReportToMicrosoftEnabled()(*bool)
    GetIsReviewEmailNotificationEnabled()(*bool)
    SetCustomizedNotificationSenderEmailAddress(value *string)()
    SetCustomizedReportRecipientEmailAddress(value *string)()
    SetIsAlwaysReportEnabledForUsers(value *bool)()
    SetIsAskMeEnabledForUsers(value *bool)()
    SetIsCustomizedMessageEnabled(value *bool)()
    SetIsCustomizedMessageEnabledForPhishing(value *bool)()
    SetIsCustomizedNotificationSenderEnabled(value *bool)()
    SetIsNeverReportEnabledForUsers(value *bool)()
    SetIsOrganizationBrandingEnabled(value *bool)()
    SetIsReportFromQuarantineEnabled(value *bool)()
    SetIsReportToCustomizedEmailAddressEnabled(value *bool)()
    SetIsReportToMicrosoftEnabled(value *bool)()
    SetIsReviewEmailNotificationEnabled(value *bool)()
}
