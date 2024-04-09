package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServicePrincipalCreationPoliciesItemExcludesRequestBuilder provides operations to manage the excludes property of the microsoft.graph.servicePrincipalCreationPolicy entity.
type ServicePrincipalCreationPoliciesItemExcludesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServicePrincipalCreationPoliciesItemExcludesRequestBuilderGetQueryParameters get excludes from policies
type ServicePrincipalCreationPoliciesItemExcludesRequestBuilderGetQueryParameters struct {
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
// ServicePrincipalCreationPoliciesItemExcludesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalCreationPoliciesItemExcludesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalCreationPoliciesItemExcludesRequestBuilderGetQueryParameters
}
// ServicePrincipalCreationPoliciesItemExcludesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalCreationPoliciesItemExcludesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByServicePrincipalCreationConditionSetId provides operations to manage the excludes property of the microsoft.graph.servicePrincipalCreationPolicy entity.
// returns a *ServicePrincipalCreationPoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder when successful
func (m *ServicePrincipalCreationPoliciesItemExcludesRequestBuilder) ByServicePrincipalCreationConditionSetId(servicePrincipalCreationConditionSetId string)(*ServicePrincipalCreationPoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if servicePrincipalCreationConditionSetId != "" {
        urlTplParams["servicePrincipalCreationConditionSet%2Did"] = servicePrincipalCreationConditionSetId
    }
    return NewServicePrincipalCreationPoliciesItemExcludesServicePrincipalCreationConditionSetItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewServicePrincipalCreationPoliciesItemExcludesRequestBuilderInternal instantiates a new ServicePrincipalCreationPoliciesItemExcludesRequestBuilder and sets the default values.
func NewServicePrincipalCreationPoliciesItemExcludesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalCreationPoliciesItemExcludesRequestBuilder) {
    m := &ServicePrincipalCreationPoliciesItemExcludesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/servicePrincipalCreationPolicies/{servicePrincipalCreationPolicy%2Did}/excludes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewServicePrincipalCreationPoliciesItemExcludesRequestBuilder instantiates a new ServicePrincipalCreationPoliciesItemExcludesRequestBuilder and sets the default values.
func NewServicePrincipalCreationPoliciesItemExcludesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalCreationPoliciesItemExcludesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalCreationPoliciesItemExcludesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ServicePrincipalCreationPoliciesItemExcludesCountRequestBuilder when successful
func (m *ServicePrincipalCreationPoliciesItemExcludesRequestBuilder) Count()(*ServicePrincipalCreationPoliciesItemExcludesCountRequestBuilder) {
    return NewServicePrincipalCreationPoliciesItemExcludesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get excludes from policies
// returns a ServicePrincipalCreationConditionSetCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServicePrincipalCreationPoliciesItemExcludesRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalCreationPoliciesItemExcludesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalCreationConditionSetCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetCollectionResponseable), nil
}
// Post create new navigation property to excludes for policies
// returns a ServicePrincipalCreationConditionSetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServicePrincipalCreationPoliciesItemExcludesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable, requestConfiguration *ServicePrincipalCreationPoliciesItemExcludesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get excludes from policies
// returns a *RequestInformation when successful
func (m *ServicePrincipalCreationPoliciesItemExcludesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalCreationPoliciesItemExcludesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to excludes for policies
// returns a *RequestInformation when successful
func (m *ServicePrincipalCreationPoliciesItemExcludesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable, requestConfiguration *ServicePrincipalCreationPoliciesItemExcludesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServicePrincipalCreationPoliciesItemExcludesRequestBuilder when successful
func (m *ServicePrincipalCreationPoliciesItemExcludesRequestBuilder) WithUrl(rawUrl string)(*ServicePrincipalCreationPoliciesItemExcludesRequestBuilder) {
    return NewServicePrincipalCreationPoliciesItemExcludesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
