package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder provides operations to call the addApps method.
type AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilderInternal instantiates a new AddAppsRequestBuilder and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/addApps", pathParameters),
    }
    return m
}
// NewAndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder instantiates a new AddAppsRequestBuilder and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action addApps
func (m *AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder) Post(ctx context.Context, body AndroidManagedStoreAccountEnterpriseSettingsAddAppsPostRequestBodyable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action addApps
func (m *AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AndroidManagedStoreAccountEnterpriseSettingsAddAppsPostRequestBodyable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder) WithUrl(rawUrl string)(*AndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsAddAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
