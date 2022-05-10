package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1b9902a2ef86bd951509e07e954ebdfcdb294d3b1e2b19720e00a93fa56ddc57 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/associatewithhubsites"
    i45e1ee212da0df658d58fd49805788760ad2db75fc50e6b6dcaa7dfcd00d3978 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/publish"
    i4863bf5c7586e4ad81fac7ab2caed19b418476f5f5675581ed4af3ed8985a288 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    i5037350d16f6cf4dece9c26765ed1ccfef9e4c65d559a5de3ce35af4c86eac3b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/unpublish"
    i7302a22b9931382d97b6c0ae96a44eeba17b344f3d44ed2e3f30f8f0f59ad8e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/base"
    i74075aa2a530d424999ae192de09b7496b812900c2559619ecc5f5d35d237ea0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/basetypes"
    iaa097b581c6199fb3dfdaf1f55ad057eabfc418daa12fda6229a111298944fc8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/columnlinks"
    iae645ef5bc4621f5053cb12c5ff8e22e42600be43d3a4dcb482b14f5a67fc2cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/columns"
    id017cbb288ce2d5950b91131681474db54e26f7b06e75d489706ddc90507dabd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/ispublished"
    if1c51b409818818be051c95b10092e9778aa1345a4d75dcc4b3f02772723ebb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/columnpositions"
    i26c0a95473fca7f3815d1e9a3056a4392983bc9e6245967e1159b481e65bc277 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/columns/item"
    i584aea776e37bd3d5c9637909dd1ee1ff8521ed3349782a607bfe93fca4578b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/columnpositions/item"
    i8f38eccc41aee4ac69d2f13ff1b3b1281484c4fe6531a54deec4ca710c401ef5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/columnlinks/item"
    if7aa086d5ef08ec3cc087e5d4539c65ae02287dc267dab98e681a4bbdcdbfebe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/drives/item/list/contenttypes/item/basetypes/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ContentTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ContentTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ContentTypeItemRequestBuilderGetQueryParameters
}
// ContentTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssociateWithHubSites the associateWithHubSites property
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i1b9902a2ef86bd951509e07e954ebdfcdb294d3b1e2b19720e00a93fa56ddc57.AssociateWithHubSitesRequestBuilder) {
    return i1b9902a2ef86bd951509e07e954ebdfcdb294d3b1e2b19720e00a93fa56ddc57.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*i7302a22b9931382d97b6c0ae96a44eeba17b344f3d44ed2e3f30f8f0f59ad8e6.BaseRequestBuilder) {
    return i7302a22b9931382d97b6c0ae96a44eeba17b344f3d44ed2e3f30f8f0f59ad8e6.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i74075aa2a530d424999ae192de09b7496b812900c2559619ecc5f5d35d237ea0.BaseTypesRequestBuilder) {
    return i74075aa2a530d424999ae192de09b7496b812900c2559619ecc5f5d35d237ea0.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.drives.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*if7aa086d5ef08ec3cc087e5d4539c65ae02287dc267dab98e681a4bbdcdbfebe.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return if7aa086d5ef08ec3cc087e5d4539c65ae02287dc267dab98e681a4bbdcdbfebe.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*iaa097b581c6199fb3dfdaf1f55ad057eabfc418daa12fda6229a111298944fc8.ColumnLinksRequestBuilder) {
    return iaa097b581c6199fb3dfdaf1f55ad057eabfc418daa12fda6229a111298944fc8.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.drives.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i8f38eccc41aee4ac69d2f13ff1b3b1281484c4fe6531a54deec4ca710c401ef5.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return i8f38eccc41aee4ac69d2f13ff1b3b1281484c4fe6531a54deec4ca710c401ef5.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*if1c51b409818818be051c95b10092e9778aa1345a4d75dcc4b3f02772723ebb9.ColumnPositionsRequestBuilder) {
    return if1c51b409818818be051c95b10092e9778aa1345a4d75dcc4b3f02772723ebb9.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.drives.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i584aea776e37bd3d5c9637909dd1ee1ff8521ed3349782a607bfe93fca4578b5.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i584aea776e37bd3d5c9637909dd1ee1ff8521ed3349782a607bfe93fca4578b5.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*iae645ef5bc4621f5053cb12c5ff8e22e42600be43d3a4dcb482b14f5a67fc2cd.ColumnsRequestBuilder) {
    return iae645ef5bc4621f5053cb12c5ff8e22e42600be43d3a4dcb482b14f5a67fc2cd.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.drives.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i26c0a95473fca7f3815d1e9a3056a4392983bc9e6245967e1159b481e65bc277.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i26c0a95473fca7f3815d1e9a3056a4392983bc9e6245967e1159b481e65bc277.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/drives/{drive%2Did}/list/contentTypes/{contentType%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToDefaultContentLocation the copyToDefaultContentLocation property
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i4863bf5c7586e4ad81fac7ab2caed19b418476f5f5675581ed4af3ed8985a288.CopyToDefaultContentLocationRequestBuilder) {
    return i4863bf5c7586e4ad81fac7ab2caed19b418476f5f5675581ed4af3ed8985a288.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for me
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property contentTypes for me
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property contentTypes in me
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property contentTypes in me
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *ContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property contentTypes for me
func (m *ContentTypeItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property contentTypes for me
func (m *ContentTypeItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ContentTypeItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ContentTypeItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*id017cbb288ce2d5950b91131681474db54e26f7b06e75d489706ddc90507dabd.IsPublishedRequestBuilder) {
    return id017cbb288ce2d5950b91131681474db54e26f7b06e75d489706ddc90507dabd.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in me
func (m *ContentTypeItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property contentTypes in me
func (m *ContentTypeItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, requestConfiguration *ContentTypeItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Publish the publish property
func (m *ContentTypeItemRequestBuilder) Publish()(*i45e1ee212da0df658d58fd49805788760ad2db75fc50e6b6dcaa7dfcd00d3978.PublishRequestBuilder) {
    return i45e1ee212da0df658d58fd49805788760ad2db75fc50e6b6dcaa7dfcd00d3978.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i5037350d16f6cf4dece9c26765ed1ccfef9e4c65d559a5de3ce35af4c86eac3b.UnpublishRequestBuilder) {
    return i5037350d16f6cf4dece9c26765ed1ccfef9e4c65d559a5de3ce35af4c86eac3b.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
