package injection

import (
	"context"
	"reflect"

	"gofr.dev/pkg/gofr"
)

const (
	ErrInstanceNotFound = "instance not found"
)

type Injection struct {
	instances          map[reflect.Type]BaseModel // Change from *BaseModel to BaseModel
	scopedInstances    map[reflect.Type]func() (BaseModel, error)
	transientInstances map[reflect.Type]func() (BaseModel, error)
}

type BaseModel interface {
}

func NewInjection() *Injection {
	return &Injection{
		instances:          make(map[reflect.Type]BaseModel),
		scopedInstances:    make(map[reflect.Type]func() (BaseModel, error)),
		transientInstances: make(map[reflect.Type]func() (BaseModel, error)),
	}
}

func (i *Injection) AddSingleton(obj BaseModel) {
	tpe := reflect.TypeOf(obj)
	i.instances[tpe] = obj // Store the BaseModel interface directly, not a pointer to it
}

func (i *Injection) AddScoped(obj BaseModel, scopedMaker func() (BaseModel, error)) {
	tpe := reflect.TypeOf(obj)
	i.scopedInstances[tpe] = scopedMaker
}

func (i *Injection) AddTransient(obj BaseModel, transientMaker func() (BaseModel, error)) {
	tpe := reflect.TypeOf(obj)
	i.transientInstances[tpe] = transientMaker
}

func (i *Injection) GetInstance(ctx *gofr.Context, obj BaseModel) (any, context.Context, error) {
	tpe := reflect.TypeOf(obj)
	exists := ctx.Value(tpe.String())
	if exists != nil {
		return exists, ctx, nil
	}
	if instance, ok := i.instances[tpe]; ok {
		return instance, ctx, nil
	}
	if scopedMaker, ok := i.scopedInstances[tpe]; ok {
		instance, err := scopedMaker()
		if err != nil {
			return nil, ctx, err
		}
		ctx.Context = context.WithValue(ctx.Context, tpe.String(), instance)
		return instance, ctx, nil
	}
	if transientMaker, ok := i.transientInstances[tpe]; ok {
		instance, err := transientMaker()
		if err != nil {
			return nil, ctx, err
		}
		newContext := context.Background()
		newContext = context.WithValue(newContext, tpe.String(), instance)
		return instance, newContext, nil
	}
	return nil, ctx, &InjectionError{
		Message:      ErrInstanceNotFound,
		InjectedType: tpe,
	}
}
