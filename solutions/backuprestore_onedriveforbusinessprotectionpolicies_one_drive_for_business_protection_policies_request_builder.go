package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder provides operations to manage the oneDriveForBusinessProtectionPolicies property of the microsoft.graph.backupRestoreRoot entity.
type BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderGetQueryParameters the list of OneDrive for Business protection policies in the tenant.
type BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderGetQueryParameters struct {
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
// BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderGetQueryParameters
}
// BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByOneDriveForBusinessProtectionPolicyId provides operations to manage the oneDriveForBusinessProtectionPolicies property of the microsoft.graph.backupRestoreRoot entity.
// returns a *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder) ByOneDriveForBusinessProtectionPolicyId(oneDriveForBusinessProtectionPolicyId string)(*BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if oneDriveForBusinessProtectionPolicyId != "" {
        urlTplParams["oneDriveForBusinessProtectionPolicy%2Did"] = oneDriveForBusinessProtectionPolicyId
    }
    return NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderInternal instantiates a new BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder) {
    m := &BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/oneDriveForBusinessProtectionPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder instantiates a new BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder and sets the default values.
func NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BackuprestoreOnedriveforbusinessprotectionpoliciesCountRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder) Count()(*BackuprestoreOnedriveforbusinessprotectionpoliciesCountRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessprotectionpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of OneDrive for Business protection policies in the tenant.
// returns a OneDriveForBusinessProtectionPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessProtectionPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOneDriveForBusinessProtectionPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessProtectionPolicyCollectionResponseable), nil
}
// Post create a protection policy for the OneDrive service in Microsoft 365. When the policy is created, its state is set to inactive. Users can also provide a list of protection units under the policy.
// returns a OneDriveForBusinessProtectionPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/backuprestoreroot-post-onedriveforbusinessprotectionpolicies?view=graph-rest-beta
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessProtectionPolicyable, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessProtectionPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the list of OneDrive for Business protection policies in the tenant.
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a protection policy for the OneDrive service in Microsoft 365. When the policy is created, its state is set to inactive. Users can also provide a list of protection units under the policy.
// returns a *RequestInformation when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OneDriveForBusinessProtectionPolicyable, requestConfiguration *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder when successful
func (m *BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder) WithUrl(rawUrl string)(*BackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder) {
    return NewBackuprestoreOnedriveforbusinessprotectionpoliciesOneDriveForBusinessProtectionPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
