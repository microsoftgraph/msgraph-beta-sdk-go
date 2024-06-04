package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder provides operations to manage the oneDriveForBusinessProtectionPolicies property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderGetQueryParameters the list of OneDrive for Business protection policies in the tenant.
type BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderGetQueryParameters
}
// BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderInternal instantiates a new BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) {
    m := &BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/oneDriveForBusinessProtectionPolicies/{oneDriveForBusinessProtectionPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder instantiates a new BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property oneDriveForBusinessProtectionPolicies for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DriveInclusionRules provides operations to manage the driveInclusionRules property of the microsoft.graph.oneDriveForBusinessProtectionPolicy entity.
// returns a *BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveInclusionRulesRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) DriveInclusionRules()(*BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveInclusionRulesRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveInclusionRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DriveProtectionUnits provides operations to manage the driveProtectionUnits property of the microsoft.graph.oneDriveForBusinessProtectionPolicy entity.
// returns a *BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveprotectionunitsDriveProtectionUnitsRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) DriveProtectionUnits()(*BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveprotectionunitsDriveProtectionUnitsRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveprotectionunitsDriveProtectionUnitsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of OneDrive for Business protection policies in the tenant.
// returns a OneDriveForBusinessProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessProtectionPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOneDriveForBusinessProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessProtectionPolicyable), nil
}
// Patch update the protection policy for the OneDrive service in Microsoft 365. This method adds a driveProtectionUnit to or removes it from a oneDriveForBusinessProtectionPolicy object.
// returns a OneDriveForBusinessProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onedriveforbusinessprotectionpolicy-update?view=graph-rest-beta
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessProtectionPolicyable, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessProtectionPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOneDriveForBusinessProtectionPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessProtectionPolicyable), nil
}
// ToDeleteRequestInformation delete navigation property oneDriveForBusinessProtectionPolicies for solutions
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of OneDrive for Business protection policies in the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the protection policy for the OneDrive service in Microsoft 365. This method adds a driveProtectionUnit to or removes it from a oneDriveForBusinessProtectionPolicy object.
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessProtectionPolicyable, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
