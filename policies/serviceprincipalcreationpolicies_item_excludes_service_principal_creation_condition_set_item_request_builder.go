package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder provides operations to manage the excludes property of the microsoft.graph.servicePrincipalCreationPolicy entity.
type ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderGetQueryParameters get excludes from policies
type ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderGetQueryParameters
}
// ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderInternal instantiates a new ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder and sets the default values.
func NewServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) {
    m := &ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/servicePrincipalCreationPolicies/{servicePrincipalCreationPolicy%2Did}/excludes/{servicePrincipalCreationConditionSet%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder instantiates a new ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder and sets the default values.
func NewServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property excludes for policies
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get excludes from policies
// returns a ServicePrincipalCreationConditionSetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalCreationConditionSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable), nil
}
// Patch update the navigation property excludes in policies
// returns a ServicePrincipalCreationConditionSetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable, requestConfiguration *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalCreationConditionSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable), nil
}
// ToDeleteRequestInformation delete navigation property excludes for policies
// returns a *RequestInformation when successful
func (m *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get excludes from policies
// returns a *RequestInformation when successful
func (m *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property excludes in policies
// returns a *RequestInformation when successful
func (m *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable, requestConfiguration *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder when successful
func (m *ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) WithUrl(rawUrl string)(*ServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) {
    return NewServiceprincipalcreationpoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
