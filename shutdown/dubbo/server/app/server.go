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

package main

import (
	"log"
	"net/http"
	"time"
)

import (
	"github.com/apache/dubbo-go-pixiu/pkg/common/constant"
)

func main() {
	http.HandleFunc("/com.dubbogo.pixiu.TripleUserService/TestByDubbo", testFunc)
	log.Println("Starting sample server ...")
	log.Fatal(http.ListenAndServe(":20001", nil))
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.Header().Set(constant.HeaderKeyContextType, constant.HeaderValueJsonUtf8)
	w.Write([]byte("{\"message\":\"receive\"}"))
}
