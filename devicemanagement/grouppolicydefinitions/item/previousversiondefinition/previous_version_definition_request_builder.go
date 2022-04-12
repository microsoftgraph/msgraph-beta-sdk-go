package previousversiondefinition

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i6410495101249363dc0f95e5b3e0c72dae0fb79f298b963eeb94e762f2229ec8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/category"
    i74328421e57816d37bc08d88c5a29578e3d131e71500cd6303c5ebfd8ad7a1f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/presentations"
    i76dcada0ec02bd3dd3b1c880caaf69fd97aab0007527200a2cb47fc6fd50d460 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/nextversiondefinition"
    iec10cf7d92cea4a05bd61c522f717307dbdc7c7f04798646f228a85f3200e9d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/definitionfile"
    i9d1674ace1a7d9c480f97ff9a7222067b7470543dff13a236c3eea16a5b144ea "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicydefinitions/item/previousversiondefinition/presentations/item"
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
func (m *PreviousVersionDefinitionRequestBuilder) Category()(*i6410495101249363dc0f95e5b3e0c72dae0fb79f298b963eeb94e762f2229ec8.CategoryRequestBuilder) {
    return i6410495101249363dc0f95e5b3e0c72dae0fb79f298b963eeb94e762f2229ec8.NewCategoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPreviousVersionDefinitionRequestBuilderInternal instantiates a new PreviousVersionDefinitionRequestBuilder and sets the default values.
func NewPreviousVersionDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PreviousVersionDefinitionRequestBuilder) {
    m := &PreviousVersionDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyDefinitions/{groupPolicyDefinition%2Did}/previousVersionDefinition{?%24select,%24expand}";
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
func (m *PreviousVersionDefinitionRequestBuilder) DefinitionFile()(*iec10cf7d92cea4a05bd61c522f717307dbdc7c7f04798646f228a85f3200e9d6.DefinitionFileRequestBuilder) {
    return iec10cf7d92cea4a05bd61c522f717307dbdc7c7f04798646f228a85f3200e9d6.NewDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// NextVersionDefinition the nextVersionDefinition property
func (m *PreviousVersionDefinitionRequestBuilder) NextVersionDefinition()(*i76dcada0ec02bd3dd3b1c880caaf69fd97aab0007527200a2cb47fc6fd50d460.NextVersionDefinitionRequestBuilder) {
    return i76dcada0ec02bd3dd3b1c880caaf69fd97aab0007527200a2cb47fc6fd50d460.NewNextVersionDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *PreviousVersionDefinitionRequestBuilder) Presentations()(*i74328421e57816d37bc08d88c5a29578e3d131e71500cd6303c5ebfd8ad7a1f1.PresentationsRequestBuilder) {
    return i74328421e57816d37bc08d88c5a29578e3d131e71500cd6303c5ebfd8ad7a1f1.NewPresentationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresentationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyDefinitions.item.previousVersionDefinition.presentations.item collection
func (m *PreviousVersionDefinitionRequestBuilder) PresentationsById(id string)(*i9d1674ace1a7d9c480f97ff9a7222067b7470543dff13a236c3eea16a5b144ea.GroupPolicyPresentationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyPresentation%2Did"] = id
    }
    return i9d1674ace1a7d9c480f97ff9a7222067b7470543dff13a236c3eea16a5b144ea.NewGroupPolicyPresentationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
