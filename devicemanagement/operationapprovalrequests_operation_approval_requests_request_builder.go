package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationapprovalrequestsOperationApprovalRequestsRequestBuilder provides operations to manage the operationApprovalRequests property of the microsoft.graph.deviceManagement entity.
type OperationapprovalrequestsOperationApprovalRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationapprovalrequestsOperationApprovalRequestsRequestBuilderGetQueryParameters the Operation Approval Requests
type OperationapprovalrequestsOperationApprovalRequestsRequestBuilderGetQueryParameters struct {
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
// OperationapprovalrequestsOperationApprovalRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalrequestsOperationApprovalRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OperationapprovalrequestsOperationApprovalRequestsRequestBuilderGetQueryParameters
}
// OperationapprovalrequestsOperationApprovalRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalrequestsOperationApprovalRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByOperationApprovalRequestId provides operations to manage the operationApprovalRequests property of the microsoft.graph.deviceManagement entity.
// returns a *OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder when successful
func (m *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) ByOperationApprovalRequestId(operationApprovalRequestId string)(*OperationapprovalrequestsOperationApprovalRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if operationApprovalRequestId != "" {
        urlTplParams["operationApprovalRequest%2Did"] = operationApprovalRequestId
    }
    return NewOperationapprovalrequestsOperationApprovalRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// CancelMyRequest provides operations to call the cancelMyRequest method.
// returns a *OperationapprovalrequestsCancelmyrequestCancelMyRequestRequestBuilder when successful
func (m *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) CancelMyRequest()(*OperationapprovalrequestsCancelmyrequestCancelMyRequestRequestBuilder) {
    return NewOperationapprovalrequestsCancelmyrequestCancelMyRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewOperationapprovalrequestsOperationApprovalRequestsRequestBuilderInternal instantiates a new OperationapprovalrequestsOperationApprovalRequestsRequestBuilder and sets the default values.
func NewOperationapprovalrequestsOperationApprovalRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) {
    m := &OperationapprovalrequestsOperationApprovalRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewOperationapprovalrequestsOperationApprovalRequestsRequestBuilder instantiates a new OperationapprovalrequestsOperationApprovalRequestsRequestBuilder and sets the default values.
func NewOperationapprovalrequestsOperationApprovalRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationapprovalrequestsOperationApprovalRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *OperationapprovalrequestsCountRequestBuilder when successful
func (m *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) Count()(*OperationapprovalrequestsCountRequestBuilder) {
    return NewOperationapprovalrequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the Operation Approval Requests
// returns a OperationApprovalRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationapprovalrequestsOperationApprovalRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOperationApprovalRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestCollectionResponseable), nil
}
// Post create new navigation property to operationApprovalRequests for deviceManagement
// returns a OperationApprovalRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, requestConfiguration *OperationapprovalrequestsOperationApprovalRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// RetrieveMyRequestByIdWithId provides operations to call the retrieveMyRequestById method.
// returns a *OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder when successful
func (m *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) RetrieveMyRequestByIdWithId(id *string)(*OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder) {
    return NewOperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, id)
}
// RetrieveMyRequests provides operations to call the retrieveMyRequests method.
// returns a *OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder when successful
func (m *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) RetrieveMyRequests()(*OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder) {
    return NewOperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveRequestStatus provides operations to call the retrieveRequestStatus method.
// returns a *OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder when successful
func (m *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) RetrieveRequestStatus()(*OperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilder) {
    return NewOperationapprovalrequestsRetrieverequeststatusRetrieveRequestStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the Operation Approval Requests
// returns a *RequestInformation when successful
func (m *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationapprovalrequestsOperationApprovalRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to operationApprovalRequests for deviceManagement
// returns a *RequestInformation when successful
func (m *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, requestConfiguration *OperationapprovalrequestsOperationApprovalRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder when successful
func (m *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) WithUrl(rawUrl string)(*OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) {
    return NewOperationapprovalrequestsOperationApprovalRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
