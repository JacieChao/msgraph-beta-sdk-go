package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder provides operations to call the evaluateRemoval method.
type ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderInternal instantiates a new SecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder) {
    m := &ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/security/informationProtection/sensitivityLabels/security.evaluateRemoval", pathParameters),
    }
    return m
}
// NewItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder instantiates a new SecurityEvaluateRemovalRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action evaluateRemoval
func (m *ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder) Post(ctx context.Context, body ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalEvaluateRemovalResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalEvaluateRemovalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalEvaluateRemovalResponseable), nil
}
// ToPostRequestInformation invoke action evaluateRemoval
func (m *ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalEvaluateRemovalPostRequestBodyable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsSecurityEvaluateRemovalRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
