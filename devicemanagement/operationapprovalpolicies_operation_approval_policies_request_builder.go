package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder provides operations to manage the operationApprovalPolicies property of the microsoft.graph.deviceManagement entity.
type OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderGetQueryParameters the Operation Approval Policies
type OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderGetQueryParameters struct {
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
// OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderGetQueryParameters
}
// OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByOperationApprovalPolicyId provides operations to manage the operationApprovalPolicies property of the microsoft.graph.deviceManagement entity.
// returns a *OperationapprovalpoliciesOperationApprovalPolicyItemRequestBuilder when successful
func (m *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) ByOperationApprovalPolicyId(operationApprovalPolicyId string)(*OperationapprovalpoliciesOperationApprovalPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if operationApprovalPolicyId != "" {
        urlTplParams["operationApprovalPolicy%2Did"] = operationApprovalPolicyId
    }
    return NewOperationapprovalpoliciesOperationApprovalPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewOperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderInternal instantiates a new OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder and sets the default values.
func NewOperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) {
    m := &OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewOperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder instantiates a new OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder and sets the default values.
func NewOperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *OperationapprovalpoliciesCountRequestBuilder when successful
func (m *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) Count()(*OperationapprovalpoliciesCountRequestBuilder) {
    return NewOperationapprovalpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the Operation Approval Policies
// returns a OperationApprovalPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOperationApprovalPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalPolicyCollectionResponseable), nil
}
// Post create new navigation property to operationApprovalPolicies for deviceManagement
// returns a OperationApprovalPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalPolicyable, requestConfiguration *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOperationApprovalPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalPolicyable), nil
}
// RetrieveApprovableOperations provides operations to call the retrieveApprovableOperations method.
// returns a *OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder when successful
func (m *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) RetrieveApprovableOperations()(*OperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilder) {
    return NewOperationapprovalpoliciesRetrieveapprovableoperationsRetrieveApprovableOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveOperationsRequiringApproval provides operations to call the retrieveOperationsRequiringApproval method.
// returns a *OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder when successful
func (m *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) RetrieveOperationsRequiringApproval()(*OperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilder) {
    return NewOperationapprovalpoliciesRetrieveoperationsrequiringapprovalRetrieveOperationsRequiringApprovalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the Operation Approval Policies
// returns a *RequestInformation when successful
func (m *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to operationApprovalPolicies for deviceManagement
// returns a *RequestInformation when successful
func (m *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalPolicyable, requestConfiguration *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder when successful
func (m *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) WithUrl(rawUrl string)(*OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) {
    return NewOperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
