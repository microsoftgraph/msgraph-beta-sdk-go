package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
type ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderGetQueryParameters collection of ChromeOSOnboardingSettings settings associated with account.
type ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderGetQueryParameters
}
// ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByChromeOSOnboardingSettingsId provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
// returns a *ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder when successful
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) ByChromeOSOnboardingSettingsId(chromeOSOnboardingSettingsId string)(*ChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if chromeOSOnboardingSettingsId != "" {
        urlTplParams["chromeOSOnboardingSettings%2Did"] = chromeOSOnboardingSettingsId
    }
    return NewChromeosonboardingsettingsChromeOSOnboardingSettingsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Connect provides operations to call the connect method.
// returns a *ChromeosonboardingsettingsConnectRequestBuilder when successful
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) Connect()(*ChromeosonboardingsettingsConnectRequestBuilder) {
    return NewChromeosonboardingsettingsConnectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderInternal instantiates a new ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder and sets the default values.
func NewChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) {
    m := &ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/chromeOSOnboardingSettings{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder instantiates a new ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder and sets the default values.
func NewChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ChromeosonboardingsettingsCountRequestBuilder when successful
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) Count()(*ChromeosonboardingsettingsCountRequestBuilder) {
    return NewChromeosonboardingsettingsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Disconnect provides operations to call the disconnect method.
// returns a *ChromeosonboardingsettingsDisconnectRequestBuilder when successful
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) Disconnect()(*ChromeosonboardingsettingsDisconnectRequestBuilder) {
    return NewChromeosonboardingsettingsDisconnectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of ChromeOSOnboardingSettings settings associated with account.
// returns a ChromeOSOnboardingSettingsCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChromeOSOnboardingSettingsCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsCollectionResponseable), nil
}
// Post create new navigation property to chromeOSOnboardingSettings for deviceManagement
// returns a ChromeOSOnboardingSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable, requestConfiguration *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation collection of ChromeOSOnboardingSettings settings associated with account.
// returns a *RequestInformation when successful
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to chromeOSOnboardingSettings for deviceManagement
// returns a *RequestInformation when successful
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable, requestConfiguration *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder when successful
func (m *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) WithUrl(rawUrl string)(*ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) {
    return NewChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
