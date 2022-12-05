package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetQueryParameters specifies the assets to include in the audience.
type AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetQueryParameters
}
// AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddMembers provides operations to call the addMembers method.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) AddMembers()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddMembersById provides operations to call the addMembersById method.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) AddMembersById()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersByIdRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderInternal instantiates a new UpdatableAssetItemRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) {
    m := &AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/members/{updatableAsset%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder instantiates a new UpdatableAssetItemRequestBuilder and sets the default values.
func NewAdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property members for admin
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation specifies the assets to include in the audience.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property members in admin
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property members for admin
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get specifies the assets to include in the audience.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable), nil
}
// Patch update the navigation property members in admin
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable), nil
}
// RemoveMembers provides operations to call the removeMembers method.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) RemoveMembers()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveMembersById provides operations to call the removeMembersById method.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) RemoveMembersById()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersByIdRequestBuilder) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
