/*
 *
 * Copyright 2018 gRPC authors.
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

// Binary client is an example client.
package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	"github.com/qiyouForSql/grpcforunconflict/channelz/service"
	"github.com/qiyouForSql/grpcforunconflict/credentials/insecure"
	pb "github.com/qiyouForSql/grpcforunconflict/examples/helloworld/helloworld"
	"github.com/qiyouForSql/grpcforunconflict/resolver"
	"github.com/qiyouForSql/grpcforunconflict/resolver/manual"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	/***** Set up the server serving channelz service. *****/
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	s := grpcforunconflict.NewServer()
	service.RegisterChannelzServiceToServer(s)
	go s.Serve(lis)
	defer s.Stop()

	/***** Initialize manual resolver and Dial *****/
	r := manual.NewBuilderWithScheme("whatever")
	// Set up a connection to the server.
	conn, err := grpcforunconflict.Dial(r.Scheme()+":///test.server", grpcforunconflict.WithTransportCredentials(insecure.NewCredentials()), grpcforunconflict.WithResolvers(r), grpcforunconflict.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Manually provide resolved addresses for the target.
	r.UpdateState(resolver.State{Addresses: []resolver.Address{{Addr: ":10001"}, {Addr: ":10002"}, {Addr: ":10003"}}})

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.

	/***** Make 100 SayHello RPCs *****/
	for i := 0; i < 100; i++ {
		// Setting a 150ms timeout on the RPC.
		ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Printf("could not greet: %v", err)
		} else {
			log.Printf("Greeting: %s", r.Message)
		}
	}

	/***** Wait for user exiting the program *****/
	// Unless you exit the program (e.g. CTRL+C), channelz data will be available for querying.
	// Users can take time to examine and learn about the info provided by channelz.
	select {}
}
