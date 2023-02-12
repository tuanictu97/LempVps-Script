//go:build wireinject
// +build wireinject

package inject

// The build tag makes sure the stub is not built in the final build.

import (
	"context"

	"github.com/google/wire"
	"io.github.tuanictu97/api/internal/module/rbac"
	"io.github.tuanictu97/api/internal/module/sys"
	"io.github.tuanictu97/api/internal/x/utilx"
) // end

func BuildInjector(ctx context.Context) (*Injector, func(), error) {
	wire.Build(
		InitAuth,
		InitCache,
		InitDB,
		wire.NewSet(wire.Struct(new(Injector), "*")),
		utilx.TransRepoSet,
		rbac.Set,
		sys.Set,
	) // end
	return new(Injector), nil, nil
}
