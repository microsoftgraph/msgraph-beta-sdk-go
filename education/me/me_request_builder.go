package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i066e6c0b7a628617371533e62ce3ebdd3bb41009a8fbfe0051aebeee5a091763 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments"
    i1197007b396c8142142d808832d472fa4a0ca1796dcda8056a9c759ced2a4005 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/schools"
    i1b5284c067c2acfcb9ec377599aa90647d0e8c253bdeac21d56033eae704633e "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/rubrics"
    i5f1f4bf3151bbb3488a322f91d88a1e3277a2a2e390a895db3a6a5a0dafd3bd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/user"
    i5f4f009fc61fc7227f7443b57e61741a07864a630e39eba9bf7e097c583915e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/taughtclasses"
    ia9901f126435d74fc86b3bf99ec52d925f16943c99de2d04f1f3f010f4c84fbc "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/classes"
    i19868cee8abfa86080d68db45779c91f9b2670f8887c97f48ef60e07b22031c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/schools/item"
    i400d8aa2bba747bb59e46da289af76154496fa6c06cb9a12266b15362b1cfa5b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/taughtclasses/item"
    i45b49c5a37e86013b854ac973ddf755ed5024fc1a6c2dda426383115a09e7f90 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/classes/item"
    i719722077dd2c46836e8fc87facd02a34341429e3162704b105c41c28deda2b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/rubrics/item"
    ieba30bd3da2220b37fa7fcf6200b1473597f54cc0733c3310394e58e26492403 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item"
)

// MeRequestBuilder provides operations to manage the me property of the microsoft.graph.educationRoot entity.
type MeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeRequestBuilderGetQueryParameters get me from education
type MeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeRequestBuilderGetQueryParameters
}
// MeRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments the assignments property
func (m *MeRequestBuilder) Assignments()(*i066e6c0b7a628617371533e62ce3ebdd3bb41009a8fbfe0051aebeee5a091763.AssignmentsRequestBuilder) {
    return i066e6c0b7a628617371533e62ce3ebdd3bb41009a8fbfe0051aebeee5a091763.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item collection
func (m *MeRequestBuilder) AssignmentsById(id string)(*ieba30bd3da2220b37fa7fcf6200b1473597f54cc0733c3310394e58e26492403.EducationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignment%2Did"] = id
    }
    return ieba30bd3da2220b37fa7fcf6200b1473597f54cc0733c3310394e58e26492403.NewEducationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Classes the classes property
func (m *MeRequestBuilder) Classes()(*ia9901f126435d74fc86b3bf99ec52d925f16943c99de2d04f1f3f010f4c84fbc.ClassesRequestBuilder) {
    return ia9901f126435d74fc86b3bf99ec52d925f16943c99de2d04f1f3f010f4c84fbc.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.classes.item collection
func (m *MeRequestBuilder) ClassesById(id string)(*i45b49c5a37e86013b854ac973ddf755ed5024fc1a6c2dda426383115a09e7f90.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass%2Did"] = id
    }
    return i45b49c5a37e86013b854ac973ddf755ed5024fc1a6c2dda426383115a09e7f90.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMeRequestBuilderInternal instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeRequestBuilder) {
    m := &MeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeRequestBuilder instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property me for education
func (m *MeRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property me for education
func (m *MeRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *MeRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get me from education
func (m *MeRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get me from education
func (m *MeRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property me in education
func (m *MeRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationUserable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property me in education
func (m *MeRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationUserable, requestConfiguration *MeRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property me for education
func (m *MeRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get me from education
func (m *MeRequestBuilder) Get(ctx context.Context, requestConfiguration *MeRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationUserable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationUserable), nil
}
// Patch update the navigation property me in education
func (m *MeRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationUserable, requestConfiguration *MeRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationUserable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationUserable), nil
}
// Rubrics the rubrics property
func (m *MeRequestBuilder) Rubrics()(*i1b5284c067c2acfcb9ec377599aa90647d0e8c253bdeac21d56033eae704633e.RubricsRequestBuilder) {
    return i1b5284c067c2acfcb9ec377599aa90647d0e8c253bdeac21d56033eae704633e.NewRubricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RubricsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.rubrics.item collection
func (m *MeRequestBuilder) RubricsById(id string)(*i719722077dd2c46836e8fc87facd02a34341429e3162704b105c41c28deda2b9.EducationRubricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationRubric%2Did"] = id
    }
    return i719722077dd2c46836e8fc87facd02a34341429e3162704b105c41c28deda2b9.NewEducationRubricItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Schools the schools property
func (m *MeRequestBuilder) Schools()(*i1197007b396c8142142d808832d472fa4a0ca1796dcda8056a9c759ced2a4005.SchoolsRequestBuilder) {
    return i1197007b396c8142142d808832d472fa4a0ca1796dcda8056a9c759ced2a4005.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchoolsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.schools.item collection
func (m *MeRequestBuilder) SchoolsById(id string)(*i19868cee8abfa86080d68db45779c91f9b2670f8887c97f48ef60e07b22031c2.EducationSchoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSchool%2Did"] = id
    }
    return i19868cee8abfa86080d68db45779c91f9b2670f8887c97f48ef60e07b22031c2.NewEducationSchoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TaughtClasses the taughtClasses property
func (m *MeRequestBuilder) TaughtClasses()(*i5f4f009fc61fc7227f7443b57e61741a07864a630e39eba9bf7e097c583915e3.TaughtClassesRequestBuilder) {
    return i5f4f009fc61fc7227f7443b57e61741a07864a630e39eba9bf7e097c583915e3.NewTaughtClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaughtClassesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.taughtClasses.item collection
func (m *MeRequestBuilder) TaughtClassesById(id string)(*i400d8aa2bba747bb59e46da289af76154496fa6c06cb9a12266b15362b1cfa5b.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass%2Did"] = id
    }
    return i400d8aa2bba747bb59e46da289af76154496fa6c06cb9a12266b15362b1cfa5b.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// User the user property
func (m *MeRequestBuilder) User()(*i5f1f4bf3151bbb3488a322f91d88a1e3277a2a2e390a895db3a6a5a0dafd3bd0.UserRequestBuilder) {
    return i5f1f4bf3151bbb3488a322f91d88a1e3277a2a2e390a895db3a6a5a0dafd3bd0.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
