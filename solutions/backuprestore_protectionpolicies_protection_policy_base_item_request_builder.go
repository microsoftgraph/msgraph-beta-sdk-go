package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder provides operations to manage the protectionPolicies property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderGetQueryParameters list of protection policies in the tenant.
type BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderGetQueryParameters
}
// BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activate provides operations to call the activate method.
// returns a *BackuprestoreProtectionpoliciesItemActivateRequestBuilder when successful
func (m *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) Activate()(*BackuprestoreProtectionpoliciesItemActivateRequestBuilder) {
    return NewBackuprestoreProtectionpoliciesItemActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderInternal instantiates a new BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder and sets the default values.
func NewBackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) {
    m := &BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/protectionPolicies/{protectionPolicyBase%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder instantiates a new BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder and sets the default values.
func NewBackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Deactivate provides operations to call the deactivate method.
// returns a *BackuprestoreProtectionpoliciesItemDeactivateRequestBuilder when successful
func (m *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) Deactivate()(*BackuprestoreProtectionpoliciesItemDeactivateRequestBuilder) {
    return NewBackuprestoreProtectionpoliciesItemDeactivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a protection policy. Read the properties and relationships of a protectionPolicyBase object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/protectionpolicybase-delete?view=graph-rest-beta
func (m *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of protection policies in the tenant.
// returns a ProtectionPolicyBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProtectionPolicyBaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProtectionPolicyBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProtectionPolicyBaseable), nil
}
// Patch update the navigation property protectionPolicies in solutions
// returns a ProtectionPolicyBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProtectionPolicyBaseable, requestConfiguration *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProtectionPolicyBaseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateProtectionPolicyBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProtectionPolicyBaseable), nil
}
// ToDeleteRequestInformation delete a protection policy. Read the properties and relationships of a protectionPolicyBase object.
// returns a *RequestInformation when successful
func (m *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of protection policies in the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property protectionPolicies in solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProtectionPolicyBaseable, requestConfiguration *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder when successful
func (m *BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder) {
    return NewBackuprestoreProtectionpoliciesProtectionPolicyBaseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
