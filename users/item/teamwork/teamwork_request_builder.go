package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i06112e4408d86eb023403d868493f6d21a2dcf420e7f41d3a057ee729c764019 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/teamwork/installedapps"
    i0c24dbd98f4f28c3ef8c12b843a0024c5d3e14adc15bf422622902bf94f67c69 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/teamwork/sendactivitynotification"
    i2ef02a19144871716098da11db760923e9baa87e04ae919f86982b61f015bfe0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/teamwork/associatedteams"
    ia3d4ee14dbf90a5db4cc41e8be01d9bc04a43f710d0ea8d9b5dc0cea55539a31 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/teamwork/installedapps/item"
    ibe8b9c72788a00966b19c87cdf9abefc88338cebbaf0ba7c1155316660033e0c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/teamwork/associatedteams/item"
)

// TeamworkRequestBuilder provides operations to manage the teamwork property of the microsoft.graph.user entity.
type TeamworkRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamworkRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamworkRequestBuilderGetQueryParameters a container for Microsoft Teams features available for the user. Read-only. Nullable.
type TeamworkRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamworkRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamworkRequestBuilderGetQueryParameters
}
// TeamworkRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamworkRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssociatedTeams provides operations to manage the associatedTeams property of the microsoft.graph.userTeamwork entity.
func (m *TeamworkRequestBuilder) AssociatedTeams()(*i2ef02a19144871716098da11db760923e9baa87e04ae919f86982b61f015bfe0.AssociatedTeamsRequestBuilder) {
    return i2ef02a19144871716098da11db760923e9baa87e04ae919f86982b61f015bfe0.NewAssociatedTeamsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssociatedTeamsById provides operations to manage the associatedTeams property of the microsoft.graph.userTeamwork entity.
func (m *TeamworkRequestBuilder) AssociatedTeamsById(id string)(*ibe8b9c72788a00966b19c87cdf9abefc88338cebbaf0ba7c1155316660033e0c.AssociatedTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["associatedTeamInfo%2Did"] = id
    }
    return ibe8b9c72788a00966b19c87cdf9abefc88338cebbaf0ba7c1155316660033e0c.NewAssociatedTeamInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTeamworkRequestBuilderInternal instantiates a new TeamworkRequestBuilder and sets the default values.
func NewTeamworkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkRequestBuilder) {
    m := &TeamworkRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/teamwork{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamworkRequestBuilder instantiates a new TeamworkRequestBuilder and sets the default values.
func NewTeamworkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamworkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamworkRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property teamwork for users
func (m *TeamworkRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamworkRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *TeamworkRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamworkRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property teamwork in users
func (m *TeamworkRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable, requestConfiguration *TeamworkRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property teamwork for users
func (m *TeamworkRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamworkRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a container for Microsoft Teams features available for the user. Read-only. Nullable.
func (m *TeamworkRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamworkRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserTeamworkFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable), nil
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.userTeamwork entity.
func (m *TeamworkRequestBuilder) InstalledApps()(*i06112e4408d86eb023403d868493f6d21a2dcf420e7f41d3a057ee729c764019.InstalledAppsRequestBuilder) {
    return i06112e4408d86eb023403d868493f6d21a2dcf420e7f41d3a057ee729c764019.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.userTeamwork entity.
func (m *TeamworkRequestBuilder) InstalledAppsById(id string)(*ia3d4ee14dbf90a5db4cc41e8be01d9bc04a43f710d0ea8d9b5dc0cea55539a31.UserScopeTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userScopeTeamsAppInstallation%2Did"] = id
    }
    return ia3d4ee14dbf90a5db4cc41e8be01d9bc04a43f710d0ea8d9b5dc0cea55539a31.NewUserScopeTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property teamwork in users
func (m *TeamworkRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable, requestConfiguration *TeamworkRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserTeamworkFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserTeamworkable), nil
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
func (m *TeamworkRequestBuilder) SendActivityNotification()(*i0c24dbd98f4f28c3ef8c12b843a0024c5d3e14adc15bf422622902bf94f67c69.SendActivityNotificationRequestBuilder) {
    return i0c24dbd98f4f28c3ef8c12b843a0024c5d3e14adc15bf422622902bf94f67c69.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
