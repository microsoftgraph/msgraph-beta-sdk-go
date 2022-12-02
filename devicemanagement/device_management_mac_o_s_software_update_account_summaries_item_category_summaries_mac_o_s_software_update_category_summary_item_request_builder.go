package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder provides operations to manage the categorySummaries property of the microsoft.graph.macOSSoftwareUpdateAccountSummary entity.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters summary of the updates by category.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetQueryParameters
}
// DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal instantiates a new MacOSSoftwareUpdateCategorySummaryItemRequestBuilder and sets the default values.
func NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    m := &DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary%2Did}/categorySummaries/{macOSSoftwareUpdateCategorySummary%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder instantiates a new MacOSSoftwareUpdateCategorySummaryItemRequestBuilder and sets the default values.
func NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property categorySummaries for deviceManagement
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation summary of the updates by category.
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property categorySummaries in deviceManagement
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property categorySummaries for deviceManagement
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get summary of the updates by category.
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable), nil
}
// Patch update the navigation property categorySummaries in deviceManagement
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateCategorySummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable), nil
}
// UpdateStateSummaries provides operations to manage the updateStateSummaries property of the microsoft.graph.macOSSoftwareUpdateCategorySummary entity.
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) UpdateStateSummaries()(*DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesRequestBuilder) {
    return NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateStateSummariesById provides operations to manage the updateStateSummaries property of the microsoft.graph.macOSSoftwareUpdateCategorySummary entity.
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesMacOSSoftwareUpdateCategorySummaryItemRequestBuilder) UpdateStateSummariesById(id string)(*DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["macOSSoftwareUpdateStateSummary%2Did"] = id
    }
    return NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesItemUpdateStateSummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
