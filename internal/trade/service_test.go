package trade

import (
	"os"
	"testing"
)

func Test_tradeManager_isFairTrade(t *testing.T) {
	type args struct {
		indexOne int
		indexTwo int
	}
	tests := []struct {
		name string
		tm   tradeManager
		args args
		want bool
	}{
		{
			name: "[T0] fair trade",
			tm:   tradeManager{},
			args: args{
				indexOne: 60,
				indexTwo: 50,
			},
			want: true,
		},
		{
			name: "[T1] fair trade",
			tm:   tradeManager{},
			args: args{
				indexOne: 45,
				indexTwo: 50,
			},
			want: true,
		},
		{
			name: "[T2] unfair trade",
			tm:   tradeManager{},
			args: args{
				indexOne: 140,
				indexTwo: 50,
			},
			want: false,
		},
	}
	os.Setenv("FAIR_INDEX", "20")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := tradeManager{}
			if got := tm.isFairTrade(tt.args.indexOne, tt.args.indexTwo); got != tt.want {
				t.Errorf("tradeManager.isFairTrade() = %v, want %v", got, tt.want)
			}
		})
	}
}
