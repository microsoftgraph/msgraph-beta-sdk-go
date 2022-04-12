package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1c4ae4453c2c0d51315d87389303abac00b2ee54b8eea32575bbc9542c97d724 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivereports"
    i3218d64770c5fb2ac858c46ce0b21eca8069f173345a45e2d0127524fe6244f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof"
    i8930ff877b98d2d829bb75e1bc731872e28123f14d7f4b26b94740f4457ce1f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/checkmembergroups"
    i8fabe20041a357bb8ca64ddefe7e31da4ced67bdca403d7afda823be6b34b940 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/getmembergroups"
    iaaabcbb8c60e122516c587d50e74e474d377c15f6aa8bfc4e7dc2b333a641821 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/checkmemberobjects"
    ic3d1f9a92a97f0f59be23d433068e16bff96fc111767647a70509ef431c85678 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/manager"
    id9d85e0537a55572b46a9bf049286ac01bd7daf191686fb4bf68f252f6cbd91a "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/directreports"
    ie27f01d64cec3543d1978d296394fa58b3dc2e67d00d36ab3e99c2e4d2f0c223 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/getmemberobjects"
    iea26a03fcd10d6c2d94266d5351c1957bf25a3b77f638458f322bc7464fb7e24 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/restore"
    if78bb81960e1b033a5e4bd2fd9dc5ad58550b0ce6b418788078e981aa388315d "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof"
    i46940b4c384df6665b2c83dbbf830182b33c3cc9604ee88b827b88500a2821a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/directreports/item"
    i685f30515a7f1399908cd02e9a255c21fb09379803dcc80c4440037dc4aa822e "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivememberof/item"
    ib4f564b6a52f86e74bf28b16db5ffd60326a6a4a07cee512fc66c526d0aeb45d "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/memberof/item"
    if03b6865af61bfce396b43696c04c2e15f4e54440b1d01ba87692c9dcca4b7eb "github.com/microsoftgraph/msgraph-beta-sdk-go/contacts/item/transitivereports/item"
)

// OrgContactItemRequestBuilder provides operations to manage the collection of orgContact entities.
type OrgContactItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrgContactItemRequestBuilderDeleteOptions options for Delete
type OrgContactItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// OrgContactItemRequestBuilderGetOptions options for Get
type OrgContactItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OrgContactItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// OrgContactItemRequestBuilderGetQueryParameters get entity from contacts by key
type OrgContactItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OrgContactItemRequestBuilderPatchOptions options for Patch
type OrgContactItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CheckMemberGroups the checkMemberGroups property
func (m *OrgContactItemRequestBuilder) CheckMemberGroups()(*i8930ff877b98d2d829bb75e1bc731872e28123f14d7f4b26b94740f4457ce1f0.CheckMemberGroupsRequestBuilder) {
    return i8930ff877b98d2d829bb75e1bc731872e28123f14d7f4b26b94740f4457ce1f0.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *OrgContactItemRequestBuilder) CheckMemberObjects()(*iaaabcbb8c60e122516c587d50e74e474d377c15f6aa8bfc4e7dc2b333a641821.CheckMemberObjectsRequestBuilder) {
    return iaaabcbb8c60e122516c587d50e74e474d377c15f6aa8bfc4e7dc2b333a641821.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewOrgContactItemRequestBuilderInternal instantiates a new OrgContactItemRequestBuilder and sets the default values.
func NewOrgContactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrgContactItemRequestBuilder) {
    m := &OrgContactItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOrgContactItemRequestBuilder instantiates a new OrgContactItemRequestBuilder and sets the default values.
func NewOrgContactItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrgContactItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrgContactItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from contacts
func (m *OrgContactItemRequestBuilder) CreateDeleteRequestInformation(options *OrgContactItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from contacts by key
func (m *OrgContactItemRequestBuilder) CreateGetRequestInformation(options *OrgContactItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in contacts
func (m *OrgContactItemRequestBuilder) CreatePatchRequestInformation(options *OrgContactItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from contacts
func (m *OrgContactItemRequestBuilder) Delete(options *OrgContactItemRequestBuilderDeleteOptions)(error) {
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
// DirectReports the directReports property
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
        urlTplParams["directoryObject%2Did"] = id
    }
    return i46940b4c384df6665b2c83dbbf830182b33c3cc9604ee88b827b88500a2821a7.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from contacts by key
func (m *OrgContactItemRequestBuilder) Get(options *OrgContactItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrgContactFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactable), nil
}
// GetMemberGroups the getMemberGroups property
func (m *OrgContactItemRequestBuilder) GetMemberGroups()(*i8fabe20041a357bb8ca64ddefe7e31da4ced67bdca403d7afda823be6b34b940.GetMemberGroupsRequestBuilder) {
    return i8fabe20041a357bb8ca64ddefe7e31da4ced67bdca403d7afda823be6b34b940.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *OrgContactItemRequestBuilder) GetMemberObjects()(*ie27f01d64cec3543d1978d296394fa58b3dc2e67d00d36ab3e99c2e4d2f0c223.GetMemberObjectsRequestBuilder) {
    return ie27f01d64cec3543d1978d296394fa58b3dc2e67d00d36ab3e99c2e4d2f0c223.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Manager the manager property
func (m *OrgContactItemRequestBuilder) Manager()(*ic3d1f9a92a97f0f59be23d433068e16bff96fc111767647a70509ef431c85678.ManagerRequestBuilder) {
    return ic3d1f9a92a97f0f59be23d433068e16bff96fc111767647a70509ef431c85678.NewManagerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOf the memberOf property
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
        urlTplParams["directoryObject%2Did"] = id
    }
    return ib4f564b6a52f86e74bf28b16db5ffd60326a6a4a07cee512fc66c526d0aeb45d.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in contacts
func (m *OrgContactItemRequestBuilder) Patch(options *OrgContactItemRequestBuilderPatchOptions)(error) {
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
// Restore the restore property
func (m *OrgContactItemRequestBuilder) Restore()(*iea26a03fcd10d6c2d94266d5351c1957bf25a3b77f638458f322bc7464fb7e24.RestoreRequestBuilder) {
    return iea26a03fcd10d6c2d94266d5351c1957bf25a3b77f638458f322bc7464fb7e24.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
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
        urlTplParams["directoryObject%2Did"] = id
    }
    return i685f30515a7f1399908cd02e9a255c21fb09379803dcc80c4440037dc4aa822e.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveReports the transitiveReports property
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
        urlTplParams["directoryObject%2Did"] = id
    }
    return if03b6865af61bfce396b43696c04c2e15f4e54440b1d01ba87692c9dcca4b7eb.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
