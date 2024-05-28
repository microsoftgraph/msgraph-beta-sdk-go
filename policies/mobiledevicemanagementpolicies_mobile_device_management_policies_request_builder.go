package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder provides operations to manage the mobileDeviceManagementPolicies property of the microsoft.graph.policyRoot entity.
type MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderGetQueryParameters get a list of the mobilityManagementPolicy objects and their properties.
type MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderGetQueryParameters struct {
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
// MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderGetQueryParameters
}
// MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMobilityManagementPolicyId provides operations to manage the mobileDeviceManagementPolicies property of the microsoft.graph.policyRoot entity.
// returns a *MobiledevicemanagementpoliciesMobilityManagementPolicyItemRequestBuilder when successful
func (m *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder) ByMobilityManagementPolicyId(mobilityManagementPolicyId string)(*MobiledevicemanagementpoliciesMobilityManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mobilityManagementPolicyId != "" {
        urlTplParams["mobilityManagementPolicy%2Did"] = mobilityManagementPolicyId
    }
    return NewMobiledevicemanagementpoliciesMobilityManagementPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderInternal instantiates a new MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder and sets the default values.
func NewMobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder) {
    m := &MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/mobileDeviceManagementPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder instantiates a new MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder and sets the default values.
func NewMobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobiledevicemanagementpoliciesCountRequestBuilder when successful
func (m *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder) Count()(*MobiledevicemanagementpoliciesCountRequestBuilder) {
    return NewMobiledevicemanagementpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the mobilityManagementPolicy objects and their properties.
// returns a MobilityManagementPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mobiledevicemanagementpolicies-list?view=graph-rest-beta
func (m *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobilityManagementPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobilityManagementPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobilityManagementPolicyCollectionResponseable), nil
}
// Post create new navigation property to mobileDeviceManagementPolicies for policies
// returns a MobilityManagementPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobilityManagementPolicyable, requestConfiguration *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobilityManagementPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobilityManagementPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobilityManagementPolicyable), nil
}
// ToGetRequestInformation get a list of the mobilityManagementPolicy objects and their properties.
// returns a *RequestInformation when successful
func (m *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to mobileDeviceManagementPolicies for policies
// returns a *RequestInformation when successful
func (m *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobilityManagementPolicyable, requestConfiguration *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder when successful
func (m *MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder) WithUrl(rawUrl string)(*MobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder) {
    return NewMobiledevicemanagementpoliciesMobileDeviceManagementPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
