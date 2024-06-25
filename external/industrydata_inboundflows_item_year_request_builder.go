package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataInboundflowsItemYearRequestBuilder provides operations to manage the year property of the microsoft.graph.industryData.inboundFlow entity.
type IndustrydataInboundflowsItemYearRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataInboundflowsItemYearRequestBuilderGetQueryParameters the year associated to the data that this flow brings in.
type IndustrydataInboundflowsItemYearRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustrydataInboundflowsItemYearRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataInboundflowsItemYearRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataInboundflowsItemYearRequestBuilderGetQueryParameters
}
// NewIndustrydataInboundflowsItemYearRequestBuilderInternal instantiates a new IndustrydataInboundflowsItemYearRequestBuilder and sets the default values.
func NewIndustrydataInboundflowsItemYearRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataInboundflowsItemYearRequestBuilder) {
    m := &IndustrydataInboundflowsItemYearRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/inboundFlows/{inboundFlow%2Did}/year{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIndustrydataInboundflowsItemYearRequestBuilder instantiates a new IndustrydataInboundflowsItemYearRequestBuilder and sets the default values.
func NewIndustrydataInboundflowsItemYearRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataInboundflowsItemYearRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataInboundflowsItemYearRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the year associated to the data that this flow brings in.
// returns a YearTimePeriodDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataInboundflowsItemYearRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataInboundflowsItemYearRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.YearTimePeriodDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateYearTimePeriodDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.YearTimePeriodDefinitionable), nil
}
// ToGetRequestInformation the year associated to the data that this flow brings in.
// returns a *RequestInformation when successful
func (m *IndustrydataInboundflowsItemYearRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataInboundflowsItemYearRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IndustrydataInboundflowsItemYearRequestBuilder when successful
func (m *IndustrydataInboundflowsItemYearRequestBuilder) WithUrl(rawUrl string)(*IndustrydataInboundflowsItemYearRequestBuilder) {
    return NewIndustrydataInboundflowsItemYearRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
