package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilder provides operations to manage the driveItem property of the microsoft.graph.itemActivityOLD entity.
type UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilderGetQueryParameters get driveItem from users
type UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilderGetQueryParameters
}
// NewUsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewUsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilder) {
    m := &UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/drives/{drive%2Did}/items/{driveItem%2Did}/listItem/activities/{itemActivityOLD%2Did}/driveItem{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilder instantiates a new DriveItemRequestBuilder and sets the default values.
func NewUsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilder) Content()(*UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemContentRequestBuilder) {
    return NewUsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get driveItem from users
func (m *UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get get driveItem from users
func (m *UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemDrivesItemItemsItemListItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
