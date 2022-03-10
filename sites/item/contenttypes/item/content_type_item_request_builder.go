package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
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
    ic5b939ddec677078eada74ae96425a6dd8ae52980fe6a4ce4ef57a15ff2c7be8 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/basetypes/item"
    ie3518f6b63ce04526457403d54384566536a4de177b91d352bd57d81171f4cd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/contenttypes/item/columnpositions/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.site entity.
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
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types defined for this site.
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i0eae2228a06806e82320cf19a03d576b42621cc4a2dd8806ce532a57bfebb362.AssociateWithHubSitesRequestBuilder) {
    return i0eae2228a06806e82320cf19a03d576b42621cc4a2dd8806ce532a57bfebb362.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*i7e8f4985115eafb5e1b9a5dcfae2a3c836415e064af553f7801cdbb55757ca11.BaseRequestBuilder) {
    return i7e8f4985115eafb5e1b9a5dcfae2a3c836415e064af553f7801cdbb55757ca11.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i7e506eefecd6b2935e76b82d621a57aa9ed417d501afb7d19965f3434a172010.BaseTypesRequestBuilder) {
    return i7e506eefecd6b2935e76b82d621a57aa9ed417d501afb7d19965f3434a172010.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*ic5b939ddec677078eada74ae96425a6dd8ae52980fe6a4ce4ef57a15ff2c7be8.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return ic5b939ddec677078eada74ae96425a6dd8ae52980fe6a4ce4ef57a15ff2c7be8.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i105f5ec8936f20cc343ef53f19692d9484bb0e2317204f4e0a83cb86af419abf.ColumnLinksRequestBuilder) {
    return i105f5ec8936f20cc343ef53f19692d9484bb0e2317204f4e0a83cb86af419abf.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i88d2724fa2dada76f95cc3b7d1c2d9bd0f61469930ea96ff9d0dc112126c5a8e.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i88d2724fa2dada76f95cc3b7d1c2d9bd0f61469930ea96ff9d0dc112126c5a8e.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i6d9ef8ddda3a3e7ebcb0a0d568b23e622439e98ff5309351bc26797ae8d500cb.ColumnPositionsRequestBuilder) {
    return i6d9ef8ddda3a3e7ebcb0a0d568b23e622439e98ff5309351bc26797ae8d500cb.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*ie3518f6b63ce04526457403d54384566536a4de177b91d352bd57d81171f4cd7.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ie3518f6b63ce04526457403d54384566536a4de177b91d352bd57d81171f4cd7.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*ifd4412476ff102f7c92f4e4a139b13990d10cc2647ddfa6e976c2dd5a98653ce.ColumnsRequestBuilder) {
    return ifd4412476ff102f7c92f4e4a139b13990d10cc2647ddfa6e976c2dd5a98653ce.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i9e8f3bdc2a2da10dbbb9f032cac127a9f38b1132156d876cec226651df880b26.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i9e8f3bdc2a2da10dbbb9f032cac127a9f38b1132156d876cec226651df880b26.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i910a70f3e4a25581ff8d025ecc5ab7dd183edf0f6e7af5d6e64e49e265867068.CopyToDefaultContentLocationRequestBuilder) {
    return i910a70f3e4a25581ff8d025ecc5ab7dd183edf0f6e7af5d6e64e49e265867068.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for sites
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
// CreateGetRequestInformation the collection of content types defined for this site.
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
// CreatePatchRequestInformation update the navigation property contentTypes in sites
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
// Delete delete navigation property contentTypes for sites
func (m *ContentTypeItemRequestBuilder) Delete(options *ContentTypeItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types defined for this site.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i943b393012430b44d1cf499bc2d8e50d28fa67a47c6be6689e906105bd5eb4e3.IsPublishedRequestBuilder) {
    return i943b393012430b44d1cf499bc2d8e50d28fa67a47c6be6689e906105bd5eb4e3.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in sites
func (m *ContentTypeItemRequestBuilder) Patch(options *ContentTypeItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeItemRequestBuilder) Publish()(*if8b488d84d1b1415d1c88a73424e76e4a1a1efd4526e5164d0604809ea5c8d13.PublishRequestBuilder) {
    return if8b488d84d1b1415d1c88a73424e76e4a1a1efd4526e5164d0604809ea5c8d13.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i20d9f259a492016a967dd6a10e660b98616fd7d4888ef6ee8111990cf5517382.UnpublishRequestBuilder) {
    return i20d9f259a492016a967dd6a10e660b98616fd7d4888ef6ee8111990cf5517382.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
