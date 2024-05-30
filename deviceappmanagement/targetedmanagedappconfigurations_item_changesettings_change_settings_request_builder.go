package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder provides operations to call the changeSettings method.
type TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilderInternal instantiates a new TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder and sets the default values.
func NewTargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder) {
    m := &TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration%2Did}/changeSettings", pathParameters),
    }
    return m
}
// NewTargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder instantiates a new TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder and sets the default values.
func NewTargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action changeSettings
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder) Post(ctx context.Context, body TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsPostRequestBodyable, requestConfiguration *TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action changeSettings
// returns a *RequestInformation when successful
func (m *TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsPostRequestBodyable, requestConfiguration *TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder when successful
func (m *TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder) WithUrl(rawUrl string)(*TargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder) {
    return NewTargetedmanagedappconfigurationsItemChangesettingsChangeSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
