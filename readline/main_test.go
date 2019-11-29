package main

import "testing"

func TestGetNewName(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{
			name: `normal test`,
			in:   "gcr.io/google-samples/node-hello:1.0",
			want: DeploymentName + "/gcr.io_google-samples_node-hello:1.0",
		},
		{
			name: `dockerhub default`,
			in:   "nginx",
			want: DeploymentName + "/hub.docker.com_library_nginx",
		},
		{
			name: `dockerhub with projectname`,
			in:   "lyft/envoy:1.2",
			want: DeploymentName + "/hub.docker.com_lyft_envoy:1.2",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNewName(tt.in)
			if got != tt.want {
				t.Errorf(`(%v) = %v; want "%v"`, tt.in, got, tt.want)
			}
		})
	}
}
