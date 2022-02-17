package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0eae2228a06806e82320cf19a03d576b42621cc4a2dd8806ce532a57bfebb362 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/associatewithhubsites"
    i105f5ec8936f20cc343ef53f19692d9484bb0e2317204f4e0a83cb86af419abf "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columnlinks"
    i20d9f259a492016a967dd6a10e660b98616fd7d4888ef6ee8111990cf5517382 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/unpublish"
    i6d9ef8ddda3a3e7ebcb0a0d568b23e622439e98ff5309351bc26797ae8d500cb "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columnpositions"
    i7e506eefecd6b2935e76b82d621a57aa9ed417d501afb7d19965f3434a172010 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/basetypes"
    i7e8f4985115eafb5e1b9a5dcfae2a3c836415e064af553f7801cdbb55757ca11 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/base"
    i910a70f3e4a25581ff8d025ecc5ab7dd183edf0f6e7af5d6e64e49e265867068 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/copytodefaultcontentlocation"
    i943b393012430b44d1cf499bc2d8e50d28fa67a47c6be6689e906105bd5eb4e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/ispublished"
    if8b488d84d1b1415d1c88a73424e76e4a1a1efd4526e5164d0604809ea5c8d13 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/publish"
    ifd4412476ff102f7c92f4e4a139b13990d10cc2647ddfa6e976c2dd5a98653ce "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columns"
    i88d2724fa2dada76f95cc3b7d1c2d9bd0f61469930ea96ff9d0dc112126c5a8e "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columnlinks/item"
    i9e8f3bdc2a2da10dbbb9f032cac127a9f38b1132156d876cec226651df880b26 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columns/item"
)

// ContentTypeRequestBuilder builds and executes requests for operations under \sites\{site-id}\contentTypes\{contentType-id}
type ContentTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContentTypeRequestBuilderDeleteOptions options for Delete
type ContentTypeRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeRequestBuilderGetOptions options for Get
type ContentTypeRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContentTypeRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeRequestBuilderGetQueryParameters the collection of content types defined for this site.
type ContentTypeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContentTypeRequestBuilderPatchOptions options for Patch
type ContentTypeRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContentTypeRequestBuilder) AssociateWithHubSites()(*i0eae2228a06806e82320cf19a03d576b42621cc4a2dd8806ce532a57bfebb362.AssociateWithHubSitesRequestBuilder) {
    return i0eae2228a06806e82320cf19a03d576b42621cc4a2dd8806ce532a57bfebb362.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Base()(*i7e8f4985115eafb5e1b9a5dcfae2a3c836415e064af553f7801cdbb55757ca11.BaseRequestBuilder) {
    return i7e8f4985115eafb5e1b9a5dcfae2a3c836415e064af553f7801cdbb55757ca11.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) BaseTypes()(*i7e506eefecd6b2935e76b82d621a57aa9ed417d501afb7d19965f3434a172010.BaseTypesRequestBuilder) {
    return i7e506eefecd6b2935e76b82d621a57aa9ed417d501afb7d19965f3434a172010.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnLinks()(*i105f5ec8936f20cc343ef53f19692d9484bb0e2317204f4e0a83cb86af419abf.ColumnLinksRequestBuilder) {
    return i105f5ec8936f20cc343ef53f19692d9484bb0e2317204f4e0a83cb86af419abf.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.contentTypes.item.columnLinks.item collection
func (m *ContentTypeRequestBuilder) ColumnLinksById(id string)(*i88d2724fa2dada76f95cc3b7d1c2d9bd0f61469930ea96ff9d0dc112126c5a8e.ColumnLinkRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i88d2724fa2dada76f95cc3b7d1c2d9bd0f61469930ea96ff9d0dc112126c5a8e.NewColumnLinkRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnPositions()(*i6d9ef8ddda3a3e7ebcb0a0d568b23e622439e98ff5309351bc26797ae8d500cb.ColumnPositionsRequestBuilder) {
    return i6d9ef8ddda3a3e7ebcb0a0d568b23e622439e98ff5309351bc26797ae8d500cb.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Columns()(*ifd4412476ff102f7c92f4e4a139b13990d10cc2647ddfa6e976c2dd5a98653ce.ColumnsRequestBuilder) {
    return ifd4412476ff102f7c92f4e4a139b13990d10cc2647ddfa6e976c2dd5a98653ce.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.contentTypes.item.columns.item collection
func (m *ContentTypeRequestBuilder) ColumnsById(id string)(*i9e8f3bdc2a2da10dbbb9f032cac127a9f38b1132156d876cec226651df880b26.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i9e8f3bdc2a2da10dbbb9f032cac127a9f38b1132156d876cec226651df880b26.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeRequestBuilderInternal instantiates a new ContentTypeRequestBuilder and sets the default values.
func NewContentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    m := &ContentTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeRequestBuilder instantiates a new ContentTypeRequestBuilder and sets the default values.
func NewContentTypeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeRequestBuilder) CopyToDefaultContentLocation()(*i910a70f3e4a25581ff8d025ecc5ab7dd183edf0f6e7af5d6e64e49e265867068.CopyToDefaultContentLocationRequestBuilder) {
    return i910a70f3e4a25581ff8d025ecc5ab7dd183edf0f6e7af5d6e64e49e265867068.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the collection of content types defined for this site.
func (m *ContentTypeRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of content types defined for this site.
func (m *ContentTypeRequestBuilder) CreateGetRequestInformation(options *ContentTypeRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the collection of content types defined for this site.
func (m *ContentTypeRequestBuilder) CreatePatchRequestInformation(options *ContentTypeRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the collection of content types defined for this site.
func (m *ContentTypeRequestBuilder) Delete(options *ContentTypeRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types defined for this site.
func (m *ContentTypeRequestBuilder) Get(options *ContentTypeRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContentType() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentType), nil
}
// IsPublished builds and executes requests for operations under \sites\{site-id}\contentTypes\{contentType-id}\microsoft.graph.isPublished()
func (m *ContentTypeRequestBuilder) IsPublished()(*i943b393012430b44d1cf499bc2d8e50d28fa67a47c6be6689e906105bd5eb4e3.IsPublishedRequestBuilder) {
    return i943b393012430b44d1cf499bc2d8e50d28fa67a47c6be6689e906105bd5eb4e3.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the collection of content types defined for this site.
func (m *ContentTypeRequestBuilder) Patch(options *ContentTypeRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeRequestBuilder) Publish()(*if8b488d84d1b1415d1c88a73424e76e4a1a1efd4526e5164d0604809ea5c8d13.PublishRequestBuilder) {
    return if8b488d84d1b1415d1c88a73424e76e4a1a1efd4526e5164d0604809ea5c8d13.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Unpublish()(*i20d9f259a492016a967dd6a10e660b98616fd7d4888ef6ee8111990cf5517382.UnpublishRequestBuilder) {
    return i20d9f259a492016a967dd6a10e660b98616fd7d4888ef6ee8111990cf5517382.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
