# `oauth2_proxy` test application

Pre-work - setup UPS and download proxy binary:

```bash
# Example for Microsoft Azure AD
cf create-user-provided-service proxy_ups -p @<(cat <<EOF
{
    "OAUTH2_PROXY_AZURE_TENANT": "xxx",
    "OAUTH2_PROXY_CLIENT_ID": "xxx",
    "OAUTH2_PROXY_CLIENT_SECRET": "xxx",
    "OAUTH2_PROXY_COOKIE_SECRET": "xxx",
    "OAUTH2_PROXY_EMAIL_DOMAINS": "xxx",
    "OAUTH2_PROXY_PROVIDER": "azure",
    "OAUTH2_PROXY_REDIRECT_URL": "https://cc-test-ae.example.com/oauth2/callback"
}
EOF
)

# Example for AWS Cognito
cf create-user-provided-service proxy_ups -p @<(cat <<EOF
{
    "OAUTH2_PROXY_CLIENT_ID": "xxx",
    "OAUTH2_PROXY_CLIENT_SECRET": "xxx",
    "OAUTH2_PROXY_COOKIE_SECRET": "xxx",
    "OAUTH2_PROXY_EMAIL_DOMAINS": "*",
    "OAUTH2_PROXY_OIDC_ISSUER_URL": "https://cognito-idp.ap-southeast-2.amazonaws.com/{id}",
    "OAUTH2_PROXY_PROVIDER": "oidc",
    "OAUTH2_PROXY_SCOPE": "openid",
    "OAUTH2_PROXY_REDIRECT_URL": "https://cc-test-ae.example.com/oauth2/callback"
}
EOF
)

# Download the binary
OAUTH2_PROXY_VERSION="3.2.0"
curl -L https://github.com/pusher/oauth2_proxy/releases/download/v${OAUTH2_PROXY_VERSION}/oauth2_proxy-v${OAUTH2_PROXY_VERSION}.linux-amd64.go1.11.tar.gz | \
    tar -Ozxf - release/oauth2_proxy-linux-amd64 > ./scripts/oauth2_proxy
chmod a+x ./scripts/oauth2_proxy
```

Then:

```bash
cf push
```

Done!
