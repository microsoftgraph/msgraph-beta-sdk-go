package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder provides operations to manage the operationApprovalRequests property of the microsoft.graph.deviceManagement entity.
type OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderGetQueryParameters the Operation Approval Requests
type OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderGetQueryParameters
}
// OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Approve provides operations to call the approve method.
// returns a *OperationapprovalrequestsItemApproveRequestBuilder when successful
func (m *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) Approve()(*OperationapprovalrequestsItemApproveRequestBuilder) {
    return NewOperationapprovalrequestsItemApproveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CancelApproval provides operations to call the cancelApproval method.
// returns a *OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder when successful
func (m *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) CancelApproval()(*OperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilder) {
    return NewOperationapprovalrequestsItemCancelapprovalCancelApprovalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewOperationapprovalrequestsOperationApprovalRequestItemRequestBuilderInternal instantiates a new OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder and sets the default values.
func NewOperationapprovalrequestsOperationApprovalRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) {
    m := &OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/{operationApprovalRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewOperationapprovalrequestsOperationApprovalRequestItemRequestBuilder instantiates a new OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder and sets the default values.
func NewOperationapprovalrequestsOperationApprovalRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationapprovalrequestsOperationApprovalRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property operationApprovalRequests for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the Operation Approval Requests
// returns a OperationApprovalRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOperationApprovalRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable), nil
}
// Patch update the navigation property operationApprovalRequests in deviceManagement
// returns a OperationApprovalRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, requestConfiguration *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOperationApprovalRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable), nil
}
// Reject provides operations to call the reject method.
// returns a *OperationapprovalrequestsItemRejectRequestBuilder when successful
func (m *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) Reject()(*OperationapprovalrequestsItemRejectRequestBuilder) {
    return NewOperationapprovalrequestsItemRejectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property operationApprovalRequests for deviceManagement
// returns a *RequestInformation when successful
func (m *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the Operation Approval Requests
// returns a *RequestInformation when successful
func (m *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property operationApprovalRequests in deviceManagement
// returns a *RequestInformation when successful
func (m *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, requestConfiguration *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder when successful
func (m *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) WithUrl(rawUrl string)(*OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) {
    return NewOperationapprovalrequestsOperationApprovalRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
