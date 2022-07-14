package health_check

import (
	"testing"
)

func TestHealthCheckService_Pong(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "when pong method is call should return Pong :)",
			want: "Pong :)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &HealthCheckService{}
			if got := this.Pong(); got != tt.want {
				t.Errorf("Pong() = %v, want %v", got, tt.want)
			}
		})
	}
}
