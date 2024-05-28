package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder provides operations to call the targetApps method.
type TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilderInternal instantiates a new TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder and sets the default values.
func NewTargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder) {
    m := &TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration%2Did}/targetApps", pathParameters),
    }
    return m
}
// NewTargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder instantiates a new TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder and sets the default values.
func NewTargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action targetApps
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder) Post(ctx context.Context, body TargetedmanagedappconfigurationsItemTargetappsTargetAppsPostRequestBodyable, requestConfiguration *TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action targetApps
// returns a *RequestInformation when successful
func (m *TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TargetedmanagedappconfigurationsItemTargetappsTargetAppsPostRequestBodyable, requestConfiguration *TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder when successful
func (m *TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder) WithUrl(rawUrl string)(*TargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder) {
    return NewTargetedmanagedappconfigurationsItemTargetappsTargetAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
