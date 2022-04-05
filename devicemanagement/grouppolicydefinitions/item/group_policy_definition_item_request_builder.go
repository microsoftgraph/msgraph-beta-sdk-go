package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0084c64e69f2de1db99e3db4eebd9d8a41444eb2f0f9d19cc03286f0b896875a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/nextversiondefinition"
    i91fccd52e0d78587541f985ebb0b59be750da1e83e56b1da10b05e4eb4dc73ac "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/definitionfile"
    ia5547c0682358c94012164a326f029053ec5f1ddfbb91573107909f094b955bc "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/presentations"
    ib5e5e1e82d57aaa3cb2752c79932c323ee118f915bbc0dc5e5ab7600a66a7f27 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/category"
    ide1c03ea82181c831c7c8a172c8a94f8d4edd80cc28a64d3a11e7c2c58c4ab3e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition"
    i755dd99b2492e3a413ac1208e4de2ede29dfb284de02d498fb55c3247fe139f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/presentations/item"
)

// GroupPolicyDefinitionItemRequestBuilder provides operations to manage the groupPolicyDefinitions property of the microsoft.graph.deviceManagement entity.
type GroupPolicyDefinitionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupPolicyDefinitionItemRequestBuilderDeleteOptions options for Delete
type GroupPolicyDefinitionItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// GroupPolicyDefinitionItemRequestBuilderGetOptions options for Get
type GroupPolicyDefinitionItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *GroupPolicyDefinitionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// GroupPolicyDefinitionItemRequestBuilderGetQueryParameters the available group policy definitions for this account.
type GroupPolicyDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GroupPolicyDefinitionItemRequestBuilderPatchOptions options for Patch
type GroupPolicyDefinitionItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Category the category property
func (m *GroupPolicyDefinitionItemRequestBuilder) Category()(*ib5e5e1e82d57aaa3cb2752c79932c323ee118f915bbc0dc5e5ab7600a66a7f27.CategoryRequestBuilder) {
    return ib5e5e1e82d57aaa3cb2752c79932c323ee118f915bbc0dc5e5ab7600a66a7f27.NewCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupPolicyDefinitionItemRequestBuilderInternal instantiates a new GroupPolicyDefinitionItemRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionItemRequestBuilder) {
    m := &GroupPolicyDefinitionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyDefinitionItemRequestBuilder instantiates a new GroupPolicyDefinitionItemRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groupPolicyDefinitions for deviceManagement
func (m *GroupPolicyDefinitionItemRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyDefinitionItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the available group policy definitions for this account.
func (m *GroupPolicyDefinitionItemRequestBuilder) CreateGetRequestInformation(options *GroupPolicyDefinitionItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property groupPolicyDefinitions in deviceManagement
func (m *GroupPolicyDefinitionItemRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyDefinitionItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DefinitionFile the definitionFile property
func (m *GroupPolicyDefinitionItemRequestBuilder) DefinitionFile()(*i91fccd52e0d78587541f985ebb0b59be750da1e83e56b1da10b05e4eb4dc73ac.DefinitionFileRequestBuilder) {
    return i91fccd52e0d78587541f985ebb0b59be750da1e83e56b1da10b05e4eb4dc73ac.NewDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property groupPolicyDefinitions for deviceManagement
func (m *GroupPolicyDefinitionItemRequestBuilder) Delete(options *GroupPolicyDefinitionItemRequestBuilderDeleteOptions)(error) {
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
// Get the available group policy definitions for this account.
func (m *GroupPolicyDefinitionItemRequestBuilder) Get(options *GroupPolicyDefinitionItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable), nil
}
// NextVersionDefinition the nextVersionDefinition property
func (m *GroupPolicyDefinitionItemRequestBuilder) NextVersionDefinition()(*i0084c64e69f2de1db99e3db4eebd9d8a41444eb2f0f9d19cc03286f0b896875a.NextVersionDefinitionRequestBuilder) {
    return i0084c64e69f2de1db99e3db4eebd9d8a41444eb2f0f9d19cc03286f0b896875a.NewNextVersionDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property groupPolicyDefinitions in deviceManagement
func (m *GroupPolicyDefinitionItemRequestBuilder) Patch(options *GroupPolicyDefinitionItemRequestBuilderPatchOptions)(error) {
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
// Presentations the presentations property
func (m *GroupPolicyDefinitionItemRequestBuilder) Presentations()(*ia5547c0682358c94012164a326f029053ec5f1ddfbb91573107909f094b955bc.PresentationsRequestBuilder) {
    return ia5547c0682358c94012164a326f029053ec5f1ddfbb91573107909f094b955bc.NewPresentationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresentationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyDefinitions.item.presentations.item collection
func (m *GroupPolicyDefinitionItemRequestBuilder) PresentationsById(id string)(*i755dd99b2492e3a413ac1208e4de2ede29dfb284de02d498fb55c3247fe139f0.GroupPolicyPresentationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyPresentation_id"] = id
    }
    return i755dd99b2492e3a413ac1208e4de2ede29dfb284de02d498fb55c3247fe139f0.NewGroupPolicyPresentationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PreviousVersionDefinition the previousVersionDefinition property
func (m *GroupPolicyDefinitionItemRequestBuilder) PreviousVersionDefinition()(*ide1c03ea82181c831c7c8a172c8a94f8d4edd80cc28a64d3a11e7c2c58c4ab3e.PreviousVersionDefinitionRequestBuilder) {
    return ide1c03ea82181c831c7c8a172c8a94f8d4edd80cc28a64d3a11e7c2c58c4ab3e.NewPreviousVersionDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
