package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder casts the previous resource to windowsWebApp.
type MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsWebApp
type MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphwindowswebappAssignmentsRequestBuilder when successful
func (m *MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder) Assignments()(*MobileappsItemGraphwindowswebappAssignmentsRequestBuilder) {
    return NewMobileappsItemGraphwindowswebappAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphwindowswebappCategoriesRequestBuilder when successful
func (m *MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder) Categories()(*MobileappsItemGraphwindowswebappCategoriesRequestBuilder) {
    return NewMobileappsItemGraphwindowswebappCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilderInternal instantiates a new MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder) {
    m := &MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.windowsWebApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder instantiates a new MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsWebApp
// returns a WindowsWebAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsWebAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsWebAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsWebAppable), nil
}
// Relationships provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphwindowswebappRelationshipsRequestBuilder when successful
func (m *MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder) Relationships()(*MobileappsItemGraphwindowswebappRelationshipsRequestBuilder) {
    return NewMobileappsItemGraphwindowswebappRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsWebApp
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder when successful
func (m *MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder) {
    return NewMobileappsItemGraphwindowswebappGraphWindowsWebAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
