package primarychannel

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0449b81134c61c669c1db7ab85cfc49f222cd55d433bc73bade8c1b2998c1c62 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/removeemail"
    i382f67431e8fad97449e69db32556bff1823eb0f75e897c8708039d9b4e0daaa "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/completemigration"
    i640eb9d5e7144c0052d2d0cef299ce1fa11d20a248e996caef4a723c672b60de "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname"
    i7edc7f5259d5db562c84cc16ebd1d6c953b609207c7e242543c9b7df1fc71235 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/members"
    i9093c9fedf01f6c908b2550374c62ef1fe3375446859db6fec5041c358d90336 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/filesfolder"
    i91609f8b20925576bd47ef331e1eea32f67eb9006d4d07418261aaaa95899dbd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/messages"
    ibded2e76ab803fb8c99bb7f58a5eede3e99873317d67026d66b065b87b55e86b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/sharedwithteams"
    idbfdc5f80a6d286aa4e4df5e8f287898af0e023ab3d74fd479733e2ef4def31d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/provisionemail"
    if6ef420b7c287eb7c1bfbc3a9ccefc5cf6e65d3858316d9cc4edd8f4055c06b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/tabs"
    i06994b42d6f2b79a353c090540e8d4f118668adf056428c7d7fcbe8b08523302 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/members/item"
    i2925bf60bfc540b5241b63815da7a81b7bfe4cd5d5bef0452e27f46c194c51d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/sharedwithteams/item"
    i3ef987913bc93d2b6ab62eee2c31a168a9bfcf76ad15cf70c0b046c1ca35911d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/tabs/item"
    i44a689dd15b418268e0970efd9d4eaafc230d92f7d2ac2fc421a57194102733a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/primarychannel/messages/item"
)

// PrimaryChannelRequestBuilder provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
type PrimaryChannelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrimaryChannelRequestBuilderDeleteOptions options for Delete
type PrimaryChannelRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// PrimaryChannelRequestBuilderGetOptions options for Get
type PrimaryChannelRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrimaryChannelRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// PrimaryChannelRequestBuilderGetQueryParameters the general channel for the team.
type PrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrimaryChannelRequestBuilderPatchOptions options for Patch
type PrimaryChannelRequestBuilderPatchOptions struct {
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
func (m *PrimaryChannelRequestBuilder) CompleteMigration()(*i382f67431e8fad97449e69db32556bff1823eb0f75e897c8708039d9b4e0daaa.CompleteMigrationRequestBuilder) {
    return i382f67431e8fad97449e69db32556bff1823eb0f75e897c8708039d9b4e0daaa.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrimaryChannelRequestBuilderInternal instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    m := &PrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/primaryChannel{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrimaryChannelRequestBuilder instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property primaryChannel for groups
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformation(options *PrimaryChannelRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the general channel for the team.
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformation(options *PrimaryChannelRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property primaryChannel in groups
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformation(options *PrimaryChannelRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property primaryChannel for groups
func (m *PrimaryChannelRequestBuilder) Delete(options *PrimaryChannelRequestBuilderDeleteOptions)(error) {
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
func (m *PrimaryChannelRequestBuilder) DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName()(*i640eb9d5e7144c0052d2d0cef299ce1fa11d20a248e996caef4a723c672b60de.DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    return i640eb9d5e7144c0052d2d0cef299ce1fa11d20a248e996caef4a723c672b60de.NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder the filesFolder property
func (m *PrimaryChannelRequestBuilder) FilesFolder()(*i9093c9fedf01f6c908b2550374c62ef1fe3375446859db6fec5041c358d90336.FilesFolderRequestBuilder) {
    return i9093c9fedf01f6c908b2550374c62ef1fe3375446859db6fec5041c358d90336.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the general channel for the team.
func (m *PrimaryChannelRequestBuilder) Get(options *PrimaryChannelRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
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
func (m *PrimaryChannelRequestBuilder) Members()(*i7edc7f5259d5db562c84cc16ebd1d6c953b609207c7e242543c9b7df1fc71235.MembersRequestBuilder) {
    return i7edc7f5259d5db562c84cc16ebd1d6c953b609207c7e242543c9b7df1fc71235.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.primaryChannel.members.item collection
func (m *PrimaryChannelRequestBuilder) MembersById(id string)(*i06994b42d6f2b79a353c090540e8d4f118668adf056428c7d7fcbe8b08523302.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i06994b42d6f2b79a353c090540e8d4f118668adf056428c7d7fcbe8b08523302.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *PrimaryChannelRequestBuilder) Messages()(*i91609f8b20925576bd47ef331e1eea32f67eb9006d4d07418261aaaa95899dbd.MessagesRequestBuilder) {
    return i91609f8b20925576bd47ef331e1eea32f67eb9006d4d07418261aaaa95899dbd.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.primaryChannel.messages.item collection
func (m *PrimaryChannelRequestBuilder) MessagesById(id string)(*i44a689dd15b418268e0970efd9d4eaafc230d92f7d2ac2fc421a57194102733a.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return i44a689dd15b418268e0970efd9d4eaafc230d92f7d2ac2fc421a57194102733a.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property primaryChannel in groups
func (m *PrimaryChannelRequestBuilder) Patch(options *PrimaryChannelRequestBuilderPatchOptions)(error) {
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
func (m *PrimaryChannelRequestBuilder) ProvisionEmail()(*idbfdc5f80a6d286aa4e4df5e8f287898af0e023ab3d74fd479733e2ef4def31d.ProvisionEmailRequestBuilder) {
    return idbfdc5f80a6d286aa4e4df5e8f287898af0e023ab3d74fd479733e2ef4def31d.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *PrimaryChannelRequestBuilder) RemoveEmail()(*i0449b81134c61c669c1db7ab85cfc49f222cd55d433bc73bade8c1b2998c1c62.RemoveEmailRequestBuilder) {
    return i0449b81134c61c669c1db7ab85cfc49f222cd55d433bc73bade8c1b2998c1c62.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams the sharedWithTeams property
func (m *PrimaryChannelRequestBuilder) SharedWithTeams()(*ibded2e76ab803fb8c99bb7f58a5eede3e99873317d67026d66b065b87b55e86b.SharedWithTeamsRequestBuilder) {
    return ibded2e76ab803fb8c99bb7f58a5eede3e99873317d67026d66b065b87b55e86b.NewSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.primaryChannel.sharedWithTeams.item collection
func (m *PrimaryChannelRequestBuilder) SharedWithTeamsById(id string)(*i2925bf60bfc540b5241b63815da7a81b7bfe4cd5d5bef0452e27f46c194c51d9.SharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return i2925bf60bfc540b5241b63815da7a81b7bfe4cd5d5bef0452e27f46c194c51d9.NewSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs the tabs property
func (m *PrimaryChannelRequestBuilder) Tabs()(*if6ef420b7c287eb7c1bfbc3a9ccefc5cf6e65d3858316d9cc4edd8f4055c06b8.TabsRequestBuilder) {
    return if6ef420b7c287eb7c1bfbc3a9ccefc5cf6e65d3858316d9cc4edd8f4055c06b8.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.team.primaryChannel.tabs.item collection
func (m *PrimaryChannelRequestBuilder) TabsById(id string)(*i3ef987913bc93d2b6ab62eee2c31a168a9bfcf76ad15cf70c0b046c1ca35911d.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return i3ef987913bc93d2b6ab62eee2c31a168a9bfcf76ad15cf70c0b046c1ca35911d.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
