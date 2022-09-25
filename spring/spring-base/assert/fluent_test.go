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

package assert_test

import (
	"testing"

	"github.com/go-spring/spring-base/assert"
)

func TestFluentTrue(t *testing.T) {

	Case(t, func(g *assert.MockT) {
		assert.That(g, true).IsTrue()
	})

	Case(t, func(g *assert.MockT) {
		g.EXPECT().Error([]interface{}{"got false but expect true"})
		assert.That(g, false).IsTrue()
	})

	Case(t, func(g *assert.MockT) {
		g.EXPECT().Error([]interface{}{"got false but expect true; param (index=0)"})
		assert.Bool(g, false).IsTrue("param (index=0)")
	})
}

func TestFluentHasPrefix(t *testing.T) {

	Case(t, func(g *assert.MockT) {
		assert.That(g, "hello, world!").HasPrefix("hello")
	})

	Case(t, func(g *assert.MockT) {
		g.EXPECT().Error([]interface{}{"'hello, world!' doesn't have prefix 'xxx'"})
		assert.That(g, "hello, world!").HasPrefix("xxx")
	})

	Case(t, func(g *assert.MockT) {
		g.EXPECT().Error([]interface{}{"'hello, world!' doesn't have prefix 'xxx'; param (index=0)"})
		assert.String(g, "hello, world!").HasPrefix("xxx", "param (index=0)")
	})
}
