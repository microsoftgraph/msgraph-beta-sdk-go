package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder provides operations to manage the permissionGrantPreApprovalPolicies property of the microsoft.graph.servicePrincipal entity.
type ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetQueryParameters retrieve the permissionGrantPreApprovalPolicy object for the servicePrincipal.
type ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetQueryParameters
}
// NewItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderInternal instantiates a new ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder and sets the default values.
func NewItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) {
    m := &ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/permissionGrantPreApprovalPolicies/{permissionGrantPreApprovalPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder instantiates a new ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder and sets the default values.
func NewItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieve the permissionGrantPreApprovalPolicy object for the servicePrincipal.
// returns a PermissionGrantPreApprovalPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-list-permissiongrantpreapprovalpolicies?view=graph-rest-beta
func (m *ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPreApprovalPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionGrantPreApprovalPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPreApprovalPolicyable), nil
}
// ToGetRequestInformation retrieve the permissionGrantPreApprovalPolicy object for the servicePrincipal.
// returns a *RequestInformation when successful
func (m *ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder when successful
func (m *ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) {
    return NewItemPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
