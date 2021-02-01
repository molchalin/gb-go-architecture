package main

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"unsafe"

	"github.com/google/uuid"
)

const (
	CtxRequestIDKey = "X-Request-Id"
)

func printCtxValuesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		printContextInternals(r.Context(), false)
		next.ServeHTTP(w, r)
	})
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := r.Context().Value(CtxRequestIDKey)
		if v == nil {
			v = uuid.New().String()
			ctxWithRID := context.WithValue(r.Context(), CtxRequestIDKey, v)
			r = r.WithContext(ctxWithRID)
		}
		next.ServeHTTP(w, r)
	})
}

func printContextInternals(ctx interface{}, inner bool) {
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				printContextInternals(reflectValue.Interface(), true)
			} else {
				fmt.Printf("field name: %+v\n", reflectField.Name)
				fmt.Printf("value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Printf("context is empty (int)\n")
	}
}
