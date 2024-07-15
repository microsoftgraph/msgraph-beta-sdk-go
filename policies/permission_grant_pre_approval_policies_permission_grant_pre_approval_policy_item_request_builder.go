package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder provides operations to manage the permissionGrantPreApprovalPolicies property of the microsoft.graph.policyRoot entity.
type PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetQueryParameters read the properties and relationships of a permissionGrantPreApprovalPolicy object.
type PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetQueryParameters
}
// PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderInternal instantiates a new PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder and sets the default values.
func NewPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) {
    m := &PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/permissionGrantPreApprovalPolicies/{permissionGrantPreApprovalPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder instantiates a new PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder and sets the default values.
func NewPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a permissionGrantPreApprovalPolicy object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissiongrantpreapprovalpolicy-delete?view=graph-rest-beta
func (m *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a permissionGrantPreApprovalPolicy object.
// returns a PermissionGrantPreApprovalPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissiongrantpreapprovalpolicy-get?view=graph-rest-beta
func (m *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPreApprovalPolicyable, error) {
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
// Patch update the properties of a permissionGrantPreApprovalPolicy object.
// returns a PermissionGrantPreApprovalPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissiongrantpreapprovalpolicy-update?view=graph-rest-beta
func (m *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPreApprovalPolicyable, requestConfiguration *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPreApprovalPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete a permissionGrantPreApprovalPolicy object.
// returns a *RequestInformation when successful
func (m *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a permissionGrantPreApprovalPolicy object.
// returns a *RequestInformation when successful
func (m *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a permissionGrantPreApprovalPolicy object.
// returns a *RequestInformation when successful
func (m *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionGrantPreApprovalPolicyable, requestConfiguration *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder when successful
func (m *PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) WithUrl(rawUrl string)(*PermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder) {
    return NewPermissionGrantPreApprovalPoliciesPermissionGrantPreApprovalPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
