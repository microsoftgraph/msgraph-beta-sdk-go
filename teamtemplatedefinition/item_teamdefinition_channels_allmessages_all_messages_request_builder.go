package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder provides operations to call the allMessages method.
type ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilderGetQueryParameters invoke function allMessages
type ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilderGetQueryParameters struct {
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
// ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilderGetQueryParameters
}
// NewItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilderInternal instantiates a new ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder and sets the default values.
func NewItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder) {
    m := &ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/channels/allMessages(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder instantiates a new ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder and sets the default values.
func NewItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function allMessages
// Deprecated: This method is obsolete. Use GetAsAllMessagesGetResponse instead.
// returns a ItemTeamdefinitionChannelsAllmessagesAllMessagesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilderGetRequestConfiguration)(ItemTeamdefinitionChannelsAllmessagesAllMessagesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamdefinitionChannelsAllmessagesAllMessagesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamdefinitionChannelsAllmessagesAllMessagesResponseable), nil
}
// GetAsAllMessagesGetResponse invoke function allMessages
// returns a ItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder) GetAsAllMessagesGetResponse(ctx context.Context, requestConfiguration *ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilderGetRequestConfiguration)(ItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamdefinitionChannelsAllmessagesAllMessagesGetResponseable), nil
}
// ToGetRequestInformation invoke function allMessages
// returns a *RequestInformation when successful
func (m *ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder when successful
func (m *ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder) WithUrl(rawUrl string)(*ItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder) {
    return NewItemTeamdefinitionChannelsAllmessagesAllMessagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
