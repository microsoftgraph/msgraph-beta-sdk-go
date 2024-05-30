package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder provides operations to call the retrieveMyRequests method.
type OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderGetQueryParameters invoke function retrieveMyRequests
type OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderGetQueryParameters struct {
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
// OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderGetQueryParameters
}
// NewOperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderInternal instantiates a new OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder and sets the default values.
func NewOperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder) {
    m := &OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/retrieveMyRequests(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewOperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder instantiates a new OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder and sets the default values.
func NewOperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function retrieveMyRequests
// Deprecated: This method is obsolete. Use GetAsRetrieveMyRequestsGetResponse instead.
// returns a OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderGetRequestConfiguration)(OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsResponseable), nil
}
// GetAsRetrieveMyRequestsGetResponse invoke function retrieveMyRequests
// returns a OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder) GetAsRetrieveMyRequestsGetResponse(ctx context.Context, requestConfiguration *OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderGetRequestConfiguration)(OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsGetResponseable), nil
}
// ToGetRequestInformation invoke function retrieveMyRequests
// returns a *RequestInformation when successful
func (m *OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder when successful
func (m *OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder) WithUrl(rawUrl string)(*OperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder) {
    return NewOperationapprovalrequestsRetrievemyrequestsRetrieveMyRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
