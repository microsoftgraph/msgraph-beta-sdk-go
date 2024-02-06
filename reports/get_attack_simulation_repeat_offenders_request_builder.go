package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetAttackSimulationRepeatOffendersRequestBuilder provides operations to call the getAttackSimulationRepeatOffenders method.
type GetAttackSimulationRepeatOffendersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetAttackSimulationRepeatOffendersRequestBuilderGetQueryParameters invoke function getAttackSimulationRepeatOffenders
type GetAttackSimulationRepeatOffendersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetAttackSimulationRepeatOffendersRequestBuilderGetQueryParameters
}
// NewGetAttackSimulationRepeatOffendersRequestBuilderInternal instantiates a new GetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAttackSimulationRepeatOffendersRequestBuilder) {
    m := &GetAttackSimulationRepeatOffendersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getAttackSimulationRepeatOffenders(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGetAttackSimulationRepeatOffendersRequestBuilder instantiates a new GetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewGetAttackSimulationRepeatOffendersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAttackSimulationRepeatOffendersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getAttackSimulationRepeatOffenders
// Deprecated: This method is obsolete. Use GetAsGetAttackSimulationRepeatOffendersGetResponse instead.
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) Get(ctx context.Context, requestConfiguration *GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(GetAttackSimulationRepeatOffendersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetAttackSimulationRepeatOffendersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetAttackSimulationRepeatOffendersResponseable), nil
}
// GetAsGetAttackSimulationRepeatOffendersGetResponse invoke function getAttackSimulationRepeatOffenders
// Deprecated: This report function api is deprecated and will stop returning data on August 20, 2022. Api is now moved to /reports/security. Please use the new API. as of 2022-05/Tasks_And_Plans
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) GetAsGetAttackSimulationRepeatOffendersGetResponse(ctx context.Context, requestConfiguration *GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(GetAttackSimulationRepeatOffendersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetAttackSimulationRepeatOffendersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetAttackSimulationRepeatOffendersGetResponseable), nil
}
// ToGetRequestInformation invoke function getAttackSimulationRepeatOffenders
// Deprecated: This report function api is deprecated and will stop returning data on August 20, 2022. Api is now moved to /reports/security. Please use the new API. as of 2022-05/Tasks_And_Plans
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: This report function api is deprecated and will stop returning data on August 20, 2022. Api is now moved to /reports/security. Please use the new API. as of 2022-05/Tasks_And_Plans
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) WithUrl(rawUrl string)(*GetAttackSimulationRepeatOffendersRequestBuilder) {
    return NewGetAttackSimulationRepeatOffendersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
