package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContentTypeItemRequestBuilderDeleteOptions options for Delete
type ContentTypeItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetOptions options for Get
type ContentTypeItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContentTypeItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContentTypeItemRequestBuilderPatchOptions options for Patch
type ContentTypeItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*ica1f4a86eeb18758022cf5081583a57bee454a99bee0b0089c9259acd8d33b96.AssociateWithHubSitesRequestBuilder) {
    return ica1f4a86eeb18758022cf5081583a57bee454a99bee0b0089c9259acd8d33b96.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*i8600db2d55639da431b37656207b0c98ddc59b50dcd1397b8765ca0850ce4342.BaseRequestBuilder) {
    return i8600db2d55639da431b37656207b0c98ddc59b50dcd1397b8765ca0850ce4342.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["contentType_id1"] = id
    }
    return ib65bfaaa1198a07dd17420f2942fd6f01e1954d8fdc2b8857294e8c22e9ac6ed.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["columnLink_id"] = id
    }
    return i9893720f698491c8a4882b1eaab041e316c1b228975788f5b2a18dc2122e3f8c.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["columnDefinition_id"] = id
    }
    return i53d995fb4a9fdbb93b96ae64f9b566f4db8941eee407e8f956eb99c64192a5c7.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["columnDefinition_id"] = id
    }
    return i23c412dbeb647c0450a16421812d00c779dbdafb19526b5ca20dbda018843812.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/sites/{site_id}/lists/{list_id}/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i8c5a0674d846ab57c33dff9b351ff3ff0180750aa37b597c6f834debbb778789.CopyToDefaultContentLocationRequestBuilder) {
    return i8c5a0674d846ab57c33dff9b351ff3ff0180750aa37b597c6f834debbb778789.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for groups
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation(options *ContentTypeItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property contentTypes in groups
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(options *ContentTypeItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
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
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable), nil
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
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeItemRequestBuilder) Publish()(*ibd5a8f4f43ecbd93666f4ced2d9185fbd0273067379d33fc59dede358ba3dd72.PublishRequestBuilder) {
    return ibd5a8f4f43ecbd93666f4ced2d9185fbd0273067379d33fc59dede358ba3dd72.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*iefdb1b0c46303dda2a483e20bc2adb6cfec5b355ef20ca9330ad1a2eb52d1a30.UnpublishRequestBuilder) {
    return iefdb1b0c46303dda2a483e20bc2adb6cfec5b355ef20ca9330ad1a2eb52d1a30.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
