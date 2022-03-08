package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Chatable 
type Chatable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetChatType()(*ChatType)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInstalledApps()([]TeamsAppInstallationable)
    GetLastMessagePreview()(ChatMessageInfoable)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMembers()([]ConversationMemberable)
    GetMessages()([]ChatMessageable)
    GetOnlineMeetingInfo()(TeamworkOnlineMeetingInfoable)
    GetOperations()([]TeamsAsyncOperationable)
    GetPermissionGrants()([]ResourceSpecificPermissionGrantable)
    GetTabs()([]TeamsTabable)
    GetTenantId()(*string)
    GetTopic()(*string)
    GetViewpoint()(ChatViewpointable)
    GetWebUrl()(*string)
    SetChatType(value *ChatType)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInstalledApps(value []TeamsAppInstallationable)()
    SetLastMessagePreview(value ChatMessageInfoable)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMembers(value []ConversationMemberable)()
    SetMessages(value []ChatMessageable)()
    SetOnlineMeetingInfo(value TeamworkOnlineMeetingInfoable)()
    SetOperations(value []TeamsAsyncOperationable)()
    SetPermissionGrants(value []ResourceSpecificPermissionGrantable)()
    SetTabs(value []TeamsTabable)()
    SetTenantId(value *string)()
    SetTopic(value *string)()
    SetViewpoint(value ChatViewpointable)()
    SetWebUrl(value *string)()
}
