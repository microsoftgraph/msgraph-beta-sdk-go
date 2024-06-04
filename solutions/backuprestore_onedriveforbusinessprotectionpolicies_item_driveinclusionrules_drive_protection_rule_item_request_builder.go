package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder provides operations to manage the driveInclusionRules property of the microsoft.graph.oneDriveForBusinessProtectionPolicy entity.
type BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetQueryParameters get a protection rule that's associated with a protection policy. You can use this operation to get mailbox, drive, and site protection rules. An inclusion rule indicates that a protection policy should contain protection units that match the specified rule criteria. The initial status of a protection rule upon creation is active. After the rule is applied, the state is either completed or completedWithErrors.
type BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetQueryParameters
}
// NewBackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilderInternal instantiates a new BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) {
    m := &BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/oneDriveForBusinessProtectionPolicies/{oneDriveForBusinessProtectionPolicy%2Did}/driveInclusionRules/{driveProtectionRule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder instantiates a new BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a protection rule that's associated with a protection policy. You can use this operation to get mailbox, drive, and site protection rules. An inclusion rule indicates that a protection policy should contain protection units that match the specified rule criteria. The initial status of a protection rule upon creation is active. After the rule is applied, the state is either completed or completedWithErrors.
// returns a DriveProtectionRuleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/protectionrulebase-get?view=graph-rest-beta
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveProtectionRuleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveProtectionRuleable), nil
}
// ToGetRequestInformation get a protection rule that's associated with a protection policy. You can use this operation to get mailbox, drive, and site protection rules. An inclusion rule indicates that a protection policy should contain protection units that match the specified rule criteria. The initial status of a protection rule upon creation is active. After the rule is applied, the state is either completed or completedWithErrors.
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessprotectionpoliciesItemDriveinclusionrulesDriveProtectionRuleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
