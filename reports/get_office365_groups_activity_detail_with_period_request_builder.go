package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetOffice365GroupsActivityDetailWithPeriodRequestBuilder provides operations to call the getOffice365GroupsActivityDetail method.
type GetOffice365GroupsActivityDetailWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetQueryParameters invoke function getOffice365GroupsActivityDetail
type GetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetQueryParameters struct {
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
// GetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetQueryParameters
}
// NewGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal instantiates a new GetOffice365GroupsActivityDetailWithPeriodRequestBuilder and sets the default values.
func NewGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    m := &GetOffice365GroupsActivityDetailWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365GroupsActivityDetail(period='{period}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetOffice365GroupsActivityDetailWithPeriodRequestBuilder instantiates a new GetOffice365GroupsActivityDetailWithPeriodRequestBuilder and sets the default values.
func NewGetOffice365GroupsActivityDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365GroupsActivityDetail
// Deprecated: This method is obsolete. Use GetAsGetOffice365GroupsActivityDetailWithPeriodGetResponse instead.
// returns a GetOffice365GroupsActivityDetailWithPeriodResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration)(GetOffice365GroupsActivityDetailWithPeriodResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetOffice365GroupsActivityDetailWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetOffice365GroupsActivityDetailWithPeriodResponseable), nil
}
// GetAsGetOffice365GroupsActivityDetailWithPeriodGetResponse invoke function getOffice365GroupsActivityDetail
// returns a GetOffice365GroupsActivityDetailWithPeriodGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) GetAsGetOffice365GroupsActivityDetailWithPeriodGetResponse(ctx context.Context, requestConfiguration *GetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration)(GetOffice365GroupsActivityDetailWithPeriodGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetOffice365GroupsActivityDetailWithPeriodGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetOffice365GroupsActivityDetailWithPeriodGetResponseable), nil
}
// ToGetRequestInformation invoke function getOffice365GroupsActivityDetail
// returns a *RequestInformation when successful
func (m *GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetOffice365GroupsActivityDetailWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetOffice365GroupsActivityDetailWithPeriodRequestBuilder when successful
func (m *GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return NewGetOffice365GroupsActivityDetailWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
