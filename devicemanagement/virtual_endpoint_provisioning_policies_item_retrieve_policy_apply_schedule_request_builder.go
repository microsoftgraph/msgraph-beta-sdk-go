package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder provides operations to call the retrievePolicyApplySchedule method.
type VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilderInternal instantiates a new VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder and sets the default values.
func NewVirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder) {
    m := &VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/provisioningPolicies/{cloudPcProvisioningPolicy%2Did}/retrievePolicyApplySchedule()", pathParameters),
    }
    return m
}
// NewVirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder instantiates a new VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder and sets the default values.
func NewVirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function retrievePolicyApplySchedule
// returns a CloudPcPolicyScheduledApplyActionDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcPolicyScheduledApplyActionDetailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcPolicyScheduledApplyActionDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcPolicyScheduledApplyActionDetailable), nil
}
// ToGetRequestInformation invoke function retrievePolicyApplySchedule
// returns a *RequestInformation when successful
func (m *VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder when successful
func (m *VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder) {
    return NewVirtualEndpointProvisioningPoliciesItemRetrievePolicyApplyScheduleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
