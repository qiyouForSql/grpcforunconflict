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

// Binary client is an example client to illustrate the use of the stats handler.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/qiyouForSql/grpcforunconflict/credentials/insecure"
	echogrpc "github.com/qiyouForSql/grpcforunconflict/examples/features/proto/echo"
	echopb "github.com/qiyouForSql/grpcforunconflict/examples/features/proto/echo"
	"github.com/qiyouForSql/grpcforunconflict/examples/features/stats_monitoring/statshandler"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	flag.Parse()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(statshandler.New()),
	}
	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("failed to connect to server %q: %v", *addr, err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := echogrpc.NewEchoClient(conn)

	resp, err := c.UnaryEcho(ctx, &echopb.EchoRequest{Message: "stats handler demo"})
	if err != nil {
		log.Fatalf("unexpected error from UnaryEcho: %v", err)
	}
	log.Printf("RPC response: %s", resp.Message)
}
