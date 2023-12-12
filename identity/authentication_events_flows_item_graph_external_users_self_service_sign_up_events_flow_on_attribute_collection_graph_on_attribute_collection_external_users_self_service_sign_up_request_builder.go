package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder casts the previous resource to onAttributeCollectionExternalUsersSelfServiceSignUp.
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attributes provides operations to manage the attributes property of the microsoft.graph.onAttributeCollectionExternalUsersSelfServiceSignUp entity.
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder) Attributes()(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesRequestBuilder) {
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpAttributesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilderInternal instantiates a new GraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder) {
    m := &AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/graph.onAttributeCollectionExternalUsersSelfServiceSignUp", pathParameters),
    }
    return m
}
// NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder instantiates a new GraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder and sets the default values.
func NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the items of type microsoft.graph.onAttributeCollectionExternalUsersSelfServiceSignUp in the microsoft.graph.onAttributeCollectionHandler collection
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.onAttributeCollectionExternalUsersSelfServiceSignUp in the microsoft.graph.onAttributeCollectionHandler collection
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder) WithUrl(rawUrl string)(*AuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder) {
    return NewAuthenticationEventsFlowsItemGraphExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionGraphOnAttributeCollectionExternalUsersSelfServiceSignUpRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
