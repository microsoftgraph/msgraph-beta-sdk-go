package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i29849b7bdbfecd87e4bbd3df4a5e86328cd6ecaa3abc6305f327b6fbe08c2513 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools/item/classes"
    i73b5871e98e80de0617052da16b736c5337770c4ccdfb726277a00424b92fe7d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools/item/users"
    icfb6e32f97bd4235dc90841a8db94d6ae91f82157354db59ed54dd9f182a6096 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools/item/administrativeunit"
    i23a7fb5a1ffc4cf80cf349afab1b5f8cfb3d3a3a12e2e62e24930a8787b14069 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools/item/classes/item"
    if8d8bcc13aaebd6e3ce9c01fcd6bb1063bdbebeded04ae846764acbc189974a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools/item/users/item"
)

// EducationSchoolItemRequestBuilder provides operations to manage the schools property of the microsoft.graph.educationRoot entity.
type EducationSchoolItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationSchoolItemRequestBuilderDeleteOptions options for Delete
type EducationSchoolItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EducationSchoolItemRequestBuilderGetOptions options for Get
type EducationSchoolItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *EducationSchoolItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EducationSchoolItemRequestBuilderGetQueryParameters get schools from education
type EducationSchoolItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationSchoolItemRequestBuilderPatchOptions options for Patch
type EducationSchoolItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSchoolable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AdministrativeUnit the administrativeUnit property
func (m *EducationSchoolItemRequestBuilder) AdministrativeUnit()(*icfb6e32f97bd4235dc90841a8db94d6ae91f82157354db59ed54dd9f182a6096.AdministrativeUnitRequestBuilder) {
    return icfb6e32f97bd4235dc90841a8db94d6ae91f82157354db59ed54dd9f182a6096.NewAdministrativeUnitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Classes the classes property
func (m *EducationSchoolItemRequestBuilder) Classes()(*i29849b7bdbfecd87e4bbd3df4a5e86328cd6ecaa3abc6305f327b6fbe08c2513.ClassesRequestBuilder) {
    return i29849b7bdbfecd87e4bbd3df4a5e86328cd6ecaa3abc6305f327b6fbe08c2513.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.schools.item.classes.item collection
func (m *EducationSchoolItemRequestBuilder) ClassesById(id string)(*i23a7fb5a1ffc4cf80cf349afab1b5f8cfb3d3a3a12e2e62e24930a8787b14069.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass_id"] = id
    }
    return i23a7fb5a1ffc4cf80cf349afab1b5f8cfb3d3a3a12e2e62e24930a8787b14069.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationSchoolItemRequestBuilderInternal instantiates a new EducationSchoolItemRequestBuilder and sets the default values.
func NewEducationSchoolItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSchoolItemRequestBuilder) {
    m := &EducationSchoolItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/schools/{educationSchool_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationSchoolItemRequestBuilder instantiates a new EducationSchoolItemRequestBuilder and sets the default values.
func NewEducationSchoolItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSchoolItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSchoolItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property schools for education
func (m *EducationSchoolItemRequestBuilder) CreateDeleteRequestInformation(options *EducationSchoolItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get schools from education
func (m *EducationSchoolItemRequestBuilder) CreateGetRequestInformation(options *EducationSchoolItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property schools in education
func (m *EducationSchoolItemRequestBuilder) CreatePatchRequestInformation(options *EducationSchoolItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property schools for education
func (m *EducationSchoolItemRequestBuilder) Delete(options *EducationSchoolItemRequestBuilderDeleteOptions)(error) {
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
// Get get schools from education
func (m *EducationSchoolItemRequestBuilder) Get(options *EducationSchoolItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSchoolable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSchoolFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSchoolable), nil
}
// Patch update the navigation property schools in education
func (m *EducationSchoolItemRequestBuilder) Patch(options *EducationSchoolItemRequestBuilderPatchOptions)(error) {
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
// Users the users property
func (m *EducationSchoolItemRequestBuilder) Users()(*i73b5871e98e80de0617052da16b736c5337770c4ccdfb726277a00424b92fe7d.UsersRequestBuilder) {
    return i73b5871e98e80de0617052da16b736c5337770c4ccdfb726277a00424b92fe7d.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.schools.item.users.item collection
func (m *EducationSchoolItemRequestBuilder) UsersById(id string)(*if8d8bcc13aaebd6e3ce9c01fcd6bb1063bdbebeded04ae846764acbc189974a5.EducationUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationUser_id"] = id
    }
    return if8d8bcc13aaebd6e3ce9c01fcd6bb1063bdbebeded04ae846764acbc189974a5.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
