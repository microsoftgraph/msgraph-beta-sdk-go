package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemFindroomsFindRoomsRequestBuilder provides operations to call the findRooms method.
type ItemFindroomsFindRoomsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemFindroomsFindRoomsRequestBuilderGetQueryParameters invoke function findRooms
type ItemFindroomsFindRoomsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemFindroomsFindRoomsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemFindroomsFindRoomsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemFindroomsFindRoomsRequestBuilderGetQueryParameters
}
// NewItemFindroomsFindRoomsRequestBuilderInternal instantiates a new ItemFindroomsFindRoomsRequestBuilder and sets the default values.
func NewItemFindroomsFindRoomsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFindroomsFindRoomsRequestBuilder) {
    m := &ItemFindroomsFindRoomsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/findRooms(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemFindroomsFindRoomsRequestBuilder instantiates a new ItemFindroomsFindRoomsRequestBuilder and sets the default values.
func NewItemFindroomsFindRoomsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemFindroomsFindRoomsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemFindroomsFindRoomsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function findRooms
// Deprecated: This method is obsolete. Use GetAsFindRoomsGetResponse instead.
// returns a ItemFindroomsFindRoomsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemFindroomsFindRoomsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemFindroomsFindRoomsRequestBuilderGetRequestConfiguration)(ItemFindroomsFindRoomsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemFindroomsFindRoomsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemFindroomsFindRoomsResponseable), nil
}
// GetAsFindRoomsGetResponse invoke function findRooms
// returns a ItemFindroomsFindRoomsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemFindroomsFindRoomsRequestBuilder) GetAsFindRoomsGetResponse(ctx context.Context, requestConfiguration *ItemFindroomsFindRoomsRequestBuilderGetRequestConfiguration)(ItemFindroomsFindRoomsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemFindroomsFindRoomsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemFindroomsFindRoomsGetResponseable), nil
}
// ToGetRequestInformation invoke function findRooms
// returns a *RequestInformation when successful
func (m *ItemFindroomsFindRoomsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemFindroomsFindRoomsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemFindroomsFindRoomsRequestBuilder when successful
func (m *ItemFindroomsFindRoomsRequestBuilder) WithUrl(rawUrl string)(*ItemFindroomsFindRoomsRequestBuilder) {
    return NewItemFindroomsFindRoomsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
