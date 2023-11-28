package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder provides operations to call the syncApps method.
type AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilderInternal instantiates a new SyncAppsRequestBuilder and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/syncApps", pathParameters),
    }
    return m
}
// NewAndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder instantiates a new SyncAppsRequestBuilder and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action syncApps
func (m *AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder) Post(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action syncApps
func (m *AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder) WithUrl(rawUrl string)(*AndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsSyncAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
