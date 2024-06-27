package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphWindowsAppXRequestBuilder casts the previous resource to windowsAppX.
type MobileAppsItemGraphWindowsAppXRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphWindowsAppXRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsAppX
type MobileAppsItemGraphWindowsAppXRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsItemGraphWindowsAppXRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphWindowsAppXRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphWindowsAppXRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileAppsItemGraphWindowsAppXAssignmentsRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsAppXRequestBuilder) Assignments()(*MobileAppsItemGraphWindowsAppXAssignmentsRequestBuilder) {
    return NewMobileAppsItemGraphWindowsAppXAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileAppsItemGraphWindowsAppXCategoriesRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsAppXRequestBuilder) Categories()(*MobileAppsItemGraphWindowsAppXCategoriesRequestBuilder) {
    return NewMobileAppsItemGraphWindowsAppXCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsItemGraphWindowsAppXRequestBuilderInternal instantiates a new MobileAppsItemGraphWindowsAppXRequestBuilder and sets the default values.
func NewMobileAppsItemGraphWindowsAppXRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphWindowsAppXRequestBuilder) {
    m := &MobileAppsItemGraphWindowsAppXRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.windowsAppX{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphWindowsAppXRequestBuilder instantiates a new MobileAppsItemGraphWindowsAppXRequestBuilder and sets the default values.
func NewMobileAppsItemGraphWindowsAppXRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphWindowsAppXRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphWindowsAppXRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentVersions provides operations to manage the contentVersions property of the microsoft.graph.mobileLobApp entity.
// returns a *MobileAppsItemGraphWindowsAppXContentVersionsRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsAppXRequestBuilder) ContentVersions()(*MobileAppsItemGraphWindowsAppXContentVersionsRequestBuilder) {
    return NewMobileAppsItemGraphWindowsAppXContentVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsAppX
// returns a WindowsAppXable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphWindowsAppXRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphWindowsAppXRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAppXable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAppXFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAppXable), nil
}
// Relationships provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
// returns a *MobileAppsItemGraphWindowsAppXRelationshipsRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsAppXRequestBuilder) Relationships()(*MobileAppsItemGraphWindowsAppXRelationshipsRequestBuilder) {
    return NewMobileAppsItemGraphWindowsAppXRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsAppX
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphWindowsAppXRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphWindowsAppXRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileAppsItemGraphWindowsAppXRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsAppXRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphWindowsAppXRequestBuilder) {
    return NewMobileAppsItemGraphWindowsAppXRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
