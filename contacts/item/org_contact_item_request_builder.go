package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i1c4ae4453c2c0d51315d87389303abac00b2ee54b8eea32575bbc9542c97d724 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivereports"
    i3218d64770c5fb2ac858c46ce0b21eca8069f173345a45e2d0127524fe6244f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i8930ff877b98d2d829bb75e1bc731872e28123f14d7f4b26b94740f4457ce1f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/checkmembergroups"
    i8fabe20041a357bb8ca64ddefe7e31da4ced67bdca403d7afda823be6b34b940 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/getmembergroups"
    iaaabcbb8c60e122516c587d50e74e474d377c15f6aa8bfc4e7dc2b333a641821 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/checkmemberobjects"
    ic3d1f9a92a97f0f59be23d433068e16bff96fc111767647a70509ef431c85678 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/manager"
    id9d85e0537a55572b46a9bf049286ac01bd7daf191686fb4bf68f252f6cbd91a "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/directreports"
    ie27f01d64cec3543d1978d296394fa58b3dc2e67d00d36ab3e99c2e4d2f0c223 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/getmemberobjects"
    iea26a03fcd10d6c2d94266d5351c1957bf25a3b77f638458f322bc7464fb7e24 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/restore"
    if78bb81960e1b033a5e4bd2fd9dc5ad58550b0ce6b418788078e981aa388315d "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i46940b4c384df6665b2c83dbbf830182b33c3cc9604ee88b827b88500a2821a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/directreports/item"
    i685f30515a7f1399908cd02e9a255c21fb09379803dcc80c4440037dc4aa822e "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/item"
    ib4f564b6a52f86e74bf28b16db5ffd60326a6a4a07cee512fc66c526d0aeb45d "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/item"
    if03b6865af61bfce396b43696c04c2e15f4e54440b1d01ba87692c9dcca4b7eb "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivereports/item"
)

// OrgContactItemRequestBuilder provides operations to manage the collection of orgContact entities.
type OrgContactItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OrgContactItemRequestBuilderDeleteOptions options for Delete
type OrgContactItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OrgContactItemRequestBuilderGetOptions options for Get
type OrgContactItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OrgContactItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OrgContactItemRequestBuilderGetQueryParameters get entity from contacts by key
type OrgContactItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OrgContactItemRequestBuilderPatchOptions options for Patch
type OrgContactItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContactable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *OrgContactItemRequestBuilder) CheckMemberGroups()(*i8930ff877b98d2d829bb75e1bc731872e28123f14d7f4b26b94740f4457ce1f0.CheckMemberGroupsRequestBuilder) {
    return i8930ff877b98d2d829bb75e1bc731872e28123f14d7f4b26b94740f4457ce1f0.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactItemRequestBuilder) CheckMemberObjects()(*iaaabcbb8c60e122516c587d50e74e474d377c15f6aa8bfc4e7dc2b333a641821.CheckMemberObjectsRequestBuilder) {
    return iaaabcbb8c60e122516c587d50e74e474d377c15f6aa8bfc4e7dc2b333a641821.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrgContactItemRequestBuilderInternal instantiates a new OrgContactItemRequestBuilder and sets the default values.
func NewOrgContactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrgContactItemRequestBuilder) {
    m := &OrgContactItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrgContactItemRequestBuilder instantiates a new OrgContactItemRequestBuilder and sets the default values.
func NewOrgContactItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OrgContactItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrgContactItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from contacts
func (m *OrgContactItemRequestBuilder) CreateDeleteRequestInformation(options *OrgContactItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from contacts by key
func (m *OrgContactItemRequestBuilder) CreateGetRequestInformation(options *OrgContactItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in contacts
func (m *OrgContactItemRequestBuilder) CreatePatchRequestInformation(options *OrgContactItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from contacts
func (m *OrgContactItemRequestBuilder) Delete(options *OrgContactItemRequestBuilderDeleteOptions)(error) {
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
func (m *OrgContactItemRequestBuilder) DirectReports()(*id9d85e0537a55572b46a9bf049286ac01bd7daf191686fb4bf68f252f6cbd91a.DirectReportsRequestBuilder) {
    return id9d85e0537a55572b46a9bf049286ac01bd7daf191686fb4bf68f252f6cbd91a.NewDirectReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.contacts.item.directReports.item collection
func (m *OrgContactItemRequestBuilder) DirectReportsById(id string)(*i46940b4c384df6665b2c83dbbf830182b33c3cc9604ee88b827b88500a2821a7.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i46940b4c384df6665b2c83dbbf830182b33c3cc9604ee88b827b88500a2821a7.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from contacts by key
func (m *OrgContactItemRequestBuilder) Get(options *OrgContactItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContactable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOrgContactFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OrgContactable), nil
}
func (m *OrgContactItemRequestBuilder) GetMemberGroups()(*i8fabe20041a357bb8ca64ddefe7e31da4ced67bdca403d7afda823be6b34b940.GetMemberGroupsRequestBuilder) {
    return i8fabe20041a357bb8ca64ddefe7e31da4ced67bdca403d7afda823be6b34b940.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactItemRequestBuilder) GetMemberObjects()(*ie27f01d64cec3543d1978d296394fa58b3dc2e67d00d36ab3e99c2e4d2f0c223.GetMemberObjectsRequestBuilder) {
    return ie27f01d64cec3543d1978d296394fa58b3dc2e67d00d36ab3e99c2e4d2f0c223.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactItemRequestBuilder) Manager()(*ic3d1f9a92a97f0f59be23d433068e16bff96fc111767647a70509ef431c85678.ManagerRequestBuilder) {
    return ic3d1f9a92a97f0f59be23d433068e16bff96fc111767647a70509ef431c85678.NewManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactItemRequestBuilder) MemberOf()(*i3218d64770c5fb2ac858c46ce0b21eca8069f173345a45e2d0127524fe6244f5.MemberOfRequestBuilder) {
    return i3218d64770c5fb2ac858c46ce0b21eca8069f173345a45e2d0127524fe6244f5.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.contacts.item.memberOf.item collection
func (m *OrgContactItemRequestBuilder) MemberOfById(id string)(*ib4f564b6a52f86e74bf28b16db5ffd60326a6a4a07cee512fc66c526d0aeb45d.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ib4f564b6a52f86e74bf28b16db5ffd60326a6a4a07cee512fc66c526d0aeb45d.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in contacts
func (m *OrgContactItemRequestBuilder) Patch(options *OrgContactItemRequestBuilderPatchOptions)(error) {
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
func (m *OrgContactItemRequestBuilder) Restore()(*iea26a03fcd10d6c2d94266d5351c1957bf25a3b77f638458f322bc7464fb7e24.RestoreRequestBuilder) {
    return iea26a03fcd10d6c2d94266d5351c1957bf25a3b77f638458f322bc7464fb7e24.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OrgContactItemRequestBuilder) TransitiveMemberOf()(*if78bb81960e1b033a5e4bd2fd9dc5ad58550b0ce6b418788078e981aa388315d.TransitiveMemberOfRequestBuilder) {
    return if78bb81960e1b033a5e4bd2fd9dc5ad58550b0ce6b418788078e981aa388315d.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.contacts.item.transitiveMemberOf.item collection
func (m *OrgContactItemRequestBuilder) TransitiveMemberOfById(id string)(*i685f30515a7f1399908cd02e9a255c21fb09379803dcc80c4440037dc4aa822e.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i685f30515a7f1399908cd02e9a255c21fb09379803dcc80c4440037dc4aa822e.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OrgContactItemRequestBuilder) TransitiveReports()(*i1c4ae4453c2c0d51315d87389303abac00b2ee54b8eea32575bbc9542c97d724.TransitiveReportsRequestBuilder) {
    return i1c4ae4453c2c0d51315d87389303abac00b2ee54b8eea32575bbc9542c97d724.NewTransitiveReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveReportsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.contacts.item.transitiveReports.item collection
func (m *OrgContactItemRequestBuilder) TransitiveReportsById(id string)(*if03b6865af61bfce396b43696c04c2e15f4e54440b1d01ba87692c9dcca4b7eb.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return if03b6865af61bfce396b43696c04c2e15f4e54440b1d01ba87692c9dcca4b7eb.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
