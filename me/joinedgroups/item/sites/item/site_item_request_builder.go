package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i17b68979b33f3f73fe620f39125bd0a458d134d9b2c94810d6b48fab948080ae "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/drive"
    i2170fc85f3abb4b9d05121b631a7f3b44602fe299fe2dc7c360adf4f1ac6a3ec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/contenttypes"
    i27233fe5d0e513e210991aeb1311472b8bf0db878b94979c1e62c26debbe97f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/columns"
    i435413fd5289981ea160f52373f3a91d69c453d19933cd5c74eeb22875d31bcd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/analytics"
    i934c04a240d2ea1b05af3c8a464b25d0a3c75a487debd676adf6fe3909b10688 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/termstore"
    i9be005967d3fe038e6f437f2c6e5b8ce1922a4f17bbc7e738c81583a616776dd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/onenote"
    ia9b5fd1d639368be14e7e594bf53811675bf0fb8b8e5f09c2b7a893f73ac78ac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/lists"
    ib04c4a23dd086f002ff3a4fec80d6729bba184920a27e770cb15328668b267e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/items"
    id2c3494771770008c9a909c5d337ee937a13ccfe6dc486af8cb3812cee6e8c0c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/externalcolumns"
    id7894033129dcd068545a3f7f3146627e2a99a736ebba6ad9a38807407ee7c68 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/drives"
    ide6e9a0189fd82a1ed144a80bb76fbe435ce9f93ab0a8745c7cd12c2338e008d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/permissions"
    ie23950a3c869f6f1f082ec3ff78582f92e3767808014295f7bf87c843c8e1994 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/sites"
    iea864e7c766b479e411f1ac10a41a8b0d3e759c75d19bdfa2327c93de409d8d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/pages"
    if4a1aa8fae3e0b33f21742ed66d0b1309fb68af2abf18fc7a8a3bfeb05473028 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/operations"
    i2a955cc499653986081b6df6c88f3e682efb7af26336bc709c37c532a4c37066 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/columns/item"
    i2ae4cf71bda535c637bc94c8f0adc90e2873d63af64d49019707f6fb5909d5f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/pages/item"
    i60ee4789d233eeb8975f64280ceaa2b1a60e6b2d537faa5fa8bf8d7111866b19 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/permissions/item"
    i70e62e70ff611ef5c15f857c08697ff928c2aef1817657a6b702326ef644f34e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/items/item"
    i8f990428c9af573c23239c02781225fcf781db4a12fe429feced300e841baa5c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/drives/item"
    ia80302f5156c3c8940a4c182426442826a5cda01cfb9e6d6fcdea2922a637cd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/externalcolumns/item"
    iabf37a1f662b37be27b4f4f07f111736a11dfb2070841c44e0a1ab4690085fb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/operations/item"
    ibde4808b20750f16db031392131d9c17b2f13d63fd5c328a2f6dae5d3276f5be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/sites/item"
    ic472504bb69fe0b751d0a44d27d60c70cbef1fe09891d2f622a090f12d2ae6e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/contenttypes/item"
    ie9df4709cd8900f524f132aea4f972a17649b22438a87c2cb46b4426a75b5f98 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/lists/item"
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
func (m *SiteItemRequestBuilder) Analytics()(*i435413fd5289981ea160f52373f3a91d69c453d19933cd5c74eeb22875d31bcd.AnalyticsRequestBuilder) {
    return i435413fd5289981ea160f52373f3a91d69c453d19933cd5c74eeb22875d31bcd.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns the columns property
func (m *SiteItemRequestBuilder) Columns()(*i27233fe5d0e513e210991aeb1311472b8bf0db878b94979c1e62c26debbe97f3.ColumnsRequestBuilder) {
    return i27233fe5d0e513e210991aeb1311472b8bf0db878b94979c1e62c26debbe97f3.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.columns.item collection
func (m *SiteItemRequestBuilder) ColumnsById(id string)(*i2a955cc499653986081b6df6c88f3e682efb7af26336bc709c37c532a4c37066.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i2a955cc499653986081b6df6c88f3e682efb7af26336bc709c37c532a4c37066.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSiteItemRequestBuilderInternal instantiates a new SiteItemRequestBuilder and sets the default values.
func NewSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SiteItemRequestBuilder) {
    m := &SiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/sites/{site%2Did}{?%24select,%24expand}";
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
func (m *SiteItemRequestBuilder) ContentTypes()(*i2170fc85f3abb4b9d05121b631a7f3b44602fe299fe2dc7c360adf4f1ac6a3ec.ContentTypesRequestBuilder) {
    return i2170fc85f3abb4b9d05121b631a7f3b44602fe299fe2dc7c360adf4f1ac6a3ec.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.contentTypes.item collection
func (m *SiteItemRequestBuilder) ContentTypesById(id string)(*ic472504bb69fe0b751d0a44d27d60c70cbef1fe09891d2f622a090f12d2ae6e3.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return ic472504bb69fe0b751d0a44d27d60c70cbef1fe09891d2f622a090f12d2ae6e3.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property sites for me
func (m *SiteItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property sites for me
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
// CreatePatchRequestInformation update the navigation property sites in me
func (m *SiteItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property sites in me
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
// Delete delete navigation property sites for me
func (m *SiteItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property sites for me
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
func (m *SiteItemRequestBuilder) Drive()(*i17b68979b33f3f73fe620f39125bd0a458d134d9b2c94810d6b48fab948080ae.DriveRequestBuilder) {
    return i17b68979b33f3f73fe620f39125bd0a458d134d9b2c94810d6b48fab948080ae.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives the drives property
func (m *SiteItemRequestBuilder) Drives()(*id7894033129dcd068545a3f7f3146627e2a99a736ebba6ad9a38807407ee7c68.DrivesRequestBuilder) {
    return id7894033129dcd068545a3f7f3146627e2a99a736ebba6ad9a38807407ee7c68.NewDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.drives.item collection
func (m *SiteItemRequestBuilder) DrivesById(id string)(*i8f990428c9af573c23239c02781225fcf781db4a12fe429feced300e841baa5c.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return i8f990428c9af573c23239c02781225fcf781db4a12fe429feced300e841baa5c.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExternalColumns the externalColumns property
func (m *SiteItemRequestBuilder) ExternalColumns()(*id2c3494771770008c9a909c5d337ee937a13ccfe6dc486af8cb3812cee6e8c0c.ExternalColumnsRequestBuilder) {
    return id2c3494771770008c9a909c5d337ee937a13ccfe6dc486af8cb3812cee6e8c0c.NewExternalColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.externalColumns.item collection
func (m *SiteItemRequestBuilder) ExternalColumnsById(id string)(*ia80302f5156c3c8940a4c182426442826a5cda01cfb9e6d6fcdea2922a637cd3.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return ia80302f5156c3c8940a4c182426442826a5cda01cfb9e6d6fcdea2922a637cd3.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *SiteItemRequestBuilder) Items()(*ib04c4a23dd086f002ff3a4fec80d6729bba184920a27e770cb15328668b267e9.ItemsRequestBuilder) {
    return ib04c4a23dd086f002ff3a4fec80d6729bba184920a27e770cb15328668b267e9.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.items.item collection
func (m *SiteItemRequestBuilder) ItemsById(id string)(*i70e62e70ff611ef5c15f857c08697ff928c2aef1817657a6b702326ef644f34e.BaseItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["baseItem%2Did"] = id
    }
    return i70e62e70ff611ef5c15f857c08697ff928c2aef1817657a6b702326ef644f34e.NewBaseItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Lists the lists property
func (m *SiteItemRequestBuilder) Lists()(*ia9b5fd1d639368be14e7e594bf53811675bf0fb8b8e5f09c2b7a893f73ac78ac.ListsRequestBuilder) {
    return ia9b5fd1d639368be14e7e594bf53811675bf0fb8b8e5f09c2b7a893f73ac78ac.NewListsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.lists.item collection
func (m *SiteItemRequestBuilder) ListsById(id string)(*ie9df4709cd8900f524f132aea4f972a17649b22438a87c2cb46b4426a75b5f98.ListItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["list%2Did"] = id
    }
    return ie9df4709cd8900f524f132aea4f972a17649b22438a87c2cb46b4426a75b5f98.NewListItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote the onenote property
func (m *SiteItemRequestBuilder) Onenote()(*i9be005967d3fe038e6f437f2c6e5b8ce1922a4f17bbc7e738c81583a616776dd.OnenoteRequestBuilder) {
    return i9be005967d3fe038e6f437f2c6e5b8ce1922a4f17bbc7e738c81583a616776dd.NewOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations the operations property
func (m *SiteItemRequestBuilder) Operations()(*if4a1aa8fae3e0b33f21742ed66d0b1309fb68af2abf18fc7a8a3bfeb05473028.OperationsRequestBuilder) {
    return if4a1aa8fae3e0b33f21742ed66d0b1309fb68af2abf18fc7a8a3bfeb05473028.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.operations.item collection
func (m *SiteItemRequestBuilder) OperationsById(id string)(*iabf37a1f662b37be27b4f4f07f111736a11dfb2070841c44e0a1ab4690085fb2.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return iabf37a1f662b37be27b4f4f07f111736a11dfb2070841c44e0a1ab4690085fb2.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages the pages property
func (m *SiteItemRequestBuilder) Pages()(*iea864e7c766b479e411f1ac10a41a8b0d3e759c75d19bdfa2327c93de409d8d0.PagesRequestBuilder) {
    return iea864e7c766b479e411f1ac10a41a8b0d3e759c75d19bdfa2327c93de409d8d0.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.pages.item collection
func (m *SiteItemRequestBuilder) PagesById(id string)(*i2ae4cf71bda535c637bc94c8f0adc90e2873d63af64d49019707f6fb5909d5f1.SitePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sitePage%2Did"] = id
    }
    return i2ae4cf71bda535c637bc94c8f0adc90e2873d63af64d49019707f6fb5909d5f1.NewSitePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property sites in me
func (m *SiteItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Siteable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property sites in me
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
func (m *SiteItemRequestBuilder) Permissions()(*ide6e9a0189fd82a1ed144a80bb76fbe435ce9f93ab0a8745c7cd12c2338e008d.PermissionsRequestBuilder) {
    return ide6e9a0189fd82a1ed144a80bb76fbe435ce9f93ab0a8745c7cd12c2338e008d.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.permissions.item collection
func (m *SiteItemRequestBuilder) PermissionsById(id string)(*i60ee4789d233eeb8975f64280ceaa2b1a60e6b2d537faa5fa8bf8d7111866b19.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return i60ee4789d233eeb8975f64280ceaa2b1a60e6b2d537faa5fa8bf8d7111866b19.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites the sites property
func (m *SiteItemRequestBuilder) Sites()(*ie23950a3c869f6f1f082ec3ff78582f92e3767808014295f7bf87c843c8e1994.SitesRequestBuilder) {
    return ie23950a3c869f6f1f082ec3ff78582f92e3767808014295f7bf87c843c8e1994.NewSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.sites.item collection
func (m *SiteItemRequestBuilder) SitesById(id string)(*ibde4808b20750f16db031392131d9c17b2f13d63fd5c328a2f6dae5d3276f5be.SiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did1"] = id
    }
    return ibde4808b20750f16db031392131d9c17b2f13d63fd5c328a2f6dae5d3276f5be.NewSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TermStore the termStore property
func (m *SiteItemRequestBuilder) TermStore()(*i934c04a240d2ea1b05af3c8a464b25d0a3c75a487debd676adf6fe3909b10688.TermStoreRequestBuilder) {
    return i934c04a240d2ea1b05af3c8a464b25d0a3c75a487debd676adf6fe3909b10688.NewTermStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
