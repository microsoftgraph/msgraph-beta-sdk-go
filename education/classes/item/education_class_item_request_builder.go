package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationClassItemRequestBuilderDeleteOptions options for Delete
type EducationClassItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationClassItemRequestBuilderGetOptions options for Get
type EducationClassItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationClassItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationClassItemRequestBuilderGetQueryParameters get classes from education
type EducationClassItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationClassItemRequestBuilderPatchOptions options for Patch
type EducationClassItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationClassable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
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
        urlTplParams["educationCategory_id"] = id
    }
    return id33c9a2adb44743656c66ac6287fd646205cf808807397442c562d7261ab81af.NewEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationClassItemRequestBuilder) AssignmentDefaults()(*ibdb5b1525c6e65ee517f1f36e31fe5660be846f8bc83c1582f2717d92b8b77b4.AssignmentDefaultsRequestBuilder) {
    return ibdb5b1525c6e65ee517f1f36e31fe5660be846f8bc83c1582f2717d92b8b77b4.NewAssignmentDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["educationAssignment_id"] = id
    }
    return i1a0550da4eee947582946109e27a5d8b97a5dd93365d78fb41b93f4fcb684613.NewEducationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationClassItemRequestBuilder) AssignmentSettings()(*i1a2560c3e16777ce4016204fe1a193e56b0af48fa14ae66e31d31833ad245f65.AssignmentSettingsRequestBuilder) {
    return i1a2560c3e16777ce4016204fe1a193e56b0af48fa14ae66e31d31833ad245f65.NewAssignmentSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEducationClassItemRequestBuilderInternal instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewEducationClassItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationClassItemRequestBuilder) {
    m := &EducationClassItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationClassItemRequestBuilder instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewEducationClassItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationClassItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationClassItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property classes for education
func (m *EducationClassItemRequestBuilder) CreateDeleteRequestInformation(options *EducationClassItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get classes from education
func (m *EducationClassItemRequestBuilder) CreateGetRequestInformation(options *EducationClassItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property classes in education
func (m *EducationClassItemRequestBuilder) CreatePatchRequestInformation(options *EducationClassItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property classes for education
func (m *EducationClassItemRequestBuilder) Delete(options *EducationClassItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get classes from education
func (m *EducationClassItemRequestBuilder) Get(options *EducationClassItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationClassable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEducationClassFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationClassable), nil
}
func (m *EducationClassItemRequestBuilder) Group()(*if2d0fdc70365b40fab81c748f0f6e31abf6fb9e7d22890e73bd7c57afc299ac6.GroupRequestBuilder) {
    return if2d0fdc70365b40fab81c748f0f6e31abf6fb9e7d22890e73bd7c57afc299ac6.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["educationUser_id"] = id
    }
    return i168885acb42722fc703411df677333e0795a5a687835bc4830b53e95a8ddc301.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property classes in education
func (m *EducationClassItemRequestBuilder) Patch(options *EducationClassItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["educationSchool_id"] = id
    }
    return i9dcf5d71a4b8b21e2df678df6182539084fb161769f694781827d339e20e09c8.NewEducationSchoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["educationUser_id"] = id
    }
    return i79ac00214ff83c8c45d677f0030298feea62079c3b3d345b2f0b38922f74f671.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
