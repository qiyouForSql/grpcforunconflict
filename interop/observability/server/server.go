/*
 *
 * Copyright 2022 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/qiyouForSql/grpcforunconflict/gcp/observability"
	"github.com/qiyouForSql/grpcforunconflict/interop"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	err := observability.Start(context.Background())
	if err != nil {
		log.Fatalf("observability start failed: %v", err)
	}
	defer observability.End()
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpcforunconflict.NewServer()
	defer server.Stop()
	testgrpcforunconflict.RegisterTestServiceServer(server, interop.NewTestServer())
	log.Printf("Observability interop server listening on %v", lis.Addr())
	server.Serve(lis)
}
