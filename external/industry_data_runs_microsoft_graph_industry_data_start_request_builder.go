package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder provides operations to call the start method.
type IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilderInternal instantiates a new IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder and sets the default values.
func NewIndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder) {
    m := &IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/runs/microsoft.graph.industryData.start", pathParameters),
    }
    return m
}
// NewIndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder instantiates a new IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder and sets the default values.
func NewIndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilderInternal(urlParams, requestAdapter)
}
// Post start a new industryDataRun. Industry data automates a run every 12 hours; however, users can use the start action to perform an on-demand run. The industry data service throttles the start of runs, allowing up to five successful runs every 12 hours. An on-demand run causes the next system-scheduled run to be skipped if it is set to run within the next 10 hours.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder) Post(ctx context.Context, requestConfiguration *IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation start a new industryDataRun. Industry data automates a run every 12 hours; however, users can use the start action to perform an on-demand run. The industry data service throttles the start of runs, allowing up to five successful runs every 12 hours. An on-demand run causes the next system-scheduled run to be skipped if it is set to run within the next 10 hours.
// returns a *RequestInformation when successful
func (m *IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder when successful
func (m *IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder) WithUrl(rawUrl string)(*IndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder) {
    return NewIndustryDataRunsMicrosoftGraphIndustryDataStartRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
