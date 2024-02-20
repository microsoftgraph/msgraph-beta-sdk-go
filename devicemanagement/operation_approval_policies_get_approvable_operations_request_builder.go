package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationApprovalPoliciesGetApprovableOperationsRequestBuilder provides operations to call the getApprovableOperations method.
type OperationApprovalPoliciesGetApprovableOperationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationApprovalPoliciesGetApprovableOperationsRequestBuilderGetQueryParameters invoke function getApprovableOperations
type OperationApprovalPoliciesGetApprovableOperationsRequestBuilderGetQueryParameters struct {
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
// OperationApprovalPoliciesGetApprovableOperationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationApprovalPoliciesGetApprovableOperationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OperationApprovalPoliciesGetApprovableOperationsRequestBuilderGetQueryParameters
}
// NewOperationApprovalPoliciesGetApprovableOperationsRequestBuilderInternal instantiates a new OperationApprovalPoliciesGetApprovableOperationsRequestBuilder and sets the default values.
func NewOperationApprovalPoliciesGetApprovableOperationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalPoliciesGetApprovableOperationsRequestBuilder) {
    m := &OperationApprovalPoliciesGetApprovableOperationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalPolicies/getApprovableOperations(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewOperationApprovalPoliciesGetApprovableOperationsRequestBuilder instantiates a new OperationApprovalPoliciesGetApprovableOperationsRequestBuilder and sets the default values.
func NewOperationApprovalPoliciesGetApprovableOperationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalPoliciesGetApprovableOperationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationApprovalPoliciesGetApprovableOperationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getApprovableOperations
// Deprecated: This method is obsolete. Use GetAsGetApprovableOperationsGetResponse instead.
// returns a OperationApprovalPoliciesGetApprovableOperationsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationApprovalPoliciesGetApprovableOperationsRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationApprovalPoliciesGetApprovableOperationsRequestBuilderGetRequestConfiguration)(OperationApprovalPoliciesGetApprovableOperationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalPoliciesGetApprovableOperationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalPoliciesGetApprovableOperationsResponseable), nil
}
// GetAsGetApprovableOperationsGetResponse invoke function getApprovableOperations
// returns a OperationApprovalPoliciesGetApprovableOperationsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationApprovalPoliciesGetApprovableOperationsRequestBuilder) GetAsGetApprovableOperationsGetResponse(ctx context.Context, requestConfiguration *OperationApprovalPoliciesGetApprovableOperationsRequestBuilderGetRequestConfiguration)(OperationApprovalPoliciesGetApprovableOperationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalPoliciesGetApprovableOperationsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalPoliciesGetApprovableOperationsGetResponseable), nil
}
// ToGetRequestInformation invoke function getApprovableOperations
// returns a *RequestInformation when successful
func (m *OperationApprovalPoliciesGetApprovableOperationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationApprovalPoliciesGetApprovableOperationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationApprovalPoliciesGetApprovableOperationsRequestBuilder when successful
func (m *OperationApprovalPoliciesGetApprovableOperationsRequestBuilder) WithUrl(rawUrl string)(*OperationApprovalPoliciesGetApprovableOperationsRequestBuilder) {
    return NewOperationApprovalPoliciesGetApprovableOperationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
