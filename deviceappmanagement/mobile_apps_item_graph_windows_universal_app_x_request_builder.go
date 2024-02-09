package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphWindowsUniversalAppXRequestBuilder casts the previous resource to windowsUniversalAppX.
type MobileAppsItemGraphWindowsUniversalAppXRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphWindowsUniversalAppXRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsUniversalAppX
type MobileAppsItemGraphWindowsUniversalAppXRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsItemGraphWindowsUniversalAppXRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphWindowsUniversalAppXRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphWindowsUniversalAppXRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileAppsItemGraphWindowsUniversalAppXAssignmentsRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) Assignments()(*MobileAppsItemGraphWindowsUniversalAppXAssignmentsRequestBuilder) {
    return NewMobileAppsItemGraphWindowsUniversalAppXAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileAppsItemGraphWindowsUniversalAppXCategoriesRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) Categories()(*MobileAppsItemGraphWindowsUniversalAppXCategoriesRequestBuilder) {
    return NewMobileAppsItemGraphWindowsUniversalAppXCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CommittedContainedApps provides operations to manage the committedContainedApps property of the microsoft.graph.windowsUniversalAppX entity.
// returns a *MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) CommittedContainedApps()(*MobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilder) {
    return NewMobileAppsItemGraphWindowsUniversalAppXCommittedContainedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsItemGraphWindowsUniversalAppXRequestBuilderInternal instantiates a new MobileAppsItemGraphWindowsUniversalAppXRequestBuilder and sets the default values.
func NewMobileAppsItemGraphWindowsUniversalAppXRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) {
    m := &MobileAppsItemGraphWindowsUniversalAppXRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.windowsUniversalAppX{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphWindowsUniversalAppXRequestBuilder instantiates a new MobileAppsItemGraphWindowsUniversalAppXRequestBuilder and sets the default values.
func NewMobileAppsItemGraphWindowsUniversalAppXRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphWindowsUniversalAppXRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentVersions provides operations to manage the contentVersions property of the microsoft.graph.mobileLobApp entity.
// returns a *MobileAppsItemGraphWindowsUniversalAppXContentVersionsRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) ContentVersions()(*MobileAppsItemGraphWindowsUniversalAppXContentVersionsRequestBuilder) {
    return NewMobileAppsItemGraphWindowsUniversalAppXContentVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsUniversalAppX
// returns a WindowsUniversalAppXable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphWindowsUniversalAppXRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUniversalAppXable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsUniversalAppXFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsUniversalAppXable), nil
}
// Relationships provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
// returns a *MobileAppsItemGraphWindowsUniversalAppXRelationshipsRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) Relationships()(*MobileAppsItemGraphWindowsUniversalAppXRelationshipsRequestBuilder) {
    return NewMobileAppsItemGraphWindowsUniversalAppXRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.windowsUniversalAppX
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphWindowsUniversalAppXRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileAppsItemGraphWindowsUniversalAppXRequestBuilder when successful
func (m *MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphWindowsUniversalAppXRequestBuilder) {
    return NewMobileAppsItemGraphWindowsUniversalAppXRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
