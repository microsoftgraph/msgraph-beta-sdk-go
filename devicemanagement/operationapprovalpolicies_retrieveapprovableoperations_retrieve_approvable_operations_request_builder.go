package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder provides operations to call the retrieveApprovableOperations method.
type OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderGetQueryParameters invoke function retrieveApprovableOperations
type OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderGetQueryParameters struct {
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
// OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderGetQueryParameters
}
// NewOperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderInternal instantiates a new OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder and sets the default values.
func NewOperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder) {
    m := &OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalPolicies/retrieveApprovableOperations(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewOperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder instantiates a new OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder and sets the default values.
func NewOperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function retrieveApprovableOperations
// Deprecated: This method is obsolete. Use GetAsRetrieveApprovableOperationsGetResponse instead.
// returns a OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderGetRequestConfiguration)(OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsResponseable), nil
}
// GetAsRetrieveApprovableOperationsGetResponse invoke function retrieveApprovableOperations
// returns a OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder) GetAsRetrieveApprovableOperationsGetResponse(ctx context.Context, requestConfiguration *OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderGetRequestConfiguration)(OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsGetResponseable), nil
}
// ToGetRequestInformation invoke function retrieveApprovableOperations
// returns a *RequestInformation when successful
func (m *OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder when successful
func (m *OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder) WithUrl(rawUrl string)(*OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder) {
    return NewOperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
