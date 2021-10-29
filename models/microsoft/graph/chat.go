package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Chat struct {
    Entity
    // Specifies the type of chat. Possible values are:group, oneOnOne and meeting.
    chatType *ChatType;
    // Date and time at which the chat was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A collection of all the apps in the chat. Nullable.
    installedApps []TeamsAppInstallation;
    // Preview of the last message sent in the chat. Null if no messages have been sent in the chat. Currently, only the list chats operation supports this property.
    lastMessagePreview *ChatMessageInfo;
    // Date and time at which the chat was renamed or list of members were last changed. Read-only.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A collection of all the members in the chat. Nullable.
    members []ConversationMember;
    // A collection of all the messages in the chat. Nullable.
    messages []ChatMessage;
    // Represents details about an online meeting. If the chat isn't associated with an online meeting, the property is empty. Read-only.
    onlineMeetingInfo *TeamworkOnlineMeetingInfo;
    // A collection of all the Teams async operations that ran or are running on the chat. Nullable.
    operations []TeamsAsyncOperation;
    // A collection of permissions granted to apps for the chat.
    permissionGrants []ResourceSpecificPermissionGrant;
    // 
    tabs []TeamsTab;
    // The identifier of the tenant in which the chat was created. Read-only.
    tenantId *string;
    // (Optional) Subject or topic for the chat. Only available for group chats.
    topic *string;
    // Represents caller-specific information about the chat, such as last message read date and time. This property is populated only when the request is made in a delegated context.
    viewpoint *ChatViewpoint;
    // A hyperlink that will go to the chat in Microsoft Teams. This URL should be treated as an opaque blob, and not parsed. Read-only.
    webUrl *string;
}
// Instantiates a new chat and sets the default values.
func NewChat()(*Chat) {
    m := &Chat{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the chatType property value. Specifies the type of chat. Possible values are:group, oneOnOne and meeting.
func (m *Chat) GetChatType()(*ChatType) {
    if m == nil {
        return nil
    } else {
        return m.chatType
    }
}
// Gets the createdDateTime property value. Date and time at which the chat was created. Read-only.
func (m *Chat) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the installedApps property value. A collection of all the apps in the chat. Nullable.
func (m *Chat) GetInstalledApps()([]TeamsAppInstallation) {
    if m == nil {
        return nil
    } else {
        return m.installedApps
    }
}
// Gets the lastMessagePreview property value. Preview of the last message sent in the chat. Null if no messages have been sent in the chat. Currently, only the list chats operation supports this property.
func (m *Chat) GetLastMessagePreview()(*ChatMessageInfo) {
    if m == nil {
        return nil
    } else {
        return m.lastMessagePreview
    }
}
// Gets the lastUpdatedDateTime property value. Date and time at which the chat was renamed or list of members were last changed. Read-only.
func (m *Chat) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// Gets the members property value. A collection of all the members in the chat. Nullable.
func (m *Chat) GetMembers()([]ConversationMember) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Gets the messages property value. A collection of all the messages in the chat. Nullable.
func (m *Chat) GetMessages()([]ChatMessage) {
    if m == nil {
        return nil
    } else {
        return m.messages
    }
}
// Gets the onlineMeetingInfo property value. Represents details about an online meeting. If the chat isn't associated with an online meeting, the property is empty. Read-only.
func (m *Chat) GetOnlineMeetingInfo()(*TeamworkOnlineMeetingInfo) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetingInfo
    }
}
// Gets the operations property value. A collection of all the Teams async operations that ran or are running on the chat. Nullable.
func (m *Chat) GetOperations()([]TeamsAsyncOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// Gets the permissionGrants property value. A collection of permissions granted to apps for the chat.
func (m *Chat) GetPermissionGrants()([]ResourceSpecificPermissionGrant) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrants
    }
}
// Gets the tabs property value. 
func (m *Chat) GetTabs()([]TeamsTab) {
    if m == nil {
        return nil
    } else {
        return m.tabs
    }
}
// Gets the tenantId property value. The identifier of the tenant in which the chat was created. Read-only.
func (m *Chat) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Gets the topic property value. (Optional) Subject or topic for the chat. Only available for group chats.
func (m *Chat) GetTopic()(*string) {
    if m == nil {
        return nil
    } else {
        return m.topic
    }
}
// Gets the viewpoint property value. Represents caller-specific information about the chat, such as last message read date and time. This property is populated only when the request is made in a delegated context.
func (m *Chat) GetViewpoint()(*ChatViewpoint) {
    if m == nil {
        return nil
    } else {
        return m.viewpoint
    }
}
// Gets the webUrl property value. A hyperlink that will go to the chat in Microsoft Teams. This URL should be treated as an opaque blob, and not parsed. Read-only.
func (m *Chat) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *Chat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["chatType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatType)
        if err != nil {
            return err
        }
        cast := val.(ChatType)
        m.SetChatType(&cast)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["installedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAppInstallation() })
        if err != nil {
            return err
        }
        res := make([]TeamsAppInstallation, len(val))
        for i, v := range val {
            res[i] = *(v.(*TeamsAppInstallation))
        }
        m.SetInstalledApps(res)
        return nil
    }
    res["lastMessagePreview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageInfo() })
        if err != nil {
            return err
        }
        m.SetLastMessagePreview(val.(*ChatMessageInfo))
        return nil
    }
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastUpdatedDateTime(val)
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConversationMember() })
        if err != nil {
            return err
        }
        res := make([]ConversationMember, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConversationMember))
        }
        m.SetMembers(res)
        return nil
    }
    res["messages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessage() })
        if err != nil {
            return err
        }
        res := make([]ChatMessage, len(val))
        for i, v := range val {
            res[i] = *(v.(*ChatMessage))
        }
        m.SetMessages(res)
        return nil
    }
    res["onlineMeetingInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkOnlineMeetingInfo() })
        if err != nil {
            return err
        }
        m.SetOnlineMeetingInfo(val.(*TeamworkOnlineMeetingInfo))
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAsyncOperation() })
        if err != nil {
            return err
        }
        res := make([]TeamsAsyncOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*TeamsAsyncOperation))
        }
        m.SetOperations(res)
        return nil
    }
    res["permissionGrants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceSpecificPermissionGrant() })
        if err != nil {
            return err
        }
        res := make([]ResourceSpecificPermissionGrant, len(val))
        for i, v := range val {
            res[i] = *(v.(*ResourceSpecificPermissionGrant))
        }
        m.SetPermissionGrants(res)
        return nil
    }
    res["tabs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsTab() })
        if err != nil {
            return err
        }
        res := make([]TeamsTab, len(val))
        for i, v := range val {
            res[i] = *(v.(*TeamsTab))
        }
        m.SetTabs(res)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    res["topic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTopic(val)
        return nil
    }
    res["viewpoint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatViewpoint() })
        if err != nil {
            return err
        }
        m.SetViewpoint(val.(*ChatViewpoint))
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *Chat) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Chat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChatType() != nil {
        cast := m.GetChatType().String()
        err = writer.WriteStringValue("chatType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInstalledApps()))
        for i, v := range m.GetInstalledApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("installedApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastMessagePreview", m.GetLastMessagePreview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onlineMeetingInfo", m.GetOnlineMeetingInfo())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPermissionGrants()))
        for i, v := range m.GetPermissionGrants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("permissionGrants", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTabs()))
        for i, v := range m.GetTabs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tabs", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("topic", m.GetTopic())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("viewpoint", m.GetViewpoint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the chatType property value. Specifies the type of chat. Possible values are:group, oneOnOne and meeting.
// Parameters:
//  - value : Value to set for the chatType property.
func (m *Chat) SetChatType(value *ChatType)() {
    m.chatType = value
}
// Sets the createdDateTime property value. Date and time at which the chat was created. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Chat) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the installedApps property value. A collection of all the apps in the chat. Nullable.
// Parameters:
//  - value : Value to set for the installedApps property.
func (m *Chat) SetInstalledApps(value []TeamsAppInstallation)() {
    m.installedApps = value
}
// Sets the lastMessagePreview property value. Preview of the last message sent in the chat. Null if no messages have been sent in the chat. Currently, only the list chats operation supports this property.
// Parameters:
//  - value : Value to set for the lastMessagePreview property.
func (m *Chat) SetLastMessagePreview(value *ChatMessageInfo)() {
    m.lastMessagePreview = value
}
// Sets the lastUpdatedDateTime property value. Date and time at which the chat was renamed or list of members were last changed. Read-only.
// Parameters:
//  - value : Value to set for the lastUpdatedDateTime property.
func (m *Chat) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// Sets the members property value. A collection of all the members in the chat. Nullable.
// Parameters:
//  - value : Value to set for the members property.
func (m *Chat) SetMembers(value []ConversationMember)() {
    m.members = value
}
// Sets the messages property value. A collection of all the messages in the chat. Nullable.
// Parameters:
//  - value : Value to set for the messages property.
func (m *Chat) SetMessages(value []ChatMessage)() {
    m.messages = value
}
// Sets the onlineMeetingInfo property value. Represents details about an online meeting. If the chat isn't associated with an online meeting, the property is empty. Read-only.
// Parameters:
//  - value : Value to set for the onlineMeetingInfo property.
func (m *Chat) SetOnlineMeetingInfo(value *TeamworkOnlineMeetingInfo)() {
    m.onlineMeetingInfo = value
}
// Sets the operations property value. A collection of all the Teams async operations that ran or are running on the chat. Nullable.
// Parameters:
//  - value : Value to set for the operations property.
func (m *Chat) SetOperations(value []TeamsAsyncOperation)() {
    m.operations = value
}
// Sets the permissionGrants property value. A collection of permissions granted to apps for the chat.
// Parameters:
//  - value : Value to set for the permissionGrants property.
func (m *Chat) SetPermissionGrants(value []ResourceSpecificPermissionGrant)() {
    m.permissionGrants = value
}
// Sets the tabs property value. 
// Parameters:
//  - value : Value to set for the tabs property.
func (m *Chat) SetTabs(value []TeamsTab)() {
    m.tabs = value
}
// Sets the tenantId property value. The identifier of the tenant in which the chat was created. Read-only.
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *Chat) SetTenantId(value *string)() {
    m.tenantId = value
}
// Sets the topic property value. (Optional) Subject or topic for the chat. Only available for group chats.
// Parameters:
//  - value : Value to set for the topic property.
func (m *Chat) SetTopic(value *string)() {
    m.topic = value
}
// Sets the viewpoint property value. Represents caller-specific information about the chat, such as last message read date and time. This property is populated only when the request is made in a delegated context.
// Parameters:
//  - value : Value to set for the viewpoint property.
func (m *Chat) SetViewpoint(value *ChatViewpoint)() {
    m.viewpoint = value
}
// Sets the webUrl property value. A hyperlink that will go to the chat in Microsoft Teams. This URL should be treated as an opaque blob, and not parsed. Read-only.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *Chat) SetWebUrl(value *string)() {
    m.webUrl = value
}
