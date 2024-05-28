package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetQueryParameters a collection of team members who have access to the shared channel.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetQueryParameters
}
// ByConversationMemberId provides operations to manage the allowedMembers property of the microsoft.graph.sharedWithChannelTeamInfo entity.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) ByConversationMemberId(conversationMemberId string)(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if conversationMemberId != "" {
        urlTplParams["conversationMember%2Did"] = conversationMemberId
    }
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderInternal instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) {
    m := &TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/primaryChannel/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder instantiates a new TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder and sets the default values.
func NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) Count()(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of team members who have access to the shared channel.
// returns a ConversationMemberCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConversationMemberCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversationMemberCollectionResponseable), nil
}
// ToGetRequestInformation a collection of team members who have access to the shared channel.
// returns a *RequestInformation when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder when successful
func (m *TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) WithUrl(rawUrl string)(*TeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder) {
    return NewTeamtemplatesItemDefinitionsItemTeamdefinitionPrimarychannelSharedwithteamsItemAllowedmembersAllowedMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
