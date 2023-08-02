package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
type TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderGetQueryParameters get the list of incoming channels (channels shared with a team).
type TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderGetQueryParameters struct {
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
// TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderGetQueryParameters
}
// ByChannelId provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder) ByChannelId(channelId string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if channelId != "" {
        urlTplParams["channel%2Did"] = channelId
    }
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderInternal instantiates a new IncomingChannelsRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/incomingChannels{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder instantiates a new IncomingChannelsRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder) Count()(*TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsCountRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the list of incoming channels (channels shared with a team).
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-list-incomingchannels?view=graph-rest-1.0
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChannelCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation get the list of incoming channels (channels shared with a team).
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionIncomingChannelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
