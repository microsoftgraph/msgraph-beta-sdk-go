package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i15f37890fd55ac82b8528efa2c29ff81a969a940b9bc9fb7d06dd2685cdac5e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments"
    i1bf348583d80598c9676a7d9a21e7756cd5ff18e1039e00b611e0b2f9919a3fd "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/schools"
    i2b5fb8ced6d46b1790e2c676445d42b0d843d84349474ada5e6fa2dddc8bbdbf "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/classes"
    i5f1177027388a1108a3f7b3004b1c6991c2cb44b7ef2b62a156ff94dc12f20a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/rubrics"
    i89d89ec0623f7cf34f3adc86d9dbf1b4d79c69280a133bebe65ac66d61cbb93f "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/taughtclasses"
    if69eac37b7cfbf6d4343393881192e3fbb1a33fd9f5e895e7c21e9296d2dc71b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/user"
    i0f11e129cc0a86c04aeab4bef7a1bec96ec8dc4595bbcd883507aeaf4f1aaeda "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/rubrics/item"
    i6bc86d513ec01bc025e0973f496b3d684d0b9855eb8f39c0a2cedefc89f6ae63 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item"
)

type EducationUserRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EducationUserRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *EducationUserRequestBuilder) Assignments()(*i15f37890fd55ac82b8528efa2c29ff81a969a940b9bc9fb7d06dd2685cdac5e3.AssignmentsRequestBuilder) {
    return i15f37890fd55ac82b8528efa2c29ff81a969a940b9bc9fb7d06dd2685cdac5e3.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationUserRequestBuilder) AssignmentsById(id string)(*i6bc86d513ec01bc025e0973f496b3d684d0b9855eb8f39c0a2cedefc89f6ae63.EducationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationAssignment_id"] = id
    }
    return i6bc86d513ec01bc025e0973f496b3d684d0b9855eb8f39c0a2cedefc89f6ae63.NewEducationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationUserRequestBuilder) Classes()(*i2b5fb8ced6d46b1790e2c676445d42b0d843d84349474ada5e6fa2dddc8bbdbf.ClassesRequestBuilder) {
    return i2b5fb8ced6d46b1790e2c676445d42b0d843d84349474ada5e6fa2dddc8bbdbf.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEducationUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationUserRequestBuilder) {
    m := &EducationUserRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/education/users/{educationUser_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewEducationUserRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationUserRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EducationUserRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EducationUserRequestBuilder) CreateGetRequestInformation(q func (value *EducationUserRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EducationUserRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EducationUserRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EducationUserRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationUserRequestBuilder) Get(q func (value *EducationUserRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationUser, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationUser() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationUser), nil
}
func (m *EducationUserRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationUserRequestBuilder) Rubrics()(*i5f1177027388a1108a3f7b3004b1c6991c2cb44b7ef2b62a156ff94dc12f20a7.RubricsRequestBuilder) {
    return i5f1177027388a1108a3f7b3004b1c6991c2cb44b7ef2b62a156ff94dc12f20a7.NewRubricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationUserRequestBuilder) RubricsById(id string)(*i0f11e129cc0a86c04aeab4bef7a1bec96ec8dc4595bbcd883507aeaf4f1aaeda.EducationRubricRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationRubric_id"] = id
    }
    return i0f11e129cc0a86c04aeab4bef7a1bec96ec8dc4595bbcd883507aeaf4f1aaeda.NewEducationRubricRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationUserRequestBuilder) Schools()(*i1bf348583d80598c9676a7d9a21e7756cd5ff18e1039e00b611e0b2f9919a3fd.SchoolsRequestBuilder) {
    return i1bf348583d80598c9676a7d9a21e7756cd5ff18e1039e00b611e0b2f9919a3fd.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationUserRequestBuilder) TaughtClasses()(*i89d89ec0623f7cf34f3adc86d9dbf1b4d79c69280a133bebe65ac66d61cbb93f.TaughtClassesRequestBuilder) {
    return i89d89ec0623f7cf34f3adc86d9dbf1b4d79c69280a133bebe65ac66d61cbb93f.NewTaughtClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationUserRequestBuilder) User()(*if69eac37b7cfbf6d4343393881192e3fbb1a33fd9f5e895e7c21e9296d2dc71b.UserRequestBuilder) {
    return if69eac37b7cfbf6d4343393881192e3fbb1a33fd9f5e895e7c21e9296d2dc71b.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
