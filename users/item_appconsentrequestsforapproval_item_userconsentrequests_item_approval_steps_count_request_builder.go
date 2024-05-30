package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder provides operations to count the resources in the collection.
type ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilderGetQueryParameters
}
// NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilderInternal instantiates a new ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder and sets the default values.
func NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder) {
    m := &ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/appConsentRequestsForApproval/{appConsentRequest%2Did}/userConsentRequests/{userConsentRequest%2Did}/approval/steps/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder instantiates a new ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder and sets the default values.
func NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder when successful
func (m *ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder) WithUrl(rawUrl string)(*ItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder) {
    return NewItemAppconsentrequestsforapprovalItemUserconsentrequestsItemApprovalStepsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
