package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder provides operations to call the retrieveApprovableOperations method.
type OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilderGetQueryParameters invoke function retrieveApprovableOperations
type OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilderGetQueryParameters struct {
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
// OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilderGetQueryParameters
}
// NewOperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilderInternal instantiates a new OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder and sets the default values.
func NewOperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder) {
    m := &OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalPolicies/retrieveApprovableOperations(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewOperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder instantiates a new OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder and sets the default values.
func NewOperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function retrieveApprovableOperations
// Deprecated: This method is obsolete. Use GetAsRetrieveApprovableOperationsGetResponse instead.
// returns a OperationApprovalPoliciesRetrieveApprovableOperationsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilderGetRequestConfiguration)(OperationApprovalPoliciesRetrieveApprovableOperationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalPoliciesRetrieveApprovableOperationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalPoliciesRetrieveApprovableOperationsResponseable), nil
}
// GetAsRetrieveApprovableOperationsGetResponse invoke function retrieveApprovableOperations
// returns a OperationApprovalPoliciesRetrieveApprovableOperationsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder) GetAsRetrieveApprovableOperationsGetResponse(ctx context.Context, requestConfiguration *OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilderGetRequestConfiguration)(OperationApprovalPoliciesRetrieveApprovableOperationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalPoliciesRetrieveApprovableOperationsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalPoliciesRetrieveApprovableOperationsGetResponseable), nil
}
// ToGetRequestInformation invoke function retrieveApprovableOperations
// returns a *RequestInformation when successful
func (m *OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder when successful
func (m *OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder) WithUrl(rawUrl string)(*OperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder) {
    return NewOperationApprovalPoliciesRetrieveApprovableOperationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
