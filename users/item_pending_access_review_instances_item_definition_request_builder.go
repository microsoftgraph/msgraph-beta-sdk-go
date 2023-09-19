package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingAccessReviewInstancesItemDefinitionRequestBuilder provides operations to manage the definition property of the microsoft.graph.accessReviewInstance entity.
type ItemPendingAccessReviewInstancesItemDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingAccessReviewInstancesItemDefinitionRequestBuilderGetQueryParameters there is exactly one accessReviewScheduleDefinition associated with each instance. It is the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
type ItemPendingAccessReviewInstancesItemDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPendingAccessReviewInstancesItemDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPendingAccessReviewInstancesItemDefinitionRequestBuilderGetQueryParameters
}
// NewItemPendingAccessReviewInstancesItemDefinitionRequestBuilderInternal instantiates a new DefinitionRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemDefinitionRequestBuilder) {
    m := &ItemPendingAccessReviewInstancesItemDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/definition{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemPendingAccessReviewInstancesItemDefinitionRequestBuilder instantiates a new DefinitionRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingAccessReviewInstancesItemDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get there is exactly one accessReviewScheduleDefinition associated with each instance. It is the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
func (m *ItemPendingAccessReviewInstancesItemDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDefinitionRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewScheduleDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewScheduleDefinitionable), nil
}
// ToGetRequestInformation there is exactly one accessReviewScheduleDefinition associated with each instance. It is the parent schedule for the instance, where instances are created for each recurrence of a review definition and each group selected to review by the definition.
func (m *ItemPendingAccessReviewInstancesItemDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemPendingAccessReviewInstancesItemDefinitionRequestBuilder) WithUrl(rawUrl string)(*ItemPendingAccessReviewInstancesItemDefinitionRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
