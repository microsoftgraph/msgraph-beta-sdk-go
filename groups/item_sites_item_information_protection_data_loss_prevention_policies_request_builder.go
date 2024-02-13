package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
type ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetQueryParameters get dataLossPreventionPolicies from groups
type ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetQueryParameters
}
// ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDataLossPreventionPolicyId provides operations to manage the dataLossPreventionPolicies property of the microsoft.graph.informationProtection entity.
// returns a *ItemSitesItemInformationProtectionDataLossPreventionPoliciesDataLossPreventionPolicyItemRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) ByDataLossPreventionPolicyId(dataLossPreventionPolicyId string)(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesDataLossPreventionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if dataLossPreventionPolicyId != "" {
        urlTplParams["dataLossPreventionPolicy%2Did"] = dataLossPreventionPolicyId
    }
    return NewItemSitesItemInformationProtectionDataLossPreventionPoliciesDataLossPreventionPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderInternal instantiates a new ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) {
    m := &ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/dataLossPreventionPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder instantiates a new ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemInformationProtectionDataLossPreventionPoliciesCountRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) Count()(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesCountRequestBuilder) {
    return NewItemSitesItemInformationProtectionDataLossPreventionPoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Evaluate provides operations to call the evaluate method.
// returns a *ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) Evaluate()(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilder) {
    return NewItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get dataLossPreventionPolicies from groups
// returns a DataLossPreventionPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataLossPreventionPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataLossPreventionPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataLossPreventionPolicyCollectionResponseable), nil
}
// Post create new navigation property to dataLossPreventionPolicies for groups
// returns a DataLossPreventionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataLossPreventionPolicyable, requestConfiguration *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataLossPreventionPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataLossPreventionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataLossPreventionPolicyable), nil
}
// ToGetRequestInformation get dataLossPreventionPolicies from groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to dataLossPreventionPolicies for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataLossPreventionPolicyable, requestConfiguration *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/dataLossPreventionPolicies", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder) {
    return NewItemSitesItemInformationProtectionDataLossPreventionPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
