package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder provides operations to call the getOperationsRequiringApproval method.
type OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilderGetQueryParameters invoke function getOperationsRequiringApproval
type OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilderGetQueryParameters struct {
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
// OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilderGetQueryParameters
}
// NewOperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilderInternal instantiates a new GetOperationsRequiringApprovalRequestBuilder and sets the default values.
func NewOperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder) {
    m := &OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalPolicies/getOperationsRequiringApproval(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewOperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder instantiates a new GetOperationsRequiringApprovalRequestBuilder and sets the default values.
func NewOperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getOperationsRequiringApproval
// Deprecated: This method is obsolete. Use GetAsGetOperationsRequiringApprovalGetResponse instead.
func (m *OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilderGetRequestConfiguration)(OperationApprovalPoliciesGetOperationsRequiringApprovalResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalPoliciesGetOperationsRequiringApprovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalPoliciesGetOperationsRequiringApprovalResponseable), nil
}
// GetAsGetOperationsRequiringApprovalGetResponse invoke function getOperationsRequiringApproval
func (m *OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder) GetAsGetOperationsRequiringApprovalGetResponse(ctx context.Context, requestConfiguration *OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilderGetRequestConfiguration)(OperationApprovalPoliciesGetOperationsRequiringApprovalGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalPoliciesGetOperationsRequiringApprovalGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalPoliciesGetOperationsRequiringApprovalGetResponseable), nil
}
// ToGetRequestInformation invoke function getOperationsRequiringApproval
func (m *OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder) WithUrl(rawUrl string)(*OperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder) {
    return NewOperationApprovalPoliciesGetOperationsRequiringApprovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
