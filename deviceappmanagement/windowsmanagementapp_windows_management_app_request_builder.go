package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsmanagementappWindowsManagementAppRequestBuilder provides operations to manage the windowsManagementApp property of the microsoft.graph.deviceAppManagement entity.
type WindowsmanagementappWindowsManagementAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsmanagementappWindowsManagementAppRequestBuilderGetQueryParameters windows management app.
type WindowsmanagementappWindowsManagementAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsmanagementappWindowsManagementAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsmanagementappWindowsManagementAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsmanagementappWindowsManagementAppRequestBuilderGetQueryParameters
}
// NewWindowsmanagementappWindowsManagementAppRequestBuilderInternal instantiates a new WindowsmanagementappWindowsManagementAppRequestBuilder and sets the default values.
func NewWindowsmanagementappWindowsManagementAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsmanagementappWindowsManagementAppRequestBuilder) {
    m := &WindowsmanagementappWindowsManagementAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/windowsManagementApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsmanagementappWindowsManagementAppRequestBuilder instantiates a new WindowsmanagementappWindowsManagementAppRequestBuilder and sets the default values.
func NewWindowsmanagementappWindowsManagementAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsmanagementappWindowsManagementAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsmanagementappWindowsManagementAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Get windows management app.
// returns a WindowsManagementAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsmanagementappWindowsManagementAppRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsmanagementappWindowsManagementAppRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagementAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsManagementAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagementAppable), nil
}
// ToGetRequestInformation windows management app.
// returns a *RequestInformation when successful
func (m *WindowsmanagementappWindowsManagementAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsmanagementappWindowsManagementAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsmanagementappWindowsManagementAppRequestBuilder when successful
func (m *WindowsmanagementappWindowsManagementAppRequestBuilder) WithUrl(rawUrl string)(*WindowsmanagementappWindowsManagementAppRequestBuilder) {
    return NewWindowsmanagementappWindowsManagementAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
