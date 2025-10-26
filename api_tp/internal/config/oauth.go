// config/oauth.go
package config

import (
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "golang.org/x/oauth2/yandex"
    "golang.org/x/oauth2/vk"
)


var (
    GoogleOAuthConfig = &oauth2.Config{
        ClientID:     "your-google-client-id",
        ClientSecret: "your-google-client-secret",
        RedirectURL:  "http://localhost:8080/auth/google/callback",
        Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
        Endpoint:     google.Endpoint,
    }
    
    YandexOAuthConfig = &oauth2.Config{
        ClientID:     "your-yandex-client-id",
        ClientSecret: "your-yandex-client-secret",
        RedirectURL:  "http://localhost:8080/auth/yandex/callback",
        Scopes:       []string{"login:email", "login:info", "login:avatar"},
        Endpoint:     yandex.Endpoint,
    }
    
    VKOAuthConfig = &oauth2.Config{
        ClientID:     "your-vk-client-id",
        ClientSecret: "your-vk-client-secret",
        RedirectURL:  "http://localhost:8080/auth/vk/callback",
        Scopes:       []string{"email", "photos"},
        Endpoint:     vk.Endpoint,
    }
)