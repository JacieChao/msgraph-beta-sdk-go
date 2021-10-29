package dataclassification

import (
    i073bf3b84f9c5c09e315a22772bce801c5a9a6d430187bdc73f389f2cbca1815 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchdatastores"
    i373b8bc295a82bf2e640c37a2dcfed846455f2e1735c48e5cb25dbe3d1217798 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchuploadagents"
    i3cff8239684e50220c01c5330bc2d3ce5fe0516a4602276e77c0c2ef9881baeb "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifytext"
    i4fea7db9ec3dfb8eee5d0f28c2b8b50636edd00c4d449297a2541a4ef2b44e2e "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/evaluatelabeljobs"
    i6594c9c85b2fe737b67f8e7096b7bac27a77e943585e903afc1c799b3c1391e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/jobs"
    i6d2bd6b47f18fd25d8706f12f070fd69369097c3c3fa36e4f0972447bd008caf "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/sensitivitylabels"
    i89b152cd04bb38d005cae24645ee5cda597c6283fbeeb11796e9c2aebeda49a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifyfilejobs"
    i8fb53346ef40e76f80dc166fa4d90c7b83b5c39805b01bfbb5c0714b774ae860 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifyexactmatches"
    ic6b82df222a8a700bc0ac43b62566ac43d9007a0e2ba62d0e8397cce2c308d2e "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifyfile"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie0685d2de7afc64537caa2ec0727e7ebc4b4f61202921a69878ce943d829154d "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/evaluatedlppoliciesjobs"
    if188e6305418a24469cd7682c1cd0ca10f00ccdfaa5a763f68d87201ba832986 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/sensitivetypes"
    ifbc5c976485c8ce6802245796a7bca4914bfd3e80f624a761511266b13799b70 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifytextjobs"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i27eb8dd6f1e19fde6babe9ea8109f5b765c6eb4971d149bdf4ac63504c239dbf "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/sensitivitylabels/item"
    i45cfe1305c2a625c914638131757b68e2e595331b1cb835c1b30df24ea16f31f "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifytextjobs/item"
    i4792e57c8d9125efc146f0b2e40e63e9de13f444548ccb305eb4736a1a5be3a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifyfilejobs/item"
    i4d55b1afcc01d53334597316e2857e2085d3a6c2a3947cdde8abd9a1f813161f "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/evaluatedlppoliciesjobs/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6031e0e8286fe11a8cf1dce07b526e8cf9498b49a741cc2cc1f28e944ef6db8b "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchuploadagents/item"
    i787f440e60431af3738f5d22582c9098279e6cf44887787414a87feaede3518a "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/jobs/item"
    i7eaf730996e637195332ee006d958f53fdc086a573ef5e8ccc04f904d293e94c "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/sensitivetypes/item"
    i839259d04bc1e0752cd1dfcdc547564ab43c76e7472ad42df51a84e7b5459230 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/exactmatchdatastores/item"
    ic8fa6b5aea70c4c29088711c4efe56023dceaa26babb0879a81e2e2b0ac67178 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifytext/item"
    id639ff5c7f5dd2905bbbcf9dbcd6a7912904f3f8da7f947d8b3eec42a0d9bfd4 "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/evaluatelabeljobs/item"
    idd1a815d516e23b374b9b127ceaa407b8f4e99a7b2d6b9faf512ed6c6581965f "github.com/microsoftgraph/msgraph-beta-sdk-go/dataclassification/classifyfile/item"
)

type DataClassificationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DataClassificationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *DataClassificationRequestBuilder) ClassifyExactMatches()(*i8fb53346ef40e76f80dc166fa4d90c7b83b5c39805b01bfbb5c0714b774ae860.ClassifyExactMatchesRequestBuilder) {
    return i8fb53346ef40e76f80dc166fa4d90c7b83b5c39805b01bfbb5c0714b774ae860.NewClassifyExactMatchesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ClassifyFile()(*ic6b82df222a8a700bc0ac43b62566ac43d9007a0e2ba62d0e8397cce2c308d2e.ClassifyFileRequestBuilder) {
    return ic6b82df222a8a700bc0ac43b62566ac43d9007a0e2ba62d0e8397cce2c308d2e.NewClassifyFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ClassifyFileById(id string)(*idd1a815d516e23b374b9b127ceaa407b8f4e99a7b2d6b9faf512ed6c6581965f.FileClassificationRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fileClassificationRequest_id"] = id
    }
    return idd1a815d516e23b374b9b127ceaa407b8f4e99a7b2d6b9faf512ed6c6581965f.NewFileClassificationRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ClassifyFileJobs()(*i89b152cd04bb38d005cae24645ee5cda597c6283fbeeb11796e9c2aebeda49a9.ClassifyFileJobsRequestBuilder) {
    return i89b152cd04bb38d005cae24645ee5cda597c6283fbeeb11796e9c2aebeda49a9.NewClassifyFileJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ClassifyFileJobsById(id string)(*i4792e57c8d9125efc146f0b2e40e63e9de13f444548ccb305eb4736a1a5be3a5.JobResponseBaseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase_id"] = id
    }
    return i4792e57c8d9125efc146f0b2e40e63e9de13f444548ccb305eb4736a1a5be3a5.NewJobResponseBaseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ClassifyText()(*i3cff8239684e50220c01c5330bc2d3ce5fe0516a4602276e77c0c2ef9881baeb.ClassifyTextRequestBuilder) {
    return i3cff8239684e50220c01c5330bc2d3ce5fe0516a4602276e77c0c2ef9881baeb.NewClassifyTextRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ClassifyTextById(id string)(*ic8fa6b5aea70c4c29088711c4efe56023dceaa26babb0879a81e2e2b0ac67178.TextClassificationRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["textClassificationRequest_id"] = id
    }
    return ic8fa6b5aea70c4c29088711c4efe56023dceaa26babb0879a81e2e2b0ac67178.NewTextClassificationRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ClassifyTextJobs()(*ifbc5c976485c8ce6802245796a7bca4914bfd3e80f624a761511266b13799b70.ClassifyTextJobsRequestBuilder) {
    return ifbc5c976485c8ce6802245796a7bca4914bfd3e80f624a761511266b13799b70.NewClassifyTextJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ClassifyTextJobsById(id string)(*i45cfe1305c2a625c914638131757b68e2e595331b1cb835c1b30df24ea16f31f.JobResponseBaseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase_id"] = id
    }
    return i45cfe1305c2a625c914638131757b68e2e595331b1cb835c1b30df24ea16f31f.NewJobResponseBaseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewDataClassificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DataClassificationRequestBuilder) {
    m := &DataClassificationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/dataClassification{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDataClassificationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DataClassificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataClassificationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DataClassificationRequestBuilder) CreateGetRequestInformation(q func (value *DataClassificationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DataClassificationRequestBuilderGetQueryParameters)
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
func (m *DataClassificationRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DataClassificationService, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DataClassificationRequestBuilder) EvaluateDlpPoliciesJobs()(*ie0685d2de7afc64537caa2ec0727e7ebc4b4f61202921a69878ce943d829154d.EvaluateDlpPoliciesJobsRequestBuilder) {
    return ie0685d2de7afc64537caa2ec0727e7ebc4b4f61202921a69878ce943d829154d.NewEvaluateDlpPoliciesJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) EvaluateDlpPoliciesJobsById(id string)(*i4d55b1afcc01d53334597316e2857e2085d3a6c2a3947cdde8abd9a1f813161f.JobResponseBaseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase_id"] = id
    }
    return i4d55b1afcc01d53334597316e2857e2085d3a6c2a3947cdde8abd9a1f813161f.NewJobResponseBaseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) EvaluateLabelJobs()(*i4fea7db9ec3dfb8eee5d0f28c2b8b50636edd00c4d449297a2541a4ef2b44e2e.EvaluateLabelJobsRequestBuilder) {
    return i4fea7db9ec3dfb8eee5d0f28c2b8b50636edd00c4d449297a2541a4ef2b44e2e.NewEvaluateLabelJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) EvaluateLabelJobsById(id string)(*id639ff5c7f5dd2905bbbcf9dbcd6a7912904f3f8da7f947d8b3eec42a0d9bfd4.JobResponseBaseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase_id"] = id
    }
    return id639ff5c7f5dd2905bbbcf9dbcd6a7912904f3f8da7f947d8b3eec42a0d9bfd4.NewJobResponseBaseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ExactMatchDataStores()(*i073bf3b84f9c5c09e315a22772bce801c5a9a6d430187bdc73f389f2cbca1815.ExactMatchDataStoresRequestBuilder) {
    return i073bf3b84f9c5c09e315a22772bce801c5a9a6d430187bdc73f389f2cbca1815.NewExactMatchDataStoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ExactMatchDataStoresById(id string)(*i839259d04bc1e0752cd1dfcdc547564ab43c76e7472ad42df51a84e7b5459230.ExactMatchDataStoreRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["exactMatchDataStore_id"] = id
    }
    return i839259d04bc1e0752cd1dfcdc547564ab43c76e7472ad42df51a84e7b5459230.NewExactMatchDataStoreRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ExactMatchUploadAgents()(*i373b8bc295a82bf2e640c37a2dcfed846455f2e1735c48e5cb25dbe3d1217798.ExactMatchUploadAgentsRequestBuilder) {
    return i373b8bc295a82bf2e640c37a2dcfed846455f2e1735c48e5cb25dbe3d1217798.NewExactMatchUploadAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) ExactMatchUploadAgentsById(id string)(*i6031e0e8286fe11a8cf1dce07b526e8cf9498b49a741cc2cc1f28e944ef6db8b.ExactMatchUploadAgentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["exactMatchUploadAgent_id"] = id
    }
    return i6031e0e8286fe11a8cf1dce07b526e8cf9498b49a741cc2cc1f28e944ef6db8b.NewExactMatchUploadAgentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) Get(q func (value *DataClassificationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DataClassificationService, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDataClassificationService() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DataClassificationService), nil
}
func (m *DataClassificationRequestBuilder) Jobs()(*i6594c9c85b2fe737b67f8e7096b7bac27a77e943585e903afc1c799b3c1391e0.JobsRequestBuilder) {
    return i6594c9c85b2fe737b67f8e7096b7bac27a77e943585e903afc1c799b3c1391e0.NewJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) JobsById(id string)(*i787f440e60431af3738f5d22582c9098279e6cf44887787414a87feaede3518a.JobResponseBaseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["jobResponseBase_id"] = id
    }
    return i787f440e60431af3738f5d22582c9098279e6cf44887787414a87feaede3518a.NewJobResponseBaseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DataClassificationService, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *DataClassificationRequestBuilder) SensitiveTypes()(*if188e6305418a24469cd7682c1cd0ca10f00ccdfaa5a763f68d87201ba832986.SensitiveTypesRequestBuilder) {
    return if188e6305418a24469cd7682c1cd0ca10f00ccdfaa5a763f68d87201ba832986.NewSensitiveTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) SensitiveTypesById(id string)(*i7eaf730996e637195332ee006d958f53fdc086a573ef5e8ccc04f904d293e94c.SensitiveTypeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitiveType_id"] = id
    }
    return i7eaf730996e637195332ee006d958f53fdc086a573ef5e8ccc04f904d293e94c.NewSensitiveTypeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) SensitivityLabels()(*i6d2bd6b47f18fd25d8706f12f070fd69369097c3c3fa36e4f0972447bd008caf.SensitivityLabelsRequestBuilder) {
    return i6d2bd6b47f18fd25d8706f12f070fd69369097c3c3fa36e4f0972447bd008caf.NewSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DataClassificationRequestBuilder) SensitivityLabelsById(id string)(*i27eb8dd6f1e19fde6babe9ea8109f5b765c6eb4971d149bdf4ac63504c239dbf.SensitivityLabelRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel_id"] = id
    }
    return i27eb8dd6f1e19fde6babe9ea8109f5b765c6eb4971d149bdf4ac63504c239dbf.NewSensitivityLabelRequestBuilderInternal(urlTplParams, m.requestAdapter);
}