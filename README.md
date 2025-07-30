# v3-go-sdk

##### This is currently in development.  Do not use for production.

## Usage
1. Most simple way is to use the Interface defined in client_factory.go, `ApiClientActions`.
2. Implement the 3 methods for Your client factory:  
`HandleAuthorization` - this method is for handling the authorization step in the OAuth2 flow, can access the OAuth config, to fetch the authorization URL.  
`SaveTokens` - this method is for saving the OAuth2 tokens. An Oauth tokens object is provided as an input argument.   
`FetchTokens` - this method is used for fetching extisting OAuth2 tokens, so the Authorization flow could be skipped, if an existing OAuth2 connection has been created previously. It must return an `oauth2.Token` object.
3. `scoro.GetClient` will return a API client object that has all the CRUD operations implemented. It takes the client config and the implementation of the client factory object as arguments.
4. You can create the `ApiConfig` object manually, or use the `GetAPIClientConfig`.
5. See the example client factory implementation in `examples/example_client.go`.
