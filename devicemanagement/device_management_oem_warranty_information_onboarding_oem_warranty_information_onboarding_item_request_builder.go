package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder provides operations to manage the oemWarrantyInformationOnboarding property of the microsoft.graph.deviceManagement entity.
type DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetQueryParameters list of OEM Warranty Statuses
type DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetQueryParameters
}
// DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderInternal instantiates a new OemWarrantyInformationOnboardingItemRequestBuilder and sets the default values.
func NewDeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) {
    m := &DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/oemWarrantyInformationOnboarding/{oemWarrantyInformationOnboarding%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder instantiates a new OemWarrantyInformationOnboardingItemRequestBuilder and sets the default values.
func NewDeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property oemWarrantyInformationOnboarding for deviceManagement
func (m *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of OEM Warranty Statuses
func (m *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property oemWarrantyInformationOnboarding in deviceManagement
func (m *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable, requestConfiguration *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property oemWarrantyInformationOnboarding for deviceManagement
func (m *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Disable provides operations to call the disable method.
func (m *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) Disable()(*DeviceManagementOemWarrantyInformationOnboardingItemDisableRequestBuilder) {
    return NewDeviceManagementOemWarrantyInformationOnboardingItemDisableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Enable provides operations to call the enable method.
func (m *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) Enable()(*DeviceManagementOemWarrantyInformationOnboardingItemEnableRequestBuilder) {
    return NewDeviceManagementOemWarrantyInformationOnboardingItemEnableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get list of OEM Warranty Statuses
func (m *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOemWarrantyInformationOnboardingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable), nil
}
// Patch update the navigation property oemWarrantyInformationOnboarding in deviceManagement
func (m *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable, requestConfiguration *DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOemWarrantyInformationOnboardingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OemWarrantyInformationOnboardingable), nil
}
