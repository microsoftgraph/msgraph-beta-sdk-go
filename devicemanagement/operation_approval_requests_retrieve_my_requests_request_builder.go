package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OperationApprovalRequestsRetrieveMyRequestsRequestBuilder provides operations to call the retrieveMyRequests method.
type OperationApprovalRequestsRetrieveMyRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OperationApprovalRequestsRetrieveMyRequestsRequestBuilderGetQueryParameters invoke function retrieveMyRequests
type OperationApprovalRequestsRetrieveMyRequestsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// OperationApprovalRequestsRetrieveMyRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OperationApprovalRequestsRetrieveMyRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OperationApprovalRequestsRetrieveMyRequestsRequestBuilderGetQueryParameters
}
// NewOperationApprovalRequestsRetrieveMyRequestsRequestBuilderInternal instantiates a new RetrieveMyRequestsRequestBuilder and sets the default values.
func NewOperationApprovalRequestsRetrieveMyRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalRequestsRetrieveMyRequestsRequestBuilder) {
    m := &OperationApprovalRequestsRetrieveMyRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/operationApprovalRequests/retrieveMyRequests(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    return m
}
// NewOperationApprovalRequestsRetrieveMyRequestsRequestBuilder instantiates a new RetrieveMyRequestsRequestBuilder and sets the default values.
func NewOperationApprovalRequestsRetrieveMyRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OperationApprovalRequestsRetrieveMyRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOperationApprovalRequestsRetrieveMyRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function retrieveMyRequests
// Deprecated: This method is obsolete. Use GetAsRetrieveMyRequestsGetResponse instead.
func (m *OperationApprovalRequestsRetrieveMyRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *OperationApprovalRequestsRetrieveMyRequestsRequestBuilderGetRequestConfiguration)(OperationApprovalRequestsRetrieveMyRequestsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalRequestsRetrieveMyRequestsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalRequestsRetrieveMyRequestsResponseable), nil
}
// GetAsRetrieveMyRequestsGetResponse invoke function retrieveMyRequests
func (m *OperationApprovalRequestsRetrieveMyRequestsRequestBuilder) GetAsRetrieveMyRequestsGetResponse(ctx context.Context, requestConfiguration *OperationApprovalRequestsRetrieveMyRequestsRequestBuilderGetRequestConfiguration)(OperationApprovalRequestsRetrieveMyRequestsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOperationApprovalRequestsRetrieveMyRequestsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(OperationApprovalRequestsRetrieveMyRequestsGetResponseable), nil
}
// ToGetRequestInformation invoke function retrieveMyRequests
func (m *OperationApprovalRequestsRetrieveMyRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OperationApprovalRequestsRetrieveMyRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *OperationApprovalRequestsRetrieveMyRequestsRequestBuilder) WithUrl(rawUrl string)(*OperationApprovalRequestsRetrieveMyRequestsRequestBuilder) {
    return NewOperationApprovalRequestsRetrieveMyRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
