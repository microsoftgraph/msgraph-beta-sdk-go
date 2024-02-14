package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder provides operations to manage the allChannels property of the microsoft.graph.team entity.
type TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderGetQueryParameters get the list of channels either in this team or shared with this team (incoming channels).
type TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderGetQueryParameters struct {
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
// TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderGetQueryParameters
}
// ByChannelId provides operations to manage the allChannels property of the microsoft.graph.team entity.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsChannelItemRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder) ByChannelId(channelId string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if channelId != "" {
        urlTplParams["channel%2Did"] = channelId
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderInternal instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/allChannels{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsCountRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder) Count()(*TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsCountRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the list of channels either in this team or shared with this team (incoming channels).
// returns a ChannelCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-list-allchannels?view=graph-rest-1.0
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChannelCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChannelCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChannelCollectionResponseable), nil
}
// ToGetRequestInformation get the list of channels either in this team or shared with this team (incoming channels).
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder) WithUrl(rawUrl string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionAllChannelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
