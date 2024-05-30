package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder provides operations to call the retrieveReviewStatus method.
type ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilderInternal instantiates a new ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder and sets the default values.
func NewItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder) {
    m := &ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/retrieveReviewStatus()", pathParameters),
    }
    return m
}
// NewItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder instantiates a new ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder and sets the default values.
func NewItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function retrieveReviewStatus
// returns a CloudPcReviewStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcReviewStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable), nil
}
// ToGetRequestInformation invoke function retrieveReviewStatus
// returns a *RequestInformation when successful
func (m *ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder when successful
func (m *ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder) WithUrl(rawUrl string)(*ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder) {
    return NewItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
