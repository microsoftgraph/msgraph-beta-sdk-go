package root

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0a759423e57f7ca082727db4b6d8a8898f512a525053fb3d63f0cefc721c1de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/children"
    i10967f3f015a7553858f3d7655160d79e82c92744e557898f8831df427a7f9d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/permissions"
    i4b907a46b4abff913d399e355773c295402e30948ecc48bcf4772e19463a9500 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/listitem"
    i88dce30fb62e2ab6e1303dbddc766fe68f945ca739b73a6d018b65303d9c5775 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/activities"
    iac5ca1ab5d5a17b0fdb527a4da0f9b513197ceb75ce9fcc85bff3e84289324fa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/subscriptions"
    ic2e3b15a1a3f5031cf2b93c3d125a1acb16a0d17f5a7a65ca4107e15a3b48d8c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/analytics"
    ic6c9cb759cb9b2c3f6222bf33b4b6f9d3f9296bb6826defd9993e0e12b9fba70 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/thumbnails"
    icefd7a6794e7c2da403f53d85781b693358d81ec7ec22ba7b8f4b25bbe37762a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/content"
    ie7091c6c63c7c2e80421ae52f1dfb3a98528a74fb87172e4ce6fb2b77b798df1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/versions"
    i3d1f0ab7aefec569aff57bf4dc231634efabab69a7619a43f36d4e2e890fe320 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/permissions/item"
    i716bf1df0c193166f142b93a995c91fe156031ae2c01e24f510c1f5f7e220bd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/children/item"
    i78e57eeeebb0af099ea095c21f51ff791b4ba57237dc7fa541b1801e0e801f2b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/activities/item"
    i810772176a264e9126a5f0431010acd740e4e5e4a9d033eb71a2a7adba995046 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/thumbnails/item"
    i84445af8dde775bdfa662f8cc61cb3171dba639fa77ac6a8dc56b5e2c999d63c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/versions/item"
    i85104fb17131a6a1f94c64c9c63e18daa3a300536f9fc44378087cb1a38a3332 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/root/subscriptions/item"
)

// RootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type RootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RootRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RootRequestBuilderGetQueryParameters the root folder of the drive. Read-only.
type RootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RootRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RootRequestBuilderGetQueryParameters
}
// RootRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RootRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities the activities property
func (m *RootRequestBuilder) Activities()(*i88dce30fb62e2ab6e1303dbddc766fe68f945ca739b73a6d018b65303d9c5775.ActivitiesRequestBuilder) {
    return i88dce30fb62e2ab6e1303dbddc766fe68f945ca739b73a6d018b65303d9c5775.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*i78e57eeeebb0af099ea095c21f51ff791b4ba57237dc7fa541b1801e0e801f2b.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i78e57eeeebb0af099ea095c21f51ff791b4ba57237dc7fa541b1801e0e801f2b.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*ic2e3b15a1a3f5031cf2b93c3d125a1acb16a0d17f5a7a65ca4107e15a3b48d8c.AnalyticsRequestBuilder) {
    return ic2e3b15a1a3f5031cf2b93c3d125a1acb16a0d17f5a7a65ca4107e15a3b48d8c.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*i0a759423e57f7ca082727db4b6d8a8898f512a525053fb3d63f0cefc721c1de3.ChildrenRequestBuilder) {
    return i0a759423e57f7ca082727db4b6d8a8898f512a525053fb3d63f0cefc721c1de3.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i716bf1df0c193166f142b93a995c91fe156031ae2c01e24f510c1f5f7e220bd0.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i716bf1df0c193166f142b93a995c91fe156031ae2c01e24f510c1f5f7e220bd0.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/drives/{drive%2Did}/root{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRootRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the content property
func (m *RootRequestBuilder) Content()(*icefd7a6794e7c2da403f53d85781b693358d81ec7ec22ba7b8f4b25bbe37762a.ContentRequestBuilder) {
    return icefd7a6794e7c2da403f53d85781b693358d81ec7ec22ba7b8f4b25bbe37762a.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for me
func (m *RootRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for me
func (m *RootRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RootRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property root in me
func (m *RootRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in me
func (m *RootRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property root for me
func (m *RootRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property root for me
func (m *RootRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *RootRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the root folder of the drive. Read-only.
func (m *RootRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the root folder of the drive. Read-only.
func (m *RootRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *RootRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// ListItem the listItem property
func (m *RootRequestBuilder) ListItem()(*i4b907a46b4abff913d399e355773c295402e30948ecc48bcf4772e19463a9500.ListItemRequestBuilder) {
    return i4b907a46b4abff913d399e355773c295402e30948ecc48bcf4772e19463a9500.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in me
func (m *RootRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property root in me
func (m *RootRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *RootRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Permissions the permissions property
func (m *RootRequestBuilder) Permissions()(*i10967f3f015a7553858f3d7655160d79e82c92744e557898f8831df427a7f9d0.PermissionsRequestBuilder) {
    return i10967f3f015a7553858f3d7655160d79e82c92744e557898f8831df427a7f9d0.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i3d1f0ab7aefec569aff57bf4dc231634efabab69a7619a43f36d4e2e890fe320.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i3d1f0ab7aefec569aff57bf4dc231634efabab69a7619a43f36d4e2e890fe320.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*iac5ca1ab5d5a17b0fdb527a4da0f9b513197ceb75ce9fcc85bff3e84289324fa.SubscriptionsRequestBuilder) {
    return iac5ca1ab5d5a17b0fdb527a4da0f9b513197ceb75ce9fcc85bff3e84289324fa.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i85104fb17131a6a1f94c64c9c63e18daa3a300536f9fc44378087cb1a38a3332.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i85104fb17131a6a1f94c64c9c63e18daa3a300536f9fc44378087cb1a38a3332.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*ic6c9cb759cb9b2c3f6222bf33b4b6f9d3f9296bb6826defd9993e0e12b9fba70.ThumbnailsRequestBuilder) {
    return ic6c9cb759cb9b2c3f6222bf33b4b6f9d3f9296bb6826defd9993e0e12b9fba70.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i810772176a264e9126a5f0431010acd740e4e5e4a9d033eb71a2a7adba995046.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i810772176a264e9126a5f0431010acd740e4e5e4a9d033eb71a2a7adba995046.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*ie7091c6c63c7c2e80421ae52f1dfb3a98528a74fb87172e4ce6fb2b77b798df1.VersionsRequestBuilder) {
    return ie7091c6c63c7c2e80421ae52f1dfb3a98528a74fb87172e4ce6fb2b77b798df1.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i84445af8dde775bdfa662f8cc61cb3171dba639fa77ac6a8dc56b5e2c999d63c.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return i84445af8dde775bdfa662f8cc61cb3171dba639fa77ac6a8dc56b5e2c999d63c.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
