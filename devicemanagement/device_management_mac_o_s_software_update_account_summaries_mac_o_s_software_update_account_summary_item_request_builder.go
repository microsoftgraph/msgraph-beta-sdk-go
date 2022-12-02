package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters the MacOS software update account summaries for this account.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters
}
// DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CategorySummaries provides operations to manage the categorySummaries property of the microsoft.graph.macOSSoftwareUpdateAccountSummary entity.
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CategorySummaries()(*DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder) {
    return NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategorySummariesById provides operations to manage the categorySummaries property of the microsoft.graph.macOSSoftwareUpdateAccountSummary entity.
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CategorySummariesById(id string)(*DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["macOSSoftwareUpdateCategorySummary%2Did"] = id
    }
    return NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal instantiates a new MacOSSoftwareUpdateAccountSummaryItemRequestBuilder and sets the default values.
func NewDeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    m := &DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder instantiates a new MacOSSoftwareUpdateAccountSummaryItemRequestBuilder and sets the default values.
func NewDeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property macOSSoftwareUpdateAccountSummaries for deviceManagement
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the MacOS software update account summaries for this account.
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property macOSSoftwareUpdateAccountSummaries in deviceManagement
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property macOSSoftwareUpdateAccountSummaries for deviceManagement
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the MacOS software update account summaries for this account.
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable), nil
}
// Patch update the navigation property macOSSoftwareUpdateAccountSummaries in deviceManagement
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable), nil
}
