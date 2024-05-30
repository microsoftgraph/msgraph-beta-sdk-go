package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder provides operations to call the getAttackSimulationRepeatOffenders method.
type GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilderGetQueryParameters list the users of a tenant who have yielded to attacks more than once in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
type GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilderGetQueryParameters struct {
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
// GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilderGetQueryParameters
}
// NewGetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilderInternal instantiates a new GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewGetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder) {
    m := &GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getAttackSimulationRepeatOffenders(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder instantiates a new GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewGetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list the users of a tenant who have yielded to attacks more than once in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// Deprecated: This method is obsolete. Use GetAsGetAttackSimulationRepeatOffendersGetResponse instead.
// returns a GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getattacksimulationrepeatoffenders?view=graph-rest-beta
func (m *GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder) Get(ctx context.Context, requestConfiguration *GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersResponseable), nil
}
// GetAsGetAttackSimulationRepeatOffendersGetResponse list the users of a tenant who have yielded to attacks more than once in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// Deprecated: This report function api is deprecated and will stop returning data on August 20, 2022. Api is now moved to /reports/security. Please use the new API. as of 2022-05/Tasks_And_Plans
// returns a GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getattacksimulationrepeatoffenders?view=graph-rest-beta
func (m *GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder) GetAsGetAttackSimulationRepeatOffendersGetResponse(ctx context.Context, requestConfiguration *GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersGetResponseable), nil
}
// ToGetRequestInformation list the users of a tenant who have yielded to attacks more than once in attack simulation and training campaigns. This function supports @odata.nextLink for pagination.
// Deprecated: This report function api is deprecated and will stop returning data on August 20, 2022. Api is now moved to /reports/security. Please use the new API. as of 2022-05/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder when successful
func (m *GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder) WithUrl(rawUrl string)(*GetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder) {
    return NewGetattacksimulationrepeatoffendersGetAttackSimulationRepeatOffendersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
