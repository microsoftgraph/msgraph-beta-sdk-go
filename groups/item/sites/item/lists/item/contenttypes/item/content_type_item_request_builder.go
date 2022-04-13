package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i406acee98e174f55be5c00f9a9944c963f17ea0d3d73146c4282398e65ded440 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/ispublished"
    i4ff70d63f353c245e3289e522701321b799064d358151b0ca79fd2873a834762 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columns"
    i7bd409a67e70ab70d9cc710822987f4586a327e808a55f1ac24dd0378ae889d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/basetypes"
    i7f64c863306e1e1293f0f12a26eabc7421a188a0d1f5263d7ba3e2ffe6252f26 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columnpositions"
    i8600db2d55639da431b37656207b0c98ddc59b50dcd1397b8765ca0850ce4342 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/base"
    i86dbe3439572a444ae30dd4bdb86d652d790107bc159f836b83b1865e60555a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columnlinks"
    i8c5a0674d846ab57c33dff9b351ff3ff0180750aa37b597c6f834debbb778789 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/copytodefaultcontentlocation"
    ibd5a8f4f43ecbd93666f4ced2d9185fbd0273067379d33fc59dede358ba3dd72 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/publish"
    ica1f4a86eeb18758022cf5081583a57bee454a99bee0b0089c9259acd8d33b96 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/associatewithhubsites"
    iefdb1b0c46303dda2a483e20bc2adb6cfec5b355ef20ca9330ad1a2eb52d1a30 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/unpublish"
    i23c412dbeb647c0450a16421812d00c779dbdafb19526b5ca20dbda018843812 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columns/item"
    i53d995fb4a9fdbb93b96ae64f9b566f4db8941eee407e8f956eb99c64192a5c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columnpositions/item"
    i9893720f698491c8a4882b1eaab041e316c1b228975788f5b2a18dc2122e3f8c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/columnlinks/item"
    ib65bfaaa1198a07dd17420f2942fd6f01e1954d8fdc2b8857294e8c22e9ac6ed "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/lists/item/contenttypes/item/basetypes/item"
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
// ContentTypeItemRequestBuilderDeleteOptions options for Delete
type ContentTypeItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ContentTypeItemRequestBuilderGetOptions options for Get
type ContentTypeItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ContentTypeItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ContentTypeItemRequestBuilderPatchOptions options for Patch
type ContentTypeItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AssociateWithHubSites the associateWithHubSites property
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*ica1f4a86eeb18758022cf5081583a57bee454a99bee0b0089c9259acd8d33b96.AssociateWithHubSitesRequestBuilder) {
    return ica1f4a86eeb18758022cf5081583a57bee454a99bee0b0089c9259acd8d33b96.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*i8600db2d55639da431b37656207b0c98ddc59b50dcd1397b8765ca0850ce4342.BaseRequestBuilder) {
    return i8600db2d55639da431b37656207b0c98ddc59b50dcd1397b8765ca0850ce4342.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i7bd409a67e70ab70d9cc710822987f4586a327e808a55f1ac24dd0378ae889d0.BaseTypesRequestBuilder) {
    return i7bd409a67e70ab70d9cc710822987f4586a327e808a55f1ac24dd0378ae889d0.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*ib65bfaaa1198a07dd17420f2942fd6f01e1954d8fdc2b8857294e8c22e9ac6ed.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return ib65bfaaa1198a07dd17420f2942fd6f01e1954d8fdc2b8857294e8c22e9ac6ed.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i86dbe3439572a444ae30dd4bdb86d652d790107bc159f836b83b1865e60555a5.ColumnLinksRequestBuilder) {
    return i86dbe3439572a444ae30dd4bdb86d652d790107bc159f836b83b1865e60555a5.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i9893720f698491c8a4882b1eaab041e316c1b228975788f5b2a18dc2122e3f8c.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return i9893720f698491c8a4882b1eaab041e316c1b228975788f5b2a18dc2122e3f8c.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i7f64c863306e1e1293f0f12a26eabc7421a188a0d1f5263d7ba3e2ffe6252f26.ColumnPositionsRequestBuilder) {
    return i7f64c863306e1e1293f0f12a26eabc7421a188a0d1f5263d7ba3e2ffe6252f26.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i53d995fb4a9fdbb93b96ae64f9b566f4db8941eee407e8f956eb99c64192a5c7.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i53d995fb4a9fdbb93b96ae64f9b566f4db8941eee407e8f956eb99c64192a5c7.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*i4ff70d63f353c245e3289e522701321b799064d358151b0ca79fd2873a834762.ColumnsRequestBuilder) {
    return i4ff70d63f353c245e3289e522701321b799064d358151b0ca79fd2873a834762.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.lists.item.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i23c412dbeb647c0450a16421812d00c779dbdafb19526b5ca20dbda018843812.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i23c412dbeb647c0450a16421812d00c779dbdafb19526b5ca20dbda018843812.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}{?%24select,%24expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i8c5a0674d846ab57c33dff9b351ff3ff0180750aa37b597c6f834debbb778789.CopyToDefaultContentLocationRequestBuilder) {
    return i8c5a0674d846ab57c33dff9b351ff3ff0180750aa37b597c6f834debbb778789.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation(options *ContentTypeItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(options *ContentTypeItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) Delete(options *ContentTypeItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i406acee98e174f55be5c00f9a9944c963f17ea0d3d73146c4282398e65ded440.IsPublishedRequestBuilder) {
    return i406acee98e174f55be5c00f9a9944c963f17ea0d3d73146c4282398e65ded440.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) Patch(options *ContentTypeItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Publish the publish property
func (m *ContentTypeItemRequestBuilder) Publish()(*ibd5a8f4f43ecbd93666f4ced2d9185fbd0273067379d33fc59dede358ba3dd72.PublishRequestBuilder) {
    return ibd5a8f4f43ecbd93666f4ced2d9185fbd0273067379d33fc59dede358ba3dd72.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*iefdb1b0c46303dda2a483e20bc2adb6cfec5b355ef20ca9330ad1a2eb52d1a30.UnpublishRequestBuilder) {
    return iefdb1b0c46303dda2a483e20bc2adb6cfec5b355ef20ca9330ad1a2eb52d1a30.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
