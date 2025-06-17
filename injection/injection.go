package injection

import (
	"context"
	"reflect"
)

const (
	ErrInstanceNotFound = "instance not found"
)

type Injection struct {
	instances          map[reflect.Type]*interface{}
	scopedInstances    map[reflect.Type]func() (*interface{}, error)
	transientInstances map[reflect.Type]func() (*interface{}, error)
}

func NewInjection() *Injection {
	return &Injection{
		instances:          make(map[reflect.Type]*interface{}),
		scopedInstances:    make(map[reflect.Type]func() (*interface{}, error)),
		transientInstances: make(map[reflect.Type]func() (*interface{}, error)),
	}
}

func (i *Injection) AddSingleton(obj *interface{}) {
	tpe := reflect.TypeOf(&obj)
	i.instances[tpe] = obj
}

func (i *Injection) AddScoped(obj *interface{}, scopedMaker func() (*interface{}, error)) {
	tpe := reflect.TypeOf(&obj)
	i.scopedInstances[tpe] = scopedMaker
}

func (i *Injection) AddTransient(obj *interface{}, transientMaker func() (*interface{}, error)) {
	tpe := reflect.TypeOf(&obj)
	i.transientInstances[tpe] = transientMaker
}

func (i *Injection) GetInstance(ctx context.Context, obj *interface{}) (*interface{}, context.Context, error) {
	tpe := reflect.TypeOf(&obj)
	exists := ctx.Value(tpe.String())
	if exists != nil {
		return &exists, ctx, nil
	}
	if instance, ok := i.instances[tpe]; ok {
		return instance, ctx, nil
	}
	if scopedMaker, ok := i.scopedInstances[tpe]; ok {
		instance, err := scopedMaker()
		if err != nil {
			return nil, ctx, err
		}
		ctx = context.WithValue(ctx, tpe.String(), instance)
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
