package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
    i250dd9b6cf0b441ae9ad6dd8b424bf0341fb39d6f13e1a8af39fd07070d35a31 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/parentgroup"
    ia3754d5d285507891626f5df6f89aa875778e56b68ce460f5b6e3fbb7a57e1fa "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/terms"
    ib68e940b5c91c0899ffae7e529e330e092c8c17f617049423242c363373ec567 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/children"
    ibe1ea46390b123f01791e91de5a75be01f76ff0077cac9a5a9a13ab8206c77c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/relations"
    i0d9b9b9cd2472b796e46b377506082941cbd9afdf632cf6216e386ac455aae38 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/children/item"
    i36d775ae30be636ca23bb6a064d3de5e1f1722461bdb24649dba586dfb4181ce "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/relations/item"
    id049073a70a9b38a6f8a94285b2dbcec1acb00a0e4dfd4863e659b573e369b55 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/termstore/sets/item/terms/item"
)

// SetItemRequestBuilder provides operations to manage the sets property of the microsoft.graph.termStore.store entity.
type SetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SetItemRequestBuilderDeleteOptions options for Delete
type SetItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SetItemRequestBuilderGetOptions options for Get
type SetItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SetItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SetItemRequestBuilderGetQueryParameters collection of all sets available in the term store.
type SetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SetItemRequestBuilderPatchOptions options for Patch
type SetItemRequestBuilderPatchOptions struct {
    // 
    Body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Children the children property
func (m *SetItemRequestBuilder) Children()(*ib68e940b5c91c0899ffae7e529e330e092c8c17f617049423242c363373ec567.ChildrenRequestBuilder) {
    return ib68e940b5c91c0899ffae7e529e330e092c8c17f617049423242c363373ec567.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.termStore.sets.item.children.item collection
func (m *SetItemRequestBuilder) ChildrenById(id string)(*i0d9b9b9cd2472b796e46b377506082941cbd9afdf632cf6216e386ac455aae38.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return i0d9b9b9cd2472b796e46b377506082941cbd9afdf632cf6216e386ac455aae38.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSetItemRequestBuilderInternal instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetItemRequestBuilder) {
    m := &SetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStore/sets/{set%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetItemRequestBuilder instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sets for sites
func (m *SetItemRequestBuilder) CreateDeleteRequestInformation(options *SetItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation collection of all sets available in the term store.
func (m *SetItemRequestBuilder) CreateGetRequestInformation(options *SetItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sets in sites
func (m *SetItemRequestBuilder) CreatePatchRequestInformation(options *SetItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property sets for sites
func (m *SetItemRequestBuilder) Delete(options *SetItemRequestBuilderDeleteOptions)(error) {
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
// Get collection of all sets available in the term store.
func (m *SetItemRequestBuilder) Get(options *SetItemRequestBuilderGetOptions)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateSetFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable), nil
}
// ParentGroup the parentGroup property
func (m *SetItemRequestBuilder) ParentGroup()(*i250dd9b6cf0b441ae9ad6dd8b424bf0341fb39d6f13e1a8af39fd07070d35a31.ParentGroupRequestBuilder) {
    return i250dd9b6cf0b441ae9ad6dd8b424bf0341fb39d6f13e1a8af39fd07070d35a31.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sets in sites
func (m *SetItemRequestBuilder) Patch(options *SetItemRequestBuilderPatchOptions)(error) {
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
// Relations the relations property
func (m *SetItemRequestBuilder) Relations()(*ibe1ea46390b123f01791e91de5a75be01f76ff0077cac9a5a9a13ab8206c77c7.RelationsRequestBuilder) {
    return ibe1ea46390b123f01791e91de5a75be01f76ff0077cac9a5a9a13ab8206c77c7.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.termStore.sets.item.relations.item collection
func (m *SetItemRequestBuilder) RelationsById(id string)(*i36d775ae30be636ca23bb6a064d3de5e1f1722461bdb24649dba586dfb4181ce.RelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return i36d775ae30be636ca23bb6a064d3de5e1f1722461bdb24649dba586dfb4181ce.NewRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Terms the terms property
func (m *SetItemRequestBuilder) Terms()(*ia3754d5d285507891626f5df6f89aa875778e56b68ce460f5b6e3fbb7a57e1fa.TermsRequestBuilder) {
    return ia3754d5d285507891626f5df6f89aa875778e56b68ce460f5b6e3fbb7a57e1fa.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.termStore.sets.item.terms.item collection
func (m *SetItemRequestBuilder) TermsById(id string)(*id049073a70a9b38a6f8a94285b2dbcec1acb00a0e4dfd4863e659b573e369b55.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return id049073a70a9b38a6f8a94285b2dbcec1acb00a0e4dfd4863e659b573e369b55.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
