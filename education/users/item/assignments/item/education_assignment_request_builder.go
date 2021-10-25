package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5b10b16160956c55311bec6d730fa2056fd37872e249abeca309071996a8fc0b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/resources"
    i8050f70ff077aa7a00ee30e1a3c86c0f0615a85e77ec2d73dc4e1bafa3f4e16f "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/setupresourcesfolder"
    ib73813a5d39b89f70827c0fa1818289fccec256ca0ea65df5135c2bf448aa519 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/rubric"
    ic194e429a42abea4fc9457129ec19b14e4af236b99092c2a55f726379f0ecc39 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions"
    idf7d52d5e40c4a2f0a279dbefcf41f586915062fc10338040bd70c8bde7211d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/publish"
    ie3e3ed607d0779e161141593c134093354a4e6852436b9197a40b9ac9a01eb0e "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/categories"
    i105c230efb20cac97a293b5469f4f607fe6cb947c167d5ee426b70bce49e9df5 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item"
    ibe13b36acf757c0c212481dd71384ef998a32ff3577f6c574e9f1785f7bb86f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/resources/item"
    iee68e3ee77c1c1596d86d9cdb46e4cba567ac32c664bddc07320057cef108c0c "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/categories/item"
)

type EducationAssignmentRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EducationAssignmentRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *EducationAssignmentRequestBuilder) Categories()(*ie3e3ed607d0779e161141593c134093354a4e6852436b9197a40b9ac9a01eb0e.CategoriesRequestBuilder) {
    return ie3e3ed607d0779e161141593c134093354a4e6852436b9197a40b9ac9a01eb0e.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) CategoriesById(id string)(*iee68e3ee77c1c1596d86d9cdb46e4cba567ac32c664bddc07320057cef108c0c.EducationCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationCategory_id"] = id
    }
    return iee68e3ee77c1c1596d86d9cdb46e4cba567ac32c664bddc07320057cef108c0c.NewEducationCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewEducationAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    m := &EducationAssignmentRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/education/users/{educationUser_id}/assignments/{educationAssignment_id}{?select,expand}";
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
func NewEducationAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EducationAssignmentRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationAssignmentRequestBuilder) CreateGetRequestInformation(q func (value *EducationAssignmentRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EducationAssignmentRequestBuilderGetQueryParameters)
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
func (m *EducationAssignmentRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationAssignment, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationAssignmentRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationAssignmentRequestBuilder) Get(q func (value *EducationAssignmentRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationAssignment() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationAssignment), nil
}
func (m *EducationAssignmentRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationAssignment, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationAssignmentRequestBuilder) Publish()(*idf7d52d5e40c4a2f0a279dbefcf41f586915062fc10338040bd70c8bde7211d7.PublishRequestBuilder) {
    return idf7d52d5e40c4a2f0a279dbefcf41f586915062fc10338040bd70c8bde7211d7.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Resources()(*i5b10b16160956c55311bec6d730fa2056fd37872e249abeca309071996a8fc0b.ResourcesRequestBuilder) {
    return i5b10b16160956c55311bec6d730fa2056fd37872e249abeca309071996a8fc0b.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) ResourcesById(id string)(*ibe13b36acf757c0c212481dd71384ef998a32ff3577f6c574e9f1785f7bb86f2.EducationAssignmentResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationAssignmentResource_id"] = id
    }
    return ibe13b36acf757c0c212481dd71384ef998a32ff3577f6c574e9f1785f7bb86f2.NewEducationAssignmentResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Rubric()(*ib73813a5d39b89f70827c0fa1818289fccec256ca0ea65df5135c2bf448aa519.RubricRequestBuilder) {
    return ib73813a5d39b89f70827c0fa1818289fccec256ca0ea65df5135c2bf448aa519.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) SetUpResourcesFolder()(*i8050f70ff077aa7a00ee30e1a3c86c0f0615a85e77ec2d73dc4e1bafa3f4e16f.SetUpResourcesFolderRequestBuilder) {
    return i8050f70ff077aa7a00ee30e1a3c86c0f0615a85e77ec2d73dc4e1bafa3f4e16f.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Submissions()(*ic194e429a42abea4fc9457129ec19b14e4af236b99092c2a55f726379f0ecc39.SubmissionsRequestBuilder) {
    return ic194e429a42abea4fc9457129ec19b14e4af236b99092c2a55f726379f0ecc39.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) SubmissionsById(id string)(*i105c230efb20cac97a293b5469f4f607fe6cb947c167d5ee426b70bce49e9df5.EducationSubmissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationSubmission_id"] = id
    }
    return i105c230efb20cac97a293b5469f4f607fe6cb947c167d5ee426b70bce49e9df5.NewEducationSubmissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
