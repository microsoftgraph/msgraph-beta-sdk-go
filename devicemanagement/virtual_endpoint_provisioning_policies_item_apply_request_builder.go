package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder provides operations to call the apply method.
type VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointProvisioningPoliciesItemApplyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointProvisioningPoliciesItemApplyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointProvisioningPoliciesItemApplyRequestBuilderInternal instantiates a new VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder and sets the default values.
func NewVirtualEndpointProvisioningPoliciesItemApplyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder) {
    m := &VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/provisioningPolicies/{cloudPcProvisioningPolicy%2Did}/apply", pathParameters),
    }
    return m
}
// NewVirtualEndpointProvisioningPoliciesItemApplyRequestBuilder instantiates a new VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder and sets the default values.
func NewVirtualEndpointProvisioningPoliciesItemApplyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointProvisioningPoliciesItemApplyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post apply the current provisioning policy configuration to all Cloud PC devices under a specified policy. Currently, the region is the only policy setting that you can apply.
// Deprecated: The onPremisesConnectionId property is deprecated and will stop returning on July 30, 2023. as of 2023-03/onPremisesConnectionId
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcprovisioningpolicy-apply?view=graph-rest-beta
func (m *VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder) Post(ctx context.Context, body VirtualEndpointProvisioningPoliciesItemApplyPostRequestBodyable, requestConfiguration *VirtualEndpointProvisioningPoliciesItemApplyRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation apply the current provisioning policy configuration to all Cloud PC devices under a specified policy. Currently, the region is the only policy setting that you can apply.
// Deprecated: The onPremisesConnectionId property is deprecated and will stop returning on July 30, 2023. as of 2023-03/onPremisesConnectionId
// returns a *RequestInformation when successful
func (m *VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointProvisioningPoliciesItemApplyPostRequestBodyable, requestConfiguration *VirtualEndpointProvisioningPoliciesItemApplyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The onPremisesConnectionId property is deprecated and will stop returning on July 30, 2023. as of 2023-03/onPremisesConnectionId
// returns a *VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder when successful
func (m *VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointProvisioningPoliciesItemApplyRequestBuilder) {
    return NewVirtualEndpointProvisioningPoliciesItemApplyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
