package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder provides operations to call the validate method.
type IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilderInternal instantiates a new MicrosoftGraphIndustryDataValidateRequestBuilder and sets the default values.
func NewIndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder) {
    m := &IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/dataConnectors/{industryDataConnector%2Did}/microsoft.graph.industryData.validate", pathParameters),
    }
    return m
}
// NewIndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder instantiates a new MicrosoftGraphIndustryDataValidateRequestBuilder and sets the default values.
func NewIndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform validations applicable for the specific instance of the data connector. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-industrydataconnector-validate?view=graph-rest-1.0
func (m *IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder) Post(ctx context.Context, requestConfiguration *IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation perform validations applicable for the specific instance of the data connector. This API is available in the following national cloud deployments.
func (m *IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder) WithUrl(rawUrl string)(*IndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder) {
    return NewIndustryDataDataConnectorsItemMicrosoftGraphIndustryDataValidateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
