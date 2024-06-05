package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreBackupRestoreRequestBuilder provides operations to manage the backupRestore property of the microsoft.graph.solutionsRoot entity.
type BackuprestoreBackupRestoreRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreBackupRestoreRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreBackupRestoreRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackuprestoreBackupRestoreRequestBuilderGetQueryParameters get the serviceStatus of the Microsoft 365 Backup Storage service in a tenant.
type BackuprestoreBackupRestoreRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackuprestoreBackupRestoreRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreBackupRestoreRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreBackupRestoreRequestBuilderGetQueryParameters
}
// BackuprestoreBackupRestoreRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreBackupRestoreRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackuprestoreBackupRestoreRequestBuilderInternal instantiates a new BackuprestoreBackupRestoreRequestBuilder and sets the default values.
func NewBackuprestoreBackupRestoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreBackupRestoreRequestBuilder) {
    m := &BackuprestoreBackupRestoreRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackuprestoreBackupRestoreRequestBuilder instantiates a new BackuprestoreBackupRestoreRequestBuilder and sets the default values.
func NewBackuprestoreBackupRestoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreBackupRestoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreBackupRestoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property backupRestore for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreBackupRestoreRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackuprestoreBackupRestoreRequestBuilderDeleteRequestConfiguration)(error) {
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
// DriveInclusionRules provides operations to manage the driveInclusionRules property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) DriveInclusionRules()(*BackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilder) {
    return NewBackuprestoreDriveinclusionrulesDriveInclusionRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DriveProtectionUnits provides operations to manage the driveProtectionUnits property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreDriveprotectionunitsDriveProtectionUnitsRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) DriveProtectionUnits()(*BackuprestoreDriveprotectionunitsDriveProtectionUnitsRequestBuilder) {
    return NewBackuprestoreDriveprotectionunitsDriveProtectionUnitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Enable provides operations to call the enable method.
// returns a *BackuprestoreEnableRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) Enable()(*BackuprestoreEnableRequestBuilder) {
    return NewBackuprestoreEnableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExchangeProtectionPolicies provides operations to manage the exchangeProtectionPolicies property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreExchangeprotectionpoliciesExchangeProtectionPoliciesRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) ExchangeProtectionPolicies()(*BackuprestoreExchangeprotectionpoliciesExchangeProtectionPoliciesRequestBuilder) {
    return NewBackuprestoreExchangeprotectionpoliciesExchangeProtectionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExchangeRestoreSessions provides operations to manage the exchangeRestoreSessions property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreExchangerestoresessionsExchangeRestoreSessionsRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) ExchangeRestoreSessions()(*BackuprestoreExchangerestoresessionsExchangeRestoreSessionsRequestBuilder) {
    return NewBackuprestoreExchangerestoresessionsExchangeRestoreSessionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the serviceStatus of the Microsoft 365 Backup Storage service in a tenant.
// returns a BackupRestoreRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/backuprestoreroot-get?view=graph-rest-beta
func (m *BackuprestoreBackupRestoreRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreBackupRestoreRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BackupRestoreRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBackupRestoreRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BackupRestoreRootable), nil
}
// MailboxInclusionRules provides operations to manage the mailboxInclusionRules property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreMailboxinclusionrulesMailboxInclusionRulesRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) MailboxInclusionRules()(*BackuprestoreMailboxinclusionrulesMailboxInclusionRulesRequestBuilder) {
    return NewBackuprestoreMailboxinclusionrulesMailboxInclusionRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MailboxProtectionUnits provides operations to manage the mailboxProtectionUnits property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreMailboxprotectionunitsMailboxProtectionUnitsRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) MailboxProtectionUnits()(*BackuprestoreMailboxprotectionunitsMailboxProtectionUnitsRequestBuilder) {
    return NewBackuprestoreMailboxprotectionunitsMailboxProtectionUnitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OneDriveForBusinessProtectionPolicies provides operations to manage the oneDriveForBusinessProtectionPolicies property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) OneDriveForBusinessProtectionPolicies()(*BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OneDriveForBusinessRestoreSessions provides operations to manage the oneDriveForBusinessRestoreSessions property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) OneDriveForBusinessRestoreSessions()(*BackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessrestoresessionsOneDriveForBusinessRestoreSessionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property backupRestore in solutions
// returns a BackupRestoreRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreBackupRestoreRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BackupRestoreRootable, requestConfiguration *BackuprestoreBackupRestoreRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BackupRestoreRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBackupRestoreRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BackupRestoreRootable), nil
}
// ProtectionPolicies provides operations to manage the protectionPolicies property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreProtectionpoliciesProtectionPoliciesRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) ProtectionPolicies()(*BackuprestoreProtectionpoliciesProtectionPoliciesRequestBuilder) {
    return NewBackuprestoreProtectionpoliciesProtectionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ProtectionUnits provides operations to manage the protectionUnits property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreProtectionunitsProtectionUnitsRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) ProtectionUnits()(*BackuprestoreProtectionunitsProtectionUnitsRequestBuilder) {
    return NewBackuprestoreProtectionunitsProtectionUnitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RestorePoints provides operations to manage the restorePoints property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreRestorepointsRestorePointsRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) RestorePoints()(*BackuprestoreRestorepointsRestorePointsRequestBuilder) {
    return NewBackuprestoreRestorepointsRestorePointsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RestoreSessions provides operations to manage the restoreSessions property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreRestoresessionsRestoreSessionsRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) RestoreSessions()(*BackuprestoreRestoresessionsRestoreSessionsRequestBuilder) {
    return NewBackuprestoreRestoresessionsRestoreSessionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceApps provides operations to manage the serviceApps property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreServiceappsServiceAppsRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) ServiceApps()(*BackuprestoreServiceappsServiceAppsRequestBuilder) {
    return NewBackuprestoreServiceappsServiceAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharePointProtectionPolicies provides operations to manage the sharePointProtectionPolicies property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreSharepointprotectionpoliciesSharePointProtectionPoliciesRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) SharePointProtectionPolicies()(*BackuprestoreSharepointprotectionpoliciesSharePointProtectionPoliciesRequestBuilder) {
    return NewBackuprestoreSharepointprotectionpoliciesSharePointProtectionPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharePointRestoreSessions provides operations to manage the sharePointRestoreSessions property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) SharePointRestoreSessions()(*BackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilder) {
    return NewBackuprestoreSharepointrestoresessionsSharePointRestoreSessionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SiteInclusionRules provides operations to manage the siteInclusionRules property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) SiteInclusionRules()(*BackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilder) {
    return NewBackuprestoreSiteinclusionrulesSiteInclusionRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SiteProtectionUnits provides operations to manage the siteProtectionUnits property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) SiteProtectionUnits()(*BackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilder) {
    return NewBackuprestoreSiteprotectionunitsSiteProtectionUnitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property backupRestore for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreBackupRestoreRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the serviceStatus of the Microsoft 365 Backup Storage service in a tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreBackupRestoreRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property backupRestore in solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BackupRestoreRootable, requestConfiguration *BackuprestoreBackupRestoreRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreBackupRestoreRequestBuilder when successful
func (m *BackuprestoreBackupRestoreRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreBackupRestoreRequestBuilder) {
    return NewBackuprestoreBackupRestoreRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
