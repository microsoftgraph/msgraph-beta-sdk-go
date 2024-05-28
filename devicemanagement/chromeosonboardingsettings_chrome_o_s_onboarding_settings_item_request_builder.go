package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
type ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderGetQueryParameters collection of ChromeOSOnboardingSettings settings associated with account.
type ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderGetQueryParameters
}
// ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderInternal instantiates a new ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder and sets the default values.
func NewChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder) {
    m := &ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/chromeOSOnboardingSettings/{chromeOSOnboardingSettings%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder instantiates a new ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder and sets the default values.
func NewChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property chromeOSOnboardingSettings for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get collection of ChromeOSOnboardingSettings settings associated with account.
// returns a ChromeOSOnboardingSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChromeOSOnboardingSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable), nil
}
// Patch update the navigation property chromeOSOnboardingSettings in deviceManagement
// returns a ChromeOSOnboardingSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable, requestConfiguration *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChromeOSOnboardingSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable), nil
}
// ToDeleteRequestInformation delete navigation property chromeOSOnboardingSettings for deviceManagement
// returns a *RequestInformation when successful
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of ChromeOSOnboardingSettings settings associated with account.
// returns a *RequestInformation when successful
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property chromeOSOnboardingSettings in deviceManagement
// returns a *RequestInformation when successful
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable, requestConfiguration *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder when successful
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder) WithUrl(rawUrl string)(*ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder) {
    return NewChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
