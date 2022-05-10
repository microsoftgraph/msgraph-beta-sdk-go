package root

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0906b36a4c6c7e8fa300e6712af3d1532d16969d436731907cbbcd07ec2c1303 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/versions"
    i1a2313e6ac103bc2f750a67ff7be5e8c0347553f28f2d941ac24c07416fcdb8b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/children"
    i4d302afe3c097d2397bea4b437d9ba2c45601731ba2b79d26ab3673a5da462df "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/subscriptions"
    i63baa5e7acb91aba07e2546fca385eaf695da1babdc82b21a9601c22d143e605 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/thumbnails"
    i65467600fd8d2b15564ff06f3e2178353d38b7a279223bcb861cd765e98c081c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/listitem"
    i6d1dcea335dca28dac4247e85ca94f26f27536cb9f174a0a8c8a7c8f973d41f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/permissions"
    i76d6ec189516c6c2050bcee195e5aad1f8f54affb6812643ed67dfb5af0bcff2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/content"
    ibc59dad39e3ee7ffdda6571da29b4285ac212503475fe3e1f8fd76f4e76b5ec6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/analytics"
    if00d7b9db4e3f0660dc5afd77b5e5ed0c6a2cd132701d329744707de547b1d6b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/activities"
    i15918befbf4792dc63e42c43f80e3d4503f9ac067aaed621978c3627dee349b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/subscriptions/item"
    i283b11728cdb492b1361c595b4b848744ee5622316255c6dbe99cad9d5a80932 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/thumbnails/item"
    i30d1f92bb23b5e835bc1a9ed02a97822eb5f22f8e4439f89f16a9d1a6b430584 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/permissions/item"
    icc9ebe92276d0da995e6238fcc76c634f5d802c7feb614ac82169113fdd55a98 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/versions/item"
    icd5f43bcc46b5ce4bd3436841eb61b953b961386788edd00c7bf3dc6db4f1ac9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/activities/item"
    ifef3593b8c1b16ef0ff79ccf7152d3e207f8c74ff86e88380968964af1e745d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/drives/item/root/children/item"
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
func (m *RootRequestBuilder) Activities()(*if00d7b9db4e3f0660dc5afd77b5e5ed0c6a2cd132701d329744707de547b1d6b.ActivitiesRequestBuilder) {
    return if00d7b9db4e3f0660dc5afd77b5e5ed0c6a2cd132701d329744707de547b1d6b.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*icd5f43bcc46b5ce4bd3436841eb61b953b961386788edd00c7bf3dc6db4f1ac9.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return icd5f43bcc46b5ce4bd3436841eb61b953b961386788edd00c7bf3dc6db4f1ac9.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics the analytics property
func (m *RootRequestBuilder) Analytics()(*ibc59dad39e3ee7ffdda6571da29b4285ac212503475fe3e1f8fd76f4e76b5ec6.AnalyticsRequestBuilder) {
    return ibc59dad39e3ee7ffdda6571da29b4285ac212503475fe3e1f8fd76f4e76b5ec6.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children the children property
func (m *RootRequestBuilder) Children()(*i1a2313e6ac103bc2f750a67ff7be5e8c0347553f28f2d941ac24c07416fcdb8b.ChildrenRequestBuilder) {
    return i1a2313e6ac103bc2f750a67ff7be5e8c0347553f28f2d941ac24c07416fcdb8b.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*ifef3593b8c1b16ef0ff79ccf7152d3e207f8c74ff86e88380968964af1e745d7.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return ifef3593b8c1b16ef0ff79ccf7152d3e207f8c74ff86e88380968964af1e745d7.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/drives/{drive%2Did}/root{?%24select,%24expand}";
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
func (m *RootRequestBuilder) Content()(*i76d6ec189516c6c2050bcee195e5aad1f8f54affb6812643ed67dfb5af0bcff2.ContentRequestBuilder) {
    return i76d6ec189516c6c2050bcee195e5aad1f8f54affb6812643ed67dfb5af0bcff2.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for users
func (m *RootRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property root for users
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
// CreatePatchRequestInformation update the navigation property root in users
func (m *RootRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property root in users
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
// Delete delete navigation property root for users
func (m *RootRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property root for users
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
func (m *RootRequestBuilder) ListItem()(*i65467600fd8d2b15564ff06f3e2178353d38b7a279223bcb861cd765e98c081c.ListItemRequestBuilder) {
    return i65467600fd8d2b15564ff06f3e2178353d38b7a279223bcb861cd765e98c081c.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in users
func (m *RootRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property root in users
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
func (m *RootRequestBuilder) Permissions()(*i6d1dcea335dca28dac4247e85ca94f26f27536cb9f174a0a8c8a7c8f973d41f0.PermissionsRequestBuilder) {
    return i6d1dcea335dca28dac4247e85ca94f26f27536cb9f174a0a8c8a7c8f973d41f0.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i30d1f92bb23b5e835bc1a9ed02a97822eb5f22f8e4439f89f16a9d1a6b430584.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i30d1f92bb23b5e835bc1a9ed02a97822eb5f22f8e4439f89f16a9d1a6b430584.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Subscriptions the subscriptions property
func (m *RootRequestBuilder) Subscriptions()(*i4d302afe3c097d2397bea4b437d9ba2c45601731ba2b79d26ab3673a5da462df.SubscriptionsRequestBuilder) {
    return i4d302afe3c097d2397bea4b437d9ba2c45601731ba2b79d26ab3673a5da462df.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i15918befbf4792dc63e42c43f80e3d4503f9ac067aaed621978c3627dee349b6.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return i15918befbf4792dc63e42c43f80e3d4503f9ac067aaed621978c3627dee349b6.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails the thumbnails property
func (m *RootRequestBuilder) Thumbnails()(*i63baa5e7acb91aba07e2546fca385eaf695da1babdc82b21a9601c22d143e605.ThumbnailsRequestBuilder) {
    return i63baa5e7acb91aba07e2546fca385eaf695da1babdc82b21a9601c22d143e605.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i283b11728cdb492b1361c595b4b848744ee5622316255c6dbe99cad9d5a80932.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return i283b11728cdb492b1361c595b4b848744ee5622316255c6dbe99cad9d5a80932.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Versions the versions property
func (m *RootRequestBuilder) Versions()(*i0906b36a4c6c7e8fa300e6712af3d1532d16969d436731907cbbcd07ec2c1303.VersionsRequestBuilder) {
    return i0906b36a4c6c7e8fa300e6712af3d1532d16969d436731907cbbcd07ec2c1303.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*icc9ebe92276d0da995e6238fcc76c634f5d802c7feb614ac82169113fdd55a98.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return icc9ebe92276d0da995e6238fcc76c634f5d802c7feb614ac82169113fdd55a98.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
