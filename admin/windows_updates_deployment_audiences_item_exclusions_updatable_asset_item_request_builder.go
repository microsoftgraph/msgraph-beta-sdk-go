package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder provides operations to manage the exclusions property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderGetQueryParameters specifies the assets to exclude from the audience.
type WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderGetQueryParameters
}
// WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderInternal instantiates a new UpdatableAssetItemRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) {
    m := &WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deploymentAudiences/{deploymentAudience%2Did}/exclusions/{updatableAsset%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder instantiates a new UpdatableAssetItemRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property exclusions for admin
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get specifies the assets to exclude from the audience.
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable), nil
}
// Patch update the navigation property exclusions in admin
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable), nil
}
// ToDeleteRequestInformation delete navigation property exclusions for admin
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation specifies the assets to exclude from the audience.
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update the navigation property exclusions in admin
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WindowsUpdatesAddMembers provides operations to call the addMembers method.
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) WindowsUpdatesAddMembers()(*WindowsUpdatesDeploymentAudiencesItemExclusionsItemWindowsUpdatesAddMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemWindowsUpdatesAddMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsUpdatesAddMembersById provides operations to call the addMembersById method.
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) WindowsUpdatesAddMembersById()(*WindowsUpdatesDeploymentAudiencesItemExclusionsItemWindowsUpdatesAddMembersByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemWindowsUpdatesAddMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsUpdatesRemoveMembers provides operations to call the removeMembers method.
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) WindowsUpdatesRemoveMembers()(*WindowsUpdatesDeploymentAudiencesItemExclusionsItemWindowsUpdatesRemoveMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemWindowsUpdatesRemoveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsUpdatesRemoveMembersById provides operations to call the removeMembersById method.
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsUpdatableAssetItemRequestBuilder) WindowsUpdatesRemoveMembersById()(*WindowsUpdatesDeploymentAudiencesItemExclusionsItemWindowsUpdatesRemoveMembersByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemWindowsUpdatesRemoveMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
