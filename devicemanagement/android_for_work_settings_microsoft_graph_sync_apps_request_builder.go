package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilder provides operations to call the syncApps method.
type AndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilderInternal instantiates a new MicrosoftGraphSyncAppsRequestBuilder and sets the default values.
func NewAndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilder) {
    m := &AndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidForWorkSettings/microsoft.graph.syncApps";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilder instantiates a new MicrosoftGraphSyncAppsRequestBuilder and sets the default values.
func NewAndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action syncApps
func (m *AndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilder) Post(ctx context.Context, requestConfiguration *AndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action syncApps
func (m *AndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AndroidForWorkSettingsMicrosoftGraphSyncAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
