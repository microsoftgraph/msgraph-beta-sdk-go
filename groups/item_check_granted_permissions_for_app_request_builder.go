package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCheckGrantedPermissionsForAppRequestBuilder provides operations to call the checkGrantedPermissionsForApp method.
type ItemCheckGrantedPermissionsForAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCheckGrantedPermissionsForAppRequestBuilderInternal instantiates a new ItemCheckGrantedPermissionsForAppRequestBuilder and sets the default values.
func NewItemCheckGrantedPermissionsForAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCheckGrantedPermissionsForAppRequestBuilder) {
    m := &ItemCheckGrantedPermissionsForAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/checkGrantedPermissionsForApp", pathParameters),
    }
    return m
}
// NewItemCheckGrantedPermissionsForAppRequestBuilder instantiates a new ItemCheckGrantedPermissionsForAppRequestBuilder and sets the default values.
func NewItemCheckGrantedPermissionsForAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCheckGrantedPermissionsForAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCheckGrantedPermissionsForAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action checkGrantedPermissionsForApp
// Deprecated: This method is obsolete. Use PostAsCheckGrantedPermissionsForAppPostResponse instead.
// returns a ItemCheckGrantedPermissionsForAppResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCheckGrantedPermissionsForAppRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration)(ItemCheckGrantedPermissionsForAppResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCheckGrantedPermissionsForAppResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCheckGrantedPermissionsForAppResponseable), nil
}
// PostAsCheckGrantedPermissionsForAppPostResponse invoke action checkGrantedPermissionsForApp
// returns a ItemCheckGrantedPermissionsForAppPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCheckGrantedPermissionsForAppRequestBuilder) PostAsCheckGrantedPermissionsForAppPostResponse(ctx context.Context, requestConfiguration *ItemCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration)(ItemCheckGrantedPermissionsForAppPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCheckGrantedPermissionsForAppPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCheckGrantedPermissionsForAppPostResponseable), nil
}
// ToPostRequestInformation invoke action checkGrantedPermissionsForApp
// returns a *RequestInformation when successful
func (m *ItemCheckGrantedPermissionsForAppRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCheckGrantedPermissionsForAppRequestBuilder when successful
func (m *ItemCheckGrantedPermissionsForAppRequestBuilder) WithUrl(rawUrl string)(*ItemCheckGrantedPermissionsForAppRequestBuilder) {
    return NewItemCheckGrantedPermissionsForAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
