package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder provides operations to manage the userAttribute property of the microsoft.graph.identityUserFlowAttributeAssignment entity.
type B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilderGetQueryParameters the user attribute that you want to add to your user flow.
type B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilderGetQueryParameters
}
// NewB2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilderInternal instantiates a new B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder and sets the default values.
func NewB2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder) {
    m := &B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/userAttributeAssignments/{identityUserFlowAttributeAssignment%2Did}/userAttribute{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewB2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder instantiates a new B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder and sets the default values.
func NewB2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the user attribute that you want to add to your user flow.
// returns a IdentityUserFlowAttributeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityUserFlowAttributeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityUserFlowAttributeable), nil
}
// ToGetRequestInformation the user attribute that you want to add to your user flow.
// returns a *RequestInformation when successful
func (m *B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder when successful
func (m *B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder) WithUrl(rawUrl string)(*B2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder) {
    return NewB2xUserFlowsItemUserAttributeAssignmentsItemUserAttributeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
