package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder provides operations to manage the iosManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
type IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderGetQueryParameters iOS managed app policies.
type IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderGetQueryParameters struct {
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
// IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderGetQueryParameters
}
// IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByIosManagedAppProtectionId provides operations to manage the iosManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
// returns a *IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) ByIosManagedAppProtectionId(iosManagedAppProtectionId string)(*IosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if iosManagedAppProtectionId != "" {
        urlTplParams["iosManagedAppProtection%2Did"] = iosManagedAppProtectionId
    }
    return NewIosmanagedappprotectionsIosManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewIosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderInternal instantiates a new IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder and sets the default values.
func NewIosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) {
    m := &IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/iosManagedAppProtections{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder instantiates a new IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder and sets the default values.
func NewIosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *IosmanagedappprotectionsCountRequestBuilder when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) Count()(*IosmanagedappprotectionsCountRequestBuilder) {
    return NewIosmanagedappprotectionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get iOS managed app policies.
// returns a IosManagedAppProtectionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosManagedAppProtectionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosManagedAppProtectionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosManagedAppProtectionCollectionResponseable), nil
}
// HasPayloadLinks provides operations to call the hasPayloadLinks method.
// returns a *IosmanagedappprotectionsHaspayloadlinksHasPayloadLinksRequestBuilder when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) HasPayloadLinks()(*IosmanagedappprotectionsHaspayloadlinksHasPayloadLinksRequestBuilder) {
    return NewIosmanagedappprotectionsHaspayloadlinksHasPayloadLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to iosManagedAppProtections for deviceAppManagement
// returns a IosManagedAppProtectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosManagedAppProtectionable, requestConfiguration *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosManagedAppProtectionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosManagedAppProtectionable), nil
}
// ToGetRequestInformation iOS managed app policies.
// returns a *RequestInformation when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to iosManagedAppProtections for deviceAppManagement
// returns a *RequestInformation when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosManagedAppProtectionable, requestConfiguration *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder when successful
func (m *IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) WithUrl(rawUrl string)(*IosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder) {
    return NewIosmanagedappprotectionsIosManagedAppProtectionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
