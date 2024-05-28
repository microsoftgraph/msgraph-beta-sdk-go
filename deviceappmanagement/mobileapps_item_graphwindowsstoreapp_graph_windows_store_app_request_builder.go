package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder casts the previous resource to windowsStoreApp.
type MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsStoreApp
type MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphwindowsstoreappAssignmentsRequestBuilder when successful
func (m *MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder) Assignments()(*MobileappsItemGraphwindowsstoreappAssignmentsRequestBuilder) {
    return NewMobileappsItemGraphwindowsstoreappAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphwindowsstoreappCategoriesRequestBuilder when successful
func (m *MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder) Categories()(*MobileappsItemGraphwindowsstoreappCategoriesRequestBuilder) {
    return NewMobileappsItemGraphwindowsstoreappCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderInternal instantiates a new MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder) {
    m := &MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.windowsStoreApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder instantiates a new MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsStoreApp
// returns a WindowsStoreAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsStoreAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsStoreAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsStoreAppable), nil
}
// Relationships provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphwindowsstoreappRelationshipsRequestBuilder when successful
func (m *MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder) Relationships()(*MobileappsItemGraphwindowsstoreappRelationshipsRequestBuilder) {
    return NewMobileappsItemGraphwindowsstoreappRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsStoreApp
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder when successful
func (m *MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder) {
    return NewMobileappsItemGraphwindowsstoreappGraphWindowsStoreAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
