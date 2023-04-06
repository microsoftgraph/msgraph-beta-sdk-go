package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder provides operations to call the getProvisionedCloudPCs method.
type VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetQueryParameters invoke function getProvisionedCloudPCs
type VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetQueryParameters
}
// NewVirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderInternal instantiates a new GetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, groupId *string, servicePlanId *string)(*VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) {
    m := &VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/getProvisionedCloudPCs(groupId='{groupId}',servicePlanId='{servicePlanId}'){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    if groupId != nil {
        m.BaseRequestBuilder.PathParameters["groupId"] = *groupId
    }
    if servicePlanId != nil {
        m.BaseRequestBuilder.PathParameters["servicePlanId"] = *servicePlanId
    }
    return m
}
// NewVirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder instantiates a new GetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function getProvisionedCloudPCs
func (m *VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration)(VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdResponseable), nil
}
// ToGetRequestInformation invoke function getProvisionedCloudPCs
func (m *VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsGetProvisionedCloudPCsWithGroupIdWithServicePlanIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
