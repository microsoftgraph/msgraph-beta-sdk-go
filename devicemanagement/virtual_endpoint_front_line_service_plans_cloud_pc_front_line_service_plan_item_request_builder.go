package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder provides operations to manage the frontLineServicePlans property of the microsoft.graph.virtualEndpoint entity.
type VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetQueryParameters front-line service plans for a Cloud PC.
type VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetQueryParameters
}
// VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderInternal instantiates a new VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder and sets the default values.
func NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) {
    m := &VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/frontLineServicePlans/{cloudPcFrontLineServicePlan%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder instantiates a new VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder and sets the default values.
func NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property frontLineServicePlans for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get front-line service plans for a Cloud PC.
// returns a CloudPcFrontLineServicePlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcFrontLineServicePlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable), nil
}
// Patch update the navigation property frontLineServicePlans in deviceManagement
// returns a CloudPcFrontLineServicePlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcFrontLineServicePlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable), nil
}
// ToDeleteRequestInformation delete navigation property frontLineServicePlans for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation front-line service plans for a Cloud PC.
// returns a *RequestInformation when successful
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property frontLineServicePlans in deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcFrontLineServicePlanable, requestConfiguration *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder when successful
func (m *VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder) {
    return NewVirtualEndpointFrontLineServicePlansCloudPcFrontLineServicePlanItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
