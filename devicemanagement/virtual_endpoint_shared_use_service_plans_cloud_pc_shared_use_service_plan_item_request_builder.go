package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder provides operations to manage the sharedUseServicePlans property of the microsoft.graph.virtualEndpoint entity.
type VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderGetQueryParameters read the properties and relationships of a cloudPcSharedUseServicePlan object.
type VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderGetQueryParameters
}
// VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderInternal instantiates a new VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder and sets the default values.
func NewVirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) {
    m := &VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/sharedUseServicePlans/{cloudPcSharedUseServicePlan%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder instantiates a new VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder and sets the default values.
func NewVirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sharedUseServicePlans for deviceManagement
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a cloudPcSharedUseServicePlan object.
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a CloudPcSharedUseServicePlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcshareduseserviceplan-get?view=graph-rest-1.0
func (m *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcSharedUseServicePlanable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcSharedUseServicePlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcSharedUseServicePlanable), nil
}
// Patch update the navigation property sharedUseServicePlans in deviceManagement
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a CloudPcSharedUseServicePlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcSharedUseServicePlanable, requestConfiguration *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcSharedUseServicePlanable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcSharedUseServicePlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcSharedUseServicePlanable), nil
}
// ToDeleteRequestInformation delete navigation property sharedUseServicePlans for deviceManagement
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a *RequestInformation when successful
func (m *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a cloudPcSharedUseServicePlan object.
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a *RequestInformation when successful
func (m *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sharedUseServicePlans in deviceManagement
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a *RequestInformation when successful
func (m *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcSharedUseServicePlanable, requestConfiguration *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder when successful
func (m *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) {
    return NewVirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
