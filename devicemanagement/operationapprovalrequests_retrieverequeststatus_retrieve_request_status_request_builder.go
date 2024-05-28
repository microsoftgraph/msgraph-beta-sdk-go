package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder provides operations to call the retrieveRequestStatus method.
type OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilderInternal instantiates a new OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder and sets the default values.
func NewOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder) {
    m := &OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/retrieveRequestStatus", pathParameters),
    }
    return m
}
// NewOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder instantiates a new OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder and sets the default values.
func NewOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action retrieveRequestStatus
// returns a OperationApprovalRequestEntityStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder) Post(ctx context.Context, body OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBodyable, requestConfiguration *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestEntityStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *RequestInformation when successful
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusPostRequestBodyable, requestConfiguration *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder when successful
func (m *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder) WithUrl(rawUrl string)(*OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder) {
    return NewOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
