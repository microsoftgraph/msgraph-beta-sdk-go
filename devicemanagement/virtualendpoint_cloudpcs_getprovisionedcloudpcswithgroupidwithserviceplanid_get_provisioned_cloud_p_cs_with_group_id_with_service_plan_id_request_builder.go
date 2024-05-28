package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder provides operations to call the getProvisionedCloudPCs method.
type VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetQueryParameters get all provisioned Cloud PCs of a specific service plan for users under a Microsoft Entra user group.
type VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetQueryParameters struct {
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
// VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetQueryParameters
}
// NewVirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderInternal instantiates a new VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, groupId *string, servicePlanId *string)(*VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) {
    m := &VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/getProvisionedCloudPCs(groupId='{groupId}',servicePlanId='{servicePlanId}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if groupId != nil {
        m.BaseRequestBuilder.PathParameters["groupId"] = *groupId
    }
    if servicePlanId != nil {
        m.BaseRequestBuilder.PathParameters["servicePlanId"] = *servicePlanId
    }
    return m
}
// NewVirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder instantiates a new VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get get all provisioned Cloud PCs of a specific service plan for users under a Microsoft Entra user group.
// Deprecated: This method is obsolete. Use GetAsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponse instead.
// returns a VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getprovisionedcloudpcs?view=graph-rest-beta
func (m *VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration)(VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdResponseable), nil
}
// GetAsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponse get all provisioned Cloud PCs of a specific service plan for users under a Microsoft Entra user group.
// returns a VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-getprovisionedcloudpcs?view=graph-rest-beta
func (m *VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) GetAsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponse(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration)(VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdGetResponseable), nil
}
// ToGetRequestInformation get all provisioned Cloud PCs of a specific service plan for users under a Microsoft Entra user group.
// returns a *RequestInformation when successful
func (m *VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder when successful
func (m *VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) {
    return NewVirtualendpointCloudpcsGetprovisionedcloudpcswithgroupidwithserviceplanidGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
