package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ChromeOSOnboardingSettingsConnectRequestBuilder provides operations to call the connect method.
type ChromeOSOnboardingSettingsConnectRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChromeOSOnboardingSettingsConnectRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChromeOSOnboardingSettingsConnectRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChromeOSOnboardingSettingsConnectRequestBuilderInternal instantiates a new ChromeOSOnboardingSettingsConnectRequestBuilder and sets the default values.
func NewChromeOSOnboardingSettingsConnectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChromeOSOnboardingSettingsConnectRequestBuilder) {
    m := &ChromeOSOnboardingSettingsConnectRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/chromeOSOnboardingSettings/connect", pathParameters),
    }
    return m
}
// NewChromeOSOnboardingSettingsConnectRequestBuilder instantiates a new ChromeOSOnboardingSettingsConnectRequestBuilder and sets the default values.
func NewChromeOSOnboardingSettingsConnectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChromeOSOnboardingSettingsConnectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChromeOSOnboardingSettingsConnectRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action connect
// returns a *ChromeOSOnboardingStatus when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ChromeOSOnboardingSettingsConnectRequestBuilder) Post(ctx context.Context, body ChromeOSOnboardingSettingsConnectPostRequestBodyable, requestConfiguration *ChromeOSOnboardingSettingsConnectRequestBuilderPostRequestConfiguration)(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingStatus, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendEnum(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseChromeOSOnboardingStatus, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingStatus), nil
}
// ToPostRequestInformation invoke action connect
// returns a *RequestInformation when successful
func (m *ChromeOSOnboardingSettingsConnectRequestBuilder) ToPostRequestInformation(ctx context.Context, body ChromeOSOnboardingSettingsConnectPostRequestBodyable, requestConfiguration *ChromeOSOnboardingSettingsConnectRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ChromeOSOnboardingSettingsConnectRequestBuilder when successful
func (m *ChromeOSOnboardingSettingsConnectRequestBuilder) WithUrl(rawUrl string)(*ChromeOSOnboardingSettingsConnectRequestBuilder) {
    return NewChromeOSOnboardingSettingsConnectRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
