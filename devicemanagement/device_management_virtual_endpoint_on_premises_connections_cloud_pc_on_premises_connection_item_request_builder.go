package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
type DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
type DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters
}
// DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal instantiates a new CloudPcOnPremisesConnectionItemRequestBuilder and sets the default values.
func NewDeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) {
    m := &DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder instantiates a new CloudPcOnPremisesConnectionItemRequestBuilder and sets the default values.
func NewDeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onPremisesConnections for deviceManagement
func (m *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property onPremisesConnections in deviceManagement
func (m *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, requestConfiguration *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property onPremisesConnections for deviceManagement
func (m *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcOnPremisesConnectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable), nil
}
// Patch update the navigation property onPremisesConnections in deviceManagement
func (m *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, requestConfiguration *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcOnPremisesConnectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOnPremisesConnectionable), nil
}
// RunHealthChecks provides operations to call the runHealthChecks method.
func (m *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) RunHealthChecks()(*DeviceManagementVirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilder) {
    return NewDeviceManagementVirtualEndpointOnPremisesConnectionsItemRunHealthChecksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateAdDomainPassword provides operations to call the updateAdDomainPassword method.
func (m *DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) UpdateAdDomainPassword()(*DeviceManagementVirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilder) {
    return NewDeviceManagementVirtualEndpointOnPremisesConnectionsItemUpdateAdDomainPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
