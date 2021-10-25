package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1902f3a0ccc4b4d090900e03e4ad1c30b5e8404be560e8d2ae81ccf93c63bc9b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/setupresourcesfolder"
    i34f2686ac9e572a0aab5c6d3a460d4aab368203139a21a9e6676179cbc76f29d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions"
    i7aec1a2ebd08a150442dd3f2dab446f096282c891d32dee633746b4a74177b6d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/resources"
    i9f6aff0f0fb85711382cc277517e1973a86d4d579f3d7e2f1c175fad1b973d28 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/categories"
    iaaa8edd814077594dcfe08fabae894b6b945f70a45186a958cdb96a4f3c66924 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/rubric"
    iba118f4097921de3e02dbaeb597a1cf1cc65f08ed80fe4ffdd011bbdeb936ed2 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/publish"
    i5da31198b4698c9632f11b1d4cb2162d37d7be431577136dcb6670fb9876d353 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item"
    i7406f0fbcb81f97517ec46d1a396eee4f12972305c2bd1f4a72ae53516d8d8b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/resources/item"
    ib56c03b6b9d132ee7fc45579e6690dcf8e4f40924a1bccf9043eee4b1c8322d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/categories/item"
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
func (m *EducationAssignmentRequestBuilder) Categories()(*i9f6aff0f0fb85711382cc277517e1973a86d4d579f3d7e2f1c175fad1b973d28.CategoriesRequestBuilder) {
    return i9f6aff0f0fb85711382cc277517e1973a86d4d579f3d7e2f1c175fad1b973d28.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) CategoriesById(id string)(*ib56c03b6b9d132ee7fc45579e6690dcf8e4f40924a1bccf9043eee4b1c8322d8.EducationCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationCategory_id"] = id
    }
    return ib56c03b6b9d132ee7fc45579e6690dcf8e4f40924a1bccf9043eee4b1c8322d8.NewEducationCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewEducationAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    m := &EducationAssignmentRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/education/me/assignments/{educationAssignment_id}{?select,expand}";
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
func (m *EducationAssignmentRequestBuilder) Publish()(*iba118f4097921de3e02dbaeb597a1cf1cc65f08ed80fe4ffdd011bbdeb936ed2.PublishRequestBuilder) {
    return iba118f4097921de3e02dbaeb597a1cf1cc65f08ed80fe4ffdd011bbdeb936ed2.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Resources()(*i7aec1a2ebd08a150442dd3f2dab446f096282c891d32dee633746b4a74177b6d.ResourcesRequestBuilder) {
    return i7aec1a2ebd08a150442dd3f2dab446f096282c891d32dee633746b4a74177b6d.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) ResourcesById(id string)(*i7406f0fbcb81f97517ec46d1a396eee4f12972305c2bd1f4a72ae53516d8d8b3.EducationAssignmentResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationAssignmentResource_id"] = id
    }
    return i7406f0fbcb81f97517ec46d1a396eee4f12972305c2bd1f4a72ae53516d8d8b3.NewEducationAssignmentResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Rubric()(*iaaa8edd814077594dcfe08fabae894b6b945f70a45186a958cdb96a4f3c66924.RubricRequestBuilder) {
    return iaaa8edd814077594dcfe08fabae894b6b945f70a45186a958cdb96a4f3c66924.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) SetUpResourcesFolder()(*i1902f3a0ccc4b4d090900e03e4ad1c30b5e8404be560e8d2ae81ccf93c63bc9b.SetUpResourcesFolderRequestBuilder) {
    return i1902f3a0ccc4b4d090900e03e4ad1c30b5e8404be560e8d2ae81ccf93c63bc9b.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Submissions()(*i34f2686ac9e572a0aab5c6d3a460d4aab368203139a21a9e6676179cbc76f29d.SubmissionsRequestBuilder) {
    return i34f2686ac9e572a0aab5c6d3a460d4aab368203139a21a9e6676179cbc76f29d.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) SubmissionsById(id string)(*i5da31198b4698c9632f11b1d4cb2162d37d7be431577136dcb6670fb9876d353.EducationSubmissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationSubmission_id"] = id
    }
    return i5da31198b4698c9632f11b1d4cb2162d37d7be431577136dcb6670fb9876d353.NewEducationSubmissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
