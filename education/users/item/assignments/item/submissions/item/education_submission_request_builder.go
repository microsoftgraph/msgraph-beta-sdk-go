package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i03a5515c9c388939cec0c62975b5200f7cc84414e6ff4962c81d48008588c4ca "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/submit"
    i164979e9c6e40c75579e19bae874710bbe54a4757fa82f177ce989d7159b8115 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/submittedresources"
    i29e598c0e3132502e75ebc5f97d630c0d489e1a3695a279765b136eb63851edb "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/unsubmit"
    i620d352f6d923470da79c3942b8c8aaaf958d8aaa2e3d4eb26af40a92833870f "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/resources"
    i6ba46d6835c0a47a27d09f0253c579859dcd40873b53932a91d1e8d864ea4af4 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/outcomes"
    i80a1281eb0b576d298b514fc90347e66750486fa7e9fdfd1403829774fe5acba "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/return_escpaped"
    if15e4fd1817bfddf93807ec99a8a0cbce18f2c72a43161efa316e4ff2ad07585 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/setupresourcesfolder"
    idcb18db22001d0c8560f4e04f592d2ab741e41fdccd32e2bbd2826405e8162b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/outcomes/item"
    ided68156635ca8883976581f28a920199cc5242c12a0b15c16f989f304e5824b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/resources/item"
    ie8d02aac9f22a617e7db20c68df056fad4bb1fb1425faaa25cf4ef67594043ec "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/submittedresources/item"
)

type EducationSubmissionRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EducationSubmissionRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewEducationSubmissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionRequestBuilder) {
    m := &EducationSubmissionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/education/users/{educationUser_id}/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}{?select,expand}";
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
func NewEducationSubmissionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSubmissionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EducationSubmissionRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationSubmissionRequestBuilder) CreateGetRequestInformation(q func (value *EducationSubmissionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EducationSubmissionRequestBuilderGetQueryParameters)
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
func (m *EducationSubmissionRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmission, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationSubmissionRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationSubmissionRequestBuilder) Get(q func (value *EducationSubmissionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmission, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationSubmission() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmission), nil
}
func (m *EducationSubmissionRequestBuilder) Outcomes()(*i6ba46d6835c0a47a27d09f0253c579859dcd40873b53932a91d1e8d864ea4af4.OutcomesRequestBuilder) {
    return i6ba46d6835c0a47a27d09f0253c579859dcd40873b53932a91d1e8d864ea4af4.NewOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) OutcomesById(id string)(*idcb18db22001d0c8560f4e04f592d2ab741e41fdccd32e2bbd2826405e8162b6.EducationOutcomeRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationOutcome_id"] = id
    }
    return idcb18db22001d0c8560f4e04f592d2ab741e41fdccd32e2bbd2826405e8162b6.NewEducationOutcomeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmission, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationSubmissionRequestBuilder) Resources()(*i620d352f6d923470da79c3942b8c8aaaf958d8aaa2e3d4eb26af40a92833870f.ResourcesRequestBuilder) {
    return i620d352f6d923470da79c3942b8c8aaaf958d8aaa2e3d4eb26af40a92833870f.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) ResourcesById(id string)(*ided68156635ca8883976581f28a920199cc5242c12a0b15c16f989f304e5824b.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return ided68156635ca8883976581f28a920199cc5242c12a0b15c16f989f304e5824b.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Return_escpaped()(*i80a1281eb0b576d298b514fc90347e66750486fa7e9fdfd1403829774fe5acba.ReturnRequestBuilder) {
    return i80a1281eb0b576d298b514fc90347e66750486fa7e9fdfd1403829774fe5acba.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SetUpResourcesFolder()(*if15e4fd1817bfddf93807ec99a8a0cbce18f2c72a43161efa316e4ff2ad07585.SetUpResourcesFolderRequestBuilder) {
    return if15e4fd1817bfddf93807ec99a8a0cbce18f2c72a43161efa316e4ff2ad07585.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Submit()(*i03a5515c9c388939cec0c62975b5200f7cc84414e6ff4962c81d48008588c4ca.SubmitRequestBuilder) {
    return i03a5515c9c388939cec0c62975b5200f7cc84414e6ff4962c81d48008588c4ca.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SubmittedResources()(*i164979e9c6e40c75579e19bae874710bbe54a4757fa82f177ce989d7159b8115.SubmittedResourcesRequestBuilder) {
    return i164979e9c6e40c75579e19bae874710bbe54a4757fa82f177ce989d7159b8115.NewSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SubmittedResourcesById(id string)(*ie8d02aac9f22a617e7db20c68df056fad4bb1fb1425faaa25cf4ef67594043ec.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return ie8d02aac9f22a617e7db20c68df056fad4bb1fb1425faaa25cf4ef67594043ec.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Unsubmit()(*i29e598c0e3132502e75ebc5f97d630c0d489e1a3695a279765b136eb63851edb.UnsubmitRequestBuilder) {
    return i29e598c0e3132502e75ebc5f97d630c0d489e1a3695a279765b136eb63851edb.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
