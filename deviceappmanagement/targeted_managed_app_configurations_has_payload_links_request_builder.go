package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder provides operations to call the hasPayloadLinks method.
type TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTargetedManagedAppConfigurationsHasPayloadLinksRequestBuilderInternal instantiates a new TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationsHasPayloadLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder) {
    m := &TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/hasPayloadLinks", pathParameters),
    }
    return m
}
// NewTargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder instantiates a new TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppConfigurationsHasPayloadLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasPayloadLinks
// Deprecated: This method is obsolete. Use PostAsHasPayloadLinksPostResponse instead.
// returns a TargetedManagedAppConfigurationsHasPayloadLinksResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder) Post(ctx context.Context, body TargetedManagedAppConfigurationsHasPayloadLinksPostRequestBodyable, requestConfiguration *TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilderPostRequestConfiguration)(TargetedManagedAppConfigurationsHasPayloadLinksResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTargetedManagedAppConfigurationsHasPayloadLinksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TargetedManagedAppConfigurationsHasPayloadLinksResponseable), nil
}
// PostAsHasPayloadLinksPostResponse invoke action hasPayloadLinks
// returns a TargetedManagedAppConfigurationsHasPayloadLinksPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder) PostAsHasPayloadLinksPostResponse(ctx context.Context, body TargetedManagedAppConfigurationsHasPayloadLinksPostRequestBodyable, requestConfiguration *TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilderPostRequestConfiguration)(TargetedManagedAppConfigurationsHasPayloadLinksPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTargetedManagedAppConfigurationsHasPayloadLinksPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TargetedManagedAppConfigurationsHasPayloadLinksPostResponseable), nil
}
// ToPostRequestInformation invoke action hasPayloadLinks
// returns a *RequestInformation when successful
func (m *TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder) ToPostRequestInformation(ctx context.Context, body TargetedManagedAppConfigurationsHasPayloadLinksPostRequestBodyable, requestConfiguration *TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder when successful
func (m *TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder) WithUrl(rawUrl string)(*TargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder) {
    return NewTargetedManagedAppConfigurationsHasPayloadLinksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
