// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder provides operations to manage the sharePointProtectionPolicies property of the microsoft.graph.backupRestoreRoot entity.
type BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderGetQueryParameters the list of SharePoint protection policies in the tenant.
type BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderGetQueryParameters
}
// BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderInternal instantiates a new BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder and sets the default values.
func NewBackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) {
    m := &BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/sharePointProtectionPolicies/{sharePointProtectionPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder instantiates a new BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder and sets the default values.
func NewBackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sharePointProtectionPolicies for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of SharePoint protection policies in the tenant.
// returns a SharePointProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointProtectionPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSharePointProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointProtectionPolicyable), nil
}
// Patch update a SharePoint protection policy. This method adds a siteprotectionunit to or removes it from the protection policy.
// returns a SharePointProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/sharepointprotectionpolicy-update?view=graph-rest-beta
func (m *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointProtectionPolicyable, requestConfiguration *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointProtectionPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSharePointProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointProtectionPolicyable), nil
}
// SiteInclusionRules provides operations to manage the siteInclusionRules property of the microsoft.graph.sharePointProtectionPolicy entity.
// returns a *BackupRestoreSharePointProtectionPoliciesItemSiteInclusionRulesRequestBuilder when successful
func (m *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) SiteInclusionRules()(*BackupRestoreSharePointProtectionPoliciesItemSiteInclusionRulesRequestBuilder) {
    return NewBackupRestoreSharePointProtectionPoliciesItemSiteInclusionRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SiteProtectionUnits provides operations to manage the siteProtectionUnits property of the microsoft.graph.sharePointProtectionPolicy entity.
// returns a *BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsRequestBuilder when successful
func (m *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) SiteProtectionUnits()(*BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsRequestBuilder) {
    return NewBackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SiteProtectionUnitsBulkAdditionJobs provides operations to manage the siteProtectionUnitsBulkAdditionJobs property of the microsoft.graph.sharePointProtectionPolicy entity.
// returns a *BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsRequestBuilder when successful
func (m *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) SiteProtectionUnitsBulkAdditionJobs()(*BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsRequestBuilder) {
    return NewBackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property sharePointProtectionPolicies for solutions
// returns a *RequestInformation when successful
func (m *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of SharePoint protection policies in the tenant.
// returns a *RequestInformation when successful
func (m *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update a SharePoint protection policy. This method adds a siteprotectionunit to or removes it from the protection policy.
// returns a *RequestInformation when successful
func (m *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharePointProtectionPolicyable, requestConfiguration *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder when successful
func (m *BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder) {
    return NewBackupRestoreSharePointProtectionPoliciesSharePointProtectionPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
