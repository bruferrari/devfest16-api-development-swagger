package app

import (
    mcache "github.com/go-macaron/cache"
    "github.com/go-macaron/gzip"
    "github.com/go-macaron/i18n"
    "github.com/go-macaron/jade"
    "github.com/go-macaron/session"
    "github.com/bruferrari/devfest16-api-development-swagger/conf"
    "github.com/bruferrari/devfest16-api-development-swagger/lib/cache"
    "github.com/bruferrari/devfest16-api-development-swagger/lib/context"
    "github.com/bruferrari/devfest16-api-development-swagger/lib/template"
    "github.com/bruferrari/devfest16-api-development-swagger/handler"
    "gopkg.in/macaron.v1"
)

func SetupMiddlewares(app *macaron.Macaron) {
    app.Use(macaron.Logger())
    app.Use(macaron.Recovery())
    app.Use(gzip.Gziper())
    app.Use(macaron.Static("public"))
    app.Use(i18n.I18n(i18n.Options{
    Directory: "locale",
    Langs:     []string{"pt-BR", "en-US"},
    Names:     []string{"Português do Brasil", "American English"},
    }))
    app.Use(jade.Renderer(jade.Options{
    Directory: "public/templates",
    Funcs:     template.FuncMaps(),
    }))
    app.Use(macaron.Renderer())
    app.Use(mcache.Cacher(
    cache.Option(conf.Cfg.Section("").Key("cache_adapter").Value()),
    ))
    app.Use(session.Sessioner())
    app.Use(context.Contexter())
}

func SetupRoutes(app *macaron.Macaron) {
    
        route(app, "DELETE", "/v1/media/{media-id}/comments", handler.MediaMediaIdCommentsDelete)
    
        route(app, "GET", "/v1/media/{media-id}/comments", handler.MediaMediaIdCommentsGet)
    
        route(app, "POST", "/v1/media/{media-id}/comments", handler.MediaMediaIdCommentsPost)
    
        route(app, "GET", "/v1/geographies/{geo-id}/media/recent", handler.GeographiesGeoIdMediaRecentGet)
    
        route(app, "DELETE", "/v1/media/{media-id}/likes", handler.MediaMediaIdLikesDelete)
    
        route(app, "GET", "/v1/media/{media-id}/likes", handler.MediaMediaIdLikesGet)
    
        route(app, "POST", "/v1/media/{media-id}/likes", handler.MediaMediaIdLikesPost)
    
        route(app, "GET", "/v1/locations/{location-id}", handler.LocationsLocationIdGet)
    
        route(app, "GET", "/v1/locations/{location-id}/media/recent", handler.LocationsLocationIdMediaRecentGet)
    
        route(app, "GET", "/v1/locations/search", handler.LocationsSearchGet)
    
        route(app, "GET", "/v1/locations/{location-id}/media/recent", handler.LocationsLocationIdMediaRecentGet)
    
        route(app, "GET", "/v1/media1/{shortcode}", handler.Media1ShortcodeGet)
    
        route(app, "POST", "/v1/media/{media-id}/comments", handler.MediaMediaIdCommentsPost)
    
        route(app, "GET", "/v1/media/{media-id}", handler.MediaMediaIdGet)
    
        route(app, "GET", "/v1/media/{media-id}/likes", handler.MediaMediaIdLikesGet)
    
        route(app, "GET", "/v1/media/popular", handler.MediaPopularGet)
    
        route(app, "GET", "/v1/media/search", handler.MediaSearchGet)
    
        route(app, "GET", "/v1/users/self/requested-by", handler.UsersSelfRequestedByGet)
    
        route(app, "GET", "/v1/users/{user-id}/followed-by", handler.UsersUserIdFollowedByGet)
    
        route(app, "GET", "/v1/users/{user-id}/follows", handler.UsersUserIdFollowsGet)
    
        route(app, "POST", "/v1/users/{user-id}/relationship", handler.UsersUserIdRelationshipPost)
    
        route(app, "GET", "/v1/tags/search", handler.TagsSearchGet)
    
        route(app, "GET", "/v1/tags/{tag-name}", handler.TagsTagNameGet)
    
        route(app, "GET", "/v1/tags/{tag-name}/media/recent", handler.TagsTagNameMediaRecentGet)
    
        route(app, "GET", "/v1/users/search", handler.UsersSearchGet)
    
        route(app, "GET", "/v1/users/self/feed", handler.UsersSelfFeedGet)
    
        route(app, "GET", "/v1/users/self/media/liked", handler.UsersSelfMediaLikedGet)
    
        route(app, "GET", "/v1/users/{user-id}", handler.UsersUserIdGet)
    
        route(app, "GET", "/v1/users/{user-id}/media/recent", handler.UsersUserIdMediaRecentGet)
    
}

func route(app *macaron.Macaron, method, path string, handler ...macaron.Handler) {
    switch method {
        case "GET":
        app.Get(path, handler...)
        case "POST":
        app.Post(path, handler...)
        case "PUT":
        app.Put(path, handler...)
        case "DELETE":
        app.Delete(path, handler...)
        case "HEAD":
        app.Head(path, handler...)
        case "OPTIONS":
        app.Options(path, handler...)
        case "PATCH":
        app.Patch(path, handler...)
    }
}