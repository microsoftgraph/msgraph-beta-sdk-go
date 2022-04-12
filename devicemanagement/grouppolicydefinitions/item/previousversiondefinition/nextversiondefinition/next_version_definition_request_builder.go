package nextversiondefinition

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4d4e245af6721bc39b0ed5b42130efdfd5a1fc3654c3bd63cfbabf940a72ebb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/nextversiondefinition/definitionfile"
    i8c068a336e694856444efa8c9530bbfab994d44a29f10414147eea9838587f10 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/nextversiondefinition/presentations"
    icd163ce32e48cec41a54cf51929022d155cff3b928a52ef12e2c4fce1bf256aa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/nextversiondefinition/category"
    i116532920a384b7a633bc588077a2000112ae5f7dfa6fb6cf05bdc16cb9610a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/nextversiondefinition/presentations/item"
)

// NextVersionDefinitionRequestBuilder provides operations to manage the nextVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
type NextVersionDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NextVersionDefinitionRequestBuilderDeleteOptions options for Delete
type NextVersionDefinitionRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NextVersionDefinitionRequestBuilderGetOptions options for Get
type NextVersionDefinitionRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *NextVersionDefinitionRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NextVersionDefinitionRequestBuilderGetQueryParameters definition of the next version of this definition
type NextVersionDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// NextVersionDefinitionRequestBuilderPatchOptions options for Patch
type NextVersionDefinitionRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Category the category property
func (m *NextVersionDefinitionRequestBuilder) Category()(*icd163ce32e48cec41a54cf51929022d155cff3b928a52ef12e2c4fce1bf256aa.CategoryRequestBuilder) {
    return icd163ce32e48cec41a54cf51929022d155cff3b928a52ef12e2c4fce1bf256aa.NewCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewNextVersionDefinitionRequestBuilderInternal instantiates a new NextVersionDefinitionRequestBuilder and sets the default values.
func NewNextVersionDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NextVersionDefinitionRequestBuilder) {
    m := &NextVersionDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/previousVersionDefinition/nextVersionDefinition{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNextVersionDefinitionRequestBuilder instantiates a new NextVersionDefinitionRequestBuilder and sets the default values.
func NewNextVersionDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*NextVersionDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNextVersionDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property nextVersionDefinition for deviceManagement
func (m *NextVersionDefinitionRequestBuilder) CreateDeleteRequestInformation(options *NextVersionDefinitionRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation definition of the next version of this definition
func (m *NextVersionDefinitionRequestBuilder) CreateGetRequestInformation(options *NextVersionDefinitionRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property nextVersionDefinition in deviceManagement
func (m *NextVersionDefinitionRequestBuilder) CreatePatchRequestInformation(options *NextVersionDefinitionRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *NextVersionDefinitionRequestBuilder) DefinitionFile()(*i4d4e245af6721bc39b0ed5b42130efdfd5a1fc3654c3bd63cfbabf940a72ebb3.DefinitionFileRequestBuilder) {
    return i4d4e245af6721bc39b0ed5b42130efdfd5a1fc3654c3bd63cfbabf940a72ebb3.NewDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property nextVersionDefinition for deviceManagement
func (m *NextVersionDefinitionRequestBuilder) Delete(options *NextVersionDefinitionRequestBuilderDeleteOptions)(error) {
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
// Get definition of the next version of this definition
func (m *NextVersionDefinitionRequestBuilder) Get(options *NextVersionDefinitionRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
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
// Patch update the navigation property nextVersionDefinition in deviceManagement
func (m *NextVersionDefinitionRequestBuilder) Patch(options *NextVersionDefinitionRequestBuilderPatchOptions)(error) {
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
func (m *NextVersionDefinitionRequestBuilder) Presentations()(*i8c068a336e694856444efa8c9530bbfab994d44a29f10414147eea9838587f10.PresentationsRequestBuilder) {
    return i8c068a336e694856444efa8c9530bbfab994d44a29f10414147eea9838587f10.NewPresentationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresentationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyDefinitions.item.previousVersionDefinition.nextVersionDefinition.presentations.item collection
func (m *NextVersionDefinitionRequestBuilder) PresentationsById(id string)(*i116532920a384b7a633bc588077a2000112ae5f7dfa6fb6cf05bdc16cb9610a9.GroupPolicyPresentationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyPresentation%2Did"] = id
    }
    return i116532920a384b7a633bc588077a2000112ae5f7dfa6fb6cf05bdc16cb9610a9.NewGroupPolicyPresentationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
