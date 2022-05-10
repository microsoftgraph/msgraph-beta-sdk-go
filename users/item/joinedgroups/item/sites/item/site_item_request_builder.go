package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i13544e2d8aa9994bd010d34ed0ea70bac4c40c9957efb17a9c1abdbae0eaecf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/columns"
    i1479843598a69e3bbc076430c6f0596e5946f1da72fa81a3f660bf15fbae343e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/termstore"
    i1ca43ffcf305f2fac8c710e3bedd1f56f13a2caf9a07d207ea920c31c22a013b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/analytics"
    i1e79b58e80e6a71186bd81f26af3d29fb517b862a9daf7bb0d17e952f0318ce0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/items"
    i237f5c8a56879e0e746a07dbed81675621fb3e59e829d29e510e2bbd7636bef1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/permissions"
    i61dd201b8cbd4403133f98a782c6bb177200c06f5ba89f8581bf1d4e396e17c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/externalcolumns"
    i7ab7bea9e09b4f256695fbc33d3245c1ea71636d79985a249166223d387c0f75 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/operations"
    i8d733ccce229c9049050007dfdc1eedcd26df58aeba7e013b8b570a5284f9b86 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/onenote"
    ia5f4717b3bd819a92c0763dd046c872f7fe3fdb24e57290f78237727174c3339 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/drive"
    iaf8da9a307fb3f653d5b1915276d32d347c952d8d4b76d376e65aa526a30b68d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/contenttypes"
    iafeecbba9f77222d4d0992948495ee2ec37d6b03926157a26fd0df561b9d081d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/sites"
    ic1b087c8847224d2062a3c47327f019ee8f455bbbd038fd925dba5b333c31714 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/lists"
    icb73fecf7ecf71182eb6f99f3d8c9cdb943a8d632adc3bfbb5dbdf06be3fc1d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/drives"
    ie55865ab3e66dfb4f79d9beae2055c61b756cd87344722b07f64527606d6d91f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/pages"
    i09fc093eccde79e94fc33f2fd74705dea81838cec2d9bebe4e7ebc2d455dfaa5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/contenttypes/item"
    i39a23dcb7d0b40e29450d10f35630912f65c6f807a0f7c6b604ebf06067e2fe2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/externalcolumns/item"
    i414b158083f36e2f446419077f090f93f0e04a3fe73341f5efda18a8c03ebcca "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/permissions/item"
    i41e2fd899d45a48d6c2222aa34264c8a7c40389c9e2675f85a8bbb8844e100b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/sites/item"
    i46363e4c574a649f10d7a4e1a7e057455de9b24fc2238c38b5be8a860bc442cf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/columns/item"
    i61750667379c51833b4f2ef880e0590b18c2a2764c7574fa4b78d0c2acde4188 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/operations/item"
    i8460d64e083f260f11f5628152e32b47ba640c5ab9150fb47fa0ca2ee1c68abe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/pages/item"
    iebea6253cc4ad0c0c4ea5977da373d7e11439140c8169d9ca70d66e9f13c8239 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/drives/item"
    if250a18e28948ec9dd9e98c6b6086d37e24bba6ad53a864ecbfa26595c8a23fe "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/lists/item"
    if45e6f360d2744e6ae92d23cad00928c3b7f83ca17d77b7f66e215dc85e0d36f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/sites/item/items/item"
)

// SiteItemRequestBuilder provides operations to manage the sites property of the microsoft.graph.group entity.
type SiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SiteItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SiteItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SiteItemRequestBuilderGetQueryParameters the list of SharePoint sites in this group. Access the default site with /sites/root.
type SiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SiteItemRequestBuilderGetQueryParameters
}
// SiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics the analytics property
func (m *SiteItemRequestBuilder) Analytics()(*i1ca43ffcf305f2fac8c710e3bedd1f56f13a2caf9a07d207ea920c31c22a013b.AnalyticsRequestBuilder) {
    return i1ca43ffcf305f2fac8c710e3bedd1f56f13a2caf9a07d207ea920c31c22a013b.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns the columns property
func (m *SiteItemRequestBuilder) Columns()(*i13544e2d8aa9994bd010d34ed0ea70bac4c40c9957efb17a9c1abdbae0eaecf2.ColumnsRequestBuilder) {
    return i13544e2d8aa9994bd010d34ed0ea70bac4c40c9957efb17a9c1abdbae0eaecf2.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.sites.item.columns.item collection
func (m *SiteItemRequestBuilder) ColumnsById(id string)(*i46363e4c574a649f10d7a4e1a7e057455de9b24fc2238c38b5be8a860bc442cf.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i46363e4c574a649f10d7a4e1a7e057455de9b24fc2238c38b5be8a860bc442cf.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SiteItemRequestBuilder) {
    m := &SiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/sites/{site%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSiteItemRequestBuilder instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes the contentTypes property
func (m *SiteItemRequestBuilder) ContentTypes()(*iaf8da9a307fb3f653d5b1915276d32d347c952d8d4b76d376e65aa526a30b68d.ContentTypesRequestBuilder) {
    return iaf8da9a307fb3f653d5b1915276d32d347c952d8d4b76d376e65aa526a30b68d.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.sites.item.contentTypes.item collection
func (m *SiteItemRequestBuilder) ContentTypesById(id string)(*i09fc093eccde79e94fc33f2fd74705dea81838cec2d9bebe4e7ebc2d455dfaa5.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return i09fc093eccde79e94fc33f2fd74705dea81838cec2d9bebe4e7ebc2d455dfaa5.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property sites for users
func (m *SiteItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property sites for users
func (m *SiteItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *SiteItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *SiteItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *SiteItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sites in users
func (m *SiteItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property sites in users
func (m *SiteItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, requestConfiguration *SiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property sites for users
func (m *SiteItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property sites for users
func (m *SiteItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *SiteItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Drive the drive property
func (m *SiteItemRequestBuilder) Drive()(*ia5f4717b3bd819a92c0763dd046c872f7fe3fdb24e57290f78237727174c3339.DriveRequestBuilder) {
    return ia5f4717b3bd819a92c0763dd046c872f7fe3fdb24e57290f78237727174c3339.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives the drives property
func (m *SiteItemRequestBuilder) Drives()(*icb73fecf7ecf71182eb6f99f3d8c9cdb943a8d632adc3bfbb5dbdf06be3fc1d0.DrivesRequestBuilder) {
    return icb73fecf7ecf71182eb6f99f3d8c9cdb943a8d632adc3bfbb5dbdf06be3fc1d0.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.sites.item.drives.item collection
func (m *SiteItemRequestBuilder) DrivesById(id string)(*iebea6253cc4ad0c0c4ea5977da373d7e11439140c8169d9ca70d66e9f13c8239.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return iebea6253cc4ad0c0c4ea5977da373d7e11439140c8169d9ca70d66e9f13c8239.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExternalColumns the externalColumns property
func (m *SiteItemRequestBuilder) ExternalColumns()(*i61dd201b8cbd4403133f98a782c6bb177200c06f5ba89f8581bf1d4e396e17c4.ExternalColumnsRequestBuilder) {
    return i61dd201b8cbd4403133f98a782c6bb177200c06f5ba89f8581bf1d4e396e17c4.NewExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.sites.item.externalColumns.item collection
func (m *SiteItemRequestBuilder) ExternalColumnsById(id string)(*i39a23dcb7d0b40e29450d10f35630912f65c6f807a0f7c6b604ebf06067e2fe2.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i39a23dcb7d0b40e29450d10f35630912f65c6f807a0f7c6b604ebf06067e2fe2.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *SiteItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *SiteItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *SiteItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSiteFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable), nil
}
// Items the items property
func (m *SiteItemRequestBuilder) Items()(*i1e79b58e80e6a71186bd81f26af3d29fb517b862a9daf7bb0d17e952f0318ce0.ItemsRequestBuilder) {
    return i1e79b58e80e6a71186bd81f26af3d29fb517b862a9daf7bb0d17e952f0318ce0.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.sites.item.items.item collection
func (m *SiteItemRequestBuilder) ItemsById(id string)(*if45e6f360d2744e6ae92d23cad00928c3b7f83ca17d77b7f66e215dc85e0d36f.BaseItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseItem%2Did"] = id
    }
    return if45e6f360d2744e6ae92d23cad00928c3b7f83ca17d77b7f66e215dc85e0d36f.NewBaseItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Lists the lists property
func (m *SiteItemRequestBuilder) Lists()(*ic1b087c8847224d2062a3c47327f019ee8f455bbbd038fd925dba5b333c31714.ListsRequestBuilder) {
    return ic1b087c8847224d2062a3c47327f019ee8f455bbbd038fd925dba5b333c31714.NewListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.sites.item.lists.item collection
func (m *SiteItemRequestBuilder) ListsById(id string)(*if250a18e28948ec9dd9e98c6b6086d37e24bba6ad53a864ecbfa26595c8a23fe.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["list%2Did"] = id
    }
    return if250a18e28948ec9dd9e98c6b6086d37e24bba6ad53a864ecbfa26595c8a23fe.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote the onenote property
func (m *SiteItemRequestBuilder) Onenote()(*i8d733ccce229c9049050007dfdc1eedcd26df58aeba7e013b8b570a5284f9b86.OnenoteRequestBuilder) {
    return i8d733ccce229c9049050007dfdc1eedcd26df58aeba7e013b8b570a5284f9b86.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations the operations property
func (m *SiteItemRequestBuilder) Operations()(*i7ab7bea9e09b4f256695fbc33d3245c1ea71636d79985a249166223d387c0f75.OperationsRequestBuilder) {
    return i7ab7bea9e09b4f256695fbc33d3245c1ea71636d79985a249166223d387c0f75.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.sites.item.operations.item collection
func (m *SiteItemRequestBuilder) OperationsById(id string)(*i61750667379c51833b4f2ef880e0590b18c2a2764c7574fa4b78d0c2acde4188.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return i61750667379c51833b4f2ef880e0590b18c2a2764c7574fa4b78d0c2acde4188.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages the pages property
func (m *SiteItemRequestBuilder) Pages()(*ie55865ab3e66dfb4f79d9beae2055c61b756cd87344722b07f64527606d6d91f.PagesRequestBuilder) {
    return ie55865ab3e66dfb4f79d9beae2055c61b756cd87344722b07f64527606d6d91f.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.sites.item.pages.item collection
func (m *SiteItemRequestBuilder) PagesById(id string)(*i8460d64e083f260f11f5628152e32b47ba640c5ab9150fb47fa0ca2ee1c68abe.SitePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sitePage%2Did"] = id
    }
    return i8460d64e083f260f11f5628152e32b47ba640c5ab9150fb47fa0ca2ee1c68abe.NewSitePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property sites in users
func (m *SiteItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property sites in users
func (m *SiteItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable, requestConfiguration *SiteItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *SiteItemRequestBuilder) Permissions()(*i237f5c8a56879e0e746a07dbed81675621fb3e59e829d29e510e2bbd7636bef1.PermissionsRequestBuilder) {
    return i237f5c8a56879e0e746a07dbed81675621fb3e59e829d29e510e2bbd7636bef1.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.sites.item.permissions.item collection
func (m *SiteItemRequestBuilder) PermissionsById(id string)(*i414b158083f36e2f446419077f090f93f0e04a3fe73341f5efda18a8c03ebcca.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i414b158083f36e2f446419077f090f93f0e04a3fe73341f5efda18a8c03ebcca.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites the sites property
func (m *SiteItemRequestBuilder) Sites()(*iafeecbba9f77222d4d0992948495ee2ec37d6b03926157a26fd0df561b9d081d.SitesRequestBuilder) {
    return iafeecbba9f77222d4d0992948495ee2ec37d6b03926157a26fd0df561b9d081d.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.sites.item.sites.item collection
func (m *SiteItemRequestBuilder) SitesById(id string)(*i41e2fd899d45a48d6c2222aa34264c8a7c40389c9e2675f85a8bbb8844e100b8.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did1"] = id
    }
    return i41e2fd899d45a48d6c2222aa34264c8a7c40389c9e2675f85a8bbb8844e100b8.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TermStore the termStore property
func (m *SiteItemRequestBuilder) TermStore()(*i1479843598a69e3bbc076430c6f0596e5946f1da72fa81a3f660bf15fbae343e.TermStoreRequestBuilder) {
    return i1479843598a69e3bbc076430c6f0596e5946f1da72fa81a3f660bf15fbae343e.NewTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
