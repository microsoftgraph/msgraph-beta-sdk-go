package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder provides operations to manage the categorySummaries property of the microsoft.graph.macOSSoftwareUpdateAccountSummary entity.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderGetQueryParameters summary of the updates by category.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderGetQueryParameters struct {
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
// DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderGetQueryParameters
}
// DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderInternal instantiates a new CategorySummariesRequestBuilder and sets the default values.
func NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder) {
    m := &DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary%2Did}/categorySummaries{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder instantiates a new CategorySummariesRequestBuilder and sets the default values.
func NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder) Count()(*DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesCountRequestBuilder) {
    return NewDeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation summary of the updates by category.
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to categorySummaries for deviceManagement
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get summary of the updates by category.
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateCategorySummaryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryCollectionResponseable), nil
}
// Post create new navigation property to categorySummaries for deviceManagement
func (m *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, requestConfiguration *DeviceManagementMacOSSoftwareUpdateAccountSummariesItemCategorySummariesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateCategorySummaryable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
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
