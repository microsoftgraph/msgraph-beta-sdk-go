package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder provides operations to manage the windowsManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
type WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderGetQueryParameters windows managed app policies.
type WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderGetQueryParameters struct {
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
// WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderGetQueryParameters
}
// WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsManagedAppProtectionId provides operations to manage the windowsManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
// returns a *WindowsmanagedappprotectionsWindowsManagedAppProtectionItemRequestBuilder when successful
func (m *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder) ByWindowsManagedAppProtectionId(windowsManagedAppProtectionId string)(*WindowsmanagedappprotectionsWindowsManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsManagedAppProtectionId != "" {
        urlTplParams["windowsManagedAppProtection%2Did"] = windowsManagedAppProtectionId
    }
    return NewWindowsmanagedappprotectionsWindowsManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderInternal instantiates a new WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder and sets the default values.
func NewWindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder) {
    m := &WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsManagedAppProtections{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder instantiates a new WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder and sets the default values.
func NewWindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsmanagedappprotectionsCountRequestBuilder when successful
func (m *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder) Count()(*WindowsmanagedappprotectionsCountRequestBuilder) {
    return NewWindowsmanagedappprotectionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get windows managed app policies.
// returns a WindowsManagedAppProtectionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsManagedAppProtectionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionCollectionResponseable), nil
}
// Post create new navigation property to windowsManagedAppProtections for deviceAppManagement
// returns a WindowsManagedAppProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, requestConfiguration *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable), nil
}
// ToGetRequestInformation windows managed app policies.
// returns a *RequestInformation when successful
func (m *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to windowsManagedAppProtections for deviceAppManagement
// returns a *RequestInformation when successful
func (m *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, requestConfiguration *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder when successful
func (m *WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder) WithUrl(rawUrl string)(*WindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder) {
    return NewWindowsmanagedappprotectionsWindowsManagedAppProtectionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
