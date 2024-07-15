package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder provides operations to call the applyConfig method.
type VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointProvisioningPoliciesApplyConfigRequestBuilderInternal instantiates a new VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder and sets the default values.
func NewVirtualEndpointProvisioningPoliciesApplyConfigRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder) {
    m := &VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/provisioningPolicies/applyConfig", pathParameters),
    }
    return m
}
// NewVirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder instantiates a new VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder and sets the default values.
func NewVirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointProvisioningPoliciesApplyConfigRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the provisioning policy configuration for a set of Cloud PC devices by their IDs. This method supports retry and allows you to apply the configuration to a subset of Cloud PCs initially to test.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcprovisioningpolicy-applyconfig?view=graph-rest-beta
func (m *VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder) Post(ctx context.Context, body VirtualEndpointProvisioningPoliciesApplyConfigPostRequestBodyable, requestConfiguration *VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation update the provisioning policy configuration for a set of Cloud PC devices by their IDs. This method supports retry and allows you to apply the configuration to a subset of Cloud PCs initially to test.
// returns a *RequestInformation when successful
func (m *VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointProvisioningPoliciesApplyConfigPostRequestBodyable, requestConfiguration *VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder when successful
func (m *VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder) {
    return NewVirtualEndpointProvisioningPoliciesApplyConfigRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
