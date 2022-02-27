package template

var (
	ApiProto = `syntax = "proto3";

package api.{{.SCName}};

import "google/api/annotations.proto";

option go_package = "{{.ModPath}}/api/{{.SCName}};{{.SCName}}";
option java_multiple_files = true;
option java_package = "api.{{.SCName}}";

service {{.BCName}} {
	// Sends a greeting
	rpc SayHello (HelloRequest) returns (HelloReply)  {
		option (google.api.http) = {
			get: "/helloworld/{name}"
		};
	}
}

// The request message containing the user's name.
message HelloRequest {
	string name = 1;
}

// The response message containing the greetings
message HelloReply {
	string message = 1;
}`

)