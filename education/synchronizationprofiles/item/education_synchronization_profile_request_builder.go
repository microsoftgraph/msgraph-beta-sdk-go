package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i04d76fbaad9f08921b0fefffd4ae909a8fbadb2da7b557711ada64363e2e763d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/profilestatus"
    i9915f3a66e2de398bdfae23c9b58bdf2c4a7c373437f422e62cc19ab1b37c01d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/errors"
    ibd324584f412d6082609e09d36338222a1ea33342c17f036ead958c71c7e57d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/resume"
    ibfdf2171a3f86b83740d33b9eb574ed52ef8e274f8b0c78649787ec4351bfa5d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/pause"
    icea2a0d7fcf828c9af01c7bba86f54c7305fbb4d6c4cea549b5ee5ece1723710 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/uploadurl"
    id814594e10d99c016c54d9229da3f00c037b42a01b52dd4efad6b8da5d373fb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/reset"
    if70ade9d653b89ed0849b345d9e94a2327f9abc57ba5b75b6511e5e4013d1aea "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/start"
    i88e077efbe7af1476bbdd89a3d093be25e50e57c1d6a1e16d2b9f2406ba01c90 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/errors/item"
)

type EducationSynchronizationProfileRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EducationSynchronizationProfileRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewEducationSynchronizationProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSynchronizationProfileRequestBuilder) {
    m := &EducationSynchronizationProfileRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/education/synchronizationProfiles/{educationSynchronizationProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewEducationSynchronizationProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSynchronizationProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSynchronizationProfileRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EducationSynchronizationProfileRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationSynchronizationProfileRequestBuilder) CreateGetRequestInformation(q func (value *EducationSynchronizationProfileRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EducationSynchronizationProfileRequestBuilderGetQueryParameters)
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
func (m *EducationSynchronizationProfileRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSynchronizationProfile, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationSynchronizationProfileRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationSynchronizationProfileRequestBuilder) Errors()(*i9915f3a66e2de398bdfae23c9b58bdf2c4a7c373437f422e62cc19ab1b37c01d.ErrorsRequestBuilder) {
    return i9915f3a66e2de398bdfae23c9b58bdf2c4a7c373437f422e62cc19ab1b37c01d.NewErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSynchronizationProfileRequestBuilder) ErrorsById(id string)(*i88e077efbe7af1476bbdd89a3d093be25e50e57c1d6a1e16d2b9f2406ba01c90.EducationSynchronizationErrorRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSynchronizationError_id"] = id
    }
    return i88e077efbe7af1476bbdd89a3d093be25e50e57c1d6a1e16d2b9f2406ba01c90.NewEducationSynchronizationErrorRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSynchronizationProfileRequestBuilder) Get(q func (value *EducationSynchronizationProfileRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSynchronizationProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationSynchronizationProfile() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSynchronizationProfile), nil
}
func (m *EducationSynchronizationProfileRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSynchronizationProfile, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationSynchronizationProfileRequestBuilder) Pause()(*ibfdf2171a3f86b83740d33b9eb574ed52ef8e274f8b0c78649787ec4351bfa5d.PauseRequestBuilder) {
    return ibfdf2171a3f86b83740d33b9eb574ed52ef8e274f8b0c78649787ec4351bfa5d.NewPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSynchronizationProfileRequestBuilder) ProfileStatus()(*i04d76fbaad9f08921b0fefffd4ae909a8fbadb2da7b557711ada64363e2e763d.ProfileStatusRequestBuilder) {
    return i04d76fbaad9f08921b0fefffd4ae909a8fbadb2da7b557711ada64363e2e763d.NewProfileStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSynchronizationProfileRequestBuilder) Reset()(*id814594e10d99c016c54d9229da3f00c037b42a01b52dd4efad6b8da5d373fb6.ResetRequestBuilder) {
    return id814594e10d99c016c54d9229da3f00c037b42a01b52dd4efad6b8da5d373fb6.NewResetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSynchronizationProfileRequestBuilder) Resume()(*ibd324584f412d6082609e09d36338222a1ea33342c17f036ead958c71c7e57d6.ResumeRequestBuilder) {
    return ibd324584f412d6082609e09d36338222a1ea33342c17f036ead958c71c7e57d6.NewResumeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSynchronizationProfileRequestBuilder) Start()(*if70ade9d653b89ed0849b345d9e94a2327f9abc57ba5b75b6511e5e4013d1aea.StartRequestBuilder) {
    return if70ade9d653b89ed0849b345d9e94a2327f9abc57ba5b75b6511e5e4013d1aea.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSynchronizationProfileRequestBuilder) UploadUrl()(*icea2a0d7fcf828c9af01c7bba86f54c7305fbb4d6c4cea549b5ee5ece1723710.UploadUrlRequestBuilder) {
    return icea2a0d7fcf828c9af01c7bba86f54c7305fbb4d6c4cea549b5ee5ece1723710.NewUploadUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
