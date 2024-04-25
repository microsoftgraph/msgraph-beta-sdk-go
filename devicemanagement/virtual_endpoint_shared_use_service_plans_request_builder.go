package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointSharedUseServicePlansRequestBuilder provides operations to manage the sharedUseServicePlans property of the microsoft.graph.virtualEndpoint entity.
type VirtualEndpointSharedUseServicePlansRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointSharedUseServicePlansRequestBuilderGetQueryParameters get sharedUseServicePlans from deviceManagement
type VirtualEndpointSharedUseServicePlansRequestBuilderGetQueryParameters struct {
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
// VirtualEndpointSharedUseServicePlansRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointSharedUseServicePlansRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointSharedUseServicePlansRequestBuilderGetQueryParameters
}
// VirtualEndpointSharedUseServicePlansRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointSharedUseServicePlansRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCloudPcSharedUseServicePlanId provides operations to manage the sharedUseServicePlans property of the microsoft.graph.virtualEndpoint entity.
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a *VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder when successful
func (m *VirtualEndpointSharedUseServicePlansRequestBuilder) ByCloudPcSharedUseServicePlanId(cloudPcSharedUseServicePlanId string)(*VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if cloudPcSharedUseServicePlanId != "" {
        urlTplParams["cloudPcSharedUseServicePlan%2Did"] = cloudPcSharedUseServicePlanId
    }
    return NewVirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEndpointSharedUseServicePlansRequestBuilderInternal instantiates a new VirtualEndpointSharedUseServicePlansRequestBuilder and sets the default values.
func NewVirtualEndpointSharedUseServicePlansRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointSharedUseServicePlansRequestBuilder) {
    m := &VirtualEndpointSharedUseServicePlansRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/sharedUseServicePlans{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualEndpointSharedUseServicePlansRequestBuilder instantiates a new VirtualEndpointSharedUseServicePlansRequestBuilder and sets the default values.
func NewVirtualEndpointSharedUseServicePlansRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointSharedUseServicePlansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointSharedUseServicePlansRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualEndpointSharedUseServicePlansCountRequestBuilder when successful
func (m *VirtualEndpointSharedUseServicePlansRequestBuilder) Count()(*VirtualEndpointSharedUseServicePlansCountRequestBuilder) {
    return NewVirtualEndpointSharedUseServicePlansCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get sharedUseServicePlans from deviceManagement
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a CloudPcSharedUseServicePlanCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointSharedUseServicePlansRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointSharedUseServicePlansRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcSharedUseServicePlanCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcSharedUseServicePlanCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcSharedUseServicePlanCollectionResponseable), nil
}
// Post create new navigation property to sharedUseServicePlans for deviceManagement
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a CloudPcSharedUseServicePlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointSharedUseServicePlansRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcSharedUseServicePlanable, requestConfiguration *VirtualEndpointSharedUseServicePlansRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcSharedUseServicePlanable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get sharedUseServicePlans from deviceManagement
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a *RequestInformation when successful
func (m *VirtualEndpointSharedUseServicePlansRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointSharedUseServicePlansRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to sharedUseServicePlans for deviceManagement
// Deprecated: The sharedUseServicePlans property is deprecated and will not be supported starting Oct 8, 2023. This property will not be included as part of the API response. as of 2023-03/sharedUseServicePlans
// returns a *RequestInformation when successful
func (m *VirtualEndpointSharedUseServicePlansRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcSharedUseServicePlanable, requestConfiguration *VirtualEndpointSharedUseServicePlansRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *VirtualEndpointSharedUseServicePlansRequestBuilder when successful
func (m *VirtualEndpointSharedUseServicePlansRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointSharedUseServicePlansRequestBuilder) {
    return NewVirtualEndpointSharedUseServicePlansRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
