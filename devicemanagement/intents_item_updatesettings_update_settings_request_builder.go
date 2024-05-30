package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IntentsItemUpdatesettingsUpdateSettingsRequestBuilder provides operations to call the updateSettings method.
type IntentsItemUpdatesettingsUpdateSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IntentsItemUpdatesettingsUpdateSettingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemUpdatesettingsUpdateSettingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIntentsItemUpdatesettingsUpdateSettingsRequestBuilderInternal instantiates a new IntentsItemUpdatesettingsUpdateSettingsRequestBuilder and sets the default values.
func NewIntentsItemUpdatesettingsUpdateSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemUpdatesettingsUpdateSettingsRequestBuilder) {
    m := &IntentsItemUpdatesettingsUpdateSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}/updateSettings", pathParameters),
    }
    return m
}
// NewIntentsItemUpdatesettingsUpdateSettingsRequestBuilder instantiates a new IntentsItemUpdatesettingsUpdateSettingsRequestBuilder and sets the default values.
func NewIntentsItemUpdatesettingsUpdateSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemUpdatesettingsUpdateSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntentsItemUpdatesettingsUpdateSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateSettings
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemUpdatesettingsUpdateSettingsRequestBuilder) Post(ctx context.Context, body IntentsItemUpdatesettingsUpdateSettingsPostRequestBodyable, requestConfiguration *IntentsItemUpdatesettingsUpdateSettingsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action updateSettings
// returns a *RequestInformation when successful
func (m *IntentsItemUpdatesettingsUpdateSettingsRequestBuilder) ToPostRequestInformation(ctx context.Context, body IntentsItemUpdatesettingsUpdateSettingsPostRequestBodyable, requestConfiguration *IntentsItemUpdatesettingsUpdateSettingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IntentsItemUpdatesettingsUpdateSettingsRequestBuilder when successful
func (m *IntentsItemUpdatesettingsUpdateSettingsRequestBuilder) WithUrl(rawUrl string)(*IntentsItemUpdatesettingsUpdateSettingsRequestBuilder) {
    return NewIntentsItemUpdatesettingsUpdateSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
