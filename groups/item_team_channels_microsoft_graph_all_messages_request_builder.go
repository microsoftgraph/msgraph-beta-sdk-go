package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilder provides operations to call the allMessages method.
type ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilderGetQueryParameters invoke function allMessages
type ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilderGetQueryParameters
}
// NewItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilderInternal instantiates a new MicrosoftGraphAllMessagesRequestBuilder and sets the default values.
func NewItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilder) {
    m := &ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/channels/microsoft.graph.allMessages(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilder instantiates a new MicrosoftGraphAllMessagesRequestBuilder and sets the default values.
func NewItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function allMessages
func (m *ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilderGetRequestConfiguration)(ItemTeamChannelsMicrosoftGraphAllMessagesAllMessagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemTeamChannelsMicrosoftGraphAllMessagesAllMessagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamChannelsMicrosoftGraphAllMessagesAllMessagesResponseable), nil
}
// ToGetRequestInformation invoke function allMessages
func (m *ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamChannelsMicrosoftGraphAllMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
