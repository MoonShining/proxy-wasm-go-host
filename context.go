/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package proxywasm

const ProxyWasmABI_0_1_0 string = "proxy_abi_version_0_1_0"

type ABIContext struct {
	Imports  ImportsHandler
	Instance WasmInstance
}

func (a *ABIContext) Name() string {
	return ProxyWasmABI_0_1_0
}

func (a *ABIContext) GetExports() Exports {
	return a
}

func (a *ABIContext) SetImports(imports ImportsHandler) {
	a.Imports = imports
}

func (a *ABIContext) SetInstance(instance WasmInstance) {
	a.Instance = instance
}