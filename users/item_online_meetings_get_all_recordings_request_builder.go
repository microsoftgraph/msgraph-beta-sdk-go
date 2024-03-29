package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnlineMeetingsGetAllRecordingsRequestBuilder provides operations to call the getAllRecordings method.
type ItemOnlineMeetingsGetAllRecordingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlineMeetingsGetAllRecordingsRequestBuilderGetQueryParameters invoke function getAllRecordings
type ItemOnlineMeetingsGetAllRecordingsRequestBuilderGetQueryParameters struct {
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
// ItemOnlineMeetingsGetAllRecordingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsGetAllRecordingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnlineMeetingsGetAllRecordingsRequestBuilderGetQueryParameters
}
// NewItemOnlineMeetingsGetAllRecordingsRequestBuilderInternal instantiates a new ItemOnlineMeetingsGetAllRecordingsRequestBuilder and sets the default values.
func NewItemOnlineMeetingsGetAllRecordingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsGetAllRecordingsRequestBuilder) {
    m := &ItemOnlineMeetingsGetAllRecordingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/getAllRecordings(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemOnlineMeetingsGetAllRecordingsRequestBuilder instantiates a new ItemOnlineMeetingsGetAllRecordingsRequestBuilder and sets the default values.
func NewItemOnlineMeetingsGetAllRecordingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsGetAllRecordingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlineMeetingsGetAllRecordingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getAllRecordings
// Deprecated: This method is obsolete. Use GetAsGetAllRecordingsGetResponse instead.
// returns a ItemOnlineMeetingsGetAllRecordingsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlineMeetingsGetAllRecordingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlineMeetingsGetAllRecordingsRequestBuilderGetRequestConfiguration)(ItemOnlineMeetingsGetAllRecordingsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOnlineMeetingsGetAllRecordingsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOnlineMeetingsGetAllRecordingsResponseable), nil
}
// GetAsGetAllRecordingsGetResponse invoke function getAllRecordings
// returns a ItemOnlineMeetingsGetAllRecordingsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlineMeetingsGetAllRecordingsRequestBuilder) GetAsGetAllRecordingsGetResponse(ctx context.Context, requestConfiguration *ItemOnlineMeetingsGetAllRecordingsRequestBuilderGetRequestConfiguration)(ItemOnlineMeetingsGetAllRecordingsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOnlineMeetingsGetAllRecordingsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOnlineMeetingsGetAllRecordingsGetResponseable), nil
}
// ToGetRequestInformation invoke function getAllRecordings
// returns a *RequestInformation when successful
func (m *ItemOnlineMeetingsGetAllRecordingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlineMeetingsGetAllRecordingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnlineMeetingsGetAllRecordingsRequestBuilder when successful
func (m *ItemOnlineMeetingsGetAllRecordingsRequestBuilder) WithUrl(rawUrl string)(*ItemOnlineMeetingsGetAllRecordingsRequestBuilder) {
    return NewItemOnlineMeetingsGetAllRecordingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
