package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder provides operations to manage the operationApprovalRequests property of the microsoft.graph.deviceManagement entity.
type OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderGetQueryParameters the Operation Approval Requests
type OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderGetQueryParameters
}
// OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Approve provides operations to call the approve method.
// returns a *OperationApprovalRequestsItemApproveRequestBuilder when successful
func (m *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) Approve()(*OperationApprovalRequestsItemApproveRequestBuilder) {
    return NewOperationApprovalRequestsItemApproveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CancelApproval provides operations to call the cancelApproval method.
// returns a *OperationApprovalRequestsItemCancelApprovalRequestBuilder when successful
func (m *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) CancelApproval()(*OperationApprovalRequestsItemCancelApprovalRequestBuilder) {
    return NewOperationApprovalRequestsItemCancelApprovalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewOperationApprovalRequestsOperationApprovalRequestItemRequestBuilderInternal instantiates a new OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder and sets the default values.
func NewOperationApprovalRequestsOperationApprovalRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) {
    m := &OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/{operationApprovalRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewOperationApprovalRequestsOperationApprovalRequestItemRequestBuilder instantiates a new OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder and sets the default values.
func NewOperationApprovalRequestsOperationApprovalRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationApprovalRequestsOperationApprovalRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property operationApprovalRequests for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, error) {
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
func (m *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, requestConfiguration *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, error) {
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
// returns a *OperationApprovalRequestsItemRejectRequestBuilder when successful
func (m *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) Reject()(*OperationApprovalRequestsItemRejectRequestBuilder) {
    return NewOperationApprovalRequestsItemRejectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property operationApprovalRequests for deviceManagement
// returns a *RequestInformation when successful
func (m *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/operationApprovalRequests/{operationApprovalRequest%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the Operation Approval Requests
// returns a *RequestInformation when successful
func (m *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, requestConfiguration *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/operationApprovalRequests/{operationApprovalRequest%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder when successful
func (m *OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) WithUrl(rawUrl string)(*OperationApprovalRequestsOperationApprovalRequestItemRequestBuilder) {
    return NewOperationApprovalRequestsOperationApprovalRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
