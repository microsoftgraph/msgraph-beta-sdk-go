package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder provides operations to manage the multiTenantOrganizationIdentitySynchronization property of the microsoft.graph.policyTemplate entity.
type CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderGetQueryParameters get the cross-tenant access policy template with user synchronization settings for a multi-tenant organization.
type CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderGetQueryParameters
}
// CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderInternal instantiates a new MultiTenantOrganizationIdentitySynchronizationRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder) {
    m := &CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/crossTenantAccessPolicy/templates/multiTenantOrganizationIdentitySynchronization{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder instantiates a new MultiTenantOrganizationIdentitySynchronizationRequestBuilder and sets the default values.
func NewCrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property multiTenantOrganizationIdentitySynchronization for policies
func (m *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder) Delete(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get the cross-tenant access policy template with user synchronization settings for a multi-tenant organization.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganizationidentitysyncpolicytemplate-get?view=graph-rest-1.0
func (m *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder) Get(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationIdentitySyncPolicyTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMultiTenantOrganizationIdentitySyncPolicyTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationIdentitySyncPolicyTemplateable), nil
}
// Patch update the cross-tenant access policy template with user synchronization settings for a multi-tenant organization.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganizationidentitysyncpolicytemplate-update?view=graph-rest-1.0
func (m *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationIdentitySyncPolicyTemplateable, requestConfiguration *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationIdentitySyncPolicyTemplateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMultiTenantOrganizationIdentitySyncPolicyTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationIdentitySyncPolicyTemplateable), nil
}
// ResetToDefaultSettings provides operations to call the resetToDefaultSettings method.
func (m *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder) ResetToDefaultSettings()(*CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationResetToDefaultSettingsRequestBuilder) {
    return NewCrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationResetToDefaultSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property multiTenantOrganizationIdentitySynchronization for policies
func (m *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get the cross-tenant access policy template with user synchronization settings for a multi-tenant organization.
func (m *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the cross-tenant access policy template with user synchronization settings for a multi-tenant organization.
func (m *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationIdentitySyncPolicyTemplateable, requestConfiguration *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder) WithUrl(rawUrl string)(*CrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder) {
    return NewCrossTenantAccessPolicyTemplatesMultiTenantOrganizationIdentitySynchronizationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
