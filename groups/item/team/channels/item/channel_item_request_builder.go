package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i005211e94469a36a987062d942e8e3e4b9f530d02f325b79d83acbc72a7296b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/tabs"
    i0f949d8700cdd20dd5464492ee35ce01485a1a3ffdf50bce12c98acee6a8f39b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/members"
    i21e401c8131ed289d503e008a3d43d5445a6979bb2088584c87107e5f9a7efd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/provisionemail"
    i242c35c5532ed25cf08fb018cf513312bd1fa96dd2b1e8b989dd0b1bb2ae55c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/sharedwithteams"
    i7164a133c07ed245aaded02aa0ea05e0f0699c32d668c579eb4a326e743be552 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/completemigration"
    i75075de7416f40b42eeae5c62e3a1830b76c46700fb92956f8a2994416edf871 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/removeemail"
    i7a14e44089123d3294daf6fdf163348604cebc83dee8f9f2c1702c99c3a388ca "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/filesfolder"
    iab02a926f6b68b941538c030223f00d8ee799b702aeb9af0f7dea46b8e014905 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/messages"
    ide7ba64a450154e1df13d6eb89f40b8db2d075fb2910e1ab8e2e618a974585ac "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname"
    i1c349b81f3eac6900d6c57af5d7581e337b2ca590948d172e4adae3852bcc581 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/sharedwithteams/item"
    i50b60f6b70ccabd26064958871bb043fc875c38b3cdd16b1828e04c8a5a21880 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/tabs/item"
    i654c2adcf04d8630b1cb24e632177c047ce85757bf63c3ea55ed853f8c2c06da "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/messages/item"
    i833055126432577ccf7abb5c852b0232dd1f9933c3113ff1d863ba3fde085be2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/channels/item/members/item"
)

// ChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.team entity.
type ChannelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChannelItemRequestBuilderDeleteOptions options for Delete
type ChannelItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChannelItemRequestBuilderGetOptions options for Get
type ChannelItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChannelItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChannelItemRequestBuilderGetQueryParameters the collection of channels and messages associated with the team.
type ChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChannelItemRequestBuilderPatchOptions options for Patch
type ChannelItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CompleteMigration the completeMigration property
func (m *ChannelItemRequestBuilder) CompleteMigration()(*i7164a133c07ed245aaded02aa0ea05e0f0699c32d668c579eb4a326e743be552.CompleteMigrationRequestBuilder) {
    return i7164a133c07ed245aaded02aa0ea05e0f0699c32d668c579eb4a326e743be552.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelItemRequestBuilder) {
    m := &ChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/channels/{channel%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChannelItemRequestBuilder instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property channels for groups
func (m *ChannelItemRequestBuilder) CreateDeleteRequestInformation(options *ChannelItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) CreateGetRequestInformation(options *ChannelItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property channels in groups
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformation(options *ChannelItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property channels for groups
func (m *ChannelItemRequestBuilder) Delete(options *ChannelItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName provides operations to call the doesUserHaveAccess method.
func (m *ChannelItemRequestBuilder) DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName()(*ide7ba64a450154e1df13d6eb89f40b8db2d075fb2910e1ab8e2e618a974585ac.DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    return ide7ba64a450154e1df13d6eb89f40b8db2d075fb2910e1ab8e2e618a974585ac.NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder the filesFolder property
func (m *ChannelItemRequestBuilder) FilesFolder()(*i7a14e44089123d3294daf6fdf163348604cebc83dee8f9f2c1702c99c3a388ca.FilesFolderRequestBuilder) {
    return i7a14e44089123d3294daf6fdf163348604cebc83dee8f9f2c1702c99c3a388ca.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) Get(options *ChannelItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable), nil
}
// Members the members property
func (m *ChannelItemRequestBuilder) Members()(*i0f949d8700cdd20dd5464492ee35ce01485a1a3ffdf50bce12c98acee6a8f39b.MembersRequestBuilder) {
    return i0f949d8700cdd20dd5464492ee35ce01485a1a3ffdf50bce12c98acee6a8f39b.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.channels.item.members.item collection
func (m *ChannelItemRequestBuilder) MembersById(id string)(*i833055126432577ccf7abb5c852b0232dd1f9933c3113ff1d863ba3fde085be2.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i833055126432577ccf7abb5c852b0232dd1f9933c3113ff1d863ba3fde085be2.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *ChannelItemRequestBuilder) Messages()(*iab02a926f6b68b941538c030223f00d8ee799b702aeb9af0f7dea46b8e014905.MessagesRequestBuilder) {
    return iab02a926f6b68b941538c030223f00d8ee799b702aeb9af0f7dea46b8e014905.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.channels.item.messages.item collection
func (m *ChannelItemRequestBuilder) MessagesById(id string)(*i654c2adcf04d8630b1cb24e632177c047ce85757bf63c3ea55ed853f8c2c06da.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return i654c2adcf04d8630b1cb24e632177c047ce85757bf63c3ea55ed853f8c2c06da.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in groups
func (m *ChannelItemRequestBuilder) Patch(options *ChannelItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ProvisionEmail the provisionEmail property
func (m *ChannelItemRequestBuilder) ProvisionEmail()(*i21e401c8131ed289d503e008a3d43d5445a6979bb2088584c87107e5f9a7efd9.ProvisionEmailRequestBuilder) {
    return i21e401c8131ed289d503e008a3d43d5445a6979bb2088584c87107e5f9a7efd9.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *ChannelItemRequestBuilder) RemoveEmail()(*i75075de7416f40b42eeae5c62e3a1830b76c46700fb92956f8a2994416edf871.RemoveEmailRequestBuilder) {
    return i75075de7416f40b42eeae5c62e3a1830b76c46700fb92956f8a2994416edf871.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams the sharedWithTeams property
func (m *ChannelItemRequestBuilder) SharedWithTeams()(*i242c35c5532ed25cf08fb018cf513312bd1fa96dd2b1e8b989dd0b1bb2ae55c6.SharedWithTeamsRequestBuilder) {
    return i242c35c5532ed25cf08fb018cf513312bd1fa96dd2b1e8b989dd0b1bb2ae55c6.NewSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.channels.item.sharedWithTeams.item collection
func (m *ChannelItemRequestBuilder) SharedWithTeamsById(id string)(*i1c349b81f3eac6900d6c57af5d7581e337b2ca590948d172e4adae3852bcc581.SharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return i1c349b81f3eac6900d6c57af5d7581e337b2ca590948d172e4adae3852bcc581.NewSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs the tabs property
func (m *ChannelItemRequestBuilder) Tabs()(*i005211e94469a36a987062d942e8e3e4b9f530d02f325b79d83acbc72a7296b6.TabsRequestBuilder) {
    return i005211e94469a36a987062d942e8e3e4b9f530d02f325b79d83acbc72a7296b6.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.channels.item.tabs.item collection
func (m *ChannelItemRequestBuilder) TabsById(id string)(*i50b60f6b70ccabd26064958871bb043fc875c38b3cdd16b1828e04c8a5a21880.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return i50b60f6b70ccabd26064958871bb043fc875c38b3cdd16b1828e04c8a5a21880.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
