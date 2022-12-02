package drive

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DriveRequestBuilder provides operations to manage the drive singleton.
type DriveRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DriveRequestBuilderGetQueryParameters get drive
type DriveRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DriveRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DriveRequestBuilderGetQueryParameters
}
// DriveRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) Activities()(*DriveActivitiesRequestBuilder) {
    return NewDriveActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) ActivitiesById(id string)(*DriveActivitiesItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return NewDriveActivitiesItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Bundles provides operations to manage the bundles property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) Bundles()(*DriveBundlesRequestBuilder) {
    return NewDriveBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById provides operations to manage the bundles property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) BundlesById(id string)(*DriveBundlesDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDriveBundlesDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveRequestBuilderInternal instantiates a new DriveRequestBuilder and sets the default values.
func NewDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveRequestBuilder) {
    m := &DriveRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveRequestBuilder instantiates a new DriveRequestBuilder and sets the default values.
func NewDriveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get drive
func (m *DriveRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DriveRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update drive
func (m *DriveRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *DriveRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Following provides operations to manage the following property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) Following()(*DriveFollowingRequestBuilder) {
    return NewDriveFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById provides operations to manage the following property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) FollowingById(id string)(*DriveFollowingDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDriveFollowingDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get drive
func (m *DriveRequestBuilder) Get(ctx context.Context, requestConfiguration *DriveRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
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
func (m *DriveRequestBuilder) Items()(*DriveItemsRequestBuilder) {
    return NewDriveItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById provides operations to manage the items property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) ItemsById(id string)(*DriveItemsDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDriveItemsDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List provides operations to manage the list property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) List()(*DriveListRequestBuilder) {
    return NewDriveListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update drive
func (m *DriveRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *DriveRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
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
func (m *DriveRequestBuilder) Recent()(*DriveRecentRequestBuilder) {
    return NewDriveRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Root provides operations to manage the root property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) Root()(*DriveRootRequestBuilder) {
    return NewDriveRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveRequestBuilder) SearchWithQ(q *string)(*DriveSearchWithQRequestBuilder) {
    return NewDriveSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe provides operations to call the sharedWithMe method.
func (m *DriveRequestBuilder) SharedWithMe()(*DriveSharedWithMeRequestBuilder) {
    return NewDriveSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Special provides operations to manage the special property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) Special()(*DriveSpecialRequestBuilder) {
    return NewDriveSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById provides operations to manage the special property of the microsoft.graph.drive entity.
func (m *DriveRequestBuilder) SpecialById(id string)(*DriveSpecialDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDriveSpecialDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
