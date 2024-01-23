package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationApprovalRequestsRetrieveRequestStatusRequestBuilder provides operations to call the retrieveRequestStatus method.
type OperationApprovalRequestsRetrieveRequestStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationApprovalRequestsRetrieveRequestStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationApprovalRequestsRetrieveRequestStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOperationApprovalRequestsRetrieveRequestStatusRequestBuilderInternal instantiates a new RetrieveRequestStatusRequestBuilder and sets the default values.
func NewOperationApprovalRequestsRetrieveRequestStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalRequestsRetrieveRequestStatusRequestBuilder) {
    m := &OperationApprovalRequestsRetrieveRequestStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/retrieveRequestStatus", pathParameters),
    }
    return m
}
// NewOperationApprovalRequestsRetrieveRequestStatusRequestBuilder instantiates a new RetrieveRequestStatusRequestBuilder and sets the default values.
func NewOperationApprovalRequestsRetrieveRequestStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalRequestsRetrieveRequestStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationApprovalRequestsRetrieveRequestStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action retrieveRequestStatus
func (m *OperationApprovalRequestsRetrieveRequestStatusRequestBuilder) Post(ctx context.Context, body OperationApprovalRequestsRetrieveRequestStatusPostRequestBodyable, requestConfiguration *OperationApprovalRequestsRetrieveRequestStatusRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestEntityStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOperationApprovalRequestEntityStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestEntityStatusable), nil
}
// ToPostRequestInformation invoke action retrieveRequestStatus
func (m *OperationApprovalRequestsRetrieveRequestStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body OperationApprovalRequestsRetrieveRequestStatusPostRequestBodyable, requestConfiguration *OperationApprovalRequestsRetrieveRequestStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OperationApprovalRequestsRetrieveRequestStatusRequestBuilder) WithUrl(rawUrl string)(*OperationApprovalRequestsRetrieveRequestStatusRequestBuilder) {
    return NewOperationApprovalRequestsRetrieveRequestStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
