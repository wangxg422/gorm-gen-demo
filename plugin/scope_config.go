package plugin

import "context"

type scopeCtxKey struct{}

func WithScopeConfig(ctx context.Context, cfg ScopeConfig) context.Context {
    return context.WithValue(ctx, scopeCtxKey{}, cfg)
}

func GetScopeConfig(ctx context.Context) (ScopeConfig, bool) {
    v := ctx.Value(scopeCtxKey{})
    if v == nil {
        return ScopeConfig{}, false
    }
    return v.(ScopeConfig), true
}
