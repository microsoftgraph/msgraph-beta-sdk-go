package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder provides operations to manage the permissionGrantPreApprovalPolicies property of the microsoft.graph.servicePrincipal entity.
type ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilderGetQueryParameters retrieve the permissionGrantPreApprovalPolicy object for the servicePrincipal.
type ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilderGetQueryParameters struct {
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
// ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilderGetQueryParameters
}
// ByPermissionGrantPreApprovalPolicyId provides operations to manage the permissionGrantPreApprovalPolicies property of the microsoft.graph.servicePrincipal entity.
// returns a *ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder when successful
func (m *ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder) ByPermissionGrantPreApprovalPolicyId(permissionGrantPreApprovalPolicyId string)(*ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if permissionGrantPreApprovalPolicyId != "" {
        urlTplParams["permissionGrantPreApprovalPolicy%2Did"] = permissionGrantPreApprovalPolicyId
    }
    return NewItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilderInternal instantiates a new ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder and sets the default values.
func NewItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder) {
    m := &ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/permissionGrantPreApprovalPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder instantiates a new ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder and sets the default values.
func NewItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemPermissiongrantpreapprovalpoliciesCountRequestBuilder when successful
func (m *ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder) Count()(*ItemPermissiongrantpreapprovalpoliciesCountRequestBuilder) {
    return NewItemPermissiongrantpreapprovalpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the permissionGrantPreApprovalPolicy object for the servicePrincipal.
// returns a PermissionGrantPreApprovalPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPreApprovalPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionGrantPreApprovalPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPreApprovalPolicyCollectionResponseable), nil
}
// ToGetRequestInformation retrieve the permissionGrantPreApprovalPolicy object for the servicePrincipal.
// returns a *RequestInformation when successful
func (m *ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder when successful
func (m *ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder) WithUrl(rawUrl string)(*ItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder) {
    return NewItemPermissiongrantpreapprovalpoliciesPermissionGrantPreApprovalPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
