package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder provides operations to call the reset method.
type IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilderInternal instantiates a new IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder and sets the default values.
func NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder) {
    m := &IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/outboundProvisioningFlowSets/{outboundProvisioningFlowSet%2Did}/provisioningFlows/{provisioningFlow%2Did}/microsoft.graph.industryData.reset", pathParameters),
    }
    return m
}
// NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder instantiates a new IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder and sets the default values.
func NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reset a provisioningFlow. This action reprovisions all current data as if it were the initial run, and clears out the existing errors associated with the flow.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-provisioningflow-reset?view=graph-rest-beta
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder) Post(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation reset a provisioningFlow. This action reprovisions all current data as if it were the initial run, and clears out the existing errors associated with the flow.
// returns a *RequestInformation when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder when successful
func (m *IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder) WithUrl(rawUrl string)(*IndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder) {
    return NewIndustrydataOutboundprovisioningflowsetsItemProvisioningflowsItemMicrosoftgraphindustrydataresetMicrosoftGraphIndustryDataResetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
