package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i15f37890fd55ac82b8528efa2c29ff81a969a940b9bc9fb7d06dd2685cdac5e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments"
    i1bf348583d80598c9676a7d9a21e7756cd5ff18e1039e00b611e0b2f9919a3fd "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/schools"
    i2b5fb8ced6d46b1790e2c676445d42b0d843d84349474ada5e6fa2dddc8bbdbf "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/classes"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i5f1177027388a1108a3f7b3004b1c6991c2cb44b7ef2b62a156ff94dc12f20a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/rubrics"
    i89d89ec0623f7cf34f3adc86d9dbf1b4d79c69280a133bebe65ac66d61cbb93f "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/taughtclasses"
    if69eac37b7cfbf6d4343393881192e3fbb1a33fd9f5e895e7c21e9296d2dc71b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/user"
    i0f11e129cc0a86c04aeab4bef7a1bec96ec8dc4595bbcd883507aeaf4f1aaeda "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/rubrics/item"
    i65dac0e57837e375e66f86d3d45e4195ac0556bbd9feeeec3fc32d72883e5ede "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/classes/item"
    i6bc86d513ec01bc025e0973f496b3d684d0b9855eb8f39c0a2cedefc89f6ae63 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item"
    icb302a6ba25d6a0b246265f93572109598e31817fe673fd41657d25d2f112d1a "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/schools/item"
    if965f0e5730d3de38d9b0d48d0ecb28f888caa600902ee5890f5fea7880567f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/taughtclasses/item"
)

// EducationUserItemRequestBuilder provides operations to manage the users property of the microsoft.graph.educationRoot entity.
type EducationUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationUserItemRequestBuilderDeleteOptions options for Delete
type EducationUserItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationUserItemRequestBuilderGetOptions options for Get
type EducationUserItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationUserItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationUserItemRequestBuilderGetQueryParameters get users from education
type EducationUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationUserItemRequestBuilderPatchOptions options for Patch
type EducationUserItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationUserable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EducationUserItemRequestBuilder) Assignments()(*i15f37890fd55ac82b8528efa2c29ff81a969a940b9bc9fb7d06dd2685cdac5e3.AssignmentsRequestBuilder) {
    return i15f37890fd55ac82b8528efa2c29ff81a969a940b9bc9fb7d06dd2685cdac5e3.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.assignments.item collection
func (m *EducationUserItemRequestBuilder) AssignmentsById(id string)(*i6bc86d513ec01bc025e0973f496b3d684d0b9855eb8f39c0a2cedefc89f6ae63.EducationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignment_id"] = id
    }
    return i6bc86d513ec01bc025e0973f496b3d684d0b9855eb8f39c0a2cedefc89f6ae63.NewEducationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationUserItemRequestBuilder) Classes()(*i2b5fb8ced6d46b1790e2c676445d42b0d843d84349474ada5e6fa2dddc8bbdbf.ClassesRequestBuilder) {
    return i2b5fb8ced6d46b1790e2c676445d42b0d843d84349474ada5e6fa2dddc8bbdbf.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.classes.item collection
func (m *EducationUserItemRequestBuilder) ClassesById(id string)(*i65dac0e57837e375e66f86d3d45e4195ac0556bbd9feeeec3fc32d72883e5ede.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass_id"] = id
    }
    return i65dac0e57837e375e66f86d3d45e4195ac0556bbd9feeeec3fc32d72883e5ede.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationUserItemRequestBuilderInternal instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewEducationUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationUserItemRequestBuilder) {
    m := &EducationUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationUserItemRequestBuilder instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewEducationUserItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property users for education
func (m *EducationUserItemRequestBuilder) CreateDeleteRequestInformation(options *EducationUserItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get users from education
func (m *EducationUserItemRequestBuilder) CreateGetRequestInformation(options *EducationUserItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property users in education
func (m *EducationUserItemRequestBuilder) CreatePatchRequestInformation(options *EducationUserItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property users for education
func (m *EducationUserItemRequestBuilder) Delete(options *EducationUserItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get users from education
func (m *EducationUserItemRequestBuilder) Get(options *EducationUserItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationUserable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEducationUserFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationUserable), nil
}
// Patch update the navigation property users in education
func (m *EducationUserItemRequestBuilder) Patch(options *EducationUserItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationUserItemRequestBuilder) Rubrics()(*i5f1177027388a1108a3f7b3004b1c6991c2cb44b7ef2b62a156ff94dc12f20a7.RubricsRequestBuilder) {
    return i5f1177027388a1108a3f7b3004b1c6991c2cb44b7ef2b62a156ff94dc12f20a7.NewRubricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RubricsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.rubrics.item collection
func (m *EducationUserItemRequestBuilder) RubricsById(id string)(*i0f11e129cc0a86c04aeab4bef7a1bec96ec8dc4595bbcd883507aeaf4f1aaeda.EducationRubricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationRubric_id"] = id
    }
    return i0f11e129cc0a86c04aeab4bef7a1bec96ec8dc4595bbcd883507aeaf4f1aaeda.NewEducationRubricItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationUserItemRequestBuilder) Schools()(*i1bf348583d80598c9676a7d9a21e7756cd5ff18e1039e00b611e0b2f9919a3fd.SchoolsRequestBuilder) {
    return i1bf348583d80598c9676a7d9a21e7756cd5ff18e1039e00b611e0b2f9919a3fd.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchoolsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.schools.item collection
func (m *EducationUserItemRequestBuilder) SchoolsById(id string)(*icb302a6ba25d6a0b246265f93572109598e31817fe673fd41657d25d2f112d1a.EducationSchoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSchool_id"] = id
    }
    return icb302a6ba25d6a0b246265f93572109598e31817fe673fd41657d25d2f112d1a.NewEducationSchoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationUserItemRequestBuilder) TaughtClasses()(*i89d89ec0623f7cf34f3adc86d9dbf1b4d79c69280a133bebe65ac66d61cbb93f.TaughtClassesRequestBuilder) {
    return i89d89ec0623f7cf34f3adc86d9dbf1b4d79c69280a133bebe65ac66d61cbb93f.NewTaughtClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaughtClassesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.taughtClasses.item collection
func (m *EducationUserItemRequestBuilder) TaughtClassesById(id string)(*if965f0e5730d3de38d9b0d48d0ecb28f888caa600902ee5890f5fea7880567f4.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass_id"] = id
    }
    return if965f0e5730d3de38d9b0d48d0ecb28f888caa600902ee5890f5fea7880567f4.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationUserItemRequestBuilder) User()(*if69eac37b7cfbf6d4343393881192e3fbb1a33fd9f5e895e7c21e9296d2dc71b.UserRequestBuilder) {
    return if69eac37b7cfbf6d4343393881192e3fbb1a33fd9f5e895e7c21e9296d2dc71b.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
