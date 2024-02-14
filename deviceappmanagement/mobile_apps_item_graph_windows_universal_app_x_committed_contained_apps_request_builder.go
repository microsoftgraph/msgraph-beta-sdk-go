package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder provides operations to manage the committedContainedApps property of the microsoft.graph.windowsUniversalAppX entity.
type MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderGetQueryParameters the collection of contained apps in the committed mobileAppContent of a windowsUniversalAppX app.
type MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderGetQueryParameters struct {
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
// MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderGetQueryParameters
}
// MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMobileContainedAppId provides operations to manage the committedContainedApps property of the microsoft.graph.windowsUniversalAppX entity.
// returns a *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsMobileContainedAppItemRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder) ByMobileContainedAppId(mobileContainedAppId string)(*MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsMobileContainedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mobileContainedAppId != "" {
        urlTplParams["mobileContainedApp%2Did"] = mobileContainedAppId
    }
    return NewMobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsMobileContainedAppItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderInternal instantiates a new MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder and sets the default values.
func NewMobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder) {
    m := &MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.windowsUniversalAppX/committedContainedApps{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder instantiates a new MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder and sets the default values.
func NewMobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsCountRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder) Count()(*MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsCountRequestBuilder) {
    return NewMobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of contained apps in the committed mobileAppContent of a windowsUniversalAppX app.
// returns a MobileContainedAppCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileContainedAppCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppCollectionResponseable), nil
}
// Post create new navigation property to committedContainedApps for deviceAppManagement
// returns a MobileContainedAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppable, requestConfiguration *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileContainedAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppable), nil
}
// ToGetRequestInformation the collection of contained apps in the committed mobileAppContent of a windowsUniversalAppX app.
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to committedContainedApps for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppable, requestConfiguration *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.windowsUniversalAppX/committedContainedApps", m.BaseRequestBuilder.PathParameters)
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
// returns a *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder) {
    return NewMobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
