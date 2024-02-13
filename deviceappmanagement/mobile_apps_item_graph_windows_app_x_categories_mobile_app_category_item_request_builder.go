package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
type MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilderGetQueryParameters the list of categories for this app.
type MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilderGetQueryParameters
}
// NewMobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilderInternal instantiates a new MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder) {
    m := &MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.windowsAppX/categories/{mobileAppCategory%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder instantiates a new MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the list of categories for this app.
// returns a MobileAppCategoryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCategoryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppCategoryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppCategoryable), nil
}
// ToGetRequestInformation the list of categories for this app.
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder) {
    return NewMobileAppsItemGraphWindowsAppXCategoriesMobileAppCategoryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
