package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder provides operations to call the retrieveMyRequestById method.
type OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilderInternal instantiates a new OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder and sets the default values.
func NewOperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, id *string)(*OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder) {
    m := &OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/retrieveMyRequestById(id='{id}')", pathParameters),
    }
    if id != nil {
        m.BaseRequestBuilder.PathParameters["id"] = *id
    }
    return m
}
// NewOperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder instantiates a new OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder and sets the default values.
func NewOperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function retrieveMyRequestById
// returns a OperationApprovalRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OperationApprovalRequestable, error) {
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
// ToGetRequestInformation invoke function retrieveMyRequestById
// returns a *RequestInformation when successful
func (m *OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder when successful
func (m *OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder) WithUrl(rawUrl string)(*OperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder) {
    return NewOperationapprovalrequestsRetrievemyrequestbyidwithidRetrieveMyRequestByIdWithIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
