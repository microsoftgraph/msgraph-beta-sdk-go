package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppCatalogPackagesRequestBuilder provides operations to manage the mobileAppCatalogPackages property of the microsoft.graph.deviceAppManagement entity.
type MobileAppCatalogPackagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppCatalogPackagesRequestBuilderGetQueryParameters mobileAppCatalogPackage entities.
type MobileAppCatalogPackagesRequestBuilderGetQueryParameters struct {
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
// MobileAppCatalogPackagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppCatalogPackagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppCatalogPackagesRequestBuilderGetQueryParameters
}
// MobileAppCatalogPackagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppCatalogPackagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMobileAppCatalogPackageId provides operations to manage the mobileAppCatalogPackages property of the microsoft.graph.deviceAppManagement entity.
// returns a *MobileAppCatalogPackagesMobileAppCatalogPackageItemRequestBuilder when successful
func (m *MobileAppCatalogPackagesRequestBuilder) ByMobileAppCatalogPackageId(mobileAppCatalogPackageId string)(*MobileAppCatalogPackagesMobileAppCatalogPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mobileAppCatalogPackageId != "" {
        urlTplParams["mobileAppCatalogPackage%2Did"] = mobileAppCatalogPackageId
    }
    return NewMobileAppCatalogPackagesMobileAppCatalogPackageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppCatalogPackagesRequestBuilderInternal instantiates a new MobileAppCatalogPackagesRequestBuilder and sets the default values.
func NewMobileAppCatalogPackagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppCatalogPackagesRequestBuilder) {
    m := &MobileAppCatalogPackagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileAppCatalogPackages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileAppCatalogPackagesRequestBuilder instantiates a new MobileAppCatalogPackagesRequestBuilder and sets the default values.
func NewMobileAppCatalogPackagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppCatalogPackagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppCatalogPackagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileAppCatalogPackagesCountRequestBuilder when successful
func (m *MobileAppCatalogPackagesRequestBuilder) Count()(*MobileAppCatalogPackagesCountRequestBuilder) {
    return NewMobileAppCatalogPackagesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get mobileAppCatalogPackage entities.
// returns a MobileAppCatalogPackageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppCatalogPackagesRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppCatalogPackagesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCatalogPackageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppCatalogPackageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCatalogPackageCollectionResponseable), nil
}
// Post create new navigation property to mobileAppCatalogPackages for deviceAppManagement
// returns a MobileAppCatalogPackageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppCatalogPackagesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCatalogPackageable, requestConfiguration *MobileAppCatalogPackagesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCatalogPackageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppCatalogPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCatalogPackageable), nil
}
// ToGetRequestInformation mobileAppCatalogPackage entities.
// returns a *RequestInformation when successful
func (m *MobileAppCatalogPackagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppCatalogPackagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to mobileAppCatalogPackages for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppCatalogPackagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCatalogPackageable, requestConfiguration *MobileAppCatalogPackagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/deviceAppManagement/mobileAppCatalogPackages", m.BaseRequestBuilder.PathParameters)
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
// returns a *MobileAppCatalogPackagesRequestBuilder when successful
func (m *MobileAppCatalogPackagesRequestBuilder) WithUrl(rawUrl string)(*MobileAppCatalogPackagesRequestBuilder) {
    return NewMobileAppCatalogPackagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
