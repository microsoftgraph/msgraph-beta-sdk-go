package previousversiondefinition

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4f3a86479d34cc32eaf91379a13629e71edf7037f7872864bcaeb9786910219f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/nextversiondefinition/previousversiondefinition/definitionfile"
    i7169e4b822d16779856f45354e143287fa3065cfbaef25c4c9e8f5aef3d4021c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/nextversiondefinition/previousversiondefinition/category"
    ibc555e9ae1332724d8a8c3dc23ae92e0e4de83dca6c2fc964ac483e2859fc119 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/nextversiondefinition/previousversiondefinition/presentations"
    i442f1022f6c43666754932bce9fea28984732b5c9032691588101f7054b95f7e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/nextversiondefinition/previousversiondefinition/presentations/item"
)

// PreviousVersionDefinitionRequestBuilder provides operations to manage the previousVersionDefinition property of the microsoft.graph.groupPolicyDefinition entity.
type PreviousVersionDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PreviousVersionDefinitionRequestBuilderDeleteOptions options for Delete
type PreviousVersionDefinitionRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// PreviousVersionDefinitionRequestBuilderGetOptions options for Get
type PreviousVersionDefinitionRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PreviousVersionDefinitionRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// PreviousVersionDefinitionRequestBuilderGetQueryParameters definition of the previous version of this definition
type PreviousVersionDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PreviousVersionDefinitionRequestBuilderPatchOptions options for Patch
type PreviousVersionDefinitionRequestBuilderPatchOptions struct {
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
func (m *PreviousVersionDefinitionRequestBuilder) Category()(*i7169e4b822d16779856f45354e143287fa3065cfbaef25c4c9e8f5aef3d4021c.CategoryRequestBuilder) {
    return i7169e4b822d16779856f45354e143287fa3065cfbaef25c4c9e8f5aef3d4021c.NewCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPreviousVersionDefinitionRequestBuilderInternal instantiates a new PreviousVersionDefinitionRequestBuilder and sets the default values.
func NewPreviousVersionDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreviousVersionDefinitionRequestBuilder) {
    m := &PreviousVersionDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/nextVersionDefinition/previousVersionDefinition{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPreviousVersionDefinitionRequestBuilder instantiates a new PreviousVersionDefinitionRequestBuilder and sets the default values.
func NewPreviousVersionDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreviousVersionDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPreviousVersionDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property previousVersionDefinition for deviceManagement
func (m *PreviousVersionDefinitionRequestBuilder) CreateDeleteRequestInformation(options *PreviousVersionDefinitionRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation definition of the previous version of this definition
func (m *PreviousVersionDefinitionRequestBuilder) CreateGetRequestInformation(options *PreviousVersionDefinitionRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property previousVersionDefinition in deviceManagement
func (m *PreviousVersionDefinitionRequestBuilder) CreatePatchRequestInformation(options *PreviousVersionDefinitionRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PreviousVersionDefinitionRequestBuilder) DefinitionFile()(*i4f3a86479d34cc32eaf91379a13629e71edf7037f7872864bcaeb9786910219f.DefinitionFileRequestBuilder) {
    return i4f3a86479d34cc32eaf91379a13629e71edf7037f7872864bcaeb9786910219f.NewDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property previousVersionDefinition for deviceManagement
func (m *PreviousVersionDefinitionRequestBuilder) Delete(options *PreviousVersionDefinitionRequestBuilderDeleteOptions)(error) {
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
// Get definition of the previous version of this definition
func (m *PreviousVersionDefinitionRequestBuilder) Get(options *PreviousVersionDefinitionRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
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
// Patch update the navigation property previousVersionDefinition in deviceManagement
func (m *PreviousVersionDefinitionRequestBuilder) Patch(options *PreviousVersionDefinitionRequestBuilderPatchOptions)(error) {
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
func (m *PreviousVersionDefinitionRequestBuilder) Presentations()(*ibc555e9ae1332724d8a8c3dc23ae92e0e4de83dca6c2fc964ac483e2859fc119.PresentationsRequestBuilder) {
    return ibc555e9ae1332724d8a8c3dc23ae92e0e4de83dca6c2fc964ac483e2859fc119.NewPresentationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresentationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyDefinitions.item.nextVersionDefinition.previousVersionDefinition.presentations.item collection
func (m *PreviousVersionDefinitionRequestBuilder) PresentationsById(id string)(*i442f1022f6c43666754932bce9fea28984732b5c9032691588101f7054b95f7e.GroupPolicyPresentationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyPresentation%2Did"] = id
    }
    return i442f1022f6c43666754932bce9fea28984732b5c9032691588101f7054b95f7e.NewGroupPolicyPresentationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
