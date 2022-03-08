package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4d2cb6355c05070b43e771667e3c932735e55500af7d9008d0667dc15158dbfd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columnlinks"
    i584a99730fbd2ece5c6bfeeaf5256d826a03e37884ca6438f650d5236ca3485b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/ispublished"
    i6158e8fe936ce7a9232caf8ccfc3d8619a86e3d9493d7860a852c5b13e90ab25 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/basetypes"
    i854362c6353a95e813ebb2720a7c790bfb67cfb95b830527c59f7966139d71c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/associatewithhubsites"
    i97c73e8bccb2aeec40e8c91eaf704c572ef18478ca0f71947d7c3b651d907ea4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    i9def72881302efcfa7f2e3cc87905ee20db183df30ca0f6e13d8da1a294d04e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/publish"
    ic2b61d41ac59272716a65a67f10b958477f6f42c3287e4324c941b6257334a50 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columnpositions"
    idf8dd2eb707934e76443176905248e470c89a626df0ff62a74cc94ad0b884239 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columns"
    ie8f8fe074cdc7e704e1945d78071313b03611e728052a551ae404ac364ab9d82 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/base"
    ieb6b07f19a05128118ad274cf509ece9882db22095de3b1dc2bec37b991bfc9a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/unpublish"
    i0180fc6328d895b6a5187ca5c60771e42333dc3b52a00fcfd8aba46845553404 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columns/item"
    i63c74afb1978caaa67f675dda58750a522b0c0ed5def8dee9adfe882940e58ff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columnlinks/item"
    i9de1ac96113c37c6db0105ac34c7abae8b4de298ef6731e0d312197e57b7706a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/basetypes/item"
    ie43a627d2d2c6d4d2ad9648e748314d9e7fddb4990ac1ada48f7463ff2422d00 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/list/contenttypes/item/columnpositions/item"
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i854362c6353a95e813ebb2720a7c790bfb67cfb95b830527c59f7966139d71c3.AssociateWithHubSitesRequestBuilder) {
    return i854362c6353a95e813ebb2720a7c790bfb67cfb95b830527c59f7966139d71c3.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*ie8f8fe074cdc7e704e1945d78071313b03611e728052a551ae404ac364ab9d82.BaseRequestBuilder) {
    return ie8f8fe074cdc7e704e1945d78071313b03611e728052a551ae404ac364ab9d82.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i6158e8fe936ce7a9232caf8ccfc3d8619a86e3d9493d7860a852c5b13e90ab25.BaseTypesRequestBuilder) {
    return i6158e8fe936ce7a9232caf8ccfc3d8619a86e3d9493d7860a852c5b13e90ab25.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i9de1ac96113c37c6db0105ac34c7abae8b4de298ef6731e0d312197e57b7706a.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return i9de1ac96113c37c6db0105ac34c7abae8b4de298ef6731e0d312197e57b7706a.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i4d2cb6355c05070b43e771667e3c932735e55500af7d9008d0667dc15158dbfd.ColumnLinksRequestBuilder) {
    return i4d2cb6355c05070b43e771667e3c932735e55500af7d9008d0667dc15158dbfd.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i63c74afb1978caaa67f675dda58750a522b0c0ed5def8dee9adfe882940e58ff.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i63c74afb1978caaa67f675dda58750a522b0c0ed5def8dee9adfe882940e58ff.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*ic2b61d41ac59272716a65a67f10b958477f6f42c3287e4324c941b6257334a50.ColumnPositionsRequestBuilder) {
    return ic2b61d41ac59272716a65a67f10b958477f6f42c3287e4324c941b6257334a50.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*ie43a627d2d2c6d4d2ad9648e748314d9e7fddb4990ac1ada48f7463ff2422d00.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ie43a627d2d2c6d4d2ad9648e748314d9e7fddb4990ac1ada48f7463ff2422d00.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*idf8dd2eb707934e76443176905248e470c89a626df0ff62a74cc94ad0b884239.ColumnsRequestBuilder) {
    return idf8dd2eb707934e76443176905248e470c89a626df0ff62a74cc94ad0b884239.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i0180fc6328d895b6a5187ca5c60771e42333dc3b52a00fcfd8aba46845553404.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i0180fc6328d895b6a5187ca5c60771e42333dc3b52a00fcfd8aba46845553404.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/list/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i97c73e8bccb2aeec40e8c91eaf704c572ef18478ca0f71947d7c3b651d907ea4.CopyToDefaultContentLocationRequestBuilder) {
    return i97c73e8bccb2aeec40e8c91eaf704c572ef18478ca0f71947d7c3b651d907ea4.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for users
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
// CreatePatchRequestInformation update the navigation property contentTypes in users
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
// Delete delete navigation property contentTypes for users
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i584a99730fbd2ece5c6bfeeaf5256d826a03e37884ca6438f650d5236ca3485b.IsPublishedRequestBuilder) {
    return i584a99730fbd2ece5c6bfeeaf5256d826a03e37884ca6438f650d5236ca3485b.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in users
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
func (m *ContentTypeItemRequestBuilder) Publish()(*i9def72881302efcfa7f2e3cc87905ee20db183df30ca0f6e13d8da1a294d04e1.PublishRequestBuilder) {
    return i9def72881302efcfa7f2e3cc87905ee20db183df30ca0f6e13d8da1a294d04e1.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*ieb6b07f19a05128118ad274cf509ece9882db22095de3b1dc2bec37b991bfc9a.UnpublishRequestBuilder) {
    return ieb6b07f19a05128118ad274cf509ece9882db22095de3b1dc2bec37b991bfc9a.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
