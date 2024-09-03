package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackupRestoreProtectionPoliciesItemActivateRequestBuilder provides operations to call the activate method.
type BackupRestoreProtectionPoliciesItemActivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreProtectionPoliciesItemActivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreProtectionPoliciesItemActivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBackupRestoreProtectionPoliciesItemActivateRequestBuilderInternal instantiates a new BackupRestoreProtectionPoliciesItemActivateRequestBuilder and sets the default values.
func NewBackupRestoreProtectionPoliciesItemActivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreProtectionPoliciesItemActivateRequestBuilder) {
    m := &BackupRestoreProtectionPoliciesItemActivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/protectionPolicies/{protectionPolicyBase%2Did}/activate", pathParameters),
    }
    return m
}
// NewBackupRestoreProtectionPoliciesItemActivateRequestBuilder instantiates a new BackupRestoreProtectionPoliciesItemActivateRequestBuilder and sets the default values.
func NewBackupRestoreProtectionPoliciesItemActivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreProtectionPoliciesItemActivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreProtectionPoliciesItemActivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post activate a protectionPolicyBase. Currently, only one active backup policy per underlying service is supported (that is, one for OneDrive accounts, one for SharePoint sites, and one for Exchange Online users). You can add or remove artifacts (sites or user accounts) to or from each active policy.
// returns a ProtectionPolicyBaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/protectionpolicybase-activate?view=graph-rest-beta
func (m *BackupRestoreProtectionPoliciesItemActivateRequestBuilder) Post(ctx context.Context, requestConfiguration *BackupRestoreProtectionPoliciesItemActivateRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ProtectionPolicyBaseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation activate a protectionPolicyBase. Currently, only one active backup policy per underlying service is supported (that is, one for OneDrive accounts, one for SharePoint sites, and one for Exchange Online users). You can add or remove artifacts (sites or user accounts) to or from each active policy.
// returns a *RequestInformation when successful
func (m *BackupRestoreProtectionPoliciesItemActivateRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreProtectionPoliciesItemActivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *BackupRestoreProtectionPoliciesItemActivateRequestBuilder when successful
func (m *BackupRestoreProtectionPoliciesItemActivateRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreProtectionPoliciesItemActivateRequestBuilder) {
    return NewBackupRestoreProtectionPoliciesItemActivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
