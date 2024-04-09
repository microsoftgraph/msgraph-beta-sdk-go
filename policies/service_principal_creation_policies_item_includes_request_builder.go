package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServicePrincipalCreationPoliciesItemIncludesRequestBuilder provides operations to manage the includes property of the microsoft.graph.servicePrincipalCreationPolicy entity.
type ServicePrincipalCreationPoliciesItemIncludesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServicePrincipalCreationPoliciesItemIncludesRequestBuilderGetQueryParameters get includes from policies
type ServicePrincipalCreationPoliciesItemIncludesRequestBuilderGetQueryParameters struct {
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
// ServicePrincipalCreationPoliciesItemIncludesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalCreationPoliciesItemIncludesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalCreationPoliciesItemIncludesRequestBuilderGetQueryParameters
}
// ServicePrincipalCreationPoliciesItemIncludesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalCreationPoliciesItemIncludesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByServicePrincipalCreationConditionSetId provides operations to manage the includes property of the microsoft.graph.servicePrincipalCreationPolicy entity.
// returns a *ServicePrincipalCreationPoliciesItemIncludesServicePrincipalCreationConditionSetItemRequestBuilder when successful
func (m *ServicePrincipalCreationPoliciesItemIncludesRequestBuilder) ByServicePrincipalCreationConditionSetId(servicePrincipalCreationConditionSetId string)(*ServicePrincipalCreationPoliciesItemIncludesServicePrincipalCreationConditionSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if servicePrincipalCreationConditionSetId != "" {
        urlTplParams["servicePrincipalCreationConditionSet%2Did"] = servicePrincipalCreationConditionSetId
    }
    return NewServicePrincipalCreationPoliciesItemIncludesServicePrincipalCreationConditionSetItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewServicePrincipalCreationPoliciesItemIncludesRequestBuilderInternal instantiates a new ServicePrincipalCreationPoliciesItemIncludesRequestBuilder and sets the default values.
func NewServicePrincipalCreationPoliciesItemIncludesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalCreationPoliciesItemIncludesRequestBuilder) {
    m := &ServicePrincipalCreationPoliciesItemIncludesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/servicePrincipalCreationPolicies/{servicePrincipalCreationPolicy%2Did}/includes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewServicePrincipalCreationPoliciesItemIncludesRequestBuilder instantiates a new ServicePrincipalCreationPoliciesItemIncludesRequestBuilder and sets the default values.
func NewServicePrincipalCreationPoliciesItemIncludesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalCreationPoliciesItemIncludesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalCreationPoliciesItemIncludesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ServicePrincipalCreationPoliciesItemIncludesCountRequestBuilder when successful
func (m *ServicePrincipalCreationPoliciesItemIncludesRequestBuilder) Count()(*ServicePrincipalCreationPoliciesItemIncludesCountRequestBuilder) {
    return NewServicePrincipalCreationPoliciesItemIncludesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get includes from policies
// returns a ServicePrincipalCreationConditionSetCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServicePrincipalCreationPoliciesItemIncludesRequestBuilder) Get(ctx context.Context, requestConfiguration *ServicePrincipalCreationPoliciesItemIncludesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetCollectionResponseable, error) {
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
// Post create new navigation property to includes for policies
// returns a ServicePrincipalCreationConditionSetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServicePrincipalCreationPoliciesItemIncludesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable, requestConfiguration *ServicePrincipalCreationPoliciesItemIncludesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable, error) {
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
// ToGetRequestInformation get includes from policies
// returns a *RequestInformation when successful
func (m *ServicePrincipalCreationPoliciesItemIncludesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServicePrincipalCreationPoliciesItemIncludesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to includes for policies
// returns a *RequestInformation when successful
func (m *ServicePrincipalCreationPoliciesItemIncludesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalCreationConditionSetable, requestConfiguration *ServicePrincipalCreationPoliciesItemIncludesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServicePrincipalCreationPoliciesItemIncludesRequestBuilder when successful
func (m *ServicePrincipalCreationPoliciesItemIncludesRequestBuilder) WithUrl(rawUrl string)(*ServicePrincipalCreationPoliciesItemIncludesRequestBuilder) {
    return NewServicePrincipalCreationPoliciesItemIncludesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
