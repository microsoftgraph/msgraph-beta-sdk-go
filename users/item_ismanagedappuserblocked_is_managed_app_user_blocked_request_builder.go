package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder provides operations to call the isManagedAppUserBlocked method.
type ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilderInternal instantiates a new ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder and sets the default values.
func NewItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder) {
    m := &ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/isManagedAppUserBlocked()", pathParameters),
    }
    return m
}
// NewItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder instantiates a new ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder and sets the default values.
func NewItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the blocked state of a managed app user.
// Deprecated: This method is obsolete. Use GetAsIsManagedAppUserBlockedGetResponse instead.
// returns a ItemIsmanagedappuserblockedIsManagedAppUserBlockedResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilderGetRequestConfiguration)(ItemIsmanagedappuserblockedIsManagedAppUserBlockedResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemIsmanagedappuserblockedIsManagedAppUserBlockedResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemIsmanagedappuserblockedIsManagedAppUserBlockedResponseable), nil
}
// GetAsIsManagedAppUserBlockedGetResponse gets the blocked state of a managed app user.
// returns a ItemIsmanagedappuserblockedIsManagedAppUserBlockedGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder) GetAsIsManagedAppUserBlockedGetResponse(ctx context.Context, requestConfiguration *ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilderGetRequestConfiguration)(ItemIsmanagedappuserblockedIsManagedAppUserBlockedGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemIsmanagedappuserblockedIsManagedAppUserBlockedGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemIsmanagedappuserblockedIsManagedAppUserBlockedGetResponseable), nil
}
// ToGetRequestInformation gets the blocked state of a managed app user.
// returns a *RequestInformation when successful
func (m *ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder when successful
func (m *ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder) WithUrl(rawUrl string)(*ItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder) {
    return NewItemIsmanagedappuserblockedIsManagedAppUserBlockedRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
