package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i38c36241584d283abed567d4941512ed9e66c68e569683934efaa70fd2818554 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/sharepoint"
    i3bebed6a38e396d320d4ac9b584cfacdff3f1460850f658e519a126116ae40b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows"
    i49dcf340f8344c7025f40bd5c0ae4cd9418dca301a765e4ac73ade0fec4a7fea "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/reportsettings"
    ic4ec23e71b979f81b15c68b4aa659bac1d7db6b385d7b2e876a5399aacb43a6e "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/serviceannouncement"
)

// AdminRequestBuilder provides operations to manage the admin singleton.
type AdminRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminRequestBuilderGetQueryParameters get admin
type AdminRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdminRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdminRequestBuilderGetQueryParameters
}
// AdminRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdminRequestBuilderInternal instantiates a new AdminRequestBuilder and sets the default values.
func NewAdminRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminRequestBuilder) {
    m := &AdminRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdminRequestBuilder instantiates a new AdminRequestBuilder and sets the default values.
func NewAdminRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get admin
func (m *AdminRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get admin
func (m *AdminRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AdminRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update admin
func (m *AdminRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Adminable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update admin
func (m *AdminRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Adminable, requestConfiguration *AdminRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get admin
func (m *AdminRequestBuilder) Get(ctx context.Context, requestConfiguration *AdminRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Adminable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdminFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Adminable), nil
}
// Patch update admin
func (m *AdminRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Adminable, requestConfiguration *AdminRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Adminable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdminFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Adminable), nil
}
// ReportSettings the reportSettings property
func (m *AdminRequestBuilder) ReportSettings()(*i49dcf340f8344c7025f40bd5c0ae4cd9418dca301a765e4ac73ade0fec4a7fea.ReportSettingsRequestBuilder) {
    return i49dcf340f8344c7025f40bd5c0ae4cd9418dca301a765e4ac73ade0fec4a7fea.NewReportSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServiceAnnouncement the serviceAnnouncement property
func (m *AdminRequestBuilder) ServiceAnnouncement()(*ic4ec23e71b979f81b15c68b4aa659bac1d7db6b385d7b2e876a5399aacb43a6e.ServiceAnnouncementRequestBuilder) {
    return ic4ec23e71b979f81b15c68b4aa659bac1d7db6b385d7b2e876a5399aacb43a6e.NewServiceAnnouncementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Sharepoint the sharepoint property
func (m *AdminRequestBuilder) Sharepoint()(*i38c36241584d283abed567d4941512ed9e66c68e569683934efaa70fd2818554.SharepointRequestBuilder) {
    return i38c36241584d283abed567d4941512ed9e66c68e569683934efaa70fd2818554.NewSharepointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Windows the windows property
func (m *AdminRequestBuilder) Windows()(*i3bebed6a38e396d320d4ac9b584cfacdff3f1460850f658e519a126116ae40b8.WindowsRequestBuilder) {
    return i3bebed6a38e396d320d4ac9b584cfacdff3f1460850f658e519a126116ae40b8.NewWindowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
