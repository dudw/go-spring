/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package log_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/go-spring/spring-base/log"
)

func TestDefault(t *testing.T) {

	log.SetLevel(log.TraceLevel)
	defer log.Reset()

	log.Trace("a", "=", "1")
	log.Tracef("a=%d", 1)

	log.Trace(func() []interface{} { return log.T("a", "=", "1") })
	log.Tracef("a=%d", func() []interface{} { return log.T(1) })

	log.Debug("a", "=", "1")
	log.Debugf("a=%d", 1)

	log.Debug(func() []interface{} { return log.T("a", "=", "1") })
	log.Debugf("a=%d", func() []interface{} { return log.T(1) })

	log.Info("a", "=", "1")
	log.Infof("a=%d", 1)

	log.Info(func() []interface{} { return log.T("a", "=", "1") })
	log.Infof("a=%d", func() []interface{} { return log.T(1) })

	log.Warn("a", "=", "1")
	log.Warnf("a=%d", 1)

	log.Warn(func() []interface{} { return log.T("a", "=", "1") })
	log.Warnf("a=%d", func() []interface{} { return log.T(1) })

	log.Error("a", "=", "1")
	log.Errorf("a=%d", 1)

	log.Error(func() []interface{} { return log.T("a", "=", "1") })
	log.Errorf("a=%d", func() []interface{} { return log.T(1) })

	t.Run("panic#00", func(t *testing.T) {
		defer func() { fmt.Println(recover()) }()
		log.Panic("error")
	})

	t.Run("panic#01", func(t *testing.T) {
		defer func() { fmt.Println(recover()) }()
		log.Panic(errors.New("error"))
	})

	t.Run("panic#02", func(t *testing.T) {
		defer func() { fmt.Println(recover()) }()
		log.Panicf("error:%d", 404)
	})

	log.Fatal("a", "=", "1")
	log.Fatalf("a=%d", 1)
}

type traceIDKeyType int

var traceIDKey traceIDKeyType

func myOutput(level log.Level, e *log.Entry) {

	msg := e.GetMsg()
	tag := e.GetTag()
	if len(tag) > 0 {
		tag += " "
	}

	strCtx := func(ctx context.Context) string {
		if ctx == nil {
			return ""
		}
		v := ctx.Value(traceIDKey)
		if v == nil {
			return ""
		}
		return "trace_id:" + v.(string)
	}(e.GetCtx())

	line := e.GetLine()
	file := e.GetFile()
	strLevel := strings.ToUpper(level.String())
	fmt.Printf("[%s] %s:%d %s %s%s\n", strLevel, file, line, strCtx, tag, msg)
}

func TestEntry(t *testing.T) {
	ctx := context.WithValue(context.TODO(), traceIDKey, "0689")

	log.SetLevel(log.TraceLevel)
	log.SetOutput(myOutput)
	defer log.Reset()

	logger := log.Ctx(ctx)
	logger.Trace("level:", "trace")
	logger.Tracef("level:%s", "trace")
	logger.Debug("level:", "debug")
	logger.Debugf("level:%s", "debug")
	logger.Info("level:", "info")
	logger.Infof("level:%s", "info")
	logger.Warn("level:", "warn")
	logger.Warnf("level:%s", "warn")
	logger.Error("level:", "error")
	logger.Errorf("level:%s", "error")
	logger.Panic("level:", "panic")
	logger.Panicf("level:%s", "panic")
	logger.Fatal("level:", "fatal")
	logger.Fatalf("level:%s", "fatal")

	logger.Trace(func() []interface{} { return log.T("level:", "trace") })
	logger.Tracef("level:%s", func() []interface{} { return log.T("trace") })
	logger.Debug(func() []interface{} { return log.T("level:", "debug") })
	logger.Debugf("level:%s", func() []interface{} { return log.T("debug") })
	logger.Info(func() []interface{} { return log.T("level:", "info") })
	logger.Infof("level:%s", func() []interface{} { return log.T("info") })
	logger.Warn(func() []interface{} { return log.T("level:", "warn") })
	logger.Warnf("level:%s", func() []interface{} { return log.T("warn") })
	logger.Error(func() []interface{} { return log.T("level:", "error") })
	logger.Errorf("level:%s", func() []interface{} { return log.T("error") })

	logger = logger.Tag("__in")
	logger.Trace("level:", "trace")
	logger.Tracef("level:%s", "trace")
	logger.Debug("level:", "debug")
	logger.Debugf("level:%s", "debug")
	logger.Info("level:", "info")
	logger.Infof("level:%s", "info")
	logger.Warn("level:", "warn")
	logger.Warnf("level:%s", "warn")
	logger.Error("level:", "error")
	logger.Errorf("level:%s", "error")
	logger.Panic("level:", "panic")
	logger.Panicf("level:%s", "panic")
	logger.Fatal("level:", "fatal")
	logger.Fatalf("level:%s", "fatal")

	logger = log.Tag("__in")
	logger.Ctx(ctx).Trace("level:", "trace")
	logger.Ctx(ctx).Tracef("level:%s", "trace")
	logger.Ctx(ctx).Debug("level:", "debug")
	logger.Ctx(ctx).Debugf("level:%s", "debug")
	logger.Ctx(ctx).Info("level:", "info")
	logger.Ctx(ctx).Infof("level:%s", "info")
	logger.Ctx(ctx).Warn("level:", "warn")
	logger.Ctx(ctx).Warnf("level:%s", "warn")
	logger.Ctx(ctx).Error("level:", "error")
	logger.Ctx(ctx).Errorf("level:%s", "error")
	logger.Ctx(ctx).Panic("level:", "panic")
	logger.Ctx(ctx).Panicf("level:%s", "panic")
	logger.Ctx(ctx).Fatal("level:", "fatal")
	logger.Ctx(ctx).Fatalf("level:%s", "fatal")
}
