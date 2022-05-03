package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1a2560c3e16777ce4016204fe1a193e56b0af48fa14ae66e31d31833ad245f65 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignmentsettings"
    i447743c6be252c60dde89565d4e1d2abcfde356cbc33dc813bbd667af4340b3f "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/schools"
    i4dc519cf190b35025011540ce15503c6dfbae11c1971b05ccbe2c89907556782 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/members"
    i8829116957e541ac2071bd39ad0c5693e10cda61555863d88c5b45462262992c "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/teachers"
    i8e2e53960e8a6a06380d4ef2d8f516a18d60590f7da782cbf7c461851a4b2fe9 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments"
    ibdb5b1525c6e65ee517f1f36e31fe5660be846f8bc83c1582f2717d92b8b77b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignmentdefaults"
    ie2d9988dc2165607477b233458b6b41efb14b20ef871a6f5407c79527c1fe3d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignmentcategories"
    if2d0fdc70365b40fab81c748f0f6e31abf6fb9e7d22890e73bd7c57afc299ac6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/group"
    i168885acb42722fc703411df677333e0795a5a687835bc4830b53e95a8ddc301 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/members/item"
    i1a0550da4eee947582946109e27a5d8b97a5dd93365d78fb41b93f4fcb684613 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item"
    i79ac00214ff83c8c45d677f0030298feea62079c3b3d345b2f0b38922f74f671 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/teachers/item"
    i9dcf5d71a4b8b21e2df678df6182539084fb161769f694781827d339e20e09c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/schools/item"
    id33c9a2adb44743656c66ac6287fd646205cf808807397442c562d7261ab81af "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignmentcategories/item"
)

// EducationClassItemRequestBuilder provides operations to manage the classes property of the microsoft.graph.educationRoot entity.
type EducationClassItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EducationClassItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationClassItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EducationClassItemRequestBuilderGetQueryParameters get classes from education
type EducationClassItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EducationClassItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationClassItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EducationClassItemRequestBuilderGetQueryParameters
}
// EducationClassItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationClassItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignmentCategories the assignmentCategories property
func (m *EducationClassItemRequestBuilder) AssignmentCategories()(*ie2d9988dc2165607477b233458b6b41efb14b20ef871a6f5407c79527c1fe3d1.AssignmentCategoriesRequestBuilder) {
    return ie2d9988dc2165607477b233458b6b41efb14b20ef871a6f5407c79527c1fe3d1.NewAssignmentCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentCategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.classes.item.assignmentCategories.item collection
func (m *EducationClassItemRequestBuilder) AssignmentCategoriesById(id string)(*id33c9a2adb44743656c66ac6287fd646205cf808807397442c562d7261ab81af.EducationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory%2Did"] = id
    }
    return id33c9a2adb44743656c66ac6287fd646205cf808807397442c562d7261ab81af.NewEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignmentDefaults the assignmentDefaults property
func (m *EducationClassItemRequestBuilder) AssignmentDefaults()(*ibdb5b1525c6e65ee517f1f36e31fe5660be846f8bc83c1582f2717d92b8b77b4.AssignmentDefaultsRequestBuilder) {
    return ibdb5b1525c6e65ee517f1f36e31fe5660be846f8bc83c1582f2717d92b8b77b4.NewAssignmentDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *EducationClassItemRequestBuilder) Assignments()(*i8e2e53960e8a6a06380d4ef2d8f516a18d60590f7da782cbf7c461851a4b2fe9.AssignmentsRequestBuilder) {
    return i8e2e53960e8a6a06380d4ef2d8f516a18d60590f7da782cbf7c461851a4b2fe9.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.classes.item.assignments.item collection
func (m *EducationClassItemRequestBuilder) AssignmentsById(id string)(*i1a0550da4eee947582946109e27a5d8b97a5dd93365d78fb41b93f4fcb684613.EducationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignment%2Did"] = id
    }
    return i1a0550da4eee947582946109e27a5d8b97a5dd93365d78fb41b93f4fcb684613.NewEducationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignmentSettings the assignmentSettings property
func (m *EducationClassItemRequestBuilder) AssignmentSettings()(*i1a2560c3e16777ce4016204fe1a193e56b0af48fa14ae66e31d31833ad245f65.AssignmentSettingsRequestBuilder) {
    return i1a2560c3e16777ce4016204fe1a193e56b0af48fa14ae66e31d31833ad245f65.NewAssignmentSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEducationClassItemRequestBuilderInternal instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewEducationClassItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationClassItemRequestBuilder) {
    m := &EducationClassItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationClassItemRequestBuilder instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewEducationClassItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationClassItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationClassItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property classes for education
func (m *EducationClassItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property classes for education
func (m *EducationClassItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EducationClassItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration get classes from education
func (m *EducationClassItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get classes from education
func (m *EducationClassItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EducationClassItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property classes in education
func (m *EducationClassItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationClassable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property classes in education
func (m *EducationClassItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationClassable, requestConfiguration *EducationClassItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteWithResponseHandler delete navigation property classes for education
func (m *EducationClassItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *EducationClassItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property classes for education
func (m *EducationClassItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *EducationClassItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler get classes from education
func (m *EducationClassItemRequestBuilder) GetWithResponseHandler(requestConfiguration *EducationClassItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationClassable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler get classes from education
func (m *EducationClassItemRequestBuilder) GetWithResponseHandler(requestConfiguration *EducationClassItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationClassable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationClassFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationClassable), nil
}
// Group the group property
func (m *EducationClassItemRequestBuilder) Group()(*if2d0fdc70365b40fab81c748f0f6e31abf6fb9e7d22890e73bd7c57afc299ac6.GroupRequestBuilder) {
    return if2d0fdc70365b40fab81c748f0f6e31abf6fb9e7d22890e73bd7c57afc299ac6.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Members the members property
func (m *EducationClassItemRequestBuilder) Members()(*i4dc519cf190b35025011540ce15503c6dfbae11c1971b05ccbe2c89907556782.MembersRequestBuilder) {
    return i4dc519cf190b35025011540ce15503c6dfbae11c1971b05ccbe2c89907556782.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.classes.item.members.item collection
func (m *EducationClassItemRequestBuilder) MembersById(id string)(*i168885acb42722fc703411df677333e0795a5a687835bc4830b53e95a8ddc301.EducationUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationUser%2Did"] = id
    }
    return i168885acb42722fc703411df677333e0795a5a687835bc4830b53e95a8ddc301.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property classes in education
func (m *EducationClassItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationClassable, requestConfiguration *EducationClassItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property classes in education
func (m *EducationClassItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationClassable, requestConfiguration *EducationClassItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Schools the schools property
func (m *EducationClassItemRequestBuilder) Schools()(*i447743c6be252c60dde89565d4e1d2abcfde356cbc33dc813bbd667af4340b3f.SchoolsRequestBuilder) {
    return i447743c6be252c60dde89565d4e1d2abcfde356cbc33dc813bbd667af4340b3f.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchoolsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.classes.item.schools.item collection
func (m *EducationClassItemRequestBuilder) SchoolsById(id string)(*i9dcf5d71a4b8b21e2df678df6182539084fb161769f694781827d339e20e09c8.EducationSchoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSchool%2Did"] = id
    }
    return i9dcf5d71a4b8b21e2df678df6182539084fb161769f694781827d339e20e09c8.NewEducationSchoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Teachers the teachers property
func (m *EducationClassItemRequestBuilder) Teachers()(*i8829116957e541ac2071bd39ad0c5693e10cda61555863d88c5b45462262992c.TeachersRequestBuilder) {
    return i8829116957e541ac2071bd39ad0c5693e10cda61555863d88c5b45462262992c.NewTeachersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeachersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.classes.item.teachers.item collection
func (m *EducationClassItemRequestBuilder) TeachersById(id string)(*i79ac00214ff83c8c45d677f0030298feea62079c3b3d345b2f0b38922f74f671.EducationUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationUser%2Did"] = id
    }
    return i79ac00214ff83c8c45d677f0030298feea62079c3b3d345b2f0b38922f74f671.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
