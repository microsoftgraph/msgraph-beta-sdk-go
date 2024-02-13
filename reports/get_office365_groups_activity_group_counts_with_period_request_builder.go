package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityGroupCounts method.
type GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetQueryParameters invoke function getOffice365GroupsActivityGroupCounts
type GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetQueryParameters struct {
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
// GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetQueryParameters
}
// NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal instantiates a new GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder and sets the default values.
func NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    m := &GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365GroupsActivityGroupCounts(period='{period}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder instantiates a new GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder and sets the default values.
func NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365GroupsActivityGroupCounts
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a GetOffice365GroupsActivityGroupCountsWithPeriodResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration)(GetOffice365GroupsActivityGroupCountsWithPeriodResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetOffice365GroupsActivityGroupCountsWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetOffice365GroupsActivityGroupCountsWithPeriodResponseable), nil
}
// GetAsGetOffice365GroupsActivityGroupCountsWithPeriodGetResponse invoke function getOffice365GroupsActivityGroupCounts
// returns a GetOffice365GroupsActivityGroupCountsWithPeriodGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) GetAsGetOffice365GroupsActivityGroupCountsWithPeriodGetResponse(ctx context.Context, requestConfiguration *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration)(GetOffice365GroupsActivityGroupCountsWithPeriodGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetOffice365GroupsActivityGroupCountsWithPeriodGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetOffice365GroupsActivityGroupCountsWithPeriodGetResponseable), nil
}
// ToGetRequestInformation invoke function getOffice365GroupsActivityGroupCounts
// returns a *RequestInformation when successful
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder when successful
func (m *GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
