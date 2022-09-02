package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0af18b3722f8a3fb442c5ba71f6304b653abe260406248d912b7093b3070c2e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname"
    i27cdacac1faf7ccf98b5489b1c37c00f34d852b59fa321eaa64e787d8bbc8759 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/provisionemail"
    i3d5b954cc594ac52c8068026f95d895591feb0e0ac93da7263456e59f30a1ba7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/removeemail"
    i430959b95f2aa3abfa37dcb98a2c3decf4c3300997e66066bff2b934ec5cfb62 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/members"
    i5a0d561c3987ff1c11f81d2b99423a6fd8c44d989307b9b4dc23b967da1baf21 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/completemigration"
    i7c49d37bc8a276bcb1738d203f4f83dc88c80e92437943471040b89124508d54 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/sharedwithteams"
    ic4dda6b0817ac926212dc2ea24b10997c07cb5f9a064f06af79ce8adff1ba1ca "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/messages"
    idbf8affae0c8851f2cf34764dc23d6df4539e08b1cbe1406ed8d79fcbd5251cb "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/tabs"
    ieef334d66ca01e1a9abeee81f9c42528fb192bb0f4e3b566e1006c4d792ef0ee "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/filesfolder"
    i164ae4c97bb1279df4e3326b6b4a1e622a1715574d897ae1f6c29c842d38a7d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/members/item"
    i64e6d1a204c5a8adce619b9d12d97fc76d415e5d88e0d81f4d6a20c578db5c9f "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/tabs/item"
    ib7fd317e805e6bcfebef0c2c48b8a9fdede7325d640446179b33e3d5ec06867d "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/sharedwithteams/item"
    ie02be09053704f8658f76e06dbaf2919bfbd29b3fc81181747b236ec49b5f0f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/teamwork/deletedteams/item/channels/item/messages/item"
)

// ChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.deletedTeam entity.
type ChannelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChannelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChannelItemRequestBuilderGetQueryParameters the channels those are either shared with this deleted team or created in this deleted team.
type ChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChannelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChannelItemRequestBuilderGetQueryParameters
}
// ChannelItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChannelItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration the completeMigration property
func (m *ChannelItemRequestBuilder) CompleteMigration()(*i5a0d561c3987ff1c11f81d2b99423a6fd8c44d989307b9b4dc23b967da1baf21.CompleteMigrationRequestBuilder) {
    return i5a0d561c3987ff1c11f81d2b99423a6fd8c44d989307b9b4dc23b967da1baf21.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelItemRequestBuilder) {
    m := &ChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property channels for teamwork
func (m *ChannelItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property channels for teamwork
func (m *ChannelItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ChannelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the channels those are either shared with this deleted team or created in this deleted team.
func (m *ChannelItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the channels those are either shared with this deleted team or created in this deleted team.
func (m *ChannelItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ChannelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property channels in teamwork
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property channels in teamwork
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ChannelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property channels for teamwork
func (m *ChannelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ChannelItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName provides operations to call the doesUserHaveAccess method.
func (m *ChannelItemRequestBuilder) DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName()(*i0af18b3722f8a3fb442c5ba71f6304b653abe260406248d912b7093b3070c2e7.DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    return i0af18b3722f8a3fb442c5ba71f6304b653abe260406248d912b7093b3070c2e7.NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder the filesFolder property
func (m *ChannelItemRequestBuilder) FilesFolder()(*ieef334d66ca01e1a9abeee81f9c42528fb192bb0f4e3b566e1006c4d792ef0ee.FilesFolderRequestBuilder) {
    return ieef334d66ca01e1a9abeee81f9c42528fb192bb0f4e3b566e1006c4d792ef0ee.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the channels those are either shared with this deleted team or created in this deleted team.
func (m *ChannelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ChannelItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable), nil
}
// Members the members property
func (m *ChannelItemRequestBuilder) Members()(*i430959b95f2aa3abfa37dcb98a2c3decf4c3300997e66066bff2b934ec5cfb62.MembersRequestBuilder) {
    return i430959b95f2aa3abfa37dcb98a2c3decf4c3300997e66066bff2b934ec5cfb62.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.deletedTeams.item.channels.item.members.item collection
func (m *ChannelItemRequestBuilder) MembersById(id string)(*i164ae4c97bb1279df4e3326b6b4a1e622a1715574d897ae1f6c29c842d38a7d7.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i164ae4c97bb1279df4e3326b6b4a1e622a1715574d897ae1f6c29c842d38a7d7.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *ChannelItemRequestBuilder) Messages()(*ic4dda6b0817ac926212dc2ea24b10997c07cb5f9a064f06af79ce8adff1ba1ca.MessagesRequestBuilder) {
    return ic4dda6b0817ac926212dc2ea24b10997c07cb5f9a064f06af79ce8adff1ba1ca.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.deletedTeams.item.channels.item.messages.item collection
func (m *ChannelItemRequestBuilder) MessagesById(id string)(*ie02be09053704f8658f76e06dbaf2919bfbd29b3fc81181747b236ec49b5f0f5.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return ie02be09053704f8658f76e06dbaf2919bfbd29b3fc81181747b236ec49b5f0f5.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in teamwork
func (m *ChannelItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *ChannelItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ProvisionEmail the provisionEmail property
func (m *ChannelItemRequestBuilder) ProvisionEmail()(*i27cdacac1faf7ccf98b5489b1c37c00f34d852b59fa321eaa64e787d8bbc8759.ProvisionEmailRequestBuilder) {
    return i27cdacac1faf7ccf98b5489b1c37c00f34d852b59fa321eaa64e787d8bbc8759.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *ChannelItemRequestBuilder) RemoveEmail()(*i3d5b954cc594ac52c8068026f95d895591feb0e0ac93da7263456e59f30a1ba7.RemoveEmailRequestBuilder) {
    return i3d5b954cc594ac52c8068026f95d895591feb0e0ac93da7263456e59f30a1ba7.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams the sharedWithTeams property
func (m *ChannelItemRequestBuilder) SharedWithTeams()(*i7c49d37bc8a276bcb1738d203f4f83dc88c80e92437943471040b89124508d54.SharedWithTeamsRequestBuilder) {
    return i7c49d37bc8a276bcb1738d203f4f83dc88c80e92437943471040b89124508d54.NewSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.deletedTeams.item.channels.item.sharedWithTeams.item collection
func (m *ChannelItemRequestBuilder) SharedWithTeamsById(id string)(*ib7fd317e805e6bcfebef0c2c48b8a9fdede7325d640446179b33e3d5ec06867d.SharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return ib7fd317e805e6bcfebef0c2c48b8a9fdede7325d640446179b33e3d5ec06867d.NewSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs the tabs property
func (m *ChannelItemRequestBuilder) Tabs()(*idbf8affae0c8851f2cf34764dc23d6df4539e08b1cbe1406ed8d79fcbd5251cb.TabsRequestBuilder) {
    return idbf8affae0c8851f2cf34764dc23d6df4539e08b1cbe1406ed8d79fcbd5251cb.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teamwork.deletedTeams.item.channels.item.tabs.item collection
func (m *ChannelItemRequestBuilder) TabsById(id string)(*i64e6d1a204c5a8adce619b9d12d97fc76d415e5d88e0d81f4d6a20c578db5c9f.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return i64e6d1a204c5a8adce619b9d12d97fc76d415e5d88e0d81f4d6a20c578db5c9f.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
