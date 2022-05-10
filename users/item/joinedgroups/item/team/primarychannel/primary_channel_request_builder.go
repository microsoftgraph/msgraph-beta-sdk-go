package primarychannel

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i397ac06d851bb412f3e763646f7677ea8be38c007ba637ac3eb96716caccbfb8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/sharedwithteams"
    i3ddd9fd419703e5f354e823fb09475a314add9efc6dce0ef592334b6f84edbb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/tabs"
    i629106f20aa076485bbd67ba9aee272c3e172866a557ff8003dd1a8c902844a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname"
    i71cccfc88b3e07ed99e5cbd41a97895f45bdb55b845032fde139b556d29ddcd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/filesfolder"
    i81d984bd0d172ff844eacc8982c8ab9bcafbaa17c4edca42ccca7ddf06534517 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/completemigration"
    i97c827d79b72e066c8fce453b58b08e8675ac391c8664490f64bd67cae080b57 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/provisionemail"
    ib9473c7ad4ca6126b108f14dafb729b09a4e8e71bdb6fa75913e76aad4cb5e4b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/members"
    ida72567330d3f23b04c4d2811d46100a31ef20b03eff5ce6c6ada733449f2807 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/removeemail"
    ief402bdaf8a83132e930912035955979e759a9ab82792718946f0007e3ae62f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/messages"
    i68542c3810b5bed8c180d0ee4cd3f5a7023f20044328baf64aa5c06ace5ec6af "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/members/item"
    i91995b0936addb540e028aaef5ddd32a9d5a22ff6a722b4edba2c75aa65c7ee6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/messages/item"
    icc69305ae45817774219c8d6cfa9e8d1dbcfffff223b34f877c694b8eb83adcc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/sharedwithteams/item"
    ida15749b553fd3a8f88b5f107a631df8c28915ca618d6b591389105b86a6167f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/team/primarychannel/tabs/item"
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
// PrimaryChannelRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrimaryChannelRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrimaryChannelRequestBuilderGetQueryParameters the general channel for the team.
type PrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrimaryChannelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrimaryChannelRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrimaryChannelRequestBuilderGetQueryParameters
}
// PrimaryChannelRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrimaryChannelRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration the completeMigration property
func (m *PrimaryChannelRequestBuilder) CompleteMigration()(*i81d984bd0d172ff844eacc8982c8ab9bcafbaa17c4edca42ccca7ddf06534517.CompleteMigrationRequestBuilder) {
    return i81d984bd0d172ff844eacc8982c8ab9bcafbaa17c4edca42ccca7ddf06534517.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrimaryChannelRequestBuilderInternal instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrimaryChannelRequestBuilder) {
    m := &PrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/team/primaryChannel{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property primaryChannel for users
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property primaryChannel for users
func (m *PrimaryChannelRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *PrimaryChannelRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the general channel for the team.
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the general channel for the team.
func (m *PrimaryChannelRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PrimaryChannelRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property primaryChannel in users
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property primaryChannel in users
func (m *PrimaryChannelRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *PrimaryChannelRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property primaryChannel for users
func (m *PrimaryChannelRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property primaryChannel for users
func (m *PrimaryChannelRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *PrimaryChannelRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName provides operations to call the doesUserHaveAccess method.
func (m *PrimaryChannelRequestBuilder) DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalName()(*i629106f20aa076485bbd67ba9aee272c3e172866a557ff8003dd1a8c902844a4.DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    return i629106f20aa076485bbd67ba9aee272c3e172866a557ff8003dd1a8c902844a4.NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesFolder the filesFolder property
func (m *PrimaryChannelRequestBuilder) FilesFolder()(*i71cccfc88b3e07ed99e5cbd41a97895f45bdb55b845032fde139b556d29ddcd9.FilesFolderRequestBuilder) {
    return i71cccfc88b3e07ed99e5cbd41a97895f45bdb55b845032fde139b556d29ddcd9.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the general channel for the team.
func (m *PrimaryChannelRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the general channel for the team.
func (m *PrimaryChannelRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *PrimaryChannelRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable), nil
}
// Members the members property
func (m *PrimaryChannelRequestBuilder) Members()(*ib9473c7ad4ca6126b108f14dafb729b09a4e8e71bdb6fa75913e76aad4cb5e4b.MembersRequestBuilder) {
    return ib9473c7ad4ca6126b108f14dafb729b09a4e8e71bdb6fa75913e76aad4cb5e4b.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.primaryChannel.members.item collection
func (m *PrimaryChannelRequestBuilder) MembersById(id string)(*i68542c3810b5bed8c180d0ee4cd3f5a7023f20044328baf64aa5c06ace5ec6af.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i68542c3810b5bed8c180d0ee4cd3f5a7023f20044328baf64aa5c06ace5ec6af.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *PrimaryChannelRequestBuilder) Messages()(*ief402bdaf8a83132e930912035955979e759a9ab82792718946f0007e3ae62f0.MessagesRequestBuilder) {
    return ief402bdaf8a83132e930912035955979e759a9ab82792718946f0007e3ae62f0.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.primaryChannel.messages.item collection
func (m *PrimaryChannelRequestBuilder) MessagesById(id string)(*i91995b0936addb540e028aaef5ddd32a9d5a22ff6a722b4edba2c75aa65c7ee6.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return i91995b0936addb540e028aaef5ddd32a9d5a22ff6a722b4edba2c75aa65c7ee6.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property primaryChannel in users
func (m *PrimaryChannelRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property primaryChannel in users
func (m *PrimaryChannelRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Channelable, requestConfiguration *PrimaryChannelRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ProvisionEmail the provisionEmail property
func (m *PrimaryChannelRequestBuilder) ProvisionEmail()(*i97c827d79b72e066c8fce453b58b08e8675ac391c8664490f64bd67cae080b57.ProvisionEmailRequestBuilder) {
    return i97c827d79b72e066c8fce453b58b08e8675ac391c8664490f64bd67cae080b57.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *PrimaryChannelRequestBuilder) RemoveEmail()(*ida72567330d3f23b04c4d2811d46100a31ef20b03eff5ce6c6ada733449f2807.RemoveEmailRequestBuilder) {
    return ida72567330d3f23b04c4d2811d46100a31ef20b03eff5ce6c6ada733449f2807.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeams the sharedWithTeams property
func (m *PrimaryChannelRequestBuilder) SharedWithTeams()(*i397ac06d851bb412f3e763646f7677ea8be38c007ba637ac3eb96716caccbfb8.SharedWithTeamsRequestBuilder) {
    return i397ac06d851bb412f3e763646f7677ea8be38c007ba637ac3eb96716caccbfb8.NewSharedWithTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedWithTeamsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.primaryChannel.sharedWithTeams.item collection
func (m *PrimaryChannelRequestBuilder) SharedWithTeamsById(id string)(*icc69305ae45817774219c8d6cfa9e8d1dbcfffff223b34f877c694b8eb83adcc.SharedWithChannelTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedWithChannelTeamInfo%2Did"] = id
    }
    return icc69305ae45817774219c8d6cfa9e8d1dbcfffff223b34f877c694b8eb83adcc.NewSharedWithChannelTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tabs the tabs property
func (m *PrimaryChannelRequestBuilder) Tabs()(*i3ddd9fd419703e5f354e823fb09475a314add9efc6dce0ef592334b6f84edbb2.TabsRequestBuilder) {
    return i3ddd9fd419703e5f354e823fb09475a314add9efc6dce0ef592334b6f84edbb2.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.team.primaryChannel.tabs.item collection
func (m *PrimaryChannelRequestBuilder) TabsById(id string)(*ida15749b553fd3a8f88b5f107a631df8c28915ca618d6b591389105b86a6167f.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return ida15749b553fd3a8f88b5f107a631df8c28915ca618d6b591389105b86a6167f.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
