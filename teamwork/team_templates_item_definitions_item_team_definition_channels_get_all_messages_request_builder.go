package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder provides operations to call the getAllMessages method.
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilderGetQueryParameters retrieve messages across all channels in a team, including text, audio, and video conversations. To learn more about using the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // The payment model for the API
    Model *string `uriparametername:"model"`
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
// TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilderGetQueryParameters
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilderInternal instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/channels/getAllMessages(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top,model*}", pathParameters),
    }
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder instantiates a new TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve messages across all channels in a team, including text, audio, and video conversations. To learn more about using the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
// Deprecated: This method is obsolete. Use GetAsGetAllMessagesGetResponse instead.
// returns a TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-getallmessages?view=graph-rest-beta
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilderGetRequestConfiguration)(TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesResponseable), nil
}
// GetAsGetAllMessagesGetResponse retrieve messages across all channels in a team, including text, audio, and video conversations. To learn more about using the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
// returns a TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-getallmessages?view=graph-rest-beta
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder) GetAsGetAllMessagesGetResponse(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilderGetRequestConfiguration)(TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesGetResponseable), nil
}
// ToGetRequestInformation retrieve messages across all channels in a team, including text, audio, and video conversations. To learn more about using the Microsoft Teams export APIs to export content, see Export content with the Microsoft Teams export APIs.
// returns a *RequestInformation when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder when successful
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder) WithUrl(rawUrl string)(*TeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder) {
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionChannelsGetAllMessagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
