package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemDrivesDriveItemRequestBuilder provides operations to manage the drives property of the microsoft.graph.user entity.
type UsersItemDrivesDriveItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemDrivesDriveItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDrivesDriveItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemDrivesDriveItemRequestBuilderGetQueryParameters a collection of drives available for this user. Read-only.
type UsersItemDrivesDriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemDrivesDriveItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDrivesDriveItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemDrivesDriveItemRequestBuilderGetQueryParameters
}
// UsersItemDrivesDriveItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDrivesDriveItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) Activities()(*UsersItemDrivesItemActivitiesRequestBuilder) {
    return NewUsersItemDrivesItemActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) ActivitiesById(id string)(*UsersItemDrivesItemActivitiesItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return NewUsersItemDrivesItemActivitiesItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Bundles provides operations to manage the bundles property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) Bundles()(*UsersItemDrivesItemBundlesRequestBuilder) {
    return NewUsersItemDrivesItemBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById provides operations to manage the bundles property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) BundlesById(id string)(*UsersItemDrivesItemBundlesDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewUsersItemDrivesItemBundlesDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUsersItemDrivesDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewUsersItemDrivesDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDrivesDriveItemRequestBuilder) {
    m := &UsersItemDrivesDriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemDrivesDriveItemRequestBuilder instantiates a new DriveItemRequestBuilder and sets the default values.
func NewUsersItemDrivesDriveItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDrivesDriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemDrivesDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property drives for users
func (m *UsersItemDrivesDriveItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemDrivesDriveItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of drives available for this user. Read-only.
func (m *UsersItemDrivesDriveItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemDrivesDriveItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property drives in users
func (m *UsersItemDrivesDriveItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *UsersItemDrivesDriveItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property drives for users
func (m *UsersItemDrivesDriveItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemDrivesDriveItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Following provides operations to manage the following property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) Following()(*UsersItemDrivesItemFollowingRequestBuilder) {
    return NewUsersItemDrivesItemFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById provides operations to manage the following property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) FollowingById(id string)(*UsersItemDrivesItemFollowingDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewUsersItemDrivesItemFollowingDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get a collection of drives available for this user. Read-only.
func (m *UsersItemDrivesDriveItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemDrivesDriveItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) Items()(*UsersItemDrivesItemItemsRequestBuilder) {
    return NewUsersItemDrivesItemItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) ItemsById(id string)(*UsersItemDrivesItemItemsDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewUsersItemDrivesItemItemsDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List provides operations to manage the list property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) List()(*UsersItemDrivesItemListRequestBuilder) {
    return NewUsersItemDrivesItemListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property drives in users
func (m *UsersItemDrivesDriveItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *UsersItemDrivesDriveItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable), nil
}
// Recent provides operations to call the recent method.
func (m *UsersItemDrivesDriveItemRequestBuilder) Recent()(*UsersItemDrivesItemRecentRequestBuilder) {
    return NewUsersItemDrivesItemRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Root provides operations to manage the root property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) Root()(*UsersItemDrivesItemRootRequestBuilder) {
    return NewUsersItemDrivesItemRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *UsersItemDrivesDriveItemRequestBuilder) SearchWithQ(q *string)(*UsersItemDrivesItemSearchWithQRequestBuilder) {
    return NewUsersItemDrivesItemSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe provides operations to call the sharedWithMe method.
func (m *UsersItemDrivesDriveItemRequestBuilder) SharedWithMe()(*UsersItemDrivesItemSharedWithMeRequestBuilder) {
    return NewUsersItemDrivesItemSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Special provides operations to manage the special property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) Special()(*UsersItemDrivesItemSpecialRequestBuilder) {
    return NewUsersItemDrivesItemSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById provides operations to manage the special property of the microsoft.graph.drive entity.
func (m *UsersItemDrivesDriveItemRequestBuilder) SpecialById(id string)(*UsersItemDrivesItemSpecialDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewUsersItemDrivesItemSpecialDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
