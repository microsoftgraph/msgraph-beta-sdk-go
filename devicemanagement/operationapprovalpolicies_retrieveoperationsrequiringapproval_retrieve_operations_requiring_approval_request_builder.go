package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder provides operations to call the retrieveOperationsRequiringApproval method.
type OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderGetQueryParameters invoke function retrieveOperationsRequiringApproval
type OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderGetQueryParameters struct {
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
// OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderGetQueryParameters
}
// NewOperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderInternal instantiates a new OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder and sets the default values.
func NewOperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder) {
    m := &OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalPolicies/retrieveOperationsRequiringApproval(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewOperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder instantiates a new OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder and sets the default values.
func NewOperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function retrieveOperationsRequiringApproval
// Deprecated: This method is obsolete. Use GetAsRetrieveOperationsRequiringApprovalGetResponse instead.
// returns a OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderGetRequestConfiguration)(OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalResponseable), nil
}
// GetAsRetrieveOperationsRequiringApprovalGetResponse invoke function retrieveOperationsRequiringApproval
// returns a OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder) GetAsRetrieveOperationsRequiringApprovalGetResponse(ctx context.Context, requestConfiguration *OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderGetRequestConfiguration)(OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalGetResponseable), nil
}
// ToGetRequestInformation invoke function retrieveOperationsRequiringApproval
// returns a *RequestInformation when successful
func (m *OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder when successful
func (m *OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder) WithUrl(rawUrl string)(*OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder) {
    return NewOperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
